package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpAddressFamilyIpv6Redistribute struct {
	ConnectedCfg RouterBgpAddressFamilyIpv6RedistributeInstance `json:"redistribute,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstance struct {
	Connected      RouterBgpAddressFamilyIpv6RedistributeConnectedCfg  `json:"connected-cfg,omitempty"`
	FloatingIP     RouterBgpAddressFamilyIpv6RedistributeFloatingIPCfg `json:"floating-ip-cfg,omitempty"`
	IPNat          RouterBgpAddressFamilyIpv6RedistributeIPNatCfg      `json:"ip-nat-cfg,omitempty"`
	IPNatList      RouterBgpAddressFamilyIpv6RedistributeIPNatListCfg  `json:"ip-nat-list-cfg,omitempty"`
	Isis           RouterBgpAddressFamilyIpv6RedistributeIsisCfg       `json:"isis-cfg,omitempty"`
	Lw4O6          RouterBgpAddressFamilyIpv6RedistributeLw4O6Cfg      `json:"lw4o6-cfg,omitempty"`
	Nat64          RouterBgpAddressFamilyIpv6RedistributeNat64Cfg      `json:"nat64-cfg,omitempty"`
	NatMap         RouterBgpAddressFamilyIpv6RedistributeNatMapCfg     `json:"nat-map-cfg,omitempty"`
	Ospf           RouterBgpAddressFamilyIpv6RedistributeOspfCfg       `json:"ospf-cfg,omitempty"`
	Rip            RouterBgpAddressFamilyIpv6RedistributeRipCfg        `json:"rip-cfg,omitempty"`
	Static         RouterBgpAddressFamilyIpv6RedistributeStaticCfg     `json:"static-cfg,omitempty"`
	StaticNat      RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg  `json:"static-nat-cfg,omitempty"`
	UUID           string                                              `json:"uuid,omitempty"`
	OnlyFlaggedCfg RouterBgpAddressFamilyIpv6RedistributeVip           `json:"vip,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeConnectedCfg struct {
	Connected int    `json:"connected,omitempty"`
	RouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeFloatingIPCfg struct {
	FloatingIP int    `json:"floating-ip,omitempty"`
	RouteMap   string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeIPNatCfg struct {
	IPNat    int    `json:"ip-nat,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeIPNatListCfg struct {
	IPNatList int    `json:"ip-nat-list,omitempty"`
	RouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeIsisCfg struct {
	Isis     int    `json:"isis,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeLw4O6Cfg struct {
	Lw4O6    int    `json:"lw4o6,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeNat64Cfg struct {
	Nat64    int    `json:"nat64,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeNatMapCfg struct {
	NatMap   int    `json:"nat-map,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeOspfCfg struct {
	Ospf     int    `json:"ospf,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeRipCfg struct {
	Rip      int    `json:"rip,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeStaticCfg struct {
	RouteMap string `json:"route-map,omitempty"`
	Static   int    `json:"static,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg struct {
	RouteMap  string `json:"route-map,omitempty"`
	StaticNat int    `json:"static-nat,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeVip struct {
	OnlyFlagged    RouterBgpAddressFamilyIpv6RedistributeOnlyFlaggedCfg    `json:"only-flagged-cfg,omitempty"`
	OnlyNotFlagged RouterBgpAddressFamilyIpv6RedistributeOnlyNotFlaggedCfg `json:"only-not-flagged-cfg,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeOnlyFlaggedCfg struct {
	OnlyFlagged int    `json:"only-flagged,omitempty"`
	RouteMap    string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeOnlyNotFlaggedCfg struct {
	OnlyNotFlagged int    `json:"only-not-flagged,omitempty"`
	RouteMap       string `json:"route-map,omitempty"`
}

func PostRouterBgpAddressFamilyIpv6Redistribute(id string, name1 string, inst RouterBgpAddressFamilyIpv6Redistribute, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpAddressFamilyIpv6Redistribute")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/redistribute", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6Redistribute
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpAddressFamilyIpv6Redistribute", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func GetRouterBgpAddressFamilyIpv6Redistribute(id string, name1 string, host string) (*RouterBgpAddressFamilyIpv6Redistribute, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpAddressFamilyIpv6Redistribute")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/redistribute", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6Redistribute
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpAddressFamilyIpv6Redistribute", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
