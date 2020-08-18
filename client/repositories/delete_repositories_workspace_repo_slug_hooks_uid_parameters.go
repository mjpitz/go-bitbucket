// Code generated by go-swagger; DO NOT EDIT.

package repositories

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

// NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDParams creates a new DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams object
// with the default values initialized.
func NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDParams() *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDParamsWithTimeout creates a new DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDParamsWithTimeout(timeout time.Duration) *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams{

		timeout: timeout,
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDParamsWithContext creates a new DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDParamsWithContext(ctx context.Context) *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams{

		Context: ctx,
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDParamsWithHTTPClient creates a new DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRepositoriesWorkspaceRepoSlugHooksUIDParamsWithHTTPClient(client *http.Client) *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams{
		HTTPClient: client,
	}
}

/*DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams contains all the parameters to send to the API endpoint
for the delete repositories workspace repo slug hooks UID operation typically these are written to a http.Request
*/
type DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams struct {

	/*RepoSlug
	  This can either be the repository slug or the UUID of the repository,
	surrounded by curly-braces, for example: `{repository UUID}`.


	*/
	RepoSlug string
	/*UID
	  The installed webhook's id

	*/
	UID string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) WithTimeout(timeout time.Duration) *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) WithContext(ctx context.Context) *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) WithHTTPClient(client *http.Client) *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) WithRepoSlug(repoSlug string) *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUID adds the uid to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) WithUID(uid string) *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WithWorkspace adds the workspace to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) WithWorkspace(workspace string) *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the delete repositories workspace repo slug hooks UID params
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepositoriesWorkspaceRepoSlugHooksUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	// path param workspace
	if err := r.SetPathParam("workspace", o.Workspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}