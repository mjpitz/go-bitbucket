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

// NewGetSnippetsParams creates a new GetSnippetsParams object
// with the default values initialized.
func NewGetSnippetsParams() *GetSnippetsParams {
	var ()
	return &GetSnippetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnippetsParamsWithTimeout creates a new GetSnippetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnippetsParamsWithTimeout(timeout time.Duration) *GetSnippetsParams {
	var ()
	return &GetSnippetsParams{

		timeout: timeout,
	}
}

// NewGetSnippetsParamsWithContext creates a new GetSnippetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnippetsParamsWithContext(ctx context.Context) *GetSnippetsParams {
	var ()
	return &GetSnippetsParams{

		Context: ctx,
	}
}

// NewGetSnippetsParamsWithHTTPClient creates a new GetSnippetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnippetsParamsWithHTTPClient(client *http.Client) *GetSnippetsParams {
	var ()
	return &GetSnippetsParams{
		HTTPClient: client,
	}
}

/*GetSnippetsParams contains all the parameters to send to the API endpoint
for the get snippets operation typically these are written to a http.Request
*/
type GetSnippetsParams struct {

	/*Role
	  Filter down the result based on the authenticated user's role (`owner`, `contributor`, or `member`).

	*/
	Role *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get snippets params
func (o *GetSnippetsParams) WithTimeout(timeout time.Duration) *GetSnippetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snippets params
func (o *GetSnippetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snippets params
func (o *GetSnippetsParams) WithContext(ctx context.Context) *GetSnippetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snippets params
func (o *GetSnippetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snippets params
func (o *GetSnippetsParams) WithHTTPClient(client *http.Client) *GetSnippetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snippets params
func (o *GetSnippetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRole adds the role to the get snippets params
func (o *GetSnippetsParams) WithRole(role *string) *GetSnippetsParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the get snippets params
func (o *GetSnippetsParams) SetRole(role *string) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnippetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
