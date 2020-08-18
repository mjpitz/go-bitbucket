// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BranchingModel branching model
//
// swagger:model branching_model
type BranchingModel struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	// The active branch types.
	// Max Items: 4
	// Min Items: 0
	// Unique: true
	BranchTypes []*BranchingModelBranchTypesItems0 `json:"branch_types"`

	// development
	Development *BranchingModelAO1Development `json:"development,omitempty"`

	// production
	Production *BranchingModelAO1Production `json:"production,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this subtype
func (m *BranchingModel) Type() string {
	return "branching_model"
}

// SetType sets the type of this subtype
func (m *BranchingModel) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *BranchingModel) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The active branch types.
		// Max Items: 4
		// Min Items: 0
		// Unique: true
		BranchTypes []*BranchingModelBranchTypesItems0 `json:"branch_types"`

		// development
		Development *BranchingModelAO1Development `json:"development,omitempty"`

		// production
		Production *BranchingModelAO1Production `json:"production,omitempty"`

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

	var result BranchingModel

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.BranchTypes = data.BranchTypes
	result.Development = data.Development
	result.Production = data.Production

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m BranchingModel) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The active branch types.
		// Max Items: 4
		// Min Items: 0
		// Unique: true
		BranchTypes []*BranchingModelBranchTypesItems0 `json:"branch_types"`

		// development
		Development *BranchingModelAO1Development `json:"development,omitempty"`

		// production
		Production *BranchingModelAO1Production `json:"production,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}{

		BranchTypes: m.BranchTypes,

		Development: m.Development,

		Production: m.Production,
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

// Validate validates this branching model
func (m *BranchingModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranchTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevelopment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchingModel) validateBranchTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.BranchTypes) { // not required
		return nil
	}

	iBranchTypesSize := int64(len(m.BranchTypes))

	if err := validate.MinItems("branch_types", "body", iBranchTypesSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("branch_types", "body", iBranchTypesSize, 4); err != nil {
		return err
	}

	if err := validate.UniqueItems("branch_types", "body", m.BranchTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.BranchTypes); i++ {
		if swag.IsZero(m.BranchTypes[i]) { // not required
			continue
		}

		if m.BranchTypes[i] != nil {
			if err := m.BranchTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("branch_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BranchingModel) validateDevelopment(formats strfmt.Registry) error {

	if swag.IsZero(m.Development) { // not required
		return nil
	}

	if m.Development != nil {
		if err := m.Development.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("development")
			}
			return err
		}
	}

	return nil
}

func (m *BranchingModel) validateProduction(formats strfmt.Registry) error {

	if swag.IsZero(m.Production) { // not required
		return nil
	}

	if m.Production != nil {
		if err := m.Production.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("production")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchingModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchingModel) UnmarshalBinary(b []byte) error {
	var res BranchingModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BranchingModelAO1Development branching model a o1 development
//
// swagger:model BranchingModelAO1Development
type BranchingModelAO1Development struct {

	// branch
	Branch *Branch `json:"branch,omitempty"`

	// Name of the target branch. Will be listed here even when the target branch does not exist. Will be `null` if targeting the main branch and the repository is empty.
	// Required: true
	Name *string `json:"name"`

	// Indicates if the setting points at an explicit branch (`false`) or tracks the main branch (`true`).
	// Required: true
	UseMainbranch *bool `json:"use_mainbranch"`
}

// Validate validates this branching model a o1 development
func (m *BranchingModelAO1Development) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseMainbranch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchingModelAO1Development) validateBranch(formats strfmt.Registry) error {

	if swag.IsZero(m.Branch) { // not required
		return nil
	}

	if m.Branch != nil {
		if err := m.Branch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("development" + "." + "branch")
			}
			return err
		}
	}

	return nil
}

func (m *BranchingModelAO1Development) validateName(formats strfmt.Registry) error {

	if err := validate.Required("development"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BranchingModelAO1Development) validateUseMainbranch(formats strfmt.Registry) error {

	if err := validate.Required("development"+"."+"use_mainbranch", "body", m.UseMainbranch); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchingModelAO1Development) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchingModelAO1Development) UnmarshalBinary(b []byte) error {
	var res BranchingModelAO1Development
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BranchingModelAO1Production branching model a o1 production
//
// swagger:model BranchingModelAO1Production
type BranchingModelAO1Production struct {

	// branch
	Branch *Branch `json:"branch,omitempty"`

	// Name of the target branch. Will be listed here even when the target branch does not exist. Will be `null` if targeting the main branch and the repository is empty.
	// Required: true
	Name *string `json:"name"`

	// Indicates if the setting points at an explicit branch (`false`) or tracks the main branch (`true`).
	// Required: true
	UseMainbranch *bool `json:"use_mainbranch"`
}

// Validate validates this branching model a o1 production
func (m *BranchingModelAO1Production) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseMainbranch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchingModelAO1Production) validateBranch(formats strfmt.Registry) error {

	if swag.IsZero(m.Branch) { // not required
		return nil
	}

	if m.Branch != nil {
		if err := m.Branch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("production" + "." + "branch")
			}
			return err
		}
	}

	return nil
}

func (m *BranchingModelAO1Production) validateName(formats strfmt.Registry) error {

	if err := validate.Required("production"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BranchingModelAO1Production) validateUseMainbranch(formats strfmt.Registry) error {

	if err := validate.Required("production"+"."+"use_mainbranch", "body", m.UseMainbranch); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchingModelAO1Production) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchingModelAO1Production) UnmarshalBinary(b []byte) error {
	var res BranchingModelAO1Production
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BranchingModelBranchTypesItems0 branching model branch types items0
//
// swagger:model BranchingModelBranchTypesItems0
type BranchingModelBranchTypesItems0 struct {

	// The kind of branch.
	// Required: true
	// Enum: [feature bugfix release hotfix]
	Kind *string `json:"kind"`

	// The prefix for this branch type. A branch with this prefix will be classified as per `kind`. The prefix must be a valid prefix for a branch and must always exist. It cannot be blank, empty or `null`.
	// Required: true
	Prefix *string `json:"prefix"`
}

// Validate validates this branching model branch types items0
func (m *BranchingModelBranchTypesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var branchingModelBranchTypesItems0TypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["feature","bugfix","release","hotfix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		branchingModelBranchTypesItems0TypeKindPropEnum = append(branchingModelBranchTypesItems0TypeKindPropEnum, v)
	}
}

const (

	// BranchingModelBranchTypesItems0KindFeature captures enum value "feature"
	BranchingModelBranchTypesItems0KindFeature string = "feature"

	// BranchingModelBranchTypesItems0KindBugfix captures enum value "bugfix"
	BranchingModelBranchTypesItems0KindBugfix string = "bugfix"

	// BranchingModelBranchTypesItems0KindRelease captures enum value "release"
	BranchingModelBranchTypesItems0KindRelease string = "release"

	// BranchingModelBranchTypesItems0KindHotfix captures enum value "hotfix"
	BranchingModelBranchTypesItems0KindHotfix string = "hotfix"
)

// prop value enum
func (m *BranchingModelBranchTypesItems0) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, branchingModelBranchTypesItems0TypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BranchingModelBranchTypesItems0) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *BranchingModelBranchTypesItems0) validatePrefix(formats strfmt.Registry) error {

	if err := validate.Required("prefix", "body", m.Prefix); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchingModelBranchTypesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchingModelBranchTypesItems0) UnmarshalBinary(b []byte) error {
	var res BranchingModelBranchTypesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}