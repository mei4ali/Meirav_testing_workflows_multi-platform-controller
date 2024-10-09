// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SharedProcessorPoolUpdate shared processor pool update
//
// swagger:model SharedProcessorPoolUpdate
type SharedProcessorPoolUpdate struct {

	// The new name for the Shared Processor Pool; minumum of 2 characters, maximum of 12, the only special character allowed is the underscore '_'.
	Name string `json:"name,omitempty"`

	// The amount of reserved processor cores for the Shared Processor Pool; only integers allowed, no fractional values; the amount can be increased (dependent on available resources) or decreased (dependent on currently allocated resources)
	ReservedCores *int64 `json:"reservedCores,omitempty"`
}

// Validate validates this shared processor pool update
func (m *SharedProcessorPoolUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this shared processor pool update based on context it is used
func (m *SharedProcessorPoolUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SharedProcessorPoolUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SharedProcessorPoolUpdate) UnmarshalBinary(b []byte) error {
	var res SharedProcessorPoolUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}