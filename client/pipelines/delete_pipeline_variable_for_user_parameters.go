// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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

// NewDeletePipelineVariableForUserParams creates a new DeletePipelineVariableForUserParams object
// with the default values initialized.
func NewDeletePipelineVariableForUserParams() *DeletePipelineVariableForUserParams {
	var ()
	return &DeletePipelineVariableForUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePipelineVariableForUserParamsWithTimeout creates a new DeletePipelineVariableForUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePipelineVariableForUserParamsWithTimeout(timeout time.Duration) *DeletePipelineVariableForUserParams {
	var ()
	return &DeletePipelineVariableForUserParams{

		timeout: timeout,
	}
}

// NewDeletePipelineVariableForUserParamsWithContext creates a new DeletePipelineVariableForUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePipelineVariableForUserParamsWithContext(ctx context.Context) *DeletePipelineVariableForUserParams {
	var ()
	return &DeletePipelineVariableForUserParams{

		Context: ctx,
	}
}

// NewDeletePipelineVariableForUserParamsWithHTTPClient creates a new DeletePipelineVariableForUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePipelineVariableForUserParamsWithHTTPClient(client *http.Client) *DeletePipelineVariableForUserParams {
	var ()
	return &DeletePipelineVariableForUserParams{
		HTTPClient: client,
	}
}

/*DeletePipelineVariableForUserParams contains all the parameters to send to the API endpoint
for the delete pipeline variable for user operation typically these are written to a http.Request
*/
type DeletePipelineVariableForUserParams struct {

	/*SelectedUser
	  Either the UUID of the account surrounded by curly-braces, for example `{account UUID}`, OR an Atlassian Account ID.

	*/
	SelectedUser string
	/*VariableUUID
	  The UUID of the variable to delete.

	*/
	VariableUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) WithTimeout(timeout time.Duration) *DeletePipelineVariableForUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) WithContext(ctx context.Context) *DeletePipelineVariableForUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) WithHTTPClient(client *http.Client) *DeletePipelineVariableForUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSelectedUser adds the selectedUser to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) WithSelectedUser(selectedUser string) *DeletePipelineVariableForUserParams {
	o.SetSelectedUser(selectedUser)
	return o
}

// SetSelectedUser adds the selectedUser to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) SetSelectedUser(selectedUser string) {
	o.SelectedUser = selectedUser
}

// WithVariableUUID adds the variableUUID to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) WithVariableUUID(variableUUID string) *DeletePipelineVariableForUserParams {
	o.SetVariableUUID(variableUUID)
	return o
}

// SetVariableUUID adds the variableUuid to the delete pipeline variable for user params
func (o *DeletePipelineVariableForUserParams) SetVariableUUID(variableUUID string) {
	o.VariableUUID = variableUUID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePipelineVariableForUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param selected_user
	if err := r.SetPathParam("selected_user", o.SelectedUser); err != nil {
		return err
	}

	// path param variable_uuid
	if err := r.SetPathParam("variable_uuid", o.VariableUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
