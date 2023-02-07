package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

//based on ACOS 5_2_1-P4_90
type InterfaceLif struct {
	Inst struct {
		AccessList     InterfaceLifAccessList       `json:"access-list"`
		Action         string                       `json:"action" dval:"enable"`
		Bfd            InterfaceLifBfd              `json:"bfd"`
		Ifname         string                       `json:"ifname"`
		Ip             InterfaceLifIp               `json:"ip"`
		Ipv6           InterfaceLifIpv6             `json:"ipv6"`
		Isis           InterfaceLifIsis             `json:"isis"`
		Mtu            int                          `json:"mtu"`
		Name           string                       `json:"name"`
		SamplingEnable []InterfaceLifSamplingEnable `json:"sampling-enable"`
		UserTag        string                       `json:"user-tag"`
		Uuid           string                       `json:"uuid"`
	} `json:"lif"`
}

type InterfaceLifAccessList struct {
	AclId   int    `json:"acl-id"`
	AclName string `json:"acl-name"`
}

type InterfaceLifBfd struct {
	Authentication InterfaceLifBfdAuthentication `json:"authentication"`
	Echo           int                           `json:"echo"`
	Demand         int                           `json:"demand"`
	IntervalCfg    InterfaceLifBfdIntervalCfg    `json:"interval-cfg"`
	Uuid           string                        `json:"uuid"`
}

type InterfaceLifBfdAuthentication struct {
	KeyId     int    `json:"key-id"`
	Method    string `json:"method"`
	Password  string `json:"password"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLifBfdIntervalCfg struct {
	Interval   int `json:"interval"`
	MinRx      int `json:"min-rx"`
	Multiplier int `json:"multiplier"`
}

type InterfaceLifIp struct {
	Dhcp                    int                         `json:"dhcp"`
	AddressList             []InterfaceLifIpAddressList `json:"address-list"`
	AllowPromiscuousVip     int                         `json:"allow-promiscuous-vip"`
	CacheSpoofingPort       int                         `json:"cache-spoofing-port"`
	Inside                  int                         `json:"inside"`
	Outside                 int                         `json:"outside"`
	GenerateMembershipQuery int                         `json:"generate-membership-query"`
	QueryInterval           int                         `json:"query-interval" dval:"125"`
	MaxRespTime             int                         `json:"max-resp-time" dval:"100"`
	Unnumbered              int                         `json:"unnumbered"`
	Uuid                    string                      `json:"uuid"`
	Router                  InterfaceLifIpRouter        `json:"router"`
	Rip                     InterfaceLifIpRip           `json:"rip"`
	Ospf                    InterfaceLifIpOspf          `json:"ospf"`
}

type InterfaceLifIpAddressList struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceLifIpRouter struct {
	Isis InterfaceLifIpRouterIsis `json:"isis"`
}

type InterfaceLifIpRouterIsis struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpRip struct {
	Authentication  InterfaceLifIpRipAuthentication  `json:"authentication"`
	SendPacket      int                              `json:"send-packet" dval:"1"`
	ReceivePacket   int                              `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceLifIpRipSendCfg         `json:"send-cfg"`
	ReceiveCfg      InterfaceLifIpRipReceiveCfg      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceLifIpRipSplitHorizonCfg `json:"split-horizon-cfg"`
	Uuid            string                           `json:"uuid"`
}

type InterfaceLifIpRipAuthentication struct {
	Str      InterfaceLifIpRipAuthenticationStr      `json:"str"`
	Mode     InterfaceLifIpRipAuthenticationMode     `json:"mode"`
	KeyChain InterfaceLifIpRipAuthenticationKeyChain `json:"key-chain"`
}

type InterfaceLifIpRipAuthenticationStr struct {
	String string `json:"string"`
}

type InterfaceLifIpRipAuthenticationMode struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceLifIpRipAuthenticationKeyChain struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceLifIpRipSendCfg struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceLifIpRipReceiveCfg struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceLifIpRipSplitHorizonCfg struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLifIpOspf struct {
	OspfGlobal InterfaceLifIpOspfOspfGlobal   `json:"ospf-global"`
	OspfIpList []InterfaceLifIpOspfOspfIpList `json:"ospf-ip-list"`
}

