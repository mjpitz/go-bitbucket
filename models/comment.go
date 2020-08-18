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

// Comment comment
//
// swagger:model comment
type Comment struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	// content
	Content *CommentAO1Content `json:"content,omitempty"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// deleted
	Deleted bool `json:"deleted,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// inline
	Inline *CommentAO1Inline `json:"inline,omitempty"`

	// links
	Links *CommentAO1Links `json:"links,omitempty"`

	// parent
	Parent *Comment `json:"parent,omitempty"`

	// updated on
	// Format: date-time
	UpdatedOn strfmt.DateTime `json:"updated_on,omitempty"`

	// user
	User *User `json:"user,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this subtype
func (m *Comment) Type() string {
	return "comment"
}

// SetType sets the type of this subtype
func (m *Comment) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Comment) UnmarshalJSON(raw []byte) error {
	var data struct {

		// content
		Content *CommentAO1Content `json:"content,omitempty"`

		// created on
		// Format: date-time
		CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

		// deleted
		Deleted bool `json:"deleted,omitempty"`

		// id
		ID int64 `json:"id,omitempty"`

		// inline
		Inline *CommentAO1Inline `json:"inline,omitempty"`

		// links
		Links *CommentAO1Links `json:"links,omitempty"`

		// parent
		Parent *Comment `json:"parent,omitempty"`

		// updated on
		// Format: date-time
		UpdatedOn strfmt.DateTime `json:"updated_on,omitempty"`

		// user
		User *User `json:"user,omitempty"`

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

	var result Comment

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Content = data.Content
	result.CreatedOn = data.CreatedOn
	result.Deleted = data.Deleted
	result.ID = data.ID
	result.Inline = data.Inline
	result.Links = data.Links
	result.Parent = data.Parent
	result.UpdatedOn = data.UpdatedOn
	result.User = data.User

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Comment) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// content
		Content *CommentAO1Content `json:"content,omitempty"`

		// created on
		// Format: date-time
		CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

		// deleted
		Deleted bool `json:"deleted,omitempty"`

		// id
		ID int64 `json:"id,omitempty"`

		// inline
		Inline *CommentAO1Inline `json:"inline,omitempty"`

		// links
		Links *CommentAO1Links `json:"links,omitempty"`

		// parent
		Parent *Comment `json:"parent,omitempty"`

		// updated on
		// Format: date-time
		UpdatedOn strfmt.DateTime `json:"updated_on,omitempty"`

		// user
		User *User `json:"user,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}{

		Content: m.Content,

		CreatedOn: m.CreatedOn,

		Deleted: m.Deleted,

		ID: m.ID,

		Inline: m.Inline,

		Links: m.Links,

		Parent: m.Parent,

		UpdatedOn: m.UpdatedOn,

		User: m.User,
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

// Validate validates this comment
func (m *Comment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedOn(formats); err != nil {
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

func (m *Comment) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

func (m *Comment) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Comment) validateInline(formats strfmt.Registry) error {

	if swag.IsZero(m.Inline) { // not required
		return nil
	}

	if m.Inline != nil {
		if err := m.Inline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inline")
			}
			return err
		}
	}

	return nil
}

func (m *Comment) validateLinks(formats strfmt.Registry) error {

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

func (m *Comment) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *Comment) validateUpdatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_on", "body", "date-time", m.UpdatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Comment) validateUser(formats strfmt.Registry) error {

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
func (m *Comment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Comment) UnmarshalBinary(b []byte) error {
	var res Comment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CommentAO1Content comment a o1 content
//
// swagger:model CommentAO1Content
type CommentAO1Content struct {

	// The user's content rendered as HTML.
	HTML string `json:"html,omitempty"`

	// The type of markup language the raw content is to be interpreted in.
	// Enum: [markdown creole plaintext]
	Markup string `json:"markup,omitempty"`

	// The text as it was typed by a user.
	Raw string `json:"raw,omitempty"`
}

// Validate validates this comment a o1 content
func (m *CommentAO1Content) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMarkup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var commentAO1ContentTypeMarkupPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["markdown","creole","plaintext"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commentAO1ContentTypeMarkupPropEnum = append(commentAO1ContentTypeMarkupPropEnum, v)
	}
}

