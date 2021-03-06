// Code generated by go-swagger; DO NOT EDIT.

package net_fundings

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

// GetNetFundingDetailsReader is a Reader for the GetNetFundingDetails structure.
type GetNetFundingDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetFundingDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetFundingDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNetFundingDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNetFundingDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNetFundingDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNetFundingDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNetFundingDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetFundingDetailsOK creates a GetNetFundingDetailsOK with default headers values
func NewGetNetFundingDetailsOK() *GetNetFundingDetailsOK {
	return &GetNetFundingDetailsOK{}
}

/*GetNetFundingDetailsOK handles this case with default header values.

Ok
*/
type GetNetFundingDetailsOK struct {
	Payload *GetNetFundingDetailsOKBody
}

func (o *GetNetFundingDetailsOK) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/net-fundings][%d] getNetFundingDetailsOK  %+v", 200, o.Payload)
}

func (o *GetNetFundingDetailsOK) GetPayload() *GetNetFundingDetailsOKBody {
	return o.Payload
}

func (o *GetNetFundingDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetFundingDetailsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetFundingDetailsBadRequest creates a GetNetFundingDetailsBadRequest with default headers values
func NewGetNetFundingDetailsBadRequest() *GetNetFundingDetailsBadRequest {
	return &GetNetFundingDetailsBadRequest{}
}

/*GetNetFundingDetailsBadRequest handles this case with default header values.

Invalid request
*/
type GetNetFundingDetailsBadRequest struct {
	Payload *GetNetFundingDetailsBadRequestBody
}

func (o *GetNetFundingDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/net-fundings][%d] getNetFundingDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetNetFundingDetailsBadRequest) GetPayload() *GetNetFundingDetailsBadRequestBody {
	return o.Payload
}

