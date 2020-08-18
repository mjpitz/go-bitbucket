// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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

// NewUpdateDeploymentVariableParams creates a new UpdateDeploymentVariableParams object
// with the default values initialized.
func NewUpdateDeploymentVariableParams() *UpdateDeploymentVariableParams {
	var ()
	return &UpdateDeploymentVariableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeploymentVariableParamsWithTimeout creates a new UpdateDeploymentVariableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeploymentVariableParamsWithTimeout(timeout time.Duration) *UpdateDeploymentVariableParams {
	var ()
	return &UpdateDeploymentVariableParams{

		timeout: timeout,
	}
}

// NewUpdateDeploymentVariableParamsWithContext creates a new UpdateDeploymentVariableParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeploymentVariableParamsWithContext(ctx context.Context) *UpdateDeploymentVariableParams {
	var ()
	return &UpdateDeploymentVariableParams{

		Context: ctx,
	}
}

// NewUpdateDeploymentVariableParamsWithHTTPClient creates a new UpdateDeploymentVariableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeploymentVariableParamsWithHTTPClient(client *http.Client) *UpdateDeploymentVariableParams {
	var ()
	return &UpdateDeploymentVariableParams{
		HTTPClient: client,
	}
}

/*UpdateDeploymentVariableParams contains all the parameters to send to the API endpoint
for the update deployment variable operation typically these are written to a http.Request
*/
type UpdateDeploymentVariableParams struct {

	/*Body
	  The updated deployment variable.

	*/
	Body *models.DeploymentVariable
	/*EnvironmentUUID
	  The environment.

	*/
	EnvironmentUUID string
	/*RepoSlug
	  The repository.

	*/
	RepoSlug string
	/*Username
	  The account.

	*/
	Username string
	/*VariableUUID
	  The UUID of the variable to update.

	*/
	VariableUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update deployment variable params
func (o *UpdateDeploymentVariableParams) WithTimeout(timeout time.Duration) *UpdateDeploymentVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update deployment variable params
func (o *UpdateDeploymentVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update deployment variable params
func (o *UpdateDeploymentVariableParams) WithContext(ctx context.Context) *UpdateDeploymentVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update deployment variable params
func (o *UpdateDeploymentVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update deployment variable params
func (o *UpdateDeploymentVariableParams) WithHTTPClient(client *http.Client) *UpdateDeploymentVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update deployment variable params
func (o *UpdateDeploymentVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update deployment variable params
func (o *UpdateDeploymentVariableParams) WithBody(body *models.DeploymentVariable) *UpdateDeploymentVariableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update deployment variable params
func (o *UpdateDeploymentVariableParams) SetBody(body *models.DeploymentVariable) {
	o.Body = body
}

// WithEnvironmentUUID adds the environmentUUID to the update deployment variable params
func (o *UpdateDeploymentVariableParams) WithEnvironmentUUID(environmentUUID string) *UpdateDeploymentVariableParams {
	o.SetEnvironmentUUID(environmentUUID)
	return o
}

// SetEnvironmentUUID adds the environmentUuid to the update deployment variable params
func (o *UpdateDeploymentVariableParams) SetEnvironmentUUID(environmentUUID string) {
	o.EnvironmentUUID = environmentUUID
}

// WithRepoSlug adds the repoSlug to the update deployment variable params
func (o *UpdateDeploymentVariableParams) WithRepoSlug(repoSlug string) *UpdateDeploymentVariableParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the update deployment variable params
func (o *UpdateDeploymentVariableParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUsername adds the username to the update deployment variable params
func (o *UpdateDeploymentVariableParams) WithUsername(username string) *UpdateDeploymentVariableParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the update deployment variable params
func (o *UpdateDeploymentVariableParams) SetUsername(username string) {
	o.Username = username
}

// WithVariableUUID adds the variableUUID to the update deployment variable params
func (o *UpdateDeploymentVariableParams) WithVariableUUID(variableUUID string) *UpdateDeploymentVariableParams {
	o.SetVariableUUID(variableUUID)
	return o
}

// SetVariableUUID adds the variableUuid to the update deployment variable params
func (o *UpdateDeploymentVariableParams) SetVariableUUID(variableUUID string) {
	o.VariableUUID = variableUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeploymentVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param environment_uuid
	if err := r.SetPathParam("environment_uuid", o.EnvironmentUUID); err != nil {
		return err
	}

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	// path param variable_uuid
	if err := r.SetPathParam("variable_uuid", o.VariableUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
