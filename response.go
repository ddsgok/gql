package gql

import (
	"encoding/json"
	"time"

	"github.com/tidwall/gjson"
)

type Response gjson.Result

func ParseResponse(val interface{}) (Response, error) {
	if b, err := json.Marshal(val); err == nil {
		return Response(gjson.Parse(string(b))), nil
	} else {
		return emptyResponse(), err
	}
}

func emptyResponse() Response {
	return Response(gjson.Result{})
}

// String returns a string representation of the value.
func (t Response) String() string {
	return gjson.Result(t).String()
}

// Bool returns an boolean representation.
func (t Response) Bool() bool {
	return gjson.Result(t).Bool()
}

// Int returns an integer representation.
func (t Response) Int() int64 {
	return gjson.Result(t).Int()
}

// Uint returns an unsigned integer representation.
func (t Response) Uint() uint64 {
	return gjson.Result(t).Uint()
}

// Float returns an float64 representation.
func (t Response) Float() float64 {
	return gjson.Result(t).Float()
}

// Time returns a time.Time representation.
func (t Response) Time() time.Time {
	return gjson.Result(t).Time()
}

// Array returns back an array of values.
// If the result represents a non-existent value, then an empty array will be
// returned. If the result is not a JSON array, the return value will be an
// array containing one result.
func (t Response) Array() []Response {
	a := gjson.Result(t).Array()
	b := make([]Response, len(a))
	for i := range a {
		b[i] = Response(a[i])
	}

	return b
}

// IsObject returns true if the result value is a JSON object.
func (t Response) IsObject() bool {
	return gjson.Result(t).IsObject()
}

// IsArray returns true if the result value is a JSON array.
func (t Response) IsArray() bool {
	return gjson.Result(t).IsArray()
}

// ForEach iterates through values.
// If the result represents a non-existent value, then no values will be
// iterated. If the result is an Object, the iterator will pass the key and
// value of each item. If the result is an Array, the iterator will only pass
// the value of each item. If the result is not a JSON array or object, the
// iterator will pass back one value equal to the result.
func (t Response) ForEach(iterator func(key, value Response) bool) {
	it := func(key, value gjson.Result) bool {
		return iterator(Response(key), Response(value))
	}
	gjson.Result(t).ForEach(it)
}

// Map returns back an map of values. The result should be a JSON array.
func (t Response) Map() map[string]Response {
	a := gjson.Result(t).Map()
	b := make(map[string]Response, len(a))
	for i := range a {
		b[i] = Response(a[i])
	}

	return b
}

// Get searches result for the specified path.
// The result should be a JSON array or object.
func (t Response) Get(path string) Response {
	return Response(gjson.Result(t).Get(path))
}
