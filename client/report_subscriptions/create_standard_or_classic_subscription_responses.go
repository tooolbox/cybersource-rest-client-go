// Code generated by go-swagger; DO NOT EDIT.

package report_subscriptions

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

// CreateStandardOrClassicSubscriptionReader is a Reader for the CreateStandardOrClassicSubscription structure.
type CreateStandardOrClassicSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStandardOrClassicSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateStandardOrClassicSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateStandardOrClassicSubscriptionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateStandardOrClassicSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateStandardOrClassicSubscriptionOK creates a CreateStandardOrClassicSubscriptionOK with default headers values
func NewCreateStandardOrClassicSubscriptionOK() *CreateStandardOrClassicSubscriptionOK {
	return &CreateStandardOrClassicSubscriptionOK{}
}

/*CreateStandardOrClassicSubscriptionOK handles this case with default header values.

Ok
*/
type CreateStandardOrClassicSubscriptionOK struct {
}

func (o *CreateStandardOrClassicSubscriptionOK) Error() string {
	return fmt.Sprintf("[PUT /reporting/v3/predefined-report-subscriptions][%d] createStandardOrClassicSubscriptionOK ", 200)
}

func (o *CreateStandardOrClassicSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateStandardOrClassicSubscriptionCreated creates a CreateStandardOrClassicSubscriptionCreated with default headers values
func NewCreateStandardOrClassicSubscriptionCreated() *CreateStandardOrClassicSubscriptionCreated {
	return &CreateStandardOrClassicSubscriptionCreated{}
}

/*CreateStandardOrClassicSubscriptionCreated handles this case with default header values.

Created
*/
type CreateStandardOrClassicSubscriptionCreated struct {
}

func (o *CreateStandardOrClassicSubscriptionCreated) Error() string {
	return fmt.Sprintf("[PUT /reporting/v3/predefined-report-subscriptions][%d] createStandardOrClassicSubscriptionCreated ", 201)
}

func (o *CreateStandardOrClassicSubscriptionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateStandardOrClassicSubscriptionBadRequest creates a CreateStandardOrClassicSubscriptionBadRequest with default headers values
func NewCreateStandardOrClassicSubscriptionBadRequest() *CreateStandardOrClassicSubscriptionBadRequest {
	return &CreateStandardOrClassicSubscriptionBadRequest{}
}

/*CreateStandardOrClassicSubscriptionBadRequest handles this case with default header values.

Invalid request
*/
type CreateStandardOrClassicSubscriptionBadRequest struct {
	Payload *CreateStandardOrClassicSubscriptionBadRequestBody
}

func (o *CreateStandardOrClassicSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /reporting/v3/predefined-report-subscriptions][%d] createStandardOrClassicSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateStandardOrClassicSubscriptionBadRequest) GetPayload() *CreateStandardOrClassicSubscriptionBadRequestBody {
	return o.Payload
}

func (o *CreateStandardOrClassicSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateStandardOrClassicSubscriptionBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateStandardOrClassicSubscriptionBadRequestBody Error Bean
swagger:model CreateStandardOrClassicSubscriptionBadRequestBody
*/
type CreateStandardOrClassicSubscriptionBadRequestBody struct {

	// Error code
	// Required: true
	Code *string `json:"code"`

	// Correlation Id
	CorrelationID string `json:"correlationId,omitempty"`

	// Error Detail
	Detail string `json:"detail,omitempty"`

	// Error fields List
	Fields []*CreateStandardOrClassicSubscriptionBadRequestBodyFieldsItems0 `json:"fields"`

	// Localization Key Name
	LocalizationKey string `json:"localizationKey,omitempty"`

	// Error message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this create standard or classic subscription bad request body
func (o *CreateStandardOrClassicSubscriptionBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateStandardOrClassicSubscriptionBadRequestBody) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("createStandardOrClassicSubscriptionBadRequest"+"."+"code", "body", o.Code); err != nil {
		return err
	}

	return nil
}

func (o *CreateStandardOrClassicSubscriptionBadRequestBody) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		if o.Fields[i] != nil {
			if err := o.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createStandardOrClassicSubscriptionBadRequest" + "." + "fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateStandardOrClassicSubscriptionBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("createStandardOrClassicSubscriptionBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateStandardOrClassicSubscriptionBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateStandardOrClassicSubscriptionBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateStandardOrClassicSubscriptionBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateStandardOrClassicSubscriptionBadRequestBodyFieldsItems0 Provide validation failed input field details
swagger:model CreateStandardOrClassicSubscriptionBadRequestBodyFieldsItems0
*/
type CreateStandardOrClassicSubscriptionBadRequestBodyFieldsItems0 struct {

	// Localized Key Name
	LocalizationKey string `json:"localizationKey,omitempty"`

	// Error description about validation failed field
	Message string `json:"message,omitempty"`

	// Path of the failed property
	Path string `json:"path,omitempty"`
}

// Validate validates this create standard or classic subscription bad request body fields items0
func (o *CreateStandardOrClassicSubscriptionBadRequestBodyFieldsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateStandardOrClassicSubscriptionBadRequestBodyFieldsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateStandardOrClassicSubscriptionBadRequestBodyFieldsItems0) UnmarshalBinary(b []byte) error {
	var res CreateStandardOrClassicSubscriptionBadRequestBodyFieldsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateStandardOrClassicSubscriptionBody create standard or classic subscription body
swagger:model CreateStandardOrClassicSubscriptionBody
*/
type CreateStandardOrClassicSubscriptionBody struct {

	// Valid Report Definition Name
	// Required: true
	// Max Length: 80
	// Min Length: 1
	// Pattern: [a-zA-Z]+
	ReportDefinitionName *string `json:"reportDefinitionName"`

	// 'The frequency for which subscription is created. For Standard we can have DAILY, WEEKLY and MONTHLY. But for Classic we will have only DAILY.'
	// **NOTE: Do not document USER_DEFINED Frequency field in developer center**
	// Valid Values:
	// - 'DAILY'
	// - 'WEEKLY'
	// - 'MONTHLY'
	// - 'USER_DEFINED'
	//
	ReportFrequency string `json:"reportFrequency,omitempty"`

	// If the reportFrequency is User-defined, reportInterval should be in **ISO 8601 time format**
	// Please refer the following link to know more about ISO 8601 format.[Rfc Time Format](https://en.wikipedia.org/wiki/ISO_8601#Durations)
	//
	// **Example time format for 2 hours and 30 Mins:**
	//   - PT2H30M
	// **NOTE: Do not document reportInterval field in developer center**
	//
	// Pattern: ^PT((([1-9]|1[0-9]|2[0-3])H(([1-9]|[1-4][0-9]|5[0-9])M)?)|((([1-9]|1[0-9]|2[0-3])H)?([1-9]|[1-4][0-9]|5[0-9])M))$
	ReportInterval string `json:"reportInterval,omitempty"`

	// Report Format
	// Valid Values:
	//   - application/xml
	//   - text/csv
	//
	ReportMimeType string `json:"reportMimeType,omitempty"`

	// report name
	// Max Length: 128
	// Min Length: 1
	// Pattern: [a-zA-Z0-9-_ ]+
	ReportName string `json:"reportName,omitempty"`

	// This is the start day if the frequency is WEEKLY or MONTHLY. The value varies from 1-7 for WEEKLY and 1-31 for MONTHLY. For WEEKLY 1 means Sunday and 7 means Saturday. By default the value is 1.
	// Maximum: 31
	// Minimum: 1
	StartDay int64 `json:"startDay,omitempty"`

	// The hour at which the report generation should start. It should be in hhmm format. By Default it will be 0000. The format is 24 hours format.
	StartTime string `json:"startTime,omitempty"`

	// The status for subscription which is either created or updated. By default it is ACTIVE.
	// Valid Values:
	//   - ACTIVE
	//   - INACTIVE
	//
	SubscriptionStatus string `json:"subscriptionStatus,omitempty"`

	// The subscription type for which report definition is required. Valid values are CLASSIC and STANDARD.
	// Valid Values:
	//   - CLASSIC
	//   - STANDARD
	//
	// Required: true
	SubscriptionType *string `json:"subscriptionType"`

	// By Default the timezone for Standard subscription is PST. And for Classic subscription it will be GMT. If user provides any other time zone apart from PST for Standard subscription api should error out.
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this create standard or classic subscription body
func (o *CreateStandardOrClassicSubscriptionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReportDefinitionName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReportInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReportName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartDay(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubscriptionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateStandardOrClassicSubscriptionBody) validateReportDefinitionName(formats strfmt.Registry) error {

	if err := validate.Required("predefinedSubscriptionRequestBean"+"."+"reportDefinitionName", "body", o.ReportDefinitionName); err != nil {
		return err
	}

	if err := validate.MinLength("predefinedSubscriptionRequestBean"+"."+"reportDefinitionName", "body", string(*o.ReportDefinitionName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("predefinedSubscriptionRequestBean"+"."+"reportDefinitionName", "body", string(*o.ReportDefinitionName), 80); err != nil {
		return err
	}

	if err := validate.Pattern("predefinedSubscriptionRequestBean"+"."+"reportDefinitionName", "body", string(*o.ReportDefinitionName), `[a-zA-Z]+`); err != nil {
		return err
	}

	return nil
}

func (o *CreateStandardOrClassicSubscriptionBody) validateReportInterval(formats strfmt.Registry) error {

	if swag.IsZero(o.ReportInterval) { // not required
		return nil
	}

	if err := validate.Pattern("predefinedSubscriptionRequestBean"+"."+"reportInterval", "body", string(o.ReportInterval), `^PT((([1-9]|1[0-9]|2[0-3])H(([1-9]|[1-4][0-9]|5[0-9])M)?)|((([1-9]|1[0-9]|2[0-3])H)?([1-9]|[1-4][0-9]|5[0-9])M))$`); err != nil {
		return err
	}

	return nil
}

func (o *CreateStandardOrClassicSubscriptionBody) validateReportName(formats strfmt.Registry) error {

	if swag.IsZero(o.ReportName) { // not required
		return nil
	}

	if err := validate.MinLength("predefinedSubscriptionRequestBean"+"."+"reportName", "body", string(o.ReportName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("predefinedSubscriptionRequestBean"+"."+"reportName", "body", string(o.ReportName), 128); err != nil {
		return err
	}

	if err := validate.Pattern("predefinedSubscriptionRequestBean"+"."+"reportName", "body", string(o.ReportName), `[a-zA-Z0-9-_ ]+`); err != nil {
		return err
	}

	return nil
}

func (o *CreateStandardOrClassicSubscriptionBody) validateStartDay(formats strfmt.Registry) error {

	if swag.IsZero(o.StartDay) { // not required
		return nil
	}

	if err := validate.MinimumInt("predefinedSubscriptionRequestBean"+"."+"startDay", "body", int64(o.StartDay), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("predefinedSubscriptionRequestBean"+"."+"startDay", "body", int64(o.StartDay), 31, false); err != nil {
		return err
	}

	return nil
}

func (o *CreateStandardOrClassicSubscriptionBody) validateSubscriptionType(formats strfmt.Registry) error {

	if err := validate.Required("predefinedSubscriptionRequestBean"+"."+"subscriptionType", "body", o.SubscriptionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateStandardOrClassicSubscriptionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateStandardOrClassicSubscriptionBody) UnmarshalBinary(b []byte) error {
	var res CreateStandardOrClassicSubscriptionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
