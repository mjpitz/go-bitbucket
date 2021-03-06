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

// Diffstat A diffstat object that includes a summary of changes made to a file between two commits.
//
// swagger:model diffstat
type Diffstat struct {

	// lines added
	LinesAdded int64 `json:"lines_added,omitempty"`

	// lines removed
	LinesRemoved int64 `json:"lines_removed,omitempty"`

	// new
	New *CommitFile `json:"new,omitempty"`

	// old
	Old *CommitFile `json:"old,omitempty"`

	// status
	// Enum: [added removed modified renamed]
	Status string `json:"status,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`

	// diffstat additional properties
	DiffstatAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *Diffstat) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// lines added
		LinesAdded int64 `json:"lines_added,omitempty"`

		// lines removed
		LinesRemoved int64 `json:"lines_removed,omitempty"`

		// new
		New *CommitFile `json:"new,omitempty"`

		// old
		Old *CommitFile `json:"old,omitempty"`

		// status
		// Enum: [added removed modified renamed]
		Status string `json:"status,omitempty"`

		// type
		// Required: true
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv Diffstat

	rcv.LinesAdded = stage1.LinesAdded
	rcv.LinesRemoved = stage1.LinesRemoved
	rcv.New = stage1.New
	rcv.Old = stage1.Old
	rcv.Status = stage1.Status
	rcv.Type = stage1.Type
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "lines_added")
	delete(stage2, "lines_removed")
	delete(stage2, "new")
	delete(stage2, "old")
	delete(stage2, "status")
	delete(stage2, "type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.DiffstatAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m Diffstat) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// lines added
		LinesAdded int64 `json:"lines_added,omitempty"`

		// lines removed
		LinesRemoved int64 `json:"lines_removed,omitempty"`

		// new
		New *CommitFile `json:"new,omitempty"`

		// old
		Old *CommitFile `json:"old,omitempty"`

		// status
		// Enum: [added removed modified renamed]
		Status string `json:"status,omitempty"`

		// type
		// Required: true
		Type *string `json:"type"`
	}

	stage1.LinesAdded = m.LinesAdded
	stage1.LinesRemoved = m.LinesRemoved
	stage1.New = m.New
	stage1.Old = m.Old
	stage1.Status = m.Status
	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.DiffstatAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.DiffstatAdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this diffstat
func (m *Diffstat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNew(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOld(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Diffstat) validateNew(formats strfmt.Registry) error {

	if swag.IsZero(m.New) { // not required
		return nil
	}

	if m.New != nil {
		if err := m.New.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("new")
			}
			return err
		}
	}

	return nil
}

func (m *Diffstat) validateOld(formats strfmt.Registry) error {

	if swag.IsZero(m.Old) { // not required
		return nil
	}

	if m.Old != nil {
		if err := m.Old.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("old")
			}
			return err
		}
	}

	return nil
}

var diffstatTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["added","removed","modified","renamed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diffstatTypeStatusPropEnum = append(diffstatTypeStatusPropEnum, v)
	}
}

const (

	// DiffstatStatusAdded captures enum value "added"
	DiffstatStatusAdded string = "added"

	// DiffstatStatusRemoved captures enum value "removed"
	DiffstatStatusRemoved string = "removed"

	// DiffstatStatusModified captures enum value "modified"
	DiffstatStatusModified string = "modified"

	// DiffstatStatusRenamed captures enum value "renamed"
	DiffstatStatusRenamed string = "renamed"
)

// prop value enum
func (m *Diffstat) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, diffstatTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Diffstat) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Diffstat) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Diffstat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Diffstat) UnmarshalBinary(b []byte) error {
	var res Diffstat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
