// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteFunctionHandlerFunc turns a function with the right signature into a delete function handler
type DeleteFunctionHandlerFunc func(DeleteFunctionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFunctionHandlerFunc) Handle(params DeleteFunctionParams) middleware.Responder {
	return fn(params)
}

// DeleteFunctionHandler interface for that can handle valid delete function params
type DeleteFunctionHandler interface {
	Handle(DeleteFunctionParams) middleware.Responder
}

// NewDeleteFunction creates a new http.Handler for the delete function operation
func NewDeleteFunction(ctx *middleware.Context, handler DeleteFunctionHandler) *DeleteFunction {
	return &DeleteFunction{Context: ctx, Handler: handler}
}

/*DeleteFunction swagger:route DELETE /function/{functionName} Store deleteFunction

Deletes a function

*/
type DeleteFunction struct {
	Context *middleware.Context
	Handler DeleteFunctionHandler
}

func (o *DeleteFunction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteFunctionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
