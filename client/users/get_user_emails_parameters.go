// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetUserEmailsParams creates a new GetUserEmailsParams object
// with the default values initialized.
func NewGetUserEmailsParams() *GetUserEmailsParams {

	return &GetUserEmailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserEmailsParamsWithTimeout creates a new GetUserEmailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserEmailsParamsWithTimeout(timeout time.Duration) *GetUserEmailsParams {

	return &GetUserEmailsParams{

		timeout: timeout,
	}
}

// NewGetUserEmailsParamsWithContext creates a new GetUserEmailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserEmailsParamsWithContext(ctx context.Context) *GetUserEmailsParams {

	return &GetUserEmailsParams{

		Context: ctx,
	}
}

// NewGetUserEmailsParamsWithHTTPClient creates a new GetUserEmailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserEmailsParamsWithHTTPClient(client *http.Client) *GetUserEmailsParams {

	return &GetUserEmailsParams{
		HTTPClient: client,
	}
}

/*GetUserEmailsParams contains all the parameters to send to the API endpoint
for the get user emails operation typically these are written to a http.Request
*/
type GetUserEmailsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user emails params
func (o *GetUserEmailsParams) WithTimeout(timeout time.Duration) *GetUserEmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user emails params
func (o *GetUserEmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user emails params
func (o *GetUserEmailsParams) WithContext(ctx context.Context) *GetUserEmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user emails params
func (o *GetUserEmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user emails params
func (o *GetUserEmailsParams) WithHTTPClient(client *http.Client) *GetUserEmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user emails params
func (o *GetUserEmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserEmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
