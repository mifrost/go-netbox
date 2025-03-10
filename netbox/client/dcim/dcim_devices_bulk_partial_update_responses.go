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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimDevicesBulkPartialUpdateReader is a Reader for the DcimDevicesBulkPartialUpdate structure.
type DcimDevicesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDevicesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesBulkPartialUpdateOK creates a DcimDevicesBulkPartialUpdateOK with default headers values
func NewDcimDevicesBulkPartialUpdateOK() *DcimDevicesBulkPartialUpdateOK {
	return &DcimDevicesBulkPartialUpdateOK{}
}

/* DcimDevicesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDevicesBulkPartialUpdateOK dcim devices bulk partial update o k
*/
type DcimDevicesBulkPartialUpdateOK struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/devices/][%d] dcimDevicesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDevicesBulkPartialUpdateOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDevicesBulkPartialUpdateDefault creates a DcimDevicesBulkPartialUpdateDefault with default headers values
func NewDcimDevicesBulkPartialUpdateDefault(code int) *DcimDevicesBulkPartialUpdateDefault {
	return &DcimDevicesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimDevicesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimDevicesBulkPartialUpdateDefault dcim devices bulk partial update default
*/
type DcimDevicesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim devices bulk partial update default response
func (o *DcimDevicesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDevicesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/devices/][%d] dcim_devices_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimDevicesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
