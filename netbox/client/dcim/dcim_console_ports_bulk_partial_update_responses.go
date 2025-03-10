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

// DcimConsolePortsBulkPartialUpdateReader is a Reader for the DcimConsolePortsBulkPartialUpdate structure.
type DcimConsolePortsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsolePortsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortsBulkPartialUpdateOK creates a DcimConsolePortsBulkPartialUpdateOK with default headers values
func NewDcimConsolePortsBulkPartialUpdateOK() *DcimConsolePortsBulkPartialUpdateOK {
	return &DcimConsolePortsBulkPartialUpdateOK{}
}

/* DcimConsolePortsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortsBulkPartialUpdateOK dcim console ports bulk partial update o k
*/
type DcimConsolePortsBulkPartialUpdateOK struct {
	Payload *models.ConsolePort
}

func (o *DcimConsolePortsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-ports/][%d] dcimConsolePortsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimConsolePortsBulkPartialUpdateOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortsBulkPartialUpdateDefault creates a DcimConsolePortsBulkPartialUpdateDefault with default headers values
func NewDcimConsolePortsBulkPartialUpdateDefault(code int) *DcimConsolePortsBulkPartialUpdateDefault {
	return &DcimConsolePortsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimConsolePortsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimConsolePortsBulkPartialUpdateDefault dcim console ports bulk partial update default
*/
type DcimConsolePortsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console ports bulk partial update default response
func (o *DcimConsolePortsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsolePortsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-ports/][%d] dcim_console-ports_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimConsolePortsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
