// Code generated by go-swagger; DO NOT EDIT.

package taxes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VoidTaxReader is a Reader for the VoidTax structure.
type VoidTaxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoidTaxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVoidTaxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVoidTaxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewVoidTaxBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVoidTaxOK creates a VoidTaxOK with default headers values
func NewVoidTaxOK() *VoidTaxOK {
	return &VoidTaxOK{}
}

/*VoidTaxOK handles this case with default header values.

Successful response.
*/
type VoidTaxOK struct {
	Payload *VoidTaxOKBody
}

func (o *VoidTaxOK) Error() string {
	return fmt.Sprintf("[PATCH /vas/v2/tax/{id}][%d] voidTaxOK  %+v", 200, o.Payload)
}

func (o *VoidTaxOK) GetPayload() *VoidTaxOKBody {
	return o.Payload
}

func (o *VoidTaxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VoidTaxOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoidTaxBadRequest creates a VoidTaxBadRequest with default headers values
func NewVoidTaxBadRequest() *VoidTaxBadRequest {
	return &VoidTaxBadRequest{}
}

/*VoidTaxBadRequest handles this case with default header values.

Invalid request.
*/
type VoidTaxBadRequest struct {
	Payload *VoidTaxBadRequestBody
}

func (o *VoidTaxBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /vas/v2/tax/{id}][%d] voidTaxBadRequest  %+v", 400, o.Payload)
}

func (o *VoidTaxBadRequest) GetPayload() *VoidTaxBadRequestBody {
	return o.Payload
}

