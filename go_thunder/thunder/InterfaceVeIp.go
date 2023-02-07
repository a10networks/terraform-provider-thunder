package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceVeIp struct {
	InterfaceVeIPInstanceDhcp InterfaceVeIPInstance `json:"ip,omitempty"`
}

type InterfaceVeIPInstance struct {
	InterfaceVeIPInstanceAddressListIpv4Address         []InterfaceVeIPInstanceAddressList       `json:"address-list,omitempty"`
	InterfaceVeIPInstanceAllowPromiscuousVip            int                                      `json:"allow-promiscuous-vip,omitempty"`
	InterfaceVeIPInstanceClient                         int                                      `json:"client,omitempty"`
	InterfaceVeIPInstanceDhcp                           int                                      `json:"dhcp,omitempty"`
	InterfaceVeIPInstanceGenerateMembershipQuery        int                                      `json:"generate-membership-query,omitempty"`
	InterfaceVeIPInstanceHelperAddressListHelperAddress []InterfaceVeIPInstanceHelperAddressList `json:"helper-address-list,omitempty"`
	InterfaceVeIPInstanceInside                         int                                      `json:"inside,omitempty"`
	InterfaceVeIPInstanceMaxRespTime                    int                                      `json:"max-resp-time,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobal                 InterfaceVeIPInstanceOspf                `json:"ospf,omitempty"`
	InterfaceVeIPInstanceOutside                        int                                      `json:"outside,omitempty"`
	InterfaceVeIPInstanceQueryInterval                  int                                      `json:"query-interval,omitempty"`
	InterfaceVeIPInstanceRipAuthentication              InterfaceVeIPInstanceRip                 `json:"rip,omitempty"`
	InterfaceVeIPInstanceRouterIsis                     InterfaceVeIPInstanceRouter              `json:"router,omitempty"`
	InterfaceVeIPInstanceServer                         int                                      `json:"server,omitempty"`
	InterfaceVeIPInstanceSlbPartitionRedirect           int                                      `json:"slb-partition-redirect,omitempty"`
	InterfaceVeIPInstanceStatefulFirewallInside         InterfaceVeIPInstanceStatefulFirewall    `json:"stateful-firewall,omitempty"`
	InterfaceVeIPInstanceSynCookie                      int                                      `json:"syn-cookie,omitempty"`
	InterfaceVeIPInstanceTTLIgnore                      int                                      `json:"ttl-ignore,omitempty"`
	InterfaceVeIPInstanceUUID                           string                                   `json:"uuid,omitempty"`
}

type InterfaceVeIPInstanceAddressList struct {
	InterfaceVeIPInstanceAddressListIpv4Address string `json:"ipv4-address,omitempty"`
	InterfaceVeIPInstanceAddressListIpv4Netmask string `json:"ipv4-netmask,omitempty"`
}

type InterfaceVeIPInstanceHelperAddressList struct {
	InterfaceVeIPInstanceHelperAddressListHelperAddress string `json:"helper-address,omitempty"`
}

type InterfaceVeIPInstanceOspf struct {
	InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfg InterfaceVeIPInstanceOspfOspfGlobal   `json:"ospf-global,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListIPAddr            []InterfaceVeIPInstanceOspfOspfIPList `json:"ospf-ip-list,omitempty"`
}

type InterfaceVeIPInstanceRip struct {
	InterfaceVeIPInstanceRipAuthenticationStr    InterfaceVeIPInstanceRipAuthentication  `json:"authentication,omitempty"`
	InterfaceVeIPInstanceRipReceiveCfgReceive    InterfaceVeIPInstanceRipReceiveCfg      `json:"receive-cfg,omitempty"`
	InterfaceVeIPInstanceRipReceivePacket        int                                     `json:"receive-packet,omitempty"`
	InterfaceVeIPInstanceRipSendCfgSend          InterfaceVeIPInstanceRipSendCfg         `json:"send-cfg,omitempty"`
	InterfaceVeIPInstanceRipSendPacket           int                                     `json:"send-packet,omitempty"`
	InterfaceVeIPInstanceRipSplitHorizonCfgState InterfaceVeIPInstanceRipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	InterfaceVeIPInstanceRipUUID                 string                                  `json:"uuid,omitempty"`
}

type InterfaceVeIPInstanceRouter struct {
	InterfaceVeIPInstanceRouterIsisTag InterfaceVeIPInstanceRouterIsis `json:"isis,omitempty"`
}

type InterfaceVeIPInstanceStatefulFirewall struct {
	InterfaceVeIPInstanceStatefulFirewallAccessList int    `json:"access-list,omitempty"`
	InterfaceVeIPInstanceStatefulFirewallAclID      int    `json:"acl-id,omitempty"`
	InterfaceVeIPInstanceStatefulFirewallClassList  string `json:"class-list,omitempty"`
	InterfaceVeIPInstanceStatefulFirewallInside     int    `json:"inside,omitempty"`
	InterfaceVeIPInstanceStatefulFirewallOutside    int    `json:"outside,omitempty"`
	InterfaceVeIPInstanceStatefulFirewallUUID       string `json:"uuid,omitempty"`
}

