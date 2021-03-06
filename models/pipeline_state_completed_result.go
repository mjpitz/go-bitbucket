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

// PipelineStateCompletedResult pipeline state completed result
//
// swagger:model pipeline_state_completed_result
type PipelineStateCompletedResult struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	PipelineStateCompletedResultAllOf1
}

// Type gets the type of this subtype
func (m *PipelineStateCompletedResult) Type() string {
	return "pipeline_state_completed_result"
}

// SetType sets the type of this subtype
func (m *PipelineStateCompletedResult) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PipelineStateCompletedResult) UnmarshalJSON(raw []byte) error {
	var data struct {
		PipelineStateCompletedResultAllOf1
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

	var result PipelineStateCompletedResult

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.PipelineStateCompletedResultAllOf1 = data.PipelineStateCompletedResultAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PipelineStateCompletedResult) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		PipelineStateCompletedResultAllOf1
	}{

		PipelineStateCompletedResultAllOf1: m.PipelineStateCompletedResultAllOf1,
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

// Validate validates this pipeline state completed result
func (m *PipelineStateCompletedResult) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PipelineStateCompletedResultAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineStateCompletedResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineStateCompletedResult) UnmarshalBinary(b []byte) error {
	var res PipelineStateCompletedResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PipelineStateCompletedResultAllOf1 A result of a completed pipeline state.
//
// swagger:model PipelineStateCompletedResultAllOf1
type PipelineStateCompletedResultAllOf1 interface{}
