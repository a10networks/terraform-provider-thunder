package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type InterfaceTrunk struct {
	Inst struct {
		AccessList InterfaceTrunkAccessList `json:"access-list"`

		Action string `json:"action" dval:"enable"`

		Bfd InterfaceTrunkBfd868 `json:"bfd"`

		Ddos InterfaceTrunkDdos872 `json:"ddos"`

		DoAutoRecovery int `json:"do-auto-recovery"`

		GamingProtocolCompliance int `json:"gaming-protocol-compliance"`

		IcmpRateLimit InterfaceTrunkIcmpRateLimit `json:"icmp-rate-limit"`

		Icmpv6RateLimit InterfaceTrunkIcmpv6RateLimit `json:"icmpv6-rate-limit"`

		Ifnum int `json:"ifnum"`

		Ip InterfaceTrunkIp873 `json:"ip"`

		Ipv6 InterfaceTrunkIpv6897 `json:"ipv6"`

		Isis InterfaceTrunkIsis922 `json:"isis"`

		L3VlanFwdDisable int `json:"l3-vlan-fwd-disable"`

		Lw4o6 InterfaceTrunkLw4o6937 `json:"lw-4o6"`

		MacLearning string `json:"mac-learning"`

		Map InterfaceTrunkMap938 `json:"map"`

		Mtu int `json:"mtu"`

		Name string `json:"name"`

		Nptv6 InterfaceTrunkNptv6939 `json:"nptv6"`

		PortsThreshold int `json:"ports-threshold"`

		SamplingEnable []InterfaceTrunkSamplingEnable `json:"sampling-enable"`

		SpanningTree InterfaceTrunkSpanningTree940 `json:"spanning-tree"`

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

type InterfaceTrunkBfd868 struct {
	Authentication InterfaceTrunkBfdAuthentication869 `json:"authentication"`
	Echo           int                                `json:"echo"`
	Demand         int                                `json:"demand"`
	IntervalCfg    InterfaceTrunkBfdIntervalCfg870    `json:"interval-cfg"`
	Uuid           string                             `json:"uuid"`
	PerMemberPort  InterfaceTrunkBfdPerMemberPort871  `json:"per-member-port"`
}

type InterfaceTrunkBfdAuthentication869 struct {
	KeyId     int    `json:"key-id"`
	Method    string `json:"method"`
	Password  string `json:"password"`
	Encrypted string `json:"encrypted"`
}

type InterfaceTrunkBfdIntervalCfg870 struct {
	Interval   int `json:"interval"`
	MinRx      int `json:"min-rx"`
	Multiplier int `json:"multiplier"`
}

type InterfaceTrunkBfdPerMemberPort871 struct {
	LocalAddress    string `json:"local-address"`
	NeighborAddress string `json:"neighbor-address"`
	Ipv6Local       string `json:"ipv6-local"`
	Ipv6Nbr         string `json:"ipv6-nbr"`
	Uuid            string `json:"uuid"`
}

type InterfaceTrunkDdos872 struct {
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

type InterfaceTrunkIp873 struct {
	Dhcp                    int                                    `json:"dhcp"`
	AddressList             []InterfaceTrunkIpAddressList874       `json:"address-list"`
	AllowPromiscuousVip     int                                    `json:"allow-promiscuous-vip"`
	Client                  int                                    `json:"client"`
	Server                  int                                    `json:"server"`
	Dmz                     int                                    `json:"dmz"`
	CacheSpoofingPort       int                                    `json:"cache-spoofing-port"`
	HelperAddressList       []InterfaceTrunkIpHelperAddressList875 `json:"helper-address-list"`
	Nat                     InterfaceTrunkIpNat876                 `json:"nat"`
	TtlIgnore               int                                    `json:"ttl-ignore"`
	SynCookie               int                                    `json:"syn-cookie"`
	SlbPartitionRedirect    int                                    `json:"slb-partition-redirect"`
	GenerateMembershipQuery int                                    `json:"generate-membership-query"`
	QueryInterval           int                                    `json:"query-interval" dval:"125"`
	MaxRespTime             int                                    `json:"max-resp-time" dval:"100"`
	Unnumbered              int                                    `json:"unnumbered"`
	Uuid                    string                                 `json:"uuid"`
	Router                  InterfaceTrunkIpRouter877              `json:"router"`
	Rip                     InterfaceTrunkIpRip879                 `json:"rip"`
	Ospf                    InterfaceTrunkIpOspf887                `json:"ospf"`
}

type InterfaceTrunkIpAddressList874 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceTrunkIpHelperAddressList875 struct {
	HelperAddress string `json:"helper-address"`
}

type InterfaceTrunkIpNat876 struct {
	Inside  int `json:"inside"`
	Outside int `json:"outside"`
}

type InterfaceTrunkIpRouter877 struct {
	Isis InterfaceTrunkIpRouterIsis878 `json:"isis"`
}

type InterfaceTrunkIpRouterIsis878 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceTrunkIpRip879 struct {
	Authentication  InterfaceTrunkIpRipAuthentication880  `json:"authentication"`
	SendPacket      int                                   `json:"send-packet" dval:"1"`
	ReceivePacket   int                                   `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceTrunkIpRipSendCfg884         `json:"send-cfg"`
	ReceiveCfg      InterfaceTrunkIpRipReceiveCfg885      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceTrunkIpRipSplitHorizonCfg886 `json:"split-horizon-cfg"`
	Uuid            string                                `json:"uuid"`
}

type InterfaceTrunkIpRipAuthentication880 struct {
	Str      InterfaceTrunkIpRipAuthenticationStr881      `json:"str"`
	Mode     InterfaceTrunkIpRipAuthenticationMode882     `json:"mode"`
	KeyChain InterfaceTrunkIpRipAuthenticationKeyChain883 `json:"key-chain"`
}

type InterfaceTrunkIpRipAuthenticationStr881 struct {
	String string `json:"string"`
}

type InterfaceTrunkIpRipAuthenticationMode882 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceTrunkIpRipAuthenticationKeyChain883 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceTrunkIpRipSendCfg884 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceTrunkIpRipReceiveCfg885 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceTrunkIpRipSplitHorizonCfg886 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceTrunkIpOspf887 struct {
	OspfGlobal InterfaceTrunkIpOspfOspfGlobal888   `json:"ospf-global"`
	OspfIpList []InterfaceTrunkIpOspfOspfIpList895 `json:"ospf-ip-list"`
}

type InterfaceTrunkIpOspfOspfGlobal888 struct {
	AuthenticationCfg  InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg889  `json:"authentication-cfg"`
	AuthenticationKey  string                                              `json:"authentication-key"`
	BfdCfg             InterfaceTrunkIpOspfOspfGlobalBfdCfg890             `json:"bfd-cfg"`
	Cost               int                                                 `json:"cost"`
	DatabaseFilterCfg  InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg891  `json:"database-filter-cfg"`
	DeadInterval       int                                                 `json:"dead-interval" dval:"40"`
	Disable            string                                              `json:"disable"`
	HelloInterval      int                                                 `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg892 `json:"message-digest-cfg"`
	Mtu                int                                                 `json:"mtu"`
	MtuIgnore          int                                                 `json:"mtu-ignore"`
	Network            InterfaceTrunkIpOspfOspfGlobalNetwork894            `json:"network"`
	Priority           int                                                 `json:"priority" dval:"1"`
	RetransmitInterval int                                                 `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                 `json:"transmit-delay" dval:"1"`
	Uuid               string                                              `json:"uuid"`
}

type InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg889 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceTrunkIpOspfOspfGlobalBfdCfg890 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg891 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg892 struct {
	MessageDigestKey int                                                  `json:"message-digest-key"`
	Md5              InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5893 `json:"md5"`
}

type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5893 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceTrunkIpOspfOspfGlobalNetwork894 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceTrunkIpOspfOspfIpList895 struct {
	IpAddr             string                                              `json:"ip-addr"`
	Authentication     int                                                 `json:"authentication"`
	Value              string                                              `json:"value"`
	AuthenticationKey  string                                              `json:"authentication-key"`
	Cost               int                                                 `json:"cost"`
	DatabaseFilter     string                                              `json:"database-filter"`
	Out                int                                                 `json:"out"`
	DeadInterval       int                                                 `json:"dead-interval" dval:"40"`
	HelloInterval      int                                                 `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceTrunkIpOspfOspfIpListMessageDigestCfg896 `json:"message-digest-cfg"`
	MtuIgnore          int                                                 `json:"mtu-ignore"`
	Priority           int                                                 `json:"priority" dval:"1"`
	RetransmitInterval int                                                 `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                 `json:"transmit-delay" dval:"1"`
	Uuid               string                                              `json:"uuid"`
}

type InterfaceTrunkIpOspfOspfIpListMessageDigestCfg896 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceTrunkIpv6897 struct {
	AddressList   []InterfaceTrunkIpv6AddressList898 `json:"address-list"`
	Ipv6Enable    int                                `json:"ipv6-enable"`
	AccessListCfg InterfaceTrunkIpv6AccessListCfg899 `json:"access-list-cfg"`
	Nat           InterfaceTrunkIpv6Nat900           `json:"nat"`
	TtlIgnore     int                                `json:"ttl-ignore"`
	RouterAdver   InterfaceTrunkIpv6RouterAdver901   `json:"router-adver"`
	Uuid          string                             `json:"uuid"`
	Router        InterfaceTrunkIpv6Router905        `json:"router"`
	Rip           InterfaceTrunkIpv6Rip910           `json:"rip"`
	Ospf          InterfaceTrunkIpv6Ospf912          `json:"ospf"`
}

type InterfaceTrunkIpv6AddressList898 struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceTrunkIpv6AccessListCfg899 struct {
	V6AclName string `json:"v6-acl-name"`
	Inbound   int    `json:"inbound"`
}

type InterfaceTrunkIpv6Nat900 struct {
	Inside  int `json:"inside"`
	Outside int `json:"outside"`
}

type InterfaceTrunkIpv6RouterAdver901 struct {
	Action              string                                       `json:"action" dval:"disable"`
	DefaultLifetime     int                                          `json:"default-lifetime" dval:"1800"`
	HopLimit            int                                          `json:"hop-limit" dval:"255"`
	MaxInterval         int                                          `json:"max-interval" dval:"600"`
	MinInterval         int                                          `json:"min-interval" dval:"200"`
	RateLimit           int                                          `json:"rate-limit" dval:"100000"`
	ReachableTime       int                                          `json:"reachable-time"`
	RetransmitTimer     int                                          `json:"retransmit-timer"`
	Mtu                 InterfaceTrunkIpv6RouterAdverMtu902          `json:"mtu"`
	PrefixList          []InterfaceTrunkIpv6RouterAdverPrefixList903 `json:"prefix-list"`
	ManagedConfigAction string                                       `json:"managed-config-action" dval:"disable"`
	OtherConfigAction   string                                       `json:"other-config-action" dval:"disable"`
	Vrid                InterfaceTrunkIpv6RouterAdverVrid904         `json:"vrid"`
}

type InterfaceTrunkIpv6RouterAdverMtu902 struct {
	AdverMtuDisable int `json:"adver-mtu-disable" dval:"1"`
	AdverMtu        int `json:"adver-mtu"`
}

type InterfaceTrunkIpv6RouterAdverPrefixList903 struct {
	Prefix            string `json:"prefix"`
	NotAutonomous     int    `json:"not-autonomous"`
	NotOnLink         int    `json:"not-on-link"`
	PreferredLifetime int    `json:"preferred-lifetime" dval:"604800"`
	ValidLifetime     int    `json:"valid-lifetime" dval:"2592000"`
}

type InterfaceTrunkIpv6RouterAdverVrid904 struct {
	AdverVrid                int    `json:"adver-vrid"`
	UseFloatingIp            int    `json:"use-floating-ip"`
	FloatingIp               string `json:"floating-ip"`
	AdverVridDefault         int    `json:"adver-vrid-default"`
	UseFloatingIpDefaultVrid int    `json:"use-floating-ip-default-vrid"`
	FloatingIpDefaultVrid    string `json:"floating-ip-default-vrid"`
}

type InterfaceTrunkIpv6Router905 struct {
	Ripng InterfaceTrunkIpv6RouterRipng906 `json:"ripng"`
	Ospf  InterfaceTrunkIpv6RouterOspf907  `json:"ospf"`
	Isis  InterfaceTrunkIpv6RouterIsis909  `json:"isis"`
}

type InterfaceTrunkIpv6RouterRipng906 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceTrunkIpv6RouterOspf907 struct {
	AreaList []InterfaceTrunkIpv6RouterOspfAreaList908 `json:"area-list"`
	Uuid     string                                    `json:"uuid"`
}

type InterfaceTrunkIpv6RouterOspfAreaList908 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceTrunkIpv6RouterIsis909 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceTrunkIpv6Rip910 struct {
	SplitHorizonCfg InterfaceTrunkIpv6RipSplitHorizonCfg911 `json:"split-horizon-cfg"`
	Uuid            string                                  `json:"uuid"`
}

type InterfaceTrunkIpv6RipSplitHorizonCfg911 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceTrunkIpv6Ospf912 struct {
	NetworkList           []InterfaceTrunkIpv6OspfNetworkList913           `json:"network-list"`
	Bfd                   int                                              `json:"bfd"`
	Disable               int                                              `json:"disable"`
	CostCfg               []InterfaceTrunkIpv6OspfCostCfg914               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceTrunkIpv6OspfDeadIntervalCfg915       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceTrunkIpv6OspfHelloIntervalCfg916      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceTrunkIpv6OspfMtuIgnoreCfg917          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceTrunkIpv6OspfNeighborCfg918           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceTrunkIpv6OspfPriorityCfg919           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceTrunkIpv6OspfRetransmitIntervalCfg920 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceTrunkIpv6OspfTransmitDelayCfg921      `json:"transmit-delay-cfg"`
	Uuid                  string                                           `json:"uuid"`
}

type InterfaceTrunkIpv6OspfNetworkList913 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceTrunkIpv6OspfCostCfg914 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfDeadIntervalCfg915 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfHelloIntervalCfg916 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfMtuIgnoreCfg917 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfNeighborCfg918 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceTrunkIpv6OspfPriorityCfg919 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfRetransmitIntervalCfg920 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfTransmitDelayCfg921 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTrunkIsis922 struct {
	Authentication           InterfaceTrunkIsisAuthentication923             `json:"authentication"`
	BfdCfg                   InterfaceTrunkIsisBfdCfg927                     `json:"bfd-cfg"`
	CircuitType              string                                          `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceTrunkIsisCsnpIntervalList928         `json:"csnp-interval-list"`
	Padding                  int                                             `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceTrunkIsisHelloIntervalList929        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceTrunkIsisHelloIntervalMinimalList930 `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceTrunkIsisHelloMultiplierList931      `json:"hello-multiplier-list"`
	LspInterval              int                                             `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceTrunkIsisMeshGroup932                  `json:"mesh-group"`
	MetricList               []InterfaceTrunkIsisMetricList933               `json:"metric-list"`
	Network                  string                                          `json:"network"`
	PasswordList             []InterfaceTrunkIsisPasswordList934             `json:"password-list"`
	PriorityList             []InterfaceTrunkIsisPriorityList935             `json:"priority-list"`
	RetransmitInterval       int                                             `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceTrunkIsisWideMetricList936           `json:"wide-metric-list"`
	Uuid                     string                                          `json:"uuid"`
}

type InterfaceTrunkIsisAuthentication923 struct {
	SendOnlyList []InterfaceTrunkIsisAuthenticationSendOnlyList924 `json:"send-only-list"`
	ModeList     []InterfaceTrunkIsisAuthenticationModeList925     `json:"mode-list"`
	KeyChainList []InterfaceTrunkIsisAuthenticationKeyChainList926 `json:"key-chain-list"`
}

type InterfaceTrunkIsisAuthenticationSendOnlyList924 struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceTrunkIsisAuthenticationModeList925 struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceTrunkIsisAuthenticationKeyChainList926 struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceTrunkIsisBfdCfg927 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceTrunkIsisCsnpIntervalList928 struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceTrunkIsisHelloIntervalList929 struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceTrunkIsisHelloIntervalMinimalList930 struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceTrunkIsisHelloMultiplierList931 struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceTrunkIsisMeshGroup932 struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceTrunkIsisMetricList933 struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceTrunkIsisPasswordList934 struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceTrunkIsisPriorityList935 struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceTrunkIsisWideMetricList936 struct {
	WideMetric int    `json:"wide-metric" dval:"10"`
	Level      string `json:"level"`
}

type InterfaceTrunkLw4o6937 struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceTrunkMap938 struct {
	Inside      int    `json:"inside"`
	Outside     int    `json:"outside"`
	MapTInside  int    `json:"map-t-inside"`
	MapTOutside int    `json:"map-t-outside"`
	Uuid        string `json:"uuid"`
}

type InterfaceTrunkNptv6939 struct {
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

type InterfaceTrunkSpanningTree940 struct {
	AutoEdge     int                                         `json:"auto-edge" dval:"1"`
	AdminEdge    int                                         `json:"admin-edge"`
	InstanceList []InterfaceTrunkSpanningTreeInstanceList941 `json:"instance-list"`
	PathCost     int                                         `json:"path-cost"`
	Uuid         string                                      `json:"uuid"`
}

type InterfaceTrunkSpanningTreeInstanceList941 struct {
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
