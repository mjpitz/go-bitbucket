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

// PipelineTarget pipeline target
//
// swagger:model pipeline_target
type PipelineTarget struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	PipelineTargetAllOf1
}

// Type gets the type of this subtype
func (m *PipelineTarget) Type() string {
	return "pipeline_target"
}

// SetType sets the type of this subtype
func (m *PipelineTarget) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PipelineTarget) UnmarshalJSON(raw []byte) error {
	var data struct {
		PipelineTargetAllOf1
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

	var result PipelineTarget

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.PipelineTargetAllOf1 = data.PipelineTargetAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PipelineTarget) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		PipelineTargetAllOf1
	}{

		PipelineTargetAllOf1: m.PipelineTargetAllOf1,
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

// Validate validates this pipeline target
func (m *PipelineTarget) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PipelineTargetAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineTarget) UnmarshalBinary(b []byte) error {
	var res PipelineTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PipelineTargetAllOf1 A representation of the target that a pipeline executes on.
//
// swagger:model PipelineTargetAllOf1
type PipelineTargetAllOf1 interface{}