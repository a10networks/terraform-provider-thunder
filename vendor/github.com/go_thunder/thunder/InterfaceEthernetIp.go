package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceEthernetIp struct {
	InterfaceEthernetIPInstanceDhcp InterfaceEthernetIPInstance `json:"ip,omitempty"`
}

type InterfaceEthernetIPInstance struct {
	InterfaceEthernetIPInstanceAddressListIpv4Address         []InterfaceEthernetIPInstanceAddressList       `json:"address-list,omitempty"`
	InterfaceEthernetIPInstanceAllowPromiscuousVip            int                                            `json:"allow-promiscuous-vip,omitempty"`
	InterfaceEthernetIPInstanceCacheSpoofingPort              int                                            `json:"cache-spoofing-port,omitempty"`
	InterfaceEthernetIPInstanceClient                         int                                            `json:"client,omitempty"`
	InterfaceEthernetIPInstanceDhcp                           int                                            `json:"dhcp,omitempty"`
	InterfaceEthernetIPInstanceGenerateMembershipQuery        int                                            `json:"generate-membership-query,omitempty"`
	InterfaceEthernetIPInstanceHelperAddressListHelperAddress []InterfaceEthernetIPInstanceHelperAddressList `json:"helper-address-list,omitempty"`
	InterfaceEthernetIPInstanceInside                         int                                            `json:"inside,omitempty"`
	InterfaceEthernetIPInstanceMaxRespTime                    int                                            `json:"max-resp-time,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobal                 InterfaceEthernetIPInstanceOspf                `json:"ospf,omitempty"`
	InterfaceEthernetIPInstanceOutside                        int                                            `json:"outside,omitempty"`
	InterfaceEthernetIPInstanceQueryInterval                  int                                            `json:"query-interval,omitempty"`
	InterfaceEthernetIPInstanceRipAuthentication              InterfaceEthernetIPInstanceRip                 `json:"rip,omitempty"`
	InterfaceEthernetIPInstanceRouterIsis                     InterfaceEthernetIPInstanceRouter              `json:"router,omitempty"`
	InterfaceEthernetIPInstanceServer                         int                                            `json:"server,omitempty"`
	InterfaceEthernetIPInstanceSlbPartitionRedirect           int                                            `json:"slb-partition-redirect,omitempty"`
	InterfaceEthernetIPInstanceStatefulFirewallInside         InterfaceEthernetIPInstanceStatefulFirewall    `json:"stateful-firewall,omitempty"`
	InterfaceEthernetIPInstanceSynCookie                      int                                            `json:"syn-cookie,omitempty"`
	InterfaceEthernetIPInstanceTTLIgnore                      int                                            `json:"ttl-ignore,omitempty"`
	InterfaceEthernetIPInstanceUUID                           string                                         `json:"uuid,omitempty"`
}

type InterfaceEthernetIPInstanceAddressList struct {
	InterfaceEthernetIPInstanceAddressListIpv4Address string `json:"ipv4-address,omitempty"`
	InterfaceEthernetIPInstanceAddressListIpv4Netmask string `json:"ipv4-netmask,omitempty"`
}

type InterfaceEthernetIPInstanceHelperAddressList struct {
	InterfaceEthernetIPInstanceHelperAddressListHelperAddress string `json:"helper-address,omitempty"`
}

type InterfaceEthernetIPInstanceOspf struct {
	InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfg InterfaceEthernetIPInstanceOspfOspfGlobal   `json:"ospf-global,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListIPAddr            []InterfaceEthernetIPInstanceOspfOspfIPList `json:"ospf-ip-list,omitempty"`
}

type InterfaceEthernetIPInstanceRip struct {
	InterfaceEthernetIPInstanceRipAuthenticationStr    InterfaceEthernetIPInstanceRipAuthentication  `json:"authentication,omitempty"`
	InterfaceEthernetIPInstanceRipReceiveCfgReceive    InterfaceEthernetIPInstanceRipReceiveCfg      `json:"receive-cfg,omitempty"`
	InterfaceEthernetIPInstanceRipReceivePacket        int                                           `json:"receive-packet,omitempty"`
	InterfaceEthernetIPInstanceRipSendCfgSend          InterfaceEthernetIPInstanceRipSendCfg         `json:"send-cfg,omitempty"`
	InterfaceEthernetIPInstanceRipSendPacket           int                                           `json:"send-packet,omitempty"`
	InterfaceEthernetIPInstanceRipSplitHorizonCfgState InterfaceEthernetIPInstanceRipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	InterfaceEthernetIPInstanceRipUUID                 string                                        `json:"uuid,omitempty"`
}

type InterfaceEthernetIPInstanceRouter struct {
	InterfaceEthernetIPInstanceRouterIsisTag InterfaceEthernetIPInstanceRouterIsis `json:"isis,omitempty"`
}

type InterfaceEthernetIPInstanceStatefulFirewall struct {
	InterfaceEthernetIPInstanceStatefulFirewallAccessList int    `json:"access-list,omitempty"`
	InterfaceEthernetIPInstanceStatefulFirewallAclID      int    `json:"acl-id,omitempty"`
	InterfaceEthernetIPInstanceStatefulFirewallClassList  string `json:"class-list,omitempty"`
	InterfaceEthernetIPInstanceStatefulFirewallInside     int    `json:"inside,omitempty"`
	InterfaceEthernetIPInstanceStatefulFirewallOutside    int    `json:"outside,omitempty"`
	InterfaceEthernetIPInstanceStatefulFirewallUUID       string `json:"uuid,omitempty"`
}

