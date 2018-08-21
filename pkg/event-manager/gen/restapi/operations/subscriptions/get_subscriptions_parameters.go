///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSubscriptionsParams creates a new GetSubscriptionsParams object
// with the default values initialized.
func NewGetSubscriptionsParams() GetSubscriptionsParams {

	var (
		// initialize parameters with default values

		xDispatchProjectDefault = string("default")
	)

	return GetSubscriptionsParams{
		XDispatchProject: &xDispatchProjectDefault,
	}
}

// GetSubscriptionsParams contains all the bound params for the get subscriptions operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSubscriptions
type GetSubscriptionsParams struct {

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
	/*Filter based on tags
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSubscriptionsParams() beforehand.
func (o *GetSubscriptionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXDispatchOrg(r.Header[http.CanonicalHeaderKey("X-Dispatch-Org")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXDispatchProject(r.Header[http.CanonicalHeaderKey("X-Dispatch-Project")], true, route.Formats); err != nil {
		res = append(res, err)
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

func (o *GetSubscriptionsParams) bindXDispatchOrg(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *GetSubscriptionsParams) bindXDispatchProject(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetSubscriptionsParams()
		return nil
	}

	o.XDispatchProject = &raw

	if err := o.validateXDispatchProject(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetSubscriptionsParams) validateXDispatchProject(formats strfmt.Registry) error {

	if err := validate.Pattern("X-Dispatch-Project", "header", (*o.XDispatchProject), `^[\w\d][\w\d\-]*[\w\d]|[\w\d]+$`); err != nil {
		return err
	}

	return nil
}

func (o *GetSubscriptionsParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {

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
