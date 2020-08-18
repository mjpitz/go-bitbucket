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

// NewGetRepositoriesWorkspaceRepoSlugHooksUIDParams creates a new GetRepositoriesWorkspaceRepoSlugHooksUIDParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugHooksUIDParams() *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugHooksUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugHooksUIDParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugHooksUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugHooksUIDParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugHooksUIDParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugHooksUIDParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugHooksUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugHooksUIDParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugHooksUIDParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugHooksUIDParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugHooksUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugHooksUIDParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugHooksUIDParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugHooksUIDParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug hooks UID operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugHooksUIDParams struct {

	/*RepoSlug
	  This can either be the repository slug or the UUID of the repository,
	surrounded by curly-braces, for example: `{repository UUID}`.


	*/
	RepoSlug string
	/*UID
	  The installed webhook's id.

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

// WithTimeout adds the timeout to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUID adds the uid to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) WithUID(uid string) *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) SetUID(uid string) {
	o.UID = uid
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugHooksUIDParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug hooks UID params
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugHooksUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
