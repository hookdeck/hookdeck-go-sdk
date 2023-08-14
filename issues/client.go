// This file was auto-generated by Fern from our API Definition.

package issues

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
	GetIssues(ctx context.Context, request *hookdeckgo.GetIssuesRequest) (*hookdeckgo.IssueWithDataPaginatedResult, error)
	GetIssueCount(ctx context.Context, request *hookdeckgo.GetIssueCountRequest) (*hookdeckgo.IssueCount, error)
	GetIssue(ctx context.Context, id string) (*hookdeckgo.IssueWithData, error)
	UpdateIssue(ctx context.Context, id string, request *hookdeckgo.UpdateIssueRequest) (*hookdeckgo.Issue, error)
	DismissIssue(ctx context.Context, id string) (*hookdeckgo.Issue, error)
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

func (c *client) GetIssues(ctx context.Context, request *hookdeckgo.GetIssuesRequest) (*hookdeckgo.IssueWithDataPaginatedResult, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "issues"

	queryParams := make(url.Values)
	if request.Id != nil {
		queryParams.Add("id", fmt.Sprintf("%v", *request.Id))
	}
	if request.IssueTriggerId != nil {
		queryParams.Add("issue_trigger_id", fmt.Sprintf("%v", *request.IssueTriggerId))
	}
	if request.Type != nil {
		queryParams.Add("type", fmt.Sprintf("%v", *request.Type))
	}
	if request.Status != nil {
		queryParams.Add("status", fmt.Sprintf("%v", *request.Status))
	}
	if request.MergedWith != nil {
		queryParams.Add("merged_with", fmt.Sprintf("%v", *request.MergedWith))
	}
	if request.CreatedAt != nil {
		queryParams.Add("created_at", fmt.Sprintf("%v", request.CreatedAt.Format(time.RFC3339)))
	}
	if request.FirstSeenAt != nil {
		queryParams.Add("first_seen_at", fmt.Sprintf("%v", request.FirstSeenAt.Format(time.RFC3339)))
	}
	if request.LastSeenAt != nil {
		queryParams.Add("last_seen_at", fmt.Sprintf("%v", request.LastSeenAt.Format(time.RFC3339)))
	}
	if request.DismissedAt != nil {
		queryParams.Add("dismissed_at", fmt.Sprintf("%v", request.DismissedAt.Format(time.RFC3339)))
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

	var response *hookdeckgo.IssueWithDataPaginatedResult
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

func (c *client) GetIssueCount(ctx context.Context, request *hookdeckgo.GetIssueCountRequest) (*hookdeckgo.IssueCount, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "issues/count"

	queryParams := make(url.Values)
	if request.Id != nil {
		queryParams.Add("id", fmt.Sprintf("%v", *request.Id))
	}
	if request.IssueTriggerId != nil {
		queryParams.Add("issue_trigger_id", fmt.Sprintf("%v", *request.IssueTriggerId))
	}
	if request.Type != nil {
		queryParams.Add("type", fmt.Sprintf("%v", *request.Type))
	}
	if request.Status != nil {
		queryParams.Add("status", fmt.Sprintf("%v", *request.Status))
	}
	if request.MergedWith != nil {
		queryParams.Add("merged_with", fmt.Sprintf("%v", *request.MergedWith))
	}
	if request.CreatedAt != nil {
		queryParams.Add("created_at", fmt.Sprintf("%v", request.CreatedAt.Format(time.RFC3339)))
	}
	if request.FirstSeenAt != nil {
		queryParams.Add("first_seen_at", fmt.Sprintf("%v", request.FirstSeenAt.Format(time.RFC3339)))
	}
	if request.LastSeenAt != nil {
		queryParams.Add("last_seen_at", fmt.Sprintf("%v", request.LastSeenAt.Format(time.RFC3339)))
	}
	if request.DismissedAt != nil {
		queryParams.Add("dismissed_at", fmt.Sprintf("%v", request.DismissedAt.Format(time.RFC3339)))
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

	var response *hookdeckgo.IssueCount
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

func (c *client) GetIssue(ctx context.Context, id string) (*hookdeckgo.IssueWithData, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"issues/%v", id)

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

	var response *hookdeckgo.IssueWithData
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

func (c *client) UpdateIssue(ctx context.Context, id string, request *hookdeckgo.UpdateIssueRequest) (*hookdeckgo.Issue, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"issues/%v", id)

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

	var response *hookdeckgo.Issue
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPut,
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

func (c *client) DismissIssue(ctx context.Context, id string) (*hookdeckgo.Issue, error) {
	baseURL := "https://api.hookdeck.com/2023-07-01"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"issues/%v", id)

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

	var response *hookdeckgo.Issue
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodDelete,
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