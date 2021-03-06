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

// NewGetSnippetsWorkspaceEncodedIDRevisionPatchParams creates a new GetSnippetsWorkspaceEncodedIDRevisionPatchParams object
// with the default values initialized.
func NewGetSnippetsWorkspaceEncodedIDRevisionPatchParams() *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDRevisionPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnippetsWorkspaceEncodedIDRevisionPatchParamsWithTimeout creates a new GetSnippetsWorkspaceEncodedIDRevisionPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnippetsWorkspaceEncodedIDRevisionPatchParamsWithTimeout(timeout time.Duration) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDRevisionPatchParams{

		timeout: timeout,
	}
}

// NewGetSnippetsWorkspaceEncodedIDRevisionPatchParamsWithContext creates a new GetSnippetsWorkspaceEncodedIDRevisionPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnippetsWorkspaceEncodedIDRevisionPatchParamsWithContext(ctx context.Context) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDRevisionPatchParams{

		Context: ctx,
	}
}

// NewGetSnippetsWorkspaceEncodedIDRevisionPatchParamsWithHTTPClient creates a new GetSnippetsWorkspaceEncodedIDRevisionPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnippetsWorkspaceEncodedIDRevisionPatchParamsWithHTTPClient(client *http.Client) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	var ()
	return &GetSnippetsWorkspaceEncodedIDRevisionPatchParams{
		HTTPClient: client,
	}
}

/*GetSnippetsWorkspaceEncodedIDRevisionPatchParams contains all the parameters to send to the API endpoint
for the get snippets workspace encoded ID revision patch operation typically these are written to a http.Request
*/
type GetSnippetsWorkspaceEncodedIDRevisionPatchParams struct {

	/*EncodedID
	  The snippet id.

	*/
	EncodedID string
	/*Revision
	  A revspec expression. This can simply be a commit SHA1, a ref name, or a compare expression like `staging..production`.

	*/
	Revision string
	/*Spec
	  A commit SHA (e.g. `3a8b42`) or a commit range using double dot
	notation (e.g. `3a8b42..9ff173`).


	*/
	Spec string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) WithTimeout(timeout time.Duration) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) WithContext(ctx context.Context) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) WithHTTPClient(client *http.Client) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEncodedID adds the encodedID to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) WithEncodedID(encodedID string) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	o.SetEncodedID(encodedID)
	return o
}

// SetEncodedID adds the encodedId to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) SetEncodedID(encodedID string) {
	o.EncodedID = encodedID
}

// WithRevision adds the revision to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) WithRevision(revision string) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	o.SetRevision(revision)
	return o
}

// SetRevision adds the revision to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) SetRevision(revision string) {
	o.Revision = revision
}

// WithSpec adds the spec to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) WithSpec(spec string) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	o.SetSpec(spec)
	return o
}

// SetSpec adds the spec to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) SetSpec(spec string) {
	o.Spec = spec
}

// WithWorkspace adds the workspace to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) WithWorkspace(workspace string) *GetSnippetsWorkspaceEncodedIDRevisionPatchParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get snippets workspace encoded ID revision patch params
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnippetsWorkspaceEncodedIDRevisionPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param encoded_id
	if err := r.SetPathParam("encoded_id", o.EncodedID); err != nil {
		return err
	}

	// path param revision
	if err := r.SetPathParam("revision", o.Revision); err != nil {
		return err
	}

	// path param spec
	if err := r.SetPathParam("spec", o.Spec); err != nil {
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
