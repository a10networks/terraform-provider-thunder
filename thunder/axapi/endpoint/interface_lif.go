package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type InterfaceLif struct {
	Inst struct {
		AccessList InterfaceLifAccessList `json:"access-list"`

		Action string `json:"action" dval:"enable"`

		Bfd InterfaceLifBfd683 `json:"bfd"`

		Encapsulation InterfaceLifEncapsulation686 `json:"encapsulation"`

		Ifname string `json:"ifname"`

		Ip InterfaceLifIp688 `json:"ip"`

		Ipv6 InterfaceLifIpv6710 `json:"ipv6"`

		Isis InterfaceLifIsis727 `json:"isis"`

		Mtu int `json:"mtu"`

		Name string `json:"name"`

		SamplingEnable []InterfaceLifSamplingEnable `json:"sampling-enable"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"lif"`
}

type InterfaceLifAccessList struct {
	AclId   int    `json:"acl-id"`
	AclName string `json:"acl-name"`
}

type InterfaceLifBfd683 struct {
	Authentication InterfaceLifBfdAuthentication684 `json:"authentication"`
	Echo           int                              `json:"echo"`
	Demand         int                              `json:"demand"`
	IntervalCfg    InterfaceLifBfdIntervalCfg685    `json:"interval-cfg"`
	Uuid           string                           `json:"uuid"`
}

type InterfaceLifBfdAuthentication684 struct {
	KeyId     int    `json:"key-id"`
	Method    string `json:"method"`
	Password  string `json:"password"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLifBfdIntervalCfg685 struct {
	Interval   int `json:"interval"`
	MinRx      int `json:"min-rx"`
	Multiplier int `json:"multiplier"`
}

type InterfaceLifEncapsulation686 struct {
	Dot1q InterfaceLifEncapsulationDot1q687 `json:"dot1q"`
}

type InterfaceLifEncapsulationDot1q687 struct {
	Tag      int    `json:"tag"`
	Ethernet int    `json:"ethernet"`
	Trunk    int    `json:"trunk"`
	Uuid     string `json:"uuid"`
}

type InterfaceLifIp688 struct {
	Dhcp                    int                            `json:"dhcp"`
	AddressList             []InterfaceLifIpAddressList689 `json:"address-list"`
	AllowPromiscuousVip     int                            `json:"allow-promiscuous-vip"`
	CacheSpoofingPort       int                            `json:"cache-spoofing-port"`
	Client                  int                            `json:"client"`
	Server                  int                            `json:"server"`
	Dmz                     int                            `json:"dmz"`
	Inside                  int                            `json:"inside"`
	Outside                 int                            `json:"outside"`
	GenerateMembershipQuery int                            `json:"generate-membership-query"`
	QueryInterval           int                            `json:"query-interval" dval:"125"`
	MaxRespTime             int                            `json:"max-resp-time" dval:"100"`
	Unnumbered              int                            `json:"unnumbered"`
	Uuid                    string                         `json:"uuid"`
	Router                  InterfaceLifIpRouter690        `json:"router"`
	Rip                     InterfaceLifIpRip692           `json:"rip"`
	Ospf                    InterfaceLifIpOspf700          `json:"ospf"`
}

type InterfaceLifIpAddressList689 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceLifIpRouter690 struct {
	Isis InterfaceLifIpRouterIsis691 `json:"isis"`
}

type InterfaceLifIpRouterIsis691 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpRip692 struct {
	Authentication  InterfaceLifIpRipAuthentication693  `json:"authentication"`
	SendPacket      int                                 `json:"send-packet" dval:"1"`
	ReceivePacket   int                                 `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceLifIpRipSendCfg697         `json:"send-cfg"`
	ReceiveCfg      InterfaceLifIpRipReceiveCfg698      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceLifIpRipSplitHorizonCfg699 `json:"split-horizon-cfg"`
	Uuid            string                              `json:"uuid"`
}

type InterfaceLifIpRipAuthentication693 struct {
	Str      InterfaceLifIpRipAuthenticationStr694      `json:"str"`
	Mode     InterfaceLifIpRipAuthenticationMode695     `json:"mode"`
	KeyChain InterfaceLifIpRipAuthenticationKeyChain696 `json:"key-chain"`
}

type InterfaceLifIpRipAuthenticationStr694 struct {
	String string `json:"string"`
}

type InterfaceLifIpRipAuthenticationMode695 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceLifIpRipAuthenticationKeyChain696 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceLifIpRipSendCfg697 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceLifIpRipReceiveCfg698 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceLifIpRipSplitHorizonCfg699 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLifIpOspf700 struct {
	OspfGlobal InterfaceLifIpOspfOspfGlobal701   `json:"ospf-global"`
	OspfIpList []InterfaceLifIpOspfOspfIpList708 `json:"ospf-ip-list"`
}

type InterfaceLifIpOspfOspfGlobal701 struct {
	AuthenticationCfg  InterfaceLifIpOspfOspfGlobalAuthenticationCfg702  `json:"authentication-cfg"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	BfdCfg             InterfaceLifIpOspfOspfGlobalBfdCfg703             `json:"bfd-cfg"`
	Cost               int                                               `json:"cost"`
	DatabaseFilterCfg  InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg704  `json:"database-filter-cfg"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	Disable            string                                            `json:"disable"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLifIpOspfOspfGlobalMessageDigestCfg705 `json:"message-digest-cfg"`
	Mtu                int                                               `json:"mtu"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Network            InterfaceLifIpOspfOspfGlobalNetwork707            `json:"network"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceLifIpOspfOspfGlobalAuthenticationCfg702 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceLifIpOspfOspfGlobalBfdCfg703 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg704 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceLifIpOspfOspfGlobalMessageDigestCfg705 struct {
	MessageDigestKey int                                                `json:"message-digest-key"`
	Md5              InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5706 `json:"md5"`
}

type InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5706 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLifIpOspfOspfGlobalNetwork707 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceLifIpOspfOspfIpList708 struct {
	IpAddr             string                                            `json:"ip-addr"`
	Authentication     int                                               `json:"authentication"`
	Value              string                                            `json:"value"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	Cost               int                                               `json:"cost"`
	DatabaseFilter     string                                            `json:"database-filter"`
	Out                int                                               `json:"out"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLifIpOspfOspfIpListMessageDigestCfg709 `json:"message-digest-cfg"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceLifIpOspfOspfIpListMessageDigestCfg709 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceLifIpv6710 struct {
	AddressList []InterfaceLifIpv6AddressList711 `json:"address-list"`
	Ipv6Enable  int                              `json:"ipv6-enable"`
	Inside      int                              `json:"inside"`
	Outside     int                              `json:"outside"`
	Uuid        string                           `json:"uuid"`
	Router      InterfaceLifIpv6Router712        `json:"router"`
	Ospf        InterfaceLifIpv6Ospf717          `json:"ospf"`
}

type InterfaceLifIpv6AddressList711 struct {
	Ipv6Addr  string `json:"ipv6-addr"`
	Anycast   int    `json:"anycast"`
	LinkLocal int    `json:"link-local"`
}

type InterfaceLifIpv6Router712 struct {
	Ripng InterfaceLifIpv6RouterRipng713 `json:"ripng"`
	Ospf  InterfaceLifIpv6RouterOspf714  `json:"ospf"`
	Isis  InterfaceLifIpv6RouterIsis716  `json:"isis"`
}

type InterfaceLifIpv6RouterRipng713 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpv6RouterOspf714 struct {
	AreaList []InterfaceLifIpv6RouterOspfAreaList715 `json:"area-list"`
	Uuid     string                                  `json:"uuid"`
}

type InterfaceLifIpv6RouterOspfAreaList715 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceLifIpv6RouterIsis716 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpv6Ospf717 struct {
	NetworkList           []InterfaceLifIpv6OspfNetworkList718           `json:"network-list"`
	Bfd                   int                                            `json:"bfd"`
	Disable               int                                            `json:"disable"`
	CostCfg               []InterfaceLifIpv6OspfCostCfg719               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceLifIpv6OspfDeadIntervalCfg720       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceLifIpv6OspfHelloIntervalCfg721      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceLifIpv6OspfMtuIgnoreCfg722          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceLifIpv6OspfNeighborCfg723           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceLifIpv6OspfPriorityCfg724           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceLifIpv6OspfRetransmitIntervalCfg725 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceLifIpv6OspfTransmitDelayCfg726      `json:"transmit-delay-cfg"`
	Uuid                  string                                         `json:"uuid"`
}

type InterfaceLifIpv6OspfNetworkList718 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceLifIpv6OspfCostCfg719 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfDeadIntervalCfg720 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceLifIpv6OspfHelloIntervalCfg721 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLifIpv6OspfMtuIgnoreCfg722 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfNeighborCfg723 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceLifIpv6OspfPriorityCfg724 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfRetransmitIntervalCfg725 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceLifIpv6OspfTransmitDelayCfg726 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLifIsis727 struct {
	Authentication           InterfaceLifIsisAuthentication728             `json:"authentication"`
	BfdCfg                   InterfaceLifIsisBfdCfg732                     `json:"bfd-cfg"`
	CircuitType              string                                        `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceLifIsisCsnpIntervalList733         `json:"csnp-interval-list"`
	Padding                  int                                           `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceLifIsisHelloIntervalList734        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceLifIsisHelloIntervalMinimalList735 `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceLifIsisHelloMultiplierList736      `json:"hello-multiplier-list"`
	LspInterval              int                                           `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceLifIsisMeshGroup737                  `json:"mesh-group"`
	MetricList               []InterfaceLifIsisMetricList738               `json:"metric-list"`
	Network                  string                                        `json:"network"`
	PasswordList             []InterfaceLifIsisPasswordList739             `json:"password-list"`
	PriorityList             []InterfaceLifIsisPriorityList740             `json:"priority-list"`
	RetransmitInterval       int                                           `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceLifIsisWideMetricList741           `json:"wide-metric-list"`
	Uuid                     string                                        `json:"uuid"`
}

type InterfaceLifIsisAuthentication728 struct {
	SendOnlyList []InterfaceLifIsisAuthenticationSendOnlyList729 `json:"send-only-list"`
	ModeList     []InterfaceLifIsisAuthenticationModeList730     `json:"mode-list"`
	KeyChainList []InterfaceLifIsisAuthenticationKeyChainList731 `json:"key-chain-list"`
}

type InterfaceLifIsisAuthenticationSendOnlyList729 struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceLifIsisAuthenticationModeList730 struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceLifIsisAuthenticationKeyChainList731 struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceLifIsisBfdCfg732 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLifIsisCsnpIntervalList733 struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceLifIsisHelloIntervalList734 struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceLifIsisHelloIntervalMinimalList735 struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceLifIsisHelloMultiplierList736 struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceLifIsisMeshGroup737 struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceLifIsisMetricList738 struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceLifIsisPasswordList739 struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceLifIsisPriorityList740 struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceLifIsisWideMetricList741 struct {
	WideMetric int    `json:"wide-metric" dval:"10"`
	Level      string `json:"level"`
}

type InterfaceLifSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *InterfaceLif) GetId() string {
	return p.Inst.Ifname
}

func (p *InterfaceLif) getPath() string {
	return "interface/lif"
}

func (p *InterfaceLif) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLif::Post")
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

func (p *InterfaceLif) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLif::Get")
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
func (p *InterfaceLif) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLif::Put")
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

func (p *InterfaceLif) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLif::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
