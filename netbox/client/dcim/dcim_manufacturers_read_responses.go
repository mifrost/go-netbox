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

// DcimManufacturersReadReader is a Reader for the DcimManufacturersRead structure.
type DcimManufacturersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimManufacturersReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimManufacturersReadOK creates a DcimManufacturersReadOK with default headers values
func NewDcimManufacturersReadOK() *DcimManufacturersReadOK {
	return &DcimManufacturersReadOK{}
}

/* DcimManufacturersReadOK describes a response with status code 200, with default header values.

DcimManufacturersReadOK dcim manufacturers read o k
*/
type DcimManufacturersReadOK struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcimManufacturersReadOK  %+v", 200, o.Payload)
}
func (o *DcimManufacturersReadOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersReadDefault creates a DcimManufacturersReadDefault with default headers values
func NewDcimManufacturersReadDefault(code int) *DcimManufacturersReadDefault {
	return &DcimManufacturersReadDefault{
		_statusCode: code,
	}
}

/* DcimManufacturersReadDefault describes a response with status code -1, with default header values.

DcimManufacturersReadDefault dcim manufacturers read default
*/
type DcimManufacturersReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim manufacturers read default response
func (o *DcimManufacturersReadDefault) Code() int {
	return o._statusCode
}

func (o *DcimManufacturersReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcim_manufacturers_read default  %+v", o._statusCode, o.Payload)
}
func (o *DcimManufacturersReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
