///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package serviceaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/dispatch/pkg/identity-manager/gen/models"
)

// NewUpdateServiceAccountParams creates a new UpdateServiceAccountParams object
// no default values defined in spec.
func NewUpdateServiceAccountParams() UpdateServiceAccountParams {

	return UpdateServiceAccountParams{}
}

// UpdateServiceAccountParams contains all the bound params for the update service account operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateServiceAccount
type UpdateServiceAccountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Service Account object
	  Required: true
	  In: body
	*/
	Body *models.ServiceAccount
	/*Name of ServiceAccount to work on
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	ServiceAccountName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateServiceAccountParams() beforehand.
func (o *UpdateServiceAccountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ServiceAccount
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {

			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	rServiceAccountName, rhkServiceAccountName, _ := route.Params.GetOK("serviceAccountName")
	if err := o.bindServiceAccountName(rServiceAccountName, rhkServiceAccountName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateServiceAccountParams) bindServiceAccountName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ServiceAccountName = raw

	if err := o.validateServiceAccountName(formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateServiceAccountParams) validateServiceAccountName(formats strfmt.Registry) error {

	if err := validate.Pattern("serviceAccountName", "path", o.ServiceAccountName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}
