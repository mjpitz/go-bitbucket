// Code generated by go-swagger; DO NOT EDIT.

package snippets

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

// NewPostSnippetsWorkspaceParams creates a new PostSnippetsWorkspaceParams object
// with the default values initialized.
func NewPostSnippetsWorkspaceParams() *PostSnippetsWorkspaceParams {
	var ()
	return &PostSnippetsWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSnippetsWorkspaceParamsWithTimeout creates a new PostSnippetsWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSnippetsWorkspaceParamsWithTimeout(timeout time.Duration) *PostSnippetsWorkspaceParams {
	var ()
	return &PostSnippetsWorkspaceParams{

		timeout: timeout,
	}
}

// NewPostSnippetsWorkspaceParamsWithContext creates a new PostSnippetsWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSnippetsWorkspaceParamsWithContext(ctx context.Context) *PostSnippetsWorkspaceParams {
	var ()
	return &PostSnippetsWorkspaceParams{

		Context: ctx,
	}
}

// NewPostSnippetsWorkspaceParamsWithHTTPClient creates a new PostSnippetsWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSnippetsWorkspaceParamsWithHTTPClient(client *http.Client) *PostSnippetsWorkspaceParams {
	var ()
	return &PostSnippetsWorkspaceParams{
		HTTPClient: client,
	}
}

/*PostSnippetsWorkspaceParams contains all the parameters to send to the API endpoint
for the post snippets workspace operation typically these are written to a http.Request
*/
type PostSnippetsWorkspaceParams struct {

	/*Body
	  The new snippet object.

	*/
	Body *models.Snippet
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) WithTimeout(timeout time.Duration) *PostSnippetsWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) WithContext(ctx context.Context) *PostSnippetsWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) WithHTTPClient(client *http.Client) *PostSnippetsWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) WithBody(body *models.Snippet) *PostSnippetsWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) SetBody(body *models.Snippet) {
	o.Body = body
}

// WithWorkspace adds the workspace to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) WithWorkspace(workspace string) *PostSnippetsWorkspaceParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the post snippets workspace params
func (o *PostSnippetsWorkspaceParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *PostSnippetsWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
