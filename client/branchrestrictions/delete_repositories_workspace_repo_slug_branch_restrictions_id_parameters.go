// Code generated by go-swagger; DO NOT EDIT.

package branchrestrictions

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

// NewDeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams creates a new DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams object
// with the default values initialized.
func NewDeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams() *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithTimeout creates a new DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithTimeout(timeout time.Duration) *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams{

		timeout: timeout,
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithContext creates a new DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithContext(ctx context.Context) *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams{

		Context: ctx,
	}
}

// NewDeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithHTTPClient creates a new DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithHTTPClient(client *http.Client) *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	var ()
	return &DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams{
		HTTPClient: client,
	}
}

/*DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams contains all the parameters to send to the API endpoint
for the delete repositories workspace repo slug branch restrictions ID operation typically these are written to a http.Request
*/
type DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams struct {

	/*ID
	  The restriction rule's id

	*/
	ID string
	/*RepoSlug
	  This can either be the repository slug or the UUID of the repository,
	surrounded by curly-braces, for example: `{repository UUID}`.


	*/
	RepoSlug string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithTimeout(timeout time.Duration) *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithContext(ctx context.Context) *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithHTTPClient(client *http.Client) *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithID(id string) *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetID(id string) {
	o.ID = id
}

// WithRepoSlug adds the repoSlug to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithRepoSlug(repoSlug string) *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithWorkspace(workspace string) *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the delete repositories workspace repo slug branch restrictions ID params
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
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