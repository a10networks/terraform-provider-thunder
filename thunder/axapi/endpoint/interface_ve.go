package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

//based on ACOS 5_2_1-P4_90
type InterfaceVe struct {
	Inst struct {
		AccessList         InterfaceVeAccessList       `json:"access-list"`
		Action             string                      `json:"action" dval:"enable"`
		Bfd                InterfaceVeBfd              `json:"bfd"`
		Ddos               InterfaceVeDdos             `json:"ddos"`
		IcmpRateLimit      InterfaceVeIcmpRateLimit    `json:"icmp-rate-limit"`
		Icmpv6RateLimit    InterfaceVeIcmpv6RateLimit  `json:"icmpv6-rate-limit"`
		Ifnum              int                         `json:"ifnum"`
		Ip                 InterfaceVeIp               `json:"ip"`
		Ipv6               InterfaceVeIpv6             `json:"ipv6"`
		Isis               InterfaceVeIsis             `json:"isis"`
		L3VlanFwdDisable   int                         `json:"l3-vlan-fwd-disable"`
		Lw4o6              InterfaceVeLw4o6            `json:"lw-4o6"`
		Map                InterfaceVeMap              `json:"map"`
		Mtu                int                         `json:"mtu"`
		Name               string                      `json:"name"`
		Nptv6              InterfaceVeNptv6            `json:"nptv6"`
		PingSweepDetection string                      `json:"ping-sweep-detection" dval:"disable"`
		PortScanDetection  string                      `json:"port-scan-detection" dval:"disable"`
		SamplingEnable     []InterfaceVeSamplingEnable `json:"sampling-enable"`
		TrapSource         int                         `json:"trap-source"`
		UserTag            string                      `json:"user-tag"`
		Uuid               string                      `json:"uuid"`
	} `json:"ve"`
}

type InterfaceVeAccessList struct {
	AclId   int    `json:"acl-id"`
	AclName string `json:"acl-name"`
}

type InterfaceVeBfd struct {
	Authentication InterfaceVeBfdAuthentication `json:"authentication"`
	Echo           int                          `json:"echo"`
	Demand         int                          `json:"demand"`
	IntervalCfg    InterfaceVeBfdIntervalCfg    `json:"interval-cfg"`
	Uuid           string                       `json:"uuid"`
}

type InterfaceVeBfdAuthentication struct {
	KeyId     int    `json:"key-id"`
	Method    string `json:"method"`
	Password  string `json:"password"`
	Encrypted string `json:"encrypted"`
}

type InterfaceVeBfdIntervalCfg struct {
	Interval   int `json:"interval"`
	MinRx      int `json:"min-rx"`
	Multiplier int `json:"multiplier"`
}

type InterfaceVeDdos struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceVeIcmpRateLimit struct {
	Normal       int `json:"normal"`
	Lockup       int `json:"lockup"`
	LockupPeriod int `json:"lockup-period"`
}

type InterfaceVeIcmpv6RateLimit struct {
	NormalV6       int `json:"normal-v6"`
	LockupV6       int `json:"lockup-v6"`
	LockupPeriodV6 int `json:"lockup-period-v6"`
}

type InterfaceVeIp struct {
	Dhcp                    int                              `json:"dhcp"`
	AddressList             []InterfaceVeIpAddressList       `json:"address-list"`
	AllowPromiscuousVip     int                              `json:"allow-promiscuous-vip"`
	Client                  int                              `json:"client"`
	Server                  int                              `json:"server"`
	HelperAddressList       []InterfaceVeIpHelperAddressList `json:"helper-address-list"`
	Inside                  int                              `json:"inside"`
	Outside                 int                              `json:"outside"`
	TtlIgnore               int                              `json:"ttl-ignore"`
	SynCookie               int                              `json:"syn-cookie"`
	SlbPartitionRedirect    int                              `json:"slb-partition-redirect"`
	GenerateMembershipQuery int                              `json:"generate-membership-query"`
	QueryInterval           int                              `json:"query-interval" dval:"125"`
	MaxRespTime             int                              `json:"max-resp-time" dval:"100"`
	Uuid                    string                           `json:"uuid"`
	StatefulFirewall        InterfaceVeIpStatefulFirewall    `json:"stateful-firewall"`
	Router                  InterfaceVeIpRouter              `json:"router"`
	Rip                     InterfaceVeIpRip                 `json:"rip"`
	Ospf                    InterfaceVeIpOspf                `json:"ospf"`
}

