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

// Branchrestriction branchrestriction
//
// swagger:model branchrestriction
type Branchrestriction struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	// Indicates how the restriction is matched against a branch. The default is `glob`.
	// Required: true
	// Enum: [branching_model glob]
	BranchMatchKind *string `json:"branch_match_kind"`

	// Apply the restriction to branches of this type. Active when `branch_match_kind` is `branching_model`. The branch type will be calculated using the branching model configured for the repository.
	// Enum: [feature bugfix release hotfix development production]
	BranchType string `json:"branch_type,omitempty"`

	// groups
	// Min Items: 0
	Groups []*Group `json:"groups"`

	// The branch restriction status' id.
	ID int64 `json:"id,omitempty"`

	// The type of restriction that is being applied.
	// Required: true
	// Enum: [require_tasks_to_be_completed force restrict_merges enforce_merge_checks require_approvals_to_merge allow_auto_merge_when_builds_pass delete require_all_dependencies_merged push require_passing_builds_to_merge reset_pullrequest_approvals_on_change require_default_reviewer_approvals_to_merge]
	Kind *string `json:"kind"`

	// links
	Links *BranchrestrictionAO1Links `json:"links,omitempty"`

	// Apply the restriction to branches that match this pattern. Active when `branch_match_kind` is `glob`. Will be empty when `branch_match_kind` is `branching_model`.
	// Required: true
	Pattern *string `json:"pattern"`

	// users
	// Min Items: 0
	Users []*Account `json:"users"`

	// Value with kind-specific semantics: "require_approvals_to_merge" uses it to require a minimum number of approvals on a PR; "require_default_reviewer_approvals_to_merge" uses it to require a minimum number of approvals from default reviewers on a PR; "require_passing_builds_to_merge" uses it to require a minimum number of passing builds.
	Value int64 `json:"value,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this subtype
func (m *Branchrestriction) Type() string {
	return "branchrestriction"
}

// SetType sets the type of this subtype
func (m *Branchrestriction) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Branchrestriction) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Indicates how the restriction is matched against a branch. The default is `glob`.
		// Required: true
		// Enum: [branching_model glob]
		BranchMatchKind *string `json:"branch_match_kind"`

		// Apply the restriction to branches of this type. Active when `branch_match_kind` is `branching_model`. The branch type will be calculated using the branching model configured for the repository.
		// Enum: [feature bugfix release hotfix development production]
		BranchType string `json:"branch_type,omitempty"`

		// groups
		// Min Items: 0
		Groups []*Group `json:"groups"`

		// The branch restriction status' id.
		ID int64 `json:"id,omitempty"`

		// The type of restriction that is being applied.
		// Required: true
		// Enum: [require_tasks_to_be_completed force restrict_merges enforce_merge_checks require_approvals_to_merge allow_auto_merge_when_builds_pass delete require_all_dependencies_merged push require_passing_builds_to_merge reset_pullrequest_approvals_on_change require_default_reviewer_approvals_to_merge]
		Kind *string `json:"kind"`

		// links
		Links *BranchrestrictionAO1Links `json:"links,omitempty"`

		// Apply the restriction to branches that match this pattern. Active when `branch_match_kind` is `glob`. Will be empty when `branch_match_kind` is `branching_model`.
		// Required: true
		Pattern *string `json:"pattern"`

		// users
		// Min Items: 0
		Users []*Account `json:"users"`

		// Value with kind-specific semantics: "require_approvals_to_merge" uses it to require a minimum number of approvals on a PR; "require_default_reviewer_approvals_to_merge" uses it to require a minimum number of approvals from default reviewers on a PR; "require_passing_builds_to_merge" uses it to require a minimum number of passing builds.
		Value int64 `json:"value,omitempty"`

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

	var result Branchrestriction

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.BranchMatchKind = data.BranchMatchKind
	result.BranchType = data.BranchType
	result.Groups = data.Groups
	result.ID = data.ID
	result.Kind = data.Kind
	result.Links = data.Links
	result.Pattern = data.Pattern
	result.Users = data.Users
	result.Value = data.Value

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Branchrestriction) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Indicates how the restriction is matched against a branch. The default is `glob`.
		// Required: true
		// Enum: [branching_model glob]
		BranchMatchKind *string `json:"branch_match_kind"`

		// Apply the restriction to branches of this type. Active when `branch_match_kind` is `branching_model`. The branch type will be calculated using the branching model configured for the repository.
		// Enum: [feature bugfix release hotfix development production]
		BranchType string `json:"branch_type,omitempty"`

		// groups
		// Min Items: 0
		Groups []*Group `json:"groups"`

		// The branch restriction status' id.
		ID int64 `json:"id,omitempty"`

		// The type of restriction that is being applied.
		// Required: true
		// Enum: [require_tasks_to_be_completed force restrict_merges enforce_merge_checks require_approvals_to_merge allow_auto_merge_when_builds_pass delete require_all_dependencies_merged push require_passing_builds_to_merge reset_pullrequest_approvals_on_change require_default_reviewer_approvals_to_merge]
		Kind *string `json:"kind"`

		// links
		Links *BranchrestrictionAO1Links `json:"links,omitempty"`

		// Apply the restriction to branches that match this pattern. Active when `branch_match_kind` is `glob`. Will be empty when `branch_match_kind` is `branching_model`.
		// Required: true
		Pattern *string `json:"pattern"`

		// users
		// Min Items: 0
		Users []*Account `json:"users"`

		// Value with kind-specific semantics: "require_approvals_to_merge" uses it to require a minimum number of approvals on a PR; "require_default_reviewer_approvals_to_merge" uses it to require a minimum number of approvals from default reviewers on a PR; "require_passing_builds_to_merge" uses it to require a minimum number of passing builds.
		Value int64 `json:"value,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}{

		BranchMatchKind: m.BranchMatchKind,

		BranchType: m.BranchType,

		Groups: m.Groups,

		ID: m.ID,

		Kind: m.Kind,

		Links: m.Links,

		Pattern: m.Pattern,

		Users: m.Users,

		Value: m.Value,
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

