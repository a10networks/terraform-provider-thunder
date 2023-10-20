package thunder

// based on ACOS 5_2_1-P3_70
type EndpointInterfaceEthernet struct {
	Inst InterfaceEthernet `json:"ethernet"`
}
type InterfaceEthernet struct {
	AccessList              InterfaceEthernetAccessList       `json:"access-list,omitempty"`
	Action                  string                            `json:"action,omitempty"`
	AutoNegEnable           int                               `json:"auto-neg-enable,omitempty"`
	Bfd                     InterfaceEthernetBfd              `json:"bfd,omitempty"`
	CpuProcess              int                               `json:"cpu-process,omitempty"`
	CpuProcessDir           string                            `json:"cpu-process-dir,omitempty"`
	Ddos                    InterfaceEthernetDdos             `json:"ddos,omitempty"`
	Duplexity               string                            `json:"duplexity,omitempty"`
	FecForcedOff            int                               `json:"fec-forced-off,omitempty"`
	FecForcedOn             int                               `json:"fec-forced-on,omitempty"`
	FlowControl             int                               `json:"flow-control,omitempty"`
	IcmpRateLimit           InterfaceEthernetIcmpRateLimit    `json:"icmp-rate-limit,omitempty"`
	Icmpv6RateLimit         InterfaceEthernetIcmpv6RateLimit  `json:"icmpv6-rate-limit,omitempty"`
	Ifnum                   int                               `json:"ifnum,omitempty"`
	Ip                      InterfaceEthernetIp               `json:"ip,omitempty"`
	IpgBitTime              int                               `json:"ipg-bit-time,omitempty"`
	Ipv6                    InterfaceEthernetIpv6             `json:"ipv6,omitempty"`
	Isis                    InterfaceEthernetIsis             `json:"isis,omitempty"`
	L3VlanFwdDisable        int                               `json:"l3-vlan-fwd-disable,omitempty"`
	Lldp                    InterfaceEthernetLldp             `json:"lldp,omitempty"`
	LoadInterval            int                               `json:"load-interval,omitempty"`
	Lw4o6                   InterfaceEthernetLw4o6            `json:"lw-4o6,omitempty"`
	MacLearning             string                            `json:"mac-learning,omitempty"`
	Map                     InterfaceEthernetMap              `json:"map,omitempty"`
	MediaTypeCopper         int                               `json:"media-type-copper,omitempty"`
	MonitorList             []InterfaceEthernetMonitorList    `json:"monitor-list,omitempty"`
	Mtu                     int                               `json:"mtu,omitempty"`
	Name                    string                            `json:"name,omitempty"`
	Nptv6                   InterfaceEthernetNptv6            `json:"nptv6,omitempty"`
	PacketCaptureTemplate   string                            `json:"packet-capture-template,omitempty"`
	PingSweepDetection      string                            `json:"ping-sweep-detection,omitempty"`
	PortBreakout            string                            `json:"port-breakout,omitempty"`
	PortScanDetection       string                            `json:"port-scan-detection,omitempty"`
	RemoveVlanTag           int                               `json:"remove-vlan-tag,omitempty"`
	SamplingEnable          []InterfaceEthernetSamplingEnable `json:"sampling-enable,omitempty"`
	SpanningTree            InterfaceEthernetSpanningTree     `json:"spanning-tree,omitempty"`
	Speed                   string                            `json:"speed,omitempty"`
	SpeedForced10g          int                               `json:"speed-forced-10g,omitempty"`
	SpeedForced1g           int                               `json:"speed-forced-1g,omitempty"`
	SpeedForced40g          int                               `json:"speed-forced-40g,omitempty"`
	TrafficDistributionMode string                            `json:"traffic-distribution-mode,omitempty"`
	TrapSource              int                               `json:"trap-source,omitempty"`
	TrunkGroupList          []InterfaceEthernetTrunkGroupList `json:"trunk-group-list,omitempty"`
	UpdateL2Info            int                               `json:"update-l2-info,omitempty"`
	UserTag                 string                            `json:"user-tag,omitempty"`
	Uuid                    string                            `json:"uuid,omitempty"`
	VirtualWire             int                               `json:"virtual-wire,omitempty"`
	VlanLearning            string                            `json:"vlan-learning,omitempty"`
}

