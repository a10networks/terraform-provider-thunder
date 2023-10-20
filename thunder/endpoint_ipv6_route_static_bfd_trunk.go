package thunder

import (
	"bytes"
	"github.com/clarketm/json"
	"io/ioutil"
	"strconv"
	"util"
)

// based on ACOS 5_2_1-P3_70
type EndpointIpv6RouteStaticBfdTrunk struct {
	Inst Ipv6RouteStaticBfdTrunk `json:"trunk"`
}
type Ipv6RouteStaticBfdTrunk struct {
	Action        string `json:"action,omitempty"`
	TrunkNum      int    `json:"trunk-num,omitempty"`
	NexthopIpv6Ll string `json:"nexthop-ipv6-ll,omitempty"`
	Template      string `json:"template,omitempty"`
	Threshold     int    `json:"threshold,omitempty"`
	Uuid          string `json:"uuid,omitempty"`
}

func (p *EndpointIpv6RouteStaticBfdTrunk) GetId() string {
	return strconv.Itoa(p.Inst.TrunkNum) + "+" + p.Inst.NexthopIpv6Ll
}

func PostEndpointIpv6RouteStaticBfdTrunk(auth_token string, inst EndpointIpv6RouteStaticBfdTrunk, host string) error {
	logger := util.GetLoggerInstance()
	logger.Println("PostIpv6RouteStaticBfdTrunk")
	headers := GetReqHeader(auth_token)
	payloadBytes, err := json.Marshal(inst)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	resp, err := doHttp("POST", "https://"+host+"/axapi/v3/ipv6/route/static/bfd/trunk", bytes.NewReader(payloadBytes), headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", string(data))
		var m EndpointIpv6RouteStaticBfdTrunk
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("PostIpv6RouteStaticBfdTrunk", data)
		}
	}
	return err
}

func GetEndpointIpv6RouteStaticBfdTrunk(auth_token string, host string, instId string) (*EndpointIpv6RouteStaticBfdTrunk, error) {
	logger := util.GetLoggerInstance()
	logger.Println("GetIpv6RouteStaticBfdTrunk")
	headers := GetReqHeader(auth_token)
	resp, err := doHttp("GET", "https://"+host+"/axapi/v3/ipv6/route/static/bfd/trunk/"+instId, nil, headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", string(data))
		var m EndpointIpv6RouteStaticBfdTrunk
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("GetIpv6RouteStaticBfdTrunk", data)
			if err == nil {
				return &m, nil
			}
		}
	}
	return nil, err
}

func PutEndpointIpv6RouteStaticBfdTrunk(auth_token string, inst EndpointIpv6RouteStaticBfdTrunk, host string) error {
	logger := util.GetLoggerInstance()
	logger.Println("PutIpv6RouteStaticBfdTrunk")
	headers := GetReqHeader(auth_token)
	payloadBytes, err := json.Marshal(inst)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))

	resp, err := doHttp("PUT", "https://"+host+"/axapi/v3/ipv6/route/static/bfd/trunk/"+inst.GetId(), bytes.NewReader(payloadBytes), headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", string(data))
		var m EndpointIpv6RouteStaticBfdTrunk
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("PutIpv6RouteStaticBfdTrunk", data)
		}
	}
	return err
}

func DeleteEndpointIpv6RouteStaticBfdTrunk(auth_token string, host string, instId string) error {
	logger := util.GetLoggerInstance()
	logger.Println("DeleteIpv6RouteStaticBfdTrunk")
	headers := GetReqHeader(auth_token)
	resp, err := doHttp("DELETE", "https://"+host+"/axapi/v3/ipv6/route/static/bfd/trunk/"+instId, nil, headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", data)
		var m EndpointIpv6RouteStaticBfdTrunk
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("DeleteIpv6RouteStaticBfdTrunk", data)
		}
	}
	return err
}
