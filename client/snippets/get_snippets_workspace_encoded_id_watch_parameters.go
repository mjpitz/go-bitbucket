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

// NewGetSnippetsWorkspaceEncodedIDWatchParams creates a new GetSnippetsWorkspaceEncodedIDWatchParams object
// with the default values initialized.
func NewGetSnippetsWorkspaceEncodedIDWatchParams() *GetSnippetsWorkspaceEncodedIDWatchParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDWatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnippetsWorkspaceEncodedIDWatchParamsWithTimeout creates a new GetSnippetsWorkspaceEncodedIDWatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnippetsWorkspaceEncodedIDWatchParamsWithTimeout(timeout time.Duration) *GetSnippetsWorkspaceEncodedIDWatchParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDWatchParams{

		timeout: timeout,
	}
}

// NewGetSnippetsWorkspaceEncodedIDWatchParamsWithContext creates a new GetSnippetsWorkspaceEncodedIDWatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnippetsWorkspaceEncodedIDWatchParamsWithContext(ctx context.Context) *GetSnippetsWorkspaceEncodedIDWatchParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDWatchParams{

		Context: ctx,
	}
}

// NewGetSnippetsWorkspaceEncodedIDWatchParamsWithHTTPClient creates a new GetSnippetsWorkspaceEncodedIDWatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnippetsWorkspaceEncodedIDWatchParamsWithHTTPClient(client *http.Client) *GetSnippetsWorkspaceEncodedIDWatchParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDWatchParams{
		HTTPClient: client,
	}
}

/*GetSnippetsWorkspaceEncodedIDWatchParams contains all the parameters to send to the API endpoint
for the get snippets workspace encoded ID watch operation typically these are written to a http.Request
*/
type GetSnippetsWorkspaceEncodedIDWatchParams struct {

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

// WithTimeout adds the timeout to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) WithTimeout(timeout time.Duration) *GetSnippetsWorkspaceEncodedIDWatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) WithContext(ctx context.Context) *GetSnippetsWorkspaceEncodedIDWatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) WithHTTPClient(client *http.Client) *GetSnippetsWorkspaceEncodedIDWatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEncodedID adds the encodedID to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) WithEncodedID(encodedID string) *GetSnippetsWorkspaceEncodedIDWatchParams {
	o.SetEncodedID(encodedID)
	return o
}

// SetEncodedID adds the encodedId to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) SetEncodedID(encodedID string) {
	o.EncodedID = encodedID
}

// WithWorkspace adds the workspace to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) WithWorkspace(workspace string) *GetSnippetsWorkspaceEncodedIDWatchParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get snippets workspace encoded ID watch params
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnippetsWorkspaceEncodedIDWatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
