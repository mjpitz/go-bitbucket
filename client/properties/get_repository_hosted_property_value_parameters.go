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

// NewGetRepositoryHostedPropertyValueParams creates a new GetRepositoryHostedPropertyValueParams object
// with the default values initialized.
func NewGetRepositoryHostedPropertyValueParams() *GetRepositoryHostedPropertyValueParams {
	var ()
	return &GetRepositoryHostedPropertyValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoryHostedPropertyValueParamsWithTimeout creates a new GetRepositoryHostedPropertyValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoryHostedPropertyValueParamsWithTimeout(timeout time.Duration) *GetRepositoryHostedPropertyValueParams {
	var ()
	return &GetRepositoryHostedPropertyValueParams{

		timeout: timeout,
	}
}

// NewGetRepositoryHostedPropertyValueParamsWithContext creates a new GetRepositoryHostedPropertyValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoryHostedPropertyValueParamsWithContext(ctx context.Context) *GetRepositoryHostedPropertyValueParams {
	var ()
	return &GetRepositoryHostedPropertyValueParams{

		Context: ctx,
	}
}

// NewGetRepositoryHostedPropertyValueParamsWithHTTPClient creates a new GetRepositoryHostedPropertyValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoryHostedPropertyValueParamsWithHTTPClient(client *http.Client) *GetRepositoryHostedPropertyValueParams {
	var ()
	return &GetRepositoryHostedPropertyValueParams{
		HTTPClient: client,
	}
}

/*GetRepositoryHostedPropertyValueParams contains all the parameters to send to the API endpoint
for the get repository hosted property value operation typically these are written to a http.Request
*/
type GetRepositoryHostedPropertyValueParams struct {

	/*AppKey
	  The key of the Connect app.

	*/
	AppKey string
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

// WithTimeout adds the timeout to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) WithTimeout(timeout time.Duration) *GetRepositoryHostedPropertyValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) WithContext(ctx context.Context) *GetRepositoryHostedPropertyValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) WithHTTPClient(client *http.Client) *GetRepositoryHostedPropertyValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppKey adds the appKey to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) WithAppKey(appKey string) *GetRepositoryHostedPropertyValueParams {
	o.SetAppKey(appKey)
	return o
}

// SetAppKey adds the appKey to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) SetAppKey(appKey string) {
	o.AppKey = appKey
}

// WithPropertyName adds the propertyName to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) WithPropertyName(propertyName string) *GetRepositoryHostedPropertyValueParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WithRepoSlug adds the repoSlug to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) WithRepoSlug(repoSlug string) *GetRepositoryHostedPropertyValueParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) WithWorkspace(workspace string) *GetRepositoryHostedPropertyValueParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repository hosted property value params
func (o *GetRepositoryHostedPropertyValueParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoryHostedPropertyValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
