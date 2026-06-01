package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type InterfaceTrunk struct {
	Inst struct {
		AccessList InterfaceTrunkAccessList `json:"access-list"`

		Action string `json:"action" dval:"enable"`

		Bfd InterfaceTrunkBfd861 `json:"bfd"`

		Ddos InterfaceTrunkDdos865 `json:"ddos"`

		DoAutoRecovery int `json:"do-auto-recovery"`

		GamingProtocolCompliance int `json:"gaming-protocol-compliance"`

		IcmpRateLimit InterfaceTrunkIcmpRateLimit `json:"icmp-rate-limit"`

		Icmpv6RateLimit InterfaceTrunkIcmpv6RateLimit `json:"icmpv6-rate-limit"`

		Ifnum int `json:"ifnum"`

		Ip InterfaceTrunkIp866 `json:"ip"`

		Ipv6 InterfaceTrunkIpv6891 `json:"ipv6"`

		Isis InterfaceTrunkIsis917 `json:"isis"`

		L3VlanFwdDisable int `json:"l3-vlan-fwd-disable"`

		Lw4o6 InterfaceTrunkLw4o6932 `json:"lw-4o6"`

		MacLearning string `json:"mac-learning"`

		Map InterfaceTrunkMap933 `json:"map"`

		Mtu int `json:"mtu"`

		Name string `json:"name"`

		Nptv6 InterfaceTrunkNptv6934 `json:"nptv6"`

		PortsThreshold int `json:"ports-threshold"`

		SamplingEnable []InterfaceTrunkSamplingEnable `json:"sampling-enable"`

		SpanningTree InterfaceTrunkSpanningTree935 `json:"spanning-tree"`

		SyncModifyDisable int `json:"sync-modify-disable"`

		Timer int `json:"timer" dval:"10"`

		TrapSource int `json:"trap-source"`

		UpdateL2Info int `json:"update-l2-info"`

		UseHwHash int `json:"use-hw-hash"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		VirtualWire int `json:"virtual-wire"`

		VlanLearning string `json:"vlan-learning"`
	} `json:"trunk"`
}

type InterfaceTrunkAccessList struct {
	AclId   int    `json:"acl-id"`
	AclName string `json:"acl-name"`
}

type InterfaceTrunkBfd861 struct {
	Authentication InterfaceTrunkBfdAuthentication862 `json:"authentication"`
	Echo           int                                `json:"echo"`
	Demand         int                                `json:"demand"`
	IntervalCfg    InterfaceTrunkBfdIntervalCfg863    `json:"interval-cfg"`
	Uuid           string                             `json:"uuid"`
	PerMemberPort  InterfaceTrunkBfdPerMemberPort864  `json:"per-member-port"`
}

type InterfaceTrunkBfdAuthentication862 struct {
	KeyId     int    `json:"key-id"`
	Method    string `json:"method"`
	Password  string `json:"password"`
	Encrypted string `json:"encrypted"`
}

type InterfaceTrunkBfdIntervalCfg863 struct {
	Interval   int `json:"interval"`
	MinRx      int `json:"min-rx"`
	Multiplier int `json:"multiplier"`
}

type InterfaceTrunkBfdPerMemberPort864 struct {
	LocalAddress    string `json:"local-address"`
	NeighborAddress string `json:"neighbor-address"`
	Ipv6Local       string `json:"ipv6-local"`
	Ipv6Nbr         string `json:"ipv6-nbr"`
	Uuid            string `json:"uuid"`
}

type InterfaceTrunkDdos865 struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceTrunkIcmpRateLimit struct {
	Normal       int `json:"normal"`
	Lockup       int `json:"lockup"`
	LockupPeriod int `json:"lockup-period"`
}

type InterfaceTrunkIcmpv6RateLimit struct {
	NormalV6       int `json:"normal-v6"`
	LockupV6       int `json:"lockup-v6"`
	LockupPeriodV6 int `json:"lockup-period-v6"`
}

type InterfaceTrunkIp866 struct {
	Dhcp                    int                                    `json:"dhcp"`
	AddressList             []InterfaceTrunkIpAddressList867       `json:"address-list"`
	AllowPromiscuousVip     int                                    `json:"allow-promiscuous-vip"`
	Client                  int                                    `json:"client"`
	Server                  int                                    `json:"server"`
	Dmz                     int                                    `json:"dmz"`
	CacheSpoofingPort       int                                    `json:"cache-spoofing-port"`
	HelperAddressList       []InterfaceTrunkIpHelperAddressList868 `json:"helper-address-list"`
	Nat                     InterfaceTrunkIpNat869                 `json:"nat"`
	TtlIgnore               int                                    `json:"ttl-ignore"`
	SynCookie               int                                    `json:"syn-cookie"`
	SlbPartitionRedirect    int                                    `json:"slb-partition-redirect"`
	GenerateMembershipQuery int                                    `json:"generate-membership-query"`
	QueryInterval           int                                    `json:"query-interval" dval:"125"`
	MaxRespTime             int                                    `json:"max-resp-time" dval:"100"`
	Unnumbered              int                                    `json:"unnumbered"`
	Uuid                    string                                 `json:"uuid"`
	StatefulFirewall        InterfaceTrunkIpStatefulFirewall870    `json:"stateful-firewall"`
	Router                  InterfaceTrunkIpRouter871              `json:"router"`
	Rip                     InterfaceTrunkIpRip873                 `json:"rip"`
	Ospf                    InterfaceTrunkIpOspf881                `json:"ospf"`
}

type InterfaceTrunkIpAddressList867 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceTrunkIpHelperAddressList868 struct {
	HelperAddress string `json:"helper-address"`
}

type InterfaceTrunkIpNat869 struct {
	Inside  int `json:"inside"`
	Outside int `json:"outside"`
}

type InterfaceTrunkIpStatefulFirewall870 struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclId      int    `json:"acl-id"`
	Uuid       string `json:"uuid"`
}

type InterfaceTrunkIpRouter871 struct {
	Isis InterfaceTrunkIpRouterIsis872 `json:"isis"`
}

type InterfaceTrunkIpRouterIsis872 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceTrunkIpRip873 struct {
	Authentication  InterfaceTrunkIpRipAuthentication874  `json:"authentication"`
	SendPacket      int                                   `json:"send-packet" dval:"1"`
	ReceivePacket   int                                   `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceTrunkIpRipSendCfg878         `json:"send-cfg"`
	ReceiveCfg      InterfaceTrunkIpRipReceiveCfg879      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceTrunkIpRipSplitHorizonCfg880 `json:"split-horizon-cfg"`
	Uuid            string                                `json:"uuid"`
}

type InterfaceTrunkIpRipAuthentication874 struct {
	Str      InterfaceTrunkIpRipAuthenticationStr875      `json:"str"`
	Mode     InterfaceTrunkIpRipAuthenticationMode876     `json:"mode"`
	KeyChain InterfaceTrunkIpRipAuthenticationKeyChain877 `json:"key-chain"`
}

type InterfaceTrunkIpRipAuthenticationStr875 struct {
	String string `json:"string"`
}

type InterfaceTrunkIpRipAuthenticationMode876 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceTrunkIpRipAuthenticationKeyChain877 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceTrunkIpRipSendCfg878 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceTrunkIpRipReceiveCfg879 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceTrunkIpRipSplitHorizonCfg880 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceTrunkIpOspf881 struct {
	OspfGlobal InterfaceTrunkIpOspfOspfGlobal882   `json:"ospf-global"`
	OspfIpList []InterfaceTrunkIpOspfOspfIpList889 `json:"ospf-ip-list"`
}

type InterfaceTrunkIpOspfOspfGlobal882 struct {
	AuthenticationCfg  InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg883  `json:"authentication-cfg"`
	AuthenticationKey  string                                              `json:"authentication-key"`
	BfdCfg             InterfaceTrunkIpOspfOspfGlobalBfdCfg884             `json:"bfd-cfg"`
	Cost               int                                                 `json:"cost"`
	DatabaseFilterCfg  InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg885  `json:"database-filter-cfg"`
	DeadInterval       int                                                 `json:"dead-interval" dval:"40"`
	Disable            string                                              `json:"disable"`
	HelloInterval      int                                                 `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg886 `json:"message-digest-cfg"`
	Mtu                int                                                 `json:"mtu"`
	MtuIgnore          int                                                 `json:"mtu-ignore"`
	Network            InterfaceTrunkIpOspfOspfGlobalNetwork888            `json:"network"`
	Priority           int                                                 `json:"priority" dval:"1"`
	RetransmitInterval int                                                 `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                 `json:"transmit-delay" dval:"1"`
	Uuid               string                                              `json:"uuid"`
}

type InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg883 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceTrunkIpOspfOspfGlobalBfdCfg884 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg885 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg886 struct {
	MessageDigestKey int                                                  `json:"message-digest-key"`
	Md5              InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5887 `json:"md5"`
}

type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5887 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceTrunkIpOspfOspfGlobalNetwork888 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceTrunkIpOspfOspfIpList889 struct {
	IpAddr             string                                              `json:"ip-addr"`
	Authentication     int                                                 `json:"authentication"`
	Value              string                                              `json:"value"`
	AuthenticationKey  string                                              `json:"authentication-key"`
	Cost               int                                                 `json:"cost"`
	DatabaseFilter     string                                              `json:"database-filter"`
	Out                int                                                 `json:"out"`
	DeadInterval       int                                                 `json:"dead-interval" dval:"40"`
	HelloInterval      int                                                 `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTrunkIpOspfOspfIpListMessageDigestCfg890 `json:"message-digest-cfg"`
	MtuIgnore          int                                                 `json:"mtu-ignore"`
	Priority           int                                                 `json:"priority" dval:"1"`
	RetransmitInterval int                                                 `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                 `json:"transmit-delay" dval:"1"`
	Uuid               string                                              `json:"uuid"`
}

type InterfaceTrunkIpOspfOspfIpListMessageDigestCfg890 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceTrunkIpv6891 struct {
	AddressList      []InterfaceTrunkIpv6AddressList892    `json:"address-list"`
	Ipv6Enable       int                                   `json:"ipv6-enable"`
	AccessListCfg    InterfaceTrunkIpv6AccessListCfg893    `json:"access-list-cfg"`
	Nat              InterfaceTrunkIpv6Nat894              `json:"nat"`
	TtlIgnore        int                                   `json:"ttl-ignore"`
	RouterAdver      InterfaceTrunkIpv6RouterAdver895      `json:"router-adver"`
	Uuid             string                                `json:"uuid"`
	StatefulFirewall InterfaceTrunkIpv6StatefulFirewall899 `json:"stateful-firewall"`
	Router           InterfaceTrunkIpv6Router900           `json:"router"`
	Rip              InterfaceTrunkIpv6Rip905              `json:"rip"`
	Ospf             InterfaceTrunkIpv6Ospf907             `json:"ospf"`
}

type InterfaceTrunkIpv6AddressList892 struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceTrunkIpv6AccessListCfg893 struct {
	V6AclName string `json:"v6-acl-name"`
	Inbound   int    `json:"inbound"`
}

type InterfaceTrunkIpv6Nat894 struct {
	Inside  int `json:"inside"`
	Outside int `json:"outside"`
}

type InterfaceTrunkIpv6RouterAdver895 struct {
	Action              string                                       `json:"action" dval:"disable"`
	DefaultLifetime     int                                          `json:"default-lifetime" dval:"1800"`
	HopLimit            int                                          `json:"hop-limit" dval:"255"`
	MaxInterval         int                                          `json:"max-interval" dval:"600"`
	MinInterval         int                                          `json:"min-interval" dval:"200"`
	RateLimit           int                                          `json:"rate-limit" dval:"100000"`
	ReachableTime       int                                          `json:"reachable-time"`
	RetransmitTimer     int                                          `json:"retransmit-timer"`
	Mtu                 InterfaceTrunkIpv6RouterAdverMtu896          `json:"mtu"`
	PrefixList          []InterfaceTrunkIpv6RouterAdverPrefixList897 `json:"prefix-list"`
	ManagedConfigAction string                                       `json:"managed-config-action" dval:"disable"`
	OtherConfigAction   string                                       `json:"other-config-action" dval:"disable"`
	Vrid                InterfaceTrunkIpv6RouterAdverVrid898         `json:"vrid"`
}

type InterfaceTrunkIpv6RouterAdverMtu896 struct {
	AdverMtuDisable int `json:"adver-mtu-disable" dval:"1"`
	AdverMtu        int `json:"adver-mtu"`
}

type InterfaceTrunkIpv6RouterAdverPrefixList897 struct {
	Prefix            string `json:"prefix"`
	NotAutonomous     int    `json:"not-autonomous"`
	NotOnLink         int    `json:"not-on-link"`
	PreferredLifetime int    `json:"preferred-lifetime" dval:"604800"`
	ValidLifetime     int    `json:"valid-lifetime" dval:"2592000"`
}

type InterfaceTrunkIpv6RouterAdverVrid898 struct {
	AdverVrid                int    `json:"adver-vrid"`
	UseFloatingIp            int    `json:"use-floating-ip"`
	FloatingIp               string `json:"floating-ip"`
	AdverVridDefault         int    `json:"adver-vrid-default"`
	UseFloatingIpDefaultVrid int    `json:"use-floating-ip-default-vrid"`
	FloatingIpDefaultVrid    string `json:"floating-ip-default-vrid"`
}

type InterfaceTrunkIpv6StatefulFirewall899 struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclName    string `json:"acl-name"`
	Uuid       string `json:"uuid"`
}

