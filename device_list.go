package xiaomirouterapi

type authority struct {
	Admin   int `json:"admin"`
	Lan     int `json:"lan"`
	PriDisk int `json:"pridisk"`
	Wan     int `json:"wan"`
}

type deviceListIp struct {
	Active    int    `json:"active"`
	DownSpeed int64  `json:"downspeed,string"`
	Ip        string `json:"ip"`
	Online    int64  `json:"online,string"`
	UpSpeed   int64  `json:"upspeed,string"`
}

type statistics struct {
	DownSpeed int64 `json:"downspeed,string"`
	Online    int64 `json:"online,string"`
	UpSpeed   int64 `json:"upspeed,string"`
}

type list struct {
	Authority  authority      `json:"authority"`
	Icon       string         `json:"icon"`
	Ip         []deviceListIp `json:"ip"`
	Isap       int            `json:"isap"`
	Mac        string         `json:"mac"`
	Name       string         `json:"name"`
	OName      string         `json:"oname"`
	OnLine     int            `json:"online"`
	Parent     string         `json:"parent"`
	Push       int            `json:"push"`
	Statistics statistics     `json:"statistics"`
	Times      int            `json:"times"`
	Type       int            `json:"type"`
}

type DeviceListResp struct {
	Code int    `json:"code"`
	List []list `json:"list"`
	Mac  string `json:"mac"`
}

func (api *MiWifiApi) DeviceList() (*DeviceListResp, error) {
	apiURL := api.buildApiURL("misystem/devicelist", "api")
	resp := &DeviceListResp{}

	err := sendGetRequest(apiURL, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
