package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type HttpPolicy struct {
	Name HttpPolicyInstance `json:"http-policy,omitempty"`
}
type HTTPPolicyMatch struct {
	Type             string `json:"type,omitempty"`
	MatchString      string `json:"match-string,omitempty"`
	TemplateName     string `json:"template-name,omitempty"`
	ServiceGroup     string `json:"service-group,omitempty"`
	URLUnderHost     string `json:"url-under-host,omitempty"`
	OtherMatchString string `json:"other-match-string,omitempty"`
	Template         string `json:"template,omitempty"`
	OtherMatchType   string `json:"other-match-type,omitempty"`
	MatchType        string `json:"match-type,omitempty"`
}
type GeoLocationMatch struct {
	GeoLocation             string `json:"geo-location,omitempty"`
	GeoLocationServiceGroup string `json:"geo-location-service-group,omitempty"`
	GeoLocationTemplate     string `json:"geo-location-template,omitempty"`
	GeoLocationTemplateName string `json:"geo-location-template-name,omitempty"`
}
type HttpPolicyInstance struct {
	CookieName  string             `json:"cookie-name,omitempty"`
	Name        string             `json:"name,omitempty"`
	MatchString []HTTPPolicyMatch  `json:"http-policy-match,omitempty"`
	UUID        string             `json:"uuid,omitempty"`
	UserTag     string             `json:"user-tag,omitempty"`
	GeoLocation []GeoLocationMatch `json:"geo-location-match,omitempty"`
}

func GetHttpPolicy(id string, name string, host string) (*HttpPolicy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/http-policy/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HttpPolicy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GetHttpPolicy REQ RES..........................", m)
			err := check_api_status("GetHttpPolicy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostHttpPolicy(id string, sg HttpPolicy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/http-policy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HttpPolicy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostHttpPolicy", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func PutHttpPolicy(id string, name string, sg HttpPolicy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/http-policy/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HttpPolicy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PutHttpPolicy", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func DeleteHttpPolicy(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/http-policy/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HttpPolicy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
		}
	}
	return nil
}
