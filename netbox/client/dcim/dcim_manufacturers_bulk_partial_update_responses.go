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

// DcimManufacturersBulkPartialUpdateReader is a Reader for the DcimManufacturersBulkPartialUpdate structure.
type DcimManufacturersBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimManufacturersBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimManufacturersBulkPartialUpdateOK creates a DcimManufacturersBulkPartialUpdateOK with default headers values
func NewDcimManufacturersBulkPartialUpdateOK() *DcimManufacturersBulkPartialUpdateOK {
	return &DcimManufacturersBulkPartialUpdateOK{}
}

/* DcimManufacturersBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimManufacturersBulkPartialUpdateOK dcim manufacturers bulk partial update o k
*/
type DcimManufacturersBulkPartialUpdateOK struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/][%d] dcimManufacturersBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimManufacturersBulkPartialUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersBulkPartialUpdateDefault creates a DcimManufacturersBulkPartialUpdateDefault with default headers values
func NewDcimManufacturersBulkPartialUpdateDefault(code int) *DcimManufacturersBulkPartialUpdateDefault {
	return &DcimManufacturersBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimManufacturersBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimManufacturersBulkPartialUpdateDefault dcim manufacturers bulk partial update default
*/
type DcimManufacturersBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim manufacturers bulk partial update default response
func (o *DcimManufacturersBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimManufacturersBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/][%d] dcim_manufacturers_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimManufacturersBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
