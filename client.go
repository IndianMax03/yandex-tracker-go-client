// The client package provides methods, values and models for interacting with Yandex Tracker
package client

import (
	"fmt"
	"io"

	"resty.dev/v3"
)

const (
	baseURL            = "https://api.tracker.yandex.net/v2/"
	defaultContentType = "application/json"
	defaultLang        = "ru"
	defaultAuthScheme  = "OAuth"
)

// Client is a wrapper over the resty.Client type with Yandex Tracker API-specific headers and a base URL
type Client struct {
	restyClient *resty.Client
}

// Create a new Client
func New(tokenOAuth, xCloudOrgID, xOrgID, acceptLanguage string) *Client {
	var lang string
	headers := map[string]string{}

	if acceptLanguage != "ru" && acceptLanguage != "en" {
		lang = defaultLang
	}
	headers["Accept-Language"] = lang

	if xOrgID != "" {
		headers["X-Org-ID"] = xOrgID
	} else if xCloudOrgID != "" {
		headers["X-Cloud-Org-ID"] = xCloudOrgID
	}

	restyClient := resty.New()
	restyClient.SetHeaders(headers)
	restyClient.SetAuthScheme(defaultAuthScheme)
	restyClient.SetAuthToken(tokenOAuth)
	restyClient.SetBaseURL(baseURL)

	return &Client{
		restyClient: restyClient,
	}
}

// Send request to Yandex Tracker
func (c *Client) SendRequest(method, resourceURL string, requestBody, responseBody any) (resp *resty.Response, err error) {
	req := c.restyClient.R().
		SetContentType(defaultContentType).
		SetMethod(method).
		SetBody(requestBody).
		SetResult(responseBody).
		SetURL(c.restyClient.BaseURL() + resourceURL)
	resp, err = req.Send()
	return
}

// Allow logging of details of each request and response.
func (c *Client) SetDebug(debug bool) {
	c.restyClient.SetDebug(debug)
}

// Create a new issue
func (c *Client) CreateIssue(req *IssueCreateRequest) (*IssueCreateResponse, error) {
	var respBody IssueCreateResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		issuesCreateURL,
		req,
		&respBody,
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("Request failed with status code: %s. Body: %s", res.Status(), body)
	}
	return &respBody, nil
}

// Get the number of issues
func (c *Client) GetIssuesCount(req *IssueCountRequest) (int, error) {
	var respBody IssueCountResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		issuesCountURL,
		req,
		&respBody,
	)
	if err != nil {
		return 0, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return 0, fmt.Errorf("Request failed with status code: %s. Body: %s", res.Status(), body)
	}
	return int(respBody), nil
}