type InterfaceEthernetAccessList struct {
	AclId   int    `json:"acl-id,omitempty"`
	AclName string `json:"acl-name,omitempty"`
}

type InterfaceEthernetBfd struct {
	Authentication InterfaceEthernetBfdAuthentication `json:"authentication,omitempty"`
	Demand         int                                `json:"demand,omitempty"`
	Echo           int                                `json:"echo,omitempty"`
	IntervalCfg    InterfaceEthernetBfdIntervalCfg    `json:"interval-cfg,omitempty"`
	Uuid           string                             `json:"uuid,omitempty"`
}

type InterfaceEthernetDdos struct {
	Inside  int    `json:"inside,omitempty"`
	Outside int    `json:"outside,omitempty"`
	Uuid    string `json:"uuid,omitempty"`
}

type InterfaceEthernetIcmpRateLimit struct {
	Lockup       int `json:"lockup,omitempty"`
	LockupPeriod int `json:"lockup-period,omitempty"`
	Normal       int `json:"normal,omitempty"`
}

type InterfaceEthernetIcmpv6RateLimit struct {
	LockupPeriodV6 int `json:"lockup-period-v6,omitempty"`
	LockupV6       int `json:"lockup-v6,omitempty"`
	NormalV6       int `json:"normal-v6,omitempty"`
}

type InterfaceEthernetIp struct {
	AddressList             []InterfaceEthernetIpAddressList       `json:"address-list,omitempty"`
	AllowPromiscuousVip     int                                    `json:"allow-promiscuous-vip,omitempty"`
	CacheSpoofingPort       int                                    `json:"cache-spoofing-port,omitempty"`
	Client                  int                                    `json:"client,omitempty"`
	Dhcp                    int                                    `json:"dhcp,omitempty"`
	GenerateMembershipQuery int                                    `json:"generate-membership-query,omitempty"`
	HelperAddressList       []InterfaceEthernetIpHelperAddressList `json:"helper-address-list,omitempty"`
	Inside                  int                                    `json:"inside,omitempty"`
	MaxRespTime             int                                    `json:"max-resp-time,omitempty"`
	Ospf                    InterfaceEthernetIpOspf                `json:"ospf,omitempty"`
	Outside                 int                                    `json:"outside,omitempty"`
	QueryInterval           int                                    `json:"query-interval,omitempty"`
	Rip                     InterfaceEthernetIpRip                 `json:"rip,omitempty"`
	Router                  InterfaceEthernetIpRouter              `json:"router,omitempty"`
	Server                  int                                    `json:"server,omitempty"`
	SlbPartitionRedirect    int                                    `json:"slb-partition-redirect,omitempty"`
	StatefulFirewall        InterfaceEthernetIpStatefulFirewall    `json:"stateful-firewall,omitempty"`
	SynCookie               int                                    `json:"syn-cookie,omitempty"`
	TtlIgnore               int                                    `json:"ttl-ignore,omitempty"`
	Uuid                    string                                 `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6 struct {
	AccessListCfg    InterfaceEthernetIpv6AccessListCfg    `json:"access-list-cfg,omitempty"`
	AddressList      []InterfaceEthernetIpv6AddressList    `json:"address-list,omitempty"`
	Inside           int                                   `json:"inside,omitempty"`
	Ipv6Enable       int                                   `json:"ipv6-enable,omitempty"`
	Ospf             InterfaceEthernetIpv6Ospf             `json:"ospf,omitempty"`
	Outside          int                                   `json:"outside,omitempty"`
	Rip              InterfaceEthernetIpv6Rip              `json:"rip,omitempty"`
	Router           InterfaceEthernetIpv6Router           `json:"router,omitempty"`
	RouterAdver      InterfaceEthernetIpv6RouterAdver      `json:"router-adver,omitempty"`
	StatefulFirewall InterfaceEthernetIpv6StatefulFirewall `json:"stateful-firewall,omitempty"`
	TtlIgnore        int                                   `json:"ttl-ignore,omitempty"`
	Uuid             string                                `json:"uuid,omitempty"`
}

type InterfaceEthernetIsis struct {
	Authentication           InterfaceEthernetIsisAuthentication             `json:"authentication,omitempty"`
	BfdCfg                   InterfaceEthernetIsisBfdCfg                     `json:"bfd-cfg,omitempty"`
	CircuitType              string                                          `json:"circuit-type,omitempty"`
	CsnpIntervalList         []InterfaceEthernetIsisCsnpIntervalList         `json:"csnp-interval-list,omitempty"`
	HelloIntervalList        []InterfaceEthernetIsisHelloIntervalList        `json:"hello-interval-list,omitempty"`
	HelloIntervalMinimalList []InterfaceEthernetIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list,omitempty"`
	HelloMultiplierList      []InterfaceEthernetIsisHelloMultiplierList      `json:"hello-multiplier-list,omitempty"`
	LspInterval              int                                             `json:"lsp-interval,omitempty"`
	MeshGroup                InterfaceEthernetIsisMeshGroup                  `json:"mesh-group,omitempty"`
	MetricList               []InterfaceEthernetIsisMetricList               `json:"metric-list,omitempty"`
	Network                  string                                          `json:"network,omitempty"`
	Padding                  int                                             `json:"padding,omitempty"`
	PasswordList             []InterfaceEthernetIsisPasswordList             `json:"password-list,omitempty"`
	PriorityList             []InterfaceEthernetIsisPriorityList             `json:"priority-list,omitempty"`
	RetransmitInterval       int                                             `json:"retransmit-interval,omitempty"`
	Uuid                     string                                          `json:"uuid,omitempty"`
	WideMetricList           []InterfaceEthernetIsisWideMetricList           `json:"wide-metric-list,omitempty"`
}

