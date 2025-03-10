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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewIpamServicesListParams creates a new IpamServicesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesListParams() *IpamServicesListParams {
	return &IpamServicesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesListParamsWithTimeout creates a new IpamServicesListParams object
// with the ability to set a timeout on a request.
func NewIpamServicesListParamsWithTimeout(timeout time.Duration) *IpamServicesListParams {
	return &IpamServicesListParams{
		timeout: timeout,
	}
}

// NewIpamServicesListParamsWithContext creates a new IpamServicesListParams object
// with the ability to set a context for a request.
func NewIpamServicesListParamsWithContext(ctx context.Context) *IpamServicesListParams {
	return &IpamServicesListParams{
		Context: ctx,
	}
}

// NewIpamServicesListParamsWithHTTPClient creates a new IpamServicesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesListParamsWithHTTPClient(client *http.Client) *IpamServicesListParams {
	return &IpamServicesListParams{
		HTTPClient: client,
	}
}

/* IpamServicesListParams contains all the parameters to send to the API endpoint
   for the ipam services list operation.

   Typically these are written to a http.Request.
*/
type IpamServicesListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Description.
	Description *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

	// Device.
	Device *string

	// Devicen.
	Devicen *string

	// DeviceID.
	DeviceID *string

	// DeviceIDn.
	DeviceIDn *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Port.
	Port *float64

	// Protocol.
	Protocol *string

	// Protocoln.
	Protocoln *string

	// Q.
	Q *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// VirtualMachine.
	VirtualMachine *string

	// VirtualMachinen.
	VirtualMachinen *string

	// VirtualMachineID.
	VirtualMachineID *string

	// VirtualMachineIDn.
	VirtualMachineIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesListParams) WithDefaults() *IpamServicesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services list params
