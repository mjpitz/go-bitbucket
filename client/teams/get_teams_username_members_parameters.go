// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewGetTeamsUsernameMembersParams creates a new GetTeamsUsernameMembersParams object
// with the default values initialized.
func NewGetTeamsUsernameMembersParams() *GetTeamsUsernameMembersParams {
	var ()
	return &GetTeamsUsernameMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamsUsernameMembersParamsWithTimeout creates a new GetTeamsUsernameMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTeamsUsernameMembersParamsWithTimeout(timeout time.Duration) *GetTeamsUsernameMembersParams {
	var ()
	return &GetTeamsUsernameMembersParams{

		timeout: timeout,
	}
}

// NewGetTeamsUsernameMembersParamsWithContext creates a new GetTeamsUsernameMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTeamsUsernameMembersParamsWithContext(ctx context.Context) *GetTeamsUsernameMembersParams {
	var ()
	return &GetTeamsUsernameMembersParams{

		Context: ctx,
	}
}

// NewGetTeamsUsernameMembersParamsWithHTTPClient creates a new GetTeamsUsernameMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTeamsUsernameMembersParamsWithHTTPClient(client *http.Client) *GetTeamsUsernameMembersParams {
	var ()
	return &GetTeamsUsernameMembersParams{
		HTTPClient: client,
	}
}

/*GetTeamsUsernameMembersParams contains all the parameters to send to the API endpoint
for the get teams username members operation typically these are written to a http.Request
*/
type GetTeamsUsernameMembersParams struct {

	/*Username
	  This can either be the username or the UUID of the account,
	surrounded by curly-braces, for example: `{account UUID}`. An account
	is either a team or user.


	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get teams username members params
func (o *GetTeamsUsernameMembersParams) WithTimeout(timeout time.Duration) *GetTeamsUsernameMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get teams username members params
func (o *GetTeamsUsernameMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get teams username members params
func (o *GetTeamsUsernameMembersParams) WithContext(ctx context.Context) *GetTeamsUsernameMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get teams username members params
func (o *GetTeamsUsernameMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get teams username members params
func (o *GetTeamsUsernameMembersParams) WithHTTPClient(client *http.Client) *GetTeamsUsernameMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get teams username members params
func (o *GetTeamsUsernameMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get teams username members params
func (o *GetTeamsUsernameMembersParams) WithUsername(username string) *GetTeamsUsernameMembersParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get teams username members params
func (o *GetTeamsUsernameMembersParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamsUsernameMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