type InterfaceEthernetLldp struct {
	EnableCfg       InterfaceEthernetLldpEnableCfg       `json:"enable-cfg,omitempty"`
	NotificationCfg InterfaceEthernetLldpNotificationCfg `json:"notification-cfg,omitempty"`
	TxDot1Cfg       InterfaceEthernetLldpTxDot1Cfg       `json:"tx-dot1-cfg,omitempty"`
	TxTlvsCfg       InterfaceEthernetLldpTxTlvsCfg       `json:"tx-tlvs-cfg,omitempty"`
	Uuid            string                               `json:"uuid,omitempty"`
}

type InterfaceEthernetLw4o6 struct {
	Inside  int    `json:"inside,omitempty"`
	Outside int    `json:"outside,omitempty"`
	Uuid    string `json:"uuid,omitempty"`
}

type InterfaceEthernetMap struct {
	Inside      int    `json:"inside,omitempty"`
	MapTInside  int    `json:"map-t-inside,omitempty"`
	MapTOutside int    `json:"map-t-outside,omitempty"`
	Outside     int    `json:"outside,omitempty"`
	Uuid        string `json:"uuid,omitempty"`
}

type InterfaceEthernetMonitorList struct {
	MirrorIndex int    `json:"mirror-index,omitempty"`
	Monitor     string `json:"monitor,omitempty"`
	MonitorVlan int    `json:"monitor-vlan,omitempty"`
}

type InterfaceEthernetNptv6 struct {
	DomainList []InterfaceEthernetNptv6DomainList `json:"domain-list,omitempty"`
}

type InterfaceEthernetSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

