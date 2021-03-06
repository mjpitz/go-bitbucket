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

// NewDeleteSnippetsWorkspaceEncodedIDParams creates a new DeleteSnippetsWorkspaceEncodedIDParams object
// with the default values initialized.
func NewDeleteSnippetsWorkspaceEncodedIDParams() *DeleteSnippetsWorkspaceEncodedIDParams {
	var ()
	return &DeleteSnippetsWorkspaceEncodedIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnippetsWorkspaceEncodedIDParamsWithTimeout creates a new DeleteSnippetsWorkspaceEncodedIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSnippetsWorkspaceEncodedIDParamsWithTimeout(timeout time.Duration) *DeleteSnippetsWorkspaceEncodedIDParams {
	var ()
	return &DeleteSnippetsWorkspaceEncodedIDParams{

		timeout: timeout,
	}
}

// NewDeleteSnippetsWorkspaceEncodedIDParamsWithContext creates a new DeleteSnippetsWorkspaceEncodedIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSnippetsWorkspaceEncodedIDParamsWithContext(ctx context.Context) *DeleteSnippetsWorkspaceEncodedIDParams {
	var ()
	return &DeleteSnippetsWorkspaceEncodedIDParams{

		Context: ctx,
	}
}

// NewDeleteSnippetsWorkspaceEncodedIDParamsWithHTTPClient creates a new DeleteSnippetsWorkspaceEncodedIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSnippetsWorkspaceEncodedIDParamsWithHTTPClient(client *http.Client) *DeleteSnippetsWorkspaceEncodedIDParams {
	var ()
	return &DeleteSnippetsWorkspaceEncodedIDParams{
		HTTPClient: client,
	}
}

/*DeleteSnippetsWorkspaceEncodedIDParams contains all the parameters to send to the API endpoint
for the delete snippets workspace encoded ID operation typically these are written to a http.Request
*/
type DeleteSnippetsWorkspaceEncodedIDParams struct {

	/*EncodedID
	  The snippet's id.

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

// WithTimeout adds the timeout to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) WithTimeout(timeout time.Duration) *DeleteSnippetsWorkspaceEncodedIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) WithContext(ctx context.Context) *DeleteSnippetsWorkspaceEncodedIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) WithHTTPClient(client *http.Client) *DeleteSnippetsWorkspaceEncodedIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEncodedID adds the encodedID to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) WithEncodedID(encodedID string) *DeleteSnippetsWorkspaceEncodedIDParams {
	o.SetEncodedID(encodedID)
	return o
}

// SetEncodedID adds the encodedId to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) SetEncodedID(encodedID string) {
	o.EncodedID = encodedID
}

// WithWorkspace adds the workspace to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) WithWorkspace(workspace string) *DeleteSnippetsWorkspaceEncodedIDParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the delete snippets workspace encoded ID params
func (o *DeleteSnippetsWorkspaceEncodedIDParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnippetsWorkspaceEncodedIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
