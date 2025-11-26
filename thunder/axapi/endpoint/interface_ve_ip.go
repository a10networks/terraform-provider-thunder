package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
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

		Ospf InterfaceVeIpOspf1011 `json:"ospf"`

		Outside int `json:"outside"`

		QueryInterval int `json:"query-interval" dval:"125"`

		Rip InterfaceVeIpRip1019 `json:"rip"`

		Router InterfaceVeIpRouter1027 `json:"router"`

		Server int `json:"server"`

		SlbPartitionRedirect int `json:"slb-partition-redirect"`

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

type InterfaceVeIpOspf1011 struct {
	OspfGlobal InterfaceVeIpOspfOspfGlobal1012 `json:"ospf-global"`
	OspfIpList []InterfaceVeIpOspfOspfIpList   `json:"ospf-ip-list"`
}

type InterfaceVeIpOspfOspfGlobal1012 struct {
	AuthenticationCfg  InterfaceVeIpOspfOspfGlobalAuthenticationCfg1013  `json:"authentication-cfg"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	BfdCfg             InterfaceVeIpOspfOspfGlobalBfdCfg1014             `json:"bfd-cfg"`
	Cost               int                                               `json:"cost"`
	DatabaseFilterCfg  InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg1015  `json:"database-filter-cfg"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	Disable            string                                            `json:"disable"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceVeIpOspfOspfGlobalMessageDigestCfg1016 `json:"message-digest-cfg"`
	Mtu                int                                               `json:"mtu"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Network            InterfaceVeIpOspfOspfGlobalNetwork1018            `json:"network"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceVeIpOspfOspfGlobalAuthenticationCfg1013 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceVeIpOspfOspfGlobalBfdCfg1014 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg1015 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceVeIpOspfOspfGlobalMessageDigestCfg1016 struct {
	MessageDigestKey int                                                `json:"message-digest-key"`
	Md5              InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd51017 `json:"md5"`
}

type InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd51017 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceVeIpOspfOspfGlobalNetwork1018 struct {
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

type InterfaceVeIpRip1019 struct {
	Authentication  InterfaceVeIpRipAuthentication1020  `json:"authentication"`
	SendPacket      int                                 `json:"send-packet" dval:"1"`
	ReceivePacket   int                                 `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceVeIpRipSendCfg1024         `json:"send-cfg"`
	ReceiveCfg      InterfaceVeIpRipReceiveCfg1025      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceVeIpRipSplitHorizonCfg1026 `json:"split-horizon-cfg"`
	Uuid            string                              `json:"uuid"`
}

type InterfaceVeIpRipAuthentication1020 struct {
	Str      InterfaceVeIpRipAuthenticationStr1021      `json:"str"`
	Mode     InterfaceVeIpRipAuthenticationMode1022     `json:"mode"`
	KeyChain InterfaceVeIpRipAuthenticationKeyChain1023 `json:"key-chain"`
}

type InterfaceVeIpRipAuthenticationStr1021 struct {
	String string `json:"string"`
}

type InterfaceVeIpRipAuthenticationMode1022 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceVeIpRipAuthenticationKeyChain1023 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceVeIpRipSendCfg1024 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceVeIpRipReceiveCfg1025 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceVeIpRipSplitHorizonCfg1026 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceVeIpRouter1027 struct {
	Isis InterfaceVeIpRouterIsis1028 `json:"isis"`
}

type InterfaceVeIpRouterIsis1028 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
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
