package vthunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go_vthunder/vthunder"
	"io/ioutil"
	"util"
)

func getAuthHeader(host string, user string, passwd string, uri string) string {

	logger := util.GetLoggerInstance()

	var id string

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type Payload struct {
		Credentials Credentials `json:"credentials"`
	}

	type Auth struct {
		Signature   string `json:"signature"`
		Description string `json:"description"`
	}

	type Res struct {
		Auth Auth `json:"authresponse"`
	}

	data := Payload{
		Credentials: Credentials{
			Username: user,
			Password: passwd,
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
	}
	body := bytes.NewReader(payloadBytes)

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	logger.Println("[INFO] host+uri - " + host + uri)
	resp, err := go_vthunder.DoHttp("POST", "https://"+host+uri, body, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("[INFO] IThe HTTP request failed with error ")
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Res
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			logger.Println("[INFO] Unmarshal error ")
		} else {
			id = m.Auth.Signature
		}
	}
	return "A10 " + id
}