type InterfaceTrunkIpv6Router900 struct {
	Ripng InterfaceTrunkIpv6RouterRipng901 `json:"ripng"`
	Ospf  InterfaceTrunkIpv6RouterOspf902  `json:"ospf"`
	Isis  InterfaceTrunkIpv6RouterIsis904  `json:"isis"`
}

type InterfaceTrunkIpv6RouterRipng901 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceTrunkIpv6RouterOspf902 struct {
	AreaList []InterfaceTrunkIpv6RouterOspfAreaList903 `json:"area-list"`
	Uuid     string                                    `json:"uuid"`
}

type InterfaceTrunkIpv6RouterOspfAreaList903 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceTrunkIpv6RouterIsis904 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceTrunkIpv6Rip905 struct {
	SplitHorizonCfg InterfaceTrunkIpv6RipSplitHorizonCfg906 `json:"split-horizon-cfg"`
	Uuid            string                                  `json:"uuid"`
}

type InterfaceTrunkIpv6RipSplitHorizonCfg906 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceTrunkIpv6Ospf907 struct {
	NetworkList           []InterfaceTrunkIpv6OspfNetworkList908           `json:"network-list"`
	Bfd                   int                                              `json:"bfd"`
	Disable               int                                              `json:"disable"`
	CostCfg               []InterfaceTrunkIpv6OspfCostCfg909               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceTrunkIpv6OspfDeadIntervalCfg910       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceTrunkIpv6OspfHelloIntervalCfg911      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceTrunkIpv6OspfMtuIgnoreCfg912          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceTrunkIpv6OspfNeighborCfg913           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceTrunkIpv6OspfPriorityCfg914           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceTrunkIpv6OspfRetransmitIntervalCfg915 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceTrunkIpv6OspfTransmitDelayCfg916      `json:"transmit-delay-cfg"`
	Uuid                  string                                           `json:"uuid"`
}

