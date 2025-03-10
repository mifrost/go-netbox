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

package ipam

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

// NewIpamServiceTemplatesDeleteParams creates a new IpamServiceTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServiceTemplatesDeleteParams() *IpamServiceTemplatesDeleteParams {
	return &IpamServiceTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServiceTemplatesDeleteParamsWithTimeout creates a new IpamServiceTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamServiceTemplatesDeleteParamsWithTimeout(timeout time.Duration) *IpamServiceTemplatesDeleteParams {
	return &IpamServiceTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewIpamServiceTemplatesDeleteParamsWithContext creates a new IpamServiceTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewIpamServiceTemplatesDeleteParamsWithContext(ctx context.Context) *IpamServiceTemplatesDeleteParams {
	return &IpamServiceTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewIpamServiceTemplatesDeleteParamsWithHTTPClient creates a new IpamServiceTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServiceTemplatesDeleteParamsWithHTTPClient(client *http.Client) *IpamServiceTemplatesDeleteParams {
	return &IpamServiceTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/* IpamServiceTemplatesDeleteParams contains all the parameters to send to the API endpoint
   for the ipam service templates delete operation.

   Typically these are written to a http.Request.
*/
type IpamServiceTemplatesDeleteParams struct {

	/* ID.

	   A unique integer value identifying this service template.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam service templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServiceTemplatesDeleteParams) WithDefaults() *IpamServiceTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam service templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServiceTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam service templates delete params
func (o *IpamServiceTemplatesDeleteParams) WithTimeout(timeout time.Duration) *IpamServiceTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam service templates delete params
func (o *IpamServiceTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam service templates delete params
func (o *IpamServiceTemplatesDeleteParams) WithContext(ctx context.Context) *IpamServiceTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam service templates delete params
func (o *IpamServiceTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam service templates delete params
func (o *IpamServiceTemplatesDeleteParams) WithHTTPClient(client *http.Client) *IpamServiceTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam service templates delete params
func (o *IpamServiceTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam service templates delete params
func (o *IpamServiceTemplatesDeleteParams) WithID(id int64) *IpamServiceTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam service templates delete params
func (o *IpamServiceTemplatesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServiceTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
