package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type InterfaceLifIp struct {
	Inst struct {
		AddressList []InterfaceLifIpAddressList `json:"address-list"`

		AllowPromiscuousVip int `json:"allow-promiscuous-vip"`

		CacheSpoofingPort int `json:"cache-spoofing-port"`

		Client int `json:"client"`

		Dhcp int `json:"dhcp"`

		Dmz int `json:"dmz"`

		GenerateMembershipQuery int `json:"generate-membership-query"`

		Inside int `json:"inside"`

		MaxRespTime int `json:"max-resp-time" dval:"100"`

		Ospf InterfaceLifIpOspf650 `json:"ospf"`

		Outside int `json:"outside"`

		QueryInterval int `json:"query-interval" dval:"125"`

		Rip InterfaceLifIpRip658 `json:"rip"`

		Router InterfaceLifIpRouter666 `json:"router"`

		Server int `json:"server"`

		Unnumbered int `json:"unnumbered"`

		Uuid string `json:"uuid"`

		Ifname string
	} `json:"ip"`
}

type InterfaceLifIpAddressList struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceLifIpOspf650 struct {
	OspfGlobal InterfaceLifIpOspfOspfGlobal651 `json:"ospf-global"`
	OspfIpList []InterfaceLifIpOspfOspfIpList  `json:"ospf-ip-list"`
}

type InterfaceLifIpOspfOspfGlobal651 struct {
	AuthenticationCfg  InterfaceLifIpOspfOspfGlobalAuthenticationCfg652  `json:"authentication-cfg"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	BfdCfg             InterfaceLifIpOspfOspfGlobalBfdCfg653             `json:"bfd-cfg"`
	Cost               int                                               `json:"cost"`
	DatabaseFilterCfg  InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg654  `json:"database-filter-cfg"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	Disable            string                                            `json:"disable"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLifIpOspfOspfGlobalMessageDigestCfg655 `json:"message-digest-cfg"`
	Mtu                int                                               `json:"mtu"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Network            InterfaceLifIpOspfOspfGlobalNetwork657            `json:"network"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceLifIpOspfOspfGlobalAuthenticationCfg652 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceLifIpOspfOspfGlobalBfdCfg653 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg654 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceLifIpOspfOspfGlobalMessageDigestCfg655 struct {
	MessageDigestKey int                                                `json:"message-digest-key"`
	Md5              InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5656 `json:"md5"`
}

type InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5656 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLifIpOspfOspfGlobalNetwork657 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceLifIpOspfOspfIpList struct {
	IpAddr             string                                         `json:"ip-addr"`
	Authentication     int                                            `json:"authentication"`
	Value              string                                         `json:"value"`
	AuthenticationKey  string                                         `json:"authentication-key"`
	Cost               int                                            `json:"cost"`
	DatabaseFilter     string                                         `json:"database-filter"`
	Out                int                                            `json:"out"`
	DeadInterval       int                                            `json:"dead-interval" dval:"40"`
	HelloInterval      int                                            `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLifIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
	MtuIgnore          int                                            `json:"mtu-ignore"`
	Priority           int                                            `json:"priority" dval:"1"`
	RetransmitInterval int                                            `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                            `json:"transmit-delay" dval:"1"`
	Uuid               string                                         `json:"uuid"`
}

type InterfaceLifIpOspfOspfIpListMessageDigestCfg struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceLifIpRip658 struct {
	Authentication  InterfaceLifIpRipAuthentication659  `json:"authentication"`
	SendPacket      int                                 `json:"send-packet" dval:"1"`
	ReceivePacket   int                                 `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceLifIpRipSendCfg663         `json:"send-cfg"`
	ReceiveCfg      InterfaceLifIpRipReceiveCfg664      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceLifIpRipSplitHorizonCfg665 `json:"split-horizon-cfg"`
	Uuid            string                              `json:"uuid"`
}

type InterfaceLifIpRipAuthentication659 struct {
	Str      InterfaceLifIpRipAuthenticationStr660      `json:"str"`
	Mode     InterfaceLifIpRipAuthenticationMode661     `json:"mode"`
	KeyChain InterfaceLifIpRipAuthenticationKeyChain662 `json:"key-chain"`
}

type InterfaceLifIpRipAuthenticationStr660 struct {
	String string `json:"string"`
}

type InterfaceLifIpRipAuthenticationMode661 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceLifIpRipAuthenticationKeyChain662 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceLifIpRipSendCfg663 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceLifIpRipReceiveCfg664 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceLifIpRipSplitHorizonCfg665 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLifIpRouter666 struct {
	Isis InterfaceLifIpRouterIsis667 `json:"isis"`
}

type InterfaceLifIpRouterIsis667 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

func (p *InterfaceLifIp) GetId() string {
	return "1"
}

func (p *InterfaceLifIp) getPath() string {
	return "interface/lif/" + p.Inst.Ifname + "/ip"
}

func (p *InterfaceLifIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLifIp::Post")
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

func (p *InterfaceLifIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLifIp::Get")
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
func (p *InterfaceLifIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLifIp::Put")
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

func (p *InterfaceLifIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLifIp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
