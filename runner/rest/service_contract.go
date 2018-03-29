package rest

import "github.com/viant/endly/testing/validator"

//Request represents a send request
type Request struct {
	URL     string
	Method  string
	Request interface{}
	Expected interface{} `description:"If specified it will validated response as actual"`
}

//Response represents a rest response
type Response struct {
	Response interface{}
	*validator.AssertResponse
}
