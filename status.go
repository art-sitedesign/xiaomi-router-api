package xiaomirouterapi

type count struct {
	All    int `json:"all"`
	Online int `json:"online"`
}

type cpu struct {
	Core int     `json:"core"`
	Hz   string  `json:"hz"`
	Load float64 `json:"load"`
}

type dev struct {
	DevName          string      `json:"devname"`
	DownloadRaw      interface{} `json:"download,string"`
	Download         int64       `json:"-"`
	DownSpeedRaw     interface{} `json:"downspeed"`
	DownSpeed        int64       `json:"-"`
	Mac              string      `json:"mac"`
	MaxDownloadSpeed int64       `json:"maxdownloadspeed,string"`
	MaxUploadSpeed   int64       `json:"maxuploadspeed,string"`
	Online           int64       `json:"online,string"`
	UploadRaw        interface{} `json:"upload,string"`
	Upload           int64       `json:"-"`
	UpSpeedRaw       interface{} `json:"upspeed"`
	UpSpeed          int64       `json:"-"`
}

type hardware struct {
	Channel  string `json:"channel"`
	Mac      string `json:"mac"`
	Platform string `json:"platform"`
	Sn       string `json:"sn"`
	Version  string `json:"version"`
}

type mem struct {
	Hz    string  `json:"hz"`
	Total string  `json:"total"`
	Type  string  `json:"type"`
	Usage float64 `json:"usage"`
}

type wan struct {
	DevName          string `json:"devname"`
	Download         int64  `json:"download,string"`
	DownSpeed        int64  `json:"downspeed,string"`
	History          string `json:"history"`
	MaxDownloadSpeed int64  `json:"maxdownloadspeed,string"`
	MaxUploadSpeed   int64  `json:"maxuploadspeed,string"`
	Upload           int64  `json:"upload,string"`
	UpSpeed          int64  `json:"upspeed,string"`
}

type statusResp struct {
	Code        int      `json:"code"`
	Count       count    `json:"count"`
	Cpu         cpu      `json:"cpu"`
	Dev         []dev    `json:"dev"`
	Hardware    hardware `json:"hardware"`
	Mem         mem      `json:"mem"`
	Temperature float64  `json:"temperature"`
	UpTime      float64  `json:"uptime,string"`
	Wan         wan      `json:"wan"`
}

func (api *MiWifiApi) Status() (*statusResp, error) {
	apiURL := api.buildApiURL("misystem/status", "api")
	resp := &statusResp{}

	err := sendGetRequest(apiURL, &resp)
	if err != nil {
		return nil, err
	}

	// так как в некоторых полях могут приходить и string и int приходится отдельно преобразовывать
	postProcessStatusResp(resp)

	return resp, nil
}

func postProcessStatusResp(sr *statusResp) {
	for i := range sr.Dev {
		el := &sr.Dev[i]

		el.Download = interfaceToInt64(el.DownloadRaw)
		el.DownSpeed = interfaceToInt64(el.DownSpeedRaw)
		el.Upload = interfaceToInt64(el.UploadRaw)
		el.UpSpeed = interfaceToInt64(el.UpSpeedRaw)
	}
}
