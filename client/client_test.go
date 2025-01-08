package client

import (
	"context"
	"fmt"
	"github.com/sadsciencee/shopify-client-go/internal/agql"
	"github.com/sadsciencee/shopify-client-go/internal/config"
	"testing"
)

func getConfig() *config.Config {
	config.Load("../.env")
	return config.Get()
}

func getClient() Client {
	c := getConfig()
	clientManager := NewManager()
	clientConfig := Config{
		AccessToken: c.TestShopifyAccessToken,
		ShopUrl:     c.TestShopUrl,
		Version:     c.ApiVersion,
	}
	return clientManager.GetClient(clientConfig)
}

func TestClientManager(t *testing.T) {
	c := getConfig()
	clientManager := NewManager()
	clientConfig := Config{
		AccessToken: c.TestShopifyAccessToken,
		ShopUrl:     c.TestShopUrl,
		Version:     c.ApiVersion,
	}
	client := clientManager.GetClient(clientConfig)
	if client == nil {
		t.Fatal("failed to get client")
	}
	// make sure it works a second time (should pool)
	client = clientManager.GetClient(clientConfig)
	if client == nil {
		t.Fatal("failed to get client a second time")
	}
}

func TestRateLimiter(t *testing.T) {
	ctx := context.Background()
	client := getClient()
	if client == nil {
		t.Fatal("failed to get client")
	}
	const numRequests = 2000
	results := make(chan bool, numRequests)

	for i := 0; i < numRequests; i++ {
		go func(iteration int) {
			_, err := agql.GetMetafieldDefinitions(ctx, client, agql.MetafieldOwnerTypeProduct, ptr(250))
			if err != nil {
				fmt.Printf("Error on iteration %d: %v\n", iteration, err)
			}
			results <- err == nil
		}(i)
	}

	var successCount int

	for i := 0; i < numRequests; i++ {
		success := <-results
		if success {
			successCount++
		}
	}

	if successCount != 2000 {
		t.Fatal("failed to get client")
	}
}

func ptr[T any](v T) *T {
	return &v
}
