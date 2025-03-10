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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// TenancyContactGroupsPartialUpdateReader is a Reader for the TenancyContactGroupsPartialUpdate structure.
type TenancyContactGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactGroupsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactGroupsPartialUpdateOK creates a TenancyContactGroupsPartialUpdateOK with default headers values
func NewTenancyContactGroupsPartialUpdateOK() *TenancyContactGroupsPartialUpdateOK {
	return &TenancyContactGroupsPartialUpdateOK{}
}

/* TenancyContactGroupsPartialUpdateOK describes a response with status code 200, with default header values.

TenancyContactGroupsPartialUpdateOK tenancy contact groups partial update o k
*/
type TenancyContactGroupsPartialUpdateOK struct {
	Payload *models.ContactGroup
}

func (o *TenancyContactGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyContactGroupsPartialUpdateOK) GetPayload() *models.ContactGroup {
	return o.Payload
}

func (o *TenancyContactGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactGroupsPartialUpdateDefault creates a TenancyContactGroupsPartialUpdateDefault with default headers values
func NewTenancyContactGroupsPartialUpdateDefault(code int) *TenancyContactGroupsPartialUpdateDefault {
	return &TenancyContactGroupsPartialUpdateDefault{
		_statusCode: code,
	}
}

/* TenancyContactGroupsPartialUpdateDefault describes a response with status code -1, with default header values.

TenancyContactGroupsPartialUpdateDefault tenancy contact groups partial update default
*/
type TenancyContactGroupsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact groups partial update default response
func (o *TenancyContactGroupsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactGroupsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-groups/{id}/][%d] tenancy_contact-groups_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyContactGroupsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
