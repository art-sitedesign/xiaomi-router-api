package xiaomirouterapi

import (
	"fmt"

	xiaomirouterapi "pet-projects/xiaomi-router-api/app"
)

func main() {
	api := xiaomirouterapi.NewMiWifiApi("192.168.31.1")

	err := api.Auth("admin", "FLUGEGEHEIMEN")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		_ = api.Logout()
	}()

	status, err := api.Status()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", status)

	//_ = api.DeviceList()
	//_, _ = api.BandWidthTest(0)
	//_, _ = api.PPOEStatus()
	//_, _ = api.WifiDetailAll()
	//_, _ = api.CountryCode()
}
