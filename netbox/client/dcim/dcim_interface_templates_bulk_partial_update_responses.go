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

// DcimInterfaceTemplatesBulkPartialUpdateReader is a Reader for the DcimInterfaceTemplatesBulkPartialUpdate structure.
type DcimInterfaceTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfaceTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfaceTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfaceTemplatesBulkPartialUpdateOK creates a DcimInterfaceTemplatesBulkPartialUpdateOK with default headers values
func NewDcimInterfaceTemplatesBulkPartialUpdateOK() *DcimInterfaceTemplatesBulkPartialUpdateOK {
	return &DcimInterfaceTemplatesBulkPartialUpdateOK{}
}

/* DcimInterfaceTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimInterfaceTemplatesBulkPartialUpdateOK dcim interface templates bulk partial update o k
*/
type DcimInterfaceTemplatesBulkPartialUpdateOK struct {
	Payload *models.InterfaceTemplate
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interface-templates/][%d] dcimInterfaceTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfaceTemplatesBulkPartialUpdateDefault creates a DcimInterfaceTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimInterfaceTemplatesBulkPartialUpdateDefault(code int) *DcimInterfaceTemplatesBulkPartialUpdateDefault {
	return &DcimInterfaceTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimInterfaceTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimInterfaceTemplatesBulkPartialUpdateDefault dcim interface templates bulk partial update default
*/
type DcimInterfaceTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interface templates bulk partial update default response
func (o *DcimInterfaceTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interface-templates/][%d] dcim_interface-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInterfaceTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfaceTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
