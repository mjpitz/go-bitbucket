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

// DeploymentsDdevDeploymentEnvironment deployments ddev deployment environment
//
// swagger:model deployments_ddev_deployment_environment
type DeploymentsDdevDeploymentEnvironment struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	DeploymentsDdevDeploymentEnvironmentAllOf1
}

// Type gets the type of this subtype
func (m *DeploymentsDdevDeploymentEnvironment) Type() string {
	return "deployments_ddev_deployment_environment"
}

// SetType sets the type of this subtype
func (m *DeploymentsDdevDeploymentEnvironment) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DeploymentsDdevDeploymentEnvironment) UnmarshalJSON(raw []byte) error {
	var data struct {
		DeploymentsDdevDeploymentEnvironmentAllOf1
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

	var result DeploymentsDdevDeploymentEnvironment

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}
	result.DeploymentsDdevDeploymentEnvironmentAllOf1 = data.DeploymentsDdevDeploymentEnvironmentAllOf1

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DeploymentsDdevDeploymentEnvironment) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		DeploymentsDdevDeploymentEnvironmentAllOf1
	}{

		DeploymentsDdevDeploymentEnvironmentAllOf1: m.DeploymentsDdevDeploymentEnvironmentAllOf1,
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

// Validate validates this deployments ddev deployment environment
func (m *DeploymentsDdevDeploymentEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DeploymentsDdevDeploymentEnvironmentAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentsDdevDeploymentEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentsDdevDeploymentEnvironment) UnmarshalBinary(b []byte) error {
	var res DeploymentsDdevDeploymentEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeploymentsDdevDeploymentEnvironmentAllOf1 A Bitbucket Deployment Environment.
//
// swagger:model DeploymentsDdevDeploymentEnvironmentAllOf1
type DeploymentsDdevDeploymentEnvironmentAllOf1 interface{}
