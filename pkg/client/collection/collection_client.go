// Code generated by go-swagger; DO NOT EDIT.

package collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new collection API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for collection API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddCollection(params *AddCollectionParams) (*AddCollectionCreated, error)

	DeleteCollection(params *DeleteCollectionParams) (*DeleteCollectionOK, error)

	GetCollection(params *GetCollectionParams) (*GetCollectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddCollection adds a collection to the database

  Add a new collection to the database
*/
func (a *Client) AddCollection(params *AddCollectionParams) (*AddCollectionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addCollection",
		Method:             "POST",
		PathPattern:        "/v1/collection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddCollectionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCollection deletes a collection from the database

  Delete a collection from the database
*/
func (a *Client) DeleteCollection(params *DeleteCollectionParams) (*DeleteCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCollection",
		Method:             "DELETE",
		PathPattern:        "/v1/collection/{collectionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCollection gets collection information

  Get collection information
*/
func (a *Client) GetCollection(params *GetCollectionParams) (*GetCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCollection",
		Method:             "GET",
		PathPattern:        "/v1/collection/{collectionName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
