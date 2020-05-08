package xiaomirouterapi

type countryCodeResp struct {
	Code    int           `json:"code"`
	Current string        `json:"current"`
	List    []countryCode `json:"list"`
}

type countryCode struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (api *MiWifiApi) CountryCode() (*countryCodeResp, error) {
	apiURL := api.buildApiURL("xqsystem/country_code", "api")
	resp := &countryCodeResp{}

	err := sendGetRequest(apiURL, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
