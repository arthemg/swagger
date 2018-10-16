// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Jsonrepo jsonrepo
// swagger:model jsonrepo
type Jsonrepo []*JsonrepoItems0

// Validate validates this jsonrepo
func (m Jsonrepo) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// JsonrepoItems0 jsonrepo items0
// swagger:model JsonrepoItems0
type JsonrepoItems0 struct {

	// full name
	FullName string `json:"full_name,omitempty"`

	// html url
	HTMLURL string `json:"html_url,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// login
	Login string `json:"login,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this jsonrepo items0
func (m *JsonrepoItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JsonrepoItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JsonrepoItems0) UnmarshalBinary(b []byte) error {
	var res JsonrepoItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}