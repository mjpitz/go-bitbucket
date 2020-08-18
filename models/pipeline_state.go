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

// PipelineState pipeline state
//
// swagger:model pipeline_state
type PipelineState struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	PipelineStateAllOf1
}

// Type gets the type of this subtype
func (m *PipelineState) Type() string {
	return "pipeline_state"
}

// SetType sets the type of this subtype
func (m *PipelineState) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PipelineState) UnmarshalJSON(raw []byte) error {
	var data struct {
		PipelineStateAllOf1
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

	var result PipelineState

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.PipelineStateAllOf1 = data.PipelineStateAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PipelineState) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		PipelineStateAllOf1
	}{

		PipelineStateAllOf1: m.PipelineStateAllOf1,
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

// Validate validates this pipeline state
func (m *PipelineState) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PipelineStateAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineState) UnmarshalBinary(b []byte) error {
	var res PipelineState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PipelineStateAllOf1 The representation of the progress state of a pipeline.
//
// swagger:model PipelineStateAllOf1
type PipelineStateAllOf1 interface{}