type InterfaceEthernetSpanningTree struct {
	AdminEdge    int                                         `json:"admin-edge,omitempty"`
	AutoEdge     int                                         `json:"auto-edge,omitempty"`
	InstanceList []InterfaceEthernetSpanningTreeInstanceList `json:"instance-list,omitempty"`
	PathCost     int                                         `json:"path-cost,omitempty"`
	Uuid         string                                      `json:"uuid,omitempty"`
}

type InterfaceEthernetTrunkGroupList struct {
	AdminKey       int                                           `json:"admin-key,omitempty"`
	Mode           string                                        `json:"mode,omitempty"`
	PortPriority   int                                           `json:"port-priority,omitempty"`
	Timeout        string                                        `json:"timeout,omitempty"`
	TrunkNumber    int                                           `json:"trunk-number,omitempty"`
	Type           string                                        `json:"type,omitempty"`
	UdldTimeoutCfg InterfaceEthernetTrunkGroupListUdldTimeoutCfg `json:"udld-timeout-cfg,omitempty"`
	UserTag        string                                        `json:"user-tag,omitempty"`
	Uuid           string                                        `json:"uuid,omitempty"`
}

type InterfaceEthernetBfdAuthentication struct {
	Encrypted string `json:"encrypted,omitempty"`
	KeyId     int    `json:"key-id,omitempty"`
	Method    string `json:"method,omitempty"`
	Password  string `json:"password,omitempty"`
}

type InterfaceEthernetBfdIntervalCfg struct {
	Interval   int `json:"interval,omitempty"`
	MinRx      int `json:"min-rx,omitempty"`
	Multiplier int `json:"multiplier,omitempty"`
}

type InterfaceEthernetIpAddressList struct {
	Ipv4Address string `json:"ipv4-address,omitempty"`
	Ipv4Netmask string `json:"ipv4-netmask,omitempty"`
}

type InterfaceEthernetIpHelperAddressList struct {
	HelperAddress string `json:"helper-address,omitempty"`
}

type InterfaceEthernetIpOspf struct {
	OspfGlobal InterfaceEthernetIpOspfOspfGlobal   `json:"ospf-global,omitempty"`
	OspfIpList []InterfaceEthernetIpOspfOspfIpList `json:"ospf-ip-list,omitempty"`
}

type InterfaceEthernetIpRip struct {
	Authentication  InterfaceEthernetIpRipAuthentication  `json:"authentication,omitempty"`
	ReceiveCfg      InterfaceEthernetIpRipReceiveCfg      `json:"receive-cfg,omitempty"`
	ReceivePacket   int                                   `json:"receive-packet,omitempty"`
	SendCfg         InterfaceEthernetIpRipSendCfg         `json:"send-cfg,omitempty"`
	SendPacket      int                                   `json:"send-packet,omitempty"`
	SplitHorizonCfg InterfaceEthernetIpRipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	Uuid            string                                `json:"uuid,omitempty"`
}

type InterfaceEthernetIpRouter struct {
	Isis InterfaceEthernetIpRouterIsis `json:"isis,omitempty"`
}