type InterfaceVeIPInstanceOspfOspfGlobal struct {
	InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfgAuthentication  InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalAuthenticationKey                string                                                `json:"authentication-key,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalBfdCfgBfd                        InterfaceVeIPInstanceOspfOspfGlobalBfdCfg             `json:"bfd-cfg,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalCost                             int                                                   `json:"cost,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalDatabaseFilterCfgDatabaseFilter  InterfaceVeIPInstanceOspfOspfGlobalDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalDeadInterval                     int                                                   `json:"dead-interval,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalDisable                          string                                                `json:"disable,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalHelloInterval                    int                                                   `json:"hello-interval,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey []InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalMtu                              int                                                   `json:"mtu,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalMtuIgnore                        int                                                   `json:"mtu-ignore,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalNetworkBroadcast                 InterfaceVeIPInstanceOspfOspfGlobalNetwork            `json:"network,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalPriority                         int                                                   `json:"priority,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalRetransmitInterval               int                                                   `json:"retransmit-interval,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalTransmitDelay                    int                                                   `json:"transmit-delay,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalUUID                             string                                                `json:"uuid,omitempty"`
}

type InterfaceVeIPInstanceOspfOspfIPList struct {
	InterfaceVeIPInstanceOspfOspfIPListAuthentication                   int                                                   `json:"authentication,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListAuthenticationKey                string                                                `json:"authentication-key,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListCost                             int                                                   `json:"cost,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListDatabaseFilter                   string                                                `json:"database-filter,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListDeadInterval                     int                                                   `json:"dead-interval,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListHelloInterval                    int                                                   `json:"hello-interval,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListIPAddr                           string                                                `json:"ip-addr,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey []InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListMtuIgnore                        int                                                   `json:"mtu-ignore,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListOut                              int                                                   `json:"out,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListPriority                         int                                                   `json:"priority,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListRetransmitInterval               int                                                   `json:"retransmit-interval,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListTransmitDelay                    int                                                   `json:"transmit-delay,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListUUID                             string                                                `json:"uuid,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListValue                            string                                                `json:"value,omitempty"`
}

type InterfaceVeIPInstanceRipAuthentication struct {
	InterfaceVeIPInstanceRipAuthenticationKeyChainKeyChain InterfaceVeIPInstanceRipAuthenticationKeyChain `json:"key-chain,omitempty"`
	InterfaceVeIPInstanceRipAuthenticationModeMode         InterfaceVeIPInstanceRipAuthenticationMode     `json:"mode,omitempty"`
	InterfaceVeIPInstanceRipAuthenticationStrString        InterfaceVeIPInstanceRipAuthenticationStr      `json:"str,omitempty"`
}

type InterfaceVeIPInstanceRipReceiveCfg struct {
	InterfaceVeIPInstanceRipReceiveCfgReceive int    `json:"receive,omitempty"`
	InterfaceVeIPInstanceRipReceiveCfgVersion string `json:"version,omitempty"`
}

type InterfaceVeIPInstanceRipSendCfg struct {
	InterfaceVeIPInstanceRipSendCfgSend    int    `json:"send,omitempty"`
	InterfaceVeIPInstanceRipSendCfgVersion string `json:"version,omitempty"`
}

type InterfaceVeIPInstanceRipSplitHorizonCfg struct {
	InterfaceVeIPInstanceRipSplitHorizonCfgState string `json:"state,omitempty"`
}

type InterfaceVeIPInstanceRouterIsis struct {
	InterfaceVeIPInstanceRouterIsisTag  string `json:"tag,omitempty"`
	InterfaceVeIPInstanceRouterIsisUUID string `json:"uuid,omitempty"`
}

type InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfg struct {
	InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfgAuthentication int    `json:"authentication,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfgValue          string `json:"value,omitempty"`
}

type InterfaceVeIPInstanceOspfOspfGlobalBfdCfg struct {
	InterfaceVeIPInstanceOspfOspfGlobalBfdCfgBfd     int `json:"bfd,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalBfdCfgDisable int `json:"disable,omitempty"`
}

type InterfaceVeIPInstanceOspfOspfGlobalDatabaseFilterCfg struct {
	InterfaceVeIPInstanceOspfOspfGlobalDatabaseFilterCfgDatabaseFilter string `json:"database-filter,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalDatabaseFilterCfgOut            int    `json:"out,omitempty"`
}

type InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfg struct {
	InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceVeIPInstanceOspfOspfGlobalNetwork struct {
	InterfaceVeIPInstanceOspfOspfGlobalNetworkBroadcast         int `json:"broadcast,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalNetworkNonBroadcast      int `json:"non-broadcast,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalNetworkP2MpNbma          int `json:"p2mp-nbma,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalNetworkPointToMultipoint int `json:"point-to-multipoint,omitempty"`
	InterfaceVeIPInstanceOspfOspfGlobalNetworkPointToPoint      int `json:"point-to-point,omitempty"`
}

type InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfg struct {
	InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceVeIPInstanceRipAuthenticationKeyChain struct {
	InterfaceVeIPInstanceRipAuthenticationKeyChainKeyChain string `json:"key-chain,omitempty"`
}

type InterfaceVeIPInstanceRipAuthenticationMode struct {
	InterfaceVeIPInstanceRipAuthenticationModeMode string `json:"mode,omitempty"`
}

type InterfaceVeIPInstanceRipAuthenticationStr struct {
	InterfaceVeIPInstanceRipAuthenticationStrString string `json:"string,omitempty"`
}

func PostInterfaceVeIp(id string, name1 string, inst InterfaceVeIp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceVeIp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ve/"+name1+"/ip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVeIp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostInterfaceVeIp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceVeIp(id string, name1 string, host string) (*InterfaceVeIp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceVeIp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ve/"+name1+"/ip", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVeIp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetInterfaceVeIp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
