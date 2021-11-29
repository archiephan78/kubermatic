// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventRateLimitConfigItem event rate limit config item
//
// swagger:model EventRateLimitConfigItem
type EventRateLimitConfigItem struct {

	// burst
	Burst int32 `json:"burst,omitempty"`

	// cache size
	CacheSize int32 `json:"cacheSize,omitempty"`

	// QPS
	QPS int32 `json:"qps,omitempty"`
}

// Validate validates this event rate limit config item
func (m *EventRateLimitConfigItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this event rate limit config item based on context it is used
func (m *EventRateLimitConfigItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventRateLimitConfigItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventRateLimitConfigItem) UnmarshalBinary(b []byte) error {
	var res EventRateLimitConfigItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
