// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PipelineCommand An executable pipeline command.
//
// swagger:model pipeline_command
type PipelineCommand struct {

	// The executable command.
	Command string `json:"command,omitempty"`

	// The name of the command.
	Name string `json:"name,omitempty"`
}

// Validate validates this pipeline command
func (m *PipelineCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineCommand) UnmarshalBinary(b []byte) error {
	var res PipelineCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}