package xiaomirouterapi

type channelInfo struct {
	BandWidth int      `json:"bandwidth,string"`
	BandList  []string `json:"bandlist"`
	Channel   int      `json:"channel"`
}

type wifiDetailAllInfo struct {
	IfName        string      `json:"ifname"`
	ChannelInfo   channelInfo `json:"channelinfo"`
	Encryption    string      `json:"encryption"`
	BandWidth     int         `json:"bandwidth,string"`
	KickThreshold int         `json:"kickthreshold,string"`
	StatusRaw     interface{} `json:"status,string"`
	Status        int64       `json:"-"`
	Mode          string      `json:"mode"`
	BSD           int         `json:"bsd,string"`
	TXBF          int         `json:"txbf,string"`
	TXPwr         string      `json:"txpwr"`
	WeakThreshold int         `json:"weakthreshold,string"`
	Device        string      `json:"device"`
	Hidden        int         `json:"hidden,string"`
	Password      string      `json:"password"`
	Channel       int         `json:"channel,string"`
	WeakEnable    int         `json:"weakenable,string"`
	SSID          string      `json:"ssid"`
	Signal        int         `json:"signal"`
}

type wifiDetailAllResp struct {
	Code int                 `json:"code"`
	Bsd  int                 `json:"bsd"`
	Info []wifiDetailAllInfo `json:"info"`
}

func (api *MiWifiApi) WifiDetailAll() (*wifiDetailAllResp, error) {
	apiURL := api.buildApiURL("xqnetwork/wifi_detail_all", "api")
	resp := &wifiDetailAllResp{}

	err := sendGetRequest(apiURL, &resp)
	if err != nil {
		return nil, err
	}

	postProcessWifiDetailAllResp(resp)

	return resp, nil
}

func postProcessWifiDetailAllResp(wr *wifiDetailAllResp) {
	for i := range wr.Info {
		el := &wr.Info[i]

		el.Status = interfaceToInt64(el.StatusRaw)
	}
}
