// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PipelineStepStateInProgress pipeline step state in progress
//
// swagger:model pipeline_step_state_in_progress
type PipelineStepStateInProgress struct {
	PipelineStepState

	// The name of pipeline step state (IN_PROGRESS).
	// Enum: [IN_PROGRESS]
	Name string `json:"name,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PipelineStepStateInProgress) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PipelineStepState
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PipelineStepState = aO0

	// AO1
	var dataAO1 struct {
		Name string `json:"name,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Name = dataAO1.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PipelineStepStateInProgress) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PipelineStepState)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Name string `json:"name,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}

	dataAO1.Name = m.Name

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pipeline step state in progress
func (m *PipelineStepStateInProgress) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PipelineStepState
	if err := m.PipelineStepState.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pipelineStepStateInProgressTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IN_PROGRESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pipelineStepStateInProgressTypeNamePropEnum = append(pipelineStepStateInProgressTypeNamePropEnum, v)
	}
}

// property enum
func (m *PipelineStepStateInProgress) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pipelineStepStateInProgressTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PipelineStepStateInProgress) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PipelineStepStateInProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineStepStateInProgress) UnmarshalBinary(b []byte) error {
	var res PipelineStepStateInProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
