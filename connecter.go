package gql

import (
	"context"
	"os"
	"sync"
)

// Connecter represents a minimal set of functions to run a
// project using GraphQL. Connects and Disconnects, and run requests.
type Connecter interface {
	Connect(...ClientOption)
	ConnectAt(string, ...ClientOption)
	Disconnect()
	Run(*Request) (Response, error)
	RunWithSpecialResp(*Request, interface{}) error
}

// gql is a Connecter that functions with a real GraphQL
// connection.
type gql struct {
	once   sync.Once
	client *Client
}

// newConnecter returns a GraphQL connecter for production purposes. It
// connects to a real GraphQL engine to perform operations of search and
// manipulation of data.
func newConnecter() (g Connecter) {
	g = &gql{}
	return
}

// Connect to GraphQL engine. It tries to connect with GRAPHQL_URL, but
// without defining this environment variable, it tries to connect with
// default URL.
func (g *gql) Connect(opts ...ClientOption) {
	g.once.Do(func() {
		// Parse adequate gql URI.
		u := g.gqlURI()

		g.client = NewClient(u, opts...)
	})
}

// ConnectAt defined URL to GraphQL engine.
func (g *gql) ConnectAt(u string, opts ...ClientOption) {
	g.once.Do(func() {
		g.client = NewClient(u, opts...)
	})
}

// Disconnect undo the connection made. Preparing package for a new
// connection.
func (g *gql) Disconnect() {
	g.once = *new(sync.Once)

	g.client = nil
}

// Run executes the query and returns the response from the
// data field. If the request fails or the server returns an
// error, the first error will be returned.
func (g *gql) Run(req *Request) (resp Response, err error) {
	resp, err = g.client.Run(context.Background(), req)
	return
}

// RunWithSpecialResp executes the query and unmarshals the response from the
// data field into the response object.
// Pass in a nil response object to skip response parsing.
// If the request fails or the server returns an error, the first error
// will be returned.
func (g *gql) RunWithSpecialResp(req *Request, resp interface{}) (err error) {
	err = g.client.RunWithSpecialResp(context.Background(), req, resp)
	return
}

const (
	// DBUrl is the default GraphQL url that will be used to
	// connect to the engine.
	DBUrl = "http://192.168.99.100:8080/console/v1/graphql"
)

// gqlURI load the selected gql interface url, or default.
func (g *gql) gqlURI() (u string) {
	u = os.Getenv("GRAPHQL_URL")
	if len(u) == 0 {
		u = DBUrl
	}
	return
}
