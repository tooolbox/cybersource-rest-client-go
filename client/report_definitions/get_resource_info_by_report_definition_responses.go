// Code generated by go-swagger; DO NOT EDIT.

package report_definitions

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

// GetResourceInfoByReportDefinitionReader is a Reader for the GetResourceInfoByReportDefinition structure.
type GetResourceInfoByReportDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceInfoByReportDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceInfoByReportDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResourceInfoByReportDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceInfoByReportDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceInfoByReportDefinitionOK creates a GetResourceInfoByReportDefinitionOK with default headers values
func NewGetResourceInfoByReportDefinitionOK() *GetResourceInfoByReportDefinitionOK {
	return &GetResourceInfoByReportDefinitionOK{}
}

/*GetResourceInfoByReportDefinitionOK handles this case with default header values.

Ok
*/
type GetResourceInfoByReportDefinitionOK struct {
	Payload *GetResourceInfoByReportDefinitionOKBody
}

func (o *GetResourceInfoByReportDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/report-definitions/{reportDefinitionName}][%d] getResourceInfoByReportDefinitionOK  %+v", 200, o.Payload)
}

func (o *GetResourceInfoByReportDefinitionOK) GetPayload() *GetResourceInfoByReportDefinitionOKBody {
	return o.Payload
}

func (o *GetResourceInfoByReportDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetResourceInfoByReportDefinitionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceInfoByReportDefinitionBadRequest creates a GetResourceInfoByReportDefinitionBadRequest with default headers values
func NewGetResourceInfoByReportDefinitionBadRequest() *GetResourceInfoByReportDefinitionBadRequest {
	return &GetResourceInfoByReportDefinitionBadRequest{}
}

/*GetResourceInfoByReportDefinitionBadRequest handles this case with default header values.

Invalid request
*/
type GetResourceInfoByReportDefinitionBadRequest struct {
	Payload *GetResourceInfoByReportDefinitionBadRequestBody
}

func (o *GetResourceInfoByReportDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/report-definitions/{reportDefinitionName}][%d] getResourceInfoByReportDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *GetResourceInfoByReportDefinitionBadRequest) GetPayload() *GetResourceInfoByReportDefinitionBadRequestBody {
	return o.Payload
}

