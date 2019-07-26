package gql

import (
	"io"
	"net/http"
)

// Request is a GraphQL request.
type Request struct {
	q     string
	vars  map[string]interface{}
	files []File
	err   error

	// Header represent any request headers that will be set
	// when the request is made.
	Header http.Header
}

// NewRequest makes a new Request with the specified string.
func NewRequest(q string) *Request {
	req := &Request{
		q:      q,
		Header: make(map[string][]string),
	}
	return req
}

// AddHeader adds a new header to the Request.
func (req *Request) AddHeader(key, val string) *Request {
	req.Header.Add(key, val)
	return req
}

// GetHeader gets a header of the Request.
func (req *Request) GetHeader(key string) string {
	return req.Header.Get(key)
}

// SetHeader sets a header to the Request.
func (req *Request) SetHeader(key, val string) *Request {
	req.Header.Set(key, val)
	return req
}

// DelHeader deletes a header of the Request.
func (req *Request) DelHeader(key string) *Request {
	req.Header.Del(key)
	return req
}

// Var sets a variable.
func (req *Request) Var(key string, value interface{}) *Request {
	if req.vars == nil {
		req.vars = make(map[string]interface{})
	}
	req.vars[key] = value
	return req
}

// Vars gets the variables for this Request.
func (req *Request) Vars() map[string]interface{} {
	return req.vars
}

// Files gets the files in this request.
func (req *Request) Files() []File {
	return req.files
}

// Query gets the query string of this request.
func (req *Request) Query() string {
	return req.q
}

// Report stores an error to report at Run.
func (req *Request) Report(err error) *Request {
	req.err = err
	return req
}

// File sets a file to upload.
// Files are only supported with a Client that was created with
// the UseMultipartForm option.
func (req *Request) File(fieldname, filename string, r io.Reader) *Request {
	req.files = append(req.files, File{
		Field: fieldname,
		Name:  filename,
		R:     r,
	})

	return req
}

// File represents a file to upload.
type File struct {
	Field string
	Name  string
	R     io.Reader
}

// Run executes the Request with the global client.
func (req *Request) Run() (Response, error) {
	return Run(req)
}
