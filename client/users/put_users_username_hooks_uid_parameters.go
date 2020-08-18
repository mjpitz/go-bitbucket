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

// NewPutUsersUsernameHooksUIDParams creates a new PutUsersUsernameHooksUIDParams object
// with the default values initialized.
func NewPutUsersUsernameHooksUIDParams() *PutUsersUsernameHooksUIDParams {
	var ()
	return &PutUsersUsernameHooksUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutUsersUsernameHooksUIDParamsWithTimeout creates a new PutUsersUsernameHooksUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutUsersUsernameHooksUIDParamsWithTimeout(timeout time.Duration) *PutUsersUsernameHooksUIDParams {
	var ()
	return &PutUsersUsernameHooksUIDParams{

		timeout: timeout,
	}
}

// NewPutUsersUsernameHooksUIDParamsWithContext creates a new PutUsersUsernameHooksUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutUsersUsernameHooksUIDParamsWithContext(ctx context.Context) *PutUsersUsernameHooksUIDParams {
	var ()
	return &PutUsersUsernameHooksUIDParams{

		Context: ctx,
	}
}

// NewPutUsersUsernameHooksUIDParamsWithHTTPClient creates a new PutUsersUsernameHooksUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutUsersUsernameHooksUIDParamsWithHTTPClient(client *http.Client) *PutUsersUsernameHooksUIDParams {
	var ()
	return &PutUsersUsernameHooksUIDParams{
		HTTPClient: client,
	}
}

/*PutUsersUsernameHooksUIDParams contains all the parameters to send to the API endpoint
for the put users username hooks UID operation typically these are written to a http.Request
*/
type PutUsersUsernameHooksUIDParams struct {

	/*UID
	  The installed webhook's id

	*/
	UID string
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

// WithTimeout adds the timeout to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) WithTimeout(timeout time.Duration) *PutUsersUsernameHooksUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) WithContext(ctx context.Context) *PutUsersUsernameHooksUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) WithHTTPClient(client *http.Client) *PutUsersUsernameHooksUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) WithUID(uid string) *PutUsersUsernameHooksUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WithUsername adds the username to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) WithUsername(username string) *PutUsersUsernameHooksUIDParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the put users username hooks UID params
func (o *PutUsersUsernameHooksUIDParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PutUsersUsernameHooksUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}