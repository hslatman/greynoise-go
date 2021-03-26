package client

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	urllib "net/url"
	"time"

	"github.com/hslatman/greynoise-go/responses"
)

const (
	version = "0.1.0"

	defaultBaseURL           = "https://api.greynoise.io"
	defaultUserAgent         = "https://github.com/hslatman/greynoise-go"
	defaultContentTypeHeader = "application/json"
	defaultAcceptHeader      = "application/json"
	defaultTimeOut           = 30 * time.Second
)

const (
	communityEndpoint = "v3/community/%s"
	pingEndpoint      = "ping"

	// TODO: other endpoints
)

type Client struct {
	key               string
	httpClient        *http.Client
	baseURL           *urllib.URL
	userAgent         string
	contentTypeHeader string
	acceptHeader      string
}

func New(key string) (*Client, error) {

	baseURL, err := urllib.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	httpClient := http.DefaultClient
	httpClient.Timeout = defaultTimeOut

	client := &Client{
		key:               key,
		httpClient:        httpClient,
		baseURL:           baseURL,
		userAgent:         defaultUserAgent + "/" + version,
		contentTypeHeader: defaultContentTypeHeader,
		acceptHeader:      defaultAcceptHeader,
	}

	return client, nil
}

func (c *Client) Ping() (bool, error) {
	r, err := c.execute(http.MethodGet, pingEndpoint, nil)
	if err != nil {
		return false, err
	}
	return r.StatusCode == http.StatusOK, nil
}

func (c *Client) Community(ip net.IP) (responses.Community, error) {

	communityURL := fmt.Sprintf(communityEndpoint, ip.String())

	r, err := c.execute(http.MethodGet, communityURL, nil)
	if err != nil {
		return responses.Community{}, err
	}

	body := r.Body
	defer body.Close()

	// TODO: more code specific handling? like 429?
	// TODO: in case of a 404, this is more like a soft error
	// for the calling function; indicate that more nicely?
	statusCode := r.StatusCode
	if statusCode != http.StatusOK {
		result := responses.Error{Code: statusCode}
		json.NewDecoder(body).Decode(&result)
		return responses.Community{}, result
	}

	response := responses.Community{}
	json.NewDecoder(body).Decode(&response)

	return response, nil
}

func (c *Client) execute(method string, url string, body interface{}) (*http.Response, error) {

	relativeURL, err := urllib.Parse(url)
	if err != nil {
		return nil, err
	}
	requestURL := c.baseURL.ResolveReference(relativeURL)

	// TODO: handle body

	request, err := http.NewRequest(method, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", c.acceptHeader)
	request.Header.Add("Content-Type", c.contentTypeHeader)
	request.Header.Add("User-Agent", c.userAgent)
	request.Header.Add("key", c.key)

	return c.httpClient.Do(request)
}
