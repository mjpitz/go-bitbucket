// Code generated by go-swagger; DO NOT EDIT.

package source

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

// NewGetRepositoriesWorkspaceRepoSlugSrcParams creates a new GetRepositoriesWorkspaceRepoSlugSrcParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugSrcParams() *GetRepositoriesWorkspaceRepoSlugSrcParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugSrcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugSrcParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugSrcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugSrcParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugSrcParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugSrcParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugSrcParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugSrcParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugSrcParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugSrcParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugSrcParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugSrcParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugSrcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugSrcParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugSrcParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugSrcParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugSrcParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug src operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugSrcParams struct {

	/*Format
	  Instead of returning the file's contents, return the (json) meta data for it.

	*/
	Format *string
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

// WithTimeout adds the timeout to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) WithFormat(format *string) *GetRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) SetFormat(format *string) {
	o.Format = format
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug src params
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugSrcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
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
