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

// NewPutTeamsUsernameHooksUIDParams creates a new PutTeamsUsernameHooksUIDParams object
// with the default values initialized.
func NewPutTeamsUsernameHooksUIDParams() *PutTeamsUsernameHooksUIDParams {
	var ()
	return &PutTeamsUsernameHooksUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTeamsUsernameHooksUIDParamsWithTimeout creates a new PutTeamsUsernameHooksUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTeamsUsernameHooksUIDParamsWithTimeout(timeout time.Duration) *PutTeamsUsernameHooksUIDParams {
	var ()
	return &PutTeamsUsernameHooksUIDParams{

		timeout: timeout,
	}
}

// NewPutTeamsUsernameHooksUIDParamsWithContext creates a new PutTeamsUsernameHooksUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTeamsUsernameHooksUIDParamsWithContext(ctx context.Context) *PutTeamsUsernameHooksUIDParams {
	var ()
	return &PutTeamsUsernameHooksUIDParams{

		Context: ctx,
	}
}

// NewPutTeamsUsernameHooksUIDParamsWithHTTPClient creates a new PutTeamsUsernameHooksUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTeamsUsernameHooksUIDParamsWithHTTPClient(client *http.Client) *PutTeamsUsernameHooksUIDParams {
	var ()
	return &PutTeamsUsernameHooksUIDParams{
		HTTPClient: client,
	}
}

/*PutTeamsUsernameHooksUIDParams contains all the parameters to send to the API endpoint
for the put teams username hooks UID operation typically these are written to a http.Request
*/
type PutTeamsUsernameHooksUIDParams struct {

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

// WithTimeout adds the timeout to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) WithTimeout(timeout time.Duration) *PutTeamsUsernameHooksUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) WithContext(ctx context.Context) *PutTeamsUsernameHooksUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) WithHTTPClient(client *http.Client) *PutTeamsUsernameHooksUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) WithUID(uid string) *PutTeamsUsernameHooksUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WithUsername adds the username to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) WithUsername(username string) *PutTeamsUsernameHooksUIDParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the put teams username hooks UID params
func (o *PutTeamsUsernameHooksUIDParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PutTeamsUsernameHooksUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
