// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PipelineTriggerManual pipeline trigger manual
//
// swagger:model pipeline_trigger_manual
type PipelineTriggerManual struct {
	PipelineTrigger

	PipelineTriggerManualAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PipelineTriggerManual) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PipelineTrigger
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PipelineTrigger = aO0

	// AO1
	var aO1 PipelineTriggerManualAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PipelineTriggerManualAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PipelineTriggerManual) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PipelineTrigger)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PipelineTriggerManualAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pipeline trigger manual
func (m *PipelineTriggerManual) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PipelineTrigger
	if err := m.PipelineTrigger.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PipelineTriggerManualAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineTriggerManual) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineTriggerManual) UnmarshalBinary(b []byte) error {
	var res PipelineTriggerManual
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PipelineTriggerManualAllOf1 A Bitbucket Pipelines MANUAL trigger.
//
// swagger:model PipelineTriggerManualAllOf1
type PipelineTriggerManualAllOf1 interface{}
