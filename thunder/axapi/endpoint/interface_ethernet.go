package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type InterfaceEthernet struct {
	Inst struct {
		AccessList InterfaceEthernetAccessList `json:"access-list"`

		Action string `json:"action" dval:"disable"`

		AutoNegEnable int `json:"auto-neg-enable"`

		Bfd InterfaceEthernetBfd565 `json:"bfd"`

		CpuProcess int `json:"cpu-process"`

		CpuProcessDir string `json:"cpu-process-dir"`

		DacLinkTrainingEnable int `json:"dac-link-training-enable"`

		Ddos InterfaceEthernetDdos568 `json:"ddos"`

		Duplexity string `json:"duplexity" dval:"auto"`

		FecForcedOff int `json:"fec-forced-off"`

		FecForcedOn int `json:"fec-forced-on"`

		FlowControl int `json:"flow-control"`

		GamingProtocolCompliance int `json:"gaming-protocol-compliance"`

		IcmpRateLimit InterfaceEthernetIcmpRateLimit `json:"icmp-rate-limit"`

		Icmpv6RateLimit InterfaceEthernetIcmpv6RateLimit `json:"icmpv6-rate-limit"`

		Ifnum int `json:"ifnum"`

		Ip InterfaceEthernetIp569 `json:"ip"`

		IpgBitTime int `json:"ipg-bit-time" dval:"96"`

		Ipv6 InterfaceEthernetIpv6593 `json:"ipv6"`

		Isis InterfaceEthernetIsis616 `json:"isis"`

		L3VlanFwdDisable int `json:"l3-vlan-fwd-disable"`

		Lldp InterfaceEthernetLldp631 `json:"lldp"`

		LoadInterval int `json:"load-interval" dval:"300"`

		Lw4o6 InterfaceEthernetLw4o6636 `json:"lw-4o6"`

		MacLearning string `json:"mac-learning"`

		Map InterfaceEthernetMap637 `json:"map"`

		MediaTypeCopper int `json:"media-type-copper"`

		MonitorList []InterfaceEthernetMonitorList `json:"monitor-list"`

		Mtu int `json:"mtu"`

		Name string `json:"name"`

		Nptv6 InterfaceEthernetNptv6638 `json:"nptv6"`

		PacketCaptureTemplate string `json:"packet-capture-template"`

		PingSweepDetection string `json:"ping-sweep-detection" dval:"disable"`

		PortBreakout string `json:"port-breakout"`

		PortScanDetection string `json:"port-scan-detection" dval:"disable"`

		RemoveVlanTag int `json:"remove-vlan-tag"`

		SamplingEnable []InterfaceEthernetSamplingEnable `json:"sampling-enable"`

		SpanningTree InterfaceEthernetSpanningTree639 `json:"spanning-tree"`

		Speed string `json:"speed" dval:"auto"`

		SpeedForced10g int `json:"speed-forced-10g"`

		SpeedForced1g int `json:"speed-forced-1g"`

		SpeedForced40g int `json:"speed-forced-40g"`

		TrafficDistributionMode string `json:"traffic-distribution-mode"`

		TrapSource int `json:"trap-source"`

		TrunkGroupList []InterfaceEthernetTrunkGroupList `json:"trunk-group-list"`

		UpdateL2Info int `json:"update-l2-info"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		VirtualWire int `json:"virtual-wire"`

		VlanLearning string `json:"vlan-learning"`
	} `json:"ethernet"`
}

type InterfaceEthernetAccessList struct {
	AclId   int    `json:"acl-id"`
	AclName string `json:"acl-name"`
}

type InterfaceEthernetBfd565 struct {
	Authentication InterfaceEthernetBfdAuthentication566 `json:"authentication"`
	Echo           int                                   `json:"echo"`
	Demand         int                                   `json:"demand"`
	IntervalCfg    InterfaceEthernetBfdIntervalCfg567    `json:"interval-cfg"`
	Uuid           string                                `json:"uuid"`
}

type InterfaceEthernetBfdAuthentication566 struct {
	KeyId     int    `json:"key-id"`
	Method    string `json:"method"`
	Password  string `json:"password"`
	Encrypted string `json:"encrypted"`
}

type InterfaceEthernetBfdIntervalCfg567 struct {
	Interval   int `json:"interval"`
	MinRx      int `json:"min-rx"`
	Multiplier int `json:"multiplier"`
}

type InterfaceEthernetDdos568 struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceEthernetIcmpRateLimit struct {
	Normal       int `json:"normal"`
	Lockup       int `json:"lockup"`
	LockupPeriod int `json:"lockup-period"`
}

type InterfaceEthernetIcmpv6RateLimit struct {
	NormalV6       int `json:"normal-v6"`
	LockupV6       int `json:"lockup-v6"`
	LockupPeriodV6 int `json:"lockup-period-v6"`
}

type InterfaceEthernetIp569 struct {
	Dhcp                    int                                       `json:"dhcp"`
	AddressList             []InterfaceEthernetIpAddressList570       `json:"address-list"`
	AllowPromiscuousVip     int                                       `json:"allow-promiscuous-vip"`
	CacheSpoofingPort       int                                       `json:"cache-spoofing-port"`
	HelperAddressList       []InterfaceEthernetIpHelperAddressList571 `json:"helper-address-list"`
	Inside                  int                                       `json:"inside"`
	Outside                 int                                       `json:"outside"`
	TtlIgnore               int                                       `json:"ttl-ignore"`
	SynCookie               int                                       `json:"syn-cookie"`
	SlbPartitionRedirect    int                                       `json:"slb-partition-redirect"`
	GenerateMembershipQuery int                                       `json:"generate-membership-query"`
	QueryInterval           int                                       `json:"query-interval" dval:"125"`
	MaxRespTime             int                                       `json:"max-resp-time" dval:"100"`
	Client                  int                                       `json:"client"`
	Server                  int                                       `json:"server"`
	Dmz                     int                                       `json:"dmz"`
	Unnumbered              int                                       `json:"unnumbered"`
	Uuid                    string                                    `json:"uuid"`
	StatefulFirewall        InterfaceEthernetIpStatefulFirewall572    `json:"stateful-firewall"`
	Router                  InterfaceEthernetIpRouter573              `json:"router"`
	Rip                     InterfaceEthernetIpRip575                 `json:"rip"`
	Ospf                    InterfaceEthernetIpOspf583                `json:"ospf"`
}

type InterfaceEthernetIpAddressList570 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceEthernetIpHelperAddressList571 struct {
	HelperAddress string `json:"helper-address"`
}

type InterfaceEthernetIpStatefulFirewall572 struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclId      int    `json:"acl-id"`
	Uuid       string `json:"uuid"`
}

type InterfaceEthernetIpRouter573 struct {
	Isis InterfaceEthernetIpRouterIsis574 `json:"isis"`
}

type InterfaceEthernetIpRouterIsis574 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceEthernetIpRip575 struct {
	Authentication  InterfaceEthernetIpRipAuthentication576  `json:"authentication"`
	SendPacket      int                                      `json:"send-packet" dval:"1"`
	ReceivePacket   int                                      `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceEthernetIpRipSendCfg580         `json:"send-cfg"`
	ReceiveCfg      InterfaceEthernetIpRipReceiveCfg581      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceEthernetIpRipSplitHorizonCfg582 `json:"split-horizon-cfg"`
	Uuid            string                                   `json:"uuid"`
}

type InterfaceEthernetIpRipAuthentication576 struct {
	Str      InterfaceEthernetIpRipAuthenticationStr577      `json:"str"`
	Mode     InterfaceEthernetIpRipAuthenticationMode578     `json:"mode"`
	KeyChain InterfaceEthernetIpRipAuthenticationKeyChain579 `json:"key-chain"`
}

type InterfaceEthernetIpRipAuthenticationStr577 struct {
	String string `json:"string"`
}

type InterfaceEthernetIpRipAuthenticationMode578 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceEthernetIpRipAuthenticationKeyChain579 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceEthernetIpRipSendCfg580 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceEthernetIpRipReceiveCfg581 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceEthernetIpRipSplitHorizonCfg582 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceEthernetIpOspf583 struct {
	OspfGlobal InterfaceEthernetIpOspfOspfGlobal584   `json:"ospf-global"`
	OspfIpList []InterfaceEthernetIpOspfOspfIpList591 `json:"ospf-ip-list"`
}

type InterfaceEthernetIpOspfOspfGlobal584 struct {
	AuthenticationCfg  InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg585  `json:"authentication-cfg"`
	AuthenticationKey  string                                                 `json:"authentication-key"`
	BfdCfg             InterfaceEthernetIpOspfOspfGlobalBfdCfg586             `json:"bfd-cfg"`
	Cost               int                                                    `json:"cost"`
	DatabaseFilterCfg  InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg587  `json:"database-filter-cfg"`
	DeadInterval       int                                                    `json:"dead-interval" dval:"40"`
	Disable            string                                                 `json:"disable"`
	HelloInterval      int                                                    `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg588 `json:"message-digest-cfg"`
	Mtu                int                                                    `json:"mtu"`
	MtuIgnore          int                                                    `json:"mtu-ignore"`
	Network            InterfaceEthernetIpOspfOspfGlobalNetwork590            `json:"network"`
	Priority           int                                                    `json:"priority" dval:"1"`
	RetransmitInterval int                                                    `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                    `json:"transmit-delay" dval:"1"`
	Uuid               string                                                 `json:"uuid"`
}

type InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg585 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceEthernetIpOspfOspfGlobalBfdCfg586 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg587 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg588 struct {
	MessageDigestKey int                                                     `json:"message-digest-key"`
	Md5              InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5589 `json:"md5"`
}

type InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5589 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceEthernetIpOspfOspfGlobalNetwork590 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceEthernetIpOspfOspfIpList591 struct {
	IpAddr             string                                                 `json:"ip-addr"`
	Authentication     int                                                    `json:"authentication"`
	Value              string                                                 `json:"value"`
	AuthenticationKey  string                                                 `json:"authentication-key"`
	Cost               int                                                    `json:"cost"`
	DatabaseFilter     string                                                 `json:"database-filter"`
	Out                int                                                    `json:"out"`
	DeadInterval       int                                                    `json:"dead-interval" dval:"40"`
	HelloInterval      int                                                    `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceEthernetIpOspfOspfIpListMessageDigestCfg592 `json:"message-digest-cfg"`
	MtuIgnore          int                                                    `json:"mtu-ignore"`
	Priority           int                                                    `json:"priority" dval:"1"`
	RetransmitInterval int                                                    `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                    `json:"transmit-delay" dval:"1"`
	Uuid               string                                                 `json:"uuid"`
}

type InterfaceEthernetIpOspfOspfIpListMessageDigestCfg592 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceEthernetIpv6593 struct {
	AddressList      []InterfaceEthernetIpv6AddressList594    `json:"address-list"`
	Inside           int                                      `json:"inside"`
	Outside          int                                      `json:"outside"`
	Ipv6Enable       int                                      `json:"ipv6-enable"`
	TtlIgnore        int                                      `json:"ttl-ignore"`
	AccessListCfg    InterfaceEthernetIpv6AccessListCfg595    `json:"access-list-cfg"`
	RouterAdver      InterfaceEthernetIpv6RouterAdver596      `json:"router-adver"`
	Uuid             string                                   `json:"uuid"`
	StatefulFirewall InterfaceEthernetIpv6StatefulFirewall598 `json:"stateful-firewall"`
	Router           InterfaceEthernetIpv6Router599           `json:"router"`
	Rip              InterfaceEthernetIpv6Rip604              `json:"rip"`
	Ospf             InterfaceEthernetIpv6Ospf606             `json:"ospf"`
}

type InterfaceEthernetIpv6AddressList594 struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceEthernetIpv6AccessListCfg595 struct {
	V6AclName string `json:"v6-acl-name"`
	Inbound   int    `json:"inbound"`
}

type InterfaceEthernetIpv6RouterAdver596 struct {
	Action                   string                                          `json:"action" dval:"disable"`
	HopLimit                 int                                             `json:"hop-limit" dval:"255"`
	MaxInterval              int                                             `json:"max-interval" dval:"600"`
	MinInterval              int                                             `json:"min-interval" dval:"200"`
	DefaultLifetime          int                                             `json:"default-lifetime" dval:"1800"`
	RateLimit                int                                             `json:"rate-limit" dval:"100000"`
	ReachableTime            int                                             `json:"reachable-time"`
	RetransmitTimer          int                                             `json:"retransmit-timer"`
	AdverMtuDisable          int                                             `json:"adver-mtu-disable" dval:"1"`
	AdverMtu                 int                                             `json:"adver-mtu"`
	PrefixList               []InterfaceEthernetIpv6RouterAdverPrefixList597 `json:"prefix-list"`
	ManagedConfigAction      string                                          `json:"managed-config-action" dval:"disable"`
	OtherConfigAction        string                                          `json:"other-config-action" dval:"disable"`
	AdverVrid                int                                             `json:"adver-vrid"`
	UseFloatingIp            int                                             `json:"use-floating-ip"`
	FloatingIp               string                                          `json:"floating-ip"`
	AdverVridDefault         int                                             `json:"adver-vrid-default"`
	UseFloatingIpDefaultVrid int                                             `json:"use-floating-ip-default-vrid"`
	FloatingIpDefaultVrid    string                                          `json:"floating-ip-default-vrid"`
}

type InterfaceEthernetIpv6RouterAdverPrefixList597 struct {
	Prefix            string `json:"prefix"`
	NotAutonomous     int    `json:"not-autonomous"`
	NotOnLink         int    `json:"not-on-link"`
	PreferredLifetime int    `json:"preferred-lifetime" dval:"604800"`
	ValidLifetime     int    `json:"valid-lifetime" dval:"2592000"`
}

type InterfaceEthernetIpv6StatefulFirewall598 struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclName    string `json:"acl-name"`
	Uuid       string `json:"uuid"`
}

