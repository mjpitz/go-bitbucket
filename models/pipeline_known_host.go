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

// PipelineKnownHost pipeline known host
//
// swagger:model pipeline_known_host
type PipelineKnownHost struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	// The hostname of the known host.
	Hostname string `json:"hostname,omitempty"`

	// The public key of the known host.
	PublicKey *PipelineSSHPublicKey `json:"public_key,omitempty"`

	// The UUID identifying the known host.
	UUID string `json:"uuid,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this subtype
func (m *PipelineKnownHost) Type() string {
	return "pipeline_known_host"
}

// SetType sets the type of this subtype
func (m *PipelineKnownHost) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PipelineKnownHost) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The hostname of the known host.
		Hostname string `json:"hostname,omitempty"`

		// The public key of the known host.
		PublicKey *PipelineSSHPublicKey `json:"public_key,omitempty"`

		// The UUID identifying the known host.
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

	var result PipelineKnownHost

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Hostname = data.Hostname
	result.PublicKey = data.PublicKey
	result.UUID = data.UUID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PipelineKnownHost) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The hostname of the known host.
		Hostname string `json:"hostname,omitempty"`

		// The public key of the known host.
		PublicKey *PipelineSSHPublicKey `json:"public_key,omitempty"`

		// The UUID identifying the known host.
		UUID string `json:"uuid,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}{

		Hostname: m.Hostname,

		PublicKey: m.PublicKey,

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

// Validate validates this pipeline known host
func (m *PipelineKnownHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PipelineKnownHost) validatePublicKey(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicKey) { // not required
		return nil
	}

	if m.PublicKey != nil {
		if err := m.PublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PipelineKnownHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineKnownHost) UnmarshalBinary(b []byte) error {
	var res PipelineKnownHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}