package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Khan/genqlient/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"io"
	"net/http"
	"sync"
	"time"
)

// User Facing API

type Config struct {
	AccessToken string
	ShopUrl     string
	Version     string
	Retries     int // Made public for configuration
}

// Add custom response type to handle extensions and errors
type shopifyResponse struct {
	Data       interface{}       `json:"data"`
	Extensions graphQLExtensions `json:"extensions,omitempty"`
	Errors     []graphQLError    `json:"errors,omitempty"`
}

type shopifyErrorList []graphQLError

func (e shopifyErrorList) Error() string {
	if len(e) == 0 {
		return "no errors"
	}
	return e[0].Message
}

func (e shopifyErrorList) Extensions() []map[string]interface{} {
	extensions := make([]map[string]interface{}, len(e))
	for i, err := range e {
		if err.Extensions != nil {
			extensions[i] = map[string]interface{}{
				"code":          err.Extensions.Code,
				"documentation": err.Extensions.Documentation,
			}
		}
	}
	return extensions
}

type ClientManager struct {
	clients map[string]Client
	mutex   sync.RWMutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[string]Client),
	}
}

func generateClientKey(config Config) string {
	return fmt.Sprintf("%s:%s:%s", config.ShopUrl, config.Version, config.AccessToken)
}

func (manager *ClientManager) GetClient(config Config) Client {
	key := generateClientKey(config)

	manager.mutex.RLock()
	if client, exists := manager.clients[key]; exists {
		manager.mutex.RUnlock()
		return client
	}
	manager.mutex.RUnlock()

	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	// Double-check pattern in case another goroutine created the client
	if client, exists := manager.clients[key]; exists {
		return client
	}

	endpoint := fmt.Sprintf("https://%s.myshopify.com/admin/api/%v/graphql.json", config.ShopUrl, config.Version)
	httpClient := http.Client{
		Transport: &headerTransport{
			headers: map[string]string{
				"X-Shopify-Access-Token": config.AccessToken,
				"Content-Type":           "application/json",
			},
		},
	}

	client := &client{
		httpClient: &httpClient,
		endpoint:   endpoint,
		retries:    config.Retries,
	}

	manager.clients[key] = client
	return client
}

type Client interface {
	MakeRequest(
		ctx context.Context,
		req *graphql.Request,
		resp *graphql.Response,
	) error
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type client struct {
	httpClient Doer
	endpoint   string
	retries    int
}

func newClient(endpoint string, httpClient Doer, retries int) Client {
	if httpClient == nil || httpClient == (*http.Client)(nil) {
		httpClient = http.DefaultClient
	}
	return &client{httpClient, endpoint, retries}
}

// custom header transport to allow shopify access token
type headerTransport struct {
	headers map[string]string
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range t.headers {
		req.Header.Set(key, value)
	}
	return http.DefaultTransport.RoundTrip(req)
}

func convertToGQLErrors(errors []graphQLError) gqlerror.List {
	var gqlErrors gqlerror.List
	for _, err := range errors {
		gqlError := &gqlerror.Error{
			Message:   err.Message,
			Locations: []gqlerror.Location{},
		}
		if err.Extensions != nil {
			gqlError.Extensions = map[string]interface{}{
				"code":          err.Extensions.Code,
				"documentation": err.Extensions.Documentation,
			}
		}
		for _, loc := range err.Locations {
			gqlError.Locations = append(gqlError.Locations, gqlerror.Location{
				Line:   loc.Line,
				Column: loc.Column,
			})
		}
		gqlErrors = append(gqlErrors, gqlError)
	}
	return gqlErrors
}

func (c *client) MakeRequest(ctx context.Context, req *graphql.Request, resp *graphql.Response) error {
	attempts := 0
	maxRetries := c.retries
	if maxRetries == 0 {
		maxRetries = 3 // Default retries if not specified
	}

	for {
		var shopifyResp struct {
			Data       interface{}       `json:"data"`
			Extensions graphQLExtensions `json:"extensions,omitempty"`
			Errors     []graphQLError    `json:"errors,omitempty"`
		}
		shopifyResp.Data = resp.Data

		httpReq, err := c.createPostRequest(req)
		if err != nil {
			return err
		}

		if ctx != nil {
			httpReq = httpReq.WithContext(ctx)
		}

		httpResp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}

		bodyBytes, err := io.ReadAll(httpResp.Body)
		httpResp.Body.Close()
		if err != nil {
			return err
		}

		if httpResp.StatusCode != http.StatusOK {
			if attempts >= maxRetries {
				return fmt.Errorf("returned error %v: %s", httpResp.Status, bodyBytes)
			}
			attempts++
			continue
		}

		err = json.Unmarshal(bodyBytes, &shopifyResp)
		if err != nil {
			return err
		}

		// Check for throttling errors
		for _, gqlErr := range shopifyResp.Errors {
			if gqlErr.Extensions != nil && gqlErr.Extensions.Code == graphQLErrorCodeThrottled {
				if attempts >= maxRetries {
					return fmt.Errorf("max retries reached: %v", gqlErr.Message)
				}

				// Calculate retry delay based on cost
				retryAfter := shopifyResp.Extensions.Cost.RetryAfterSeconds()
				if retryAfter > 0 {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case <-time.After(time.Duration(retryAfter) * time.Second):
					}
				}

				attempts++
				continue
			}
		}

		// Convert and set errors
		resp.Errors = convertToGQLErrors(shopifyResp.Errors)

		// If we got here without errors or after handling throttling, break the loop
		break
	}

	return nil
}

func (c *client) createPostRequest(req *graphql.Request) (*http.Request, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(
		http.MethodPost,
		c.endpoint,
		bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	return httpReq, nil
}

// extensions - thanks Bold Commerce (https://github.com/bold-commerce/go-shopify)

type graphQLExtensions struct {
	Cost GraphQLCost `json:"cost"`
}

// GraphQLCost represents the cost of the graphql query
type GraphQLCost struct {
	RequestedQueryCost int                   `json:"requestedQueryCost"`
	ActualQueryCost    *int                  `json:"actualQueryCost"`
	ThrottleStatus     GraphQLThrottleStatus `json:"throttleStatus"`
}

// GraphQLThrottleStatus represents the status of the shop's rate limit points
type GraphQLThrottleStatus struct {
	MaximumAvailable   float64 `json:"maximumAvailable"`
	CurrentlyAvailable float64 `json:"currentlyAvailable"`
	RestoreRate        float64 `json:"restoreRate"`
}

type graphQLError struct {
	Message    string                  `json:"message"`
	Extensions *graphQLErrorExtensions `json:"extensions"`
	Locations  []graphQLErrorLocation  `json:"locations"`
}

type graphQLErrorExtensions struct {
	Code          string
	Documentation string
}

const (
	graphQLErrorCodeThrottled = "THROTTLED"
)

type graphQLErrorLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// RetryAfterSeconds returns the estimated retry after seconds based on
// the requested query cost and throttle status
func (c GraphQLCost) RetryAfterSeconds() float64 {
	var diff float64

	if c.ActualQueryCost != nil {
		diff = c.ThrottleStatus.CurrentlyAvailable - float64(*c.ActualQueryCost)
	} else {
		diff = c.ThrottleStatus.CurrentlyAvailable - float64(c.RequestedQueryCost)
	}

	if diff < 0 {
		return -diff / c.ThrottleStatus.RestoreRate
	}

	return 0
}
