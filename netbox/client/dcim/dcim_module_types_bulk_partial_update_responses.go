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

// DcimModuleTypesBulkPartialUpdateReader is a Reader for the DcimModuleTypesBulkPartialUpdate structure.
type DcimModuleTypesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleTypesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleTypesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleTypesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleTypesBulkPartialUpdateOK creates a DcimModuleTypesBulkPartialUpdateOK with default headers values
func NewDcimModuleTypesBulkPartialUpdateOK() *DcimModuleTypesBulkPartialUpdateOK {
	return &DcimModuleTypesBulkPartialUpdateOK{}
}

/* DcimModuleTypesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimModuleTypesBulkPartialUpdateOK dcim module types bulk partial update o k
*/
type DcimModuleTypesBulkPartialUpdateOK struct {
	Payload *models.ModuleType
}

func (o *DcimModuleTypesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-types/][%d] dcimModuleTypesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimModuleTypesBulkPartialUpdateOK) GetPayload() *models.ModuleType {
	return o.Payload
}

func (o *DcimModuleTypesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleTypesBulkPartialUpdateDefault creates a DcimModuleTypesBulkPartialUpdateDefault with default headers values
func NewDcimModuleTypesBulkPartialUpdateDefault(code int) *DcimModuleTypesBulkPartialUpdateDefault {
	return &DcimModuleTypesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimModuleTypesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimModuleTypesBulkPartialUpdateDefault dcim module types bulk partial update default
*/
type DcimModuleTypesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module types bulk partial update default response
func (o *DcimModuleTypesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimModuleTypesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-types/][%d] dcim_module-types_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimModuleTypesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleTypesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
