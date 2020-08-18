// Code generated by go-swagger; DO NOT EDIT.

package ssh

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ssh API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ssh API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteUsersUsernameSSHKeys(params *DeleteUsersUsernameSSHKeysParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUsersUsernameSSHKeysNoContent, error)

	GetUsersUsernameSSHKeys(params *GetUsersUsernameSSHKeysParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersUsernameSSHKeysOK, error)

	PostUsersUsernameSSHKeys(params *PostUsersUsernameSSHKeysParams, authInfo runtime.ClientAuthInfoWriter) (*PostUsersUsernameSSHKeysCreated, error)

	PutUsersUsernameSSHKeys(params *PutUsersUsernameSSHKeysParams, authInfo runtime.ClientAuthInfoWriter) (*PutUsersUsernameSSHKeysOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteUsersUsernameSSHKeys Deletes a specific SSH public key from a user's account

Example:
```
$ curl -X DELETE https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/{b15b6026-9c02-4626-b4ad-b905f99f763a}
```
*/
func (a *Client) DeleteUsersUsernameSSHKeys(params *DeleteUsersUsernameSSHKeysParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUsersUsernameSSHKeysNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersUsernameSSHKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUsersUsernameSSHKeys",
		Method:             "DELETE",
		PathPattern:        "/users/{username}/ssh-keys/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsersUsernameSSHKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsersUsernameSSHKeysNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteUsersUsernameSSHKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersUsernameSSHKeys Returns a paginated list of the user's SSH public keys.

Example:

```
$ curl https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys
{
    "page": 1,
    "pagelen": 10,
    "size": 1,
    "values": [
        {
            "comment": "user@myhost",
            "created_on": "2018-03-14T13:17:05.196003+00:00",
            "key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY",
            "label": "",
            "last_used": "2018-03-20T13:18:05.196003+00:00",
            "links": {
                "self": {
                    "href": "https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a"
                }
            },
            "owner": {
                "display_name": "Mark Adams",
                "links": {
                    "avatar": {
                        "href": "https://bitbucket.org/account/markadams-atl/avatar/32/"
                    },
                    "html": {
                        "href": "https://bitbucket.org/markadams-atl/"
                    },
                    "self": {
                        "href": "https://api.bitbucket.org/2.0/users/markadams-atl"
                    }
                },
                "type": "user",
                "username": "markadams-atl",
                "nickname": "markadams-atl",
                "uuid": "{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}"
            },
            "type": "ssh_key",
            "uuid": "{b15b6026-9c02-4626-b4ad-b905f99f763a}"
        }
    ]
}
```
*/
func (a *Client) GetUsersUsernameSSHKeys(params *GetUsersUsernameSSHKeysParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersUsernameSSHKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUsernameSSHKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUsernameSSHKeys",
		Method:             "GET",
		PathPattern:        "/users/{username}/ssh-keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUsernameSSHKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUsernameSSHKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsersUsernameSSHKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUsersUsernameSSHKeys Adds a new SSH public key to the specified user account and returns the resulting key.

Example:
```
$ curl -X POST -H "Content-Type: application/json" -d '{"key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY user@myhost"}' https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys

{
    "comment": "user@myhost",
    "created_on": "2018-03-14T13:17:05.196003+00:00",
    "key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY",
    "label": "",
    "last_used": "2018-03-20T13:18:05.196003+00:00",
    "links": {
        "self": {
            "href": "https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a"
        }
    },
    "owner": {
        "display_name": "Mark Adams",
        "links": {
            "avatar": {
                "href": "https://bitbucket.org/account/markadams-atl/avatar/32/"
            },
            "html": {
                "href": "https://bitbucket.org/markadams-atl/"
            },
            "self": {
                "href": "https://api.bitbucket.org/2.0/users/markadams-atl"
            }
        },
        "type": "user",
        "username": "markadams-atl",
        "nickname": "markadams-atl",
        "uuid": "{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}"
    },
    "type": "ssh_key",
    "uuid": "{b15b6026-9c02-4626-b4ad-b905f99f763a}"
}
```
*/
func (a *Client) PostUsersUsernameSSHKeys(params *PostUsersUsernameSSHKeysParams, authInfo runtime.ClientAuthInfoWriter) (*PostUsersUsernameSSHKeysCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsersUsernameSSHKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostUsersUsernameSSHKeys",
		Method:             "POST",
		PathPattern:        "/users/{username}/ssh-keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsersUsernameSSHKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUsersUsernameSSHKeysCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostUsersUsernameSSHKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutUsersUsernameSSHKeys Updates a specific SSH public key on a user's account

Note: Only the 'comment' field can be updated using this API. To modify the key or comment values, you must delete and add the key again.

Example:
```
$ curl -X PUT -H "Content-Type: application/json" -d '{"label": "Work key"}' https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/{b15b6026-9c02-4626-b4ad-b905f99f763a}

{
    "comment": "",
    "created_on": "2018-03-14T13:17:05.196003+00:00",
    "key": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKqP3Cr632C2dNhhgKVcon4ldUSAeKiku2yP9O9/bDtY",
    "label": "Work key",
    "last_used": "2018-03-20T13:18:05.196003+00:00",
    "links": {
        "self": {
            "href": "https://api.bitbucket.org/2.0/users/markadams-atl/ssh-keys/b15b6026-9c02-4626-b4ad-b905f99f763a"
        }
    },
    "owner": {
        "display_name": "Mark Adams",
        "links": {
            "avatar": {
                "href": "https://bitbucket.org/account/markadams-atl/avatar/32/"
            },
            "html": {
                "href": "https://bitbucket.org/markadams-atl/"
            },
            "self": {
                "href": "https://api.bitbucket.org/2.0/users/markadams-atl"
            }
        },
        "type": "user",
        "username": "markadams-atl",
        "nickname": "markadams-atl",
        "uuid": "{d7dd0e2d-3994-4a50-a9ee-d260b6cefdab}"
    },
    "type": "ssh_key",
    "uuid": "{b15b6026-9c02-4626-b4ad-b905f99f763a}"
}
```
*/
func (a *Client) PutUsersUsernameSSHKeys(params *PutUsersUsernameSSHKeysParams, authInfo runtime.ClientAuthInfoWriter) (*PutUsersUsernameSSHKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUsersUsernameSSHKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutUsersUsernameSSHKeys",
		Method:             "PUT",
		PathPattern:        "/users/{username}/ssh-keys/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutUsersUsernameSSHKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutUsersUsernameSSHKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutUsersUsernameSSHKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
