// This file was auto-generated by Fern from our API Definition.

package bulkretryignoredevents

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	hookdeckgosdk "github.com/hookdeck/hookdeck-go-sdk"
	core "github.com/hookdeck/hookdeck-go-sdk/core"
	io "io"
	http "net/http"
	url "net/url"
	time "time"
)

type Client interface {
	GetIgnoredEventBulkRetries(ctx context.Context, request *hookdeckgosdk.GetIgnoredEventBulkRetriesRequest) (*hookdeckgosdk.BatchOperationPaginatedResult, error)
	CreateIgnoredEventBulkRetry(ctx context.Context, request *hookdeckgosdk.CreateIgnoredEventBulkRetryRequest) (*hookdeckgosdk.BatchOperation, error)
	GenerateIgnoredEventBulkRetryPlan(ctx context.Context) (*hookdeckgosdk.GenerateIgnoredEventBulkRetryPlanResponse, error)
	GetIgnoredEventBulkRetry(ctx context.Context, id string) (*hookdeckgosdk.BatchOperation, error)
	CancelIgnoredEventBulkRetry(ctx context.Context, id string) (*hookdeckgosdk.BatchOperation, error)
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

func (c *client) GetIgnoredEventBulkRetries(ctx context.Context, request *hookdeckgosdk.GetIgnoredEventBulkRetriesRequest) (*hookdeckgosdk.BatchOperationPaginatedResult, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "bulk/ignored-events/retry"

	queryParams := make(url.Values)
	if request.CancelledAt != nil {
		queryParams.Add("cancelled_at", fmt.Sprintf("%v", request.CancelledAt.Format(time.RFC3339)))
	}
	if request.CompletedAt != nil {
		queryParams.Add("completed_at", fmt.Sprintf("%v", request.CompletedAt.Format(time.RFC3339)))
	}
	if request.CreatedAt != nil {
		queryParams.Add("created_at", fmt.Sprintf("%v", request.CreatedAt.Format(time.RFC3339)))
	}
	if request.Id != nil {
		queryParams.Add("id", fmt.Sprintf("%v", *request.Id))
	}
	if request.QueryPartialMatch != nil {
		queryParams.Add("query_partial_match", fmt.Sprintf("%v", *request.QueryPartialMatch))
	}
	if request.InProgress != nil {
		queryParams.Add("in_progress", fmt.Sprintf("%v", *request.InProgress))
	}
	if request.OrderBy != nil {
		queryParams.Add("order_by", fmt.Sprintf("%v", *request.OrderBy))
	}
	if request.Dir != nil {
		queryParams.Add("dir", fmt.Sprintf("%v", *request.Dir))
	}
	if request.Limit != nil {
		queryParams.Add("limit", fmt.Sprintf("%v", *request.Limit))
	}
	if request.Next != nil {
		queryParams.Add("next", fmt.Sprintf("%v", *request.Next))
	}
	if request.Prev != nil {
		queryParams.Add("prev", fmt.Sprintf("%v", *request.Prev))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

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
				return err
			}
			return value
		case 422:
			value := new(hookdeckgosdk.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.BatchOperationPaginatedResult
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) CreateIgnoredEventBulkRetry(ctx context.Context, request *hookdeckgosdk.CreateIgnoredEventBulkRetryRequest) (*hookdeckgosdk.BatchOperation, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "bulk/ignored-events/retry"

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
				return err
			}
			return value
		case 422:
			value := new(hookdeckgosdk.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.BatchOperation
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) GenerateIgnoredEventBulkRetryPlan(ctx context.Context) (*hookdeckgosdk.GenerateIgnoredEventBulkRetryPlanResponse, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "bulk/ignored-events/retry/plan"

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
				return err
			}
			return value
		case 422:
			value := new(hookdeckgosdk.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.GenerateIgnoredEventBulkRetryPlanResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) GetIgnoredEventBulkRetry(ctx context.Context, id string) (*hookdeckgosdk.BatchOperation, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"bulk/ignored-events/retry/%v", id)

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
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.BatchOperation
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) CancelIgnoredEventBulkRetry(ctx context.Context, id string) (*hookdeckgosdk.BatchOperation, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"bulk/ignored-events/retry/%v/cancel", id)

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
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgosdk.BatchOperation
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		nil,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}