type InterfaceEthernetIpStatefulFirewall struct {
	AccessList int    `json:"access-list,omitempty"`
	AclId      int    `json:"acl-id,omitempty"`
	ClassList  string `json:"class-list,omitempty"`
	Inside     int    `json:"inside,omitempty"`
	Outside    int    `json:"outside,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6AccessListCfg struct {
	Inbound   int    `json:"inbound,omitempty"`
	V6AclName string `json:"v6-acl-name,omitempty"`
}

type InterfaceEthernetIpv6AddressList struct {
	AddressType string `json:"address-type,omitempty"`
	Ipv6Addr    string `json:"ipv6-addr,omitempty"`
}

type InterfaceEthernetIpv6Ospf struct {
	Bfd                   int                                              `json:"bfd,omitempty"`
	CostCfg               []InterfaceEthernetIpv6OspfCostCfg               `json:"cost-cfg,omitempty"`
	DeadIntervalCfg       []InterfaceEthernetIpv6OspfDeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
	Disable               int                                              `json:"disable,omitempty"`
	HelloIntervalCfg      []InterfaceEthernetIpv6OspfHelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	MtuIgnoreCfg          []InterfaceEthernetIpv6OspfMtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	NeighborCfg           []InterfaceEthernetIpv6OspfNeighborCfg           `json:"neighbor-cfg,omitempty"`
	NetworkList           []InterfaceEthernetIpv6OspfNetworkList           `json:"network-list,omitempty"`
	PriorityCfg           []InterfaceEthernetIpv6OspfPriorityCfg           `json:"priority-cfg,omitempty"`
	RetransmitIntervalCfg []InterfaceEthernetIpv6OspfRetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	TransmitDelayCfg      []InterfaceEthernetIpv6OspfTransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	Uuid                  string                                           `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6Rip struct {
	SplitHorizonCfg InterfaceEthernetIpv6RipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	Uuid            string                                  `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6Router struct {
	Isis  InterfaceEthernetIpv6RouterIsis  `json:"isis,omitempty"`
	Ospf  InterfaceEthernetIpv6RouterOspf  `json:"ospf,omitempty"`
	Ripng InterfaceEthernetIpv6RouterRipng `json:"ripng,omitempty"`
}

type InterfaceEthernetIpv6RouterAdver struct {
	Action                   string                                       `json:"action,omitempty"`
	AdverMtu                 int                                          `json:"adver-mtu,omitempty"`
	AdverMtuDisable          int                                          `json:"adver-mtu-disable,omitempty"`
	AdverVrid                int                                          `json:"adver-vrid,omitempty"`
	AdverVridDefault         int                                          `json:"adver-vrid-default,omitempty"`
	DefaultLifetime          int                                          `json:"default-lifetime,omitempty"`
	FloatingIp               string                                       `json:"floating-ip,omitempty"`
	FloatingIpDefaultVrid    string                                       `json:"floating-ip-default-vrid,omitempty"`
	HopLimit                 int                                          `json:"hop-limit,omitempty"`
	ManagedConfigAction      string                                       `json:"managed-config-action,omitempty"`
	MaxInterval              int                                          `json:"max-interval,omitempty"`
	MinInterval              int                                          `json:"min-interval,omitempty"`
	OtherConfigAction        string                                       `json:"other-config-action,omitempty"`
	PrefixList               []InterfaceEthernetIpv6RouterAdverPrefixList `json:"prefix-list,omitempty"`
	RateLimit                int                                          `json:"rate-limit,omitempty"`
	ReachableTime            int                                          `json:"reachable-time,omitempty"`
	RetransmitTimer          int                                          `json:"retransmit-timer,omitempty"`
	UseFloatingIp            int                                          `json:"use-floating-ip,omitempty"`
	UseFloatingIpDefaultVrid int                                          `json:"use-floating-ip-default-vrid,omitempty"`
}

type InterfaceEthernetIpv6StatefulFirewall struct {
	AccessList int    `json:"access-list,omitempty"`
	AclName    string `json:"acl-name,omitempty"`
	ClassList  string `json:"class-list,omitempty"`
	Inside     int    `json:"inside,omitempty"`
	Outside    int    `json:"outside,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
}

type InterfaceEthernetIsisAuthentication struct {
	KeyChainList []InterfaceEthernetIsisAuthenticationKeyChainList `json:"key-chain-list,omitempty"`
	ModeList     []InterfaceEthernetIsisAuthenticationModeList     `json:"mode-list,omitempty"`
	SendOnlyList []InterfaceEthernetIsisAuthenticationSendOnlyList `json:"send-only-list,omitempty"`
}

type InterfaceEthernetIsisBfdCfg struct {
	Bfd     int `json:"bfd,omitempty"`
	Disable int `json:"disable,omitempty"`
}

type InterfaceEthernetIsisCsnpIntervalList struct {
	CsnpInterval int    `json:"csnp-interval,omitempty"`
	Level        string `json:"level,omitempty"`
}

type InterfaceEthernetIsisHelloIntervalList struct {
	HelloInterval int    `json:"hello-interval,omitempty"`
	Level         string `json:"level,omitempty"`
}

type InterfaceEthernetIsisHelloIntervalMinimalList struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal,omitempty"`
	Level                string `json:"level,omitempty"`
}