type InterfaceVeIpAddressList struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceVeIpHelperAddressList struct {
	HelperAddress string `json:"helper-address"`
}

type InterfaceVeIpStatefulFirewall struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclId      int    `json:"acl-id"`
	Uuid       string `json:"uuid"`
}

type InterfaceVeIpRouter struct {
	Isis InterfaceVeIpRouterIsis `json:"isis"`
}

type InterfaceVeIpRouterIsis struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceVeIpRip struct {
	Authentication  InterfaceVeIpRipAuthentication  `json:"authentication"`
	SendPacket      int                             `json:"send-packet" dval:"1"`
	ReceivePacket   int                             `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceVeIpRipSendCfg         `json:"send-cfg"`
	ReceiveCfg      InterfaceVeIpRipReceiveCfg      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceVeIpRipSplitHorizonCfg `json:"split-horizon-cfg"`
	Uuid            string                          `json:"uuid"`
}

type InterfaceVeIpRipAuthentication struct {
	Str      InterfaceVeIpRipAuthenticationStr      `json:"str"`
	Mode     InterfaceVeIpRipAuthenticationMode     `json:"mode"`
	KeyChain InterfaceVeIpRipAuthenticationKeyChain `json:"key-chain"`
}

type InterfaceVeIpRipAuthenticationStr struct {
	String string `json:"string"`
}

type InterfaceVeIpRipAuthenticationMode struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceVeIpRipAuthenticationKeyChain struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceVeIpRipSendCfg struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceVeIpRipReceiveCfg struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceVeIpRipSplitHorizonCfg struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceVeIpOspf struct {
	OspfGlobal InterfaceVeIpOspfOspfGlobal   `json:"ospf-global"`
	OspfIpList []InterfaceVeIpOspfOspfIpList `json:"ospf-ip-list"`
}

type InterfaceVeIpOspfOspfGlobal struct {
	AuthenticationCfg  InterfaceVeIpOspfOspfGlobalAuthenticationCfg  `json:"authentication-cfg"`
	AuthenticationKey  string                                        `json:"authentication-key"`
	BfdCfg             InterfaceVeIpOspfOspfGlobalBfdCfg             `json:"bfd-cfg"`
	Cost               int                                           `json:"cost"`
	DatabaseFilterCfg  InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg  `json:"database-filter-cfg"`
	DeadInterval       int                                           `json:"dead-interval" dval:"40"`
	Disable            string                                        `json:"disable"`
	HelloInterval      int                                           `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceVeIpOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg"`
	Mtu                int                                           `json:"mtu"`
	MtuIgnore          int                                           `json:"mtu-ignore"`
	Network            InterfaceVeIpOspfOspfGlobalNetwork            `json:"network"`
	Priority           int                                           `json:"priority" dval:"1"`
	RetransmitInterval int                                           `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                           `json:"transmit-delay" dval:"1"`
	Uuid               string                                        `json:"uuid"`
}

type InterfaceVeIpOspfOspfGlobalAuthenticationCfg struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceVeIpOspfOspfGlobalBfdCfg struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceVeIpOspfOspfGlobalMessageDigestCfg struct {
	MessageDigestKey int                                            `json:"message-digest-key"`
	Md5              InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5 `json:"md5"`
}

type InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceVeIpOspfOspfGlobalNetwork struct {
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

type InterfaceVeIpv6 struct {
	AddressList      []InterfaceVeIpv6AddressList    `json:"address-list"`
	Ipv6Enable       int                             `json:"ipv6-enable"`
	V6AclName        string                          `json:"v6-acl-name"`
	Inbound          int                             `json:"inbound"`
	Inside           int                             `json:"inside"`
	Outside          int                             `json:"outside"`
	TtlIgnore        int                             `json:"ttl-ignore"`
	RouterAdver      InterfaceVeIpv6RouterAdver      `json:"router-adver"`
	Uuid             string                          `json:"uuid"`
	StatefulFirewall InterfaceVeIpv6StatefulFirewall `json:"stateful-firewall"`
	Router           InterfaceVeIpv6Router           `json:"router"`
	Rip              InterfaceVeIpv6Rip              `json:"rip"`
	Ospf             InterfaceVeIpv6Ospf             `json:"ospf"`
}

type InterfaceVeIpv6AddressList struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceVeIpv6RouterAdver struct {
	Action                   string                                 `json:"action" dval:"disable"`
	DefaultLifetime          int                                    `json:"default-lifetime" dval:"1800"`
	HopLimit                 int                                    `json:"hop-limit" dval:"255"`
	MaxInterval              int                                    `json:"max-interval" dval:"600"`
	MinInterval              int                                    `json:"min-interval" dval:"200"`
	RateLimit                int                                    `json:"rate-limit" dval:"100000"`
	ReachableTime            int                                    `json:"reachable-time"`
	RetransmitTimer          int                                    `json:"retransmit-timer"`
	AdverMtuDisable          int                                    `json:"adver-mtu-disable" dval:"1"`
	AdverMtu                 int                                    `json:"adver-mtu"`
	PrefixList               []InterfaceVeIpv6RouterAdverPrefixList `json:"prefix-list"`
	ManagedConfigAction      string                                 `json:"managed-config-action" dval:"disable"`
	OtherConfigAction        string                                 `json:"other-config-action" dval:"disable"`
	AdverVrid                int                                    `json:"adver-vrid"`
	UseFloatingIp            int                                    `json:"use-floating-ip"`
	FloatingIp               string                                 `json:"floating-ip"`
	AdverVridDefault         int                                    `json:"adver-vrid-default"`
	UseFloatingIpDefaultVrid int                                    `json:"use-floating-ip-default-vrid"`
	FloatingIpDefaultVrid    string                                 `json:"floating-ip-default-vrid"`
}

type InterfaceVeIpv6RouterAdverPrefixList struct {
	Prefix            string `json:"prefix"`
	NotAutonomous     int    `json:"not-autonomous"`
	NotOnLink         int    `json:"not-on-link"`
	PreferredLifetime int    `json:"preferred-lifetime" dval:"604800"`
	ValidLifetime     int    `json:"valid-lifetime" dval:"2592000"`
}

type InterfaceVeIpv6StatefulFirewall struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclName    string `json:"acl-name"`
	Uuid       string `json:"uuid"`
}

type InterfaceVeIpv6Router struct {
	Ripng InterfaceVeIpv6RouterRipng `json:"ripng"`
	Ospf  InterfaceVeIpv6RouterOspf  `json:"ospf"`
	Isis  InterfaceVeIpv6RouterIsis  `json:"isis"`
}

type InterfaceVeIpv6RouterRipng struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceVeIpv6RouterOspf struct {
	AreaList []InterfaceVeIpv6RouterOspfAreaList `json:"area-list"`
	Uuid     string                              `json:"uuid"`
}

type InterfaceVeIpv6RouterOspfAreaList struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceVeIpv6RouterIsis struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceVeIpv6Rip struct {
	SplitHorizonCfg InterfaceVeIpv6RipSplitHorizonCfg `json:"split-horizon-cfg"`
	Uuid            string                            `json:"uuid"`
}

type InterfaceVeIpv6RipSplitHorizonCfg struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceVeIpv6Ospf struct {
	NetworkList           []InterfaceVeIpv6OspfNetworkList           `json:"network-list"`
	Bfd                   int                                        `json:"bfd"`
	Disable               int                                        `json:"disable"`
	CostCfg               []InterfaceVeIpv6OspfCostCfg               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceVeIpv6OspfDeadIntervalCfg       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceVeIpv6OspfHelloIntervalCfg      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceVeIpv6OspfMtuIgnoreCfg          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceVeIpv6OspfNeighborCfg           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceVeIpv6OspfPriorityCfg           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceVeIpv6OspfRetransmitIntervalCfg `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceVeIpv6OspfTransmitDelayCfg      `json:"transmit-delay-cfg"`
	Uuid                  string                                     `json:"uuid"`
}

