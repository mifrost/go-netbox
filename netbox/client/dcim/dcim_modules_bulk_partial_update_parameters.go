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

// NewDcimModulesBulkPartialUpdateParams creates a new DcimModulesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModulesBulkPartialUpdateParams() *DcimModulesBulkPartialUpdateParams {
	return &DcimModulesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModulesBulkPartialUpdateParamsWithTimeout creates a new DcimModulesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimModulesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimModulesBulkPartialUpdateParams {
	return &DcimModulesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimModulesBulkPartialUpdateParamsWithContext creates a new DcimModulesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimModulesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimModulesBulkPartialUpdateParams {
	return &DcimModulesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimModulesBulkPartialUpdateParamsWithHTTPClient creates a new DcimModulesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModulesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimModulesBulkPartialUpdateParams {
	return &DcimModulesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimModulesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim modules bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimModulesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableModule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim modules bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesBulkPartialUpdateParams) WithDefaults() *DcimModulesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim modules bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim modules bulk partial update params
func (o *DcimModulesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimModulesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim modules bulk partial update params
func (o *DcimModulesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim modules bulk partial update params
func (o *DcimModulesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimModulesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim modules bulk partial update params
func (o *DcimModulesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim modules bulk partial update params
func (o *DcimModulesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimModulesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim modules bulk partial update params
func (o *DcimModulesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim modules bulk partial update params
func (o *DcimModulesBulkPartialUpdateParams) WithData(data *models.WritableModule) *DcimModulesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim modules bulk partial update params
func (o *DcimModulesBulkPartialUpdateParams) SetData(data *models.WritableModule) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModulesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
