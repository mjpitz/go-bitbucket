// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SSHAccountKey ssh account key
//
// swagger:model ssh_account_key
type SSHAccountKey struct {
	SSHKey

	// owner
	Owner *Account `json:"owner,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SSHAccountKey) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SSHKey
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SSHKey = aO0

	// AO1
	var dataAO1 struct {
		Owner *Account `json:"owner,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Owner = dataAO1.Owner

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SSHAccountKey) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SSHKey)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Owner *Account `json:"owner,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}

	dataAO1.Owner = m.Owner

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ssh account key
func (m *SSHAccountKey) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SSHKey
	if err := m.SSHKey.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHAccountKey) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSHAccountKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHAccountKey) UnmarshalBinary(b []byte) error {
	var res SSHAccountKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