type InterfaceVeIpv6OspfNetworkList struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceVeIpv6OspfCostCfg struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceVeIpv6OspfDeadIntervalCfg struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceVeIpv6OspfHelloIntervalCfg struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceVeIpv6OspfMtuIgnoreCfg struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceVeIpv6OspfNeighborCfg struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceVeIpv6OspfPriorityCfg struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceVeIpv6OspfRetransmitIntervalCfg struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceVeIpv6OspfTransmitDelayCfg struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceVeIsis struct {
	Authentication           InterfaceVeIsisAuthentication             `json:"authentication"`
	BfdCfg                   InterfaceVeIsisBfdCfg                     `json:"bfd-cfg"`
	CircuitType              string                                    `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceVeIsisCsnpIntervalList         `json:"csnp-interval-list"`
	Padding                  int                                       `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceVeIsisHelloIntervalList        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceVeIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceVeIsisHelloMultiplierList      `json:"hello-multiplier-list"`
	LspInterval              int                                       `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceVeIsisMeshGroup                  `json:"mesh-group"`
	MetricList               []InterfaceVeIsisMetricList               `json:"metric-list"`
	Network                  string                                    `json:"network"`
	PasswordList             []InterfaceVeIsisPasswordList             `json:"password-list"`
	PriorityList             []InterfaceVeIsisPriorityList             `json:"priority-list"`
	RetransmitInterval       int                                       `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceVeIsisWideMetricList           `json:"wide-metric-list"`
	Uuid                     string                                    `json:"uuid"`
}

