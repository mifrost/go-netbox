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

// DcimPowerPanelsCreateReader is a Reader for the DcimPowerPanelsCreate structure.
type DcimPowerPanelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerPanelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPanelsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPanelsCreateCreated creates a DcimPowerPanelsCreateCreated with default headers values
func NewDcimPowerPanelsCreateCreated() *DcimPowerPanelsCreateCreated {
	return &DcimPowerPanelsCreateCreated{}
}

/* DcimPowerPanelsCreateCreated describes a response with status code 201, with default header values.

DcimPowerPanelsCreateCreated dcim power panels create created
*/
type DcimPowerPanelsCreateCreated struct {
	Payload *models.PowerPanel
}

func (o *DcimPowerPanelsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-panels/][%d] dcimPowerPanelsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimPowerPanelsCreateCreated) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsCreateDefault creates a DcimPowerPanelsCreateDefault with default headers values
func NewDcimPowerPanelsCreateDefault(code int) *DcimPowerPanelsCreateDefault {
	return &DcimPowerPanelsCreateDefault{
		_statusCode: code,
	}
}

/* DcimPowerPanelsCreateDefault describes a response with status code -1, with default header values.

DcimPowerPanelsCreateDefault dcim power panels create default
*/
type DcimPowerPanelsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power panels create default response
func (o *DcimPowerPanelsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPanelsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/power-panels/][%d] dcim_power-panels_create default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPowerPanelsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
