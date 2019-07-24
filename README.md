# gql [![GoDoc](https://godoc.org/github.com/ddsgok/gql?status.png)](http://godoc.org/github.com/ddsgok/gql) [![Build Status](https://travis-ci.org/ddsgok/gql.svg?branch=master)](https://travis-ci.org/ddsgok/gql) [![Go Report Card](https://goreportcard.com/badge/github.com/ddsgok/gql)](https://goreportcard.com/report/github.com/ddsgok/gql)

Low-level GraphQL client for Go.

* Simple, familiar API
* Respects `context.Context` timeouts and cancellation
* Build and execute any kind of GraphQL request
* Use strong Go types for response data
* Use variables and upload files
* Simple error handling

## Installation

Make sure you have a working Go environment. To install gql, simply run:

```bash
go get github.com/ddsgok/gql
```

## Usage

```go
import "context"

// create a client (safe to share across requests)
client := gql.NewClient("https://machinebox.io/gql")

// make a request
req := gql.NewRequest(`
    query ($key: String!) {
        items (id:$key) {
            field1
            field2
            field3
        }
    }
`)

// set any variables
req.Var("key", "value")

// set header fields
req.Header.Set("Cache-Control", "no-cache")

// define a Context for the request
ctx := context.Background()

// run it and capture the response
var respData ResponseStruct
if err := client.Run(ctx, req, &respData); err != nil {
    log.Fatal(err)
}
```

Read [graphql.queries](https://graphql.org/learn/queries/#variables) to understand how to
use variables.

### File support via multipart form data

By default, the package will send a JSON body. To enable the sending of files, you can opt to
use multipart form data instead using the `UseMultipartForm` option when you create your `Client`:

```go
client := gql.NewClient("https://machinebox.io/gql", gql.UseMultipartForm())
```

For more information, [read the godoc package documentation](http://godoc.org/github.com/ddsgok/gql) or the [blog post](https://blog.machinebox.io/a-graphql-client-library-for-go-5bffd0455878).

## Thanks

Thanks to [Chris Broadfoot](https://github.com/broady) for design help.
