//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewTenantExistsParams creates a new TenantExistsParams object
// with the default values initialized.
func NewTenantExistsParams() TenantExistsParams {

	var (
		// initialize parameters with default values

		consistencyDefault = bool(true)
	)

	return TenantExistsParams{
		Consistency: &consistencyDefault,
	}
}

// TenantExistsParams contains all the bound params for the tenant exists operation
// typically these are obtained from a http.Request
//
// swagger:parameters tenant.exists
type TenantExistsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ClassName string
	/*If consistency is true, the request will be proxied to the leader to ensure strong schema consistency
	  In: header
	  Default: true
	*/
	Consistency *bool
	/*
	  Required: true
	  In: path
	*/
	TenantName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewTenantExistsParams() beforehand.
func (o *TenantExistsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rClassName, rhkClassName, _ := route.Params.GetOK("className")
	if err := o.bindClassName(rClassName, rhkClassName, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindConsistency(r.Header[http.CanonicalHeaderKey("consistency")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rTenantName, rhkTenantName, _ := route.Params.GetOK("tenantName")
	if err := o.bindTenantName(rTenantName, rhkTenantName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClassName binds and validates parameter ClassName from path.
func (o *TenantExistsParams) bindClassName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ClassName = raw

	return nil
}

// bindConsistency binds and validates parameter Consistency from header.
func (o *TenantExistsParams) bindConsistency(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewTenantExistsParams()
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("consistency", "header", "bool", raw)
	}
	o.Consistency = &value

	return nil
}

// bindTenantName binds and validates parameter TenantName from path.
func (o *TenantExistsParams) bindTenantName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.TenantName = raw

	return nil
}
