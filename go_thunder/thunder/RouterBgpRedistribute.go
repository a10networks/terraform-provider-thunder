package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpRedistribute struct {
	ConnectedCfg RouterBgpRedistributeInstance `json:"redistribute,omitempty"`
}

type RouterBgpRedistributeInstance struct {
	Connected      RouterBgpRedistributeConnectedCfg  `json:"connected-cfg,omitempty"`
	FloatingIP     RouterBgpRedistributeFloatingIPCfg `json:"floating-ip-cfg,omitempty"`
	IPNat          RouterBgpRedistributeIPNatCfg      `json:"ip-nat-cfg,omitempty"`
	IPNatList      RouterBgpRedistributeIPNatListCfg  `json:"ip-nat-list-cfg,omitempty"`
	Isis           RouterBgpRedistributeIsisCfg       `json:"isis-cfg,omitempty"`
	Lw4O6          RouterBgpRedistributeLw4O6Cfg      `json:"lw4o6-cfg,omitempty"`
	NatMap         RouterBgpRedistributeNatMapCfg     `json:"nat-map-cfg,omitempty"`
	Ospf           RouterBgpRedistributeOspfCfg       `json:"ospf-cfg,omitempty"`
	Rip            RouterBgpRedistributeRipCfg        `json:"rip-cfg,omitempty"`
	Static         RouterBgpRedistributeStaticCfg     `json:"static-cfg,omitempty"`
	StaticNat      RouterBgpRedistributeStaticNatCfg  `json:"static-nat-cfg,omitempty"`
	UUID           string                             `json:"uuid,omitempty"`
	OnlyFlaggedCfg RouterBgpRedistributeVip           `json:"vip,omitempty"`
}

type RouterBgpRedistributeConnectedCfg struct {
	Connected int    `json:"connected,omitempty"`
	RouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeFloatingIPCfg struct {
	FloatingIP int    `json:"floating-ip,omitempty"`
	RouteMap   string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeIPNatCfg struct {
	IPNat    int    `json:"ip-nat,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeIPNatListCfg struct {
	IPNatList int    `json:"ip-nat-list,omitempty"`
	RouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeIsisCfg struct {
	Isis     int    `json:"isis,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeLw4O6Cfg struct {
	Lw4O6    int    `json:"lw4o6,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeNatMapCfg struct {
	NatMap   int    `json:"nat-map,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeOspfCfg struct {
	Ospf     int    `json:"ospf,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeRipCfg struct {
	Rip      int    `json:"rip,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeStaticCfg struct {
	RouteMap string `json:"route-map,omitempty"`
	Static   int    `json:"static,omitempty"`
}

type RouterBgpRedistributeStaticNatCfg struct {
	RouteMap  string `json:"route-map,omitempty"`
	StaticNat int    `json:"static-nat,omitempty"`
}

type RouterBgpRedistributeVip struct {
	OnlyFlagged    RouterBgpRedistributeOnlyFlaggedCfg    `json:"only-flagged-cfg,omitempty"`
	OnlyNotFlagged RouterBgpRedistributeOnlyNotFlaggedCfg `json:"only-not-flagged-cfg,omitempty"`
}

type RouterBgpRedistributeOnlyFlaggedCfg struct {
	OnlyFlagged int    `json:"only-flagged,omitempty"`
	RouteMap    string `json:"route-map,omitempty"`
}

type RouterBgpRedistributeOnlyNotFlaggedCfg struct {
	OnlyNotFlagged int    `json:"only-not-flagged,omitempty"`
	RouteMap       string `json:"route-map,omitempty"`
}

func PostRouterBgpRedistribute(id string, name1 string, inst RouterBgpRedistribute, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpRedistribute")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/redistribute", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpRedistribute
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpRedistribute", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgpRedistribute(id string, name1 string, host string) (*RouterBgpRedistribute, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpRedistribute")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/redistribute", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpRedistribute
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpRedistribute", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
