// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IssueComment issue comment
//
// swagger:model issue_comment
type IssueComment struct {
	Comment

	// issue
	Issue *Issue `json:"issue,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *IssueComment) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Comment
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Comment = aO0

	// AO1
	var dataAO1 struct {
		Issue *Issue `json:"issue,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Issue = dataAO1.Issue

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m IssueComment) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Comment)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Issue *Issue `json:"issue,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}

	dataAO1.Issue = m.Issue

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this issue comment
func (m *IssueComment) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Comment
	if err := m.Comment.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueComment) validateIssue(formats strfmt.Registry) error {

	if swag.IsZero(m.Issue) { // not required
		return nil
	}

	if m.Issue != nil {
		if err := m.Issue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueComment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueComment) UnmarshalBinary(b []byte) error {
	var res IssueComment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
