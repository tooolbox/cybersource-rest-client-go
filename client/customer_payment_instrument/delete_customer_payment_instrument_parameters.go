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

// NewDeleteCustomerPaymentInstrumentParams creates a new DeleteCustomerPaymentInstrumentParams object
// with the default values initialized.
func NewDeleteCustomerPaymentInstrumentParams() *DeleteCustomerPaymentInstrumentParams {
	var ()
	return &DeleteCustomerPaymentInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCustomerPaymentInstrumentParamsWithTimeout creates a new DeleteCustomerPaymentInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCustomerPaymentInstrumentParamsWithTimeout(timeout time.Duration) *DeleteCustomerPaymentInstrumentParams {
	var ()
	return &DeleteCustomerPaymentInstrumentParams{

		timeout: timeout,
	}
}

// NewDeleteCustomerPaymentInstrumentParamsWithContext creates a new DeleteCustomerPaymentInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCustomerPaymentInstrumentParamsWithContext(ctx context.Context) *DeleteCustomerPaymentInstrumentParams {
	var ()
	return &DeleteCustomerPaymentInstrumentParams{

		Context: ctx,
	}
}

// NewDeleteCustomerPaymentInstrumentParamsWithHTTPClient creates a new DeleteCustomerPaymentInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCustomerPaymentInstrumentParamsWithHTTPClient(client *http.Client) *DeleteCustomerPaymentInstrumentParams {
	var ()
	return &DeleteCustomerPaymentInstrumentParams{
		HTTPClient: client,
	}
}

/*DeleteCustomerPaymentInstrumentParams contains all the parameters to send to the API endpoint
for the delete customer payment instrument operation typically these are written to a http.Request
*/
type DeleteCustomerPaymentInstrumentParams struct {

	/*CustomerTokenID
	  The TokenId of a customer.

	*/
	CustomerTokenID string
	/*PaymentInstrumentTokenID
	  The TokenId of a payment instrument.

	*/
	PaymentInstrumentTokenID string
	/*ProfileID
	  The id of a profile containing user specific TMS configuration.

	*/
	ProfileID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) WithTimeout(timeout time.Duration) *DeleteCustomerPaymentInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) WithContext(ctx context.Context) *DeleteCustomerPaymentInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) WithHTTPClient(client *http.Client) *DeleteCustomerPaymentInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerTokenID adds the customerTokenID to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) WithCustomerTokenID(customerTokenID string) *DeleteCustomerPaymentInstrumentParams {
	o.SetCustomerTokenID(customerTokenID)
	return o
}

// SetCustomerTokenID adds the customerTokenId to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) SetCustomerTokenID(customerTokenID string) {
	o.CustomerTokenID = customerTokenID
}

// WithPaymentInstrumentTokenID adds the paymentInstrumentTokenID to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) WithPaymentInstrumentTokenID(paymentInstrumentTokenID string) *DeleteCustomerPaymentInstrumentParams {
	o.SetPaymentInstrumentTokenID(paymentInstrumentTokenID)
	return o
}

// SetPaymentInstrumentTokenID adds the paymentInstrumentTokenId to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) SetPaymentInstrumentTokenID(paymentInstrumentTokenID string) {
	o.PaymentInstrumentTokenID = paymentInstrumentTokenID
}

// WithProfileID adds the profileID to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) WithProfileID(profileID *string) *DeleteCustomerPaymentInstrumentParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the delete customer payment instrument params
func (o *DeleteCustomerPaymentInstrumentParams) SetProfileID(profileID *string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCustomerPaymentInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerTokenId
	if err := r.SetPathParam("customerTokenId", o.CustomerTokenID); err != nil {
		return err
	}

	// path param paymentInstrumentTokenId
	if err := r.SetPathParam("paymentInstrumentTokenId", o.PaymentInstrumentTokenID); err != nil {
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
