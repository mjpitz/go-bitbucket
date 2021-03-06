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
)

// NewDeleteDeploymentVariableParams creates a new DeleteDeploymentVariableParams object
// with the default values initialized.
func NewDeleteDeploymentVariableParams() *DeleteDeploymentVariableParams {
	var ()
	return &DeleteDeploymentVariableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeploymentVariableParamsWithTimeout creates a new DeleteDeploymentVariableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDeploymentVariableParamsWithTimeout(timeout time.Duration) *DeleteDeploymentVariableParams {
	var ()
	return &DeleteDeploymentVariableParams{

		timeout: timeout,
	}
}

// NewDeleteDeploymentVariableParamsWithContext creates a new DeleteDeploymentVariableParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDeploymentVariableParamsWithContext(ctx context.Context) *DeleteDeploymentVariableParams {
	var ()
	return &DeleteDeploymentVariableParams{

		Context: ctx,
	}
}

// NewDeleteDeploymentVariableParamsWithHTTPClient creates a new DeleteDeploymentVariableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDeploymentVariableParamsWithHTTPClient(client *http.Client) *DeleteDeploymentVariableParams {
	var ()
	return &DeleteDeploymentVariableParams{
		HTTPClient: client,
	}
}

/*DeleteDeploymentVariableParams contains all the parameters to send to the API endpoint
for the delete deployment variable operation typically these are written to a http.Request
*/
type DeleteDeploymentVariableParams struct {

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
	  The UUID of the variable to delete.

	*/
	VariableUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) WithTimeout(timeout time.Duration) *DeleteDeploymentVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) WithContext(ctx context.Context) *DeleteDeploymentVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) WithHTTPClient(client *http.Client) *DeleteDeploymentVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentUUID adds the environmentUUID to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) WithEnvironmentUUID(environmentUUID string) *DeleteDeploymentVariableParams {
	o.SetEnvironmentUUID(environmentUUID)
	return o
}

// SetEnvironmentUUID adds the environmentUuid to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) SetEnvironmentUUID(environmentUUID string) {
	o.EnvironmentUUID = environmentUUID
}

// WithRepoSlug adds the repoSlug to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) WithRepoSlug(repoSlug string) *DeleteDeploymentVariableParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUsername adds the username to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) WithUsername(username string) *DeleteDeploymentVariableParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) SetUsername(username string) {
	o.Username = username
}

// WithVariableUUID adds the variableUUID to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) WithVariableUUID(variableUUID string) *DeleteDeploymentVariableParams {
	o.SetVariableUUID(variableUUID)
	return o
}

// SetVariableUUID adds the variableUuid to the delete deployment variable params
func (o *DeleteDeploymentVariableParams) SetVariableUUID(variableUUID string) {
	o.VariableUUID = variableUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeploymentVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
