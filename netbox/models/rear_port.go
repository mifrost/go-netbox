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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RearPort rear port
//
// swagger:model RearPort
type RearPort struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable peer
	//
	//
	// Return the appropriate serializer for the cable termination model.
	//
	// Read Only: true
	CablePeer map[string]*string `json:"cable_peer,omitempty"`

	// Cable peer type
	// Read Only: true
	CablePeerType string `json:"cable_peer_type,omitempty"`

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Positions
	// Maximum: 1024
	// Minimum: 1
	Positions int64 `json:"positions,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	// Required: true
	Type *RearPortType `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rear port
func (m *RearPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RearPort) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", m.Color, 6); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", m.Color, `^[0-9a-f]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validatePositions(formats strfmt.Registry) error {
	if swag.IsZero(m.Positions) { // not required
		return nil
	}

	if err := validate.MinimumInt("positions", "body", m.Positions, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("positions", "body", m.Positions, 1024, false); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RearPort) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rear port based on the context it is used
func (m *RearPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RearPort) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *RearPort) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *RearPort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RearPort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RearPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPort) UnmarshalBinary(b []byte) error {
	var res RearPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RearPortType Type
//
// swagger:model RearPortType
type RearPortType struct {

	// label
	// Required: true
	// Enum: [8P8C 8P6C 8P4C 8P2C 6P6C 6P4C 6P2C 4P4C 4P2C GG45 TERA 4P TERA 2P TERA 1P 110 Punch BNC F Connector N Connector MRJ21 FC LC LC/APC LSH LSH/APC MPO MTRJ SC SC/APC ST CS SN SMA 905 SMA 906 URM-P2 URM-P4 URM-P8 Splice]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c 6p6c 6p4c 6p2c 4p4c 4p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc f n mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st cs sn sma-905 sma-906 urm-p2 urm-p4 urm-p8 splice]
	Value *string `json:"value"`
}

// Validate validates this rear port type
func (m *RearPortType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rearPortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8P8C","8P6C","8P4C","8P2C","6P6C","6P4C","6P2C","4P4C","4P2C","GG45","TERA 4P","TERA 2P","TERA 1P","110 Punch","BNC","F Connector","N Connector","MRJ21","FC","LC","LC/APC","LSH","LSH/APC","MPO","MTRJ","SC","SC/APC","ST","CS","SN","SMA 905","SMA 906","URM-P2","URM-P4","URM-P8","Splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTypeTypeLabelPropEnum = append(rearPortTypeTypeLabelPropEnum, v)
	}
}

const (

	// RearPortTypeLabelNr8P8C captures enum value "8P8C"
	RearPortTypeLabelNr8P8C string = "8P8C"

	// RearPortTypeLabelNr8P6C captures enum value "8P6C"
	RearPortTypeLabelNr8P6C string = "8P6C"

	// RearPortTypeLabelNr8P4C captures enum value "8P4C"
	RearPortTypeLabelNr8P4C string = "8P4C"

	// RearPortTypeLabelNr8P2C captures enum value "8P2C"
	RearPortTypeLabelNr8P2C string = "8P2C"

	// RearPortTypeLabelNr6P6C captures enum value "6P6C"
	RearPortTypeLabelNr6P6C string = "6P6C"

	// RearPortTypeLabelNr6P4C captures enum value "6P4C"
	RearPortTypeLabelNr6P4C string = "6P4C"

	// RearPortTypeLabelNr6P2C captures enum value "6P2C"
	RearPortTypeLabelNr6P2C string = "6P2C"

	// RearPortTypeLabelNr4P4C captures enum value "4P4C"
	RearPortTypeLabelNr4P4C string = "4P4C"

	// RearPortTypeLabelNr4P2C captures enum value "4P2C"
	RearPortTypeLabelNr4P2C string = "4P2C"

	// RearPortTypeLabelGG45 captures enum value "GG45"
	RearPortTypeLabelGG45 string = "GG45"

	// RearPortTypeLabelTERA4P captures enum value "TERA 4P"
	RearPortTypeLabelTERA4P string = "TERA 4P"

	// RearPortTypeLabelTERA2P captures enum value "TERA 2P"
	RearPortTypeLabelTERA2P string = "TERA 2P"

	// RearPortTypeLabelTERA1P captures enum value "TERA 1P"
	RearPortTypeLabelTERA1P string = "TERA 1P"

	// RearPortTypeLabelNr110Punch captures enum value "110 Punch"
	RearPortTypeLabelNr110Punch string = "110 Punch"

	// RearPortTypeLabelBNC captures enum value "BNC"
	RearPortTypeLabelBNC string = "BNC"

	// RearPortTypeLabelFConnector captures enum value "F Connector"
	RearPortTypeLabelFConnector string = "F Connector"

	// RearPortTypeLabelNConnector captures enum value "N Connector"
	RearPortTypeLabelNConnector string = "N Connector"

	// RearPortTypeLabelMRJ21 captures enum value "MRJ21"
	RearPortTypeLabelMRJ21 string = "MRJ21"

	// RearPortTypeLabelFC captures enum value "FC"
	RearPortTypeLabelFC string = "FC"

	// RearPortTypeLabelLC captures enum value "LC"
	RearPortTypeLabelLC string = "LC"

	// RearPortTypeLabelLCAPC captures enum value "LC/APC"
	RearPortTypeLabelLCAPC string = "LC/APC"

	// RearPortTypeLabelLSH captures enum value "LSH"
	RearPortTypeLabelLSH string = "LSH"

	// RearPortTypeLabelLSHAPC captures enum value "LSH/APC"
	RearPortTypeLabelLSHAPC string = "LSH/APC"

	// RearPortTypeLabelMPO captures enum value "MPO"
	RearPortTypeLabelMPO string = "MPO"

	// RearPortTypeLabelMTRJ captures enum value "MTRJ"
	RearPortTypeLabelMTRJ string = "MTRJ"

	// RearPortTypeLabelSC captures enum value "SC"
	RearPortTypeLabelSC string = "SC"

	// RearPortTypeLabelSCAPC captures enum value "SC/APC"
	RearPortTypeLabelSCAPC string = "SC/APC"

	// RearPortTypeLabelST captures enum value "ST"
	RearPortTypeLabelST string = "ST"

	// RearPortTypeLabelCS captures enum value "CS"
	RearPortTypeLabelCS string = "CS"

	// RearPortTypeLabelSN captures enum value "SN"
	RearPortTypeLabelSN string = "SN"

	// RearPortTypeLabelSMA905 captures enum value "SMA 905"
	RearPortTypeLabelSMA905 string = "SMA 905"

	// RearPortTypeLabelSMA906 captures enum value "SMA 906"
	RearPortTypeLabelSMA906 string = "SMA 906"

	// RearPortTypeLabelURMDashP2 captures enum value "URM-P2"
	RearPortTypeLabelURMDashP2 string = "URM-P2"

	// RearPortTypeLabelURMDashP4 captures enum value "URM-P4"
	RearPortTypeLabelURMDashP4 string = "URM-P4"

	// RearPortTypeLabelURMDashP8 captures enum value "URM-P8"
	RearPortTypeLabelURMDashP8 string = "URM-P8"

	// RearPortTypeLabelSplice captures enum value "Splice"
	RearPortTypeLabelSplice string = "Splice"
)

// prop value enum
func (m *RearPortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rearPortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","6p6c","6p4c","6p2c","4p4c","4p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","f","n","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st","cs","sn","sma-905","sma-906","urm-p2","urm-p4","urm-p8","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTypeTypeValuePropEnum = append(rearPortTypeTypeValuePropEnum, v)
	}
}

const (

	// RearPortTypeValueNr8p8c captures enum value "8p8c"
	RearPortTypeValueNr8p8c string = "8p8c"

	// RearPortTypeValueNr8p6c captures enum value "8p6c"
	RearPortTypeValueNr8p6c string = "8p6c"

	// RearPortTypeValueNr8p4c captures enum value "8p4c"
	RearPortTypeValueNr8p4c string = "8p4c"

	// RearPortTypeValueNr8p2c captures enum value "8p2c"
	RearPortTypeValueNr8p2c string = "8p2c"

	// RearPortTypeValueNr6p6c captures enum value "6p6c"
	RearPortTypeValueNr6p6c string = "6p6c"

	// RearPortTypeValueNr6p4c captures enum value "6p4c"
	RearPortTypeValueNr6p4c string = "6p4c"

	// RearPortTypeValueNr6p2c captures enum value "6p2c"
	RearPortTypeValueNr6p2c string = "6p2c"

	// RearPortTypeValueNr4p4c captures enum value "4p4c"
	RearPortTypeValueNr4p4c string = "4p4c"

	// RearPortTypeValueNr4p2c captures enum value "4p2c"
	RearPortTypeValueNr4p2c string = "4p2c"

	// RearPortTypeValueGg45 captures enum value "gg45"
	RearPortTypeValueGg45 string = "gg45"

	// RearPortTypeValueTeraDash4p captures enum value "tera-4p"
	RearPortTypeValueTeraDash4p string = "tera-4p"

	// RearPortTypeValueTeraDash2p captures enum value "tera-2p"
	RearPortTypeValueTeraDash2p string = "tera-2p"

	// RearPortTypeValueTeraDash1p captures enum value "tera-1p"
	RearPortTypeValueTeraDash1p string = "tera-1p"

	// RearPortTypeValueNr110DashPunch captures enum value "110-punch"
	RearPortTypeValueNr110DashPunch string = "110-punch"

	// RearPortTypeValueBnc captures enum value "bnc"
	RearPortTypeValueBnc string = "bnc"

	// RearPortTypeValueF captures enum value "f"
	RearPortTypeValueF string = "f"

	// RearPortTypeValueN captures enum value "n"
	RearPortTypeValueN string = "n"

	// RearPortTypeValueMrj21 captures enum value "mrj21"
	RearPortTypeValueMrj21 string = "mrj21"

	// RearPortTypeValueFc captures enum value "fc"
	RearPortTypeValueFc string = "fc"

	// RearPortTypeValueLc captures enum value "lc"
	RearPortTypeValueLc string = "lc"

	// RearPortTypeValueLcDashApc captures enum value "lc-apc"
	RearPortTypeValueLcDashApc string = "lc-apc"

	// RearPortTypeValueLsh captures enum value "lsh"
	RearPortTypeValueLsh string = "lsh"

	// RearPortTypeValueLshDashApc captures enum value "lsh-apc"
	RearPortTypeValueLshDashApc string = "lsh-apc"

	// RearPortTypeValueMpo captures enum value "mpo"
	RearPortTypeValueMpo string = "mpo"

	// RearPortTypeValueMtrj captures enum value "mtrj"
	RearPortTypeValueMtrj string = "mtrj"

	// RearPortTypeValueSc captures enum value "sc"
	RearPortTypeValueSc string = "sc"

	// RearPortTypeValueScDashApc captures enum value "sc-apc"
	RearPortTypeValueScDashApc string = "sc-apc"

	// RearPortTypeValueSt captures enum value "st"
	RearPortTypeValueSt string = "st"

	// RearPortTypeValueCs captures enum value "cs"
	RearPortTypeValueCs string = "cs"

	// RearPortTypeValueSn captures enum value "sn"
	RearPortTypeValueSn string = "sn"

	// RearPortTypeValueSmaDash905 captures enum value "sma-905"
	RearPortTypeValueSmaDash905 string = "sma-905"

	// RearPortTypeValueSmaDash906 captures enum value "sma-906"
	RearPortTypeValueSmaDash906 string = "sma-906"

	// RearPortTypeValueUrmDashP2 captures enum value "urm-p2"
	RearPortTypeValueUrmDashP2 string = "urm-p2"

	// RearPortTypeValueUrmDashP4 captures enum value "urm-p4"
	RearPortTypeValueUrmDashP4 string = "urm-p4"

	// RearPortTypeValueUrmDashP8 captures enum value "urm-p8"
	RearPortTypeValueUrmDashP8 string = "urm-p8"

	// RearPortTypeValueSplice captures enum value "splice"
	RearPortTypeValueSplice string = "splice"
)

// prop value enum
func (m *RearPortType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rear port type based on context it is used
func (m *RearPortType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RearPortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPortType) UnmarshalBinary(b []byte) error {
	var res RearPortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
