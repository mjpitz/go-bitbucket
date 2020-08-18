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

	"github.com/mjpitz/go-bitbucket/models"
)

// NewPostRepositoriesWorkspaceRepoSlugParams creates a new PostRepositoriesWorkspaceRepoSlugParams object
// with the default values initialized.
func NewPostRepositoriesWorkspaceRepoSlugParams() *PostRepositoriesWorkspaceRepoSlugParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugParamsWithTimeout creates a new PostRepositoriesWorkspaceRepoSlugParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRepositoriesWorkspaceRepoSlugParamsWithTimeout(timeout time.Duration) *PostRepositoriesWorkspaceRepoSlugParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugParams{

		timeout: timeout,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugParamsWithContext creates a new PostRepositoriesWorkspaceRepoSlugParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRepositoriesWorkspaceRepoSlugParamsWithContext(ctx context.Context) *PostRepositoriesWorkspaceRepoSlugParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugParams{

		Context: ctx,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugParamsWithHTTPClient creates a new PostRepositoriesWorkspaceRepoSlugParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRepositoriesWorkspaceRepoSlugParamsWithHTTPClient(client *http.Client) *PostRepositoriesWorkspaceRepoSlugParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugParams{
		HTTPClient: client,
	}
}

/*PostRepositoriesWorkspaceRepoSlugParams contains all the parameters to send to the API endpoint
for the post repositories workspace repo slug operation typically these are written to a http.Request
*/
type PostRepositoriesWorkspaceRepoSlugParams struct {

	/*Body
	  The repository that is to be created. Note that most object elements are optional. Elements "owner" and "full_name" are ignored as the URL implies them.

	*/
	Body *models.Repository
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

// WithTimeout adds the timeout to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) WithTimeout(timeout time.Duration) *PostRepositoriesWorkspaceRepoSlugParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) WithContext(ctx context.Context) *PostRepositoriesWorkspaceRepoSlugParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) WithHTTPClient(client *http.Client) *PostRepositoriesWorkspaceRepoSlugParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) WithBody(body *models.Repository) *PostRepositoriesWorkspaceRepoSlugParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) SetBody(body *models.Repository) {
	o.Body = body
}

// WithRepoSlug adds the repoSlug to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) WithRepoSlug(repoSlug string) *PostRepositoriesWorkspaceRepoSlugParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) WithWorkspace(workspace string) *PostRepositoriesWorkspaceRepoSlugParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the post repositories workspace repo slug params
func (o *PostRepositoriesWorkspaceRepoSlugParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesWorkspaceRepoSlugParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
