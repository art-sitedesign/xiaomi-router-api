package xiaomirouterapi

import "fmt"

type BandWidthTest struct {
	Code       int     `json:"code"`
	Download   float64 `json:"download"`
	Upload     float64 `json:"upload"`
	BandWidth  float64 `json:"bandwidth"`
	BandWidth2 float64 `json:"bandwidth2"`
}

func (api *MiWifiApi) BandWidthTest(history int) (*BandWidthTest, error) {
	path := "misystem/bandwidth_test"
	if history > 0 {
		path += fmt.Sprintf("?history=%d", history)
	}

	apiURL := api.buildApiURL(path, "api")
	resp := &BandWidthTest{}

	err := sendGetRequest(apiURL, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
