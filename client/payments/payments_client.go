// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreatePayment processes a payment

Authorize the payment for the transaction.

*/
func (a *Client) CreatePayment(params *CreatePaymentParams) (*CreatePaymentCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePaymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPayment",
		Method:             "POST",
		PathPattern:        "/pts/v2/payments/",
		ProducesMediaTypes: []string{"application/json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePaymentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePaymentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
