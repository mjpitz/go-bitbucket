// Code generated by go-swagger; DO NOT EDIT.

package properties

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

// NewDeleteCommitHostedPropertyValueParams creates a new DeleteCommitHostedPropertyValueParams object
// with the default values initialized.
func NewDeleteCommitHostedPropertyValueParams() *DeleteCommitHostedPropertyValueParams {
	var ()
	return &DeleteCommitHostedPropertyValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCommitHostedPropertyValueParamsWithTimeout creates a new DeleteCommitHostedPropertyValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCommitHostedPropertyValueParamsWithTimeout(timeout time.Duration) *DeleteCommitHostedPropertyValueParams {
	var ()
	return &DeleteCommitHostedPropertyValueParams{

		timeout: timeout,
	}
}

// NewDeleteCommitHostedPropertyValueParamsWithContext creates a new DeleteCommitHostedPropertyValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCommitHostedPropertyValueParamsWithContext(ctx context.Context) *DeleteCommitHostedPropertyValueParams {
	var ()
	return &DeleteCommitHostedPropertyValueParams{

		Context: ctx,
	}
}

// NewDeleteCommitHostedPropertyValueParamsWithHTTPClient creates a new DeleteCommitHostedPropertyValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCommitHostedPropertyValueParamsWithHTTPClient(client *http.Client) *DeleteCommitHostedPropertyValueParams {
	var ()
	return &DeleteCommitHostedPropertyValueParams{
		HTTPClient: client,
	}
}

/*DeleteCommitHostedPropertyValueParams contains all the parameters to send to the API endpoint
for the delete commit hosted property value operation typically these are written to a http.Request
*/
type DeleteCommitHostedPropertyValueParams struct {

	/*AppKey
	  The key of the Connect app.

	*/
	AppKey string
	/*Commit
	  The commit.

	*/
	Commit string
	/*PropertyName
	  The name of the property.

	*/
	PropertyName string
	/*RepoSlug
	  The repository.

	*/
	RepoSlug string
	/*Workspace
	  The repository container; either the workspace slug or the UUID in curly braces.

	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) WithTimeout(timeout time.Duration) *DeleteCommitHostedPropertyValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) WithContext(ctx context.Context) *DeleteCommitHostedPropertyValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) WithHTTPClient(client *http.Client) *DeleteCommitHostedPropertyValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppKey adds the appKey to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) WithAppKey(appKey string) *DeleteCommitHostedPropertyValueParams {
	o.SetAppKey(appKey)
	return o
}

// SetAppKey adds the appKey to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) SetAppKey(appKey string) {
	o.AppKey = appKey
}

// WithCommit adds the commit to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) WithCommit(commit string) *DeleteCommitHostedPropertyValueParams {
	o.SetCommit(commit)
	return o
}

// SetCommit adds the commit to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) SetCommit(commit string) {
	o.Commit = commit
}

// WithPropertyName adds the propertyName to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) WithPropertyName(propertyName string) *DeleteCommitHostedPropertyValueParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WithRepoSlug adds the repoSlug to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) WithRepoSlug(repoSlug string) *DeleteCommitHostedPropertyValueParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) WithWorkspace(workspace string) *DeleteCommitHostedPropertyValueParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the delete commit hosted property value params
func (o *DeleteCommitHostedPropertyValueParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCommitHostedPropertyValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_key
	if err := r.SetPathParam("app_key", o.AppKey); err != nil {
		return err
	}

	// path param commit
	if err := r.SetPathParam("commit", o.Commit); err != nil {
		return err
	}

	// path param property_name
	if err := r.SetPathParam("property_name", o.PropertyName); err != nil {
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
