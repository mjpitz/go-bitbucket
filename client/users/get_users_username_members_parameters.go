// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetUsersUsernameMembersParams creates a new GetUsersUsernameMembersParams object
// with the default values initialized.
func NewGetUsersUsernameMembersParams() *GetUsersUsernameMembersParams {
	var ()
	return &GetUsersUsernameMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersUsernameMembersParamsWithTimeout creates a new GetUsersUsernameMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersUsernameMembersParamsWithTimeout(timeout time.Duration) *GetUsersUsernameMembersParams {
	var ()
	return &GetUsersUsernameMembersParams{

		timeout: timeout,
	}
}

// NewGetUsersUsernameMembersParamsWithContext creates a new GetUsersUsernameMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersUsernameMembersParamsWithContext(ctx context.Context) *GetUsersUsernameMembersParams {
	var ()
	return &GetUsersUsernameMembersParams{

		Context: ctx,
	}
}

// NewGetUsersUsernameMembersParamsWithHTTPClient creates a new GetUsersUsernameMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersUsernameMembersParamsWithHTTPClient(client *http.Client) *GetUsersUsernameMembersParams {
	var ()
	return &GetUsersUsernameMembersParams{
		HTTPClient: client,
	}
}

/*GetUsersUsernameMembersParams contains all the parameters to send to the API endpoint
for the get users username members operation typically these are written to a http.Request
*/
type GetUsersUsernameMembersParams struct {

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

// WithTimeout adds the timeout to the get users username members params
func (o *GetUsersUsernameMembersParams) WithTimeout(timeout time.Duration) *GetUsersUsernameMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users username members params
func (o *GetUsersUsernameMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users username members params
func (o *GetUsersUsernameMembersParams) WithContext(ctx context.Context) *GetUsersUsernameMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users username members params
func (o *GetUsersUsernameMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users username members params
func (o *GetUsersUsernameMembersParams) WithHTTPClient(client *http.Client) *GetUsersUsernameMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users username members params
func (o *GetUsersUsernameMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get users username members params
func (o *GetUsersUsernameMembersParams) WithUsername(username string) *GetUsersUsernameMembersParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get users username members params
func (o *GetUsersUsernameMembersParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersUsernameMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
