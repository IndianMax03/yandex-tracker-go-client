// Package client provides methods, values and urls for interacting with Yandex Tracker
package client

import (
	"fmt"
	"io"
	"net/url"
	"strconv"

	model "github.com/IndianMax03/yandex-tracker-go-client/model"
	"resty.dev/v3"
)

const (
	baseURL            = "https://api.tracker.yandex.net/v2/"
	defaultContentType = "application/json"
	defaultLang        = "ru"
	defaultAuthScheme  = "OAuth"
	defaultPerPage     = 50
)

// Client is a wrapper over the resty.Client type with Yandex Tracker API-specific headers and a base URL
type Client struct {
	restyClient *resty.Client
}

// New Yandex Tracker Client
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

// SendRequest sends request to Yandex Tracker
func (c *Client) SendRequest(
	method,
	resourceURL string,
	queryParams map[string]string,
	multiplyQueryParams url.Values,
	pathParams map[string]string,
	requestBody,
	responseBody any,
) (resp *resty.Response, err error) {
	req := c.restyClient.R().
		SetContentType(defaultContentType).
		SetMethod(method).
		SetBody(requestBody).
		SetResult(responseBody).
		SetURL(c.restyClient.BaseURL() + resourceURL).
		SetQueryParams(queryParams).
		SetQueryParamsFromValues(multiplyQueryParams).
		SetPathParams(pathParams)
	resp, err = req.Send()
	return
}

// SendMultipartRequest sends multipart request to Yandex Tracker
func (c *Client) SendMultipartRequest(
	method,
	resourceURL string,
	queryParams map[string]string,
	multiplyQueryParams url.Values,
	pathParams map[string]string,
	requestBody *resty.MultipartField,
	responseBody any,
) (resp *resty.Response, err error) {
	req := c.restyClient.R().
		SetContentType(defaultContentType).
		SetMethod(method).
		SetMultipartFields(requestBody).
		SetResult(responseBody).
		SetURL(c.restyClient.BaseURL() + resourceURL).
		SetQueryParams(queryParams).
		SetQueryParamsFromValues(multiplyQueryParams).
		SetPathParams(pathParams)

	resp, err = req.Send()
	return
}

// SetDebug allows logging of details of each request and response.
func (c *Client) SetDebug(debug bool) {
	c.restyClient.SetDebug(debug)
}

