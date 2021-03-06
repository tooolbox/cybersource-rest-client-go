// Code generated by go-swagger; DO NOT EDIT.

package void

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// VoidPaymentReader is a Reader for the VoidPayment structure.
type VoidPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoidPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVoidPaymentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVoidPaymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewVoidPaymentBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVoidPaymentCreated creates a VoidPaymentCreated with default headers values
func NewVoidPaymentCreated() *VoidPaymentCreated {
	return &VoidPaymentCreated{}
}

/*VoidPaymentCreated handles this case with default header values.

Successful response.
*/
type VoidPaymentCreated struct {
	Payload *VoidPaymentCreatedBody
}

func (o *VoidPaymentCreated) Error() string {
	return fmt.Sprintf("[POST /pts/v2/payments/{id}/voids][%d] voidPaymentCreated  %+v", 201, o.Payload)
}

func (o *VoidPaymentCreated) GetPayload() *VoidPaymentCreatedBody {
	return o.Payload
}

func (o *VoidPaymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VoidPaymentCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoidPaymentBadRequest creates a VoidPaymentBadRequest with default headers values
func NewVoidPaymentBadRequest() *VoidPaymentBadRequest {
	return &VoidPaymentBadRequest{}
}

/*VoidPaymentBadRequest handles this case with default header values.

Invalid request.
*/
type VoidPaymentBadRequest struct {
	Payload *VoidPaymentBadRequestBody
}

func (o *VoidPaymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /pts/v2/payments/{id}/voids][%d] voidPaymentBadRequest  %+v", 400, o.Payload)
}

func (o *VoidPaymentBadRequest) GetPayload() *VoidPaymentBadRequestBody {
	return o.Payload
}

