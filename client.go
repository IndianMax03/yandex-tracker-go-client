// The client package provides methods, values and models for interacting with Yandex Tracker
// The client package provides methods, values and models for interacting with Yandex Tracker
package client

import (
	"fmt"
	"io"
	"strconv"

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
func (c *Client) SendRequest(method, resourceURL string, queryParams map[string]string, requestBody, responseBody any) (resp *resty.Response, err error) {
	req := c.restyClient.R().
		SetContentType(defaultContentType).
		SetMethod(method).
		SetBody(requestBody).
		SetResult(responseBody).
		SetURL(c.restyClient.BaseURL() + resourceURL).
		SetQueryParams(queryParams)
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
		nil,
		req,
		&respBody,
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}
	return &respBody, nil
}

// Get the number of issues
func (c *Client) GetIssuesCount(req *IssueCountRequest) (int, error) {
	var respBody IssueCountResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		issuesCountURL,
		nil,
		req,
		&respBody,
	)
	if err != nil {
		return 0, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return 0, fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}
	return int(respBody), nil
}

// Find issues using pagination
func (c *Client) SearchIssuesPage(req *IssueSearchRequest, pageReq *PageRequest) (*IssueSearchResponse, *PageResponse, error) {
	if pageReq.PerPage <= 0 {
		pageReq.PerPage = 5
	}
	if pageReq.Page <= 0 {
		pageReq.Page = 1
	}
	queryParams := make(map[string]string)
	queryParams["perPage"] = strconv.Itoa(pageReq.PerPage)
	queryParams["page"] = strconv.Itoa(pageReq.Page)

	var respBody IssueSearchResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		issuesSearchURL,
		queryParams,
		req,
		&respBody,
	)
	if err != nil {
		return nil, nil, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, nil, fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}

	totalPages, _ := strconv.Atoi(res.Header().Get("X-Total-Pages"))
	totalCount, _ := strconv.Atoi(res.Header().Get("X-Total-Count"))
	pageResp := PageResponse{
		TotalPages: totalPages,
		TotalCount: totalCount,
	}
	return &respBody, &pageResp, nil
}

// Find all issues
func (c *Client) SearchAllIssues(req *IssueSearchRequest) (*IssueSearchResponse, error) {
	currentPage := 1
	pageReq := PageRequest{
		Page:    currentPage,
		PerPage: 50,
	}

	if result, pag, err := c.SearchIssuesPage(req, &pageReq); err != nil {
		return nil, err
	} else {
		totalPages := pag.TotalPages
		for currentPage < totalPages {
			currentPage += 1
			pageReq.Page = currentPage
			resp, _, _ := c.SearchIssuesPage(req, &pageReq)
			*result = append(*result, *resp...)
		}
		return result, nil
	}
}
