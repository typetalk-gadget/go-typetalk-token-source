package source

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/vvatanabe/expiremap"
	"golang.org/x/oauth2"
)

const (
	accessTokenKey  = "access-token"
	refreshTokenKey = "refresh-token"
)

type TokenSource struct {
	ClientID     string
	ClientSecret string
	Scope        string
	store        expiremap.Map
}

func (s *TokenSource) Token() (*oauth2.Token, error) {

	if v, ok := s.store.Load(accessTokenKey); ok {
		return v.(*oauth2.Token), nil
	}

	var (
		t   *accessToken
		err error
	)
	if v, ok := s.store.Load(refreshTokenKey); ok {
		t, err = refreshAccessToken(s.ClientID, s.ClientSecret, v.(string))
	} else {
		t, err = getAccessToken(s.ClientID, s.ClientSecret, s.Scope)
	}
	if err != nil {
		return nil, err
	}

	tokenExpire := time.Duration(t.ExpiresIn-10) * time.Second
	refreshExpire := 23*time.Hour + 59*time.Minute

	token := &oauth2.Token{
		AccessToken:  t.AccessToken,
		TokenType:    t.TokenType,
		RefreshToken: t.RefreshToken,
		Expiry:       time.Now().Add(tokenExpire),
	}

	s.store.Store(accessTokenKey, token, expiremap.Expire(tokenExpire))
	s.store.Store(refreshTokenKey, token.RefreshToken, expiremap.Expire(refreshExpire))

	return token, nil
}

func getAccessToken(clientID, clientSecret, scope string) (*accessToken, error) {
	form := make(url.Values)
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)
	form.Add("grant_type", "client_credentials")
	form.Add("scope", scope)
	oauth2resp, err := http.PostForm("https://typetalk.com/oauth2/access_token", form)
	if err != nil {
		return nil, fmt.Errorf("credential request error :%w", err)
	}
	var accessToken accessToken
	err = json.NewDecoder(oauth2resp.Body).Decode(&accessToken)
	if err != nil {
		return nil, fmt.Errorf("credential decode error :%w", err)
	}
	return &accessToken, nil
}

func refreshAccessToken(clientID, clientSecret, refreshToken string) (*accessToken, error) {
	form := make(url.Values)
	form.Add("client_id", clientID)
	form.Add("client_secret", clientSecret)
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", refreshToken)
	oauth2resp, err := http.PostForm("https://typetalk.com/oauth2/access_token", form)
	if err != nil {
		return nil, fmt.Errorf("refresh accessToken error :%w", err)
	}
	var accessToken accessToken
	err = json.NewDecoder(oauth2resp.Body).Decode(&accessToken)
	if err != nil {
		return nil, fmt.Errorf("credential decode error :%w", err)
	}
	return &accessToken, nil
}

type accessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
