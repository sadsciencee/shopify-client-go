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

## How To Contribute
The types provided in this library only exist to improve inference and developer experience. However, it is my goal to eventually have all enums and mutation inputs hardcoded in this library to reduce boilerplate in user code.

If you add enums or input structs to your project, and feel like paying it forward, feel free to send a PR with the new types and I'll merge it.

You can hook up your own types for anything that's not built into the package using genqlient bindings. genqlient will do a good job generating the structs and enums, so you can pretty much just copy paste the enums into a type module and add the .ToInput() and .ToValue() conversion methods. Copilot one-shots those methods every time, but sometimes input enums don't match query response enums because Shopify leaves deprecated options in the schema for query responses longer than mutation inputs.



## Contact
david@ucoastweb.com
