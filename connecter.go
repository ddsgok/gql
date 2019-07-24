package gql

import (
	"context"
	"os"
	"sync"
)

type Connecter interface {
	Connect(...ClientOption)
	ConnectAt(string, ...ClientOption)
	Disconnect()
	Run(*Request) (Response, error)
	RunWithSpecialResp(*Request, interface{}) error
}

type gql struct {
	once   sync.Once
	client *Client
}

func newConnecter() (g Connecter) {
	g = &gql{}
	return
}

func (g *gql) Connect(opts ...ClientOption) {
	g.once.Do(func() {
		// Parse adequate gql URI.
		u := g.gqlURI()

		g.client = NewClient(u, opts...)
	})
}

func (g *gql) ConnectAt(u string, opts ...ClientOption) {
	g.once.Do(func() {
		g.client = NewClient(u, opts...)
	})
}

func (g *gql) Disconnect() {
	g.once = *new(sync.Once)

	g.client = nil
}

func (g *gql) Run(req *Request) (resp Response, err error) {
	resp, err = g.client.Run(context.Background(), req)
	return
}

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
