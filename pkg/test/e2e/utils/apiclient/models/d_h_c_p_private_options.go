// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DHCPPrivateOptions DHCPExtraOptions defines Extra DHCP options for a VM.
//
// swagger:model DHCPPrivateOptions
type DHCPPrivateOptions struct {

	// Option is an Integer value from 224-254
	// Required.
	Option int64 `json:"option,omitempty"`

	// Value is a String value for the Option provided
	// Required.
	Value string `json:"value,omitempty"`
}

// Validate validates this d h c p private options
func (m *DHCPPrivateOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this d h c p private options based on context it is used
func (m *DHCPPrivateOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DHCPPrivateOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DHCPPrivateOptions) UnmarshalBinary(b []byte) error {
	var res DHCPPrivateOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
