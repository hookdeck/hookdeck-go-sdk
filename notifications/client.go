// This file was auto-generated by Fern from our API Definition.

package notifications

import (
	context "context"
	fmt "fmt"
	hookdeckgosdk "github.com/hookdeck/hookdeck-go-sdk"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	http "net/http"
)

type Client interface {
	ListCustomDomains(ctx context.Context, teamId string) (hookdeckgosdk.ListCustomDomainSchema, error)
	AddCustomDomain(ctx context.Context, teamId string, request *hookdeckgosdk.AddCustomHostname) (*hookdeckgosdk.AddCustomHostname, error)
	DeleteCustomDomain(ctx context.Context, teamId string, domainId string) (*hookdeckgosdk.DeleteCustomDomainSchema, error)
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (c *client) ListCustomDomains(ctx context.Context, teamId string) (hookdeckgosdk.ListCustomDomainSchema, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"teams/%v/custom_domains", teamId)

	var response hookdeckgosdk.ListCustomDomainSchema
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) AddCustomDomain(ctx context.Context, teamId string, request *hookdeckgosdk.AddCustomHostname) (*hookdeckgosdk.AddCustomHostname, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"teams/%v/custom_domains", teamId)

	var response *hookdeckgosdk.AddCustomHostname
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) DeleteCustomDomain(ctx context.Context, teamId string, domainId string) (*hookdeckgosdk.DeleteCustomDomainSchema, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"teams/%v/custom_domains/%v", teamId, domainId)

	var response *hookdeckgosdk.DeleteCustomDomainSchema
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodDelete,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
