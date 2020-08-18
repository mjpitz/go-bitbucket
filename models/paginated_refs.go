// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaginatedRefs A paginated list of refs.
//
// swagger:model paginated_refs
type PaginatedRefs struct {

	// Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	// Format: uri
	Next strfmt.URI `json:"next,omitempty"`

	// Page number of the current results. This is an optional element that is not provided in all responses.
	// Minimum: 1
	Page int64 `json:"page,omitempty"`

	// Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	// Minimum: 1
	Pagelen int64 `json:"pagelen,omitempty"`

	// Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	// Format: uri
	Previous strfmt.URI `json:"previous,omitempty"`

	// Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	// Minimum: 0
	Size *int64 `json:"size,omitempty"`

	// values
	// Min Items: 0
	// Unique: true
	Values []*Ref `json:"values"`
}

// Validate validates this paginated refs
func (m *PaginatedRefs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagelen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedRefs) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("next", "body", "uri", m.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedRefs) validatePage(formats strfmt.Registry) error {

	if swag.IsZero(m.Page) { // not required
		return nil
	}

	if err := validate.MinimumInt("page", "body", int64(m.Page), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedRefs) validatePagelen(formats strfmt.Registry) error {

	if swag.IsZero(m.Pagelen) { // not required
		return nil
	}

	if err := validate.MinimumInt("pagelen", "body", int64(m.Pagelen), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedRefs) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(m.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("previous", "body", "uri", m.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedRefs) validateSize(formats strfmt.Registry) error {

	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if err := validate.MinimumInt("size", "body", int64(*m.Size), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedRefs) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	iValuesSize := int64(len(m.Values))

	if err := validate.MinItems("values", "body", iValuesSize, 0); err != nil {
		return err
	}

	if err := validate.UniqueItems("values", "body", m.Values); err != nil {
		return err
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginatedRefs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginatedRefs) UnmarshalBinary(b []byte) error {
	var res PaginatedRefs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
