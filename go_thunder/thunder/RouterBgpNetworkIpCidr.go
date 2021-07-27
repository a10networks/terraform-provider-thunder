package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpNetworkIpCidr struct {
	NetworkIpv4Cidr RouterBgpNetworkIPCidrInstance `json:"ip-cidr,omitempty"`
}

type RouterBgpNetworkIPCidrInstance struct {
	Backdoor        int    `json:"backdoor,omitempty"`
	CommValue       string `json:"comm-value,omitempty"`
	Description     string `json:"description,omitempty"`
	NetworkIpv4Cidr string `json:"network-ipv4-cidr,omitempty"`
	RouteMap        string `json:"route-map,omitempty"`
	UUID            string `json:"uuid,omitempty"`
}

func PostRouterBgpNetworkIpCidr(id string, name1 string, inst RouterBgpNetworkIpCidr, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpNetworkIpCidr")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/network/ip-cidr", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNetworkIpCidr
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpNetworkIpCidr", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgpNetworkIpCidr(id string, name1 string, name2 string, host string) (*RouterBgpNetworkIpCidr, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpNetworkIpCidr")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/network/ip-cidr/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNetworkIpCidr
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpNetworkIpCidr", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterBgpNetworkIpCidr(id string, name1 string, name2 string, inst RouterBgpNetworkIpCidr, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterBgpNetworkIpCidr")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/network/ip-cidr/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNetworkIpCidr
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterBgpNetworkIpCidr", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteRouterBgpNetworkIpCidr(id string, name1 string, name2 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteRouterBgpNetworkIpCidr")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/network/ip-cidr/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNetworkIpCidr
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteRouterBgpNetworkIpCidr", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