type InterfaceEthernetIpv6Router599 struct {
	Ripng InterfaceEthernetIpv6RouterRipng600 `json:"ripng"`
	Ospf  InterfaceEthernetIpv6RouterOspf601  `json:"ospf"`
	Isis  InterfaceEthernetIpv6RouterIsis603  `json:"isis"`
}

type InterfaceEthernetIpv6RouterRipng600 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceEthernetIpv6RouterOspf601 struct {
	AreaList []InterfaceEthernetIpv6RouterOspfAreaList602 `json:"area-list"`
	Uuid     string                                       `json:"uuid"`
}

type InterfaceEthernetIpv6RouterOspfAreaList602 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceEthernetIpv6RouterIsis603 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceEthernetIpv6Rip604 struct {
	SplitHorizonCfg InterfaceEthernetIpv6RipSplitHorizonCfg605 `json:"split-horizon-cfg"`
	Uuid            string                                     `json:"uuid"`
}

type InterfaceEthernetIpv6RipSplitHorizonCfg605 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceEthernetIpv6Ospf606 struct {
	NetworkList           []InterfaceEthernetIpv6OspfNetworkList607           `json:"network-list"`
	Bfd                   int                                                 `json:"bfd"`
	Disable               int                                                 `json:"disable"`
	CostCfg               []InterfaceEthernetIpv6OspfCostCfg608               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceEthernetIpv6OspfDeadIntervalCfg609       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceEthernetIpv6OspfHelloIntervalCfg610      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceEthernetIpv6OspfMtuIgnoreCfg611          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceEthernetIpv6OspfNeighborCfg612           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceEthernetIpv6OspfPriorityCfg613           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceEthernetIpv6OspfRetransmitIntervalCfg614 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceEthernetIpv6OspfTransmitDelayCfg615      `json:"transmit-delay-cfg"`
	Uuid                  string                                              `json:"uuid"`
}

