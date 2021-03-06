// Code generated by go-swagger; DO NOT EDIT.

package void

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewVoidPaymentParams creates a new VoidPaymentParams object
// with the default values initialized.
func NewVoidPaymentParams() *VoidPaymentParams {
	var ()
	return &VoidPaymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoidPaymentParamsWithTimeout creates a new VoidPaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoidPaymentParamsWithTimeout(timeout time.Duration) *VoidPaymentParams {
	var ()
	return &VoidPaymentParams{

		timeout: timeout,
	}
}

// NewVoidPaymentParamsWithContext creates a new VoidPaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoidPaymentParamsWithContext(ctx context.Context) *VoidPaymentParams {
	var ()
	return &VoidPaymentParams{

		Context: ctx,
	}
}

// NewVoidPaymentParamsWithHTTPClient creates a new VoidPaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoidPaymentParamsWithHTTPClient(client *http.Client) *VoidPaymentParams {
	var ()
	return &VoidPaymentParams{
		HTTPClient: client,
	}
}

/*VoidPaymentParams contains all the parameters to send to the API endpoint
for the void payment operation typically these are written to a http.Request
*/
type VoidPaymentParams struct {

	/*ID
	  The payment ID returned from a previous payment request.

	*/
	ID string
	/*VoidPaymentRequest*/
	VoidPaymentRequest VoidPaymentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the void payment params
func (o *VoidPaymentParams) WithTimeout(timeout time.Duration) *VoidPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the void payment params
func (o *VoidPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the void payment params
func (o *VoidPaymentParams) WithContext(ctx context.Context) *VoidPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the void payment params
func (o *VoidPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the void payment params
func (o *VoidPaymentParams) WithHTTPClient(client *http.Client) *VoidPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the void payment params
func (o *VoidPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the void payment params
func (o *VoidPaymentParams) WithID(id string) *VoidPaymentParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the void payment params
func (o *VoidPaymentParams) SetID(id string) {
	o.ID = id
}

// WithVoidPaymentRequest adds the voidPaymentRequest to the void payment params
func (o *VoidPaymentParams) WithVoidPaymentRequest(voidPaymentRequest VoidPaymentBody) *VoidPaymentParams {
	o.SetVoidPaymentRequest(voidPaymentRequest)
	return o
}

// SetVoidPaymentRequest adds the voidPaymentRequest to the void payment params
func (o *VoidPaymentParams) SetVoidPaymentRequest(voidPaymentRequest VoidPaymentBody) {
	o.VoidPaymentRequest = voidPaymentRequest
}

// WriteToRequest writes these params to a swagger request
func (o *VoidPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.VoidPaymentRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
