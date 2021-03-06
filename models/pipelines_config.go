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

// PipelinesConfig pipelines config
//
// swagger:model pipelines_config
type PipelinesConfig struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	// Whether Pipelines is enabled for the repository.
	Enabled bool `json:"enabled,omitempty"`

	// repository
	Repository *Repository `json:"repository,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this subtype
func (m *PipelinesConfig) Type() string {
	return "pipelines_config"
}

// SetType sets the type of this subtype
func (m *PipelinesConfig) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PipelinesConfig) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Whether Pipelines is enabled for the repository.
		Enabled bool `json:"enabled,omitempty"`

		// repository
		Repository *Repository `json:"repository,omitempty"`

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

	var result PipelinesConfig

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Enabled = data.Enabled
	result.Repository = data.Repository

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PipelinesConfig) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Whether Pipelines is enabled for the repository.
		Enabled bool `json:"enabled,omitempty"`

		// repository
		Repository *Repository `json:"repository,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}{

		Enabled: m.Enabled,

		Repository: m.Repository,
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

// Validate validates this pipelines config
func (m *PipelinesConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PipelinesConfig) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PipelinesConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelinesConfig) UnmarshalBinary(b []byte) error {
	var res PipelinesConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
