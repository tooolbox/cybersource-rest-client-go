// Code generated by go-swagger; DO NOT EDIT.

package customer_payment_instrument

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostCustomerPaymentInstrumentParams creates a new PostCustomerPaymentInstrumentParams object
// with the default values initialized.
func NewPostCustomerPaymentInstrumentParams() *PostCustomerPaymentInstrumentParams {
	var ()
	return &PostCustomerPaymentInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCustomerPaymentInstrumentParamsWithTimeout creates a new PostCustomerPaymentInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCustomerPaymentInstrumentParamsWithTimeout(timeout time.Duration) *PostCustomerPaymentInstrumentParams {
	var ()
	return &PostCustomerPaymentInstrumentParams{

		timeout: timeout,
	}
}

// NewPostCustomerPaymentInstrumentParamsWithContext creates a new PostCustomerPaymentInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCustomerPaymentInstrumentParamsWithContext(ctx context.Context) *PostCustomerPaymentInstrumentParams {
	var ()
	return &PostCustomerPaymentInstrumentParams{

		Context: ctx,
	}
}

// NewPostCustomerPaymentInstrumentParamsWithHTTPClient creates a new PostCustomerPaymentInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCustomerPaymentInstrumentParamsWithHTTPClient(client *http.Client) *PostCustomerPaymentInstrumentParams {
	var ()
	return &PostCustomerPaymentInstrumentParams{
		HTTPClient: client,
	}
}

/*PostCustomerPaymentInstrumentParams contains all the parameters to send to the API endpoint
for the post customer payment instrument operation typically these are written to a http.Request
*/
type PostCustomerPaymentInstrumentParams struct {

	/*CustomerTokenID
	  The TokenId of a customer.

	*/
	CustomerTokenID string
	/*PostCustomerPaymentInstrumentRequest*/
	PostCustomerPaymentInstrumentRequest PostCustomerPaymentInstrumentBody
	/*ProfileID
	  The id of a profile containing user specific TMS configuration.

	*/
	ProfileID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) WithTimeout(timeout time.Duration) *PostCustomerPaymentInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) WithContext(ctx context.Context) *PostCustomerPaymentInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) WithHTTPClient(client *http.Client) *PostCustomerPaymentInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerTokenID adds the customerTokenID to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) WithCustomerTokenID(customerTokenID string) *PostCustomerPaymentInstrumentParams {
	o.SetCustomerTokenID(customerTokenID)
	return o
}

// SetCustomerTokenID adds the customerTokenId to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) SetCustomerTokenID(customerTokenID string) {
	o.CustomerTokenID = customerTokenID
}

// WithPostCustomerPaymentInstrumentRequest adds the postCustomerPaymentInstrumentRequest to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) WithPostCustomerPaymentInstrumentRequest(postCustomerPaymentInstrumentRequest PostCustomerPaymentInstrumentBody) *PostCustomerPaymentInstrumentParams {
	o.SetPostCustomerPaymentInstrumentRequest(postCustomerPaymentInstrumentRequest)
	return o
}

// SetPostCustomerPaymentInstrumentRequest adds the postCustomerPaymentInstrumentRequest to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) SetPostCustomerPaymentInstrumentRequest(postCustomerPaymentInstrumentRequest PostCustomerPaymentInstrumentBody) {
	o.PostCustomerPaymentInstrumentRequest = postCustomerPaymentInstrumentRequest
}

// WithProfileID adds the profileID to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) WithProfileID(profileID *string) *PostCustomerPaymentInstrumentParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the post customer payment instrument params
func (o *PostCustomerPaymentInstrumentParams) SetProfileID(profileID *string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *PostCustomerPaymentInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerTokenId
	if err := r.SetPathParam("customerTokenId", o.CustomerTokenID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.PostCustomerPaymentInstrumentRequest); err != nil {
		return err
	}

	if o.ProfileID != nil {

		// header param profile-id
		if err := r.SetHeaderParam("profile-id", *o.ProfileID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
