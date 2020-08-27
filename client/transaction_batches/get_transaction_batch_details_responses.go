// Code generated by go-swagger; DO NOT EDIT.

package transaction_batches

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
)

// GetTransactionBatchDetailsReader is a Reader for the GetTransactionBatchDetails structure.
type GetTransactionBatchDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionBatchDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransactionBatchDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTransactionBatchDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTransactionBatchDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTransactionBatchDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTransactionBatchDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewGetTransactionBatchDetailsBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTransactionBatchDetailsOK creates a GetTransactionBatchDetailsOK with default headers values
func NewGetTransactionBatchDetailsOK() *GetTransactionBatchDetailsOK {
	return &GetTransactionBatchDetailsOK{}
}

/*GetTransactionBatchDetailsOK handles this case with default header values.

OK
*/
type GetTransactionBatchDetailsOK struct {
}

func (o *GetTransactionBatchDetailsOK) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batch-details/{id}][%d] getTransactionBatchDetailsOK ", 200)
}

func (o *GetTransactionBatchDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTransactionBatchDetailsBadRequest creates a GetTransactionBatchDetailsBadRequest with default headers values
func NewGetTransactionBatchDetailsBadRequest() *GetTransactionBatchDetailsBadRequest {
	return &GetTransactionBatchDetailsBadRequest{}
}

/*GetTransactionBatchDetailsBadRequest handles this case with default header values.

Bad Request
*/
type GetTransactionBatchDetailsBadRequest struct {
	Payload *GetTransactionBatchDetailsBadRequestBody
}

func (o *GetTransactionBatchDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batch-details/{id}][%d] getTransactionBatchDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransactionBatchDetailsBadRequest) GetPayload() *GetTransactionBatchDetailsBadRequestBody {
	return o.Payload
}

func (o *GetTransactionBatchDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchDetailsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBatchDetailsUnauthorized creates a GetTransactionBatchDetailsUnauthorized with default headers values
func NewGetTransactionBatchDetailsUnauthorized() *GetTransactionBatchDetailsUnauthorized {
	return &GetTransactionBatchDetailsUnauthorized{}
}

/*GetTransactionBatchDetailsUnauthorized handles this case with default header values.

Not Authorized
*/
type GetTransactionBatchDetailsUnauthorized struct {
	Payload *GetTransactionBatchDetailsUnauthorizedBody
}

func (o *GetTransactionBatchDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batch-details/{id}][%d] getTransactionBatchDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTransactionBatchDetailsUnauthorized) GetPayload() *GetTransactionBatchDetailsUnauthorizedBody {
	return o.Payload
}

func (o *GetTransactionBatchDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchDetailsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBatchDetailsForbidden creates a GetTransactionBatchDetailsForbidden with default headers values
func NewGetTransactionBatchDetailsForbidden() *GetTransactionBatchDetailsForbidden {
	return &GetTransactionBatchDetailsForbidden{}
}

/*GetTransactionBatchDetailsForbidden handles this case with default header values.

No Authenticated
*/
type GetTransactionBatchDetailsForbidden struct {
	Payload *GetTransactionBatchDetailsForbiddenBody
}

func (o *GetTransactionBatchDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batch-details/{id}][%d] getTransactionBatchDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetTransactionBatchDetailsForbidden) GetPayload() *GetTransactionBatchDetailsForbiddenBody {
	return o.Payload
}

func (o *GetTransactionBatchDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchDetailsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBatchDetailsNotFound creates a GetTransactionBatchDetailsNotFound with default headers values
func NewGetTransactionBatchDetailsNotFound() *GetTransactionBatchDetailsNotFound {
	return &GetTransactionBatchDetailsNotFound{}
}

/*GetTransactionBatchDetailsNotFound handles this case with default header values.

No Reports Found
*/
type GetTransactionBatchDetailsNotFound struct {
	Payload *GetTransactionBatchDetailsNotFoundBody
}

func (o *GetTransactionBatchDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batch-details/{id}][%d] getTransactionBatchDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetTransactionBatchDetailsNotFound) GetPayload() *GetTransactionBatchDetailsNotFoundBody {
	return o.Payload
}

