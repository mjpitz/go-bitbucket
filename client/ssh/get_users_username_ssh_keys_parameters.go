// Code generated by go-swagger; DO NOT EDIT.

package ssh

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

// NewGetUsersUsernameSSHKeysParams creates a new GetUsersUsernameSSHKeysParams object
// with the default values initialized.
func NewGetUsersUsernameSSHKeysParams() *GetUsersUsernameSSHKeysParams {
	var ()
	return &GetUsersUsernameSSHKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersUsernameSSHKeysParamsWithTimeout creates a new GetUsersUsernameSSHKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersUsernameSSHKeysParamsWithTimeout(timeout time.Duration) *GetUsersUsernameSSHKeysParams {
	var ()
	return &GetUsersUsernameSSHKeysParams{

		timeout: timeout,
	}
}

// NewGetUsersUsernameSSHKeysParamsWithContext creates a new GetUsersUsernameSSHKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersUsernameSSHKeysParamsWithContext(ctx context.Context) *GetUsersUsernameSSHKeysParams {
	var ()
	return &GetUsersUsernameSSHKeysParams{

		Context: ctx,
	}
}

// NewGetUsersUsernameSSHKeysParamsWithHTTPClient creates a new GetUsersUsernameSSHKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersUsernameSSHKeysParamsWithHTTPClient(client *http.Client) *GetUsersUsernameSSHKeysParams {
	var ()
	return &GetUsersUsernameSSHKeysParams{
		HTTPClient: client,
	}
}

/*GetUsersUsernameSSHKeysParams contains all the parameters to send to the API endpoint
for the get users username SSH keys operation typically these are written to a http.Request
*/
type GetUsersUsernameSSHKeysParams struct {

	/*Username
	  The account's UUID, account_id, or username. Note that username has been deprecated due to [privacy changes](https://developer.atlassian.com/cloud/bitbucket/bitbucket-api-changes-gdpr/#removal-of-usernames-from-user-referencing-apis).

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users username SSH keys params
func (o *GetUsersUsernameSSHKeysParams) WithTimeout(timeout time.Duration) *GetUsersUsernameSSHKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users username SSH keys params
func (o *GetUsersUsernameSSHKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users username SSH keys params
func (o *GetUsersUsernameSSHKeysParams) WithContext(ctx context.Context) *GetUsersUsernameSSHKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users username SSH keys params
func (o *GetUsersUsernameSSHKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users username SSH keys params
func (o *GetUsersUsernameSSHKeysParams) WithHTTPClient(client *http.Client) *GetUsersUsernameSSHKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users username SSH keys params
func (o *GetUsersUsernameSSHKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get users username SSH keys params
func (o *GetUsersUsernameSSHKeysParams) WithUsername(username string) *GetUsersUsernameSSHKeysParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get users username SSH keys params
func (o *GetUsersUsernameSSHKeysParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersUsernameSSHKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
