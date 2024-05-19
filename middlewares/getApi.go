package middlewares

import (
	"encoding/json"
	"goserver/database"
	"io/ioutil"
	"log"
	"net/http"
)

func GetApi() string {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Printf("Unable to grab the data from the Api Endpoint:%w", err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var ipResp database.IPResponse
	err = json.Unmarshal(data, &ipResp)
	if err != nil {
		return ""
	}
	return ipResp.IP
}
