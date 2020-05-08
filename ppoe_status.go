package xiaomirouterapi

type ppoeStatusIp struct {
	Mask    string `json:"mask"`
	Address string `json:"address"`
}

type PpoeStatus struct {
	Code   int          `json:"code"`
	Status int          `json:"status"`
	Proto  string       `json:"proto"`
	Dns    []string     `json:"dns"`
	Gw     string       `json:"gw"`
	Ip     ppoeStatusIp `json:"ip"`
}

func (api *MiWifiApi) PPOEStatus() (*PpoeStatus, error) {
	apiURL := api.buildApiURL("xqnetwork/pppoe_status", "api")
	resp := &PpoeStatus{}

	err := sendGetRequest(apiURL, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
