// Code generated by go-swagger; DO NOT EDIT.

package net_fundings

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

// NewGetNetFundingDetailsParams creates a new GetNetFundingDetailsParams object
// with the default values initialized.
func NewGetNetFundingDetailsParams() *GetNetFundingDetailsParams {
	var ()
	return &GetNetFundingDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetFundingDetailsParamsWithTimeout creates a new GetNetFundingDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetFundingDetailsParamsWithTimeout(timeout time.Duration) *GetNetFundingDetailsParams {
	var ()
	return &GetNetFundingDetailsParams{

		timeout: timeout,
	}
}

// NewGetNetFundingDetailsParamsWithContext creates a new GetNetFundingDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetFundingDetailsParamsWithContext(ctx context.Context) *GetNetFundingDetailsParams {
	var ()
	return &GetNetFundingDetailsParams{

		Context: ctx,
	}
}

// NewGetNetFundingDetailsParamsWithHTTPClient creates a new GetNetFundingDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetFundingDetailsParamsWithHTTPClient(client *http.Client) *GetNetFundingDetailsParams {
	var ()
	return &GetNetFundingDetailsParams{
		HTTPClient: client,
	}
}

/*GetNetFundingDetailsParams contains all the parameters to send to the API endpoint
for the get net funding details operation typically these are written to a http.Request
*/
type GetNetFundingDetailsParams struct {

	/*EndTime
	  Valid report End Time in **ISO 8601 format**
	Please refer the following link to know more about ISO 8601 format.[Rfc Date Format](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14)

	**Example date format:**
	  - yyyy-MM-dd'T'HH:mm:ss.SSSZ (e.g. 2018-01-01T00:00:00.000Z)


	*/
	EndTime strfmt.DateTime
	/*GroupName
	  Valid CyberSource Group Name.

	*/
	GroupName *string
	/*OrganizationID
	  Valid Cybersource Organization Id

	*/
	OrganizationID *string
	/*StartTime
	  Valid report Start Time in **ISO 8601 format**
	Please refer the following link to know more about ISO 8601 format.[Rfc Date Format](https://xml2rfc.tools.ietf.org/public/rfc/html/rfc3339.html#anchor14)

	**Example date format:**
	  - yyyy-MM-dd'T'HH:mm:ss.SSSZ (e.g. 2018-01-01T00:00:00.000Z)


	*/
	StartTime strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get net funding details params
func (o *GetNetFundingDetailsParams) WithTimeout(timeout time.Duration) *GetNetFundingDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get net funding details params
func (o *GetNetFundingDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get net funding details params
func (o *GetNetFundingDetailsParams) WithContext(ctx context.Context) *GetNetFundingDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get net funding details params
func (o *GetNetFundingDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get net funding details params
func (o *GetNetFundingDetailsParams) WithHTTPClient(client *http.Client) *GetNetFundingDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get net funding details params
func (o *GetNetFundingDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTime adds the endTime to the get net funding details params
func (o *GetNetFundingDetailsParams) WithEndTime(endTime strfmt.DateTime) *GetNetFundingDetailsParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get net funding details params
func (o *GetNetFundingDetailsParams) SetEndTime(endTime strfmt.DateTime) {
	o.EndTime = endTime
}

// WithGroupName adds the groupName to the get net funding details params
func (o *GetNetFundingDetailsParams) WithGroupName(groupName *string) *GetNetFundingDetailsParams {
	o.SetGroupName(groupName)
	return o
}

// SetGroupName adds the groupName to the get net funding details params
func (o *GetNetFundingDetailsParams) SetGroupName(groupName *string) {
	o.GroupName = groupName
}

// WithOrganizationID adds the organizationID to the get net funding details params
func (o *GetNetFundingDetailsParams) WithOrganizationID(organizationID *string) *GetNetFundingDetailsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get net funding details params
func (o *GetNetFundingDetailsParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithStartTime adds the startTime to the get net funding details params
func (o *GetNetFundingDetailsParams) WithStartTime(startTime strfmt.DateTime) *GetNetFundingDetailsParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get net funding details params
func (o *GetNetFundingDetailsParams) SetStartTime(startTime strfmt.DateTime) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetFundingDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param endTime
	qrEndTime := o.EndTime
	qEndTime := qrEndTime.String()
	if qEndTime != "" {
		if err := r.SetQueryParam("endTime", qEndTime); err != nil {
			return err
		}
	}

	if o.GroupName != nil {

		// query param groupName
		var qrGroupName string
		if o.GroupName != nil {
			qrGroupName = *o.GroupName
		}
		qGroupName := qrGroupName
		if qGroupName != "" {
			if err := r.SetQueryParam("groupName", qGroupName); err != nil {
				return err
			}
		}

	}

	if o.OrganizationID != nil {

		// query param organizationId
		var qrOrganizationID string
		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID
		if qOrganizationID != "" {
			if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
				return err
			}
		}

	}

	// query param startTime
	qrStartTime := o.StartTime
	qStartTime := qrStartTime.String()
	if qStartTime != "" {
		if err := r.SetQueryParam("startTime", qStartTime); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}