type InterfaceVeIsisAuthentication struct {
	SendOnlyList []InterfaceVeIsisAuthenticationSendOnlyList `json:"send-only-list"`
	ModeList     []InterfaceVeIsisAuthenticationModeList     `json:"mode-list"`
	KeyChainList []InterfaceVeIsisAuthenticationKeyChainList `json:"key-chain-list"`
}

type InterfaceVeIsisAuthenticationSendOnlyList struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceVeIsisAuthenticationModeList struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceVeIsisAuthenticationKeyChainList struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceVeIsisBfdCfg struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceVeIsisCsnpIntervalList struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceVeIsisHelloIntervalList struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceVeIsisHelloIntervalMinimalList struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceVeIsisHelloMultiplierList struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceVeIsisMeshGroup struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceVeIsisMetricList struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceVeIsisPasswordList struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceVeIsisPriorityList struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceVeIsisWideMetricList struct {
	WideMetric int    `json:"wide-metric" dval:"10"`
	Level      string `json:"level"`
}

type InterfaceVeLw4o6 struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceVeMap struct {
	Inside      int    `json:"inside"`
	Outside     int    `json:"outside"`
	MapTInside  int    `json:"map-t-inside"`
	MapTOutside int    `json:"map-t-outside"`
	Uuid        string `json:"uuid"`
}

type InterfaceVeNptv6 struct {
	DomainList []InterfaceVeNptv6DomainList `json:"domain-list"`
}

type InterfaceVeNptv6DomainList struct {
	DomainName string `json:"domain-name"`
	BindType   string `json:"bind-type"`
	Uuid       string `json:"uuid"`
}

type InterfaceVeSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *InterfaceVe) GetId() string {
	return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceVe) getPath() string {
	return "interface/ve"
}

func (p *InterfaceVe) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVe::Post")
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

func (p *InterfaceVe) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVe::Get")
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

func (p *InterfaceVe) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVe::Put")
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

func (p *InterfaceVe) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVe::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
