package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type RouteStaticBfd struct {
	LocalIP IPBfd `json:"bfd,omitempty"`
}
type IPBfd struct {
	LocalIP   string `json:"local-ip"`
	NexthopIP string `json:"nexthop-ip"`
	UUID      string `json:"uuid"`
}

func PostIPRouteStaticBfd(id string, name string, inst RouteStaticBfd, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIPRouteStaticBfd")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/route/static/bfd", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteStaticBfd
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIPRouteStaticBfd REQ RES..........................", m)
			err := check_api_status("PostIPRouteStaticBfd", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func GetIPRouteStaticBfd(id string, name string, host string) (*RouteStaticBfd, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIPRouteStaticBfd")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/route/static/bfd/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteStaticBfd
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIPRouteStaticBfd REQ RES..........................", m)
			err := check_api_status("GetIPRouteStaticBfd", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutIPRouteStaticBfd(id string, name string, inst RouteStaticBfd, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutIPRouteStaticBfd")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/ip/route/static/bfd/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteStaticBfd
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutIPRouteStaticBfd REQ RES..........................", m)
			err := check_api_status("PutIPRouteStaticBfd", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func DeleteIPRouteStaticBfd(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteIPRouteStaticBfd")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/ip/route/static/bfd/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteStaticBfd
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
