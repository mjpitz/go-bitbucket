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

// NewGetSnippetsWorkspaceEncodedIDParams creates a new GetSnippetsWorkspaceEncodedIDParams object
// with the default values initialized.
func NewGetSnippetsWorkspaceEncodedIDParams() *GetSnippetsWorkspaceEncodedIDParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnippetsWorkspaceEncodedIDParamsWithTimeout creates a new GetSnippetsWorkspaceEncodedIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnippetsWorkspaceEncodedIDParamsWithTimeout(timeout time.Duration) *GetSnippetsWorkspaceEncodedIDParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDParams{

		timeout: timeout,
	}
}

// NewGetSnippetsWorkspaceEncodedIDParamsWithContext creates a new GetSnippetsWorkspaceEncodedIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnippetsWorkspaceEncodedIDParamsWithContext(ctx context.Context) *GetSnippetsWorkspaceEncodedIDParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDParams{

		Context: ctx,
	}
}

// NewGetSnippetsWorkspaceEncodedIDParamsWithHTTPClient creates a new GetSnippetsWorkspaceEncodedIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnippetsWorkspaceEncodedIDParamsWithHTTPClient(client *http.Client) *GetSnippetsWorkspaceEncodedIDParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDParams{
		HTTPClient: client,
	}
}

/*GetSnippetsWorkspaceEncodedIDParams contains all the parameters to send to the API endpoint
for the get snippets workspace encoded ID operation typically these are written to a http.Request
*/
type GetSnippetsWorkspaceEncodedIDParams struct {

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

// WithTimeout adds the timeout to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) WithTimeout(timeout time.Duration) *GetSnippetsWorkspaceEncodedIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) WithContext(ctx context.Context) *GetSnippetsWorkspaceEncodedIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) WithHTTPClient(client *http.Client) *GetSnippetsWorkspaceEncodedIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEncodedID adds the encodedID to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) WithEncodedID(encodedID string) *GetSnippetsWorkspaceEncodedIDParams {
	o.SetEncodedID(encodedID)
	return o
}

// SetEncodedID adds the encodedId to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) SetEncodedID(encodedID string) {
	o.EncodedID = encodedID
}

// WithWorkspace adds the workspace to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) WithWorkspace(workspace string) *GetSnippetsWorkspaceEncodedIDParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get snippets workspace encoded ID params
func (o *GetSnippetsWorkspaceEncodedIDParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnippetsWorkspaceEncodedIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
