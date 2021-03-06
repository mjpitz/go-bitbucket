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

// PipelineSSHPublicKey pipeline ssh public key
//
// swagger:model pipeline_ssh_public_key
type PipelineSSHPublicKey struct {

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`

	// The base64 encoded public key.
	Key string `json:"key,omitempty"`

	// The type of the public key.
	KeyType string `json:"key_type,omitempty"`

	// The MD5 fingerprint of the public key.
	Md5Fingerprint string `json:"md5_fingerprint,omitempty"`

	// The SHA-256 fingerprint of the public key.
	Sha256Fingerprint string `json:"sha256_fingerprint,omitempty"`

	// a o1 additional properties
	AO1AdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this subtype
func (m *PipelineSSHPublicKey) Type() string {
	return "pipeline_ssh_public_key"
}

// SetType sets the type of this subtype
func (m *PipelineSSHPublicKey) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *PipelineSSHPublicKey) UnmarshalJSON(raw []byte) error {
	var data struct {

		// The base64 encoded public key.
		Key string `json:"key,omitempty"`

		// The type of the public key.
		KeyType string `json:"key_type,omitempty"`

		// The MD5 fingerprint of the public key.
		Md5Fingerprint string `json:"md5_fingerprint,omitempty"`

		// The SHA-256 fingerprint of the public key.
		Sha256Fingerprint string `json:"sha256_fingerprint,omitempty"`

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

	var result PipelineSSHPublicKey

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Key = data.Key
	result.KeyType = data.KeyType
	result.Md5Fingerprint = data.Md5Fingerprint
	result.Sha256Fingerprint = data.Sha256Fingerprint

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m PipelineSSHPublicKey) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// The base64 encoded public key.
		Key string `json:"key,omitempty"`

		// The type of the public key.
		KeyType string `json:"key_type,omitempty"`

		// The MD5 fingerprint of the public key.
		Md5Fingerprint string `json:"md5_fingerprint,omitempty"`

		// The SHA-256 fingerprint of the public key.
		Sha256Fingerprint string `json:"sha256_fingerprint,omitempty"`

		AO1AdditionalProperties map[string]interface{} `json:"-"`
	}{

		Key: m.Key,

		KeyType: m.KeyType,

		Md5Fingerprint: m.Md5Fingerprint,

		Sha256Fingerprint: m.Sha256Fingerprint,
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

// Validate validates this pipeline ssh public key
func (m *PipelineSSHPublicKey) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineSSHPublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineSSHPublicKey) UnmarshalBinary(b []byte) error {
	var res PipelineSSHPublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
