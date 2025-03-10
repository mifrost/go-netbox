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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableRearPortTemplate writable rear port template
//
// swagger:model WritableRearPortTemplate
type WritableRearPortTemplate struct {

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
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

	// Module type
	// Required: true
	ModuleType *int64 `json:"module_type"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Positions
	// Maximum: 1024
	// Minimum: 1
	Positions int64 `json:"positions,omitempty"`

	// Type
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c 6p6c 6p4c 6p2c 4p4c 4p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc f n mrj21 fc lc lc-pc lc-upc lc-apc lsh lsh-pc lsh-upc lsh-apc mpo mtrj sc sc-pc sc-upc sc-apc st cs sn sma-905 sma-906 urm-p2 urm-p4 urm-p8 splice other]
	Type *string `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable rear port template
func (m *WritableRearPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModuleType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositions(formats); err != nil {
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

func (m *WritableRearPortTemplate) validateColor(formats strfmt.Registry) error {
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

func (m *WritableRearPortTemplate) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateModuleType(formats strfmt.Registry) error {

	if err := validate.Required("module_type", "body", m.ModuleType); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateName(formats strfmt.Registry) error {

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

func (m *WritableRearPortTemplate) validatePositions(formats strfmt.Registry) error {
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

var writableRearPortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","6p6c","6p4c","6p2c","4p4c","4p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","f","n","mrj21","fc","lc","lc-pc","lc-upc","lc-apc","lsh","lsh-pc","lsh-upc","lsh-apc","mpo","mtrj","sc","sc-pc","sc-upc","sc-apc","st","cs","sn","sma-905","sma-906","urm-p2","urm-p4","urm-p8","splice","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRearPortTemplateTypeTypePropEnum = append(writableRearPortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableRearPortTemplateTypeNr8p8c captures enum value "8p8c"
	WritableRearPortTemplateTypeNr8p8c string = "8p8c"

	// WritableRearPortTemplateTypeNr8p6c captures enum value "8p6c"
	WritableRearPortTemplateTypeNr8p6c string = "8p6c"

	// WritableRearPortTemplateTypeNr8p4c captures enum value "8p4c"
	WritableRearPortTemplateTypeNr8p4c string = "8p4c"

	// WritableRearPortTemplateTypeNr8p2c captures enum value "8p2c"
	WritableRearPortTemplateTypeNr8p2c string = "8p2c"

	// WritableRearPortTemplateTypeNr6p6c captures enum value "6p6c"
	WritableRearPortTemplateTypeNr6p6c string = "6p6c"

	// WritableRearPortTemplateTypeNr6p4c captures enum value "6p4c"
	WritableRearPortTemplateTypeNr6p4c string = "6p4c"

	// WritableRearPortTemplateTypeNr6p2c captures enum value "6p2c"
	WritableRearPortTemplateTypeNr6p2c string = "6p2c"

	// WritableRearPortTemplateTypeNr4p4c captures enum value "4p4c"
	WritableRearPortTemplateTypeNr4p4c string = "4p4c"

	// WritableRearPortTemplateTypeNr4p2c captures enum value "4p2c"
	WritableRearPortTemplateTypeNr4p2c string = "4p2c"

	// WritableRearPortTemplateTypeGg45 captures enum value "gg45"
	WritableRearPortTemplateTypeGg45 string = "gg45"

	// WritableRearPortTemplateTypeTeraDash4p captures enum value "tera-4p"
	WritableRearPortTemplateTypeTeraDash4p string = "tera-4p"

	// WritableRearPortTemplateTypeTeraDash2p captures enum value "tera-2p"
	WritableRearPortTemplateTypeTeraDash2p string = "tera-2p"

	// WritableRearPortTemplateTypeTeraDash1p captures enum value "tera-1p"
	WritableRearPortTemplateTypeTeraDash1p string = "tera-1p"

	// WritableRearPortTemplateTypeNr110DashPunch captures enum value "110-punch"
	WritableRearPortTemplateTypeNr110DashPunch string = "110-punch"

	// WritableRearPortTemplateTypeBnc captures enum value "bnc"
	WritableRearPortTemplateTypeBnc string = "bnc"

	// WritableRearPortTemplateTypeF captures enum value "f"
	WritableRearPortTemplateTypeF string = "f"

	// WritableRearPortTemplateTypeN captures enum value "n"
	WritableRearPortTemplateTypeN string = "n"

	// WritableRearPortTemplateTypeMrj21 captures enum value "mrj21"
	WritableRearPortTemplateTypeMrj21 string = "mrj21"

	// WritableRearPortTemplateTypeFc captures enum value "fc"
	WritableRearPortTemplateTypeFc string = "fc"

	// WritableRearPortTemplateTypeLc captures enum value "lc"
	WritableRearPortTemplateTypeLc string = "lc"

	// WritableRearPortTemplateTypeLcDashPc captures enum value "lc-pc"
	WritableRearPortTemplateTypeLcDashPc string = "lc-pc"

	// WritableRearPortTemplateTypeLcDashUpc captures enum value "lc-upc"
	WritableRearPortTemplateTypeLcDashUpc string = "lc-upc"

	// WritableRearPortTemplateTypeLcDashApc captures enum value "lc-apc"
	WritableRearPortTemplateTypeLcDashApc string = "lc-apc"

	// WritableRearPortTemplateTypeLsh captures enum value "lsh"
	WritableRearPortTemplateTypeLsh string = "lsh"

	// WritableRearPortTemplateTypeLshDashPc captures enum value "lsh-pc"
	WritableRearPortTemplateTypeLshDashPc string = "lsh-pc"

	// WritableRearPortTemplateTypeLshDashUpc captures enum value "lsh-upc"
	WritableRearPortTemplateTypeLshDashUpc string = "lsh-upc"

	// WritableRearPortTemplateTypeLshDashApc captures enum value "lsh-apc"
	WritableRearPortTemplateTypeLshDashApc string = "lsh-apc"

	// WritableRearPortTemplateTypeMpo captures enum value "mpo"
	WritableRearPortTemplateTypeMpo string = "mpo"

	// WritableRearPortTemplateTypeMtrj captures enum value "mtrj"
	WritableRearPortTemplateTypeMtrj string = "mtrj"

	// WritableRearPortTemplateTypeSc captures enum value "sc"
	WritableRearPortTemplateTypeSc string = "sc"

	// WritableRearPortTemplateTypeScDashPc captures enum value "sc-pc"
	WritableRearPortTemplateTypeScDashPc string = "sc-pc"

	// WritableRearPortTemplateTypeScDashUpc captures enum value "sc-upc"
	WritableRearPortTemplateTypeScDashUpc string = "sc-upc"

	// WritableRearPortTemplateTypeScDashApc captures enum value "sc-apc"
	WritableRearPortTemplateTypeScDashApc string = "sc-apc"

	// WritableRearPortTemplateTypeSt captures enum value "st"
	WritableRearPortTemplateTypeSt string = "st"

	// WritableRearPortTemplateTypeCs captures enum value "cs"
	WritableRearPortTemplateTypeCs string = "cs"

	// WritableRearPortTemplateTypeSn captures enum value "sn"
	WritableRearPortTemplateTypeSn string = "sn"

	// WritableRearPortTemplateTypeSmaDash905 captures enum value "sma-905"
	WritableRearPortTemplateTypeSmaDash905 string = "sma-905"

	// WritableRearPortTemplateTypeSmaDash906 captures enum value "sma-906"
	WritableRearPortTemplateTypeSmaDash906 string = "sma-906"

	// WritableRearPortTemplateTypeUrmDashP2 captures enum value "urm-p2"
	WritableRearPortTemplateTypeUrmDashP2 string = "urm-p2"

	// WritableRearPortTemplateTypeUrmDashP4 captures enum value "urm-p4"
	WritableRearPortTemplateTypeUrmDashP4 string = "urm-p4"

	// WritableRearPortTemplateTypeUrmDashP8 captures enum value "urm-p8"
	WritableRearPortTemplateTypeUrmDashP8 string = "urm-p8"

	// WritableRearPortTemplateTypeSplice captures enum value "splice"
	WritableRearPortTemplateTypeSplice string = "splice"

	// WritableRearPortTemplateTypeOther captures enum value "other"
	WritableRearPortTemplateTypeOther string = "other"
)

// prop value enum
func (m *WritableRearPortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableRearPortTemplateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableRearPortTemplate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable rear port template based on the context it is used
func (m *WritableRearPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
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

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRearPortTemplate) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableRearPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableRearPortTemplate) UnmarshalBinary(b []byte) error {
	var res WritableRearPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