func (o *GetTransactionBatchDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchDetailsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionBatchDetailsBadGateway creates a GetTransactionBatchDetailsBadGateway with default headers values
func NewGetTransactionBatchDetailsBadGateway() *GetTransactionBatchDetailsBadGateway {
	return &GetTransactionBatchDetailsBadGateway{}
}

/*GetTransactionBatchDetailsBadGateway handles this case with default header values.

Bad Gateway
*/
type GetTransactionBatchDetailsBadGateway struct {
	Payload *GetTransactionBatchDetailsBadGatewayBody
}

func (o *GetTransactionBatchDetailsBadGateway) Error() string {
	return fmt.Sprintf("[GET /pts/v1/transaction-batch-details/{id}][%d] getTransactionBatchDetailsBadGateway  %+v", 502, o.Payload)
}

func (o *GetTransactionBatchDetailsBadGateway) GetPayload() *GetTransactionBatchDetailsBadGatewayBody {
	return o.Payload
}

func (o *GetTransactionBatchDetailsBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTransactionBatchDetailsBadGatewayBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetTransactionBatchDetailsBadGatewayBody ptsV1TransactionBatchDetailsGet502Response
swagger:model GetTransactionBatchDetailsBadGatewayBody
*/
type GetTransactionBatchDetailsBadGatewayBody struct {

	// error information
	ErrorInformation *GetTransactionBatchDetailsBadGatewayBodyErrorInformation `json:"errorInformation,omitempty"`

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

// Validate validates this get transaction batch details bad gateway body
func (o *GetTransactionBatchDetailsBadGatewayBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchDetailsBadGatewayBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchDetailsBadGateway" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadGatewayBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadGatewayBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsBadGatewayBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsBadGatewayBodyErrorInformation get transaction batch details bad gateway body error information
swagger:model GetTransactionBatchDetailsBadGatewayBodyErrorInformation
*/
type GetTransactionBatchDetailsBadGatewayBodyErrorInformation struct {

	// The detailed message related to the status and reason listed above.
	Message string `json:"message,omitempty"`

	// The reason of status
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch details bad gateway body error information
func (o *GetTransactionBatchDetailsBadGatewayBodyErrorInformation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadGatewayBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadGatewayBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsBadGatewayBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsBadRequestBody ptsV1TransactionBatchDetailsGet400Response
swagger:model GetTransactionBatchDetailsBadRequestBody
*/
type GetTransactionBatchDetailsBadRequestBody struct {

	// error information
	ErrorInformation *GetTransactionBatchDetailsBadRequestBodyErrorInformation `json:"errorInformation,omitempty"`

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

// Validate validates this get transaction batch details bad request body
func (o *GetTransactionBatchDetailsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchDetailsBadRequestBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchDetailsBadRequest" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsBadRequestBodyErrorInformation get transaction batch details bad request body error information
swagger:model GetTransactionBatchDetailsBadRequestBodyErrorInformation
*/
type GetTransactionBatchDetailsBadRequestBodyErrorInformation struct {

	// details
	Details []*GetTransactionBatchDetailsBadRequestBodyErrorInformationDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch details bad request body error information
func (o *GetTransactionBatchDetailsBadRequestBodyErrorInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchDetailsBadRequestBodyErrorInformation) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("getTransactionBatchDetailsBadRequest" + "." + "errorInformation" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadRequestBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadRequestBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsBadRequestBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsBadRequestBodyErrorInformationDetailsItems0 get transaction batch details bad request body error information details items0
swagger:model GetTransactionBatchDetailsBadRequestBodyErrorInformationDetailsItems0
*/
type GetTransactionBatchDetailsBadRequestBodyErrorInformationDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	//
	Field string `json:"field,omitempty"`

	// The detailed message related to the status and reason listed above.
	//
	Message string `json:"message,omitempty"`
}

// Validate validates this get transaction batch details bad request body error information details items0
func (o *GetTransactionBatchDetailsBadRequestBodyErrorInformationDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadRequestBodyErrorInformationDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsBadRequestBodyErrorInformationDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsBadRequestBodyErrorInformationDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsForbiddenBody ptsV1TransactionBatchDetailsGet403Response
swagger:model GetTransactionBatchDetailsForbiddenBody
*/
type GetTransactionBatchDetailsForbiddenBody struct {

	// error information
	ErrorInformation *GetTransactionBatchDetailsForbiddenBodyErrorInformation `json:"errorInformation,omitempty"`

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

// Validate validates this get transaction batch details forbidden body
func (o *GetTransactionBatchDetailsForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchDetailsForbiddenBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchDetailsForbidden" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsForbiddenBodyErrorInformation get transaction batch details forbidden body error information
swagger:model GetTransactionBatchDetailsForbiddenBodyErrorInformation
*/
type GetTransactionBatchDetailsForbiddenBodyErrorInformation struct {

	// details
	Details []*GetTransactionBatchDetailsForbiddenBodyErrorInformationDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch details forbidden body error information
func (o *GetTransactionBatchDetailsForbiddenBodyErrorInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchDetailsForbiddenBodyErrorInformation) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("getTransactionBatchDetailsForbidden" + "." + "errorInformation" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsForbiddenBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsForbiddenBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsForbiddenBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsForbiddenBodyErrorInformationDetailsItems0 get transaction batch details forbidden body error information details items0
swagger:model GetTransactionBatchDetailsForbiddenBodyErrorInformationDetailsItems0
*/
type GetTransactionBatchDetailsForbiddenBodyErrorInformationDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	//
	Field string `json:"field,omitempty"`

	// The detailed message related to the status and reason listed above.
	//
	Message string `json:"message,omitempty"`
}

// Validate validates this get transaction batch details forbidden body error information details items0
func (o *GetTransactionBatchDetailsForbiddenBodyErrorInformationDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsForbiddenBodyErrorInformationDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsForbiddenBodyErrorInformationDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsForbiddenBodyErrorInformationDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsNotFoundBody ptsV1TransactionBatchDetailsGet404Response
swagger:model GetTransactionBatchDetailsNotFoundBody
*/
type GetTransactionBatchDetailsNotFoundBody struct {

	// error information
	ErrorInformation *GetTransactionBatchDetailsNotFoundBodyErrorInformation `json:"errorInformation,omitempty"`

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

// Validate validates this get transaction batch details not found body
func (o *GetTransactionBatchDetailsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchDetailsNotFoundBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchDetailsNotFound" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsNotFoundBodyErrorInformation get transaction batch details not found body error information
swagger:model GetTransactionBatchDetailsNotFoundBodyErrorInformation
*/
type GetTransactionBatchDetailsNotFoundBodyErrorInformation struct {

	// details
	Details []*GetTransactionBatchDetailsNotFoundBodyErrorInformationDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch details not found body error information
func (o *GetTransactionBatchDetailsNotFoundBodyErrorInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchDetailsNotFoundBodyErrorInformation) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("getTransactionBatchDetailsNotFound" + "." + "errorInformation" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsNotFoundBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsNotFoundBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsNotFoundBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsNotFoundBodyErrorInformationDetailsItems0 get transaction batch details not found body error information details items0
swagger:model GetTransactionBatchDetailsNotFoundBodyErrorInformationDetailsItems0
*/
type GetTransactionBatchDetailsNotFoundBodyErrorInformationDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	//
	Field string `json:"field,omitempty"`

	// The detailed message related to the status and reason listed above.
	//
	Message string `json:"message,omitempty"`
}

// Validate validates this get transaction batch details not found body error information details items0
func (o *GetTransactionBatchDetailsNotFoundBodyErrorInformationDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsNotFoundBodyErrorInformationDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsNotFoundBodyErrorInformationDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsNotFoundBodyErrorInformationDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsUnauthorizedBody ptsV1TransactionBatchDetailsGet401Response
swagger:model GetTransactionBatchDetailsUnauthorizedBody
*/
type GetTransactionBatchDetailsUnauthorizedBody struct {

	// error information
	ErrorInformation *GetTransactionBatchDetailsUnauthorizedBodyErrorInformation `json:"errorInformation,omitempty"`

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

// Validate validates this get transaction batch details unauthorized body
func (o *GetTransactionBatchDetailsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrorInformation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchDetailsUnauthorizedBody) validateErrorInformation(formats strfmt.Registry) error {

	if swag.IsZero(o.ErrorInformation) { // not required
		return nil
	}

	if o.ErrorInformation != nil {
		if err := o.ErrorInformation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTransactionBatchDetailsUnauthorized" + "." + "errorInformation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsUnauthorizedBodyErrorInformation get transaction batch details unauthorized body error information
swagger:model GetTransactionBatchDetailsUnauthorizedBodyErrorInformation
*/
type GetTransactionBatchDetailsUnauthorizedBodyErrorInformation struct {

	// details
	Details []*GetTransactionBatchDetailsUnauthorizedBodyErrorInformationDetailsItems0 `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get transaction batch details unauthorized body error information
func (o *GetTransactionBatchDetailsUnauthorizedBodyErrorInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTransactionBatchDetailsUnauthorizedBodyErrorInformation) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("getTransactionBatchDetailsUnauthorized" + "." + "errorInformation" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsUnauthorizedBodyErrorInformation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsUnauthorizedBodyErrorInformation) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsUnauthorizedBodyErrorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetTransactionBatchDetailsUnauthorizedBodyErrorInformationDetailsItems0 get transaction batch details unauthorized body error information details items0
swagger:model GetTransactionBatchDetailsUnauthorizedBodyErrorInformationDetailsItems0
*/
type GetTransactionBatchDetailsUnauthorizedBodyErrorInformationDetailsItems0 struct {

	// This is the flattened JSON object field name/path that is either missing or invalid.
	//
	Field string `json:"field,omitempty"`

	// The detailed message related to the status and reason listed above.
	//
	Message string `json:"message,omitempty"`
}

// Validate validates this get transaction batch details unauthorized body error information details items0
func (o *GetTransactionBatchDetailsUnauthorizedBodyErrorInformationDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetTransactionBatchDetailsUnauthorizedBodyErrorInformationDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTransactionBatchDetailsUnauthorizedBodyErrorInformationDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransactionBatchDetailsUnauthorizedBodyErrorInformationDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
