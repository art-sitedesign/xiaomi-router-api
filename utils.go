package xiaomirouterapi

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	typeID    = 0
	randMin   = 1000
	randMax   = 9999
	publicKey = "a2ffa5c9be07488bbb04a3a47d3c5f6a"
)

func genMacAddress() string {
	return "f0:18:98:3d:dc:91" //todo...
}

func genNonce() string {
	return fmt.Sprintf(
		"%d_%s_%d_%d",
		typeID,
		genMacAddress(),
		time.Now().Unix(),
		rand.Intn(randMax-randMin)+randMin,
	)
}

func genPasswordHash(nonce string, password string) string {
	hashPart := fmt.Sprintf("%x", sha1.Sum([]byte(password+publicKey)))
	fullHash := sha1.Sum([]byte(nonce + hashPart))

	return fmt.Sprintf("%x", fullHash)
}

func sendPostRequest(apiURL string, data url.Values, out interface{}) error {
	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("post request failed: %v\n%s", resp.Status, b)
	}

	err = json.Unmarshal(b, out)
	if err != nil {
		return err
	}

	return nil
}

func sendGetRequest(apiURL string, out interface{}) error {
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get request failed: %v\n%s", resp.Status, b)
	}

	err = json.Unmarshal(b, out)
	if err != nil {
		return err
	}

	return nil
}

func interfaceToInt64(i interface{}) int64 {
	var res int64

	switch i.(type) {
	case string:
		r, err := strconv.Atoi(i.(string))
		if err == nil {
			res = int64(r)
		}
	case float64:
		r := i.(float64)
		res = int64(r)
	}

	return res
}