func (o *IpamServicesListParams) WithTimeout(timeout time.Duration) *IpamServicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services list params
func (o *IpamServicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services list params
func (o *IpamServicesListParams) WithContext(ctx context.Context) *IpamServicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services list params
func (o *IpamServicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services list params
func (o *IpamServicesListParams) WithHTTPClient(client *http.Client) *IpamServicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services list params
func (o *IpamServicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam services list params
func (o *IpamServicesListParams) WithCreated(created *string) *IpamServicesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam services list params
func (o *IpamServicesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam services list params
func (o *IpamServicesListParams) WithCreatedGte(createdGte *string) *IpamServicesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam services list params
func (o *IpamServicesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam services list params
func (o *IpamServicesListParams) WithCreatedLte(createdLte *string) *IpamServicesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam services list params
func (o *IpamServicesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDescription adds the description to the ipam services list params
func (o *IpamServicesListParams) WithDescription(description *string) *IpamServicesListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the ipam services list params
func (o *IpamServicesListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the ipam services list params
func (o *IpamServicesListParams) WithDescriptionIc(descriptionIc *string) *IpamServicesListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the ipam services list params
func (o *IpamServicesListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the ipam services list params
func (o *IpamServicesListParams) WithDescriptionIe(descriptionIe *string) *IpamServicesListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the ipam services list params
func (o *IpamServicesListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the ipam services list params
func (o *IpamServicesListParams) WithDescriptionIew(descriptionIew *string) *IpamServicesListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the ipam services list params
func (o *IpamServicesListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the ipam services list params
func (o *IpamServicesListParams) WithDescriptionIsw(descriptionIsw *string) *IpamServicesListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the ipam services list params
func (o *IpamServicesListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the ipam services list params
func (o *IpamServicesListParams) WithDescriptionn(descriptionn *string) *IpamServicesListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the ipam services list params
func (o *IpamServicesListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the ipam services list params
func (o *IpamServicesListParams) WithDescriptionNic(descriptionNic *string) *IpamServicesListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the ipam services list params
func (o *IpamServicesListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the ipam services list params
func (o *IpamServicesListParams) WithDescriptionNie(descriptionNie *string) *IpamServicesListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the ipam services list params
func (o *IpamServicesListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the ipam services list params
func (o *IpamServicesListParams) WithDescriptionNiew(descriptionNiew *string) *IpamServicesListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the ipam services list params
func (o *IpamServicesListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the ipam services list params
func (o *IpamServicesListParams) WithDescriptionNisw(descriptionNisw *string) *IpamServicesListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the ipam services list params
func (o *IpamServicesListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithDevice adds the device to the ipam services list params
func (o *IpamServicesListParams) WithDevice(device *string) *IpamServicesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the ipam services list params
func (o *IpamServicesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDevicen adds the devicen to the ipam services list params
func (o *IpamServicesListParams) WithDevicen(devicen *string) *IpamServicesListParams {
	o.SetDevicen(devicen)
	return o
}

// SetDevicen adds the deviceN to the ipam services list params
func (o *IpamServicesListParams) SetDevicen(devicen *string) {
	o.Devicen = devicen
}

// WithDeviceID adds the deviceID to the ipam services list params
func (o *IpamServicesListParams) WithDeviceID(deviceID *string) *IpamServicesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the ipam services list params
func (o *IpamServicesListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIDn adds the deviceIDn to the ipam services list params
func (o *IpamServicesListParams) WithDeviceIDn(deviceIDn *string) *IpamServicesListParams {
	o.SetDeviceIDn(deviceIDn)
	return o
}

// SetDeviceIDn adds the deviceIdN to the ipam services list params
func (o *IpamServicesListParams) SetDeviceIDn(deviceIDn *string) {
	o.DeviceIDn = deviceIDn
}

// WithID adds the id to the ipam services list params
func (o *IpamServicesListParams) WithID(id *string) *IpamServicesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services list params
func (o *IpamServicesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the ipam services list params
func (o *IpamServicesListParams) WithIDGt(iDGt *string) *IpamServicesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the ipam services list params
func (o *IpamServicesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the ipam services list params
func (o *IpamServicesListParams) WithIDGte(iDGte *string) *IpamServicesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the ipam services list params
func (o *IpamServicesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the ipam services list params
func (o *IpamServicesListParams) WithIDLt(iDLt *string) *IpamServicesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the ipam services list params
func (o *IpamServicesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the ipam services list params
func (o *IpamServicesListParams) WithIDLte(iDLte *string) *IpamServicesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the ipam services list params
func (o *IpamServicesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the ipam services list params
func (o *IpamServicesListParams) WithIDn(iDn *string) *IpamServicesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam services list params
func (o *IpamServicesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the ipam services list params
func (o *IpamServicesListParams) WithLastUpdated(lastUpdated *string) *IpamServicesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam services list params
func (o *IpamServicesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam services list params
func (o *IpamServicesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamServicesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam services list params
func (o *IpamServicesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam services list params
func (o *IpamServicesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamServicesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam services list params
func (o *IpamServicesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam services list params
func (o *IpamServicesListParams) WithLimit(limit *int64) *IpamServicesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam services list params
func (o *IpamServicesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam services list params
func (o *IpamServicesListParams) WithName(name *string) *IpamServicesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam services list params
func (o *IpamServicesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the ipam services list params
func (o *IpamServicesListParams) WithNameIc(nameIc *string) *IpamServicesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the ipam services list params
func (o *IpamServicesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the ipam services list params
func (o *IpamServicesListParams) WithNameIe(nameIe *string) *IpamServicesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the ipam services list params
func (o *IpamServicesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the ipam services list params
func (o *IpamServicesListParams) WithNameIew(nameIew *string) *IpamServicesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the ipam services list params
func (o *IpamServicesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the ipam services list params
func (o *IpamServicesListParams) WithNameIsw(nameIsw *string) *IpamServicesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the ipam services list params
func (o *IpamServicesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the ipam services list params
func (o *IpamServicesListParams) WithNamen(namen *string) *IpamServicesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the ipam services list params
func (o *IpamServicesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the ipam services list params
func (o *IpamServicesListParams) WithNameNic(nameNic *string) *IpamServicesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the ipam services list params
func (o *IpamServicesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the ipam services list params
func (o *IpamServicesListParams) WithNameNie(nameNie *string) *IpamServicesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the ipam services list params
func (o *IpamServicesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the ipam services list params
func (o *IpamServicesListParams) WithNameNiew(nameNiew *string) *IpamServicesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the ipam services list params
func (o *IpamServicesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the ipam services list params
func (o *IpamServicesListParams) WithNameNisw(nameNisw *string) *IpamServicesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the ipam services list params
func (o *IpamServicesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the ipam services list params
func (o *IpamServicesListParams) WithOffset(offset *int64) *IpamServicesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam services list params
func (o *IpamServicesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPort adds the port to the ipam services list params
func (o *IpamServicesListParams) WithPort(port *float64) *IpamServicesListParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the ipam services list params
func (o *IpamServicesListParams) SetPort(port *float64) {
	o.Port = port
}

// WithProtocol adds the protocol to the ipam services list params
func (o *IpamServicesListParams) WithProtocol(protocol *string) *IpamServicesListParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the ipam services list params
func (o *IpamServicesListParams) SetProtocol(protocol *string) {
	o.Protocol = protocol
}

// WithProtocoln adds the protocoln to the ipam services list params
func (o *IpamServicesListParams) WithProtocoln(protocoln *string) *IpamServicesListParams {
	o.SetProtocoln(protocoln)
	return o
}

// SetProtocoln adds the protocolN to the ipam services list params
func (o *IpamServicesListParams) SetProtocoln(protocoln *string) {
	o.Protocoln = protocoln
}

// WithQ adds the q to the ipam services list params
func (o *IpamServicesListParams) WithQ(q *string) *IpamServicesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam services list params
func (o *IpamServicesListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the ipam services list params
func (o *IpamServicesListParams) WithTag(tag *string) *IpamServicesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam services list params
func (o *IpamServicesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam services list params
func (o *IpamServicesListParams) WithTagn(tagn *string) *IpamServicesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam services list params
func (o *IpamServicesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithVirtualMachine adds the virtualMachine to the ipam services list params
func (o *IpamServicesListParams) WithVirtualMachine(virtualMachine *string) *IpamServicesListParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the ipam services list params
func (o *IpamServicesListParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachinen adds the virtualMachinen to the ipam services list params
func (o *IpamServicesListParams) WithVirtualMachinen(virtualMachinen *string) *IpamServicesListParams {
	o.SetVirtualMachinen(virtualMachinen)
	return o
}

// SetVirtualMachinen adds the virtualMachineN to the ipam services list params
func (o *IpamServicesListParams) SetVirtualMachinen(virtualMachinen *string) {
	o.VirtualMachinen = virtualMachinen
}

// WithVirtualMachineID adds the virtualMachineID to the ipam services list params
func (o *IpamServicesListParams) WithVirtualMachineID(virtualMachineID *string) *IpamServicesListParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the ipam services list params
func (o *IpamServicesListParams) SetVirtualMachineID(virtualMachineID *string) {
	o.VirtualMachineID = virtualMachineID
}

// WithVirtualMachineIDn adds the virtualMachineIDn to the ipam services list params
func (o *IpamServicesListParams) WithVirtualMachineIDn(virtualMachineIDn *string) *IpamServicesListParams {
	o.SetVirtualMachineIDn(virtualMachineIDn)
	return o
}

// SetVirtualMachineIDn adds the virtualMachineIdN to the ipam services list params
func (o *IpamServicesListParams) SetVirtualMachineIDn(virtualMachineIDn *string) {
	o.VirtualMachineIDn = virtualMachineIDn
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
				return err
			}
		}
	}

	if o.Device != nil {

		// query param device
		var qrDevice string

		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {

			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}
	}

	if o.Devicen != nil {

		// query param device__n
		var qrDevicen string

		if o.Devicen != nil {
			qrDevicen = *o.Devicen
		}
		qDevicen := qrDevicen
		if qDevicen != "" {

			if err := r.SetQueryParam("device__n", qDevicen); err != nil {
				return err
			}
		}
	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string

		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {

			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}
	}

	if o.DeviceIDn != nil {

		// query param device_id__n
		var qrDeviceIDn string

		if o.DeviceIDn != nil {
			qrDeviceIDn = *o.DeviceIDn
		}
		qDeviceIDn := qrDeviceIDn
		if qDeviceIDn != "" {

			if err := r.SetQueryParam("device_id__n", qDeviceIDn); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Port != nil {

		// query param port
		var qrPort float64

		if o.Port != nil {
			qrPort = *o.Port
		}
		qPort := swag.FormatFloat64(qrPort)
		if qPort != "" {

			if err := r.SetQueryParam("port", qPort); err != nil {
				return err
			}
		}
	}

	if o.Protocol != nil {

		// query param protocol
		var qrProtocol string

		if o.Protocol != nil {
			qrProtocol = *o.Protocol
		}
		qProtocol := qrProtocol
		if qProtocol != "" {

			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
				return err
			}
		}
	}

	if o.Protocoln != nil {

		// query param protocol__n
		var qrProtocoln string

		if o.Protocoln != nil {
			qrProtocoln = *o.Protocoln
		}
		qProtocoln := qrProtocoln
		if qProtocoln != "" {

			if err := r.SetQueryParam("protocol__n", qProtocoln); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachine != nil {

		// query param virtual_machine
		var qrVirtualMachine string

		if o.VirtualMachine != nil {
			qrVirtualMachine = *o.VirtualMachine
		}
		qVirtualMachine := qrVirtualMachine
		if qVirtualMachine != "" {

			if err := r.SetQueryParam("virtual_machine", qVirtualMachine); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachinen != nil {

		// query param virtual_machine__n
		var qrVirtualMachinen string

		if o.VirtualMachinen != nil {
			qrVirtualMachinen = *o.VirtualMachinen
		}
		qVirtualMachinen := qrVirtualMachinen
		if qVirtualMachinen != "" {

			if err := r.SetQueryParam("virtual_machine__n", qVirtualMachinen); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachineID != nil {

		// query param virtual_machine_id
		var qrVirtualMachineID string

		if o.VirtualMachineID != nil {
			qrVirtualMachineID = *o.VirtualMachineID
		}
		qVirtualMachineID := qrVirtualMachineID
		if qVirtualMachineID != "" {

			if err := r.SetQueryParam("virtual_machine_id", qVirtualMachineID); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachineIDn != nil {

		// query param virtual_machine_id__n
		var qrVirtualMachineIDn string

		if o.VirtualMachineIDn != nil {
			qrVirtualMachineIDn = *o.VirtualMachineIDn
		}
		qVirtualMachineIDn := qrVirtualMachineIDn
		if qVirtualMachineIDn != "" {

			if err := r.SetQueryParam("virtual_machine_id__n", qVirtualMachineIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
