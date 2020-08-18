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

// NewGetTeamsUsernameHooksParams creates a new GetTeamsUsernameHooksParams object
// with the default values initialized.
func NewGetTeamsUsernameHooksParams() *GetTeamsUsernameHooksParams {
	var ()
	return &GetTeamsUsernameHooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamsUsernameHooksParamsWithTimeout creates a new GetTeamsUsernameHooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTeamsUsernameHooksParamsWithTimeout(timeout time.Duration) *GetTeamsUsernameHooksParams {
	var ()
	return &GetTeamsUsernameHooksParams{

		timeout: timeout,
	}
}

// NewGetTeamsUsernameHooksParamsWithContext creates a new GetTeamsUsernameHooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTeamsUsernameHooksParamsWithContext(ctx context.Context) *GetTeamsUsernameHooksParams {
	var ()
	return &GetTeamsUsernameHooksParams{

		Context: ctx,
	}
}

// NewGetTeamsUsernameHooksParamsWithHTTPClient creates a new GetTeamsUsernameHooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTeamsUsernameHooksParamsWithHTTPClient(client *http.Client) *GetTeamsUsernameHooksParams {
	var ()
	return &GetTeamsUsernameHooksParams{
		HTTPClient: client,
	}
}

/*GetTeamsUsernameHooksParams contains all the parameters to send to the API endpoint
for the get teams username hooks operation typically these are written to a http.Request
*/
type GetTeamsUsernameHooksParams struct {

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

// WithTimeout adds the timeout to the get teams username hooks params
func (o *GetTeamsUsernameHooksParams) WithTimeout(timeout time.Duration) *GetTeamsUsernameHooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get teams username hooks params
func (o *GetTeamsUsernameHooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get teams username hooks params
func (o *GetTeamsUsernameHooksParams) WithContext(ctx context.Context) *GetTeamsUsernameHooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get teams username hooks params
func (o *GetTeamsUsernameHooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get teams username hooks params
func (o *GetTeamsUsernameHooksParams) WithHTTPClient(client *http.Client) *GetTeamsUsernameHooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get teams username hooks params
func (o *GetTeamsUsernameHooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get teams username hooks params
func (o *GetTeamsUsernameHooksParams) WithUsername(username string) *GetTeamsUsernameHooksParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get teams username hooks params
func (o *GetTeamsUsernameHooksParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamsUsernameHooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
