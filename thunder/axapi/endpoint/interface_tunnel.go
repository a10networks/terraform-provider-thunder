package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type InterfaceTunnel struct {
	Inst struct {
		Action string `json:"action" dval:"enable"`

		Ifnum int `json:"ifnum"`

		Ip InterfaceTunnelIp972 `json:"ip"`

		Ipv6 InterfaceTunnelIpv6993 `json:"ipv6"`

		LoadInterval int `json:"load-interval" dval:"300"`

		Lw4o6 InterfaceTunnelLw4o61009 `json:"lw-4o6"`

		Map InterfaceTunnelMap1010 `json:"map"`

		Mtu int `json:"mtu"`

		Name string `json:"name"`

		PacketCaptureTemplate string `json:"packet-capture-template"`

		SamplingEnable []InterfaceTunnelSamplingEnable `json:"sampling-enable"`

		Speed int `json:"speed" dval:"10"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"tunnel"`
}

type InterfaceTunnelIp972 struct {
	Address                    InterfaceTunnelIpAddress973 `json:"address"`
	GenerateMembershipQuery    int                         `json:"generate-membership-query"`
	GenerateMembershipQueryVal int                         `json:"generate-membership-query-val" dval:"125"`
	MaxRespTime                int                         `json:"max-resp-time" dval:"100"`
	Inside                     int                         `json:"inside"`
	Outside                    int                         `json:"outside"`
	Uuid                       string                      `json:"uuid"`
	Rip                        InterfaceTunnelIpRip975     `json:"rip"`
	Ospf                       InterfaceTunnelIpOspf983    `json:"ospf"`
}

type InterfaceTunnelIpAddress973 struct {
	Dhcp  int                                `json:"dhcp"`
	IpCfg []InterfaceTunnelIpAddressIpCfg974 `json:"ip-cfg"`
}

type InterfaceTunnelIpAddressIpCfg974 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceTunnelIpRip975 struct {
	Authentication  InterfaceTunnelIpRipAuthentication976  `json:"authentication"`
	SendPacket      int                                    `json:"send-packet" dval:"1"`
	ReceivePacket   int                                    `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceTunnelIpRipSendCfg980         `json:"send-cfg"`
	ReceiveCfg      InterfaceTunnelIpRipReceiveCfg981      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceTunnelIpRipSplitHorizonCfg982 `json:"split-horizon-cfg"`
	Uuid            string                                 `json:"uuid"`
}

type InterfaceTunnelIpRipAuthentication976 struct {
	Str      InterfaceTunnelIpRipAuthenticationStr977      `json:"str"`
	Mode     InterfaceTunnelIpRipAuthenticationMode978     `json:"mode"`
	KeyChain InterfaceTunnelIpRipAuthenticationKeyChain979 `json:"key-chain"`
}

type InterfaceTunnelIpRipAuthenticationStr977 struct {
	String string `json:"string"`
}

type InterfaceTunnelIpRipAuthenticationMode978 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceTunnelIpRipAuthenticationKeyChain979 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceTunnelIpRipSendCfg980 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceTunnelIpRipReceiveCfg981 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceTunnelIpRipSplitHorizonCfg982 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceTunnelIpOspf983 struct {
	OspfGlobal InterfaceTunnelIpOspfOspfGlobal984   `json:"ospf-global"`
	OspfIpList []InterfaceTunnelIpOspfOspfIpList991 `json:"ospf-ip-list"`
}

type InterfaceTunnelIpOspfOspfGlobal984 struct {
	AuthenticationCfg  InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg985  `json:"authentication-cfg"`
	AuthenticationKey  string                                               `json:"authentication-key"`
	BfdCfg             InterfaceTunnelIpOspfOspfGlobalBfdCfg986             `json:"bfd-cfg"`
	Cost               int                                                  `json:"cost"`
	DatabaseFilterCfg  InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg987  `json:"database-filter-cfg"`
	DeadInterval       int                                                  `json:"dead-interval" dval:"40"`
	Disable            string                                               `json:"disable"`
	HelloInterval      int                                                  `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg988 `json:"message-digest-cfg"`
	Mtu                int                                                  `json:"mtu"`
	MtuIgnore          int                                                  `json:"mtu-ignore"`
	Network            InterfaceTunnelIpOspfOspfGlobalNetwork990            `json:"network"`
	Priority           int                                                  `json:"priority" dval:"1"`
	RetransmitInterval int                                                  `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                  `json:"transmit-delay" dval:"1"`
	Uuid               string                                               `json:"uuid"`
}

type InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg985 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceTunnelIpOspfOspfGlobalBfdCfg986 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg987 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg988 struct {
	MessageDigestKey int                                                   `json:"message-digest-key"`
	Md5              InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5989 `json:"md5"`
}

type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5989 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceTunnelIpOspfOspfGlobalNetwork990 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceTunnelIpOspfOspfIpList991 struct {
	IpAddr             string                                               `json:"ip-addr"`
	Authentication     int                                                  `json:"authentication"`
	Value              string                                               `json:"value"`
	AuthenticationKey  string                                               `json:"authentication-key"`
	Cost               int                                                  `json:"cost"`
	DatabaseFilter     string                                               `json:"database-filter"`
	Out                int                                                  `json:"out"`
	DeadInterval       int                                                  `json:"dead-interval" dval:"40"`
	HelloInterval      int                                                  `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTunnelIpOspfOspfIpListMessageDigestCfg992 `json:"message-digest-cfg"`
	MtuIgnore          int                                                  `json:"mtu-ignore"`
	Priority           int                                                  `json:"priority" dval:"1"`
	RetransmitInterval int                                                  `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                  `json:"transmit-delay" dval:"1"`
	Uuid               string                                               `json:"uuid"`
}

type InterfaceTunnelIpOspfOspfIpListMessageDigestCfg992 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceTunnelIpv6993 struct {
	AddressCfg []InterfaceTunnelIpv6AddressCfg994 `json:"address-cfg"`
	Ipv6Enable int                                `json:"ipv6-enable"`
	Inside     int                                `json:"inside"`
	Outside    int                                `json:"outside"`
	Uuid       string                             `json:"uuid"`
	Router     InterfaceTunnelIpv6Router995       `json:"router"`
	Ospf       InterfaceTunnelIpv6Ospf999         `json:"ospf"`
}

type InterfaceTunnelIpv6AddressCfg994 struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceTunnelIpv6Router995 struct {
	Ripng InterfaceTunnelIpv6RouterRipng996 `json:"ripng"`
	Ospf  InterfaceTunnelIpv6RouterOspf997  `json:"ospf"`
}

type InterfaceTunnelIpv6RouterRipng996 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceTunnelIpv6RouterOspf997 struct {
	AreaList []InterfaceTunnelIpv6RouterOspfAreaList998 `json:"area-list"`
	Uuid     string                                     `json:"uuid"`
}

type InterfaceTunnelIpv6RouterOspfAreaList998 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceTunnelIpv6Ospf999 struct {
	NetworkList           []InterfaceTunnelIpv6OspfNetworkList1000           `json:"network-list"`
	Bfd                   int                                                `json:"bfd"`
	Disable               int                                                `json:"disable"`
	CostCfg               []InterfaceTunnelIpv6OspfCostCfg1001               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceTunnelIpv6OspfDeadIntervalCfg1002       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceTunnelIpv6OspfHelloIntervalCfg1003      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceTunnelIpv6OspfMtuIgnoreCfg1004          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceTunnelIpv6OspfNeighborCfg1005           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceTunnelIpv6OspfPriorityCfg1006           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceTunnelIpv6OspfRetransmitIntervalCfg1007 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceTunnelIpv6OspfTransmitDelayCfg1008      `json:"transmit-delay-cfg"`
	Uuid                  string                                             `json:"uuid"`
}

type InterfaceTunnelIpv6OspfNetworkList1000 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceTunnelIpv6OspfCostCfg1001 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfDeadIntervalCfg1002 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfHelloIntervalCfg1003 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfMtuIgnoreCfg1004 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfNeighborCfg1005 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceTunnelIpv6OspfPriorityCfg1006 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfRetransmitIntervalCfg1007 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceTunnelIpv6OspfTransmitDelayCfg1008 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTunnelLw4o61009 struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceTunnelMap1010 struct {
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
