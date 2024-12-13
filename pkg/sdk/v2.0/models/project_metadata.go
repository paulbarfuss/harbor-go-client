// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectMetadata project metadata
//
// swagger:model ProjectMetadata
type ProjectMetadata struct {

	// Whether generating SBOM automatically when pushing a subject artifact. The valid values are "true", "false".
	AutoSbomGeneration *string `json:"auto_sbom_generation,omitempty"`

	// Whether scan images automatically when pushing. The valid values are "true", "false".
	AutoScan *string `json:"auto_scan,omitempty"`

	// Whether content trust is enabled or not. If it is enabled, user can't pull unsigned images from this project. The valid values are "true", "false".
	EnableContentTrust *string `json:"enable_content_trust,omitempty"`

	// Whether cosign content trust is enabled or not. If it is enabled, user can't pull images without cosign signature from this project. The valid values are "true", "false".
	EnableContentTrustCosign *string `json:"enable_content_trust_cosign,omitempty"`

	// Whether prevent the vulnerable images from running. The valid values are "true", "false".
	PreventVul *string `json:"prevent_vul,omitempty"`

	// The bandwidth limit of proxy cache, in Kbps (kilobits per second). It limits the communication between Harbor and the upstream registry, not the client and the Harbor.
	ProxySpeedKb *string `json:"proxy_speed_kb,omitempty"`

	// The public status of the project. The valid values are "true", "false".
	Public string `json:"public,omitempty"`

	// The ID of the tag retention policy for the project
	RetentionID *string `json:"retention_id,omitempty"`

	// Whether this project reuse the system level CVE allowlist as the allowlist of its own.  The valid values are "true", "false". If it is set to "true" the actual allowlist associate with this project, if any, will be ignored.
	ReuseSysCVEAllowlist *string `json:"reuse_sys_cve_allowlist,omitempty"`

	// If the vulnerability is high than severity defined here, the images can't be pulled. The valid values are "none", "low", "medium", "high", "critical".
	Severity *string `json:"severity,omitempty"`
}

// Validate validates this project metadata
func (m *ProjectMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project metadata based on context it is used
func (m *ProjectMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectMetadata) UnmarshalBinary(b []byte) error {
	var res ProjectMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