type InterfaceEthernetIpv6OspfNetworkList607 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceEthernetIpv6OspfCostCfg608 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfDeadIntervalCfg609 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfHelloIntervalCfg610 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfMtuIgnoreCfg611 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfNeighborCfg612 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceEthernetIpv6OspfPriorityCfg613 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfRetransmitIntervalCfg614 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfTransmitDelayCfg615 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceEthernetIsis616 struct {
	Authentication           InterfaceEthernetIsisAuthentication617             `json:"authentication"`
	BfdCfg                   InterfaceEthernetIsisBfdCfg621                     `json:"bfd-cfg"`
	CircuitType              string                                             `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceEthernetIsisCsnpIntervalList622         `json:"csnp-interval-list"`
	Padding                  int                                                `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceEthernetIsisHelloIntervalList623        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceEthernetIsisHelloIntervalMinimalList624 `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceEthernetIsisHelloMultiplierList625      `json:"hello-multiplier-list"`
	LspInterval              int                                                `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceEthernetIsisMeshGroup626                  `json:"mesh-group"`
	MetricList               []InterfaceEthernetIsisMetricList627               `json:"metric-list"`
	Network                  string                                             `json:"network"`
	PasswordList             []InterfaceEthernetIsisPasswordList628             `json:"password-list"`
	PriorityList             []InterfaceEthernetIsisPriorityList629             `json:"priority-list"`
	RetransmitInterval       int                                                `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceEthernetIsisWideMetricList630           `json:"wide-metric-list"`
	Uuid                     string                                             `json:"uuid"`
}

type InterfaceEthernetIsisAuthentication617 struct {
	SendOnlyList []InterfaceEthernetIsisAuthenticationSendOnlyList618 `json:"send-only-list"`
	ModeList     []InterfaceEthernetIsisAuthenticationModeList619     `json:"mode-list"`
	KeyChainList []InterfaceEthernetIsisAuthenticationKeyChainList620 `json:"key-chain-list"`
}

type InterfaceEthernetIsisAuthenticationSendOnlyList618 struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceEthernetIsisAuthenticationModeList619 struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceEthernetIsisAuthenticationKeyChainList620 struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceEthernetIsisBfdCfg621 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceEthernetIsisCsnpIntervalList622 struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceEthernetIsisHelloIntervalList623 struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceEthernetIsisHelloIntervalMinimalList624 struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceEthernetIsisHelloMultiplierList625 struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceEthernetIsisMeshGroup626 struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceEthernetIsisMetricList627 struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceEthernetIsisPasswordList628 struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceEthernetIsisPriorityList629 struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceEthernetIsisWideMetricList630 struct {
	WideMetric int    `json:"wide-metric" dval:"10"`
	Level      string `json:"level"`
}

type InterfaceEthernetLldp631 struct {
	EnableCfg       InterfaceEthernetLldpEnableCfg632       `json:"enable-cfg"`
	NotificationCfg InterfaceEthernetLldpNotificationCfg633 `json:"notification-cfg"`
	TxDot1Cfg       InterfaceEthernetLldpTxDot1Cfg634       `json:"tx-dot1-cfg"`
	TxTlvsCfg       InterfaceEthernetLldpTxTlvsCfg635       `json:"tx-tlvs-cfg"`
	Uuid            string                                  `json:"uuid"`
}

type InterfaceEthernetLldpEnableCfg632 struct {
	RtEnable int `json:"rt-enable"`
	Rx       int `json:"rx"`
	Tx       int `json:"tx"`
}

type InterfaceEthernetLldpNotificationCfg633 struct {
	Notification int `json:"notification"`
	NotifEnable  int `json:"notif-enable"`
}

type InterfaceEthernetLldpTxDot1Cfg634 struct {
	TxDot1Tlvs      int `json:"tx-dot1-tlvs"`
	LinkAggregation int `json:"link-aggregation"`
	Vlan            int `json:"vlan"`
}

type InterfaceEthernetLldpTxTlvsCfg635 struct {
	TxTlvs             int `json:"tx-tlvs"`
	Exclude            int `json:"exclude"`
	ManagementAddress  int `json:"management-address"`
	PortDescription    int `json:"port-description"`
	SystemCapabilities int `json:"system-capabilities"`
	SystemDescription  int `json:"system-description"`
	SystemName         int `json:"system-name"`
}

type InterfaceEthernetLw4o6636 struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceEthernetMap637 struct {
	Inside      int    `json:"inside"`
	Outside     int    `json:"outside"`
	MapTInside  int    `json:"map-t-inside"`
	MapTOutside int    `json:"map-t-outside"`
	Uuid        string `json:"uuid"`
}

type InterfaceEthernetMonitorList struct {
	Monitor     string `json:"monitor"`
	MirrorIndex int    `json:"mirror-index"`
	MonitorVlan int    `json:"monitor-vlan"`
}

type InterfaceEthernetNptv6638 struct {
	DomainList []InterfaceEthernetNptv6DomainList `json:"domain-list"`
}

type InterfaceEthernetNptv6DomainList struct {
	DomainName string `json:"domain-name"`
	BindType   string `json:"bind-type"`
	Uuid       string `json:"uuid"`
}

type InterfaceEthernetSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type InterfaceEthernetSpanningTree639 struct {
	AutoEdge     int                                            `json:"auto-edge" dval:"1"`
	AdminEdge    int                                            `json:"admin-edge"`
	InstanceList []InterfaceEthernetSpanningTreeInstanceList640 `json:"instance-list"`
	PathCost     int                                            `json:"path-cost"`
	Uuid         string                                         `json:"uuid"`
}

type InterfaceEthernetSpanningTreeInstanceList640 struct {
	InstanceStart int `json:"instance-start"`
	MstpPathCost  int `json:"mstp-path-cost"`
}

type InterfaceEthernetTrunkGroupList struct {
	TrunkNumber    int                                           `json:"trunk-number"`
	Type           string                                        `json:"type" dval:"static"`
	AdminKey       int                                           `json:"admin-key"`
	PortPriority   int                                           `json:"port-priority"`
	UdldTimeoutCfg InterfaceEthernetTrunkGroupListUdldTimeoutCfg `json:"udld-timeout-cfg"`
	Mode           string                                        `json:"mode" dval:"active"`
	Timeout        string                                        `json:"timeout" dval:"long"`
	Uuid           string                                        `json:"uuid"`
	UserTag        string                                        `json:"user-tag"`
}

type InterfaceEthernetTrunkGroupListUdldTimeoutCfg struct {
	Fast int `json:"fast" dval:"1000"`
	Slow int `json:"slow"`
}

func (p *InterfaceEthernet) GetId() string {
	return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceEthernet) getPath() string {
	return "interface/ethernet"
}

func (p *InterfaceEthernet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceEthernet::Post")
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

func (p *InterfaceEthernet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceEthernet::Get")
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
func (p *InterfaceEthernet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceEthernet::Put")
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

func (p *InterfaceEthernet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceEthernet::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