type InterfaceEthernetIsisHelloMultiplierList struct {
	HelloMultiplier int    `json:"hello-multiplier,omitempty"`
	Level           string `json:"level,omitempty"`
}

type InterfaceEthernetIsisMeshGroup struct {
	Blocked int `json:"blocked,omitempty"`
	Value   int `json:"value,omitempty"`
}

type InterfaceEthernetIsisMetricList struct {
	Level  string `json:"level,omitempty"`
	Metric int    `json:"metric,omitempty"`
}

type InterfaceEthernetIsisPasswordList struct {
	Level    string `json:"level,omitempty"`
	Password string `json:"password,omitempty"`
}

type InterfaceEthernetIsisPriorityList struct {
	Level    string `json:"level,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type InterfaceEthernetIsisWideMetricList struct {
	Level      string `json:"level,omitempty"`
	WideMetric int    `json:"wide-metric,omitempty"`
}

type InterfaceEthernetLldpEnableCfg struct {
	RtEnable int `json:"rt-enable,omitempty"`
	Rx       int `json:"rx,omitempty"`
	Tx       int `json:"tx,omitempty"`
}

type InterfaceEthernetLldpNotificationCfg struct {
	NotifEnable  int `json:"notif-enable,omitempty"`
	Notification int `json:"notification,omitempty"`
}

type InterfaceEthernetLldpTxDot1Cfg struct {
	LinkAggregation int `json:"link-aggregation,omitempty"`
	TxDot1Tlvs      int `json:"tx-dot1-tlvs,omitempty"`
	Vlan            int `json:"vlan,omitempty"`
}

type InterfaceEthernetLldpTxTlvsCfg struct {
	Exclude            int `json:"exclude,omitempty"`
	ManagementAddress  int `json:"management-address,omitempty"`
	PortDescription    int `json:"port-description,omitempty"`
	SystemCapabilities int `json:"system-capabilities,omitempty"`
	SystemDescription  int `json:"system-description,omitempty"`
	SystemName         int `json:"system-name,omitempty"`
	TxTlvs             int `json:"tx-tlvs,omitempty"`
}

type InterfaceEthernetNptv6DomainList struct {
	BindType   string `json:"bind-type,omitempty"`
	DomainName string `json:"domain-name,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
}

type InterfaceEthernetSpanningTreeInstanceList struct {
	InstanceStart int `json:"instance-start,omitempty"`
	MstpPathCost  int `json:"mstp-path-cost,omitempty"`
}

type InterfaceEthernetTrunkGroupListUdldTimeoutCfg struct {
	Fast int `json:"fast,omitempty"`
	Slow int `json:"slow,omitempty"`
}

type InterfaceEthernetIpOspfOspfGlobal struct {
	AuthenticationCfg  InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	AuthenticationKey  string                                              `json:"authentication-key,omitempty"`
	BfdCfg             InterfaceEthernetIpOspfOspfGlobalBfdCfg             `json:"bfd-cfg,omitempty"`
	Cost               int                                                 `json:"cost,omitempty"`
	DatabaseFilterCfg  InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	DeadInterval       int                                                 `json:"dead-interval,omitempty"`
	Disable            string                                              `json:"disable,omitempty"`
	HelloInterval      int                                                 `json:"hello-interval,omitempty"`
	MessageDigestCfg   []InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	Mtu                int                                                 `json:"mtu,omitempty"`
	MtuIgnore          int                                                 `json:"mtu-ignore,omitempty"`
	Network            InterfaceEthernetIpOspfOspfGlobalNetwork            `json:"network,omitempty"`
	Priority           int                                                 `json:"priority,omitempty"`
	RetransmitInterval int                                                 `json:"retransmit-interval,omitempty"`
	TransmitDelay      int                                                 `json:"transmit-delay,omitempty"`
	Uuid               string                                              `json:"uuid,omitempty"`
}