func (o *GetNetFundingDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetFundingDetailsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetFundingDetailsUnauthorized creates a GetNetFundingDetailsUnauthorized with default headers values
func NewGetNetFundingDetailsUnauthorized() *GetNetFundingDetailsUnauthorized {
	return &GetNetFundingDetailsUnauthorized{}
}

/*GetNetFundingDetailsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetNetFundingDetailsUnauthorized struct {
	Payload interface{}
}

func (o *GetNetFundingDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/net-fundings][%d] getNetFundingDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNetFundingDetailsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetFundingDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetFundingDetailsForbidden creates a GetNetFundingDetailsForbidden with default headers values
func NewGetNetFundingDetailsForbidden() *GetNetFundingDetailsForbidden {
	return &GetNetFundingDetailsForbidden{}
}

/*GetNetFundingDetailsForbidden handles this case with default header values.

Forbidden
*/
type GetNetFundingDetailsForbidden struct {
	Payload interface{}
}

func (o *GetNetFundingDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/net-fundings][%d] getNetFundingDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetNetFundingDetailsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetFundingDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetFundingDetailsNotFound creates a GetNetFundingDetailsNotFound with default headers values
func NewGetNetFundingDetailsNotFound() *GetNetFundingDetailsNotFound {
	return &GetNetFundingDetailsNotFound{}
}

/*GetNetFundingDetailsNotFound handles this case with default header values.

Report not found
*/
type GetNetFundingDetailsNotFound struct {
	Payload interface{}
}

func (o *GetNetFundingDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/net-fundings][%d] getNetFundingDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetNetFundingDetailsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetFundingDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetFundingDetailsInternalServerError creates a GetNetFundingDetailsInternalServerError with default headers values
func NewGetNetFundingDetailsInternalServerError() *GetNetFundingDetailsInternalServerError {
	return &GetNetFundingDetailsInternalServerError{}
}

/*GetNetFundingDetailsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetNetFundingDetailsInternalServerError struct {
	Payload *GetNetFundingDetailsInternalServerErrorBody
}

func (o *GetNetFundingDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/net-fundings][%d] getNetFundingDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNetFundingDetailsInternalServerError) GetPayload() *GetNetFundingDetailsInternalServerErrorBody {
	return o.Payload
}

func (o *GetNetFundingDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetFundingDetailsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DetailsItems0 Provides failed validation input field detail
//
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// Field in request that caused an error
	//
	Field string `json:"field,omitempty"`

	// Documented reason code
	//
	Reason string `json:"reason,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNetFundingDetailsBadRequestBody reportingV3NetFundingsGet400Response
//
// HTTP status code for client application
swagger:model GetNetFundingDetailsBadRequestBody
*/
type GetNetFundingDetailsBadRequestBody struct {

	// Error field list
	//
	// Required: true
	Details []*DetailsItems0 `json:"details"`

	// Short descriptive message to the user.
	//
	// Required: true
	Message *string `json:"message"`

	// Documented reason code
	//
	// Required: true
	Reason *string `json:"reason"`

	// Time of request in UTC.
	//
	// Required: true
	// Format: date-time
	SubmitTimeUtc *strfmt.DateTime `json:"submitTimeUtc"`
}

// Validate validates this get net funding details bad request body
func (o *GetNetFundingDetailsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmitTimeUtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetFundingDetailsBadRequestBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("getNetFundingDetailsBadRequest"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetFundingDetailsBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetFundingDetailsBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getNetFundingDetailsBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *GetNetFundingDetailsBadRequestBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("getNetFundingDetailsBadRequest"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *GetNetFundingDetailsBadRequestBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("getNetFundingDetailsBadRequest"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("getNetFundingDetailsBadRequest"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetFundingDetailsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetFundingDetailsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetNetFundingDetailsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNetFundingDetailsInternalServerErrorBody reportingV3NetFundingsGet500Response
//
// HTTP status code for client application
swagger:model GetNetFundingDetailsInternalServerErrorBody
*/
type GetNetFundingDetailsInternalServerErrorBody struct {

	// Error field list
	//
	// Required: true
	Details []*DetailsItems0 `json:"details"`

	// Short descriptive message to the user.
	//
	// Required: true
	Message *string `json:"message"`

	// Documented reason code
	//
	// Required: true
	Reason *string `json:"reason"`

	// Time of request in UTC.
	//
	// Required: true
	// Format: date-time
	SubmitTimeUtc *strfmt.DateTime `json:"submitTimeUtc"`
}

// Validate validates this get net funding details internal server error body
func (o *GetNetFundingDetailsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmitTimeUtc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetFundingDetailsInternalServerErrorBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("getNetFundingDetailsInternalServerError"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetFundingDetailsInternalServerError" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetFundingDetailsInternalServerErrorBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getNetFundingDetailsInternalServerError"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *GetNetFundingDetailsInternalServerErrorBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("getNetFundingDetailsInternalServerError"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *GetNetFundingDetailsInternalServerErrorBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("getNetFundingDetailsInternalServerError"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("getNetFundingDetailsInternalServerError"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetFundingDetailsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetFundingDetailsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetNetFundingDetailsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNetFundingDetailsOKBody reportingV3NetFundingsGet200Response
swagger:model GetNetFundingDetailsOKBody
*/
type GetNetFundingDetailsOKBody struct {

	// Valid report End Date in **ISO 8601 format**
	// **Example date format:**
	// - yyyy-MM-dd'T'HH:mm:ss.SSSZZ
	//
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty" xml:"endTime"`

	// List of Netfunding summary objects
	NetFundingSummaries []*NetFundingSummariesItems0 `json:"netFundingSummaries" xml:"NetFundingSummaries"`

	// List of new total currency wise
	NetTotal []*NetTotalItems0 `json:"netTotal" xml:"netTotal"`

	// Valid report Start Date in **ISO 8601 format**.
	// Please refer the following link to know more about ISO 8601 format.[Rfc Date Format](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14)
	//
	// **Example:**
	// - yyyy-MM-dd'T'HH:mm:ss.SSSZZ
	//
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty" xml:"startTime"`

	// List of total chargebacks currency wise
	TotalChargebacks []*TotalChargebacksItems0 `json:"totalChargebacks" xml:"totalChargebacks"`

	// List of total fees currency wise
	TotalFees []*TotalFeesItems0 `json:"totalFees" xml:"totalFees"`

	// List of total purchases currency wise
	TotalPurchases []*TotalPurchasesItems0 `json:"totalPurchases" xml:"totalPurchases"`

	// List of total refunds currency wise
	TotalRefunds []*TotalRefundsItems0 `json:"totalRefunds" xml:"totalRefunds"`
}

// Validate validates this get net funding details o k body
func (o *GetNetFundingDetailsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetFundingSummaries(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotalChargebacks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotalFees(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotalPurchases(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotalRefunds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetFundingDetailsOKBody) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(o.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("getNetFundingDetailsOK"+"."+"endTime", "body", "date-time", o.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetNetFundingDetailsOKBody) validateNetFundingSummaries(formats strfmt.Registry) error {

	if swag.IsZero(o.NetFundingSummaries) { // not required
		return nil
	}

	for i := 0; i < len(o.NetFundingSummaries); i++ {
		if swag.IsZero(o.NetFundingSummaries[i]) { // not required
			continue
		}

		if o.NetFundingSummaries[i] != nil {
			if err := o.NetFundingSummaries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetFundingDetailsOK" + "." + "netFundingSummaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetFundingDetailsOKBody) validateNetTotal(formats strfmt.Registry) error {

	if swag.IsZero(o.NetTotal) { // not required
		return nil
	}

	for i := 0; i < len(o.NetTotal); i++ {
		if swag.IsZero(o.NetTotal[i]) { // not required
			continue
		}

		if o.NetTotal[i] != nil {
			if err := o.NetTotal[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetFundingDetailsOK" + "." + "netTotal" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetFundingDetailsOKBody) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(o.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("getNetFundingDetailsOK"+"."+"startTime", "body", "date-time", o.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetNetFundingDetailsOKBody) validateTotalChargebacks(formats strfmt.Registry) error {

	if swag.IsZero(o.TotalChargebacks) { // not required
		return nil
	}

	for i := 0; i < len(o.TotalChargebacks); i++ {
		if swag.IsZero(o.TotalChargebacks[i]) { // not required
			continue
		}

		if o.TotalChargebacks[i] != nil {
			if err := o.TotalChargebacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetFundingDetailsOK" + "." + "totalChargebacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetFundingDetailsOKBody) validateTotalFees(formats strfmt.Registry) error {

	if swag.IsZero(o.TotalFees) { // not required
		return nil
	}

	for i := 0; i < len(o.TotalFees); i++ {
		if swag.IsZero(o.TotalFees[i]) { // not required
			continue
		}

		if o.TotalFees[i] != nil {
			if err := o.TotalFees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetFundingDetailsOK" + "." + "totalFees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetFundingDetailsOKBody) validateTotalPurchases(formats strfmt.Registry) error {

	if swag.IsZero(o.TotalPurchases) { // not required
		return nil
	}

	for i := 0; i < len(o.TotalPurchases); i++ {
		if swag.IsZero(o.TotalPurchases[i]) { // not required
			continue
		}

		if o.TotalPurchases[i] != nil {
			if err := o.TotalPurchases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetFundingDetailsOK" + "." + "totalPurchases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetFundingDetailsOKBody) validateTotalRefunds(formats strfmt.Registry) error {

	if swag.IsZero(o.TotalRefunds) { // not required
		return nil
	}

	for i := 0; i < len(o.TotalRefunds); i++ {
		if swag.IsZero(o.TotalRefunds[i]) { // not required
			continue
		}

		if o.TotalRefunds[i] != nil {
			if err := o.TotalRefunds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetFundingDetailsOK" + "." + "totalRefunds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetFundingDetailsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetFundingDetailsOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetFundingDetailsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*NetFundingSummariesItems0 net funding summaries items0
swagger:model NetFundingSummariesItems0
*/
type NetFundingSummariesItems0 struct {

	// conveyed amount
	ConveyedAmount string `json:"conveyedAmount,omitempty"`

	// conveyed count
	ConveyedCount int64 `json:"conveyedCount,omitempty"`

	// Valid ISO 4217 ALPHA-3 currency code
	CurrencyCode string `json:"currencyCode,omitempty"`

	// funded amount
	FundedAmount string `json:"fundedAmount,omitempty"`

	// funded count
	FundedCount int64 `json:"fundedCount,omitempty"`

	// payment sub type
	PaymentSubType string `json:"paymentSubType,omitempty"`

	// settled count
	SettledCount int64 `json:"settledCount,omitempty"`

	// Valid values:
	// - PURCHASES
	// - REFUNDS
	// - FEES
	// - CHARGEBACKS
	//
	Type string `json:"type,omitempty"`
}

// Validate validates this net funding summaries items0
func (o *NetFundingSummariesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetFundingSummariesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetFundingSummariesItems0) UnmarshalBinary(b []byte) error {
	var res NetFundingSummariesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*NetTotalItems0 net total items0
swagger:model NetTotalItems0
*/
type NetTotalItems0 struct {

	// Valid ISO 4217 ALPHA-3 currency code
	// Required: true
	Currency *string `json:"currency"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this net total items0
func (o *NetTotalItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetTotalItems0) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", o.Currency); err != nil {
		return err
	}

	return nil
}

func (o *NetTotalItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetTotalItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetTotalItems0) UnmarshalBinary(b []byte) error {
	var res NetTotalItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TotalChargebacksItems0 total chargebacks items0
swagger:model TotalChargebacksItems0
*/
type TotalChargebacksItems0 struct {

	// Valid ISO 4217 ALPHA-3 currency code
	// Required: true
	Currency *string `json:"currency"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this total chargebacks items0
func (o *TotalChargebacksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotalChargebacksItems0) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", o.Currency); err != nil {
		return err
	}

	return nil
}

func (o *TotalChargebacksItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TotalChargebacksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TotalChargebacksItems0) UnmarshalBinary(b []byte) error {
	var res TotalChargebacksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TotalFeesItems0 total fees items0
swagger:model TotalFeesItems0
*/
type TotalFeesItems0 struct {

	// Valid ISO 4217 ALPHA-3 currency code
	// Required: true
	Currency *string `json:"currency"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this total fees items0
func (o *TotalFeesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotalFeesItems0) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", o.Currency); err != nil {
		return err
	}

	return nil
}

func (o *TotalFeesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TotalFeesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TotalFeesItems0) UnmarshalBinary(b []byte) error {
	var res TotalFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TotalPurchasesItems0 total purchases items0
swagger:model TotalPurchasesItems0
*/
type TotalPurchasesItems0 struct {

	// Valid ISO 4217 ALPHA-3 currency code
	// Required: true
	Currency *string `json:"currency"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this total purchases items0
func (o *TotalPurchasesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotalPurchasesItems0) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", o.Currency); err != nil {
		return err
	}

	return nil
}

func (o *TotalPurchasesItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TotalPurchasesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TotalPurchasesItems0) UnmarshalBinary(b []byte) error {
	var res TotalPurchasesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TotalRefundsItems0 total refunds items0
swagger:model TotalRefundsItems0
*/
type TotalRefundsItems0 struct {

	// Valid ISO 4217 ALPHA-3 currency code
	// Required: true
	Currency *string `json:"currency"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this total refunds items0
func (o *TotalRefundsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TotalRefundsItems0) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", o.Currency); err != nil {
		return err
	}

	return nil
}

func (o *TotalRefundsItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TotalRefundsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TotalRefundsItems0) UnmarshalBinary(b []byte) error {
	var res TotalRefundsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
