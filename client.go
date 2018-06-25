package insteon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client to access a TeamCity API
type Client struct {
	HTTPClient    *http.Client
	baseURL       string
	authorization *Authorization
}

func New(baseURL string) *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		baseURL:    strings.TrimRight(baseURL, "/"),
	}
}

// Authenticate will authenticate user to Insteon REST API
func (c *Client) Authenticate(clientID string, username string, password string) error {
	v := url.Values{}
	v.Set("grant_type", "password")
	v.Set("client_id", clientID)
	v.Set("username", username)
	v.Set("password", password)
	data := v.Encode()

	var b BearerResponse
	req, _ := http.NewRequest("POST", c.baseURL+"/oauth2/token", bytes.NewBuffer([]byte(data)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	j, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(j, &b)
	if err != nil {
		return err
	}

	c.authorization = &Authorization{
		ClientID:     clientID,
		AccessToken:  b.Access_Token,
		RefreshToken: b.Refresh_Token,
		ExpiresIn:    b.Expires_In,
	}

	return nil
}

// RefreshToken will refresh the authorization token used to authenticate with Insteon REST API
func (c *Client) RefreshToken(refreshToken string) error {
	if c.authorization == nil {
		return errors.New("must authenticate first")
	}

	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("client_id", c.authorization.ClientID)
	v.Set("refresh_token", refreshToken)
	data := v.Encode()

	var b BearerResponse
	req, _ := http.NewRequest("POST", c.baseURL+"/oauth2/token", bytes.NewBuffer([]byte(data)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	j, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(j, &b)
	if err != nil {
		return err
	}
	c.authorization.AccessToken = b.Access_Token
	c.authorization.RefreshToken = b.Refresh_Token
	c.authorization.ExpiresIn = b.Expires_In
	return nil
}

// CommandStatus returns the result of a command posted to Insteon REST API
func (c *Client) CommandStatus(commandID int) (CommandResponse, error) {
	var cmdResponse CommandResponse
	path := fmt.Sprintf("%s/commands/%d", c.baseURL, commandID)

	retries := 2
	err := withRetry(retries, func() error {
		return c.doRequest("GET", path, nil, &cmdResponse)
	})
	return cmdResponse, err
}

// SendCommand will POST an command to Insteon /commands REST API
func (c *Client) SendCommand(command string, deviceID int) (CommandStatusResponse, error) {
	cmd := Command{
		Command:   command,
		Device_Id: deviceID,
	}
	var cmdStatusResp CommandStatusResponse
	path := c.baseURL + "/commands"
	retries := 2
	err := withRetry(retries, func() error {
		return c.doRequest("POST", path, cmd, &cmdStatusResp)
	})
	return cmdStatusResp, err
}

func (c *Client) doRequest(method string, path string, data interface{}, v interface{}) error {
	jsonCnt, err := c.doJSONRequest(method, path, data)
	if err != nil {
		return err
	}

	if v != nil {
		err = json.Unmarshal(jsonCnt, &v)
		if err != nil {
			return fmt.Errorf("json unmarshal: %s (%q)", err, string(jsonCnt))
		}
	}
	return nil
}

func (c *Client) doJSONRequest(method string, path string, data interface{}) ([]byte, error) {
	// if !strings.HasPrefix(path, "/") {
	// 	path = "/" + path
	// }

	var body io.Reader
	if data != nil {
		jsonReq, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("marshaling data: %s", err)
		}
		body = bytes.NewBuffer(jsonReq)
	}

	req, _ := http.NewRequest(method, path, body)
	req.Header.Add("Authentication", "APIKey "+c.authorization.ClientID)
	req.Header.Add("Authorization", "Bearer "+c.authorization.AccessToken)

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func withRetry(retries int, f func() error) (err error) {
	for i := 0; i < retries; i++ {
		err = f()
		if err != nil {
			log.Printf("Retry: %v / %v, error: %v\n", i, retries, err)
		} else {
			return
		}
	}
	return
}