func (o *VoidPaymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VoidPaymentBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoidPaymentBadGateway creates a VoidPaymentBadGateway with default headers values
func NewVoidPaymentBadGateway() *VoidPaymentBadGateway {
	return &VoidPaymentBadGateway{}
}

/*VoidPaymentBadGateway handles this case with default header values.

Unexpected system error or system timeout.
*/
type VoidPaymentBadGateway struct {
	Payload *VoidPaymentBadGatewayBody
}

func (o *VoidPaymentBadGateway) Error() string {
	return fmt.Sprintf("[POST /pts/v2/payments/{id}/voids][%d] voidPaymentBadGateway  %+v", 502, o.Payload)
}

func (o *VoidPaymentBadGateway) GetPayload() *VoidPaymentBadGatewayBody {
	return o.Payload
}

func (o *VoidPaymentBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VoidPaymentBadGatewayBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VoidPaymentBadGatewayBody ptsV2PaymentsVoidsPost502Response
swagger:model VoidPaymentBadGatewayBody
*/
type VoidPaymentBadGatewayBody struct {

	// The detail message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of the status.
	//
	// Possible values:
	//  - SYSTEM_ERROR
	//  - SERVER_TIMEOUT
	//  - SERVICE_TIMEOUT
	//  - INVALID_OR_MISSING_CONFIG
	//
	Reason string `json:"reason,omitempty"`

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - SERVER_ERROR
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// Example `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The `T` separates the date and the
	// time. The `Z` indicates UTC.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this void payment bad gateway body
func (o *VoidPaymentBadGatewayBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentBadGatewayBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentBadGatewayBody) UnmarshalBinary(b []byte) error {
	var res VoidPaymentBadGatewayBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidPaymentBadRequestBody ptsV2PaymentsVoidsPost400Response
swagger:model VoidPaymentBadRequestBody
*/
type VoidPaymentBadRequestBody struct {

	// details
	Details []*DetailsItems0 `json:"details"`

	// The detail message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of the status.
	//
	// Possible values:
	//  - MISSING_FIELD
	//  - INVALID_DATA
	//  - DUPLICATE_REQUEST
	//  - INVALID_MERCHANT_CONFIGURATION
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
	// Example `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The `T` separates the date and the
	// time. The `Z` indicates UTC.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`
}

// Validate validates this void payment bad request body
func (o *VoidPaymentBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidPaymentBadRequestBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("voidPaymentBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentBadRequestBody) UnmarshalBinary(b []byte) error {
	var res VoidPaymentBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidPaymentBody void payment body
swagger:model VoidPaymentBody
*/
type VoidPaymentBody struct {

	// client reference information
	ClientReferenceInformation *VoidPaymentParamsBodyClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
}

// Validate validates this void payment body
func (o *VoidPaymentBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientReferenceInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidPaymentBody) validateClientReferenceInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientReferenceInformation) { // not required
		return nil
	}

	if o.ClientReferenceInformation != nil {
		if err := o.ClientReferenceInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidPaymentRequest" + "." + "clientReferenceInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentBody) UnmarshalBinary(b []byte) error {
	var res VoidPaymentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidPaymentCreatedBody ptsV2PaymentsVoidsPost201Response
swagger:model VoidPaymentCreatedBody
*/
type VoidPaymentCreatedBody struct {

	// links
	Links *VoidPaymentCreatedBodyLinks `json:"_links,omitempty"`

	// client reference information
	ClientReferenceInformation *VoidPaymentCreatedBodyClientReferenceInformation `json:"clientReferenceInformation,omitempty"`

	// An unique identification number assigned by CyberSource to identify the submitted request. It is also appended to the endpoint of the resource.
	// Max Length: 26
	ID string `json:"id,omitempty"`

	// The status of the submitted transaction.
	//
	// Possible values:
	//  - VOIDED
	//
	Status string `json:"status,omitempty"`

	// Time of request in UTC. Format: `YYYY-MM-DDThh:mm:ssZ`
	// Example `2016-08-11T22:47:57Z` equals August 11, 2016, at 22:47:57 (10:47:57 p.m.). The `T` separates the date and the
	// time. The `Z` indicates UTC.
	//
	SubmitTimeUtc string `json:"submitTimeUtc,omitempty"`

	// void amount details
	VoidAmountDetails *VoidPaymentCreatedBodyVoidAmountDetails `json:"voidAmountDetails,omitempty"`
}

// Validate validates this void payment created body
func (o *VoidPaymentCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

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

func (o *VoidPaymentCreatedBody) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidPaymentCreated" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *VoidPaymentCreatedBody) validateClientReferenceInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ClientReferenceInformation) { // not required
		return nil
	}

	if o.ClientReferenceInformation != nil {
		if err := o.ClientReferenceInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidPaymentCreated" + "." + "clientReferenceInformation")
			}
			return err
		}
	}

	return nil
}

func (o *VoidPaymentCreatedBody) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.MaxLength("voidPaymentCreated"+"."+"id", "body", string(o.ID), 26); err != nil {
		return err
	}

	return nil
}

func (o *VoidPaymentCreatedBody) validateVoidAmountDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.VoidAmountDetails) { // not required
		return nil
	}

	if o.VoidAmountDetails != nil {
		if err := o.VoidAmountDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidPaymentCreated" + "." + "voidAmountDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentCreatedBody) UnmarshalBinary(b []byte) error {
	var res VoidPaymentCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidPaymentCreatedBodyClientReferenceInformation void payment created body client reference information
swagger:model VoidPaymentCreatedBodyClientReferenceInformation
*/
type VoidPaymentCreatedBodyClientReferenceInformation struct {

	// Client-generated order reference or tracking number. CyberSource recommends that you send a unique value for each
	// transaction so that you can perform meaningful searches for the transaction.
	//
	// For information about tracking orders, see "Tracking and Reconciling Your Orders" in [Getting Started with CyberSource Advanced for the SCMP API.](https://apps.cybersource.com/library/documentation/dev_guides/Getting_Started_SCMP/html/wwhelp/wwhimpl/js/html/wwhelp.htm)
	//
	// #### FDC Nashville Global
	// Certain circumstances can cause the processor to truncate this value to 15 or 17 characters for Level II and Level III processing, which can cause a discrepancy between the value you submit and the value included in some processor reports.
	//
	// Max Length: 50
	Code string `json:"code,omitempty"`

	// Merchant ID that was used to create the subscription or customer profile for which the service was requested.
	//
	// If your CyberSource account is enabled for Recurring Billing, this field is returned only if you are using
	// subscription sharing and if your merchant ID is in the same merchant ID pool as the owner merchant ID.
	//
	// If your CyberSource account is enabled for Payment Tokenization, this field is returned only if you are using
	// profile sharing and if your merchant ID is in the same merchant ID pool as the owner merchant ID.
	//
	// For details about how this field is used for Recurring Billing or Payment Tokenization, see the `ecp_debit_owner_merchant_id` field description in the [Electronic Check Services Using the SCMP API Guide.](https://apps.cybersource.com/library/documentation/dev_guides/EChecks_SCMP_API/html/wwhelp/wwhimpl/js/html/wwhelp.htm)
	//
	OwnerMerchantID string `json:"ownerMerchantId,omitempty"`

	// Date and time at your physical location.
	//
	// Format: `YYYYMMDDhhmmss`, where YYYY = year, MM = month, DD = day, hh = hour, mm = minutes ss = seconds
	//
	// Max Length: 14
	SubmitLocalDateTime string `json:"submitLocalDateTime,omitempty"`
}

// Validate validates this void payment created body client reference information
func (o *VoidPaymentCreatedBodyClientReferenceInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmitLocalDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidPaymentCreatedBodyClientReferenceInformation) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(o.Code) { // not required
		return nil
	}

	if err := validate.MaxLength("voidPaymentCreated"+"."+"clientReferenceInformation"+"."+"code", "body", string(o.Code), 50); err != nil {
		return err
	}

	return nil
}

func (o *VoidPaymentCreatedBodyClientReferenceInformation) validateSubmitLocalDateTime(formats strfmt.Registry) error {

	if swag.IsZero(o.SubmitLocalDateTime) { // not required
		return nil
	}

	if err := validate.MaxLength("voidPaymentCreated"+"."+"clientReferenceInformation"+"."+"submitLocalDateTime", "body", string(o.SubmitLocalDateTime), 14); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentCreatedBodyClientReferenceInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentCreatedBodyClientReferenceInformation) UnmarshalBinary(b []byte) error {
	var res VoidPaymentCreatedBodyClientReferenceInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidPaymentCreatedBodyLinks void payment created body links
swagger:model VoidPaymentCreatedBodyLinks
*/
type VoidPaymentCreatedBodyLinks struct {

	// self
	Self *VoidPaymentCreatedBodyLinksSelf `json:"self,omitempty"`
}

// Validate validates this void payment created body links
func (o *VoidPaymentCreatedBodyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidPaymentCreatedBodyLinks) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidPaymentCreated" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentCreatedBodyLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentCreatedBodyLinks) UnmarshalBinary(b []byte) error {
	var res VoidPaymentCreatedBodyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidPaymentCreatedBodyLinksSelf void payment created body links self
swagger:model VoidPaymentCreatedBodyLinksSelf
*/
type VoidPaymentCreatedBodyLinksSelf struct {

	// This is the endpoint of the resource that was created by the successful request.
	Href string `json:"href,omitempty"`

	// `method` refers to the HTTP method that you can send to the `self` endpoint to retrieve details of the resource.
	Method string `json:"method,omitempty"`
}

// Validate validates this void payment created body links self
func (o *VoidPaymentCreatedBodyLinksSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentCreatedBodyLinksSelf) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentCreatedBodyLinksSelf) UnmarshalBinary(b []byte) error {
	var res VoidPaymentCreatedBodyLinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidPaymentCreatedBodyVoidAmountDetails void payment created body void amount details
swagger:model VoidPaymentCreatedBodyVoidAmountDetails
*/
type VoidPaymentCreatedBodyVoidAmountDetails struct {

	// Currency used for the order. Use the three-character I[ISO Standard Currency Codes.](http://apps.cybersource.com/library/documentation/sbc/quickref/currencies.pdf)
	//
	// For details about currency as used in partial authorizations, see "Features for Debit Cards and Prepaid Cards" in the [Credit Card Services Using the SCMP API Guide](https://apps.cybersource.com/library/documentation/dev_guides/CC_Svcs_SCMP_API/html/wwhelp/wwhimpl/js/html/wwhelp.htm)
	//
	// For an authorization reversal (`reversalInformation`) or a capture (`processingOptions.capture` is set to `true`), you must use the same currency that you used in your payment authorization request.
	//
	// #### DCC for First Data
	// Your local currency. For details, see the `currency` field description in [Dynamic Currency Conversion For First Data Using the SCMP API](http://apps.cybersource.com/library/documentation/dev_guides/DCC_FirstData_SCMP/DCC_FirstData_SCMP_API.pdf).
	//
	// Max Length: 3
	Currency string `json:"currency,omitempty"`

	// Amount of the original transaction.
	OriginalTransactionAmount string `json:"originalTransactionAmount,omitempty"`

	// Total amount of the void.
	VoidAmount string `json:"voidAmount,omitempty"`
}

// Validate validates this void payment created body void amount details
func (o *VoidPaymentCreatedBodyVoidAmountDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VoidPaymentCreatedBodyVoidAmountDetails) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(o.Currency) { // not required
		return nil
	}

	if err := validate.MaxLength("voidPaymentCreated"+"."+"voidAmountDetails"+"."+"currency", "body", string(o.Currency), 3); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentCreatedBodyVoidAmountDetails) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentCreatedBodyVoidAmountDetails) UnmarshalBinary(b []byte) error {
	var res VoidPaymentCreatedBodyVoidAmountDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidPaymentParamsBodyClientReferenceInformation void payment params body client reference information
swagger:model VoidPaymentParamsBodyClientReferenceInformation
*/
type VoidPaymentParamsBodyClientReferenceInformation struct {

	// Client-generated order reference or tracking number. CyberSource recommends that you send a unique value for each
	// transaction so that you can perform meaningful searches for the transaction.
	//
	// For information about tracking orders, see "Tracking and Reconciling Your Orders" in [Getting Started with CyberSource Advanced for the SCMP API.](https://apps.cybersource.com/library/documentation/dev_guides/Getting_Started_SCMP/html/wwhelp/wwhimpl/js/html/wwhelp.htm)
	//
	// #### FDC Nashville Global
	// Certain circumstances can cause the processor to truncate this value to 15 or 17 characters for Level II and Level III processing, which can cause a discrepancy between the value you submit and the value included in some processor reports.
	//
	// Max Length: 50
	Code string `json:"code,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// partner
	Partner *VoidPaymentParamsBodyClientReferenceInformationPartner `json:"partner,omitempty"`
}

// Validate validates this void payment params body client reference information
func (o *VoidPaymentParamsBodyClientReferenceInformation) Validate(formats strfmt.Registry) error {
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

func (o *VoidPaymentParamsBodyClientReferenceInformation) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(o.Code) { // not required
		return nil
	}

	if err := validate.MaxLength("voidPaymentRequest"+"."+"clientReferenceInformation"+"."+"code", "body", string(o.Code), 50); err != nil {
		return err
	}

	return nil
}

func (o *VoidPaymentParamsBodyClientReferenceInformation) validatePartner(formats strfmt.Registry) error {

	if swag.IsZero(o.Partner) { // not required
		return nil
	}

	if o.Partner != nil {
		if err := o.Partner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voidPaymentRequest" + "." + "clientReferenceInformation" + "." + "partner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentParamsBodyClientReferenceInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentParamsBodyClientReferenceInformation) UnmarshalBinary(b []byte) error {
	var res VoidPaymentParamsBodyClientReferenceInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VoidPaymentParamsBodyClientReferenceInformationPartner void payment params body client reference information partner
swagger:model VoidPaymentParamsBodyClientReferenceInformationPartner
*/
type VoidPaymentParamsBodyClientReferenceInformationPartner struct {

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
	// **Note** When you see a partner ID of 999 in reports, the partner ID that was submitted is incorrect.
	//
	// Max Length: 8
	SolutionID string `json:"solutionId,omitempty"`
}

// Validate validates this void payment params body client reference information partner
func (o *VoidPaymentParamsBodyClientReferenceInformationPartner) Validate(formats strfmt.Registry) error {
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

func (o *VoidPaymentParamsBodyClientReferenceInformationPartner) validateDeveloperID(formats strfmt.Registry) error {

	if swag.IsZero(o.DeveloperID) { // not required
		return nil
	}

	if err := validate.MaxLength("voidPaymentRequest"+"."+"clientReferenceInformation"+"."+"partner"+"."+"developerId", "body", string(o.DeveloperID), 8); err != nil {
		return err
	}

	return nil
}

func (o *VoidPaymentParamsBodyClientReferenceInformationPartner) validateSolutionID(formats strfmt.Registry) error {

	if swag.IsZero(o.SolutionID) { // not required
		return nil
	}

	if err := validate.MaxLength("voidPaymentRequest"+"."+"clientReferenceInformation"+"."+"partner"+"."+"solutionId", "body", string(o.SolutionID), 8); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VoidPaymentParamsBodyClientReferenceInformationPartner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VoidPaymentParamsBodyClientReferenceInformationPartner) UnmarshalBinary(b []byte) error {
	var res VoidPaymentParamsBodyClientReferenceInformationPartner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
