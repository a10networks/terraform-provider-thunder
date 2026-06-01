package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type InterfaceTrunkIp struct {
	Inst struct {
		AddressList []InterfaceTrunkIpAddressList `json:"address-list"`

		AllowPromiscuousVip int `json:"allow-promiscuous-vip"`

		CacheSpoofingPort int `json:"cache-spoofing-port"`

		Client int `json:"client"`

		Dhcp int `json:"dhcp"`

		Dmz int `json:"dmz"`

		GenerateMembershipQuery int `json:"generate-membership-query"`

		HelperAddressList []InterfaceTrunkIpHelperAddressList `json:"helper-address-list"`

		MaxRespTime int `json:"max-resp-time" dval:"100"`

		Nat InterfaceTrunkIpNat `json:"nat"`

		Ospf InterfaceTrunkIpOspf824 `json:"ospf"`

		QueryInterval int `json:"query-interval" dval:"125"`

		Rip InterfaceTrunkIpRip832 `json:"rip"`

		Router InterfaceTrunkIpRouter840 `json:"router"`

		Server int `json:"server"`

		SlbPartitionRedirect int `json:"slb-partition-redirect"`

		StatefulFirewall InterfaceTrunkIpStatefulFirewall842 `json:"stateful-firewall"`

		SynCookie int `json:"syn-cookie"`

		TtlIgnore int `json:"ttl-ignore"`

		Unnumbered int `json:"unnumbered"`

		Uuid string `json:"uuid"`

		Ifnum string
	} `json:"ip"`
}

type InterfaceTrunkIpAddressList struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceTrunkIpHelperAddressList struct {
	HelperAddress string `json:"helper-address"`
}

type InterfaceTrunkIpNat struct {
	Inside  int `json:"inside"`
	Outside int `json:"outside"`
}

type InterfaceTrunkIpOspf824 struct {
	OspfGlobal InterfaceTrunkIpOspfOspfGlobal825 `json:"ospf-global"`
	OspfIpList []InterfaceTrunkIpOspfOspfIpList  `json:"ospf-ip-list"`
}

type InterfaceTrunkIpOspfOspfGlobal825 struct {
	AuthenticationCfg  InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg826  `json:"authentication-cfg"`
	AuthenticationKey  string                                              `json:"authentication-key"`
	BfdCfg             InterfaceTrunkIpOspfOspfGlobalBfdCfg827             `json:"bfd-cfg"`
	Cost               int                                                 `json:"cost"`
	DatabaseFilterCfg  InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg828  `json:"database-filter-cfg"`
	DeadInterval       int                                                 `json:"dead-interval" dval:"40"`
	Disable            string                                              `json:"disable"`
	HelloInterval      int                                                 `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg829 `json:"message-digest-cfg"`
	Mtu                int                                                 `json:"mtu"`
	MtuIgnore          int                                                 `json:"mtu-ignore"`
	Network            InterfaceTrunkIpOspfOspfGlobalNetwork831            `json:"network"`
	Priority           int                                                 `json:"priority" dval:"1"`
	RetransmitInterval int                                                 `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                 `json:"transmit-delay" dval:"1"`
	Uuid               string                                              `json:"uuid"`
}

type InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg826 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceTrunkIpOspfOspfGlobalBfdCfg827 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg828 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg829 struct {
	MessageDigestKey int                                                  `json:"message-digest-key"`
	Md5              InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5830 `json:"md5"`
}

type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5830 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceTrunkIpOspfOspfGlobalNetwork831 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceTrunkIpOspfOspfIpList struct {
	IpAddr             string                                           `json:"ip-addr"`
	Authentication     int                                              `json:"authentication"`
	Value              string                                           `json:"value"`
	AuthenticationKey  string                                           `json:"authentication-key"`
	Cost               int                                              `json:"cost"`
	DatabaseFilter     string                                           `json:"database-filter"`
	Out                int                                              `json:"out"`
	DeadInterval       int                                              `json:"dead-interval" dval:"40"`
	HelloInterval      int                                              `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTrunkIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
	MtuIgnore          int                                              `json:"mtu-ignore"`
	Priority           int                                              `json:"priority" dval:"1"`
	RetransmitInterval int                                              `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                              `json:"transmit-delay" dval:"1"`
	Uuid               string                                           `json:"uuid"`
}

type InterfaceTrunkIpOspfOspfIpListMessageDigestCfg struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceTrunkIpRip832 struct {
	Authentication  InterfaceTrunkIpRipAuthentication833  `json:"authentication"`
	SendPacket      int                                   `json:"send-packet" dval:"1"`
	ReceivePacket   int                                   `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceTrunkIpRipSendCfg837         `json:"send-cfg"`
	ReceiveCfg      InterfaceTrunkIpRipReceiveCfg838      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceTrunkIpRipSplitHorizonCfg839 `json:"split-horizon-cfg"`
	Uuid            string                                `json:"uuid"`
}

type InterfaceTrunkIpRipAuthentication833 struct {
	Str      InterfaceTrunkIpRipAuthenticationStr834      `json:"str"`
	Mode     InterfaceTrunkIpRipAuthenticationMode835     `json:"mode"`
	KeyChain InterfaceTrunkIpRipAuthenticationKeyChain836 `json:"key-chain"`
}

type InterfaceTrunkIpRipAuthenticationStr834 struct {
	String string `json:"string"`
}

type InterfaceTrunkIpRipAuthenticationMode835 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceTrunkIpRipAuthenticationKeyChain836 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceTrunkIpRipSendCfg837 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceTrunkIpRipReceiveCfg838 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceTrunkIpRipSplitHorizonCfg839 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceTrunkIpRouter840 struct {
	Isis InterfaceTrunkIpRouterIsis841 `json:"isis"`
}

type InterfaceTrunkIpRouterIsis841 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceTrunkIpStatefulFirewall842 struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclId      int    `json:"acl-id"`
	Uuid       string `json:"uuid"`
}

func (p *InterfaceTrunkIp) GetId() string {
	return "1"
}

func (p *InterfaceTrunkIp) getPath() string {
	return "interface/trunk/" + p.Inst.Ifnum + "/ip"
}

func (p *InterfaceTrunkIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunkIp::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *InterfaceTrunkIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunkIp::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *InterfaceTrunkIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunkIp::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}

func (p *InterfaceTrunkIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunkIp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
