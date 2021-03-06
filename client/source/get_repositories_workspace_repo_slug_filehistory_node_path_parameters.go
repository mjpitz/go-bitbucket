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

// NewGetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams creates a new GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams() *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug filehistory node path operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams struct {

	/*Node*/
	Node string
	/*Path*/
	Path string
	/*Q

	Query string to narrow down the response as per
	[filtering and sorting](../../../../../../meta/filtering).

	*/
	Q *string
	/*Renames

	When `true`, Bitbucket will follow the history of the file across
	renames (this is the default behavior). This can be turned off by
	specifying `false`.

	*/
	Renames *string
	/*RepoSlug
	  This can either be the repository slug or the UUID of the repository,
	surrounded by curly-braces, for example: `{repository UUID}`.


	*/
	RepoSlug string
	/*Sort

	Name of a response property sort the result by as per
	[filtering and sorting](../../../../../../meta/filtering#query-sort).


	*/
	Sort *string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNode adds the node to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithNode(node string) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetNode(node string) {
	o.Node = node
}

// WithPath adds the path to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithPath(path string) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetPath(path string) {
	o.Path = path
}

// WithQ adds the q to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithQ(q *string) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetQ(q *string) {
	o.Q = q
}

// WithRenames adds the renames to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithRenames(renames *string) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetRenames(renames)
	return o
}

// SetRenames adds the renames to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetRenames(renames *string) {
	o.Renames = renames
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithSort adds the sort to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithSort(sort *string) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug filehistory node path params
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugFilehistoryNodePathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param node
	if err := r.SetPathParam("node", o.Node); err != nil {
		return err
	}

	// path param path
	if err := r.SetPathParam("path", o.Path); err != nil {
		return err
	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Renames != nil {

		// query param renames
		var qrRenames string
		if o.Renames != nil {
			qrRenames = *o.Renames
		}
		qRenames := qrRenames
		if qRenames != "" {
			if err := r.SetQueryParam("renames", qRenames); err != nil {
				return err
			}
		}

	}

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

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
