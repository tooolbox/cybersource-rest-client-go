// Code generated by go-swagger; DO NOT EDIT.

package payment_instrument

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

// NewGetPaymentInstrumentParams creates a new GetPaymentInstrumentParams object
// with the default values initialized.
func NewGetPaymentInstrumentParams() *GetPaymentInstrumentParams {
	var ()
	return &GetPaymentInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentInstrumentParamsWithTimeout creates a new GetPaymentInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentInstrumentParamsWithTimeout(timeout time.Duration) *GetPaymentInstrumentParams {
	var ()
	return &GetPaymentInstrumentParams{

		timeout: timeout,
	}
}

// NewGetPaymentInstrumentParamsWithContext creates a new GetPaymentInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentInstrumentParamsWithContext(ctx context.Context) *GetPaymentInstrumentParams {
	var ()
	return &GetPaymentInstrumentParams{

		Context: ctx,
	}
}

// NewGetPaymentInstrumentParamsWithHTTPClient creates a new GetPaymentInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentInstrumentParamsWithHTTPClient(client *http.Client) *GetPaymentInstrumentParams {
	var ()
	return &GetPaymentInstrumentParams{
		HTTPClient: client,
	}
}

/*GetPaymentInstrumentParams contains all the parameters to send to the API endpoint
for the get payment instrument operation typically these are written to a http.Request
*/
type GetPaymentInstrumentParams struct {

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

// WithTimeout adds the timeout to the get payment instrument params
func (o *GetPaymentInstrumentParams) WithTimeout(timeout time.Duration) *GetPaymentInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment instrument params
func (o *GetPaymentInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment instrument params
func (o *GetPaymentInstrumentParams) WithContext(ctx context.Context) *GetPaymentInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment instrument params
func (o *GetPaymentInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment instrument params
func (o *GetPaymentInstrumentParams) WithHTTPClient(client *http.Client) *GetPaymentInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment instrument params
func (o *GetPaymentInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentInstrumentTokenID adds the paymentInstrumentTokenID to the get payment instrument params
func (o *GetPaymentInstrumentParams) WithPaymentInstrumentTokenID(paymentInstrumentTokenID string) *GetPaymentInstrumentParams {
	o.SetPaymentInstrumentTokenID(paymentInstrumentTokenID)
	return o
}

// SetPaymentInstrumentTokenID adds the paymentInstrumentTokenId to the get payment instrument params
func (o *GetPaymentInstrumentParams) SetPaymentInstrumentTokenID(paymentInstrumentTokenID string) {
	o.PaymentInstrumentTokenID = paymentInstrumentTokenID
}

// WithProfileID adds the profileID to the get payment instrument params
func (o *GetPaymentInstrumentParams) WithProfileID(profileID *string) *GetPaymentInstrumentParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the get payment instrument params
func (o *GetPaymentInstrumentParams) SetProfileID(profileID *string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
