// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SBOMOverview The generate SBOM overview information
//
// swagger:model SBOMOverview
type SBOMOverview struct {

	// Time in seconds required to create the report
	// Example: 300
	Duration int64 `json:"duration,omitempty"`

	// The end time of the generating sbom report task
	// Example: 2006-01-02T15:04:05Z
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// id of the native scan report
	// Example: 5f62c830-f996-11e9-957f-0242c0a89008
	ReportID string `json:"report_id,omitempty"`

	// The digest of the generated SBOM accessory
	SbomDigest string `json:"sbom_digest,omitempty"`

	// The status of the generating SBOM task
	ScanStatus string `json:"scan_status,omitempty"`

	// scanner
	Scanner *Scanner `json:"scanner,omitempty"`

	// The start time of the generating sbom report task
	// Example: 2006-01-02T14:04:05Z
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`
}

// Validate validates this s b o m overview
func (m *SBOMOverview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SBOMOverview) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SBOMOverview) validateScanner(formats strfmt.Registry) error {
	if swag.IsZero(m.Scanner) { // not required
		return nil
	}

	if m.Scanner != nil {
		if err := m.Scanner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanner")
			}
			return err
		}
	}

	return nil
}

func (m *SBOMOverview) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s b o m overview based on the context it is used
func (m *SBOMOverview) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SBOMOverview) contextValidateScanner(ctx context.Context, formats strfmt.Registry) error {

	if m.Scanner != nil {
		if err := m.Scanner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scanner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scanner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SBOMOverview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SBOMOverview) UnmarshalBinary(b []byte) error {
	var res SBOMOverview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
