// Code generated by go-swagger; DO NOT EDIT.

package projects

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

	"github.com/mjpitz/go-bitbucket/models"
)

// NewPutTeamsUsernameProjectsProjectKeyParams creates a new PutTeamsUsernameProjectsProjectKeyParams object
// with the default values initialized.
func NewPutTeamsUsernameProjectsProjectKeyParams() *PutTeamsUsernameProjectsProjectKeyParams {
	var ()
	return &PutTeamsUsernameProjectsProjectKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutTeamsUsernameProjectsProjectKeyParamsWithTimeout creates a new PutTeamsUsernameProjectsProjectKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutTeamsUsernameProjectsProjectKeyParamsWithTimeout(timeout time.Duration) *PutTeamsUsernameProjectsProjectKeyParams {
	var ()
	return &PutTeamsUsernameProjectsProjectKeyParams{

		timeout: timeout,
	}
}

// NewPutTeamsUsernameProjectsProjectKeyParamsWithContext creates a new PutTeamsUsernameProjectsProjectKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutTeamsUsernameProjectsProjectKeyParamsWithContext(ctx context.Context) *PutTeamsUsernameProjectsProjectKeyParams {
	var ()
	return &PutTeamsUsernameProjectsProjectKeyParams{

		Context: ctx,
	}
}

// NewPutTeamsUsernameProjectsProjectKeyParamsWithHTTPClient creates a new PutTeamsUsernameProjectsProjectKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutTeamsUsernameProjectsProjectKeyParamsWithHTTPClient(client *http.Client) *PutTeamsUsernameProjectsProjectKeyParams {
	var ()
	return &PutTeamsUsernameProjectsProjectKeyParams{
		HTTPClient: client,
	}
}

/*PutTeamsUsernameProjectsProjectKeyParams contains all the parameters to send to the API endpoint
for the put teams username projects project key operation typically these are written to a http.Request
*/
type PutTeamsUsernameProjectsProjectKeyParams struct {

	/*Body*/
	Body *models.Project
	/*ProjectKey
	  The project in question. This can either be the actual `key` assigned
	to the project or the `UUID` (surrounded by curly-braces (`{}`)).


	*/
	ProjectKey string
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

// WithTimeout adds the timeout to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) WithTimeout(timeout time.Duration) *PutTeamsUsernameProjectsProjectKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) WithContext(ctx context.Context) *PutTeamsUsernameProjectsProjectKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) WithHTTPClient(client *http.Client) *PutTeamsUsernameProjectsProjectKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) WithBody(body *models.Project) *PutTeamsUsernameProjectsProjectKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) SetBody(body *models.Project) {
	o.Body = body
}

// WithProjectKey adds the projectKey to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) WithProjectKey(projectKey string) *PutTeamsUsernameProjectsProjectKeyParams {
	o.SetProjectKey(projectKey)
	return o
}

// SetProjectKey adds the projectKey to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) SetProjectKey(projectKey string) {
	o.ProjectKey = projectKey
}

// WithUsername adds the username to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) WithUsername(username string) *PutTeamsUsernameProjectsProjectKeyParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the put teams username projects project key params
func (o *PutTeamsUsernameProjectsProjectKeyParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PutTeamsUsernameProjectsProjectKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param project_key
	if err := r.SetPathParam("project_key", o.ProjectKey); err != nil {
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
