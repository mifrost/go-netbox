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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitsCircuitTerminationsBulkPartialUpdateReader is a Reader for the CircuitsCircuitTerminationsBulkPartialUpdate structure.
type CircuitsCircuitTerminationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitTerminationsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsBulkPartialUpdateOK creates a CircuitsCircuitTerminationsBulkPartialUpdateOK with default headers values
func NewCircuitsCircuitTerminationsBulkPartialUpdateOK() *CircuitsCircuitTerminationsBulkPartialUpdateOK {
	return &CircuitsCircuitTerminationsBulkPartialUpdateOK{}
}

/* CircuitsCircuitTerminationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsBulkPartialUpdateOK circuits circuit terminations bulk partial update o k
*/
type CircuitsCircuitTerminationsBulkPartialUpdateOK struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsBulkPartialUpdateDefault creates a CircuitsCircuitTerminationsBulkPartialUpdateDefault with default headers values
func NewCircuitsCircuitTerminationsBulkPartialUpdateDefault(code int) *CircuitsCircuitTerminationsBulkPartialUpdateDefault {
	return &CircuitsCircuitTerminationsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* CircuitsCircuitTerminationsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

CircuitsCircuitTerminationsBulkPartialUpdateDefault circuits circuit terminations bulk partial update default
*/
type CircuitsCircuitTerminationsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit terminations bulk partial update default response
func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/][%d] circuits_circuit-terminations_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
