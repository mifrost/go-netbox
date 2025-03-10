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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasWebhooksBulkUpdateReader is a Reader for the ExtrasWebhooksBulkUpdate structure.
type ExtrasWebhooksBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasWebhooksBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasWebhooksBulkUpdateOK creates a ExtrasWebhooksBulkUpdateOK with default headers values
func NewExtrasWebhooksBulkUpdateOK() *ExtrasWebhooksBulkUpdateOK {
	return &ExtrasWebhooksBulkUpdateOK{}
}

/* ExtrasWebhooksBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksBulkUpdateOK extras webhooks bulk update o k
*/
type ExtrasWebhooksBulkUpdateOK struct {
	Payload *models.Webhook
}

func (o *ExtrasWebhooksBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/webhooks/][%d] extrasWebhooksBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasWebhooksBulkUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasWebhooksBulkUpdateDefault creates a ExtrasWebhooksBulkUpdateDefault with default headers values
func NewExtrasWebhooksBulkUpdateDefault(code int) *ExtrasWebhooksBulkUpdateDefault {
	return &ExtrasWebhooksBulkUpdateDefault{
		_statusCode: code,
	}
}

/* ExtrasWebhooksBulkUpdateDefault describes a response with status code -1, with default header values.

ExtrasWebhooksBulkUpdateDefault extras webhooks bulk update default
*/
type ExtrasWebhooksBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras webhooks bulk update default response
func (o *ExtrasWebhooksBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasWebhooksBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/webhooks/][%d] extras_webhooks_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasWebhooksBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasWebhooksBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