type InterfaceEthernetIpOspfOspfIpList struct {
	Authentication     int                                                 `json:"authentication,omitempty"`
	AuthenticationKey  string                                              `json:"authentication-key,omitempty"`
	Cost               int                                                 `json:"cost,omitempty"`
	DatabaseFilter     string                                              `json:"database-filter,omitempty"`
	DeadInterval       int                                                 `json:"dead-interval,omitempty"`
	HelloInterval      int                                                 `json:"hello-interval,omitempty"`
	IpAddr             string                                              `json:"ip-addr,omitempty"`
	MessageDigestCfg   []InterfaceEthernetIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	MtuIgnore          int                                                 `json:"mtu-ignore,omitempty"`
	Out                int                                                 `json:"out,omitempty"`
	Priority           int                                                 `json:"priority,omitempty"`
	RetransmitInterval int                                                 `json:"retransmit-interval,omitempty"`
	TransmitDelay      int                                                 `json:"transmit-delay,omitempty"`
	Uuid               string                                              `json:"uuid,omitempty"`
	Value              string                                              `json:"value,omitempty"`
}

type InterfaceEthernetIpRipAuthentication struct {
	KeyChain InterfaceEthernetIpRipAuthenticationKeyChain `json:"key-chain,omitempty"`
	Mode     InterfaceEthernetIpRipAuthenticationMode     `json:"mode,omitempty"`
	Str      InterfaceEthernetIpRipAuthenticationStr      `json:"str,omitempty"`
}

type InterfaceEthernetIpRipReceiveCfg struct {
	Receive int    `json:"receive,omitempty"`
	Version string `json:"version,omitempty"`
}

type InterfaceEthernetIpRipSendCfg struct {
	Send    int    `json:"send,omitempty"`
	Version string `json:"version,omitempty"`
}

type InterfaceEthernetIpRipSplitHorizonCfg struct {
	State string `json:"state,omitempty"`
}

