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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamVrfsReadReader is a Reader for the IpamVrfsRead structure.
type IpamVrfsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsReadOK creates a IpamVrfsReadOK with default headers values
func NewIpamVrfsReadOK() *IpamVrfsReadOK {
	return &IpamVrfsReadOK{}
}

/* IpamVrfsReadOK describes a response with status code 200, with default header values.

IpamVrfsReadOK ipam vrfs read o k
*/
type IpamVrfsReadOK struct {
	Payload *models.VRF
}

func (o *IpamVrfsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/vrfs/{id}/][%d] ipamVrfsReadOK  %+v", 200, o.Payload)
}
func (o *IpamVrfsReadOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVrfsReadDefault creates a IpamVrfsReadDefault with default headers values
func NewIpamVrfsReadDefault(code int) *IpamVrfsReadDefault {
	return &IpamVrfsReadDefault{
		_statusCode: code,
	}
}

/* IpamVrfsReadDefault describes a response with status code -1, with default header values.

IpamVrfsReadDefault ipam vrfs read default
*/
type IpamVrfsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vrfs read default response
func (o *IpamVrfsReadDefault) Code() int {
	return o._statusCode
}

func (o *IpamVrfsReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/vrfs/{id}/][%d] ipam_vrfs_read default  %+v", o._statusCode, o.Payload)
}
func (o *IpamVrfsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
