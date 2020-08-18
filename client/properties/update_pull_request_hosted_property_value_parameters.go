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

// NewUpdatePullRequestHostedPropertyValueParams creates a new UpdatePullRequestHostedPropertyValueParams object
// with the default values initialized.
func NewUpdatePullRequestHostedPropertyValueParams() *UpdatePullRequestHostedPropertyValueParams {
	var ()
	return &UpdatePullRequestHostedPropertyValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePullRequestHostedPropertyValueParamsWithTimeout creates a new UpdatePullRequestHostedPropertyValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePullRequestHostedPropertyValueParamsWithTimeout(timeout time.Duration) *UpdatePullRequestHostedPropertyValueParams {
	var ()
	return &UpdatePullRequestHostedPropertyValueParams{

		timeout: timeout,
	}
}

// NewUpdatePullRequestHostedPropertyValueParamsWithContext creates a new UpdatePullRequestHostedPropertyValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePullRequestHostedPropertyValueParamsWithContext(ctx context.Context) *UpdatePullRequestHostedPropertyValueParams {
	var ()
	return &UpdatePullRequestHostedPropertyValueParams{

		Context: ctx,
	}
}

// NewUpdatePullRequestHostedPropertyValueParamsWithHTTPClient creates a new UpdatePullRequestHostedPropertyValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePullRequestHostedPropertyValueParamsWithHTTPClient(client *http.Client) *UpdatePullRequestHostedPropertyValueParams {
	var ()
	return &UpdatePullRequestHostedPropertyValueParams{
		HTTPClient: client,
	}
}

/*UpdatePullRequestHostedPropertyValueParams contains all the parameters to send to the API endpoint
for the update pull request hosted property value operation typically these are written to a http.Request
*/
type UpdatePullRequestHostedPropertyValueParams struct {

	/*AppKey
	  The key of the Connect app.

	*/
	AppKey string
	/*PropertyName
	  The name of the property.

	*/
	PropertyName string
	/*PullrequestID
	  The pull request ID.

	*/
	PullrequestID string
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

// WithTimeout adds the timeout to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) WithTimeout(timeout time.Duration) *UpdatePullRequestHostedPropertyValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) WithContext(ctx context.Context) *UpdatePullRequestHostedPropertyValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) WithHTTPClient(client *http.Client) *UpdatePullRequestHostedPropertyValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppKey adds the appKey to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) WithAppKey(appKey string) *UpdatePullRequestHostedPropertyValueParams {
	o.SetAppKey(appKey)
	return o
}

// SetAppKey adds the appKey to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) SetAppKey(appKey string) {
	o.AppKey = appKey
}

// WithPropertyName adds the propertyName to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) WithPropertyName(propertyName string) *UpdatePullRequestHostedPropertyValueParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WithPullrequestID adds the pullrequestID to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) WithPullrequestID(pullrequestID string) *UpdatePullRequestHostedPropertyValueParams {
	o.SetPullrequestID(pullrequestID)
	return o
}

// SetPullrequestID adds the pullrequestId to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) SetPullrequestID(pullrequestID string) {
	o.PullrequestID = pullrequestID
}

// WithRepoSlug adds the repoSlug to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) WithRepoSlug(repoSlug string) *UpdatePullRequestHostedPropertyValueParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) WithWorkspace(workspace string) *UpdatePullRequestHostedPropertyValueParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the update pull request hosted property value params
func (o *UpdatePullRequestHostedPropertyValueParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePullRequestHostedPropertyValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_key
	if err := r.SetPathParam("app_key", o.AppKey); err != nil {
		return err
	}

	// path param property_name
	if err := r.SetPathParam("property_name", o.PropertyName); err != nil {
		return err
	}

	// path param pullrequest_id
	if err := r.SetPathParam("pullrequest_id", o.PullrequestID); err != nil {
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