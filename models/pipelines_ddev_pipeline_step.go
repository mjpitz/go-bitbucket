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

// PipelinesDdevPipelineStep pipelines ddev pipeline step
//
// swagger:model pipelines_ddev_pipeline_step
type PipelinesDdevPipelineStep struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	PipelinesDdevPipelineStepAllOf1
}

// Type gets the type of this subtype
func (m *PipelinesDdevPipelineStep) Type() string {
	return "pipelines_ddev_pipeline_step"
}

// SetType sets the type of this subtype
func (m *PipelinesDdevPipelineStep) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PipelinesDdevPipelineStep) UnmarshalJSON(raw []byte) error {
	var data struct {
		PipelinesDdevPipelineStepAllOf1
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

	var result PipelinesDdevPipelineStep

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.PipelinesDdevPipelineStepAllOf1 = data.PipelinesDdevPipelineStepAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PipelinesDdevPipelineStep) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		PipelinesDdevPipelineStepAllOf1
	}{

		PipelinesDdevPipelineStepAllOf1: m.PipelinesDdevPipelineStepAllOf1,
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

// Validate validates this pipelines ddev pipeline step
func (m *PipelinesDdevPipelineStep) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PipelinesDdevPipelineStepAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PipelinesDdevPipelineStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelinesDdevPipelineStep) UnmarshalBinary(b []byte) error {
	var res PipelinesDdevPipelineStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PipelinesDdevPipelineStepAllOf1 A step of a Bitbucket pipeline. This represents the actual result of the step execution.
//
// swagger:model PipelinesDdevPipelineStepAllOf1
type PipelinesDdevPipelineStepAllOf1 interface{}
