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

// DcimRackReservationsReadReader is a Reader for the DcimRackReservationsRead structure.
type DcimRackReservationsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackReservationsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackReservationsReadOK creates a DcimRackReservationsReadOK with default headers values
func NewDcimRackReservationsReadOK() *DcimRackReservationsReadOK {
	return &DcimRackReservationsReadOK{}
}

/* DcimRackReservationsReadOK describes a response with status code 200, with default header values.

DcimRackReservationsReadOK dcim rack reservations read o k
*/
type DcimRackReservationsReadOK struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-reservations/{id}/][%d] dcimRackReservationsReadOK  %+v", 200, o.Payload)
}
func (o *DcimRackReservationsReadOK) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackReservationsReadDefault creates a DcimRackReservationsReadDefault with default headers values
func NewDcimRackReservationsReadDefault(code int) *DcimRackReservationsReadDefault {
	return &DcimRackReservationsReadDefault{
		_statusCode: code,
	}
}

/* DcimRackReservationsReadDefault describes a response with status code -1, with default header values.

DcimRackReservationsReadDefault dcim rack reservations read default
*/
type DcimRackReservationsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack reservations read default response
func (o *DcimRackReservationsReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackReservationsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-reservations/{id}/][%d] dcim_rack-reservations_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRackReservationsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackReservationsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
