package gql

var (
	conn = newConnecter()
)

// Connect to GraphQL engine. It tries to connect with GRAPHQL_URL, but
// without defining this environment variable, it tries to connect with
// default URL.
func Connect(opts ...ClientOption) {
	conn.Connect(opts...)
}

// ConnectAt defined URL to GraphQL engine.
func ConnectAt(u string, opts ...ClientOption) {
	conn.ConnectAt(u, opts...)
}

// Disconnect undo the connection made. Preparing package for a new
// connection.
func Disconnect() {
	conn.Disconnect()
}

// Run executes the query and returns the response from the
// data field. If the request fails or the server returns an
// error, the first error will be returned.
func Run(req *Request) (resp Response, err error) {
	resp, err = conn.Run(req)
	return
}

// RunWithSpecialResp executes the query and unmarshals the response from the
// data field into the response object.
// Pass in a nil response object to skip response parsing.
// If the request fails or the server returns an error, the first error
// will be returned.
func RunWithSpecialResp(req *Request, resp interface{}) (err error) {
	err = conn.RunWithSpecialResp(req, resp)
	return
}
