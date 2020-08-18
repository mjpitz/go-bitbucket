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

// NewGetUserPermissionsWorkspacesParams creates a new GetUserPermissionsWorkspacesParams object
// with the default values initialized.
func NewGetUserPermissionsWorkspacesParams() *GetUserPermissionsWorkspacesParams {

	return &GetUserPermissionsWorkspacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserPermissionsWorkspacesParamsWithTimeout creates a new GetUserPermissionsWorkspacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserPermissionsWorkspacesParamsWithTimeout(timeout time.Duration) *GetUserPermissionsWorkspacesParams {

	return &GetUserPermissionsWorkspacesParams{

		timeout: timeout,
	}
}

// NewGetUserPermissionsWorkspacesParamsWithContext creates a new GetUserPermissionsWorkspacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserPermissionsWorkspacesParamsWithContext(ctx context.Context) *GetUserPermissionsWorkspacesParams {

	return &GetUserPermissionsWorkspacesParams{

		Context: ctx,
	}
}

// NewGetUserPermissionsWorkspacesParamsWithHTTPClient creates a new GetUserPermissionsWorkspacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserPermissionsWorkspacesParamsWithHTTPClient(client *http.Client) *GetUserPermissionsWorkspacesParams {

	return &GetUserPermissionsWorkspacesParams{
		HTTPClient: client,
	}
}

/*GetUserPermissionsWorkspacesParams contains all the parameters to send to the API endpoint
for the get user permissions workspaces operation typically these are written to a http.Request
*/
type GetUserPermissionsWorkspacesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user permissions workspaces params
func (o *GetUserPermissionsWorkspacesParams) WithTimeout(timeout time.Duration) *GetUserPermissionsWorkspacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user permissions workspaces params
func (o *GetUserPermissionsWorkspacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user permissions workspaces params
func (o *GetUserPermissionsWorkspacesParams) WithContext(ctx context.Context) *GetUserPermissionsWorkspacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user permissions workspaces params
func (o *GetUserPermissionsWorkspacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user permissions workspaces params
func (o *GetUserPermissionsWorkspacesParams) WithHTTPClient(client *http.Client) *GetUserPermissionsWorkspacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user permissions workspaces params
func (o *GetUserPermissionsWorkspacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserPermissionsWorkspacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
