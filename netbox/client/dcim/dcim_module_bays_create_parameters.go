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

package dcim

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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewDcimModuleBaysCreateParams creates a new DcimModuleBaysCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModuleBaysCreateParams() *DcimModuleBaysCreateParams {
	return &DcimModuleBaysCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModuleBaysCreateParamsWithTimeout creates a new DcimModuleBaysCreateParams object
// with the ability to set a timeout on a request.
func NewDcimModuleBaysCreateParamsWithTimeout(timeout time.Duration) *DcimModuleBaysCreateParams {
	return &DcimModuleBaysCreateParams{
		timeout: timeout,
	}
}

// NewDcimModuleBaysCreateParamsWithContext creates a new DcimModuleBaysCreateParams object
// with the ability to set a context for a request.
func NewDcimModuleBaysCreateParamsWithContext(ctx context.Context) *DcimModuleBaysCreateParams {
	return &DcimModuleBaysCreateParams{
		Context: ctx,
	}
}

// NewDcimModuleBaysCreateParamsWithHTTPClient creates a new DcimModuleBaysCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModuleBaysCreateParamsWithHTTPClient(client *http.Client) *DcimModuleBaysCreateParams {
	return &DcimModuleBaysCreateParams{
		HTTPClient: client,
	}
}

/* DcimModuleBaysCreateParams contains all the parameters to send to the API endpoint
   for the dcim module bays create operation.

   Typically these are written to a http.Request.
*/
type DcimModuleBaysCreateParams struct {

	// Data.
	Data *models.WritableModuleBay

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim module bays create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBaysCreateParams) WithDefaults() *DcimModuleBaysCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim module bays create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBaysCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim module bays create params
func (o *DcimModuleBaysCreateParams) WithTimeout(timeout time.Duration) *DcimModuleBaysCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim module bays create params
func (o *DcimModuleBaysCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim module bays create params
func (o *DcimModuleBaysCreateParams) WithContext(ctx context.Context) *DcimModuleBaysCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim module bays create params
func (o *DcimModuleBaysCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim module bays create params
func (o *DcimModuleBaysCreateParams) WithHTTPClient(client *http.Client) *DcimModuleBaysCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim module bays create params
func (o *DcimModuleBaysCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim module bays create params
func (o *DcimModuleBaysCreateParams) WithData(data *models.WritableModuleBay) *DcimModuleBaysCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim module bays create params
func (o *DcimModuleBaysCreateParams) SetData(data *models.WritableModuleBay) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModuleBaysCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