// Validate validates this branchrestriction
func (m *Branchrestriction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranchMatchKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBranchType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var branchrestrictionTypeBranchMatchKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["branching_model","glob"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		branchrestrictionTypeBranchMatchKindPropEnum = append(branchrestrictionTypeBranchMatchKindPropEnum, v)
	}
}

// property enum
func (m *Branchrestriction) validateBranchMatchKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, branchrestrictionTypeBranchMatchKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Branchrestriction) validateBranchMatchKind(formats strfmt.Registry) error {

	if err := validate.Required("branch_match_kind", "body", m.BranchMatchKind); err != nil {
		return err
	}

	// value enum
	if err := m.validateBranchMatchKindEnum("branch_match_kind", "body", *m.BranchMatchKind); err != nil {
		return err
	}

	return nil
}

var branchrestrictionTypeBranchTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["feature","bugfix","release","hotfix","development","production"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		branchrestrictionTypeBranchTypePropEnum = append(branchrestrictionTypeBranchTypePropEnum, v)
	}
}

// property enum
func (m *Branchrestriction) validateBranchTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, branchrestrictionTypeBranchTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Branchrestriction) validateBranchType(formats strfmt.Registry) error {

	if swag.IsZero(m.BranchType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBranchTypeEnum("branch_type", "body", m.BranchType); err != nil {
		return err
	}

	return nil
}

func (m *Branchrestriction) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	iGroupsSize := int64(len(m.Groups))

	if err := validate.MinItems("groups", "body", iGroupsSize, 0); err != nil {
		return err
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var branchrestrictionTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["require_tasks_to_be_completed","force","restrict_merges","enforce_merge_checks","require_approvals_to_merge","allow_auto_merge_when_builds_pass","delete","require_all_dependencies_merged","push","require_passing_builds_to_merge","reset_pullrequest_approvals_on_change","require_default_reviewer_approvals_to_merge"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		branchrestrictionTypeKindPropEnum = append(branchrestrictionTypeKindPropEnum, v)
	}
}

// property enum
func (m *Branchrestriction) validateKindEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, branchrestrictionTypeKindPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Branchrestriction) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *Branchrestriction) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *Branchrestriction) validatePattern(formats strfmt.Registry) error {

	if err := validate.Required("pattern", "body", m.Pattern); err != nil {
		return err
	}

	return nil
}

func (m *Branchrestriction) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.Users) { // not required
		return nil
	}

	iUsersSize := int64(len(m.Users))

	if err := validate.MinItems("users", "body", iUsersSize, 0); err != nil {
		return err
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Branchrestriction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Branchrestriction) UnmarshalBinary(b []byte) error {
	var res Branchrestriction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BranchrestrictionAO1Links branchrestriction a o1 links
//
// swagger:model BranchrestrictionAO1Links
type BranchrestrictionAO1Links struct {

	// self
	Self *BranchrestrictionAO1LinksSelf `json:"self,omitempty"`
}

// Validate validates this branchrestriction a o1 links
func (m *BranchrestrictionAO1Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchrestrictionAO1Links) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchrestrictionAO1Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchrestrictionAO1Links) UnmarshalBinary(b []byte) error {
	var res BranchrestrictionAO1Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BranchrestrictionAO1LinksSelf branchrestriction a o1 links self
//
// swagger:model BranchrestrictionAO1LinksSelf
type BranchrestrictionAO1LinksSelf struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this branchrestriction a o1 links self
func (m *BranchrestrictionAO1LinksSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchrestrictionAO1LinksSelf) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"self"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchrestrictionAO1LinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchrestrictionAO1LinksSelf) UnmarshalBinary(b []byte) error {
	var res BranchrestrictionAO1LinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