func (o *VoidTaxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VoidTaxBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoidTaxBadGateway creates a VoidTaxBadGateway with default headers values
func NewVoidTaxBadGateway() *VoidTaxBadGateway {
	return &VoidTaxBadGateway{}
}

/*VoidTaxBadGateway handles this case with default header values.

Unexpected system error or system timeout.
*/
type VoidTaxBadGateway struct {
	Payload *VoidTaxBadGatewayBody
}

func (o *VoidTaxBadGateway) Error() string {
	return fmt.Sprintf("[PATCH /vas/v2/tax/{id}][%d] voidTaxBadGateway  %+v", 502, o.Payload)
}

func (o *VoidTaxBadGateway) GetPayload() *VoidTaxBadGatewayBody {
	return o.Payload
}

func (o *VoidTaxBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VoidTaxBadGatewayBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VoidTaxBadGatewayBody vasV2TaxVoidsPost502Response
swagger:model VoidTaxBadGatewayBody
*/
type VoidTaxBadGatewayBody struct {

	// The detail message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of the status.
	//
	// Possible values:
	//  - SYSTEM_ERROR
	//  - SERVER_TIMEOUT
	//  - SERVICE_TIMEOUT
	//
	Reason string `json:"reason,omitempty"`

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - SERVER_ERROR
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// **Example** `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.).
	// The `T` separates the date and the time. The `Z` indicates UTC.
	//
	// Returned by authorization service.
	//
	// #### PIN debit
	// Time when the PIN debit credit, PIN debit purchase or PIN debit reversal was requested.
	//
	// Returned by PIN debit credit, PIN debit purchase or PIN debit reversal.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this void tax bad gateway body
func (o *VoidTaxBadGatewayBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VoidTaxBadGatewayBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidTaxBadGatewayBody) UnmarshalBinary(b []byte) error {
	var res VoidTaxBadGatewayBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidTaxBadRequestBody vasV2TaxVoidsPost400Response
swagger:model VoidTaxBadRequestBody
*/
type VoidTaxBadRequestBody struct {

	// details
	Details []*VoidTaxBadRequestBodyDetailsItems0 `json:"details"`

	// The detail message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of the status.
	//
	// Possible values:
	//  - INVALID_DATA
	//  - NOT_VOIDABLE
	//
	Reason string `json:"reason,omitempty"`

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - INVALID_REQUEST
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// **Example** `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.).
	// The `T` separates the date and the time. The `Z` indicates UTC.
	//
	// Returned by authorization service.
	//
	// #### PIN debit
	// Time when the PIN debit credit, PIN debit purchase or PIN debit reversal was requested.
	//
	// Returned by PIN debit credit, PIN debit purchase or PIN debit reversal.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this void tax bad request body
func (o *VoidTaxBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidTaxBadRequestBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("voidTaxBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidTaxBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidTaxBadRequestBody) UnmarshalBinary(b []byte) error {
	var res VoidTaxBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidTaxBadRequestBodyDetailsItems0 void tax bad request body details items0
swagger:model VoidTaxBadRequestBodyDetailsItems0
*/
type VoidTaxBadRequestBodyDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	Field string `json:"field,omitempty"`

	// Possible reasons for the error.
	//
	// Possible values:
	//  - MISSING_FIELD
	//  - INVALID_DATA
	//
	Reason string `json:"reason,omitempty"`
}

// Validate validates this void tax bad request body details items0
func (o *VoidTaxBadRequestBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VoidTaxBadRequestBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidTaxBadRequestBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res VoidTaxBadRequestBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidTaxBody void tax body
swagger:model VoidTaxBody
*/
type VoidTaxBody struct {

	// client reference information
	ClientReferenceInformation *VoidTaxParamsBodyClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
}

// Validate validates this void tax body
func (o *VoidTaxBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientReferenceInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidTaxBody) validateClientReferenceInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientReferenceInformation) { // not required
		return nil
	}

	if o.ClientReferenceInformation != nil {
		if err := o.ClientReferenceInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidTaxRequest" + "." + "clientReferenceInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidTaxBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidTaxBody) UnmarshalBinary(b []byte) error {
	var res VoidTaxBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidTaxOKBody vasV2TaxVoid200Response
swagger:model VoidTaxOKBody
*/
type VoidTaxOKBody struct {

	// client reference information
	ClientReferenceInformation *VoidTaxOKBodyClientReferenceInformation `json:"clientReferenceInformation,omitempty"`

	// An unique identification number to identify the submitted request. It is also appended to the endpoint of the resource.
	//
	// On incremental authorizations, this value with be the same as the identification number returned in the original authorization response.
	//
	// #### PIN debit
	// Returned for all PIN debit services.
	//
	// Max Length: 26
	ID string `json:"id,omitempty"`

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - VOIDED
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// **Example** `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.).
	// The `T` separates the date and the time. The `Z` indicates UTC.
	//
	// Returned by authorization service.
	//
	// #### PIN debit
	// Time when the PIN debit credit, PIN debit purchase or PIN debit reversal was requested.
	//
	// Returned by PIN debit credit, PIN debit purchase or PIN debit reversal.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`

	// void amount details
	VoidAmountDetails *VoidTaxOKBodyVoidAmountDetails `json:"voidAmountDetails,omitempty"`
}

// Validate validates this void tax o k body
func (o *VoidTaxOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientReferenceInformation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVoidAmountDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidTaxOKBody) validateClientReferenceInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientReferenceInformation) { // not required
		return nil
	}

	if o.ClientReferenceInformation != nil {
		if err := o.ClientReferenceInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidTaxOK" + "." + "clientReferenceInformation")
			}
			return err
		}
	}

	return nil
}

func (o *VoidTaxOKBody) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("voidTaxOK"+"."+"id", "body", string(o.ID), 26); err != nil {
		return err
	}

	return nil
}

func (o *VoidTaxOKBody) validateVoidAmountDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.VoidAmountDetails) { // not required
		return nil
	}

	if o.VoidAmountDetails != nil {
		if err := o.VoidAmountDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidTaxOK" + "." + "voidAmountDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidTaxOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidTaxOKBody) UnmarshalBinary(b []byte) error {
	var res VoidTaxOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidTaxOKBodyClientReferenceInformation void tax o k body client reference information
swagger:model VoidTaxOKBodyClientReferenceInformation
*/
type VoidTaxOKBodyClientReferenceInformation struct {

	// Merchant-generated order reference or tracking number. It is recommended that you send a unique value for each
	// transaction so that you can perform meaningful searches for the transaction.
	//
	// #### Used by
	// **Authorization**
	// Required field.
	//
	// #### PIN Debit
	// Requests for PIN debit reversals need to use the same merchant reference number that was used in the transaction that is being
	// reversed.
	//
	// Required field for all PIN Debit requests (purchase, credit, and reversal).
	//
	// #### FDC Nashville Global
	// Certain circumstances can cause the processor to truncate this value to 15 or 17 characters for Level II and Level III processing, which can cause a discrepancy between the value you submit and the value included in some processor reports.
	//
	// Max Length: 50
	Code string `json:"code,omitempty"`
}

// Validate validates this void tax o k body client reference information
func (o *VoidTaxOKBodyClientReferenceInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidTaxOKBodyClientReferenceInformation) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(o.Code) { // not required
		return nil
	}

	if err := validate.MaxLength("voidTaxOK"+"."+"clientReferenceInformation"+"."+"code", "body", string(o.Code), 50); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidTaxOKBodyClientReferenceInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidTaxOKBodyClientReferenceInformation) UnmarshalBinary(b []byte) error {
	var res VoidTaxOKBodyClientReferenceInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidTaxOKBodyVoidAmountDetails void tax o k body void amount details
swagger:model VoidTaxOKBodyVoidAmountDetails
*/
type VoidTaxOKBodyVoidAmountDetails struct {

	// Currency used for the order. Use the three-character [ISO Standard Currency Codes.](http://apps.cybersource.com/library/documentation/sbc/quickref/currencies.pdf)
	//
	// #### Used by
	// **Authorization**
	// Required field.
	//
	// **Authorization Reversal**
	// For an authorization reversal (`reversalInformation`) or a capture (`processingOptions.capture` is set to `true`), you must use the same currency that you used in your payment authorization request.
	//
	// #### PIN Debit
	// Currency for the amount you requested for the PIN debit purchase. This value is returned for partial authorizations. The issuing bank can approve a partial amount if the balance on the debit card is less than the requested transaction amount. For the possible values, see the [ISO Standard Currency Codes](https://developer.cybersource.com/library/documentation/sbc/quickref/currencies.pdf).
	// Returned by PIN debit purchase.
	//
	// For PIN debit reversal requests, you must use the same currency that was used for the PIN debit purchase or PIN debit credit that you are reversing.
	// For the possible values, see the [ISO Standard Currency Codes](https://developer.cybersource.com/library/documentation/sbc/quickref/currencies.pdf).
	//
	// Required field for PIN Debit purchase and PIN Debit credit requests.
	// Optional field for PIN Debit reversal requests.
	//
	// #### GPX
	// This field is optional for reversing an authorization or credit.
	//
	// #### DCC for First Data
	// Your local currency. For details, see the `currency` field description in [Dynamic Currency Conversion For First Data Using the SCMP API](http://apps.cybersource.com/library/documentation/dev_guides/DCC_FirstData_SCMP/DCC_FirstData_SCMP_API.pdf).
	//
	// #### Tax Calculation
	// Required for international tax and value added tax only.
	// Optional for U.S. and Canadian taxes.
	// Your local currency.
	//
	// Max Length: 3
	Currency string `json:"currency,omitempty"`

	// Total amount of the void.
	//
	// #### PIN Debit
	// Amount of the reversal.
	//
	// Returned by PIN debit reversal.
	//
	VoidAmount string `json:"voidAmount,omitempty"`
}

// Validate validates this void tax o k body void amount details
func (o *VoidTaxOKBodyVoidAmountDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidTaxOKBodyVoidAmountDetails) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(o.Currency) { // not required
		return nil
	}

	if err := validate.MaxLength("voidTaxOK"+"."+"voidAmountDetails"+"."+"currency", "body", string(o.Currency), 3); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidTaxOKBodyVoidAmountDetails) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidTaxOKBodyVoidAmountDetails) UnmarshalBinary(b []byte) error {
	var res VoidTaxOKBodyVoidAmountDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidTaxParamsBodyClientReferenceInformation void tax params body client reference information
swagger:model VoidTaxParamsBodyClientReferenceInformation
*/
type VoidTaxParamsBodyClientReferenceInformation struct {

	// Merchant-generated order reference or tracking number. It is recommended that you send a unique value for each
	// transaction so that you can perform meaningful searches for the transaction.
	//
	// #### Used by
	// **Authorization**
	// Required field.
	//
	// #### PIN Debit
	// Requests for PIN debit reversals need to use the same merchant reference number that was used in the transaction that is being
	// reversed.
	//
	// Required field for all PIN Debit requests (purchase, credit, and reversal).
	//
	// #### FDC Nashville Global
	// Certain circumstances can cause the processor to truncate this value to 15 or 17 characters for Level II and Level III processing, which can cause a discrepancy between the value you submit and the value included in some processor reports.
	//
	// Max Length: 50
	Code string `json:"code,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// partner
	Partner *VoidTaxParamsBodyClientReferenceInformationPartner `json:"partner,omitempty"`
}

// Validate validates this void tax params body client reference information
func (o *VoidTaxParamsBodyClientReferenceInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePartner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidTaxParamsBodyClientReferenceInformation) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(o.Code) { // not required
		return nil
	}

	if err := validate.MaxLength("voidTaxRequest"+"."+"clientReferenceInformation"+"."+"code", "body", string(o.Code), 50); err != nil {
		return err
	}

	return nil
}

func (o *VoidTaxParamsBodyClientReferenceInformation) validatePartner(formats strfmt.Registry) error {

	if swag.IsZero(o.Partner) { // not required
		return nil
	}

	if o.Partner != nil {
		if err := o.Partner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidTaxRequest" + "." + "clientReferenceInformation" + "." + "partner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidTaxParamsBodyClientReferenceInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidTaxParamsBodyClientReferenceInformation) UnmarshalBinary(b []byte) error {
	var res VoidTaxParamsBodyClientReferenceInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidTaxParamsBodyClientReferenceInformationPartner void tax params body client reference information partner
swagger:model VoidTaxParamsBodyClientReferenceInformationPartner
*/
type VoidTaxParamsBodyClientReferenceInformationPartner struct {

	// Identifier for the developer that helped integrate a partner solution to CyberSource.
	//
	// Send this value in all requests that are sent through the partner solutions built by that developer.
	// CyberSource assigns the ID to the developer.
	//
	// **Note** When you see a developer ID of 999 in reports, the developer ID that was submitted is incorrect.
	//
	// Max Length: 8
	DeveloperID string `json:"developerId,omitempty"`

	// Identifier for the partner that is integrated to CyberSource.
	//
	// Send this value in all requests that are sent through the partner solution. CyberSource assigns the ID to the partner.
	//
	// **Note** When you see a solutionId of 999 in reports, the solutionId that was submitted is incorrect.
	//
	// Max Length: 8
	SolutionID string `json:"solutionId,omitempty"`
}

// Validate validates this void tax params body client reference information partner
func (o *VoidTaxParamsBodyClientReferenceInformationPartner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDeveloperID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolutionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidTaxParamsBodyClientReferenceInformationPartner) validateDeveloperID(formats strfmt.Registry) error {

	if swag.IsZero(o.DeveloperID) { // not required
		return nil
	}

	if err := validate.MaxLength("voidTaxRequest"+"."+"clientReferenceInformation"+"."+"partner"+"."+"developerId", "body", string(o.DeveloperID), 8); err != nil {
		return err
	}

	return nil
}

func (o *VoidTaxParamsBodyClientReferenceInformationPartner) validateSolutionID(formats strfmt.Registry) error {

	if swag.IsZero(o.SolutionID) { // not required
		return nil
	}

	if err := validate.MaxLength("voidTaxRequest"+"."+"clientReferenceInformation"+"."+"partner"+"."+"solutionId", "body", string(o.SolutionID), 8); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidTaxParamsBodyClientReferenceInformationPartner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidTaxParamsBodyClientReferenceInformationPartner) UnmarshalBinary(b []byte) error {
	var res VoidTaxParamsBodyClientReferenceInformationPartner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}