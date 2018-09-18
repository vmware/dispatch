///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewUpdateBaseImageByNameParams creates a new UpdateBaseImageByNameParams object
// with the default values initialized.
func NewUpdateBaseImageByNameParams() UpdateBaseImageByNameParams {

	var (
		// initialize parameters with default values

		xDispatchProjectDefault = string("default")
	)

	return UpdateBaseImageByNameParams{
		XDispatchProject: &xDispatchProjectDefault,
	}
}

// UpdateBaseImageByNameParams contains all the bound params for the update base image by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateBaseImageByName
type UpdateBaseImageByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	XDispatchOrg string
	/*
	  Pattern: ^[\w\d][\w\d\-]*[\w\d]|[\w\d]+$
	  In: header
	  Default: "default"
	*/
	XDispatchProject *string
	/*Name of base image to return
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	BaseImageName string
	/*
	  In: body
	*/
	Body *v1.BaseImage
	/*Filter based on tags
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateBaseImageByNameParams() beforehand.
func (o *UpdateBaseImageByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXDispatchOrg(r.Header[http.CanonicalHeaderKey("X-Dispatch-Org")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXDispatchProject(r.Header[http.CanonicalHeaderKey("X-Dispatch-Project")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rBaseImageName, rhkBaseImageName, _ := route.Params.GetOK("baseImageName")
	if err := o.bindBaseImageName(rBaseImageName, rhkBaseImageName, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body v1.BaseImage
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {

			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateBaseImageByNameParams) bindXDispatchOrg(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Dispatch-Org", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Dispatch-Org", "header", raw); err != nil {
		return err
	}

	o.XDispatchOrg = raw

	return nil
}

func (o *UpdateBaseImageByNameParams) bindXDispatchProject(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewUpdateBaseImageByNameParams()
		return nil
	}

	o.XDispatchProject = &raw

	if err := o.validateXDispatchProject(formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateBaseImageByNameParams) validateXDispatchProject(formats strfmt.Registry) error {

	if err := validate.Pattern("X-Dispatch-Project", "header", (*o.XDispatchProject), `^[\w\d][\w\d\-]*[\w\d]|[\w\d]+$`); err != nil {
		return err
	}

	return nil
}

func (o *UpdateBaseImageByNameParams) bindBaseImageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.BaseImageName = raw

	if err := o.validateBaseImageName(formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateBaseImageByNameParams) validateBaseImageName(formats strfmt.Registry) error {

	if err := validate.Pattern("baseImageName", "path", o.BaseImageName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (o *UpdateBaseImageByNameParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	tagsIC := rawData

	if len(tagsIC) == 0 {
		return nil
	}

	var tagsIR []string
	for _, tagsIV := range tagsIC {
		tagsI := tagsIV

		tagsIR = append(tagsIR, tagsI)
	}

	o.Tags = tagsIR

	return nil
}
