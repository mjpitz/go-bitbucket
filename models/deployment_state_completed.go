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

// DeploymentStateCompleted deployment state completed
//
// swagger:model deployment_state_completed
type DeploymentStateCompleted struct {
	DeploymentState

	// The timestamp when the deployment completed.
	// Format: date-time
	CompletionDate strfmt.DateTime `json:"completion_date,omitempty"`

	// The Bitbucket account that was used to perform the deployment.
	Deployer *Account `json:"deployer,omitempty"`

	// The name of deployment state (COMPLETED).
	// Enum: [COMPLETED]
	Name string `json:"name,omitempty"`

	// The timestamp when the deployment was started.
	// Format: date-time
	StartDate strfmt.DateTime `json:"start_date,omitempty"`

	// The status of the completed deployment.
	Status *DeploymentStateCompletedStatus `json:"status,omitempty"`

	// Link to the deployment result.
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DeploymentStateCompleted) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DeploymentState
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DeploymentState = aO0

	// AO1
	var dataAO1 struct {
		CompletionDate strfmt.DateTime `json:"completion_date,omitempty"`

		Deployer *Account `json:"deployer,omitempty"`

		Name string `json:"name,omitempty"`

		StartDate strfmt.DateTime `json:"start_date,omitempty"`

		Status *DeploymentStateCompletedStatus `json:"status,omitempty"`

		URL strfmt.URI `json:"url,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CompletionDate = dataAO1.CompletionDate

	m.Deployer = dataAO1.Deployer

	m.Name = dataAO1.Name

	m.StartDate = dataAO1.StartDate

	m.Status = dataAO1.Status

	m.URL = dataAO1.URL

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DeploymentStateCompleted) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.DeploymentState)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CompletionDate strfmt.DateTime `json:"completion_date,omitempty"`

		Deployer *Account `json:"deployer,omitempty"`

		Name string `json:"name,omitempty"`

		StartDate strfmt.DateTime `json:"start_date,omitempty"`

		Status *DeploymentStateCompletedStatus `json:"status,omitempty"`

		URL strfmt.URI `json:"url,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}

	dataAO1.CompletionDate = m.CompletionDate

	dataAO1.Deployer = m.Deployer

	dataAO1.Name = m.Name

	dataAO1.StartDate = m.StartDate

	dataAO1.Status = m.Status

	dataAO1.URL = m.URL

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this deployment state completed
func (m *DeploymentStateCompleted) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DeploymentState
	if err := m.DeploymentState.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *DeploymentStateCompleted) validateCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("completion_date", "body", "date-time", m.CompletionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentStateCompleted) validateDeployer(formats strfmt.Registry) error {

	if swag.IsZero(m.Deployer) { // not required
		return nil
	}

	if m.Deployer != nil {
		if err := m.Deployer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployer")
			}
			return err
		}
	}

	return nil
}

var deploymentStateCompletedTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COMPLETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentStateCompletedTypeNamePropEnum = append(deploymentStateCompletedTypeNamePropEnum, v)
	}
}

// property enum
func (m *DeploymentStateCompleted) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentStateCompletedTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentStateCompleted) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentStateCompleted) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentStateCompleted) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentStateCompleted) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentStateCompleted) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentStateCompleted) UnmarshalBinary(b []byte) error {
	var res DeploymentStateCompleted
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
