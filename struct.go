package xiaomirouterapi

import "fmt"

type MiWifiApi struct {
	routerIP string
	token    *string
}

func NewMiWifiApi(routerIP string) *MiWifiApi {
	return &MiWifiApi{
		routerIP: routerIP,
	}
}

func (api *MiWifiApi) buildURL(path string) string {
	return fmt.Sprintf("http://%s/cgi-bin/luci/api/%s", api.routerIP, path)
}

func (api *MiWifiApi) buildApiURL(path string, apiType string) string {
	return fmt.Sprintf("http://%s/cgi-bin/luci/;stok=%s/%s/%s", api.routerIP, *api.token, apiType, path)
}

func (api *MiWifiApi) checkToken() bool {
	return api.token != nil
}
