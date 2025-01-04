# shopify-go-agql
Generate a type-safe Graphql API client for the [Shopify Admin Api](https://shopify.dev/docs/api/admin-graphql) from your existing queries and mutations in Go.

## Why This Project Exists
There are a few Go client options for Shopify, notably [bold-commerce/go-shopify](https://github.com/bold-commerce/go-shopify) which is listed in the Shopify documentation. Generally these implementations tend to abstract implementation details into models, a strategy which, although convenient, has the disadvantage of constantly having to catch up to the rapidly changing Shopify models. If you prefer a model-based approach that abstracts the implementation details, and don't need any new release features, you should use the Bold Commerce libary.

This project is a minimalistic, direct implementation of the Shopify Graphql Admin API. Instead of providing dedicated models for resources like Product and Order, it utilizes [genqlient](https://github.com/Khan/genqlient) to automatically generate a type-safe library using the queries and mutations you provide.

## What This Project Actually Does
So yeah. genqlient is doing most of the heavy lifting here. This project provides the following to save you a day or two.

- Some hardcoded enums with conversion methods
- Script to convert the instrospection .json response into a format that helps genqlient with inference
- In-client rate limiting and retries.

The rest is up to you. 

## Setup
### Config Files
Refer to /example in this repo for initial configuration requirements.

Add these files to your root directory, or place them in a package to keep your top level clean.

- /graphql contains your queries and mutations
- .env is required for introspection. SHOPIFY_API_KEY is your public key
- genqlient.yaml is a pre-filled configuration file. If you are running from root directory, no changes are needed. However, if running from inside a package, change `package: agql` to `package: your-package-name` and instead of `generated: agql/genqlient.go` use `generated: your-package-name.go`

### Scripts

Run these commands in order from the directory containing your genqlient.yaml file

- go run github.com/sadsciencee/shopify-client-go@latest introspect
- go run github.com/sadsciencee/shopify-client-go@latest gen

After running `introspect` you should see a schema.graphql file in your directory.

After running `gen` you should see the generated client file or directory.

## Usage
### Initialize Client
Use `clientManager := NewManager()` to initialize a client pool.

This library does not provide session auth or management. So to create a new client you will need to have already completed oAuth.

```go
ApiVersion := "2025-01"
clientManager := NewManager()
clientConfig := Config{
    AccessToken: auth.ShopifyAccessToken,
    ShopUrl:     auth.ShopUrl,
    Version:     ApiVersion,
}
```

### Use Client
Here is an example function that takes `client` from `github.com/sadsciencee/shopify-client-go/client` and runs a metafield definition migration from one shop to another.
In this case you would initialize a source client and dest client

```go
import (
"context"
"fmt"
"your-root-package/agql"
"github.com/sadsciencee/shopify-client-go/client"
)

func migrateMetafieldDefinitions(ctx context.Context, source client.Client, dest client.Client, ownerType agql.MetafieldOwnerType) {
    var waitGroup sync.WaitGroup
    
    response, err := agql.GetMetafieldDefinitions(ctx, source, ownerType, ptr(250))
    if err != nil {
        log.Fatal(err)
    }
    
    responseLength := len(response.MetafieldDefinitions.Edges)
    fmt.Printf("Total metafield definitions: %d\n", responseLength)
    
    for _, edge := range response.MetafieldDefinitions.Edges {
        waitGroup.Add(1)
        node := edge.Node
        
        definition := agql.MetafieldDefinitionInput{
            Namespace:   &node.Namespace,
            Key:         node.Key,
            Description: node.Description,
            Name:        node.Name,
            OwnerType:   node.OwnerType,
            Type:        node.Type.Name,
            Access: &agql.MetafieldAccessInput{
                Admin:           nil,
                Storefront:      node.Access.Storefront.ToInput(),
                CustomerAccount: node.Access.CustomerAccount.ToInput(),
            },
            Capabilities: &agql.MetafieldCapabilityCreateInput{
                SmartCollectionCondition: &agql.MetafieldCapabilitySmartCollectionConditionInput{
                    Enabled: node.Capabilities.SmartCollectionCondition.Enabled,
                },
                AdminFilterable: &agql.MetafieldCapabilityAdminFilterableInput{
                    Enabled: node.Capabilities.AdminFilterable.Enabled,
                },
            },
        }
    
        definitionCopy := definition
        
        go func(definitionInput agql.MetafieldDefinitionInput) {
            defer waitGroup.Done()
        
            result, error := agql.CreateMetafieldDefinition(ctx, dest, &definitionInput)
            if error != nil {
                log.Printf("Error creating metafield definition: %v", error)
                return
            }
        
            if result != nil && result.MetafieldDefinitionCreate != nil && len(result.MetafieldDefinitionCreate.UserErrors) > 0 {
                fmt.Println("------------------------------------------------")
                log.Println("Error creating metafield definition")
                
                for _, userError := range result.MetafieldDefinitionCreate.UserErrors {
                    fmt.Println(userError.Code, userError.Message, userError.Field)
                }
                
                fmt.Println("------------------------------------------------")
            }
        }(definitionCopy)
    }

    waitGroup.Wait()
    fmt.Printf("Metafield definitions migration complete for owner type: %s\n", ownerType)
}
```
## genqlient
For more information configuring your genqlient.yaml file, refer to the [genqlient docs](https://github.com/Khan/genqlient/blob/main/docs/genqlient.yaml).

## How To Contribute
The types provided in this library only exist to improve inference and developer experience. However, it is my goal to eventually have all enums and mutation inputs hardcoded in this library to reduce boilerplate in user code.

If you add enums or input structs to your project, and feel like paying it forward, feel free to send a PR with the new types and I'll merge it.

You can hook up your own types for anything that's not built into the package using genqlient bindings. genqlient will do a good job generating the structs and enums, so you can pretty much just copy paste the enums into a type module and add the .ToInput() and .ToValue() conversion methods. Copilot one-shots those methods every time, but sometimes input enums don't match query response enums because Shopify leaves deprecated options in the schema for query responses longer than mutation inputs.



## Contact
david@ucoastweb.com
