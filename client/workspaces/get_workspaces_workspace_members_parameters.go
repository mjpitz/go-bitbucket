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

// NewGetWorkspacesWorkspaceMembersParams creates a new GetWorkspacesWorkspaceMembersParams object
// with the default values initialized.
func NewGetWorkspacesWorkspaceMembersParams() *GetWorkspacesWorkspaceMembersParams {
	var ()
	return &GetWorkspacesWorkspaceMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkspacesWorkspaceMembersParamsWithTimeout creates a new GetWorkspacesWorkspaceMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkspacesWorkspaceMembersParamsWithTimeout(timeout time.Duration) *GetWorkspacesWorkspaceMembersParams {
	var ()
	return &GetWorkspacesWorkspaceMembersParams{

		timeout: timeout,
	}
}

// NewGetWorkspacesWorkspaceMembersParamsWithContext creates a new GetWorkspacesWorkspaceMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkspacesWorkspaceMembersParamsWithContext(ctx context.Context) *GetWorkspacesWorkspaceMembersParams {
	var ()
	return &GetWorkspacesWorkspaceMembersParams{

		Context: ctx,
	}
}

// NewGetWorkspacesWorkspaceMembersParamsWithHTTPClient creates a new GetWorkspacesWorkspaceMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkspacesWorkspaceMembersParamsWithHTTPClient(client *http.Client) *GetWorkspacesWorkspaceMembersParams {
	var ()
	return &GetWorkspacesWorkspaceMembersParams{
		HTTPClient: client,
	}
}

/*GetWorkspacesWorkspaceMembersParams contains all the parameters to send to the API endpoint
for the get workspaces workspace members operation typically these are written to a http.Request
*/
type GetWorkspacesWorkspaceMembersParams struct {

	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workspaces workspace members params
func (o *GetWorkspacesWorkspaceMembersParams) WithTimeout(timeout time.Duration) *GetWorkspacesWorkspaceMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workspaces workspace members params
func (o *GetWorkspacesWorkspaceMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workspaces workspace members params
func (o *GetWorkspacesWorkspaceMembersParams) WithContext(ctx context.Context) *GetWorkspacesWorkspaceMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workspaces workspace members params
func (o *GetWorkspacesWorkspaceMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workspaces workspace members params
func (o *GetWorkspacesWorkspaceMembersParams) WithHTTPClient(client *http.Client) *GetWorkspacesWorkspaceMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workspaces workspace members params
func (o *GetWorkspacesWorkspaceMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspace adds the workspace to the get workspaces workspace members params
func (o *GetWorkspacesWorkspaceMembersParams) WithWorkspace(workspace string) *GetWorkspacesWorkspaceMembersParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get workspaces workspace members params
func (o *GetWorkspacesWorkspaceMembersParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkspacesWorkspaceMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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