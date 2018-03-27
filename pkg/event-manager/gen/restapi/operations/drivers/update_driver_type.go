///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateDriverTypeHandlerFunc turns a function with the right signature into a update driver type handler
type UpdateDriverTypeHandlerFunc func(UpdateDriverTypeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateDriverTypeHandlerFunc) Handle(params UpdateDriverTypeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateDriverTypeHandler interface for that can handle valid update driver type params
type UpdateDriverTypeHandler interface {
	Handle(UpdateDriverTypeParams, interface{}) middleware.Responder
}

// NewUpdateDriverType creates a new http.Handler for the update driver type operation
func NewUpdateDriverType(ctx *middleware.Context, handler UpdateDriverTypeHandler) *UpdateDriverType {
	return &UpdateDriverType{Context: ctx, Handler: handler}
}

/*UpdateDriverType swagger:route PUT /drivertypes/{driverTypeName} drivers updateDriverType

Update a driver type by Name

Updates a single driver type

*/
type UpdateDriverType struct {
	Context *middleware.Context
	Handler UpdateDriverTypeHandler
}

func (o *UpdateDriverType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateDriverTypeParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
