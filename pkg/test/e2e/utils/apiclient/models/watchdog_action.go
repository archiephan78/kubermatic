// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// WatchdogAction WatchdogAction defines the watchdog action, if a watchdog gets triggered.
//
// swagger:model WatchdogAction
type WatchdogAction string

// Validate validates this watchdog action
func (m WatchdogAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this watchdog action based on context it is used
func (m WatchdogAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
