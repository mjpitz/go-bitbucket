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
	"github.com/go-openapi/validate"
)

// DeploymentRelease deployment release
//
// swagger:model deployment_release
type DeploymentRelease struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	// The commit associated with this release.
	Commit *Commit `json:"commit,omitempty"`

	// The timestamp when the release was created.
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// The name of the release.
	Name string `json:"name,omitempty"`

	// Link to the pipeline that produced the release.
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// The UUID identifying the release.
	UUID string `json:"uuid,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this subtype
func (m *DeploymentRelease) Type() string {
	return "deployment_release"
}

// SetType sets the type of this subtype
func (m *DeploymentRelease) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DeploymentRelease) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The commit associated with this release.
		Commit *Commit `json:"commit,omitempty"`

		// The timestamp when the release was created.
		// Format: date-time
		CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

		// The name of the release.
		Name string `json:"name,omitempty"`

		// Link to the pipeline that produced the release.
		// Format: uri
		URL strfmt.URI `json:"url,omitempty"`

		// The UUID identifying the release.
		UUID string `json:"uuid,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
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

	var result DeploymentRelease

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Commit = data.Commit
	result.CreatedOn = data.CreatedOn
	result.Name = data.Name
	result.URL = data.URL
	result.UUID = data.UUID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DeploymentRelease) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The commit associated with this release.
		Commit *Commit `json:"commit,omitempty"`

		// The timestamp when the release was created.
		// Format: date-time
		CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

		// The name of the release.
		Name string `json:"name,omitempty"`

		// Link to the pipeline that produced the release.
		// Format: uri
		URL strfmt.URI `json:"url,omitempty"`

		// The UUID identifying the release.
		UUID string `json:"uuid,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}{

		Commit: m.Commit,

		CreatedOn: m.CreatedOn,

		Name: m.Name,

		URL: m.URL,

		UUID: m.UUID,
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

// Validate validates this deployment release
func (m *DeploymentRelease) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentRelease) validateCommit(formats strfmt.Registry) error {

	if swag.IsZero(m.Commit) { // not required
		return nil
	}

	if m.Commit != nil {
		if err := m.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentRelease) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentRelease) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentRelease) UnmarshalBinary(b []byte) error {
	var res DeploymentRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}