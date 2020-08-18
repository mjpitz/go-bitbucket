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

// ReportData A key-value element that will be displayed along with the report.
//
// swagger:model report_data
type ReportData struct {

	// A string describing what this data field represents.
	Title string `json:"title,omitempty"`

	// The type of data contained in the value field. If not provided, then the value will be detected as a boolean, number or string.
	// Enum: [BOOLEAN DATE DURATION LINK NUMBER PERCENTAGE TEXT]
	Type string `json:"type,omitempty"`

	// The value of the data element.
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this report data
func (m *ReportData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var reportDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BOOLEAN","DATE","DURATION","LINK","NUMBER","PERCENTAGE","TEXT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportDataTypeTypePropEnum = append(reportDataTypeTypePropEnum, v)
	}
}

const (

	// ReportDataTypeBOOLEAN captures enum value "BOOLEAN"
	ReportDataTypeBOOLEAN string = "BOOLEAN"

	// ReportDataTypeDATE captures enum value "DATE"
	ReportDataTypeDATE string = "DATE"

	// ReportDataTypeDURATION captures enum value "DURATION"
	ReportDataTypeDURATION string = "DURATION"

	// ReportDataTypeLINK captures enum value "LINK"
	ReportDataTypeLINK string = "LINK"

	// ReportDataTypeNUMBER captures enum value "NUMBER"
	ReportDataTypeNUMBER string = "NUMBER"

	// ReportDataTypePERCENTAGE captures enum value "PERCENTAGE"
	ReportDataTypePERCENTAGE string = "PERCENTAGE"

	// ReportDataTypeTEXT captures enum value "TEXT"
	ReportDataTypeTEXT string = "TEXT"
)

// prop value enum
func (m *ReportData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportData) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportData) UnmarshalBinary(b []byte) error {
	var res ReportData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
