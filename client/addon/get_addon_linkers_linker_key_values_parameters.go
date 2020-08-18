// Code generated by go-swagger; DO NOT EDIT.

package addon

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

// NewGetAddonLinkersLinkerKeyValuesParams creates a new GetAddonLinkersLinkerKeyValuesParams object
// with the default values initialized.
func NewGetAddonLinkersLinkerKeyValuesParams() *GetAddonLinkersLinkerKeyValuesParams {
	var ()
	return &GetAddonLinkersLinkerKeyValuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAddonLinkersLinkerKeyValuesParamsWithTimeout creates a new GetAddonLinkersLinkerKeyValuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAddonLinkersLinkerKeyValuesParamsWithTimeout(timeout time.Duration) *GetAddonLinkersLinkerKeyValuesParams {
	var ()
	return &GetAddonLinkersLinkerKeyValuesParams{

		timeout: timeout,
	}
}

// NewGetAddonLinkersLinkerKeyValuesParamsWithContext creates a new GetAddonLinkersLinkerKeyValuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAddonLinkersLinkerKeyValuesParamsWithContext(ctx context.Context) *GetAddonLinkersLinkerKeyValuesParams {
	var ()
	return &GetAddonLinkersLinkerKeyValuesParams{

		Context: ctx,
	}
}

// NewGetAddonLinkersLinkerKeyValuesParamsWithHTTPClient creates a new GetAddonLinkersLinkerKeyValuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAddonLinkersLinkerKeyValuesParamsWithHTTPClient(client *http.Client) *GetAddonLinkersLinkerKeyValuesParams {
	var ()
	return &GetAddonLinkersLinkerKeyValuesParams{
		HTTPClient: client,
	}
}

/*GetAddonLinkersLinkerKeyValuesParams contains all the parameters to send to the API endpoint
for the get addon linkers linker key values operation typically these are written to a http.Request
*/
type GetAddonLinkersLinkerKeyValuesParams struct {

	/*LinkerKey*/
	LinkerKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get addon linkers linker key values params
func (o *GetAddonLinkersLinkerKeyValuesParams) WithTimeout(timeout time.Duration) *GetAddonLinkersLinkerKeyValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get addon linkers linker key values params
func (o *GetAddonLinkersLinkerKeyValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get addon linkers linker key values params
func (o *GetAddonLinkersLinkerKeyValuesParams) WithContext(ctx context.Context) *GetAddonLinkersLinkerKeyValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get addon linkers linker key values params
func (o *GetAddonLinkersLinkerKeyValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get addon linkers linker key values params
func (o *GetAddonLinkersLinkerKeyValuesParams) WithHTTPClient(client *http.Client) *GetAddonLinkersLinkerKeyValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get addon linkers linker key values params
func (o *GetAddonLinkersLinkerKeyValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLinkerKey adds the linkerKey to the get addon linkers linker key values params
func (o *GetAddonLinkersLinkerKeyValuesParams) WithLinkerKey(linkerKey string) *GetAddonLinkersLinkerKeyValuesParams {
	o.SetLinkerKey(linkerKey)
	return o
}

// SetLinkerKey adds the linkerKey to the get addon linkers linker key values params
func (o *GetAddonLinkersLinkerKeyValuesParams) SetLinkerKey(linkerKey string) {
	o.LinkerKey = linkerKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetAddonLinkersLinkerKeyValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param linker_key
	if err := r.SetPathParam("linker_key", o.LinkerKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
