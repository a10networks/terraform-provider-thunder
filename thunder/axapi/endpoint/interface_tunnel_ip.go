package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type InterfaceTunnelIp struct {
	Inst struct {
		Address InterfaceTunnelIpAddress `json:"address"`

		GenerateMembershipQuery int `json:"generate-membership-query"`

		GenerateMembershipQueryVal int `json:"generate-membership-query-val" dval:"125"`

		Inside int `json:"inside"`

		MaxRespTime int `json:"max-resp-time" dval:"100"`

		Ospf InterfaceTunnelIpOspf942 `json:"ospf"`

		Outside int `json:"outside"`

		Rip InterfaceTunnelIpRip950 `json:"rip"`

		Uuid string `json:"uuid"`

		Ifnum string
	} `json:"ip"`
}

type InterfaceTunnelIpAddress struct {
	Dhcp  int                             `json:"dhcp"`
	IpCfg []InterfaceTunnelIpAddressIpCfg `json:"ip-cfg"`
}

type InterfaceTunnelIpAddressIpCfg struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceTunnelIpOspf942 struct {
	OspfGlobal InterfaceTunnelIpOspfOspfGlobal943 `json:"ospf-global"`
	OspfIpList []InterfaceTunnelIpOspfOspfIpList  `json:"ospf-ip-list"`
}

type InterfaceTunnelIpOspfOspfGlobal943 struct {
	AuthenticationCfg  InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg944  `json:"authentication-cfg"`
	AuthenticationKey  string                                               `json:"authentication-key"`
	BfdCfg             InterfaceTunnelIpOspfOspfGlobalBfdCfg945             `json:"bfd-cfg"`
	Cost               int                                                  `json:"cost"`
	DatabaseFilterCfg  InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg946  `json:"database-filter-cfg"`
	DeadInterval       int                                                  `json:"dead-interval" dval:"40"`
	Disable            string                                               `json:"disable"`
	HelloInterval      int                                                  `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg947 `json:"message-digest-cfg"`
	Mtu                int                                                  `json:"mtu"`
	MtuIgnore          int                                                  `json:"mtu-ignore"`
	Network            InterfaceTunnelIpOspfOspfGlobalNetwork949            `json:"network"`
	Priority           int                                                  `json:"priority" dval:"1"`
	RetransmitInterval int                                                  `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                  `json:"transmit-delay" dval:"1"`
	Uuid               string                                               `json:"uuid"`
}

type InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg944 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceTunnelIpOspfOspfGlobalBfdCfg945 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg946 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg947 struct {
	MessageDigestKey int                                                   `json:"message-digest-key"`
	Md5              InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5948 `json:"md5"`
}

type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5948 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceTunnelIpOspfOspfGlobalNetwork949 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceTunnelIpOspfOspfIpList struct {
	IpAddr             string                                            `json:"ip-addr"`
	Authentication     int                                               `json:"authentication"`
	Value              string                                            `json:"value"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	Cost               int                                               `json:"cost"`
	DatabaseFilter     string                                            `json:"database-filter"`
	Out                int                                               `json:"out"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTunnelIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceTunnelIpOspfOspfIpListMessageDigestCfg struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceTunnelIpRip950 struct {
	Authentication  InterfaceTunnelIpRipAuthentication951  `json:"authentication"`
	SendPacket      int                                    `json:"send-packet" dval:"1"`
	ReceivePacket   int                                    `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceTunnelIpRipSendCfg955         `json:"send-cfg"`
	ReceiveCfg      InterfaceTunnelIpRipReceiveCfg956      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceTunnelIpRipSplitHorizonCfg957 `json:"split-horizon-cfg"`
	Uuid            string                                 `json:"uuid"`
}

type InterfaceTunnelIpRipAuthentication951 struct {
	Str      InterfaceTunnelIpRipAuthenticationStr952      `json:"str"`
	Mode     InterfaceTunnelIpRipAuthenticationMode953     `json:"mode"`
	KeyChain InterfaceTunnelIpRipAuthenticationKeyChain954 `json:"key-chain"`
}

type InterfaceTunnelIpRipAuthenticationStr952 struct {
	String string `json:"string"`
}

type InterfaceTunnelIpRipAuthenticationMode953 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceTunnelIpRipAuthenticationKeyChain954 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceTunnelIpRipSendCfg955 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceTunnelIpRipReceiveCfg956 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceTunnelIpRipSplitHorizonCfg957 struct {
	State string `json:"state" dval:"poisoned"`
}

func (p *InterfaceTunnelIp) GetId() string {
	return "1"
}

func (p *InterfaceTunnelIp) getPath() string {
	return "interface/tunnel/" + p.Inst.Ifnum + "/ip"
}

func (p *InterfaceTunnelIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTunnelIp::Post")
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

func (p *InterfaceTunnelIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTunnelIp::Get")
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
func (p *InterfaceTunnelIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTunnelIp::Put")
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

func (p *InterfaceTunnelIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTunnelIp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