type InterfaceTrunkIpv6OspfNetworkList908 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceTrunkIpv6OspfCostCfg909 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfDeadIntervalCfg910 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfHelloIntervalCfg911 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfMtuIgnoreCfg912 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfNeighborCfg913 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceTrunkIpv6OspfPriorityCfg914 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfRetransmitIntervalCfg915 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfTransmitDelayCfg916 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTrunkIsis917 struct {
	Authentication           InterfaceTrunkIsisAuthentication918             `json:"authentication"`
	BfdCfg                   InterfaceTrunkIsisBfdCfg922                     `json:"bfd-cfg"`
	CircuitType              string                                          `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceTrunkIsisCsnpIntervalList923         `json:"csnp-interval-list"`
	Padding                  int                                             `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceTrunkIsisHelloIntervalList924        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceTrunkIsisHelloIntervalMinimalList925 `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceTrunkIsisHelloMultiplierList926      `json:"hello-multiplier-list"`
	LspInterval              int                                             `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceTrunkIsisMeshGroup927                  `json:"mesh-group"`
	MetricList               []InterfaceTrunkIsisMetricList928               `json:"metric-list"`
	Network                  string                                          `json:"network"`
	PasswordList             []InterfaceTrunkIsisPasswordList929             `json:"password-list"`
	PriorityList             []InterfaceTrunkIsisPriorityList930             `json:"priority-list"`
	RetransmitInterval       int                                             `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceTrunkIsisWideMetricList931           `json:"wide-metric-list"`
	Uuid                     string                                          `json:"uuid"`
}

type InterfaceTrunkIsisAuthentication918 struct {
	SendOnlyList []InterfaceTrunkIsisAuthenticationSendOnlyList919 `json:"send-only-list"`
	ModeList     []InterfaceTrunkIsisAuthenticationModeList920     `json:"mode-list"`
	KeyChainList []InterfaceTrunkIsisAuthenticationKeyChainList921 `json:"key-chain-list"`
}

type InterfaceTrunkIsisAuthenticationSendOnlyList919 struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceTrunkIsisAuthenticationModeList920 struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceTrunkIsisAuthenticationKeyChainList921 struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceTrunkIsisBfdCfg922 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceTrunkIsisCsnpIntervalList923 struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceTrunkIsisHelloIntervalList924 struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceTrunkIsisHelloIntervalMinimalList925 struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceTrunkIsisHelloMultiplierList926 struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceTrunkIsisMeshGroup927 struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceTrunkIsisMetricList928 struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceTrunkIsisPasswordList929 struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceTrunkIsisPriorityList930 struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceTrunkIsisWideMetricList931 struct {
	WideMetric int    `json:"wide-metric" dval:"10"`
	Level      string `json:"level"`
}

type InterfaceTrunkLw4o6932 struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceTrunkMap933 struct {
	Inside      int    `json:"inside"`
	Outside     int    `json:"outside"`
	MapTInside  int    `json:"map-t-inside"`
	MapTOutside int    `json:"map-t-outside"`
	Uuid        string `json:"uuid"`
}

type InterfaceTrunkNptv6934 struct {
	DomainList []InterfaceTrunkNptv6DomainList `json:"domain-list"`
}

type InterfaceTrunkNptv6DomainList struct {
	DomainName string `json:"domain-name"`
	BindType   string `json:"bind-type"`
	Uuid       string `json:"uuid"`
}

type InterfaceTrunkSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type InterfaceTrunkSpanningTree935 struct {
	AutoEdge     int                                         `json:"auto-edge" dval:"1"`
	AdminEdge    int                                         `json:"admin-edge"`
	InstanceList []InterfaceTrunkSpanningTreeInstanceList936 `json:"instance-list"`
	PathCost     int                                         `json:"path-cost"`
	Uuid         string                                      `json:"uuid"`
}

type InterfaceTrunkSpanningTreeInstanceList936 struct {
	InstanceStart int `json:"instance-start"`
	MstpPathCost  int `json:"mstp-path-cost"`
}

func (p *InterfaceTrunk) GetId() string {
	return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceTrunk) getPath() string {
	return "interface/trunk"
}

func (p *InterfaceTrunk) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunk::Post")
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

func (p *InterfaceTrunk) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunk::Get")
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
func (p *InterfaceTrunk) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunk::Put")
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

func (p *InterfaceTrunk) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunk::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
