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

// Collection collection
//
// swagger:model Collection
type Collection struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// index type
	IndexType string `json:"index_type,omitempty"`

	// distance metric
	DistanceMetric string `json:"distance_metric,omitempty"`

	// embedder
	Embedder string `json:"embedder,omitempty"`

	// data type
	DataType string `json:"data_type,omitempty"`

	// index params
	IndexParams interface{} `json:"index_params,omitempty"`
}

// Validate validates this collection
func (m *Collection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Collection) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Collection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Collection) UnmarshalBinary(b []byte) error {
	var res Collection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
