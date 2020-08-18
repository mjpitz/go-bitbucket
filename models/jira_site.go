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
)

// JiraSite jira site
//
// swagger:model jira_site
type JiraSite struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	JiraSiteAllOf1
}

// Type gets the type of this subtype
func (m *JiraSite) Type() string {
	return "jira_site"
}

// SetType sets the type of this subtype
func (m *JiraSite) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *JiraSite) UnmarshalJSON(raw []byte) error {
	var data struct {
		JiraSiteAllOf1
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

	var result JiraSite

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.JiraSiteAllOf1 = data.JiraSiteAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m JiraSite) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		JiraSiteAllOf1
	}{

		JiraSiteAllOf1: m.JiraSiteAllOf1,
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

// Validate validates this jira site
func (m *JiraSite) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with JiraSiteAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *JiraSite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JiraSite) UnmarshalBinary(b []byte) error {
	var res JiraSite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JiraSiteAllOf1 A Jira Site.
//
// swagger:model JiraSiteAllOf1
type JiraSiteAllOf1 interface{}
