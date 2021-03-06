// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BillTo bill to
// swagger:model BillTo
type BillTo struct {

	// First line of the billing street address.
	//
	// This field is optional if your CyberSource account is configured for relaxed requirements for address data and expiration date. See the TMS REST Developer Guide for more information about relaxed address requirements.
	//
	// **Important**:
	// It is your responsibility to determine whether a field is required for the transaction you are requesting.
	//
	// Max Length: 60
	Address1 string `json:"address1,omitempty"`

	// Additional address information.
	// Max Length: 60
	Address2 string `json:"address2,omitempty"`

	// State or province of the billing address. For an address in the U.S. or Canada, use the State, Province, and Territory Codes for the United States and Canada.
	//
	// This field is optional if your CyberSource account is configured for relaxed requirements for address data and expiration date. See the TMS REST Developer Guide for more information about relaxed address requirements.
	//
	// **Important**:
	// It is your responsibility to determine whether a field is required for the transaction you are requesting.
	//
	// Max Length: 20
	AdministrativeArea string `json:"administrativeArea,omitempty"`

	// Name of the customer’s company.
	//
	// This field is optional if your CyberSource account is configured for relaxed requirements for address data and expiration date. See the TMS REST Developer Guide for more information about relaxed address requirements.
	//
	// **Important**:
	// It is your responsibility to determine whether a field is required for the transaction you are requesting.
	//
	// Max Length: 60
	Company string `json:"company,omitempty"`

	// Country of the billing address. Accepts input in the ISO 3166-1 standard, stores as ISO 3166-1-Alpha-2.
	//
	// This field is optional if your CyberSource account is configured for relaxed requirements for address data and expiration date. See the TMS REST Developer Guide for more information about relaxed address requirements.
	//
	// **Important** It is your responsibility to determine whether a field is required for the transaction you are requesting.
	//
	// Max Length: 3
	// Min Length: 2
	Country string `json:"country,omitempty"`

	// Customer’s email address.
	//
	// This field is optional if your CyberSource account is configured for relaxed requirements for address data and expiration date. See the TMS REST Developer Guide for more information about relaxed address requirements.
	//
	// **Important** It is your responsibility to determine whether a field is required for the transaction you are requesting.
	//
	// Max Length: 320
	Email string `json:"email,omitempty"`

	// Customer’s first name. For a credit card transaction, this name must match the name on the card.
	//
	// This field is optional if your CyberSource account is configured for relaxed requirements for address data and expiration date. See the TMS REST Developer Guide for more information about relaxed address requirements.
	//
	// **Important**:
	// It is your responsibility to determine whether a field is required for the transaction you are requesting.
	//
	// Max Length: 60
	FirstName string `json:"firstName,omitempty"`

	// Customer’s last name. For a credit card transaction, this name must match the name on the card.
	//
	// This field is optional if your CyberSource account is configured for relaxed requirements for address data and expiration date. See the TMS REST Developer Guide for more information about relaxed address requirements.
	//
	// **Important**:
	// It is your responsibility to determine whether a field is required for the transaction you are requesting.
	//
	// Max Length: 60
	LastName string `json:"lastName,omitempty"`

	// City of the billing address.
	//
	// This field is optional if your CyberSource account is configured for relaxed requirements for address data and expiration date. See the TMS REST Developer Guide for more information about relaxed address requirements.
	//
	// **Important**:
	// It is your responsibility to determine whether a field is required for the transaction you are requesting.
	//
	// Max Length: 50
	Locality string `json:"locality,omitempty"`

	// Customer phone number. When you create a customer profile, the requirements depend on the payment method:
	//   * Credit cards — optional.
	//   * Electronic checks — contact your payment processor representative to find out if this field is required or optional.
	//   * PINless debits — optional.
	//
	// Max Length: 32
	// Min Length: 6
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Postal code for the billing address. The postal code must consist of 5 to 9 digits.
	//
	// When the billing country is the U.S., the 9-digit postal code must follow this format:
	// [5 digits][dash][4 digits]
	// **Example** 12345-6789
	//
	// When the billing country is Canada, the 6-digit postal code must follow this format:
	// [alpha][numeric][alpha][space]
	// [numeric][alpha][numeric]
	// Example A1B 2C3
	//
	// This field is optional if your CyberSource account is configured for relaxed requirements for address data and expiration date. See the TMS REST Developer Guide for more information about relaxed address requirements.
	//
	// **Important**:
	// It is your responsibility to determine whether a field is required for the transaction you are requesting.
	//
	// Max Length: 10
	PostalCode string `json:"postalCode,omitempty"`
}

// Validate validates this bill to
func (m *BillTo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddress2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdministrativeArea(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompany(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillTo) validateAddress1(formats strfmt.Registry) error {

	if swag.IsZero(m.Address1) { // not required
		return nil
	}

	if err := validate.MaxLength("address1", "body", string(m.Address1), 60); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validateAddress2(formats strfmt.Registry) error {

	if swag.IsZero(m.Address2) { // not required
		return nil
	}

	if err := validate.MaxLength("address2", "body", string(m.Address2), 60); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validateAdministrativeArea(formats strfmt.Registry) error {

	if swag.IsZero(m.AdministrativeArea) { // not required
		return nil
	}

	if err := validate.MaxLength("administrativeArea", "body", string(m.AdministrativeArea), 20); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validateCompany(formats strfmt.Registry) error {

	if swag.IsZero(m.Company) { // not required
		return nil
	}

	if err := validate.MaxLength("company", "body", string(m.Company), 60); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := validate.MinLength("country", "body", string(m.Country), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("country", "body", string(m.Country), 3); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.MaxLength("email", "body", string(m.Email), 320); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validateFirstName(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstName) { // not required
		return nil
	}

	if err := validate.MaxLength("firstName", "body", string(m.FirstName), 60); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validateLastName(formats strfmt.Registry) error {

	if swag.IsZero(m.LastName) { // not required
		return nil
	}

	if err := validate.MaxLength("lastName", "body", string(m.LastName), 60); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validateLocality(formats strfmt.Registry) error {

	if swag.IsZero(m.Locality) { // not required
		return nil
	}

	if err := validate.MaxLength("locality", "body", string(m.Locality), 50); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validatePhoneNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneNumber) { // not required
		return nil
	}

	if err := validate.MinLength("phoneNumber", "body", string(m.PhoneNumber), 6); err != nil {
		return err
	}

	if err := validate.MaxLength("phoneNumber", "body", string(m.PhoneNumber), 32); err != nil {
		return err
	}

	return nil
}

func (m *BillTo) validatePostalCode(formats strfmt.Registry) error {

	if swag.IsZero(m.PostalCode) { // not required
		return nil
	}

	if err := validate.MaxLength("postalCode", "body", string(m.PostalCode), 10); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillTo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillTo) UnmarshalBinary(b []byte) error {
	var res BillTo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
