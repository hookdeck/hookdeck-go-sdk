// This file was auto-generated by Fern from our API Definition.

package bulkretryevents

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	hookdeckgo "github.com/fern-hookdeck/hookdeck-go"
	core "github.com/fern-hookdeck/hookdeck-go/core"
	io "io"
	http "net/http"
	url "net/url"
	time "time"
)

type Client interface {
	GetEventBulkRetries(ctx context.Context, request *hookdeckgo.GetEventBulkRetriesRequest) (*hookdeckgo.BatchOperationPaginatedResult, error)
	CreateEventBulkRetry(ctx context.Context, request *hookdeckgo.CreateEventBulkRetryRequest) (*hookdeckgo.BatchOperation, error)
	GenerateEventBulkRetryPlan(ctx context.Context) (*hookdeckgo.GenerateEventBulkRetryPlanResponse, error)
	GetEventBulkRetry(ctx context.Context, id string) (*hookdeckgo.BatchOperation, error)
	CancelEventBulkRetry(ctx context.Context, id string) (*hookdeckgo.BatchOperation, error)
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

func (c *client) GetEventBulkRetries(ctx context.Context, request *hookdeckgo.GetEventBulkRetriesRequest) (*hookdeckgo.BatchOperationPaginatedResult, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "bulk/events/retry"

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
			value := new(hookdeckgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(hookdeckgo.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgo.BatchOperationPaginatedResult
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

func (c *client) CreateEventBulkRetry(ctx context.Context, request *hookdeckgo.CreateEventBulkRetryRequest) (*hookdeckgo.BatchOperation, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "bulk/events/retry"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(hookdeckgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(hookdeckgo.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgo.BatchOperation
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

func (c *client) GenerateEventBulkRetryPlan(ctx context.Context) (*hookdeckgo.GenerateEventBulkRetryPlanResponse, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "bulk/events/retry/plan"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(hookdeckgo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 422:
			value := new(hookdeckgo.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgo.GenerateEventBulkRetryPlanResponse
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

func (c *client) GetEventBulkRetry(ctx context.Context, id string) (*hookdeckgo.BatchOperation, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"bulk/events/retry/%v", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(hookdeckgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgo.BatchOperation
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

func (c *client) CancelEventBulkRetry(ctx context.Context, id string) (*hookdeckgo.BatchOperation, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"bulk/events/retry/%v/cancel", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(hookdeckgo.NotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *hookdeckgo.BatchOperation
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