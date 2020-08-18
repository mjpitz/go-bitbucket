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
)

// NewGetSnippetsWorkspaceEncodedIDWatchersParams creates a new GetSnippetsWorkspaceEncodedIDWatchersParams object
// with the default values initialized.
func NewGetSnippetsWorkspaceEncodedIDWatchersParams() *GetSnippetsWorkspaceEncodedIDWatchersParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDWatchersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnippetsWorkspaceEncodedIDWatchersParamsWithTimeout creates a new GetSnippetsWorkspaceEncodedIDWatchersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnippetsWorkspaceEncodedIDWatchersParamsWithTimeout(timeout time.Duration) *GetSnippetsWorkspaceEncodedIDWatchersParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDWatchersParams{

		timeout: timeout,
	}
}

// NewGetSnippetsWorkspaceEncodedIDWatchersParamsWithContext creates a new GetSnippetsWorkspaceEncodedIDWatchersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnippetsWorkspaceEncodedIDWatchersParamsWithContext(ctx context.Context) *GetSnippetsWorkspaceEncodedIDWatchersParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDWatchersParams{

		Context: ctx,
	}
}

// NewGetSnippetsWorkspaceEncodedIDWatchersParamsWithHTTPClient creates a new GetSnippetsWorkspaceEncodedIDWatchersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnippetsWorkspaceEncodedIDWatchersParamsWithHTTPClient(client *http.Client) *GetSnippetsWorkspaceEncodedIDWatchersParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDWatchersParams{
		HTTPClient: client,
	}
}

/*GetSnippetsWorkspaceEncodedIDWatchersParams contains all the parameters to send to the API endpoint
for the get snippets workspace encoded ID watchers operation typically these are written to a http.Request
*/
type GetSnippetsWorkspaceEncodedIDWatchersParams struct {

	/*EncodedID
	  The snippet id.

	*/
	EncodedID string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) WithTimeout(timeout time.Duration) *GetSnippetsWorkspaceEncodedIDWatchersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) WithContext(ctx context.Context) *GetSnippetsWorkspaceEncodedIDWatchersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) WithHTTPClient(client *http.Client) *GetSnippetsWorkspaceEncodedIDWatchersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEncodedID adds the encodedID to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) WithEncodedID(encodedID string) *GetSnippetsWorkspaceEncodedIDWatchersParams {
	o.SetEncodedID(encodedID)
	return o
}

// SetEncodedID adds the encodedId to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) SetEncodedID(encodedID string) {
	o.EncodedID = encodedID
}

// WithWorkspace adds the workspace to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) WithWorkspace(workspace string) *GetSnippetsWorkspaceEncodedIDWatchersParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get snippets workspace encoded ID watchers params
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnippetsWorkspaceEncodedIDWatchersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param encoded_id
	if err := r.SetPathParam("encoded_id", o.EncodedID); err != nil {
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