package xiaomirouterapi

import (
	"net/http"
	"net/url"
)

const (
	authURL = "xqsystem/login"
)

type authResponse struct {
	Token string `json:"token"`
}

func (api *MiWifiApi) Auth(user string, pass string) error {
	apiUrl := api.buildURL(authURL)
	nonce := genNonce()
	passHash := genPasswordHash(nonce, pass)

	req := url.Values{}
	req.Add("username", user)
	req.Add("logtype", "2")
	req.Add("password", passHash)
	req.Add("nonce", nonce)

	resp := authResponse{}

	err := sendPostRequest(apiUrl, req, &resp)
	if err != nil {
		return err
	}

	api.token = &resp.Token

	return nil
}

func (api *MiWifiApi) Logout() error {
	apiURL := api.buildApiURL("logout", "web")
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()

	return nil
}
