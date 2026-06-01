package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type InterfaceVeIp struct {
	Inst struct {
		AddressList []InterfaceVeIpAddressList `json:"address-list"`

		AllowPromiscuousVip int `json:"allow-promiscuous-vip"`

		Client int `json:"client"`

		Dhcp int `json:"dhcp"`

		Dmz int `json:"dmz"`

		GenerateMembershipQuery int `json:"generate-membership-query"`

		HelperAddressList []InterfaceVeIpHelperAddressList `json:"helper-address-list"`

		Inside int `json:"inside"`

		MaxRespTime int `json:"max-resp-time" dval:"100"`

		Ospf InterfaceVeIpOspf1006 `json:"ospf"`

		Outside int `json:"outside"`

		QueryInterval int `json:"query-interval" dval:"125"`

		Rip InterfaceVeIpRip1014 `json:"rip"`

		Router InterfaceVeIpRouter1022 `json:"router"`

		Server int `json:"server"`

		SlbPartitionRedirect int `json:"slb-partition-redirect"`

		StatefulFirewall InterfaceVeIpStatefulFirewall1024 `json:"stateful-firewall"`

		SynCookie int `json:"syn-cookie"`

		TtlIgnore int `json:"ttl-ignore"`

		Unnumbered int `json:"unnumbered"`

		Uuid string `json:"uuid"`

		Ifnum string
	} `json:"ip"`
}

type InterfaceVeIpAddressList struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceVeIpHelperAddressList struct {
	HelperAddress string `json:"helper-address"`
}

type InterfaceVeIpOspf1006 struct {
	OspfGlobal InterfaceVeIpOspfOspfGlobal1007 `json:"ospf-global"`
	OspfIpList []InterfaceVeIpOspfOspfIpList   `json:"ospf-ip-list"`
}

type InterfaceVeIpOspfOspfGlobal1007 struct {
	AuthenticationCfg  InterfaceVeIpOspfOspfGlobalAuthenticationCfg1008  `json:"authentication-cfg"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	BfdCfg             InterfaceVeIpOspfOspfGlobalBfdCfg1009             `json:"bfd-cfg"`
	Cost               int                                               `json:"cost"`
	DatabaseFilterCfg  InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg1010  `json:"database-filter-cfg"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	Disable            string                                            `json:"disable"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceVeIpOspfOspfGlobalMessageDigestCfg1011 `json:"message-digest-cfg"`
	Mtu                int                                               `json:"mtu"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Network            InterfaceVeIpOspfOspfGlobalNetwork1013            `json:"network"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceVeIpOspfOspfGlobalAuthenticationCfg1008 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceVeIpOspfOspfGlobalBfdCfg1009 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg1010 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceVeIpOspfOspfGlobalMessageDigestCfg1011 struct {
	MessageDigestKey int                                                `json:"message-digest-key"`
	Md5              InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd51012 `json:"md5"`
}

type InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd51012 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceVeIpOspfOspfGlobalNetwork1013 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceVeIpOspfOspfIpList struct {
	IpAddr             string                                        `json:"ip-addr"`
	Authentication     int                                           `json:"authentication"`
	Value              string                                        `json:"value"`
	AuthenticationKey  string                                        `json:"authentication-key"`
	Cost               int                                           `json:"cost"`
	DatabaseFilter     string                                        `json:"database-filter"`
	Out                int                                           `json:"out"`
	DeadInterval       int                                           `json:"dead-interval" dval:"40"`
	HelloInterval      int                                           `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceVeIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
	MtuIgnore          int                                           `json:"mtu-ignore"`
	Priority           int                                           `json:"priority" dval:"1"`
	RetransmitInterval int                                           `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                           `json:"transmit-delay" dval:"1"`
	Uuid               string                                        `json:"uuid"`
}

type InterfaceVeIpOspfOspfIpListMessageDigestCfg struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceVeIpRip1014 struct {
	Authentication  InterfaceVeIpRipAuthentication1015  `json:"authentication"`
	SendPacket      int                                 `json:"send-packet" dval:"1"`
	ReceivePacket   int                                 `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceVeIpRipSendCfg1019         `json:"send-cfg"`
	ReceiveCfg      InterfaceVeIpRipReceiveCfg1020      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceVeIpRipSplitHorizonCfg1021 `json:"split-horizon-cfg"`
	Uuid            string                              `json:"uuid"`
}

type InterfaceVeIpRipAuthentication1015 struct {
	Str      InterfaceVeIpRipAuthenticationStr1016      `json:"str"`
	Mode     InterfaceVeIpRipAuthenticationMode1017     `json:"mode"`
	KeyChain InterfaceVeIpRipAuthenticationKeyChain1018 `json:"key-chain"`
}

type InterfaceVeIpRipAuthenticationStr1016 struct {
	String string `json:"string"`
}

type InterfaceVeIpRipAuthenticationMode1017 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceVeIpRipAuthenticationKeyChain1018 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceVeIpRipSendCfg1019 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceVeIpRipReceiveCfg1020 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceVeIpRipSplitHorizonCfg1021 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceVeIpRouter1022 struct {
	Isis InterfaceVeIpRouterIsis1023 `json:"isis"`
}

type InterfaceVeIpRouterIsis1023 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceVeIpStatefulFirewall1024 struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclId      int    `json:"acl-id"`
	Uuid       string `json:"uuid"`
}

func (p *InterfaceVeIp) GetId() string {
	return "1"
}

func (p *InterfaceVeIp) getPath() string {
	return "interface/ve/" + p.Inst.Ifnum + "/ip"
}

func (p *InterfaceVeIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVeIp::Post")
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

func (p *InterfaceVeIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVeIp::Get")
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
func (p *InterfaceVeIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVeIp::Put")
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

func (p *InterfaceVeIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVeIp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
