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

// IssueAttachment issue attachment
//
// swagger:model issue_attachment
type IssueAttachment struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	// links
	Links *IssueAttachmentAO1Links `json:"links,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this subtype
func (m *IssueAttachment) Type() string {
	return "issue_attachment"
}

// SetType sets the type of this subtype
func (m *IssueAttachment) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *IssueAttachment) UnmarshalJSON(raw []byte) error {
	var data struct {

		// links
		Links *IssueAttachmentAO1Links `json:"links,omitempty"`

		// name
		Name string `json:"name,omitempty"`

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

	var result IssueAttachment

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Links = data.Links
	result.Name = data.Name

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m IssueAttachment) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// links
		Links *IssueAttachmentAO1Links `json:"links,omitempty"`

		// name
		Name string `json:"name,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}{

		Links: m.Links,

		Name: m.Name,
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

// Validate validates this issue attachment
func (m *IssueAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueAttachment) validateLinks(formats strfmt.Registry) error {

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
func (m *IssueAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueAttachment) UnmarshalBinary(b []byte) error {
	var res IssueAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IssueAttachmentAO1Links issue attachment a o1 links
//
// swagger:model IssueAttachmentAO1Links
type IssueAttachmentAO1Links struct {

	// self
	Self *IssueAttachmentAO1LinksSelf `json:"self,omitempty"`
}

// Validate validates this issue attachment a o1 links
func (m *IssueAttachmentAO1Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueAttachmentAO1Links) validateSelf(formats strfmt.Registry) error {

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
func (m *IssueAttachmentAO1Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueAttachmentAO1Links) UnmarshalBinary(b []byte) error {
	var res IssueAttachmentAO1Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IssueAttachmentAO1LinksSelf issue attachment a o1 links self
//
// swagger:model IssueAttachmentAO1LinksSelf
type IssueAttachmentAO1LinksSelf struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this issue attachment a o1 links self
func (m *IssueAttachmentAO1LinksSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueAttachmentAO1LinksSelf) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"self"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueAttachmentAO1LinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueAttachmentAO1LinksSelf) UnmarshalBinary(b []byte) error {
	var res IssueAttachmentAO1LinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
