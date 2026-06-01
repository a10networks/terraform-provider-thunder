package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type InterfaceTunnel struct {
	Inst struct {
		Action string `json:"action" dval:"enable"`

		Ifnum int `json:"ifnum"`

		Ip InterfaceTunnelIp967 `json:"ip"`

		Ipv6 InterfaceTunnelIpv6988 `json:"ipv6"`

		LoadInterval int `json:"load-interval" dval:"300"`

		Lw4o6 InterfaceTunnelLw4o61004 `json:"lw-4o6"`

		Map InterfaceTunnelMap1005 `json:"map"`

		Mtu int `json:"mtu"`

		Name string `json:"name"`

		PacketCaptureTemplate string `json:"packet-capture-template"`

		SamplingEnable []InterfaceTunnelSamplingEnable `json:"sampling-enable"`

		Speed int `json:"speed" dval:"10"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"tunnel"`
}

type InterfaceTunnelIp967 struct {
	Address                    InterfaceTunnelIpAddress968 `json:"address"`
	GenerateMembershipQuery    int                         `json:"generate-membership-query"`
	GenerateMembershipQueryVal int                         `json:"generate-membership-query-val" dval:"125"`
	MaxRespTime                int                         `json:"max-resp-time" dval:"100"`
	Inside                     int                         `json:"inside"`
	Outside                    int                         `json:"outside"`
	Uuid                       string                      `json:"uuid"`
	Rip                        InterfaceTunnelIpRip970     `json:"rip"`
	Ospf                       InterfaceTunnelIpOspf978    `json:"ospf"`
}

type InterfaceTunnelIpAddress968 struct {
	Dhcp  int                                `json:"dhcp"`
	IpCfg []InterfaceTunnelIpAddressIpCfg969 `json:"ip-cfg"`
}

type InterfaceTunnelIpAddressIpCfg969 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceTunnelIpRip970 struct {
	Authentication  InterfaceTunnelIpRipAuthentication971  `json:"authentication"`
	SendPacket      int                                    `json:"send-packet" dval:"1"`
	ReceivePacket   int                                    `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceTunnelIpRipSendCfg975         `json:"send-cfg"`
	ReceiveCfg      InterfaceTunnelIpRipReceiveCfg976      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceTunnelIpRipSplitHorizonCfg977 `json:"split-horizon-cfg"`
	Uuid            string                                 `json:"uuid"`
}

type InterfaceTunnelIpRipAuthentication971 struct {
	Str      InterfaceTunnelIpRipAuthenticationStr972      `json:"str"`
	Mode     InterfaceTunnelIpRipAuthenticationMode973     `json:"mode"`
	KeyChain InterfaceTunnelIpRipAuthenticationKeyChain974 `json:"key-chain"`
}

type InterfaceTunnelIpRipAuthenticationStr972 struct {
	String string `json:"string"`
}

type InterfaceTunnelIpRipAuthenticationMode973 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceTunnelIpRipAuthenticationKeyChain974 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceTunnelIpRipSendCfg975 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceTunnelIpRipReceiveCfg976 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceTunnelIpRipSplitHorizonCfg977 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceTunnelIpOspf978 struct {
	OspfGlobal InterfaceTunnelIpOspfOspfGlobal979   `json:"ospf-global"`
	OspfIpList []InterfaceTunnelIpOspfOspfIpList986 `json:"ospf-ip-list"`
}

type InterfaceTunnelIpOspfOspfGlobal979 struct {
	AuthenticationCfg  InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg980  `json:"authentication-cfg"`
	AuthenticationKey  string                                               `json:"authentication-key"`
	BfdCfg             InterfaceTunnelIpOspfOspfGlobalBfdCfg981             `json:"bfd-cfg"`
	Cost               int                                                  `json:"cost"`
	DatabaseFilterCfg  InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg982  `json:"database-filter-cfg"`
	DeadInterval       int                                                  `json:"dead-interval" dval:"40"`
	Disable            string                                               `json:"disable"`
	HelloInterval      int                                                  `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg983 `json:"message-digest-cfg"`
	Mtu                int                                                  `json:"mtu"`
	MtuIgnore          int                                                  `json:"mtu-ignore"`
	Network            InterfaceTunnelIpOspfOspfGlobalNetwork985            `json:"network"`
	Priority           int                                                  `json:"priority" dval:"1"`
	RetransmitInterval int                                                  `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                  `json:"transmit-delay" dval:"1"`
	Uuid               string                                               `json:"uuid"`
}

type InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg980 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceTunnelIpOspfOspfGlobalBfdCfg981 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg982 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg983 struct {
	MessageDigestKey int                                                   `json:"message-digest-key"`
	Md5              InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5984 `json:"md5"`
}

type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5984 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceTunnelIpOspfOspfGlobalNetwork985 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceTunnelIpOspfOspfIpList986 struct {
	IpAddr             string                                               `json:"ip-addr"`
	Authentication     int                                                  `json:"authentication"`
	Value              string                                               `json:"value"`
	AuthenticationKey  string                                               `json:"authentication-key"`
	Cost               int                                                  `json:"cost"`
	DatabaseFilter     string                                               `json:"database-filter"`
	Out                int                                                  `json:"out"`
	DeadInterval       int                                                  `json:"dead-interval" dval:"40"`
	HelloInterval      int                                                  `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTunnelIpOspfOspfIpListMessageDigestCfg987 `json:"message-digest-cfg"`
	MtuIgnore          int                                                  `json:"mtu-ignore"`
	Priority           int                                                  `json:"priority" dval:"1"`
	RetransmitInterval int                                                  `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                  `json:"transmit-delay" dval:"1"`
	Uuid               string                                               `json:"uuid"`
}

type InterfaceTunnelIpOspfOspfIpListMessageDigestCfg987 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceTunnelIpv6988 struct {
	AddressCfg []InterfaceTunnelIpv6AddressCfg989 `json:"address-cfg"`
	Ipv6Enable int                                `json:"ipv6-enable"`
	Inside     int                                `json:"inside"`
	Outside    int                                `json:"outside"`
	Uuid       string                             `json:"uuid"`
	Router     InterfaceTunnelIpv6Router990       `json:"router"`
	Ospf       InterfaceTunnelIpv6Ospf994         `json:"ospf"`
}

type InterfaceTunnelIpv6AddressCfg989 struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceTunnelIpv6Router990 struct {
	Ripng InterfaceTunnelIpv6RouterRipng991 `json:"ripng"`
	Ospf  InterfaceTunnelIpv6RouterOspf992  `json:"ospf"`
}

type InterfaceTunnelIpv6RouterRipng991 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceTunnelIpv6RouterOspf992 struct {
	AreaList []InterfaceTunnelIpv6RouterOspfAreaList993 `json:"area-list"`
	Uuid     string                                     `json:"uuid"`
}

type InterfaceTunnelIpv6RouterOspfAreaList993 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceTunnelIpv6Ospf994 struct {
	NetworkList           []InterfaceTunnelIpv6OspfNetworkList995            `json:"network-list"`
	Bfd                   int                                                `json:"bfd"`
	Disable               int                                                `json:"disable"`
	CostCfg               []InterfaceTunnelIpv6OspfCostCfg996                `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceTunnelIpv6OspfDeadIntervalCfg997        `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceTunnelIpv6OspfHelloIntervalCfg998       `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceTunnelIpv6OspfMtuIgnoreCfg999           `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceTunnelIpv6OspfNeighborCfg1000           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceTunnelIpv6OspfPriorityCfg1001           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceTunnelIpv6OspfRetransmitIntervalCfg1002 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceTunnelIpv6OspfTransmitDelayCfg1003      `json:"transmit-delay-cfg"`
	Uuid                  string                                             `json:"uuid"`
}

type InterfaceTunnelIpv6OspfNetworkList995 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceTunnelIpv6OspfCostCfg996 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfDeadIntervalCfg997 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfHelloIntervalCfg998 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfMtuIgnoreCfg999 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfNeighborCfg1000 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceTunnelIpv6OspfPriorityCfg1001 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfRetransmitIntervalCfg1002 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfTransmitDelayCfg1003 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTunnelLw4o61004 struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceTunnelMap1005 struct {
	Inside      int    `json:"inside"`
	Outside     int    `json:"outside"`
	MapTInside  int    `json:"map-t-inside"`
	MapTOutside int    `json:"map-t-outside"`
	Uuid        string `json:"uuid"`
}

type InterfaceTunnelSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *InterfaceTunnel) GetId() string {
	return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceTunnel) getPath() string {
	return "interface/tunnel"
}

func (p *InterfaceTunnel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTunnel::Post")
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

func (p *InterfaceTunnel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTunnel::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *InterfaceTunnel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTunnel::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *InterfaceTunnel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTunnel::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
