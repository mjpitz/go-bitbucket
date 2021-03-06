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

// NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams creates a new GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams() *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug branch restrictions ID operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams struct {

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

// WithTimeout adds the timeout to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithID(id string) *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetID(id string) {
	o.ID = id
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug branch restrictions ID params
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugBranchRestrictionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
