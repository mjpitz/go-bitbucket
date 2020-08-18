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

// SubjectTypes The mapping of resource/subject types pointing to their individual event types.
//
// swagger:model subject_types
type SubjectTypes struct {

	// repository
	Repository *SubjectTypesRepository `json:"repository,omitempty"`

	// team
	Team *SubjectTypesTeam `json:"team,omitempty"`

	// user
	User *SubjectTypesUser `json:"user,omitempty"`
}

// Validate validates this subject types
func (m *SubjectTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubjectTypes) validateRepository(formats strfmt.Registry) error {

	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *SubjectTypes) validateTeam(formats strfmt.Registry) error {

	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *SubjectTypes) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubjectTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectTypes) UnmarshalBinary(b []byte) error {
	var res SubjectTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SubjectTypesRepository subject types repository
//
// swagger:model SubjectTypesRepository
type SubjectTypesRepository struct {

	// events
	Events *SubjectTypesRepositoryEvents `json:"events,omitempty"`
}

// Validate validates this subject types repository
func (m *SubjectTypesRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubjectTypesRepository) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	if m.Events != nil {
		if err := m.Events.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository" + "." + "events")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubjectTypesRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectTypesRepository) UnmarshalBinary(b []byte) error {
	var res SubjectTypesRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SubjectTypesRepositoryEvents subject types repository events
//
// swagger:model SubjectTypesRepositoryEvents
type SubjectTypesRepositoryEvents struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this subject types repository events
func (m *SubjectTypesRepositoryEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubjectTypesRepositoryEvents) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("repository"+"."+"events"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubjectTypesRepositoryEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectTypesRepositoryEvents) UnmarshalBinary(b []byte) error {
	var res SubjectTypesRepositoryEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SubjectTypesTeam subject types team
//
// swagger:model SubjectTypesTeam
type SubjectTypesTeam struct {

	// events
	Events *SubjectTypesTeamEvents `json:"events,omitempty"`
}

// Validate validates this subject types team
func (m *SubjectTypesTeam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubjectTypesTeam) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	if m.Events != nil {
		if err := m.Events.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team" + "." + "events")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubjectTypesTeam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectTypesTeam) UnmarshalBinary(b []byte) error {
	var res SubjectTypesTeam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SubjectTypesTeamEvents subject types team events
//
// swagger:model SubjectTypesTeamEvents
type SubjectTypesTeamEvents struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this subject types team events
func (m *SubjectTypesTeamEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubjectTypesTeamEvents) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("team"+"."+"events"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubjectTypesTeamEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectTypesTeamEvents) UnmarshalBinary(b []byte) error {
	var res SubjectTypesTeamEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SubjectTypesUser subject types user
//
// swagger:model SubjectTypesUser
type SubjectTypesUser struct {

	// events
	Events *SubjectTypesUserEvents `json:"events,omitempty"`
}

// Validate validates this subject types user
func (m *SubjectTypesUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubjectTypesUser) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	if m.Events != nil {
		if err := m.Events.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user" + "." + "events")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubjectTypesUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectTypesUser) UnmarshalBinary(b []byte) error {
	var res SubjectTypesUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SubjectTypesUserEvents subject types user events
//
// swagger:model SubjectTypesUserEvents
type SubjectTypesUserEvents struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this subject types user events
func (m *SubjectTypesUserEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubjectTypesUserEvents) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("user"+"."+"events"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubjectTypesUserEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectTypesUserEvents) UnmarshalBinary(b []byte) error {
	var res SubjectTypesUserEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}