type InterfaceLifIpOspfOspfGlobal struct {
	AuthenticationCfg  InterfaceLifIpOspfOspfGlobalAuthenticationCfg  `json:"authentication-cfg"`
	AuthenticationKey  string                                         `json:"authentication-key"`
	BfdCfg             InterfaceLifIpOspfOspfGlobalBfdCfg             `json:"bfd-cfg"`
	Cost               int                                            `json:"cost"`
	DatabaseFilterCfg  InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg  `json:"database-filter-cfg"`
	DeadInterval       int                                            `json:"dead-interval" dval:"40"`
	Disable            string                                         `json:"disable"`
	HelloInterval      int                                            `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLifIpOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg"`
	Mtu                int                                            `json:"mtu"`
	MtuIgnore          int                                            `json:"mtu-ignore"`
	Network            InterfaceLifIpOspfOspfGlobalNetwork            `json:"network"`
	Priority           int                                            `json:"priority" dval:"1"`
	RetransmitInterval int                                            `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                            `json:"transmit-delay" dval:"1"`
	Uuid               string                                         `json:"uuid"`
}

type InterfaceLifIpOspfOspfGlobalAuthenticationCfg struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceLifIpOspfOspfGlobalBfdCfg struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceLifIpOspfOspfGlobalMessageDigestCfg struct {
	MessageDigestKey int                                             `json:"message-digest-key"`
	Md5              InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5 `json:"md5"`
}

type InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLifIpOspfOspfGlobalNetwork struct {
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

type InterfaceLifIpv6 struct {
	AddressList []InterfaceLifIpv6AddressList `json:"address-list"`
	Ipv6Enable  int                           `json:"ipv6-enable"`
	Inside      int                           `json:"inside"`
	Outside     int                           `json:"outside"`
	Uuid        string                        `json:"uuid"`
	Router      InterfaceLifIpv6Router        `json:"router"`
	Ospf        InterfaceLifIpv6Ospf          `json:"ospf"`
}

type InterfaceLifIpv6AddressList struct {
	Ipv6Addr  string `json:"ipv6-addr"`
	Anycast   int    `json:"anycast"`
	LinkLocal int    `json:"link-local"`
}

type InterfaceLifIpv6Router struct {
	Ripng InterfaceLifIpv6RouterRipng `json:"ripng"`
	Ospf  InterfaceLifIpv6RouterOspf  `json:"ospf"`
	Isis  InterfaceLifIpv6RouterIsis  `json:"isis"`
}

type InterfaceLifIpv6RouterRipng struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpv6RouterOspf struct {
	AreaList []InterfaceLifIpv6RouterOspfAreaList `json:"area-list"`
	Uuid     string                               `json:"uuid"`
}

type InterfaceLifIpv6RouterOspfAreaList struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceLifIpv6RouterIsis struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpv6Ospf struct {
	NetworkList           []InterfaceLifIpv6OspfNetworkList           `json:"network-list"`
	Bfd                   int                                         `json:"bfd"`
	Disable               int                                         `json:"disable"`
	CostCfg               []InterfaceLifIpv6OspfCostCfg               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceLifIpv6OspfDeadIntervalCfg       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceLifIpv6OspfHelloIntervalCfg      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceLifIpv6OspfMtuIgnoreCfg          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceLifIpv6OspfNeighborCfg           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceLifIpv6OspfPriorityCfg           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceLifIpv6OspfRetransmitIntervalCfg `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceLifIpv6OspfTransmitDelayCfg      `json:"transmit-delay-cfg"`
	Uuid                  string                                      `json:"uuid"`
}

type InterfaceLifIpv6OspfNetworkList struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceLifIpv6OspfCostCfg struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfDeadIntervalCfg struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceLifIpv6OspfHelloIntervalCfg struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLifIpv6OspfMtuIgnoreCfg struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfNeighborCfg struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceLifIpv6OspfPriorityCfg struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfRetransmitIntervalCfg struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceLifIpv6OspfTransmitDelayCfg struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLifIsis struct {
	Authentication           InterfaceLifIsisAuthentication             `json:"authentication"`
	BfdCfg                   InterfaceLifIsisBfdCfg                     `json:"bfd-cfg"`
	CircuitType              string                                     `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceLifIsisCsnpIntervalList         `json:"csnp-interval-list"`
	Padding                  int                                        `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceLifIsisHelloIntervalList        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceLifIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceLifIsisHelloMultiplierList      `json:"hello-multiplier-list"`
	LspInterval              int                                        `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceLifIsisMeshGroup                  `json:"mesh-group"`
	MetricList               []InterfaceLifIsisMetricList               `json:"metric-list"`
	Network                  string                                     `json:"network"`
	PasswordList             []InterfaceLifIsisPasswordList             `json:"password-list"`
	PriorityList             []InterfaceLifIsisPriorityList             `json:"priority-list"`
	RetransmitInterval       int                                        `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceLifIsisWideMetricList           `json:"wide-metric-list"`
	Uuid                     string                                     `json:"uuid"`
}

type InterfaceLifIsisAuthentication struct {
	SendOnlyList []InterfaceLifIsisAuthenticationSendOnlyList `json:"send-only-list"`
	ModeList     []InterfaceLifIsisAuthenticationModeList     `json:"mode-list"`
	KeyChainList []InterfaceLifIsisAuthenticationKeyChainList `json:"key-chain-list"`
}

type InterfaceLifIsisAuthenticationSendOnlyList struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceLifIsisAuthenticationModeList struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceLifIsisAuthenticationKeyChainList struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceLifIsisBfdCfg struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLifIsisCsnpIntervalList struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceLifIsisHelloIntervalList struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceLifIsisHelloIntervalMinimalList struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceLifIsisHelloMultiplierList struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceLifIsisMeshGroup struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceLifIsisMetricList struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceLifIsisPasswordList struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceLifIsisPriorityList struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceLifIsisWideMetricList struct {
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
		err = json.Unmarshal(axResp, &p)
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
