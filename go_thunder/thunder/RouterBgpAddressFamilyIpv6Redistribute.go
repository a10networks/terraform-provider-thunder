package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpAddressFamilyIpv6Redistribute struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfg RouterBgpAddressFamilyIpv6RedistributeInstance `json:"redistribute,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstance struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfgConnected   RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfg  `json:"connected-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceFloatingIPCfgFloatingIP RouterBgpAddressFamilyIpv6RedistributeInstanceFloatingIPCfg `json:"floating-ip-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatCfgIPNat           RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatCfg      `json:"ip-nat-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatListCfgIPNatList   RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatListCfg  `json:"ip-nat-list-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceIsisCfgIsis             RouterBgpAddressFamilyIpv6RedistributeInstanceIsisCfg       `json:"isis-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceLw4O6CfgLw4O6           RouterBgpAddressFamilyIpv6RedistributeInstanceLw4O6Cfg      `json:"lw4o6-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceNat64CfgNat64           RouterBgpAddressFamilyIpv6RedistributeInstanceNat64Cfg      `json:"nat64-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceNatMapCfgNatMap         RouterBgpAddressFamilyIpv6RedistributeInstanceNatMapCfg     `json:"nat-map-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceOspfCfgOspf             RouterBgpAddressFamilyIpv6RedistributeInstanceOspfCfg       `json:"ospf-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceRipCfgRip               RouterBgpAddressFamilyIpv6RedistributeInstanceRipCfg        `json:"rip-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceStaticCfgStatic         RouterBgpAddressFamilyIpv6RedistributeInstanceStaticCfg     `json:"static-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceStaticNatCfgStaticNat   RouterBgpAddressFamilyIpv6RedistributeInstanceStaticNatCfg  `json:"static-nat-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceUUID                    string                                                      `json:"uuid,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfg       RouterBgpAddressFamilyIpv6RedistributeInstanceVip           `json:"vip,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfgConnected int    `json:"connected,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfgRouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceFloatingIPCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceFloatingIPCfgFloatingIP int    `json:"floating-ip,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceFloatingIPCfgRouteMap   string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatCfgIPNat    int    `json:"ip-nat,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatListCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatListCfgIPNatList int    `json:"ip-nat-list,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatListCfgRouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceIsisCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceIsisCfgIsis     int    `json:"isis,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceIsisCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceLw4O6Cfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceLw4O6CfgLw4O6    int    `json:"lw4o6,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceLw4O6CfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceNat64Cfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceNat64CfgNat64    int    `json:"nat64,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceNat64CfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceNatMapCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceNatMapCfgNatMap   int    `json:"nat-map,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceNatMapCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceOspfCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceOspfCfgOspf     int    `json:"ospf,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceOspfCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceRipCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceRipCfgRip      int    `json:"rip,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceRipCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceStaticCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceStaticCfgRouteMap string `json:"route-map,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceStaticCfgStatic   int    `json:"static,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceStaticNatCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceStaticNatCfgRouteMap  string `json:"route-map,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceStaticNatCfgStaticNat int    `json:"static-nat,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceVip struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfgOnlyFlagged       RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfg    `json:"only-flagged-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyNotFlaggedCfgOnlyNotFlagged RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyNotFlaggedCfg `json:"only-not-flagged-cfg,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfgOnlyFlagged int    `json:"only-flagged,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfgRouteMap    string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyNotFlaggedCfg struct {
	RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyNotFlaggedCfgOnlyNotFlagged int    `json:"only-not-flagged,omitempty"`
	RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyNotFlaggedCfgRouteMap       string `json:"route-map,omitempty"`
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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/redistribute", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6Redistribute
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
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
		err := json.Unmarshal(data, &m)
		if err != nil {
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