type InterfaceEthernetIPInstanceOspfOspfGlobal struct {
	InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfgAuthentication  InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationKey                string                                                      `json:"authentication-key,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalBfdCfgBfd                        InterfaceEthernetIPInstanceOspfOspfGlobalBfdCfg             `json:"bfd-cfg,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalCost                             int                                                         `json:"cost,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalDatabaseFilterCfgDatabaseFilter  InterfaceEthernetIPInstanceOspfOspfGlobalDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalDeadInterval                     int                                                         `json:"dead-interval,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalDisable                          string                                                      `json:"disable,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalHelloInterval                    int                                                         `json:"hello-interval,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey []InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalMtu                              int                                                         `json:"mtu,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalMtuIgnore                        int                                                         `json:"mtu-ignore,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalNetworkBroadcast                 InterfaceEthernetIPInstanceOspfOspfGlobalNetwork            `json:"network,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalPriority                         int                                                         `json:"priority,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalRetransmitInterval               int                                                         `json:"retransmit-interval,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalTransmitDelay                    int                                                         `json:"transmit-delay,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalUUID                             string                                                      `json:"uuid,omitempty"`
}

type InterfaceEthernetIPInstanceOspfOspfIPList struct {
	InterfaceEthernetIPInstanceOspfOspfIPListAuthentication                   int                                                         `json:"authentication,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListAuthenticationKey                string                                                      `json:"authentication-key,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListCost                             int                                                         `json:"cost,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListDatabaseFilter                   string                                                      `json:"database-filter,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListDeadInterval                     int                                                         `json:"dead-interval,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListHelloInterval                    int                                                         `json:"hello-interval,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListIPAddr                           string                                                      `json:"ip-addr,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey []InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListMtuIgnore                        int                                                         `json:"mtu-ignore,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListOut                              int                                                         `json:"out,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListPriority                         int                                                         `json:"priority,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListRetransmitInterval               int                                                         `json:"retransmit-interval,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListTransmitDelay                    int                                                         `json:"transmit-delay,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListUUID                             string                                                      `json:"uuid,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListValue                            string                                                      `json:"value,omitempty"`
}

type InterfaceEthernetIPInstanceRipAuthentication struct {
	InterfaceEthernetIPInstanceRipAuthenticationKeyChainKeyChain InterfaceEthernetIPInstanceRipAuthenticationKeyChain `json:"key-chain,omitempty"`
	InterfaceEthernetIPInstanceRipAuthenticationModeMode         InterfaceEthernetIPInstanceRipAuthenticationMode     `json:"mode,omitempty"`
	InterfaceEthernetIPInstanceRipAuthenticationStrString        InterfaceEthernetIPInstanceRipAuthenticationStr      `json:"str,omitempty"`
}

type InterfaceEthernetIPInstanceRipReceiveCfg struct {
	InterfaceEthernetIPInstanceRipReceiveCfgReceive int    `json:"receive,omitempty"`
	InterfaceEthernetIPInstanceRipReceiveCfgVersion string `json:"version,omitempty"`
}

type InterfaceEthernetIPInstanceRipSendCfg struct {
	InterfaceEthernetIPInstanceRipSendCfgSend    int    `json:"send,omitempty"`
	InterfaceEthernetIPInstanceRipSendCfgVersion string `json:"version,omitempty"`
}

type InterfaceEthernetIPInstanceRipSplitHorizonCfg struct {
	InterfaceEthernetIPInstanceRipSplitHorizonCfgState string `json:"state,omitempty"`
}

type InterfaceEthernetIPInstanceRouterIsis struct {
	InterfaceEthernetIPInstanceRouterIsisTag  string `json:"tag,omitempty"`
	InterfaceEthernetIPInstanceRouterIsisUUID string `json:"uuid,omitempty"`
}

type InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfg struct {
	InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfgAuthentication int    `json:"authentication,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfgValue          string `json:"value,omitempty"`
}

type InterfaceEthernetIPInstanceOspfOspfGlobalBfdCfg struct {
	InterfaceEthernetIPInstanceOspfOspfGlobalBfdCfgBfd     int `json:"bfd,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalBfdCfgDisable int `json:"disable,omitempty"`
}

type InterfaceEthernetIPInstanceOspfOspfGlobalDatabaseFilterCfg struct {
	InterfaceEthernetIPInstanceOspfOspfGlobalDatabaseFilterCfgDatabaseFilter string `json:"database-filter,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalDatabaseFilterCfgOut            int    `json:"out,omitempty"`
}

type InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfg struct {
	InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceEthernetIPInstanceOspfOspfGlobalNetwork struct {
	InterfaceEthernetIPInstanceOspfOspfGlobalNetworkBroadcast         int `json:"broadcast,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalNetworkNonBroadcast      int `json:"non-broadcast,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalNetworkP2MpNbma          int `json:"p2mp-nbma,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalNetworkPointToMultipoint int `json:"point-to-multipoint,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfGlobalNetworkPointToPoint      int `json:"point-to-point,omitempty"`
}

type InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfg struct {
	InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceEthernetIPInstanceRipAuthenticationKeyChain struct {
	InterfaceEthernetIPInstanceRipAuthenticationKeyChainKeyChain string `json:"key-chain,omitempty"`
}

type InterfaceEthernetIPInstanceRipAuthenticationMode struct {
	InterfaceEthernetIPInstanceRipAuthenticationModeMode string `json:"mode,omitempty"`
}

type InterfaceEthernetIPInstanceRipAuthenticationStr struct {
	InterfaceEthernetIPInstanceRipAuthenticationStrString string `json:"string,omitempty"`
}

func PostInterfaceEthernetIp(id string, name1 string, inst InterfaceEthernetIp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceEthernetIp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet/"+name1+"/ip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernetIp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostInterfaceEthernetIp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceEthernetIp(id string, name1 string, host string) (*InterfaceEthernetIp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceEthernetIp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+name1+"/ip", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernetIp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetInterfaceEthernetIp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
