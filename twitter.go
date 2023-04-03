package twitter

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/gomodule/oauth1/oauth"
)

func New(apiKey, apiSecret string) *Client {
	cli := &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  apiKey,
			Secret: apiSecret,
		},
	}
	return &Client{
		appID:     apiKey,
		appSecret: apiSecret,
		cli:       cli,
	}
}

type Client struct {
	appID     string
	appSecret string
	cli       *oauth.Client
	tempCred  *oauth.Credentials
	cred      *oauth.Credentials
}

func (c *Client) GetAuthURL(callback string) (string, error) {
	tempCred, err := c.cli.RequestTemporaryCredentials(nil, callback, nil)
	if err != nil {
		return "", err
	}
	c.tempCred = tempCred
	url := c.cli.AuthorizationURL(tempCred, nil)
	return url, nil
}

func (c *Client) GetAccessToken(oauthToken string, oauthVerifier string) (string, error) {
	if c.tempCred == nil || c.tempCred.Token != oauthToken {
		return "", ErrInvalidRequest
	}
	cred, _, err := c.cli.RequestToken(nil, c.tempCred, oauthVerifier)
	c.cred = cred
	return cred.Token, err
}

func (c *Client) Verify() (*User, error) {
	v := url.Values{}
	v.Set("include_email", "false")
	resp, err := c.cli.Get(nil, c.cred, "https://api.twitter.com/1.1/account/verify_credentials.json", v)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 500 || resp.StatusCode >= 400 {
		terr := ErrorWrapper{}
		if err := json.Unmarshal(respData, &terr); err != nil {
			return nil, err
		}

		if len(terr.Errors) != 0 {
			return nil, &terr
		}

		return nil, ErrTwitterServerError
	}

	user := User{}
	if err := json.Unmarshal(respData, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
