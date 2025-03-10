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

package extras

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

// NewExtrasImageAttachmentsBulkPartialUpdateParams creates a new ExtrasImageAttachmentsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasImageAttachmentsBulkPartialUpdateParams() *ExtrasImageAttachmentsBulkPartialUpdateParams {
	return &ExtrasImageAttachmentsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsBulkPartialUpdateParamsWithTimeout creates a new ExtrasImageAttachmentsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasImageAttachmentsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsBulkPartialUpdateParams {
	return &ExtrasImageAttachmentsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsBulkPartialUpdateParamsWithContext creates a new ExtrasImageAttachmentsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasImageAttachmentsBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsBulkPartialUpdateParams {
	return &ExtrasImageAttachmentsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasImageAttachmentsBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasImageAttachmentsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasImageAttachmentsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsBulkPartialUpdateParams {
	return &ExtrasImageAttachmentsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasImageAttachmentsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras image attachments bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasImageAttachmentsBulkPartialUpdateParams struct {

	// Data.
	Data *models.ImageAttachment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras image attachments bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) WithDefaults() *ExtrasImageAttachmentsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras image attachments bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras image attachments bulk partial update params
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments bulk partial update params
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments bulk partial update params
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments bulk partial update params
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments bulk partial update params
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments bulk partial update params
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras image attachments bulk partial update params
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) WithData(data *models.ImageAttachment) *ExtrasImageAttachmentsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras image attachments bulk partial update params
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) SetData(data *models.ImageAttachment) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
