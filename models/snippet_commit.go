// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnippetCommit snippet commit
//
// swagger:model snippet_commit
type SnippetCommit struct {
	BaseCommit

	// links
	Links *SnippetCommitAO1Links `json:"links,omitempty"`

	// snippet
	Snippet *Snippet `json:"snippet,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SnippetCommit) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseCommit
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseCommit = aO0

	// AO1
	var dataAO1 struct {
		Links *SnippetCommitAO1Links `json:"links,omitempty"`

		Snippet *Snippet `json:"snippet,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Links = dataAO1.Links

	m.Snippet = dataAO1.Snippet

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SnippetCommit) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseCommit)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Links *SnippetCommitAO1Links `json:"links,omitempty"`

		Snippet *Snippet `json:"snippet,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}

	dataAO1.Links = m.Links

	dataAO1.Snippet = m.Snippet

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this snippet commit
func (m *SnippetCommit) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseCommit
	if err := m.BaseCommit.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnippet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnippetCommit) validateLinks(formats strfmt.Registry) error {

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

func (m *SnippetCommit) validateSnippet(formats strfmt.Registry) error {

	if swag.IsZero(m.Snippet) { // not required
		return nil
	}

	if m.Snippet != nil {
		if err := m.Snippet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snippet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnippetCommit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnippetCommit) UnmarshalBinary(b []byte) error {
	var res SnippetCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnippetCommitAO1Links snippet commit a o1 links
//
// swagger:model SnippetCommitAO1Links
type SnippetCommitAO1Links struct {

	// diff
	Diff *SnippetCommitAO1LinksDiff `json:"diff,omitempty"`

	// html
	HTML *SnippetCommitAO1LinksHTML `json:"html,omitempty"`

	// self
	Self *SnippetCommitAO1LinksSelf `json:"self,omitempty"`
}

// Validate validates this snippet commit a o1 links
func (m *SnippetCommitAO1Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiff(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTML(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnippetCommitAO1Links) validateDiff(formats strfmt.Registry) error {

	if swag.IsZero(m.Diff) { // not required
		return nil
	}

	if m.Diff != nil {
		if err := m.Diff.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links" + "." + "diff")
			}
			return err
		}
	}

	return nil
}

func (m *SnippetCommitAO1Links) validateHTML(formats strfmt.Registry) error {

	if swag.IsZero(m.HTML) { // not required
		return nil
	}

	if m.HTML != nil {
		if err := m.HTML.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links" + "." + "html")
			}
			return err
		}
	}

	return nil
}

func (m *SnippetCommitAO1Links) validateSelf(formats strfmt.Registry) error {

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
func (m *SnippetCommitAO1Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnippetCommitAO1Links) UnmarshalBinary(b []byte) error {
	var res SnippetCommitAO1Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnippetCommitAO1LinksDiff snippet commit a o1 links diff
//
// swagger:model SnippetCommitAO1LinksDiff
type SnippetCommitAO1LinksDiff struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this snippet commit a o1 links diff
func (m *SnippetCommitAO1LinksDiff) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnippetCommitAO1LinksDiff) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"diff"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnippetCommitAO1LinksDiff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnippetCommitAO1LinksDiff) UnmarshalBinary(b []byte) error {
	var res SnippetCommitAO1LinksDiff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnippetCommitAO1LinksHTML snippet commit a o1 links HTML
//
// swagger:model SnippetCommitAO1LinksHTML
type SnippetCommitAO1LinksHTML struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this snippet commit a o1 links HTML
func (m *SnippetCommitAO1LinksHTML) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnippetCommitAO1LinksHTML) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"html"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnippetCommitAO1LinksHTML) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnippetCommitAO1LinksHTML) UnmarshalBinary(b []byte) error {
	var res SnippetCommitAO1LinksHTML
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnippetCommitAO1LinksSelf snippet commit a o1 links self
//
// swagger:model SnippetCommitAO1LinksSelf
type SnippetCommitAO1LinksSelf struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this snippet commit a o1 links self
func (m *SnippetCommitAO1LinksSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnippetCommitAO1LinksSelf) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"self"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnippetCommitAO1LinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnippetCommitAO1LinksSelf) UnmarshalBinary(b []byte) error {
	var res SnippetCommitAO1LinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
