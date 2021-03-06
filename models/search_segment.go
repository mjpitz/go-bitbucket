// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SearchSegment search segment
//
// swagger:model search_segment
type SearchSegment struct {

	// match
	// Read Only: true
	Match *bool `json:"match,omitempty"`

	// text
	// Read Only: true
	Text string `json:"text,omitempty"`
}

// Validate validates this search segment
func (m *SearchSegment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchSegment) UnmarshalBinary(b []byte) error {
	var res SearchSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
