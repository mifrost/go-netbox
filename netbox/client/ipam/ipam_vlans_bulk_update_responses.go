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

// IpamVlansBulkUpdateReader is a Reader for the IpamVlansBulkUpdate structure.
type IpamVlansBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlansBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlansBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlansBulkUpdateOK creates a IpamVlansBulkUpdateOK with default headers values
func NewIpamVlansBulkUpdateOK() *IpamVlansBulkUpdateOK {
	return &IpamVlansBulkUpdateOK{}
}

/* IpamVlansBulkUpdateOK describes a response with status code 200, with default header values.

IpamVlansBulkUpdateOK ipam vlans bulk update o k
*/
type IpamVlansBulkUpdateOK struct {
	Payload *models.VLAN
}

func (o *IpamVlansBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlans/][%d] ipamVlansBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamVlansBulkUpdateOK) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlansBulkUpdateDefault creates a IpamVlansBulkUpdateDefault with default headers values
func NewIpamVlansBulkUpdateDefault(code int) *IpamVlansBulkUpdateDefault {
	return &IpamVlansBulkUpdateDefault{
		_statusCode: code,
	}
}

/* IpamVlansBulkUpdateDefault describes a response with status code -1, with default header values.

IpamVlansBulkUpdateDefault ipam vlans bulk update default
*/
type IpamVlansBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlans bulk update default response
func (o *IpamVlansBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamVlansBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlans/][%d] ipam_vlans_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *IpamVlansBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlansBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
