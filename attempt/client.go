// This file was auto-generated by Fern from our API Definition.

package attempt

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	hookdeckgosdk "github.com/hookdeck/hookdeck-go-sdk"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	option "github.com/hookdeck/hookdeck-go-sdk/option"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

func (c *Client) List(
	ctx context.Context,
	request *hookdeckgosdk.AttemptListRequest,
	opts ...option.RequestOption,
) (*hookdeckgosdk.EventAttemptPaginatedResult, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.hookdeck.com/2024-09-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/attempts"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(hookdeckgosdk.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(hookdeckgosdk.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.EventAttemptPaginatedResult
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Retrieve(
	ctx context.Context,
	id string,
	opts ...option.RequestOption,
) (*hookdeckgosdk.EventAttempt, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.hookdeck.com/2024-09-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/attempts/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(hookdeckgosdk.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.EventAttempt
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:                endpointURL,
			Method:             http.MethodGet,
			MaxAttempts:        options.MaxAttempts,
			Headers:            headers,
			BodyProperties:     options.BodyProperties,
			QueryParameters:    options.QueryParameters,
			Client:             options.HTTPClient,
			Response:           &response,
			ResponseIsOptional: true,
			ErrorDecoder:       errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
