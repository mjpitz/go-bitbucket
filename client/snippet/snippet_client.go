// Code generated by go-swagger; DO NOT EDIT.

package snippet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new snippet API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for snippet API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetSnippetsWorkspaceEncodedIDFilesPath(params *GetSnippetsWorkspaceEncodedIDFilesPathParams, authInfo runtime.ClientAuthInfoWriter) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSnippetsWorkspaceEncodedIDFilesPath Convenience resource for getting to a snippet's raw files without the
need for first having to retrieve the snippet itself and having to pull
out the versioned file links.
*/
func (a *Client) GetSnippetsWorkspaceEncodedIDFilesPath(params *GetSnippetsWorkspaceEncodedIDFilesPathParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSnippetsWorkspaceEncodedIDFilesPathParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSnippetsWorkspaceEncodedIDFilesPath",
		Method:             "GET",
		PathPattern:        "/snippets/{workspace}/{encoded_id}/files/{path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSnippetsWorkspaceEncodedIDFilesPathReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}