package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceEthernetIpv6 struct {
	InterfaceEthernetIpv6InstanceAddressList InterfaceEthernetIpv6Instance `json:"ipv6,omitempty"`
}

type InterfaceEthernetIpv6Instance struct {
	InterfaceEthernetIpv6InstanceAccessListCfgV6AclName InterfaceEthernetIpv6InstanceAccessListCfg    `json:"access-list-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceAddressListIpv6Addr    []InterfaceEthernetIpv6InstanceAddressList    `json:"address-list,omitempty"`
	InterfaceEthernetIpv6InstanceInside                 int                                           `json:"inside,omitempty"`
	InterfaceEthernetIpv6InstanceIpv6Enable             int                                           `json:"ipv6-enable,omitempty"`
	InterfaceEthernetIpv6InstanceOspfNetworkList        InterfaceEthernetIpv6InstanceOspf             `json:"ospf,omitempty"`
	InterfaceEthernetIpv6InstanceOutside                int                                           `json:"outside,omitempty"`
	InterfaceEthernetIpv6InstanceRipSplitHorizonCfg     InterfaceEthernetIpv6InstanceRip              `json:"rip,omitempty"`
	InterfaceEthernetIpv6InstanceRouterRipng            InterfaceEthernetIpv6InstanceRouter           `json:"router,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverAction      InterfaceEthernetIpv6InstanceRouterAdver      `json:"router-adver,omitempty"`
	InterfaceEthernetIpv6InstanceStatefulFirewallInside InterfaceEthernetIpv6InstanceStatefulFirewall `json:"stateful-firewall,omitempty"`
	InterfaceEthernetIpv6InstanceTTLIgnore              int                                           `json:"ttl-ignore,omitempty"`
	InterfaceEthernetIpv6InstanceUUID                   string                                        `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6InstanceAccessListCfg struct {
	InterfaceEthernetIpv6InstanceAccessListCfgInbound   int    `json:"inbound,omitempty"`
	InterfaceEthernetIpv6InstanceAccessListCfgV6AclName string `json:"v6-acl-name,omitempty"`
}

type InterfaceEthernetIpv6InstanceAddressList struct {
	InterfaceEthernetIpv6InstanceAddressListAddressType string `json:"address-type,omitempty"`
	InterfaceEthernetIpv6InstanceAddressListIpv6Addr    string `json:"ipv6-addr,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspf struct {
	InterfaceEthernetIpv6InstanceOspfBfd                                     int                                                      `json:"bfd,omitempty"`
	InterfaceEthernetIpv6InstanceOspfCostCfgCost                             []InterfaceEthernetIpv6InstanceOspfCostCfg               `json:"cost-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgDeadInterval             []InterfaceEthernetIpv6InstanceOspfDeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceOspfDisable                                 int                                                      `json:"disable,omitempty"`
	InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgHelloInterval           []InterfaceEthernetIpv6InstanceOspfHelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgMtuIgnore                   []InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighbor                     []InterfaceEthernetIpv6InstanceOspfNeighborCfg           `json:"neighbor-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceOspfNetworkListBroadcastType                []InterfaceEthernetIpv6InstanceOspfNetworkList           `json:"network-list,omitempty"`
	InterfaceEthernetIpv6InstanceOspfPriorityCfgPriority                     []InterfaceEthernetIpv6InstanceOspfPriorityCfg           `json:"priority-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgRetransmitInterval []InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgTransmitDelay           []InterfaceEthernetIpv6InstanceOspfTransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceOspfUUID                                    string                                                   `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6InstanceRip struct {
	InterfaceEthernetIpv6InstanceRipSplitHorizonCfgState InterfaceEthernetIpv6InstanceRipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	InterfaceEthernetIpv6InstanceRipUUID                 string                                          `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6InstanceRouter struct {
	InterfaceEthernetIpv6InstanceRouterIsisTag      InterfaceEthernetIpv6InstanceRouterIsis  `json:"isis,omitempty"`
	InterfaceEthernetIpv6InstanceRouterOspfAreaList InterfaceEthernetIpv6InstanceRouterOspf  `json:"ospf,omitempty"`
	InterfaceEthernetIpv6InstanceRouterRipngRip     InterfaceEthernetIpv6InstanceRouterRipng `json:"ripng,omitempty"`
}

type InterfaceEthernetIpv6InstanceRouterAdver struct {
	InterfaceEthernetIpv6InstanceRouterAdverAction                   string                                               `json:"action,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverAdverMtu                 int                                                  `json:"adver-mtu,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverAdverMtuDisable          int                                                  `json:"adver-mtu-disable,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverAdverVrid                int                                                  `json:"adver-vrid,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverAdverVridDefault         int                                                  `json:"adver-vrid-default,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverDefaultLifetime          int                                                  `json:"default-lifetime,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverFloatingIP               string                                               `json:"floating-ip,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverFloatingIPDefaultVrid    string                                               `json:"floating-ip-default-vrid,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverHopLimit                 int                                                  `json:"hop-limit,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverManagedConfigAction      string                                               `json:"managed-config-action,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverMaxInterval              int                                                  `json:"max-interval,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverMinInterval              int                                                  `json:"min-interval,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverOtherConfigAction        string                                               `json:"other-config-action,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverPrefixListPrefix         []InterfaceEthernetIpv6InstanceRouterAdverPrefixList `json:"prefix-list,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverRateLimit                int                                                  `json:"rate-limit,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverReachableTime            int                                                  `json:"reachable-time,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverRetransmitTimer          int                                                  `json:"retransmit-timer,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverUseFloatingIP            int                                                  `json:"use-floating-ip,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverUseFloatingIPDefaultVrid int                                                  `json:"use-floating-ip-default-vrid,omitempty"`
}

type InterfaceEthernetIpv6InstanceStatefulFirewall struct {
	InterfaceEthernetIpv6InstanceStatefulFirewallAccessList int    `json:"access-list,omitempty"`
	InterfaceEthernetIpv6InstanceStatefulFirewallAclName    string `json:"acl-name,omitempty"`
	InterfaceEthernetIpv6InstanceStatefulFirewallClassList  string `json:"class-list,omitempty"`
	InterfaceEthernetIpv6InstanceStatefulFirewallInside     int    `json:"inside,omitempty"`
	InterfaceEthernetIpv6InstanceStatefulFirewallOutside    int    `json:"outside,omitempty"`
	InterfaceEthernetIpv6InstanceStatefulFirewallUUID       string `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspfCostCfg struct {
	InterfaceEthernetIpv6InstanceOspfCostCfgCost       int `json:"cost,omitempty"`
	InterfaceEthernetIpv6InstanceOspfCostCfgInstanceID int `json:"instance-id,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspfDeadIntervalCfg struct {
	InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgDeadInterval int `json:"dead-interval,omitempty"`
	InterfaceEthernetIpv6InstanceOspfDeadIntervalCfgInstanceID   int `json:"instance-id,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspfHelloIntervalCfg struct {
	InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgHelloInterval int `json:"hello-interval,omitempty"`
	InterfaceEthernetIpv6InstanceOspfHelloIntervalCfgInstanceID    int `json:"instance-id,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfg struct {
	InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgInstanceID int `json:"instance-id,omitempty"`
	InterfaceEthernetIpv6InstanceOspfMtuIgnoreCfgMtuIgnore  int `json:"mtu-ignore,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspfNeighborCfg struct {
	InterfaceEthernetIpv6InstanceOspfNeighborCfgNeigInst             int    `json:"neig-inst,omitempty"`
	InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighbor             string `json:"neighbor,omitempty"`
	InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighborCost         int    `json:"neighbor-cost,omitempty"`
	InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	InterfaceEthernetIpv6InstanceOspfNeighborCfgNeighborPriority     int    `json:"neighbor-priority,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspfNetworkList struct {
	InterfaceEthernetIpv6InstanceOspfNetworkListBroadcastType     string `json:"broadcast-type,omitempty"`
	InterfaceEthernetIpv6InstanceOspfNetworkListNetworkInstanceID int    `json:"network-instance-id,omitempty"`
	InterfaceEthernetIpv6InstanceOspfNetworkListP2MpNbma          int    `json:"p2mp-nbma,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspfPriorityCfg struct {
	InterfaceEthernetIpv6InstanceOspfPriorityCfgInstanceID int `json:"instance-id,omitempty"`
	InterfaceEthernetIpv6InstanceOspfPriorityCfgPriority   int `json:"priority,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfg struct {
	InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgInstanceID         int `json:"instance-id,omitempty"`
	InterfaceEthernetIpv6InstanceOspfRetransmitIntervalCfgRetransmitInterval int `json:"retransmit-interval,omitempty"`
}

type InterfaceEthernetIpv6InstanceOspfTransmitDelayCfg struct {
	InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgInstanceID    int `json:"instance-id,omitempty"`
	InterfaceEthernetIpv6InstanceOspfTransmitDelayCfgTransmitDelay int `json:"transmit-delay,omitempty"`
}

type InterfaceEthernetIpv6InstanceRipSplitHorizonCfg struct {
	InterfaceEthernetIpv6InstanceRipSplitHorizonCfgState string `json:"state,omitempty"`
}

type InterfaceEthernetIpv6InstanceRouterIsis struct {
	InterfaceEthernetIpv6InstanceRouterIsisTag  string `json:"tag,omitempty"`
	InterfaceEthernetIpv6InstanceRouterIsisUUID string `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6InstanceRouterOspf struct {
	InterfaceEthernetIpv6InstanceRouterOspfAreaListAreaIDNum []InterfaceEthernetIpv6InstanceRouterOspfAreaList `json:"area-list,omitempty"`
	InterfaceEthernetIpv6InstanceRouterOspfUUID              string                                            `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6InstanceRouterRipng struct {
	InterfaceEthernetIpv6InstanceRouterRipngRip  int    `json:"rip,omitempty"`
	InterfaceEthernetIpv6InstanceRouterRipngUUID string `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6InstanceRouterAdverPrefixList struct {
	InterfaceEthernetIpv6InstanceRouterAdverPrefixListNotAutonomous     int    `json:"not-autonomous,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverPrefixListNotOnLink         int    `json:"not-on-link,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverPrefixListPreferredLifetime int    `json:"preferred-lifetime,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverPrefixListPrefix            string `json:"prefix,omitempty"`
	InterfaceEthernetIpv6InstanceRouterAdverPrefixListValidLifetime     int    `json:"valid-lifetime,omitempty"`
}

type InterfaceEthernetIpv6InstanceRouterOspfAreaList struct {
	InterfaceEthernetIpv6InstanceRouterOspfAreaListAreaIDAddr string `json:"area-id-addr,omitempty"`
	InterfaceEthernetIpv6InstanceRouterOspfAreaListAreaIDNum  int    `json:"area-id-num,omitempty"`
	InterfaceEthernetIpv6InstanceRouterOspfAreaListInstanceID int    `json:"instance-id,omitempty"`
	InterfaceEthernetIpv6InstanceRouterOspfAreaListTag        string `json:"tag,omitempty"`
}

func PostInterfaceEthernetIpv6(id string, name1 string, inst InterfaceEthernetIpv6, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceEthernetIpv6")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet/"+name1+"/ipv6", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernetIpv6
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostInterfaceEthernetIpv6", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceEthernetIpv6(id string, name1 string, host string) (*InterfaceEthernetIpv6, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceEthernetIpv6")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+name1+"/ipv6", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernetIpv6
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetInterfaceEthernetIpv6", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
