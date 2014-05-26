package spark

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

type TokensService struct {
	client *SparkClient
}

type Token struct {
	Token     *string `json:"token"`
	ExpiresAt *string `json:"expires_at"`
	Client    *string `json:"client"`
}

type AccessToken struct {
	Value     *string `json:"access_token,omitempty"`
	Type      *string `json:"token_type,omitempty"`
	ExpiresIn *uint32 `json:"expires_in,omitempty"`
}

func (a AccessToken) String() string {
	return fmt.Sprintf("AccessToken{%v, %v, %v}", *a.Value, *a.Type, *a.ExpiresIn)
}

func (t Token) String() string {
	return fmt.Sprintf("Token{token: %v, expires_at: %v, client: %v}", *t.Token, *t.ExpiresAt, *t.Client)
}

func (s *TokensService) Login(username, password string) (AccessToken, *http.Response, error) {
	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	data.Set("grant_type", "password")
	data.Set("client_id", "github.com/pims/spark")
	data.Set("client_secret", "client_secret_here")

	body := bytes.NewBufferString(data.Encode())

	req, err := s.client.NewRequest("POST", "oauth/token", body)
	token := new(AccessToken)
	resp, err := s.client.Do(req, token)
	if err != nil {
		return AccessToken{}, resp, err
	}

	return *token, resp, err
}

func (s *TokensService) List(username, password string) ([]Token, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "v1/access_tokens", nil)
	if err != nil {
		return nil, nil, err
	}
	req.SetBasicAuth(username, password)
	tokens := new([]Token)
	resp, err := s.client.Do(req, tokens)
	if err != nil {
		return nil, resp, err
	}

	return *tokens, resp, err
}

func (s *TokensService) Delete(token, username, password string) (*http.Response, error) {
	path := fmt.Sprintf("v1/access_tokens/%s", token)
	req, err := s.client.NewRequest("DELETE", path, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, password)

	resp, err := s.client.Do(req, nil)
	return resp, err
}
