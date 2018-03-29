///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceBinding service binding
// swagger:model ServiceBinding

type ServiceBinding struct {

	// binding secret
	BindingSecret string `json:"bindingSecret,omitempty"`

	// created time
	CreatedTime int64 `json:"createdTime,omitempty"`

	// parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// reason
	Reason []string `json:"reason"`

	// secret parameters
	SecretParameters []string `json:"secretParameters"`

	// status
	Status Status `json:"status,omitempty"`
}

/* polymorph ServiceBinding bindingSecret false */

/* polymorph ServiceBinding createdTime false */

/* polymorph ServiceBinding parameters false */

/* polymorph ServiceBinding reason false */

/* polymorph ServiceBinding secretParameters false */

/* polymorph ServiceBinding status false */

// Validate validates this service binding
func (m *ServiceBinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecretParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceBinding) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	return nil
}

func (m *ServiceBinding) validateSecretParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.SecretParameters) { // not required
		return nil
	}

	return nil
}

func (m *ServiceBinding) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceBinding) UnmarshalBinary(b []byte) error {
	var res ServiceBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}