const (

	// CommentAO1ContentMarkupMarkdown captures enum value "markdown"
	CommentAO1ContentMarkupMarkdown string = "markdown"

	// CommentAO1ContentMarkupCreole captures enum value "creole"
	CommentAO1ContentMarkupCreole string = "creole"

	// CommentAO1ContentMarkupPlaintext captures enum value "plaintext"
	CommentAO1ContentMarkupPlaintext string = "plaintext"
)

// prop value enum
func (m *CommentAO1Content) validateMarkupEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commentAO1ContentTypeMarkupPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommentAO1Content) validateMarkup(formats strfmt.Registry) error {

	if swag.IsZero(m.Markup) { // not required
		return nil
	}

	// value enum
	if err := m.validateMarkupEnum("content"+"."+"markup", "body", m.Markup); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommentAO1Content) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentAO1Content) UnmarshalBinary(b []byte) error {
	var res CommentAO1Content
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CommentAO1Inline comment a o1 inline
//
// swagger:model CommentAO1Inline
type CommentAO1Inline struct {

	// The comment's anchor line in the old version of the file.
	// Minimum: 1
	From int64 `json:"from,omitempty"`

	// The path of the file this comment is anchored to.
	// Required: true
	Path *string `json:"path"`

	// The comment's anchor line in the new version of the file. If the 'from' line is also provided, this value will be removed.
	// Minimum: 1
	To int64 `json:"to,omitempty"`
}

// Validate validates this comment a o1 inline
func (m *CommentAO1Inline) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommentAO1Inline) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if err := validate.MinimumInt("inline"+"."+"from", "body", int64(m.From), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *CommentAO1Inline) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("inline"+"."+"path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *CommentAO1Inline) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	if err := validate.MinimumInt("inline"+"."+"to", "body", int64(m.To), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommentAO1Inline) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentAO1Inline) UnmarshalBinary(b []byte) error {
	var res CommentAO1Inline
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CommentAO1Links comment a o1 links
//
// swagger:model CommentAO1Links
type CommentAO1Links struct {

	// code
	Code *CommentAO1LinksCode `json:"code,omitempty"`

	// html
	HTML *CommentAO1LinksHTML `json:"html,omitempty"`

	// self
	Self *CommentAO1LinksSelf `json:"self,omitempty"`
}

// Validate validates this comment a o1 links
func (m *CommentAO1Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
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

func (m *CommentAO1Links) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	if m.Code != nil {
		if err := m.Code.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links" + "." + "code")
			}
			return err
		}
	}

	return nil
}

func (m *CommentAO1Links) validateHTML(formats strfmt.Registry) error {

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

func (m *CommentAO1Links) validateSelf(formats strfmt.Registry) error {

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
func (m *CommentAO1Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentAO1Links) UnmarshalBinary(b []byte) error {
	var res CommentAO1Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CommentAO1LinksCode comment a o1 links code
//
// swagger:model CommentAO1LinksCode
type CommentAO1LinksCode struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this comment a o1 links code
func (m *CommentAO1LinksCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommentAO1LinksCode) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"code"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommentAO1LinksCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentAO1LinksCode) UnmarshalBinary(b []byte) error {
	var res CommentAO1LinksCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CommentAO1LinksHTML comment a o1 links HTML
//
// swagger:model CommentAO1LinksHTML
type CommentAO1LinksHTML struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this comment a o1 links HTML
func (m *CommentAO1LinksHTML) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommentAO1LinksHTML) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"html"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommentAO1LinksHTML) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentAO1LinksHTML) UnmarshalBinary(b []byte) error {
	var res CommentAO1LinksHTML
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CommentAO1LinksSelf comment a o1 links self
//
// swagger:model CommentAO1LinksSelf
type CommentAO1LinksSelf struct {

	// href
	// Format: uri
	Href strfmt.URI `json:"href,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this comment a o1 links self
func (m *CommentAO1LinksSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommentAO1LinksSelf) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.FormatOf("links"+"."+"self"+"."+"href", "body", "uri", m.Href.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommentAO1LinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentAO1LinksSelf) UnmarshalBinary(b []byte) error {
	var res CommentAO1LinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
