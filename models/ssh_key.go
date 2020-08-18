// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SSHKey ssh key
//
// swagger:model ssh_key
type SSHKey struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	// The comment parsed from the SSH key (if present)
	Comment string `json:"comment,omitempty"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// The SSH public key value in OpenSSH format.
	Key string `json:"key,omitempty"`

	// The user-defined label for the SSH key
	Label string `json:"label,omitempty"`

	// last used
	// Format: date-time
	LastUsed strfmt.DateTime `json:"last_used,omitempty"`

	// links
	Links *SSHKeyAO1Links `json:"links,omitempty"`

	// The SSH key's immutable ID.
	UUID string `json:"uuid,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this subtype
func (m *SSHKey) Type() string {
	return "ssh_key"
}

// SetType sets the type of this subtype
func (m *SSHKey) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *SSHKey) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The comment parsed from the SSH key (if present)
		Comment string `json:"comment,omitempty"`

		// created on
		// Format: date-time
		CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

		// The SSH public key value in OpenSSH format.
		Key string `json:"key,omitempty"`

		// The user-defined label for the SSH key
		Label string `json:"label,omitempty"`

		// last used
		// Format: date-time
		LastUsed strfmt.DateTime `json:"last_used,omitempty"`

		// links
		Links *SSHKeyAO1Links `json:"links,omitempty"`

		// The SSH key's immutable ID.
		UUID string `json:"uuid,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result SSHKey

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Comment = data.Comment
	result.CreatedOn = data.CreatedOn
	result.Key = data.Key
	result.Label = data.Label
	result.LastUsed = data.LastUsed
	result.Links = data.Links
	result.UUID = data.UUID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m SSHKey) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The comment parsed from the SSH key (if present)
		Comment string `json:"comment,omitempty"`

		// created on
		// Format: date-time
		CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

		// The SSH public key value in OpenSSH format.
		Key string `json:"key,omitempty"`

		// The user-defined label for the SSH key
		Label string `json:"label,omitempty"`

		// last used
		// Format: date-time
		LastUsed strfmt.DateTime `json:"last_used,omitempty"`

		// links
		Links *SSHKeyAO1Links `json:"links,omitempty"`

		// The SSH key's immutable ID.
		UUID string `json:"uuid,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}{

		Comment: m.Comment,

		CreatedOn: m.CreatedOn,

		Key: m.Key,

		Label: m.Label,

		LastUsed: m.LastUsed,

		Links: m.Links,

		UUID: m.UUID,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`
	}{

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this ssh key
func (m *SSHKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUsed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKey) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SSHKey) validateLastUsed(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUsed) { // not required
		return nil
	}

	if err := validate.FormatOf("last_used", "body", "date-time", m.LastUsed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SSHKey) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSHKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHKey) UnmarshalBinary(b []byte) error {
	var res SSHKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SSHKeyAO1Links SSH key a o1 links
//
// swagger:model SSHKeyAO1Links
type SSHKeyAO1Links struct {

	// self
	Self *SSHKeyAO1LinksSelf `json:"self,omitempty"`
}

// Validate validates this SSH key a o1 links
func (m *SSHKeyAO1Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKeyAO1Links) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSHKeyAO1Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHKeyAO1Links) UnmarshalBinary(b []byte) error {
	var res SSHKeyAO1Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SSHKeyAO1LinksSelf SSH key a o1 links self
//
// swagger:model SSHKeyAO1LinksSelf
type SSHKeyAO1LinksSelf struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this SSH key a o1 links self
func (m *SSHKeyAO1LinksSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SSHKeyAO1LinksSelf) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"self"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSHKeyAO1LinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHKeyAO1LinksSelf) UnmarshalBinary(b []byte) error {
	var res SSHKeyAO1LinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}