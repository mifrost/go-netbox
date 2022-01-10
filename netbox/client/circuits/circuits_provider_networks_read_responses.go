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

// CircuitsProviderNetworksReadReader is a Reader for the CircuitsProviderNetworksRead structure.
type CircuitsProviderNetworksReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProviderNetworksReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProviderNetworksReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProviderNetworksReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProviderNetworksReadOK creates a CircuitsProviderNetworksReadOK with default headers values
func NewCircuitsProviderNetworksReadOK() *CircuitsProviderNetworksReadOK {
	return &CircuitsProviderNetworksReadOK{}
}

/* CircuitsProviderNetworksReadOK describes a response with status code 200, with default header values.

CircuitsProviderNetworksReadOK circuits provider networks read o k
*/
type CircuitsProviderNetworksReadOK struct {
	Payload *models.ProviderNetwork
}

func (o *CircuitsProviderNetworksReadOK) Error() string {
	return fmt.Sprintf("[GET /circuits/provider-networks/{id}/][%d] circuitsProviderNetworksReadOK  %+v", 200, o.Payload)
}
func (o *CircuitsProviderNetworksReadOK) GetPayload() *models.ProviderNetwork {
	return o.Payload
}

func (o *CircuitsProviderNetworksReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProviderNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProviderNetworksReadDefault creates a CircuitsProviderNetworksReadDefault with default headers values
func NewCircuitsProviderNetworksReadDefault(code int) *CircuitsProviderNetworksReadDefault {
	return &CircuitsProviderNetworksReadDefault{
		_statusCode: code,
	}
}

/* CircuitsProviderNetworksReadDefault describes a response with status code -1, with default header values.

CircuitsProviderNetworksReadDefault circuits provider networks read default
*/
type CircuitsProviderNetworksReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits provider networks read default response
func (o *CircuitsProviderNetworksReadDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsProviderNetworksReadDefault) Error() string {
	return fmt.Sprintf("[GET /circuits/provider-networks/{id}/][%d] circuits_provider-networks_read default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsProviderNetworksReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProviderNetworksReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
