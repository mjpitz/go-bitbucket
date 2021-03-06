// Code generated by go-swagger; DO NOT EDIT.

package workspaces

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

// NewGetWorkspacesWorkspaceHooksParams creates a new GetWorkspacesWorkspaceHooksParams object
// with the default values initialized.
func NewGetWorkspacesWorkspaceHooksParams() *GetWorkspacesWorkspaceHooksParams {
	var ()
	return &GetWorkspacesWorkspaceHooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkspacesWorkspaceHooksParamsWithTimeout creates a new GetWorkspacesWorkspaceHooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkspacesWorkspaceHooksParamsWithTimeout(timeout time.Duration) *GetWorkspacesWorkspaceHooksParams {
	var ()
	return &GetWorkspacesWorkspaceHooksParams{

		timeout: timeout,
	}
}

// NewGetWorkspacesWorkspaceHooksParamsWithContext creates a new GetWorkspacesWorkspaceHooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkspacesWorkspaceHooksParamsWithContext(ctx context.Context) *GetWorkspacesWorkspaceHooksParams {
	var ()
	return &GetWorkspacesWorkspaceHooksParams{

		Context: ctx,
	}
}

// NewGetWorkspacesWorkspaceHooksParamsWithHTTPClient creates a new GetWorkspacesWorkspaceHooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkspacesWorkspaceHooksParamsWithHTTPClient(client *http.Client) *GetWorkspacesWorkspaceHooksParams {
	var ()
	return &GetWorkspacesWorkspaceHooksParams{
		HTTPClient: client,
	}
}

/*GetWorkspacesWorkspaceHooksParams contains all the parameters to send to the API endpoint
for the get workspaces workspace hooks operation typically these are written to a http.Request
*/
type GetWorkspacesWorkspaceHooksParams struct {

	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workspaces workspace hooks params
func (o *GetWorkspacesWorkspaceHooksParams) WithTimeout(timeout time.Duration) *GetWorkspacesWorkspaceHooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workspaces workspace hooks params
func (o *GetWorkspacesWorkspaceHooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workspaces workspace hooks params
func (o *GetWorkspacesWorkspaceHooksParams) WithContext(ctx context.Context) *GetWorkspacesWorkspaceHooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workspaces workspace hooks params
func (o *GetWorkspacesWorkspaceHooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workspaces workspace hooks params
func (o *GetWorkspacesWorkspaceHooksParams) WithHTTPClient(client *http.Client) *GetWorkspacesWorkspaceHooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workspaces workspace hooks params
func (o *GetWorkspacesWorkspaceHooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspace adds the workspace to the get workspaces workspace hooks params
func (o *GetWorkspacesWorkspaceHooksParams) WithWorkspace(workspace string) *GetWorkspacesWorkspaceHooksParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get workspaces workspace hooks params
func (o *GetWorkspacesWorkspaceHooksParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkspacesWorkspaceHooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workspace
	if err := r.SetPathParam("workspace", o.Workspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