func (o *GetResourceInfoByReportDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetResourceInfoByReportDefinitionBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceInfoByReportDefinitionNotFound creates a GetResourceInfoByReportDefinitionNotFound with default headers values
func NewGetResourceInfoByReportDefinitionNotFound() *GetResourceInfoByReportDefinitionNotFound {
	return &GetResourceInfoByReportDefinitionNotFound{}
}

/*GetResourceInfoByReportDefinitionNotFound handles this case with default header values.

Report not found
*/
type GetResourceInfoByReportDefinitionNotFound struct {
}

func (o *GetResourceInfoByReportDefinitionNotFound) Error() string {
	return fmt.Sprintf("[GET /reporting/v3/report-definitions/{reportDefinitionName}][%d] getResourceInfoByReportDefinitionNotFound ", 404)
}

func (o *GetResourceInfoByReportDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetResourceInfoByReportDefinitionBadRequestBody reportingV3ReportDefinitionsNameGet400Response
//
// HTTP status code for client application
swagger:model GetResourceInfoByReportDefinitionBadRequestBody
*/
type GetResourceInfoByReportDefinitionBadRequestBody struct {

	// Error field list
	//
	// Required: true
	Details []*GetResourceInfoByReportDefinitionBadRequestBodyDetailsItems0 `json:"details"`

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

// Validate validates this get resource info by report definition bad request body
func (o *GetResourceInfoByReportDefinitionBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *GetResourceInfoByReportDefinitionBadRequestBody) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("getResourceInfoByReportDefinitionBadRequest"+"."+"details", "body", o.Details); err != nil {
		return err
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getResourceInfoByReportDefinitionBadRequest" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetResourceInfoByReportDefinitionBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("getResourceInfoByReportDefinitionBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *GetResourceInfoByReportDefinitionBadRequestBody) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("getResourceInfoByReportDefinitionBadRequest"+"."+"reason", "body", o.Reason); err != nil {
		return err
	}

	return nil
}

func (o *GetResourceInfoByReportDefinitionBadRequestBody) validateSubmitTimeUtc(formats strfmt.Registry) error {

	if err := validate.Required("getResourceInfoByReportDefinitionBadRequest"+"."+"submitTimeUtc", "body", o.SubmitTimeUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("getResourceInfoByReportDefinitionBadRequest"+"."+"submitTimeUtc", "body", "date-time", o.SubmitTimeUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetResourceInfoByReportDefinitionBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourceInfoByReportDefinitionBadRequestBodyDetailsItems0 Provides failed validation input field detail
//
swagger:model GetResourceInfoByReportDefinitionBadRequestBodyDetailsItems0
*/
type GetResourceInfoByReportDefinitionBadRequestBodyDetailsItems0 struct {

	// Field in request that caused an error
	//
	Field string `json:"field,omitempty"`

	// Documented reason code
	//
	Reason string `json:"reason,omitempty"`
}

// Validate validates this get resource info by report definition bad request body details items0
func (o *GetResourceInfoByReportDefinitionBadRequestBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionBadRequestBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionBadRequestBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetResourceInfoByReportDefinitionBadRequestBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourceInfoByReportDefinitionOKBody reportingV3ReportDefinitionsNameGet200Response
swagger:model GetResourceInfoByReportDefinitionOKBody
*/
type GetResourceInfoByReportDefinitionOKBody struct {

	// attributes
	Attributes []*GetResourceInfoByReportDefinitionOKBodyAttributesItems0 `json:"attributes"`

	// default settings
	DefaultSettings *GetResourceInfoByReportDefinitionOKBodyDefaultSettings `json:"defaultSettings,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// report definition Id
	ReportDefinitionID int32 `json:"reportDefinitionId,omitempty"`

	// report defintion name
	ReportDefintionName string `json:"reportDefintionName,omitempty"`

	// 'The subscription type for which report definition is required. By default the type will be CUSTOM.'
	// Valid Values:
	// - 'CLASSIC'
	// - 'CUSTOM'
	// - 'STANDARD'
	//
	SubscriptionType string `json:"subscriptionType,omitempty"`

	// supported formats
	// Unique: true
	SupportedFormats []string `json:"supportedFormats"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this get resource info by report definition o k body
func (o *GetResourceInfoByReportDefinitionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefaultSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSupportedFormats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourceInfoByReportDefinitionOKBody) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(o.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(o.Attributes); i++ {
		if swag.IsZero(o.Attributes[i]) { // not required
			continue
		}

		if o.Attributes[i] != nil {
			if err := o.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getResourceInfoByReportDefinitionOK" + "." + "attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetResourceInfoByReportDefinitionOKBody) validateDefaultSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.DefaultSettings) { // not required
		return nil
	}

	if o.DefaultSettings != nil {
		if err := o.DefaultSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getResourceInfoByReportDefinitionOK" + "." + "defaultSettings")
			}
			return err
		}
	}

	return nil
}

func (o *GetResourceInfoByReportDefinitionOKBody) validateSupportedFormats(formats strfmt.Registry) error {

	if swag.IsZero(o.SupportedFormats) { // not required
		return nil
	}

	if err := validate.UniqueItems("getResourceInfoByReportDefinitionOK"+"."+"supportedFormats", "body", o.SupportedFormats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionOKBody) UnmarshalBinary(b []byte) error {
	var res GetResourceInfoByReportDefinitionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourceInfoByReportDefinitionOKBodyAttributesItems0 get resource info by report definition o k body attributes items0
swagger:model GetResourceInfoByReportDefinitionOKBodyAttributesItems0
*/
type GetResourceInfoByReportDefinitionOKBodyAttributesItems0 struct {

	// default
	Default bool `json:"default,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Attribute Filter Type.
	FilterType string `json:"filterType,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// required
	Required bool `json:"required,omitempty"`

	// Valid values for the filter.
	SupportedType string `json:"supportedType,omitempty"`
}

// Validate validates this get resource info by report definition o k body attributes items0
func (o *GetResourceInfoByReportDefinitionOKBodyAttributesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionOKBodyAttributesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionOKBodyAttributesItems0) UnmarshalBinary(b []byte) error {
	var res GetResourceInfoByReportDefinitionOKBodyAttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourceInfoByReportDefinitionOKBodyDefaultSettings get resource info by report definition o k body default settings
swagger:model GetResourceInfoByReportDefinitionOKBodyDefaultSettings
*/
type GetResourceInfoByReportDefinitionOKBodyDefaultSettings struct {

	// List of filters to apply
	ReportFilters map[string][]string `json:"reportFilters,omitempty"`

	// Report Frequency Value
	// Valid Values:
	//   - DAILY
	//   - WEEKLY
	//   - MONTHLY
	//   - ADHOC
	//
	ReportFrequency string `json:"reportFrequency,omitempty"`

	// Report Format
	// Valid values:
	//   - application/xml
	//   - text/csv
	//
	ReportMimeType string `json:"reportMimeType,omitempty"`

	// Report Name
	ReportName string `json:"reportName,omitempty"`

	// report preferences
	ReportPreferences *GetResourceInfoByReportDefinitionOKBodyDefaultSettingsReportPreferences `json:"reportPreferences,omitempty"`

	// Start Day
	StartDay int32 `json:"startDay,omitempty"`

	// Start Time
	StartTime string `json:"startTime,omitempty"`

	// Time Zone
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this get resource info by report definition o k body default settings
func (o *GetResourceInfoByReportDefinitionOKBodyDefaultSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReportPreferences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourceInfoByReportDefinitionOKBodyDefaultSettings) validateReportPreferences(formats strfmt.Registry) error {

	if swag.IsZero(o.ReportPreferences) { // not required
		return nil
	}

	if o.ReportPreferences != nil {
		if err := o.ReportPreferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getResourceInfoByReportDefinitionOK" + "." + "defaultSettings" + "." + "reportPreferences")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionOKBodyDefaultSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionOKBodyDefaultSettings) UnmarshalBinary(b []byte) error {
	var res GetResourceInfoByReportDefinitionOKBodyDefaultSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourceInfoByReportDefinitionOKBodyDefaultSettingsReportPreferences Report Preferences
swagger:model GetResourceInfoByReportDefinitionOKBodyDefaultSettingsReportPreferences
*/
type GetResourceInfoByReportDefinitionOKBodyDefaultSettingsReportPreferences struct {

	// Specify the field naming convention to be followed in reports (applicable to only csv report formats)
	//
	// Valid values:
	// - SOAPI
	// - SCMP
	//
	FieldNameConvention string `json:"fieldNameConvention,omitempty"`

	// Indicator to determine whether negative sign infront of amount for all refunded transaction
	SignedAmounts bool `json:"signedAmounts,omitempty"`
}

// Validate validates this get resource info by report definition o k body default settings report preferences
func (o *GetResourceInfoByReportDefinitionOKBodyDefaultSettingsReportPreferences) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionOKBodyDefaultSettingsReportPreferences) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceInfoByReportDefinitionOKBodyDefaultSettingsReportPreferences) UnmarshalBinary(b []byte) error {
	var res GetResourceInfoByReportDefinitionOKBodyDefaultSettingsReportPreferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
