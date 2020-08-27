// Code generated by go-swagger; DO NOT EDIT.

package reversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new reversal API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reversal API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AuthReversal(params *AuthReversalParams) (*AuthReversalCreated, error)

	MitReversal(params *MitReversalParams) (*MitReversalCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AuthReversal processes an authorization reversal

  Include the payment ID in the POST request to reverse the payment amount.
*/
func (a *Client) AuthReversal(params *AuthReversalParams) (*AuthReversalCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthReversalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "authReversal",
		Method:             "POST",
		PathPattern:        "/pts/v2/payments/{id}/reversals",
		ProducesMediaTypes: []string{"application/hal+json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AuthReversalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AuthReversalCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for authReversal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MitReversal timeouts reversal

  This is to reverse a previous payment that merchant does not receive a reply(Mostly due to Timeout). To use this feature/API, make sure to pass unique value to field - clientReferenceInformation -> transactionId in [/pts/v2/payments](https://developer.cybersource.com/api-reference-assets/index.html#payments_payments) API call and use same transactionId in this API request payload to reverse the payment.
*/
func (a *Client) MitReversal(params *MitReversalParams) (*MitReversalCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMitReversalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mitReversal",
		Method:             "POST",
		PathPattern:        "/pts/v2/reversals/",
		ProducesMediaTypes: []string{"application/hal+json;charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MitReversalReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MitReversalCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for mitReversal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
