package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type InterfaceLif struct {
	Inst struct {
		AccessList InterfaceLifAccessList `json:"access-list"`

		Action string `json:"action" dval:"enable"`

		Bfd InterfaceLifBfd674 `json:"bfd"`

		Encapsulation InterfaceLifEncapsulation677 `json:"encapsulation"`

		Ifname string `json:"ifname"`

		Ip InterfaceLifIp679 `json:"ip"`

		Ipv6 InterfaceLifIpv6701 `json:"ipv6"`

		Isis InterfaceLifIsis718 `json:"isis"`

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

type InterfaceLifBfd674 struct {
	Authentication InterfaceLifBfdAuthentication675 `json:"authentication"`
	Echo           int                              `json:"echo"`
	Demand         int                              `json:"demand"`
	IntervalCfg    InterfaceLifBfdIntervalCfg676    `json:"interval-cfg"`
	Uuid           string                           `json:"uuid"`
}

type InterfaceLifBfdAuthentication675 struct {
	KeyId     int    `json:"key-id"`
	Method    string `json:"method"`
	Password  string `json:"password"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLifBfdIntervalCfg676 struct {
	Interval   int `json:"interval"`
	MinRx      int `json:"min-rx"`
	Multiplier int `json:"multiplier"`
}

type InterfaceLifEncapsulation677 struct {
	Dot1q InterfaceLifEncapsulationDot1q678 `json:"dot1q"`
}

type InterfaceLifEncapsulationDot1q678 struct {
	Tag      int    `json:"tag"`
	Ethernet int    `json:"ethernet"`
	Trunk    int    `json:"trunk"`
	Uuid     string `json:"uuid"`
}

type InterfaceLifIp679 struct {
	Dhcp                    int                            `json:"dhcp"`
	AddressList             []InterfaceLifIpAddressList680 `json:"address-list"`
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
	Router                  InterfaceLifIpRouter681        `json:"router"`
	Rip                     InterfaceLifIpRip683           `json:"rip"`
	Ospf                    InterfaceLifIpOspf691          `json:"ospf"`
}

type InterfaceLifIpAddressList680 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceLifIpRouter681 struct {
	Isis InterfaceLifIpRouterIsis682 `json:"isis"`
}

type InterfaceLifIpRouterIsis682 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpRip683 struct {
	Authentication  InterfaceLifIpRipAuthentication684  `json:"authentication"`
	SendPacket      int                                 `json:"send-packet" dval:"1"`
	ReceivePacket   int                                 `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceLifIpRipSendCfg688         `json:"send-cfg"`
	ReceiveCfg      InterfaceLifIpRipReceiveCfg689      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceLifIpRipSplitHorizonCfg690 `json:"split-horizon-cfg"`
	Uuid            string                              `json:"uuid"`
}

type InterfaceLifIpRipAuthentication684 struct {
	Str      InterfaceLifIpRipAuthenticationStr685      `json:"str"`
	Mode     InterfaceLifIpRipAuthenticationMode686     `json:"mode"`
	KeyChain InterfaceLifIpRipAuthenticationKeyChain687 `json:"key-chain"`
}

type InterfaceLifIpRipAuthenticationStr685 struct {
	String string `json:"string"`
}

type InterfaceLifIpRipAuthenticationMode686 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceLifIpRipAuthenticationKeyChain687 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceLifIpRipSendCfg688 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceLifIpRipReceiveCfg689 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceLifIpRipSplitHorizonCfg690 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLifIpOspf691 struct {
	OspfGlobal InterfaceLifIpOspfOspfGlobal692   `json:"ospf-global"`
	OspfIpList []InterfaceLifIpOspfOspfIpList699 `json:"ospf-ip-list"`
}

type InterfaceLifIpOspfOspfGlobal692 struct {
	AuthenticationCfg  InterfaceLifIpOspfOspfGlobalAuthenticationCfg693  `json:"authentication-cfg"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	BfdCfg             InterfaceLifIpOspfOspfGlobalBfdCfg694             `json:"bfd-cfg"`
	Cost               int                                               `json:"cost"`
	DatabaseFilterCfg  InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg695  `json:"database-filter-cfg"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	Disable            string                                            `json:"disable"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLifIpOspfOspfGlobalMessageDigestCfg696 `json:"message-digest-cfg"`
	Mtu                int                                               `json:"mtu"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Network            InterfaceLifIpOspfOspfGlobalNetwork698            `json:"network"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceLifIpOspfOspfGlobalAuthenticationCfg693 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceLifIpOspfOspfGlobalBfdCfg694 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg695 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceLifIpOspfOspfGlobalMessageDigestCfg696 struct {
	MessageDigestKey int                                                `json:"message-digest-key"`
	Md5              InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5697 `json:"md5"`
}

type InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5697 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLifIpOspfOspfGlobalNetwork698 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceLifIpOspfOspfIpList699 struct {
	IpAddr             string                                            `json:"ip-addr"`
	Authentication     int                                               `json:"authentication"`
	Value              string                                            `json:"value"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	Cost               int                                               `json:"cost"`
	DatabaseFilter     string                                            `json:"database-filter"`
	Out                int                                               `json:"out"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLifIpOspfOspfIpListMessageDigestCfg700 `json:"message-digest-cfg"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceLifIpOspfOspfIpListMessageDigestCfg700 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceLifIpv6701 struct {
	AddressList []InterfaceLifIpv6AddressList702 `json:"address-list"`
	Ipv6Enable  int                              `json:"ipv6-enable"`
	Inside      int                              `json:"inside"`
	Outside     int                              `json:"outside"`
	Uuid        string                           `json:"uuid"`
	Router      InterfaceLifIpv6Router703        `json:"router"`
	Ospf        InterfaceLifIpv6Ospf708          `json:"ospf"`
}

type InterfaceLifIpv6AddressList702 struct {
	Ipv6Addr  string `json:"ipv6-addr"`
	Anycast   int    `json:"anycast"`
	LinkLocal int    `json:"link-local"`
}

type InterfaceLifIpv6Router703 struct {
	Ripng InterfaceLifIpv6RouterRipng704 `json:"ripng"`
	Ospf  InterfaceLifIpv6RouterOspf705  `json:"ospf"`
	Isis  InterfaceLifIpv6RouterIsis707  `json:"isis"`
}

type InterfaceLifIpv6RouterRipng704 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpv6RouterOspf705 struct {
	AreaList []InterfaceLifIpv6RouterOspfAreaList706 `json:"area-list"`
	Uuid     string                                  `json:"uuid"`
}

type InterfaceLifIpv6RouterOspfAreaList706 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceLifIpv6RouterIsis707 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpv6Ospf708 struct {
	NetworkList           []InterfaceLifIpv6OspfNetworkList709           `json:"network-list"`
	Bfd                   int                                            `json:"bfd"`
	Disable               int                                            `json:"disable"`
	CostCfg               []InterfaceLifIpv6OspfCostCfg710               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceLifIpv6OspfDeadIntervalCfg711       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceLifIpv6OspfHelloIntervalCfg712      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceLifIpv6OspfMtuIgnoreCfg713          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceLifIpv6OspfNeighborCfg714           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceLifIpv6OspfPriorityCfg715           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceLifIpv6OspfRetransmitIntervalCfg716 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceLifIpv6OspfTransmitDelayCfg717      `json:"transmit-delay-cfg"`
	Uuid                  string                                         `json:"uuid"`
}

type InterfaceLifIpv6OspfNetworkList709 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceLifIpv6OspfCostCfg710 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfDeadIntervalCfg711 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceLifIpv6OspfHelloIntervalCfg712 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLifIpv6OspfMtuIgnoreCfg713 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfNeighborCfg714 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceLifIpv6OspfPriorityCfg715 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfRetransmitIntervalCfg716 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceLifIpv6OspfTransmitDelayCfg717 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLifIsis718 struct {
	Authentication           InterfaceLifIsisAuthentication719             `json:"authentication"`
	BfdCfg                   InterfaceLifIsisBfdCfg723                     `json:"bfd-cfg"`
	CircuitType              string                                        `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceLifIsisCsnpIntervalList724         `json:"csnp-interval-list"`
	Padding                  int                                           `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceLifIsisHelloIntervalList725        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceLifIsisHelloIntervalMinimalList726 `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceLifIsisHelloMultiplierList727      `json:"hello-multiplier-list"`
	LspInterval              int                                           `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceLifIsisMeshGroup728                  `json:"mesh-group"`
	MetricList               []InterfaceLifIsisMetricList729               `json:"metric-list"`
	Network                  string                                        `json:"network"`
	PasswordList             []InterfaceLifIsisPasswordList730             `json:"password-list"`
	PriorityList             []InterfaceLifIsisPriorityList731             `json:"priority-list"`
	RetransmitInterval       int                                           `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceLifIsisWideMetricList732           `json:"wide-metric-list"`
	Uuid                     string                                        `json:"uuid"`
}

type InterfaceLifIsisAuthentication719 struct {
	SendOnlyList []InterfaceLifIsisAuthenticationSendOnlyList720 `json:"send-only-list"`
	ModeList     []InterfaceLifIsisAuthenticationModeList721     `json:"mode-list"`
	KeyChainList []InterfaceLifIsisAuthenticationKeyChainList722 `json:"key-chain-list"`
}

type InterfaceLifIsisAuthenticationSendOnlyList720 struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceLifIsisAuthenticationModeList721 struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceLifIsisAuthenticationKeyChainList722 struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceLifIsisBfdCfg723 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLifIsisCsnpIntervalList724 struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceLifIsisHelloIntervalList725 struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceLifIsisHelloIntervalMinimalList726 struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceLifIsisHelloMultiplierList727 struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceLifIsisMeshGroup728 struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceLifIsisMetricList729 struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceLifIsisPasswordList730 struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceLifIsisPriorityList731 struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceLifIsisWideMetricList732 struct {
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
