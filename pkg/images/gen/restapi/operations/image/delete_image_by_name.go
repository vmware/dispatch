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

// DeleteImageByNameHandlerFunc turns a function with the right signature into a delete image by name handler
type DeleteImageByNameHandlerFunc func(DeleteImageByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteImageByNameHandlerFunc) Handle(params DeleteImageByNameParams) middleware.Responder {
	return fn(params)
}

// DeleteImageByNameHandler interface for that can handle valid delete image by name params
type DeleteImageByNameHandler interface {
	Handle(DeleteImageByNameParams) middleware.Responder
}

// NewDeleteImageByName creates a new http.Handler for the delete image by name operation
func NewDeleteImageByName(ctx *middleware.Context, handler DeleteImageByNameHandler) *DeleteImageByName {
	return &DeleteImageByName{Context: ctx, Handler: handler}
}

/*DeleteImageByName swagger:route DELETE /image/{imageName} image deleteImageByName

Deletes an image

*/
type DeleteImageByName struct {
	Context *middleware.Context
	Handler DeleteImageByNameHandler
}

func (o *DeleteImageByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteImageByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