type InterfaceEthernetIpRouterIsis struct {
	Tag  string `json:"tag,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6OspfCostCfg struct {
	Cost       int `json:"cost,omitempty"`
	InstanceId int `json:"instance-id,omitempty"`
}

type InterfaceEthernetIpv6OspfDeadIntervalCfg struct {
	DeadInterval int `json:"dead-interval,omitempty"`
	InstanceId   int `json:"instance-id,omitempty"`
}

type InterfaceEthernetIpv6OspfHelloIntervalCfg struct {
	HelloInterval int `json:"hello-interval,omitempty"`
	InstanceId    int `json:"instance-id,omitempty"`
}

type InterfaceEthernetIpv6OspfMtuIgnoreCfg struct {
	InstanceId int `json:"instance-id,omitempty"`
	MtuIgnore  int `json:"mtu-ignore,omitempty"`
}

type InterfaceEthernetIpv6OspfNeighborCfg struct {
	NeigInst             int    `json:"neig-inst,omitempty"`
	Neighbor             string `json:"neighbor,omitempty"`
	NeighborCost         int    `json:"neighbor-cost,omitempty"`
	NeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	NeighborPriority     int    `json:"neighbor-priority,omitempty"`
}

type InterfaceEthernetIpv6OspfNetworkList struct {
	BroadcastType     string `json:"broadcast-type,omitempty"`
	NetworkInstanceId int    `json:"network-instance-id,omitempty"`
	P2mpNbma          int    `json:"p2mp-nbma,omitempty"`
}

type InterfaceEthernetIpv6OspfPriorityCfg struct {
	InstanceId int `json:"instance-id,omitempty"`
	Priority   int `json:"priority,omitempty"`
}

type InterfaceEthernetIpv6OspfRetransmitIntervalCfg struct {
	InstanceId         int `json:"instance-id,omitempty"`
	RetransmitInterval int `json:"retransmit-interval,omitempty"`
}

type InterfaceEthernetIpv6OspfTransmitDelayCfg struct {
	InstanceId    int `json:"instance-id,omitempty"`
	TransmitDelay int `json:"transmit-delay,omitempty"`
}

type InterfaceEthernetIpv6RipSplitHorizonCfg struct {
	State string `json:"state,omitempty"`
}

type InterfaceEthernetIpv6RouterIsis struct {
	Tag  string `json:"tag,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6RouterOspf struct {
	AreaList []InterfaceEthernetIpv6RouterOspfAreaList `json:"area-list,omitempty"`
	Uuid     string                                    `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6RouterRipng struct {
	Rip  int    `json:"rip,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

type InterfaceEthernetIpv6RouterAdverPrefixList struct {
	NotAutonomous     int    `json:"not-autonomous,omitempty"`
	NotOnLink         int    `json:"not-on-link,omitempty"`
	PreferredLifetime int    `json:"preferred-lifetime,omitempty"`
	Prefix            string `json:"prefix,omitempty"`
	ValidLifetime     int    `json:"valid-lifetime,omitempty"`
}

type InterfaceEthernetIsisAuthenticationKeyChainList struct {
	KeyChain string `json:"key-chain,omitempty"`
	Level    string `json:"level,omitempty"`
}

type InterfaceEthernetIsisAuthenticationModeList struct {
	Level string `json:"level,omitempty"`
	Mode  string `json:"mode,omitempty"`
}

type InterfaceEthernetIsisAuthenticationSendOnlyList struct {
	Level    string `json:"level,omitempty"`
	SendOnly int    `json:"send-only,omitempty"`
}

type InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg struct {
	Authentication int    `json:"authentication,omitempty"`
	Value          string `json:"value,omitempty"`
}

type InterfaceEthernetIpOspfOspfGlobalBfdCfg struct {
	Bfd     int `json:"bfd,omitempty"`
	Disable int `json:"disable,omitempty"`
}

type InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg struct {
	DatabaseFilter string `json:"database-filter,omitempty"`
	Out            int    `json:"out,omitempty"`
}

type InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg struct {
	Md5              InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5 `json:"md5,omitempty"`
	MessageDigestKey int                                                  `json:"message-digest-key,omitempty"`
}

type InterfaceEthernetIpOspfOspfGlobalNetwork struct {
	Broadcast         int `json:"broadcast,omitempty"`
	NonBroadcast      int `json:"non-broadcast,omitempty"`
	P2mpNbma          int `json:"p2mp-nbma,omitempty"`
	PointToMultipoint int `json:"point-to-multipoint,omitempty"`
	PointToPoint      int `json:"point-to-point,omitempty"`
}

type InterfaceEthernetIpOspfOspfIpListMessageDigestCfg struct {
	Encrypted        string `json:"encrypted,omitempty"`
	Md5Value         string `json:"md5-value,omitempty"`
	MessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceEthernetIpRipAuthenticationKeyChain struct {
	KeyChain string `json:"key-chain,omitempty"`
}

type InterfaceEthernetIpRipAuthenticationMode struct {
	Mode string `json:"mode,omitempty"`
}

type InterfaceEthernetIpRipAuthenticationStr struct {
	String string `json:"string,omitempty"`
}

type InterfaceEthernetIpv6RouterOspfAreaList struct {
	AreaIdAddr string `json:"area-id-addr,omitempty"`
	AreaIdNum  int    `json:"area-id-num,omitempty"`
	InstanceId int    `json:"instance-id,omitempty"`
	Tag        string `json:"tag,omitempty"`
}

type InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5 struct {
	Encrypted string `json:"encrypted,omitempty"`
	Md5Value  string `json:"md5-value,omitempty"`
}