// CreateIssue sends request to create new issue in Yandex Tracker
func (c *Client) CreateIssue(req *model.IssueCreateRequest) (*model.IssueResponse, error) {
	var respBody model.IssueResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		issuesCreateURL,
		nil,
		nil,
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

// GetIssue sends request to find concrete issue by ID
func (c *Client) GetIssue(issueID string, includeAttachments, includeTransitions bool) (*model.IssueResponse, error) {
	var respBody model.IssueResponse
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	var multiplyQueryParams url.Values
	if includeAttachments || includeTransitions {
		values := []string{}
		if includeAttachments {
			values = append(values, "attachments")
		}
		if includeTransitions {
			values = append(values, "transitions")
		}
		multiplyQueryParams = url.Values{
			"expand": values,
		}
	}
	res, err := c.SendRequest(
		resty.MethodGet,
		issuesGetURL,
		nil,
		multiplyQueryParams,
		pathParams,
		nil,
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

// GetIssuesCount sends a request to get the number of tasks
func (c *Client) GetIssuesCount(req *model.IssueCountRequest) (int, error) {
	var respBody model.IssueCountResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		issuesCountURL,
		nil,
		nil,
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

// SearchIssuesPage sends a request to find issues using pagination
func (c *Client) SearchIssuesPage(req *model.IssueSearchRequest, pageReq *model.PageRequest) ([]model.IssueResponse, *model.PageResponse, error) {
	if pageReq.PerPage <= 0 {
		pageReq.PerPage = 5
	}
	if pageReq.Page <= 0 {
		pageReq.Page = 1
	}
	queryParams := make(map[string]string)
	queryParams["perPage"] = strconv.Itoa(pageReq.PerPage)
	queryParams["page"] = strconv.Itoa(pageReq.Page)

	var respBody []model.IssueResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		issuesSearchURL,
		queryParams,
		nil,
		nil,
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
	pageResp := model.PageResponse{
		TotalPages: totalPages,
		TotalCount: totalCount,
	}
	return respBody, &pageResp, nil
}

// SearchAllIssues sends a request to find all issues
func (c *Client) SearchAllIssues(req *model.IssueSearchRequest) ([]model.IssueResponse, error) {
	currentPage := 1
	pageReq := model.PageRequest{
		Page:    currentPage,
		PerPage: defaultPerPage,
	}

	result, pag, err := c.SearchIssuesPage(req, &pageReq)
	if err != nil {
		return nil, err
	}
	totalPages := pag.TotalPages
	for currentPage < totalPages {
		currentPage++
		pageReq.Page = currentPage
		resp, _, _ := c.SearchIssuesPage(req, &pageReq)
		result = append(result, resp...)
	}
	return result, nil
}

// ModifyIssue sends a request to modify existing issue
func (c *Client) ModifyIssue(issueID string, req *model.IssueModifyRequest) (*model.IssueResponse, error) {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	var respBody model.IssueResponse
	res, err := c.SendRequest(
		resty.MethodPatch,
		issuesModifyURL,
		nil,
		nil,
		pathParams,
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

// ModifyIssueStatus sends a request to modify existing issue status
func (c *Client) ModifyIssueStatus(issueID string, transitionID string, req *model.IssueModifyStatusRequest) ([]model.IssueModifyStatusResponse, error) {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	pathParams["transition_id"] = transitionID
	var respBody []model.IssueModifyStatusResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		issuesModifyStatusURL,
		nil,
		nil,
		pathParams,
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
	return respBody, nil
}

// GetIssueTransitions sends a request to find all possible issue transitions
func (c *Client) GetIssueTransitions(issueID string) ([]model.IssueTransitionsResponse, error) {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	var respBody []model.IssueTransitionsResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		issueGetTransitionsURL,
		nil,
		nil,
		pathParams,
		nil,
		&respBody,
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}
	return respBody, nil
}

// GetPrioritiesPage sends a request to find priorities using pagination
func (c *Client) GetPrioritiesPage(localized bool, pageReq *model.PageRequest) ([]model.PriorityResponse, *model.PageResponse, error) {
	if pageReq.PerPage <= 0 {
		pageReq.PerPage = 5
	}
	if pageReq.Page <= 0 {
		pageReq.Page = 1
	}
	queryParams := make(map[string]string)
	queryParams["perPage"] = strconv.Itoa(pageReq.PerPage)
	queryParams["page"] = strconv.Itoa(pageReq.Page)

	queryParams["localized"] = strconv.FormatBool(localized)
	var respBody []model.PriorityResponse

	res, err := c.SendRequest(
		resty.MethodGet,
		prioritiesGetURL,
		queryParams,
		nil,
		nil,
		nil,
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
	pageResp := model.PageResponse{
		TotalPages: totalPages,
		TotalCount: totalCount,
	}
	return respBody, &pageResp, nil
}

// GetAllPriorities sends a request to find all priorities
func (c *Client) GetAllPriorities(localized bool) ([]model.PriorityResponse, error) {
	currentPage := 1
	pageReq := model.PageRequest{
		Page:    currentPage,
		PerPage: defaultPerPage,
	}

	result, pag, err := c.GetPrioritiesPage(localized, &pageReq)
	if err != nil {
		return nil, err
	}
	totalPages := pag.TotalPages
	for currentPage < totalPages {
		currentPage++
		pageReq.Page = currentPage
		resp, _, _ := c.GetPrioritiesPage(localized, &pageReq)
		result = append(result, resp...)
	}
	return result, nil
}

// GetPriority sends a request to find concrete priority
func (c *Client) GetPriority(priorityID int, localized bool) (*model.PriorityResponse, error) {
	queryParams := make(map[string]string)
	queryParams["localized"] = strconv.FormatBool(localized)
	pathParams := make(map[string]string)
	pathParams["priority_id"] = strconv.Itoa(priorityID)

	var respBody model.PriorityResponse

	res, err := c.SendRequest(
		resty.MethodGet,
		priorityGetURL,
		queryParams,
		nil,
		pathParams,
		nil,
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

// CreateComment sends a request to add a comment to a issue
func (c *Client) CreateComment(issueID string, req *model.CommentRequest) (*model.CommentResponse, error) {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	var respBody model.CommentResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		issueAppendCommentURL,
		nil,
		nil,
		pathParams,
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

// GetComment sends a request to get concrete comment to a issue
func (c *Client) GetComment(issueID string, commentID int) (*model.CommentResponse, error) {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	commentStringID := strconv.Itoa(commentID)
	pathParams["comment_id"] = commentStringID
	var respBody model.CommentResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		issueGetCommentURL,
		nil,
		nil,
		pathParams,
		nil,
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

// GetXCommentsAfterY sends a request to get first X comments after comment with ID=Y (model.PageRequest.PerPage = X, model.PageRequest.FromID = Y)
func (c *Client) GetXCommentsAfterY(issueID string, commentExpand string, pageReq *model.PageRequest) ([]model.CommentResponse, *model.PageResponse, error) {
	if pageReq.PerPage <= 0 {
		pageReq.PerPage = defaultPerPage
	}
	queryParams := make(map[string]string)
	if commentExpand != model.ExpandNone {
		queryParams["expand"] = commentExpand
	}

	queryParams["perPage"] = strconv.Itoa(pageReq.PerPage)
	if pageReq.FromID > 0 {
		queryParams["id"] = strconv.Itoa(pageReq.FromID)
	}

	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	var respBody []model.CommentResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		issueGetCommentsURL,
		queryParams,
		nil,
		pathParams,
		nil,
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
	var lastID int
	if len(respBody) != 0 {
		lastID = respBody[len(respBody)-1].ID
	}
	pageResp := model.PageResponse{
		TotalPages: totalPages,
		TotalCount: totalCount,
		LastID:     lastID,
	}
	return respBody, &pageResp, nil
}

// GetCommentsAll sends requests to get all of the comments using default perPage size
func (c *Client) GetCommentsAll(issueID string, commentExpand string) ([]model.CommentResponse, error) {
	pageReq := model.PageRequest{
		PerPage: defaultPerPage,
	}

	result, pag, err := c.GetXCommentsAfterY(issueID, commentExpand, &pageReq)
	if err != nil {
		return nil, err
	}
	fromID := pag.LastID
	for fromID > 0 {
		pageReq.FromID = fromID
		resp, pagResp, _ := c.GetXCommentsAfterY(issueID, commentExpand, &pageReq)
		fromID = pagResp.LastID
		result = append(result, resp...)
	}
	return result, nil
}

// UpdateComment sends a request to update a comment to a issue
func (c *Client) UpdateComment(issueID string, commentID int, req *model.CommentUpdateRequest) (*model.CommentResponse, error) {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	commentStringID := strconv.Itoa(commentID)
	pathParams["comment_id"] = commentStringID
	var respBody model.CommentResponse
	res, err := c.SendRequest(
		resty.MethodPatch,
		issueUpdateCommentURL,
		nil,
		nil,
		pathParams,
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

// DeleteComment sends a request to delete a comment to a issue
func (c *Client) DeleteComment(issueID string, commentID int) error {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	commentStringID := strconv.Itoa(commentID)
	pathParams["comment_id"] = commentStringID
	var respBody model.CommentResponse
	res, err := c.SendRequest(
		resty.MethodDelete,
		issueDeleteCommentURL,
		nil,
		nil,
		pathParams,
		nil,
		&respBody,
	)
	if err != nil {
		return err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}
	return nil
}

// GetMyself sends request to get information about the user account on whose behalf the API call is made.
func (c *Client) GetMyself() (*model.UserResponse, error) {
	var respBody model.UserResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		myselfURL,
		nil,
		nil,
		nil,
		nil,
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

// GetUsersPage sends a request to find users using pagination
func (c *Client) GetUsersPage(pageReq *model.PageRequest) ([]model.UserResponse, *model.PageResponse, error) {
	if pageReq.PerPage <= 0 {
		pageReq.PerPage = 5
	}
	if pageReq.Page <= 0 {
		pageReq.Page = 1
	}
	queryParams := make(map[string]string)
	queryParams["perPage"] = strconv.Itoa(pageReq.PerPage)
	queryParams["page"] = strconv.Itoa(pageReq.Page)

	var respBody []model.UserResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		usersGetURL,
		queryParams,
		nil,
		nil,
		nil,
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
	pageResp := model.PageResponse{
		TotalPages: totalPages,
		TotalCount: totalCount,
	}
	return respBody, &pageResp, nil
}

// GetUsersAll sends a request to find all users
func (c *Client) GetUsersAll() ([]model.UserResponse, error) {
	currentPage := 1
	pageReq := model.PageRequest{
		Page:    currentPage,
		PerPage: defaultPerPage,
	}

	result, pag, err := c.GetUsersPage(&pageReq)
	if err != nil {
		return nil, err
	}
	totalPages := pag.TotalPages
	for currentPage < totalPages {
		currentPage++
		pageReq.Page = currentPage
		resp, _, _ := c.GetUsersPage(&pageReq)
		result = append(result, resp...)
	}
	return result, nil
}

// GetUser sends request to get information about concrete user (login is a priority).
func (c *Client) GetUser(login string, userID int) (*model.UserResponse, error) {
	pathParams := make(map[string]string)
	if login != "" {
		allNums := true
		for i := range len(login) {
			if login[i] < '0' || login[i] > '9' {
				allNums = false
				break
			}
		}
		if allNums {
			pathParams["login_or_user_id"] = fmt.Sprintf("login:%s", login)
		} else {
			pathParams["login_or_user_id"] = login
		}

	} else {
		pathParams["login_or_user_id"] = strconv.Itoa(userID)
	}
	var respBody model.UserResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		userGetURL,
		nil,
		nil,
		pathParams,
		nil,
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

// CreateComponent sends a request to create a component to a queue
func (c *Client) CreateComponent(req *model.ComponentRequest) (*model.ComponentResponse, error) {
	var respBody model.ComponentResponse
	res, err := c.SendRequest(
		resty.MethodPost,
		componentCreateURL,
		nil,
		nil,
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

// UpdateComponent sends a request to update a component to a queue
func (c *Client) UpdateComponent(componentID, componentVersion int, req *model.ComponentUpdateRequest) (*model.ComponentResponse, error) {
	queryParams := make(map[string]string)
	queryParams["version"] = strconv.Itoa(componentVersion)

	pathParams := make(map[string]string)
	pathParams["component_id"] = strconv.Itoa(componentID)

	var respBody model.ComponentResponse
	res, err := c.SendRequest(
		resty.MethodPatch,
		componentUpdateURL,
		queryParams,
		nil,
		pathParams,
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

// GetComponentsPage sends a request to get components using pagination
func (c *Client) GetComponentsPage(pageReq *model.PageRequest) ([]model.ComponentResponse, *model.PageResponse, error) {
	if pageReq.PerPage <= 0 {
		pageReq.PerPage = 5
	}
	if pageReq.Page <= 0 {
		pageReq.Page = 1
	}
	queryParams := make(map[string]string)
	queryParams["perPage"] = strconv.Itoa(pageReq.PerPage)
	queryParams["page"] = strconv.Itoa(pageReq.Page)

	var respBody []model.ComponentResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		componentsGetURL,
		queryParams,
		nil,
		nil,
		nil,
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
	pageResp := model.PageResponse{
		TotalPages: totalPages,
		TotalCount: totalCount,
	}
	return respBody, &pageResp, nil
}

// GetComponentsAll sends a request to find all components
func (c *Client) GetComponentsAll() ([]model.ComponentResponse, error) {
	currentPage := 1
	pageReq := model.PageRequest{
		Page:    currentPage,
		PerPage: defaultPerPage,
	}

	result, pag, err := c.GetComponentsPage(&pageReq)
	if err != nil {
		return nil, err
	}
	totalPages := pag.TotalPages
	for currentPage < totalPages {
		currentPage++
		pageReq.Page = currentPage
		resp, _, _ := c.GetComponentsPage(&pageReq)
		result = append(result, resp...)
	}
	return result, nil
}

// GetComponent sends request to get information about concrete component.
func (c *Client) GetComponent(componentID int) (*model.ComponentResponse, error) {
	pathParams := make(map[string]string)
	pathParams["component_id"] = strconv.Itoa(componentID)
	var respBody model.ComponentResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		componentGetURL,
		nil,
		nil,
		pathParams,
		nil,
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

// GetIssueAttachments sends request to get attachments to issue.
func (c *Client) GetIssueAttachments(issueID string) ([]model.AttachmentFileResponse, error) {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	var respBody []model.AttachmentFileResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		issueGetAttachmentsURL,
		nil,
		nil,
		pathParams,
		nil,
		&respBody,
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}
	return respBody, nil
}

// GetIssueAttachment sends request to get concrete attachment to issue.
func (c *Client) GetIssueAttachment(issueID, attachmentID string) (*model.AttachmentFileResponse, error) {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	pathParams["attachment_id"] = attachmentID
	var respBody *model.AttachmentFileResponse
	res, err := c.SendRequest(
		resty.MethodGet,
		issueGetAttachmentURL,
		nil,
		nil,
		pathParams,
		nil,
		&respBody,
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}
	return respBody, nil
}

// UploadTemporaryAttachment sends request to upload temporary attachment.
func (c *Client) UploadTemporaryAttachment(multipartReq *resty.MultipartField) (*model.AttachmentFileResponse, error) {
	var respBody *model.AttachmentFileResponse
	multipartReq.Name = "filename"
	res, err := c.SendMultipartRequest(
		resty.MethodPost,
		attachmentUploadURL,
		nil,
		nil,
		nil,
		multipartReq,
		&respBody,
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}
	return respBody, nil
}

// IssueAttachFile sends request to upload an attachment to attach to issue.
func (c *Client) IssueAttachFile(issueID string, multipartReq *resty.MultipartField) (*model.AttachmentFileResponse, error) {
	var respBody *model.AttachmentFileResponse
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	multipartReq.Name = "filename"
	res, err := c.SendMultipartRequest(
		resty.MethodPost,
		issueAttachFileURL,
		nil,
		nil,
		pathParams,
		multipartReq,
		&respBody,
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}
	return respBody, nil
}

// IssueDeleteFile sends request to delete an attachment in issue.
func (c *Client) IssueDeleteFile(issueID, fileID string) error {
	pathParams := make(map[string]string)
	pathParams["issue_id"] = issueID
	pathParams["file_id"] = fileID
	res, err := c.SendRequest(
		resty.MethodDelete,
		issueDeleteFileURL,
		nil,
		nil,
		pathParams,
		nil,
		nil,
	)
	if err != nil {
		return err
	}
	if res.IsError() {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("request failed with status code: %s. body: %s", res.Status(), body)
	}
	return nil
}
