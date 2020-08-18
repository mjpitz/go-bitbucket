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

// NewDeleteRepositoryPipelineVariableParams creates a new DeleteRepositoryPipelineVariableParams object
// with the default values initialized.
func NewDeleteRepositoryPipelineVariableParams() *DeleteRepositoryPipelineVariableParams {
	var ()
	return &DeleteRepositoryPipelineVariableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRepositoryPipelineVariableParamsWithTimeout creates a new DeleteRepositoryPipelineVariableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRepositoryPipelineVariableParamsWithTimeout(timeout time.Duration) *DeleteRepositoryPipelineVariableParams {
	var ()
	return &DeleteRepositoryPipelineVariableParams{

		timeout: timeout,
	}
}

// NewDeleteRepositoryPipelineVariableParamsWithContext creates a new DeleteRepositoryPipelineVariableParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRepositoryPipelineVariableParamsWithContext(ctx context.Context) *DeleteRepositoryPipelineVariableParams {
	var ()
	return &DeleteRepositoryPipelineVariableParams{

		Context: ctx,
	}
}

// NewDeleteRepositoryPipelineVariableParamsWithHTTPClient creates a new DeleteRepositoryPipelineVariableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRepositoryPipelineVariableParamsWithHTTPClient(client *http.Client) *DeleteRepositoryPipelineVariableParams {
	var ()
	return &DeleteRepositoryPipelineVariableParams{
		HTTPClient: client,
	}
}

/*DeleteRepositoryPipelineVariableParams contains all the parameters to send to the API endpoint
for the delete repository pipeline variable operation typically these are written to a http.Request
*/
type DeleteRepositoryPipelineVariableParams struct {

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

// WithTimeout adds the timeout to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) WithTimeout(timeout time.Duration) *DeleteRepositoryPipelineVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) WithContext(ctx context.Context) *DeleteRepositoryPipelineVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) WithHTTPClient(client *http.Client) *DeleteRepositoryPipelineVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) WithRepoSlug(repoSlug string) *DeleteRepositoryPipelineVariableParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUsername adds the username to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) WithUsername(username string) *DeleteRepositoryPipelineVariableParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) SetUsername(username string) {
	o.Username = username
}

// WithVariableUUID adds the variableUUID to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) WithVariableUUID(variableUUID string) *DeleteRepositoryPipelineVariableParams {
	o.SetVariableUUID(variableUUID)
	return o
}

// SetVariableUUID adds the variableUuid to the delete repository pipeline variable params
func (o *DeleteRepositoryPipelineVariableParams) SetVariableUUID(variableUUID string) {
	o.VariableUUID = variableUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepositoryPipelineVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
