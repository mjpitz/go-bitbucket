// Code generated by go-swagger; DO NOT EDIT.

package properties

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

// NewRetrieveUserHostedPropertyValueParams creates a new RetrieveUserHostedPropertyValueParams object
// with the default values initialized.
func NewRetrieveUserHostedPropertyValueParams() *RetrieveUserHostedPropertyValueParams {
	var ()
	return &RetrieveUserHostedPropertyValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveUserHostedPropertyValueParamsWithTimeout creates a new RetrieveUserHostedPropertyValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveUserHostedPropertyValueParamsWithTimeout(timeout time.Duration) *RetrieveUserHostedPropertyValueParams {
	var ()
	return &RetrieveUserHostedPropertyValueParams{

		timeout: timeout,
	}
}

// NewRetrieveUserHostedPropertyValueParamsWithContext creates a new RetrieveUserHostedPropertyValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveUserHostedPropertyValueParamsWithContext(ctx context.Context) *RetrieveUserHostedPropertyValueParams {
	var ()
	return &RetrieveUserHostedPropertyValueParams{

		Context: ctx,
	}
}

// NewRetrieveUserHostedPropertyValueParamsWithHTTPClient creates a new RetrieveUserHostedPropertyValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveUserHostedPropertyValueParamsWithHTTPClient(client *http.Client) *RetrieveUserHostedPropertyValueParams {
	var ()
	return &RetrieveUserHostedPropertyValueParams{
		HTTPClient: client,
	}
}

/*RetrieveUserHostedPropertyValueParams contains all the parameters to send to the API endpoint
for the retrieve user hosted property value operation typically these are written to a http.Request
*/
type RetrieveUserHostedPropertyValueParams struct {

	/*AppKey
	  The key of the Connect app.

	*/
	AppKey string
	/*PropertyName
	  The name of the property.

	*/
	PropertyName string
	/*SelectedUser
	  Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.

	*/
	SelectedUser string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) WithTimeout(timeout time.Duration) *RetrieveUserHostedPropertyValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) WithContext(ctx context.Context) *RetrieveUserHostedPropertyValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) WithHTTPClient(client *http.Client) *RetrieveUserHostedPropertyValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppKey adds the appKey to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) WithAppKey(appKey string) *RetrieveUserHostedPropertyValueParams {
	o.SetAppKey(appKey)
	return o
}

// SetAppKey adds the appKey to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) SetAppKey(appKey string) {
	o.AppKey = appKey
}

// WithPropertyName adds the propertyName to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) WithPropertyName(propertyName string) *RetrieveUserHostedPropertyValueParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WithSelectedUser adds the selectedUser to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) WithSelectedUser(selectedUser string) *RetrieveUserHostedPropertyValueParams {
	o.SetSelectedUser(selectedUser)
	return o
}

// SetSelectedUser adds the selectedUser to the retrieve user hosted property value params
func (o *RetrieveUserHostedPropertyValueParams) SetSelectedUser(selectedUser string) {
	o.SelectedUser = selectedUser
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveUserHostedPropertyValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_key
	if err := r.SetPathParam("app_key", o.AppKey); err != nil {
		return err
	}

	// path param property_name
	if err := r.SetPathParam("property_name", o.PropertyName); err != nil {
		return err
	}

	// path param selected_user
	if err := r.SetPathParam("selected_user", o.SelectedUser); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
