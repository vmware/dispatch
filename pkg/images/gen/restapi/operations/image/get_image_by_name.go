///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetImageByNameHandlerFunc turns a function with the right signature into a get image by name handler
type GetImageByNameHandlerFunc func(GetImageByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetImageByNameHandlerFunc) Handle(params GetImageByNameParams) middleware.Responder {
	return fn(params)
}

// GetImageByNameHandler interface for that can handle valid get image by name params
type GetImageByNameHandler interface {
	Handle(GetImageByNameParams) middleware.Responder
}

// NewGetImageByName creates a new http.Handler for the get image by name operation
func NewGetImageByName(ctx *middleware.Context, handler GetImageByNameHandler) *GetImageByName {
	return &GetImageByName{Context: ctx, Handler: handler}
}

/*GetImageByName swagger:route GET /image/{imageName} image getImageByName

Find image by ID

Returns a single image

*/
type GetImageByName struct {
	Context *middleware.Context
	Handler GetImageByNameHandler
}

func (o *GetImageByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetImageByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
