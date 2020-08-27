// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TMSErrorResponse t m s error response
//
// swagger:model TMSErrorResponse
type TMSErrorResponse struct {

	// errors
	// Read Only: true
	Errors []*TMSErrorResponseErrorsItems0 `json:"errors"`
}

// Validate validates this t m s error response
func (m *TMSErrorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TMSErrorResponse) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TMSErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TMSErrorResponse) UnmarshalBinary(b []byte) error {
	var res TMSErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TMSErrorResponseErrorsItems0 t m s error response errors items0
//
// swagger:model TMSErrorResponseErrorsItems0
type TMSErrorResponseErrorsItems0 struct {

	// details
	// Read Only: true
	Details []*TMSErrorResponseErrorsItems0DetailsItems0 `json:"details"`

	// The detailed message related to the type stated above.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The type of error.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this t m s error response errors items0
func (m *TMSErrorResponseErrorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TMSErrorResponseErrorsItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	for i := 0; i < len(m.Details); i++ {
		if swag.IsZero(m.Details[i]) { // not required
			continue
		}

		if m.Details[i] != nil {
			if err := m.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TMSErrorResponseErrorsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TMSErrorResponseErrorsItems0) UnmarshalBinary(b []byte) error {
	var res TMSErrorResponseErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TMSErrorResponseErrorsItems0DetailsItems0 t m s error response errors items0 details items0
//
// swagger:model TMSErrorResponseErrorsItems0DetailsItems0
type TMSErrorResponseErrorsItems0DetailsItems0 struct {

	// The location of the field that caused the error.
	// Read Only: true
	Location string `json:"location,omitempty"`

	// The name of the field that caused the error.
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this t m s error response errors items0 details items0
func (m *TMSErrorResponseErrorsItems0DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TMSErrorResponseErrorsItems0DetailsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TMSErrorResponseErrorsItems0DetailsItems0) UnmarshalBinary(b []byte) error {
	var res TMSErrorResponseErrorsItems0DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
