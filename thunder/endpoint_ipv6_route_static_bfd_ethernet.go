package thunder

import (
	"bytes"
	"github.com/clarketm/json"
	"io/ioutil"
	"net/url"
	"strconv"
	"util"
)

// based on ACOS 5_2_1-P3_70
type EndpointIpv6RouteStaticBfdEthernet struct {
	Inst Ipv6RouteStaticBfdEthernet `json:"ethernet"`
}
type Ipv6RouteStaticBfdEthernet struct {
	Action        string `json:"action,omitempty"`
	EthNum        int    `json:"eth-num,omitempty"`
	NexthopIpv6Ll string `json:"nexthop-ipv6-ll,omitempty"`
	Template      string `json:"template,omitempty"`
	Threshold     int    `json:"threshold,omitempty"`
	Uuid          string `json:"uuid,omitempty"`
}

func (p *EndpointIpv6RouteStaticBfdEthernet) GetId() string {
	return strconv.Itoa(p.Inst.EthNum) + "+" + url.QueryEscape(p.Inst.NexthopIpv6Ll)
}

func PostEndpointIpv6RouteStaticBfdEthernet(auth_token string, inst EndpointIpv6RouteStaticBfdEthernet, host string) error {
	logger := util.GetLoggerInstance()
	logger.Println("PostIpv6RouteStaticBfdEthernet")
	headers := GetReqHeader(auth_token)
	payloadBytes, err := json.Marshal(inst)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	resp, err := doHttp("POST", "https://"+host+"/axapi/v3/ipv6/route/static/bfd/ethernet", bytes.NewReader(payloadBytes), headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", string(data))
		var m EndpointIpv6RouteStaticBfdEthernet
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("PostIpv6RouteStaticBfdEthernet", data)
		}
	}
	return err
}

func GetEndpointIpv6RouteStaticBfdEthernet(auth_token string, host string, instId string) (*EndpointIpv6RouteStaticBfdEthernet, error) {
	logger := util.GetLoggerInstance()
	logger.Println("GetIpv6RouteStaticBfdEthernet")
	headers := GetReqHeader(auth_token)
	resp, err := doHttp("GET", "https://"+host+"/axapi/v3/ipv6/route/static/bfd/ethernet/"+instId, nil, headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", string(data))
		var m EndpointIpv6RouteStaticBfdEthernet
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("GetIpv6RouteStaticBfdEthernet", data)
			if err == nil {
				return &m, nil
			}
		}
	}
	return nil, err
}

func PutEndpointIpv6RouteStaticBfdEthernet(auth_token string, inst EndpointIpv6RouteStaticBfdEthernet, host string) error {
	logger := util.GetLoggerInstance()
	logger.Println("PutIpv6RouteStaticBfdEthernet")
	headers := GetReqHeader(auth_token)
	payloadBytes, err := json.Marshal(inst)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))

	resp, err := doHttp("PUT", "https://"+host+"/axapi/v3/ipv6/route/static/bfd/ethernet/"+inst.GetId(), bytes.NewReader(payloadBytes), headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", string(data))
		var m EndpointIpv6RouteStaticBfdEthernet
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("PutIpv6RouteStaticBfdEthernet", data)
		}
	}
	return err
}

func DeleteEndpointIpv6RouteStaticBfdEthernet(auth_token string, host string, instId string) error {
	logger := util.GetLoggerInstance()
	logger.Println("DeleteIpv6RouteStaticBfdEthernet")
	headers := GetReqHeader(auth_token)
	resp, err := doHttp("DELETE", "https://"+host+"/axapi/v3/ipv6/route/static/bfd/ethernet/"+instId, nil, headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", data)
		var m EndpointIpv6RouteStaticBfdEthernet
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("DeleteIpv6RouteStaticBfdEthernet", data)
		}
	}
	return err
}
