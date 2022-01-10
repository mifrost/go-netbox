// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewTenancyTenantsDeleteParams creates a new TenancyTenantsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsDeleteParams() *TenancyTenantsDeleteParams {
	return &TenancyTenantsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsDeleteParamsWithTimeout creates a new TenancyTenantsDeleteParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsDeleteParamsWithTimeout(timeout time.Duration) *TenancyTenantsDeleteParams {
	return &TenancyTenantsDeleteParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsDeleteParamsWithContext creates a new TenancyTenantsDeleteParams object
// with the ability to set a context for a request.
func NewTenancyTenantsDeleteParamsWithContext(ctx context.Context) *TenancyTenantsDeleteParams {
	return &TenancyTenantsDeleteParams{
		Context: ctx,
	}
}

// NewTenancyTenantsDeleteParamsWithHTTPClient creates a new TenancyTenantsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsDeleteParamsWithHTTPClient(client *http.Client) *TenancyTenantsDeleteParams {
	return &TenancyTenantsDeleteParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsDeleteParams contains all the parameters to send to the API endpoint
   for the tenancy tenants delete operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsDeleteParams struct {

	/* ID.

	   A unique integer value identifying this tenant.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsDeleteParams) WithDefaults() *TenancyTenantsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithTimeout(timeout time.Duration) *TenancyTenantsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithContext(ctx context.Context) *TenancyTenantsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithHTTPClient(client *http.Client) *TenancyTenantsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) WithID(id int64) *TenancyTenantsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenants delete params
func (o *TenancyTenantsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
