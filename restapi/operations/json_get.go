// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// JSONGetHandlerFunc turns a function with the right signature into a Json get handler
type JSONGetHandlerFunc func(JSONGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn JSONGetHandlerFunc) Handle(params JSONGetParams) middleware.Responder {
	return fn(params)
}

// JSONGetHandler interface for that can handle valid Json get params
type JSONGetHandler interface {
	Handle(JSONGetParams) middleware.Responder
}

// NewJSONGet creates a new http.Handler for the Json get operation
func NewJSONGet(ctx *middleware.Context, handler JSONGetHandler) *JSONGet {
	return &JSONGet{Context: ctx, Handler: handler}
}

/*JSONGet swagger:route GET /json jsonGet

Returns a list of repos from URL

*/
type JSONGet struct {
	Context *middleware.Context
	Handler JSONGetHandler
}

func (o *JSONGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewJSONGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
