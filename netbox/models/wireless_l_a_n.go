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

// WirelessLAN wireless l a n
//
// swagger:model WirelessLAN
type WirelessLAN struct {

	// auth cipher
	AuthCipher *WirelessLANAuthCipher `json:"auth_cipher,omitempty"`

	// Pre-shared key
	// Max Length: 64
	AuthPsk string `json:"auth_psk,omitempty"`

	// auth type
	AuthType *WirelessLANAuthType `json:"auth_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// group
	Group *NestedWirelessLANGroup `json:"group,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// SSID
	// Required: true
	// Max Length: 32
	// Min Length: 1
	Ssid *string `json:"ssid"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// vlan
	Vlan *NestedVLAN `json:"vlan,omitempty"`
}

// Validate validates this wireless l a n
func (m *WirelessLAN) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthCipher(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthPsk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WirelessLAN) validateAuthCipher(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthCipher) { // not required
		return nil
	}

	if m.AuthCipher != nil {
		if err := m.AuthCipher.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_cipher")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_cipher")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLAN) validateAuthPsk(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthPsk) { // not required
		return nil
	}

	if err := validate.MaxLength("auth_psk", "body", m.AuthPsk, 64); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) validateAuthType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthType) { // not required
		return nil
	}

	if m.AuthType != nil {
		if err := m.AuthType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_type")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLAN) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLAN) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) validateSsid(formats strfmt.Registry) error {

	if err := validate.Required("ssid", "body", m.Ssid); err != nil {
		return err
	}

	if err := validate.MinLength("ssid", "body", *m.Ssid, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("ssid", "body", *m.Ssid, 32); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) validateTags(formats strfmt.Registry) error {
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

func (m *WirelessLAN) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) validateVlan(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlan) { // not required
		return nil
	}

	if m.Vlan != nil {
		if err := m.Vlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this wireless l a n based on the context it is used
func (m *WirelessLAN) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthCipher(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroup(ctx, formats); err != nil {
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

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WirelessLAN) contextValidateAuthCipher(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthCipher != nil {
		if err := m.AuthCipher.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_cipher")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_cipher")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLAN) contextValidateAuthType(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthType != nil {
		if err := m.AuthType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_type")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLAN) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {
		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *WirelessLAN) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WirelessLAN) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *WirelessLAN) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlan != nil {
		if err := m.Vlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WirelessLAN) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WirelessLAN) UnmarshalBinary(b []byte) error {
	var res WirelessLAN
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WirelessLANAuthCipher Auth cipher
//
// swagger:model WirelessLANAuthCipher
type WirelessLANAuthCipher struct {

	// label
	// Required: true
	// Enum: [Auto TKIP AES]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [auto tkip aes]
	Value *string `json:"value"`
}

// Validate validates this wireless l a n auth cipher
func (m *WirelessLANAuthCipher) Validate(formats strfmt.Registry) error {
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

var wirelessLANAuthCipherTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Auto","TKIP","AES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLANAuthCipherTypeLabelPropEnum = append(wirelessLANAuthCipherTypeLabelPropEnum, v)
	}
}

const (

	// WirelessLANAuthCipherLabelAuto captures enum value "Auto"
	WirelessLANAuthCipherLabelAuto string = "Auto"

	// WirelessLANAuthCipherLabelTKIP captures enum value "TKIP"
	WirelessLANAuthCipherLabelTKIP string = "TKIP"

	// WirelessLANAuthCipherLabelAES captures enum value "AES"
	WirelessLANAuthCipherLabelAES string = "AES"
)

// prop value enum
func (m *WirelessLANAuthCipher) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLANAuthCipherTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLANAuthCipher) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("auth_cipher"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("auth_cipher"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var wirelessLANAuthCipherTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","tkip","aes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLANAuthCipherTypeValuePropEnum = append(wirelessLANAuthCipherTypeValuePropEnum, v)
	}
}

const (

	// WirelessLANAuthCipherValueAuto captures enum value "auto"
	WirelessLANAuthCipherValueAuto string = "auto"

	// WirelessLANAuthCipherValueTkip captures enum value "tkip"
	WirelessLANAuthCipherValueTkip string = "tkip"

	// WirelessLANAuthCipherValueAes captures enum value "aes"
	WirelessLANAuthCipherValueAes string = "aes"
)

// prop value enum
func (m *WirelessLANAuthCipher) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLANAuthCipherTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLANAuthCipher) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("auth_cipher"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("auth_cipher"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wireless l a n auth cipher based on context it is used
func (m *WirelessLANAuthCipher) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WirelessLANAuthCipher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WirelessLANAuthCipher) UnmarshalBinary(b []byte) error {
	var res WirelessLANAuthCipher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WirelessLANAuthType Auth type
//
// swagger:model WirelessLANAuthType
type WirelessLANAuthType struct {

	// label
	// Required: true
	// Enum: [Open WEP WPA Personal (PSK) WPA Enterprise]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [open wep wpa-personal wpa-enterprise]
	Value *string `json:"value"`
}

// Validate validates this wireless l a n auth type
func (m *WirelessLANAuthType) Validate(formats strfmt.Registry) error {
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

var wirelessLANAuthTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Open","WEP","WPA Personal (PSK)","WPA Enterprise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLANAuthTypeTypeLabelPropEnum = append(wirelessLANAuthTypeTypeLabelPropEnum, v)
	}
}

const (

	// WirelessLANAuthTypeLabelOpen captures enum value "Open"
	WirelessLANAuthTypeLabelOpen string = "Open"

	// WirelessLANAuthTypeLabelWEP captures enum value "WEP"
	WirelessLANAuthTypeLabelWEP string = "WEP"

	// WirelessLANAuthTypeLabelWPAPersonalPSK captures enum value "WPA Personal (PSK)"
	WirelessLANAuthTypeLabelWPAPersonalPSK string = "WPA Personal (PSK)"

	// WirelessLANAuthTypeLabelWPAEnterprise captures enum value "WPA Enterprise"
	WirelessLANAuthTypeLabelWPAEnterprise string = "WPA Enterprise"
)

// prop value enum
func (m *WirelessLANAuthType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLANAuthTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLANAuthType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("auth_type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("auth_type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var wirelessLANAuthTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["open","wep","wpa-personal","wpa-enterprise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wirelessLANAuthTypeTypeValuePropEnum = append(wirelessLANAuthTypeTypeValuePropEnum, v)
	}
}

const (

	// WirelessLANAuthTypeValueOpen captures enum value "open"
	WirelessLANAuthTypeValueOpen string = "open"

	// WirelessLANAuthTypeValueWep captures enum value "wep"
	WirelessLANAuthTypeValueWep string = "wep"

	// WirelessLANAuthTypeValueWpaDashPersonal captures enum value "wpa-personal"
	WirelessLANAuthTypeValueWpaDashPersonal string = "wpa-personal"

	// WirelessLANAuthTypeValueWpaDashEnterprise captures enum value "wpa-enterprise"
	WirelessLANAuthTypeValueWpaDashEnterprise string = "wpa-enterprise"
)

// prop value enum
func (m *WirelessLANAuthType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wirelessLANAuthTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WirelessLANAuthType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("auth_type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("auth_type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wireless l a n auth type based on context it is used
func (m *WirelessLANAuthType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WirelessLANAuthType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WirelessLANAuthType) UnmarshalBinary(b []byte) error {
	var res WirelessLANAuthType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
