// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Collection collection
//
// swagger:model Collection
type Collection struct {

	// name
	Name string `json:"name,omitempty"`

	// index type
	IndexType string `json:"index_type,omitempty"`

	// embedder
	Embedder string `json:"embedder,omitempty"`

	// data type
	DataType string `json:"data_type,omitempty"`

	// index params
	IndexParams interface{} `json:"index_params,omitempty"`
}

// Validate validates this collection
func (m *Collection) Validate(formats strfmt.Registry) error {
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
