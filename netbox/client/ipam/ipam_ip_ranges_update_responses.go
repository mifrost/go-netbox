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

// IpamIPRangesUpdateReader is a Reader for the IpamIPRangesUpdate structure.
type IpamIPRangesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPRangesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamIPRangesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPRangesUpdateOK creates a IpamIPRangesUpdateOK with default headers values
func NewIpamIPRangesUpdateOK() *IpamIPRangesUpdateOK {
	return &IpamIPRangesUpdateOK{}
}

/* IpamIPRangesUpdateOK describes a response with status code 200, with default header values.

IpamIPRangesUpdateOK ipam Ip ranges update o k
*/
type IpamIPRangesUpdateOK struct {
	Payload *models.IPRange
}

func (o *IpamIPRangesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-ranges/{id}/][%d] ipamIpRangesUpdateOK  %+v", 200, o.Payload)
}
func (o *IpamIPRangesUpdateOK) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPRangesUpdateDefault creates a IpamIPRangesUpdateDefault with default headers values
func NewIpamIPRangesUpdateDefault(code int) *IpamIPRangesUpdateDefault {
	return &IpamIPRangesUpdateDefault{
		_statusCode: code,
	}
}

/* IpamIPRangesUpdateDefault describes a response with status code -1, with default header values.

IpamIPRangesUpdateDefault ipam ip ranges update default
*/
type IpamIPRangesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip ranges update default response
func (o *IpamIPRangesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamIPRangesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/ip-ranges/{id}/][%d] ipam_ip-ranges_update default  %+v", o._statusCode, o.Payload)
}
func (o *IpamIPRangesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
