package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterOspfRedistribute struct {
	RedistList RouterOspfRedistributeInstance `json:"redistribute,omitempty"`
}

type RouterOspfRedistributeInstance struct {
	IPNat           int                                       `json:"ip-nat,omitempty"`
	IPNatPrefix     []RouterOspfRedistributeIPNatFloatingList `json:"ip-nat-floating-list,omitempty"`
	MetricIPNat     int                                       `json:"metric-ip-nat,omitempty"`
	MetricTypeIPNat string                                    `json:"metric-type-ip-nat,omitempty"`
	Ospf            []RouterOspfRedistributeOspfList          `json:"ospf-list,omitempty"`
	Type            []RouterOspfRedistributeRedistList        `json:"redist-list,omitempty"`
	RouteMapIPNat   string                                    `json:"route-map-ip-nat,omitempty"`
	TagIPNat        int                                       `json:"tag-ip-nat,omitempty"`
	UUID            string                                    `json:"uuid,omitempty"`
	VipAddress      []RouterOspfRedistributeVipFloatingList   `json:"vip-floating-list,omitempty"`
	TypeVip         []RouterOspfRedistributeVipList           `json:"vip-list,omitempty"`
}

type RouterOspfRedistributeIPNatFloatingList struct {
	IPNatFloatingIPForward string `json:"ip-nat-floating-IP-forward,omitempty"`
	IPNatPrefix            string `json:"ip-nat-prefix,omitempty"`
}

type RouterOspfRedistributeOspfList struct {
	MetricOspf     int    `json:"metric-ospf,omitempty"`
	MetricTypeOspf string `json:"metric-type-ospf,omitempty"`
	Ospf           int    `json:"ospf,omitempty"`
	ProcessID      int    `json:"process-id,omitempty"`
	RouteMapOspf   string `json:"route-map-ospf,omitempty"`
	TagOspf        int    `json:"tag-ospf,omitempty"`
}

type RouterOspfRedistributeRedistList struct {
	Metric     int    `json:"metric,omitempty"`
	MetricType string `json:"metric-type,omitempty"`
	RouteMap   string `json:"route-map,omitempty"`
	Tag        int    `json:"tag,omitempty"`
	Type       string `json:"type,omitempty"`
}

type RouterOspfRedistributeVipFloatingList struct {
	VipAddress           string `json:"vip-address,omitempty"`
	VipFloatingIPForward string `json:"vip-floating-IP-forward,omitempty"`
}

type RouterOspfRedistributeVipList struct {
	MetricTypeVip string `json:"metric-type-vip,omitempty"`
	MetricVip     int    `json:"metric-vip,omitempty"`
	RouteMapVip   string `json:"route-map-vip,omitempty"`
	TagVip        int    `json:"tag-vip,omitempty"`
	TypeVip       string `json:"type-vip,omitempty"`
}

func PostRouterOspfRedistribute(id string, name1 string, inst RouterOspfRedistribute, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterOspfRedistribute")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/ospf/"+name1+"/redistribute", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspfRedistribute
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterOspfRedistribute", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterOspfRedistribute(id string, name1 string, host string) (*RouterOspfRedistribute, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterOspfRedistribute")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/ospf/"+name1+"/redistribute", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspfRedistribute
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterOspfRedistribute", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
