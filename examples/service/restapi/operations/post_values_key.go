// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostValuesKeyHandlerFunc turns a function with the right signature into a post values key handler
type PostValuesKeyHandlerFunc func(PostValuesKeyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostValuesKeyHandlerFunc) Handle(params PostValuesKeyParams) middleware.Responder {
	return fn(params)
}

// PostValuesKeyHandler interface for that can handle valid post values key params
type PostValuesKeyHandler interface {
	Handle(PostValuesKeyParams) middleware.Responder
}

// NewPostValuesKey creates a new http.Handler for the post values key operation
func NewPostValuesKey(ctx *middleware.Context, handler PostValuesKeyHandler) *PostValuesKey {
	return &PostValuesKey{Context: ctx, Handler: handler}
}

/*PostValuesKey swagger:route POST /values/{key} postValuesKey

Write to DB

*/
type PostValuesKey struct {
	Context *middleware.Context
	Handler PostValuesKeyHandler
}

func (o *PostValuesKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostValuesKeyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
