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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewIpamAsnsPartialUpdateParams creates a new IpamAsnsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAsnsPartialUpdateParams() *IpamAsnsPartialUpdateParams {
	return &IpamAsnsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAsnsPartialUpdateParamsWithTimeout creates a new IpamAsnsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamAsnsPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamAsnsPartialUpdateParams {
	return &IpamAsnsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamAsnsPartialUpdateParamsWithContext creates a new IpamAsnsPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamAsnsPartialUpdateParamsWithContext(ctx context.Context) *IpamAsnsPartialUpdateParams {
	return &IpamAsnsPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamAsnsPartialUpdateParamsWithHTTPClient creates a new IpamAsnsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAsnsPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamAsnsPartialUpdateParams {
	return &IpamAsnsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamAsnsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam asns partial update operation.

   Typically these are written to a http.Request.
*/
type IpamAsnsPartialUpdateParams struct {

	// Data.
	Data *models.WritableASN

	/* ID.

	   A unique integer value identifying this ASN.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam asns partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAsnsPartialUpdateParams) WithDefaults() *IpamAsnsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam asns partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAsnsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamAsnsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) WithContext(ctx context.Context) *IpamAsnsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamAsnsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) WithData(data *models.WritableASN) *IpamAsnsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) SetData(data *models.WritableASN) {
	o.Data = data
}

// WithID adds the id to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) WithID(id int64) *IpamAsnsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam asns partial update params
func (o *IpamAsnsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAsnsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
