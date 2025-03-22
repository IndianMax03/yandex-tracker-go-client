package client

import "resty.dev/v3"

const (
	baseURL            = "https://api.tracker.yandex.net/v2/"
	defaultContentType = "application/json"
	defaultLang        = "ru"
	defaultAuthScheme  = "OAuth"
)

type Client struct {
	restyClient *resty.Client
}

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

func (c *Client) SendRequest(method, resourceURL string, requestBody, responseBody any) (resp *resty.Response, err error) {
	req := c.restyClient.R().
		SetContentType(defaultContentType).
		SetMethod(method).
		SetBody(requestBody).
		SetResult(responseBody).
		SetURL(c.restyClient.BaseURL() + resourceURL)

	req.SetDebug(true)
	resp, err = req.Send()
	return
}
