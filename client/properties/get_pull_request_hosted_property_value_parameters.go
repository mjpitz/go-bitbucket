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

// NewGetPullRequestHostedPropertyValueParams creates a new GetPullRequestHostedPropertyValueParams object
// with the default values initialized.
func NewGetPullRequestHostedPropertyValueParams() *GetPullRequestHostedPropertyValueParams {
	var ()
	return &GetPullRequestHostedPropertyValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPullRequestHostedPropertyValueParamsWithTimeout creates a new GetPullRequestHostedPropertyValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPullRequestHostedPropertyValueParamsWithTimeout(timeout time.Duration) *GetPullRequestHostedPropertyValueParams {
	var ()
	return &GetPullRequestHostedPropertyValueParams{

		timeout: timeout,
	}
}

// NewGetPullRequestHostedPropertyValueParamsWithContext creates a new GetPullRequestHostedPropertyValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPullRequestHostedPropertyValueParamsWithContext(ctx context.Context) *GetPullRequestHostedPropertyValueParams {
	var ()
	return &GetPullRequestHostedPropertyValueParams{

		Context: ctx,
	}
}

// NewGetPullRequestHostedPropertyValueParamsWithHTTPClient creates a new GetPullRequestHostedPropertyValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPullRequestHostedPropertyValueParamsWithHTTPClient(client *http.Client) *GetPullRequestHostedPropertyValueParams {
	var ()
	return &GetPullRequestHostedPropertyValueParams{
		HTTPClient: client,
	}
}

/*GetPullRequestHostedPropertyValueParams contains all the parameters to send to the API endpoint
for the get pull request hosted property value operation typically these are written to a http.Request
*/
type GetPullRequestHostedPropertyValueParams struct {

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

// WithTimeout adds the timeout to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) WithTimeout(timeout time.Duration) *GetPullRequestHostedPropertyValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) WithContext(ctx context.Context) *GetPullRequestHostedPropertyValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) WithHTTPClient(client *http.Client) *GetPullRequestHostedPropertyValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppKey adds the appKey to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) WithAppKey(appKey string) *GetPullRequestHostedPropertyValueParams {
	o.SetAppKey(appKey)
	return o
}

// SetAppKey adds the appKey to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) SetAppKey(appKey string) {
	o.AppKey = appKey
}

// WithPropertyName adds the propertyName to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) WithPropertyName(propertyName string) *GetPullRequestHostedPropertyValueParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WithPullrequestID adds the pullrequestID to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) WithPullrequestID(pullrequestID string) *GetPullRequestHostedPropertyValueParams {
	o.SetPullrequestID(pullrequestID)
	return o
}

// SetPullrequestID adds the pullrequestId to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) SetPullrequestID(pullrequestID string) {
	o.PullrequestID = pullrequestID
}

// WithRepoSlug adds the repoSlug to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) WithRepoSlug(repoSlug string) *GetPullRequestHostedPropertyValueParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) WithWorkspace(workspace string) *GetPullRequestHostedPropertyValueParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get pull request hosted property value params
func (o *GetPullRequestHostedPropertyValueParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetPullRequestHostedPropertyValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
