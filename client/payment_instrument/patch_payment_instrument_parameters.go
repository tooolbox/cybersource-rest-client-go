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

// NewPatchPaymentInstrumentParams creates a new PatchPaymentInstrumentParams object
// with the default values initialized.
func NewPatchPaymentInstrumentParams() *PatchPaymentInstrumentParams {
	var ()
	return &PatchPaymentInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPaymentInstrumentParamsWithTimeout creates a new PatchPaymentInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPaymentInstrumentParamsWithTimeout(timeout time.Duration) *PatchPaymentInstrumentParams {
	var ()
	return &PatchPaymentInstrumentParams{

		timeout: timeout,
	}
}

// NewPatchPaymentInstrumentParamsWithContext creates a new PatchPaymentInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchPaymentInstrumentParamsWithContext(ctx context.Context) *PatchPaymentInstrumentParams {
	var ()
	return &PatchPaymentInstrumentParams{

		Context: ctx,
	}
}

// NewPatchPaymentInstrumentParamsWithHTTPClient creates a new PatchPaymentInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchPaymentInstrumentParamsWithHTTPClient(client *http.Client) *PatchPaymentInstrumentParams {
	var ()
	return &PatchPaymentInstrumentParams{
		HTTPClient: client,
	}
}

/*PatchPaymentInstrumentParams contains all the parameters to send to the API endpoint
for the patch payment instrument operation typically these are written to a http.Request
*/
type PatchPaymentInstrumentParams struct {

	/*IfMatch
	  Contains an ETag value from a GET request to make the request conditional.

	*/
	IfMatch *string
	/*PatchPaymentInstrumentRequest*/
	PatchPaymentInstrumentRequest PatchPaymentInstrumentBody
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

// WithTimeout adds the timeout to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) WithTimeout(timeout time.Duration) *PatchPaymentInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) WithContext(ctx context.Context) *PatchPaymentInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) WithHTTPClient(client *http.Client) *PatchPaymentInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) WithIfMatch(ifMatch *string) *PatchPaymentInstrumentParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithPatchPaymentInstrumentRequest adds the patchPaymentInstrumentRequest to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) WithPatchPaymentInstrumentRequest(patchPaymentInstrumentRequest PatchPaymentInstrumentBody) *PatchPaymentInstrumentParams {
	o.SetPatchPaymentInstrumentRequest(patchPaymentInstrumentRequest)
	return o
}

// SetPatchPaymentInstrumentRequest adds the patchPaymentInstrumentRequest to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) SetPatchPaymentInstrumentRequest(patchPaymentInstrumentRequest PatchPaymentInstrumentBody) {
	o.PatchPaymentInstrumentRequest = patchPaymentInstrumentRequest
}

// WithPaymentInstrumentTokenID adds the paymentInstrumentTokenID to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) WithPaymentInstrumentTokenID(paymentInstrumentTokenID string) *PatchPaymentInstrumentParams {
	o.SetPaymentInstrumentTokenID(paymentInstrumentTokenID)
	return o
}

// SetPaymentInstrumentTokenID adds the paymentInstrumentTokenId to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) SetPaymentInstrumentTokenID(paymentInstrumentTokenID string) {
	o.PaymentInstrumentTokenID = paymentInstrumentTokenID
}

// WithProfileID adds the profileID to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) WithProfileID(profileID *string) *PatchPaymentInstrumentParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the patch payment instrument params
func (o *PatchPaymentInstrumentParams) SetProfileID(profileID *string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPaymentInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param if-match
		if err := r.SetHeaderParam("if-match", *o.IfMatch); err != nil {
			return err
		}

	}

	if err := r.SetBodyParam(o.PatchPaymentInstrumentRequest); err != nil {
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
