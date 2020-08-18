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

// NewGetUsersUsernameRepositoriesParams creates a new GetUsersUsernameRepositoriesParams object
// with the default values initialized.
func NewGetUsersUsernameRepositoriesParams() *GetUsersUsernameRepositoriesParams {
	var ()
	return &GetUsersUsernameRepositoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersUsernameRepositoriesParamsWithTimeout creates a new GetUsersUsernameRepositoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersUsernameRepositoriesParamsWithTimeout(timeout time.Duration) *GetUsersUsernameRepositoriesParams {
	var ()
	return &GetUsersUsernameRepositoriesParams{

		timeout: timeout,
	}
}

// NewGetUsersUsernameRepositoriesParamsWithContext creates a new GetUsersUsernameRepositoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersUsernameRepositoriesParamsWithContext(ctx context.Context) *GetUsersUsernameRepositoriesParams {
	var ()
	return &GetUsersUsernameRepositoriesParams{

		Context: ctx,
	}
}

// NewGetUsersUsernameRepositoriesParamsWithHTTPClient creates a new GetUsersUsernameRepositoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersUsernameRepositoriesParamsWithHTTPClient(client *http.Client) *GetUsersUsernameRepositoriesParams {
	var ()
	return &GetUsersUsernameRepositoriesParams{
		HTTPClient: client,
	}
}

/*GetUsersUsernameRepositoriesParams contains all the parameters to send to the API endpoint
for the get users username repositories operation typically these are written to a http.Request
*/
type GetUsersUsernameRepositoriesParams struct {

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

// WithTimeout adds the timeout to the get users username repositories params
func (o *GetUsersUsernameRepositoriesParams) WithTimeout(timeout time.Duration) *GetUsersUsernameRepositoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users username repositories params
func (o *GetUsersUsernameRepositoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users username repositories params
func (o *GetUsersUsernameRepositoriesParams) WithContext(ctx context.Context) *GetUsersUsernameRepositoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users username repositories params
func (o *GetUsersUsernameRepositoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users username repositories params
func (o *GetUsersUsernameRepositoriesParams) WithHTTPClient(client *http.Client) *GetUsersUsernameRepositoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users username repositories params
func (o *GetUsersUsernameRepositoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get users username repositories params
func (o *GetUsersUsernameRepositoriesParams) WithUsername(username string) *GetUsersUsernameRepositoriesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get users username repositories params
func (o *GetUsersUsernameRepositoriesParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersUsernameRepositoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
