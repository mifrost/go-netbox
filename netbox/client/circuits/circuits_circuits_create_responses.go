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

// CircuitsCircuitsCreateReader is a Reader for the CircuitsCircuitsCreate structure.
type CircuitsCircuitsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsCircuitsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitsCreateCreated creates a CircuitsCircuitsCreateCreated with default headers values
func NewCircuitsCircuitsCreateCreated() *CircuitsCircuitsCreateCreated {
	return &CircuitsCircuitsCreateCreated{}
}

/* CircuitsCircuitsCreateCreated describes a response with status code 201, with default header values.

CircuitsCircuitsCreateCreated circuits circuits create created
*/
type CircuitsCircuitsCreateCreated struct {
	Payload *models.Circuit
}

func (o *CircuitsCircuitsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/circuits/][%d] circuitsCircuitsCreateCreated  %+v", 201, o.Payload)
}
func (o *CircuitsCircuitsCreateCreated) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitsCreateDefault creates a CircuitsCircuitsCreateDefault with default headers values
func NewCircuitsCircuitsCreateDefault(code int) *CircuitsCircuitsCreateDefault {
	return &CircuitsCircuitsCreateDefault{
		_statusCode: code,
	}
}

/* CircuitsCircuitsCreateDefault describes a response with status code -1, with default header values.

CircuitsCircuitsCreateDefault circuits circuits create default
*/
type CircuitsCircuitsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuits create default response
func (o *CircuitsCircuitsCreateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /circuits/circuits/][%d] circuits_circuits_create default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsCircuitsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
