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

// IpamFhrpGroupsBulkPartialUpdateReader is a Reader for the IpamFhrpGroupsBulkPartialUpdate structure.
type IpamFhrpGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupsBulkPartialUpdateOK creates a IpamFhrpGroupsBulkPartialUpdateOK with default headers values
func NewIpamFhrpGroupsBulkPartialUpdateOK() *IpamFhrpGroupsBulkPartialUpdateOK {
	return &IpamFhrpGroupsBulkPartialUpdateOK{}
}

/* IpamFhrpGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamFhrpGroupsBulkPartialUpdateOK ipam fhrp groups bulk partial update o k
*/
type IpamFhrpGroupsBulkPartialUpdateOK struct {
	Payload *models.FHRPGroup
}

func (o *IpamFhrpGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/fhrp-groups/][%d] ipamFhrpGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamFhrpGroupsBulkPartialUpdateOK) GetPayload() *models.FHRPGroup {
	return o.Payload
}

func (o *IpamFhrpGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupsBulkPartialUpdateDefault creates a IpamFhrpGroupsBulkPartialUpdateDefault with default headers values
func NewIpamFhrpGroupsBulkPartialUpdateDefault(code int) *IpamFhrpGroupsBulkPartialUpdateDefault {
	return &IpamFhrpGroupsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* IpamFhrpGroupsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamFhrpGroupsBulkPartialUpdateDefault ipam fhrp groups bulk partial update default
*/
type IpamFhrpGroupsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam fhrp groups bulk partial update default response
func (o *IpamFhrpGroupsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamFhrpGroupsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/fhrp-groups/][%d] ipam_fhrp-groups_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *IpamFhrpGroupsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
