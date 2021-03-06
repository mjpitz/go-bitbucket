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

// PipelineScheduleExecution pipeline schedule execution
//
// swagger:model pipeline_schedule_execution
type PipelineScheduleExecution struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	PipelineScheduleExecutionAllOf1
}

// Type gets the type of this subtype
func (m *PipelineScheduleExecution) Type() string {
	return "pipeline_schedule_execution"
}

// SetType sets the type of this subtype
func (m *PipelineScheduleExecution) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PipelineScheduleExecution) UnmarshalJSON(raw []byte) error {
	var data struct {
		PipelineScheduleExecutionAllOf1
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

	var result PipelineScheduleExecution

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.PipelineScheduleExecutionAllOf1 = data.PipelineScheduleExecutionAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PipelineScheduleExecution) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		PipelineScheduleExecutionAllOf1
	}{

		PipelineScheduleExecutionAllOf1: m.PipelineScheduleExecutionAllOf1,
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

// Validate validates this pipeline schedule execution
func (m *PipelineScheduleExecution) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PipelineScheduleExecutionAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineScheduleExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineScheduleExecution) UnmarshalBinary(b []byte) error {
	var res PipelineScheduleExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PipelineScheduleExecutionAllOf1 A Pipelines schedule execution.
//
// swagger:model PipelineScheduleExecutionAllOf1
type PipelineScheduleExecutionAllOf1 interface{}
