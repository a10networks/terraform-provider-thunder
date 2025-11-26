package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type InterfaceVe struct {
	Inst struct {
		AccessList InterfaceVeAccessList `json:"access-list"`

		Action string `json:"action" dval:"enable"`

		Bfd InterfaceVeBfd1046 `json:"bfd"`

		Ddos InterfaceVeDdos1049 `json:"ddos"`

		GamingProtocolCompliance int `json:"gaming-protocol-compliance"`

		IcmpRateLimit InterfaceVeIcmpRateLimit `json:"icmp-rate-limit"`

		Icmpv6RateLimit InterfaceVeIcmpv6RateLimit `json:"icmpv6-rate-limit"`

		Ifnum int `json:"ifnum"`

		Ip InterfaceVeIp1050 `json:"ip"`

		Ipv6 InterfaceVeIpv61073 `json:"ipv6"`

		Isis InterfaceVeIsis1094 `json:"isis"`

		L3VlanFwdDisable int `json:"l3-vlan-fwd-disable"`

		Lw4o6 InterfaceVeLw4o61109 `json:"lw-4o6"`

		Map InterfaceVeMap1110 `json:"map"`

		Mtu int `json:"mtu"`

		Name string `json:"name"`

		Nptv6 InterfaceVeNptv61111 `json:"nptv6"`

		PingSweepDetection string `json:"ping-sweep-detection" dval:"disable"`

		PortScanDetection string `json:"port-scan-detection" dval:"disable"`

		SamplingEnable []InterfaceVeSamplingEnable `json:"sampling-enable"`

		TrapSource int `json:"trap-source"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"ve"`
}

type InterfaceVeAccessList struct {
	AclId   int    `json:"acl-id"`
	AclName string `json:"acl-name"`
}

type InterfaceVeBfd1046 struct {
	Authentication InterfaceVeBfdAuthentication1047 `json:"authentication"`
	Echo           int                              `json:"echo"`
	Demand         int                              `json:"demand"`
	IntervalCfg    InterfaceVeBfdIntervalCfg1048    `json:"interval-cfg"`
	Uuid           string                           `json:"uuid"`
}

type InterfaceVeBfdAuthentication1047 struct {
	KeyId     int    `json:"key-id"`
	Method    string `json:"method"`
	Password  string `json:"password"`
	Encrypted string `json:"encrypted"`
}

type InterfaceVeBfdIntervalCfg1048 struct {
	Interval   int `json:"interval"`
	MinRx      int `json:"min-rx"`
	Multiplier int `json:"multiplier"`
}

type InterfaceVeDdos1049 struct {
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

type InterfaceVeIp1050 struct {
	Dhcp                    int                                  `json:"dhcp"`
	AddressList             []InterfaceVeIpAddressList1051       `json:"address-list"`
	AllowPromiscuousVip     int                                  `json:"allow-promiscuous-vip"`
	Client                  int                                  `json:"client"`
	Server                  int                                  `json:"server"`
	Dmz                     int                                  `json:"dmz"`
	HelperAddressList       []InterfaceVeIpHelperAddressList1052 `json:"helper-address-list"`
	Inside                  int                                  `json:"inside"`
	Outside                 int                                  `json:"outside"`
	TtlIgnore               int                                  `json:"ttl-ignore"`
	SynCookie               int                                  `json:"syn-cookie"`
	SlbPartitionRedirect    int                                  `json:"slb-partition-redirect"`
	GenerateMembershipQuery int                                  `json:"generate-membership-query"`
	QueryInterval           int                                  `json:"query-interval" dval:"125"`
	MaxRespTime             int                                  `json:"max-resp-time" dval:"100"`
	Unnumbered              int                                  `json:"unnumbered"`
	Uuid                    string                               `json:"uuid"`
	Router                  InterfaceVeIpRouter1053              `json:"router"`
	Rip                     InterfaceVeIpRip1055                 `json:"rip"`
	Ospf                    InterfaceVeIpOspf1063                `json:"ospf"`
}

type InterfaceVeIpAddressList1051 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceVeIpHelperAddressList1052 struct {
	HelperAddress string `json:"helper-address"`
}

type InterfaceVeIpRouter1053 struct {
	Isis InterfaceVeIpRouterIsis1054 `json:"isis"`
}

type InterfaceVeIpRouterIsis1054 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceVeIpRip1055 struct {
	Authentication  InterfaceVeIpRipAuthentication1056  `json:"authentication"`
	SendPacket      int                                 `json:"send-packet" dval:"1"`
	ReceivePacket   int                                 `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceVeIpRipSendCfg1060         `json:"send-cfg"`
	ReceiveCfg      InterfaceVeIpRipReceiveCfg1061      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceVeIpRipSplitHorizonCfg1062 `json:"split-horizon-cfg"`
	Uuid            string                              `json:"uuid"`
}

type InterfaceVeIpRipAuthentication1056 struct {
	Str      InterfaceVeIpRipAuthenticationStr1057      `json:"str"`
	Mode     InterfaceVeIpRipAuthenticationMode1058     `json:"mode"`
	KeyChain InterfaceVeIpRipAuthenticationKeyChain1059 `json:"key-chain"`
}

type InterfaceVeIpRipAuthenticationStr1057 struct {
	String string `json:"string"`
}

type InterfaceVeIpRipAuthenticationMode1058 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceVeIpRipAuthenticationKeyChain1059 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceVeIpRipSendCfg1060 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceVeIpRipReceiveCfg1061 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceVeIpRipSplitHorizonCfg1062 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceVeIpOspf1063 struct {
	OspfGlobal InterfaceVeIpOspfOspfGlobal1064   `json:"ospf-global"`
	OspfIpList []InterfaceVeIpOspfOspfIpList1071 `json:"ospf-ip-list"`
}

type InterfaceVeIpOspfOspfGlobal1064 struct {
	AuthenticationCfg  InterfaceVeIpOspfOspfGlobalAuthenticationCfg1065  `json:"authentication-cfg"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	BfdCfg             InterfaceVeIpOspfOspfGlobalBfdCfg1066             `json:"bfd-cfg"`
	Cost               int                                               `json:"cost"`
	DatabaseFilterCfg  InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg1067  `json:"database-filter-cfg"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	Disable            string                                            `json:"disable"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceVeIpOspfOspfGlobalMessageDigestCfg1068 `json:"message-digest-cfg"`
	Mtu                int                                               `json:"mtu"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Network            InterfaceVeIpOspfOspfGlobalNetwork1070            `json:"network"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceVeIpOspfOspfGlobalAuthenticationCfg1065 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceVeIpOspfOspfGlobalBfdCfg1066 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg1067 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceVeIpOspfOspfGlobalMessageDigestCfg1068 struct {
	MessageDigestKey int                                                `json:"message-digest-key"`
	Md5              InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd51069 `json:"md5"`
}

type InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd51069 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceVeIpOspfOspfGlobalNetwork1070 struct {
	Broadcast         int `json:"broadcast"`
	NonBroadcast      int `json:"non-broadcast"`
	PointToPoint      int `json:"point-to-point"`
	PointToMultipoint int `json:"point-to-multipoint"`
	P2mpNbma          int `json:"p2mp-nbma"`
}

type InterfaceVeIpOspfOspfIpList1071 struct {
	IpAddr             string                                            `json:"ip-addr"`
	Authentication     int                                               `json:"authentication"`
	Value              string                                            `json:"value"`
	AuthenticationKey  string                                            `json:"authentication-key"`
	Cost               int                                               `json:"cost"`
	DatabaseFilter     string                                            `json:"database-filter"`
	Out                int                                               `json:"out"`
	DeadInterval       int                                               `json:"dead-interval" dval:"40"`
	HelloInterval      int                                               `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceVeIpOspfOspfIpListMessageDigestCfg1072 `json:"message-digest-cfg"`
	MtuIgnore          int                                               `json:"mtu-ignore"`
	Priority           int                                               `json:"priority" dval:"1"`
	RetransmitInterval int                                               `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                               `json:"transmit-delay" dval:"1"`
	Uuid               string                                            `json:"uuid"`
}

type InterfaceVeIpOspfOspfIpListMessageDigestCfg1072 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceVeIpv61073 struct {
	AddressList []InterfaceVeIpv6AddressList1074 `json:"address-list"`
	Ipv6Enable  int                              `json:"ipv6-enable"`
	V6AclName   string                           `json:"v6-acl-name"`
	Inbound     int                              `json:"inbound"`
	Inside      int                              `json:"inside"`
	Outside     int                              `json:"outside"`
	TtlIgnore   int                              `json:"ttl-ignore"`
	RouterAdver InterfaceVeIpv6RouterAdver1075   `json:"router-adver"`
	Uuid        string                           `json:"uuid"`
	Router      InterfaceVeIpv6Router1077        `json:"router"`
	Rip         InterfaceVeIpv6Rip1082           `json:"rip"`
	Ospf        InterfaceVeIpv6Ospf1084          `json:"ospf"`
}

type InterfaceVeIpv6AddressList1074 struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceVeIpv6RouterAdver1075 struct {
	Action                   string                                     `json:"action" dval:"disable"`
	DefaultLifetime          int                                        `json:"default-lifetime" dval:"1800"`
	HopLimit                 int                                        `json:"hop-limit" dval:"255"`
	MaxInterval              int                                        `json:"max-interval" dval:"600"`
	MinInterval              int                                        `json:"min-interval" dval:"200"`
	RateLimit                int                                        `json:"rate-limit" dval:"100000"`
	ReachableTime            int                                        `json:"reachable-time"`
	RetransmitTimer          int                                        `json:"retransmit-timer"`
	AdverMtuDisable          int                                        `json:"adver-mtu-disable" dval:"1"`
	AdverMtu                 int                                        `json:"adver-mtu"`
	PrefixList               []InterfaceVeIpv6RouterAdverPrefixList1076 `json:"prefix-list"`
	ManagedConfigAction      string                                     `json:"managed-config-action" dval:"disable"`
	OtherConfigAction        string                                     `json:"other-config-action" dval:"disable"`
	AdverVrid                int                                        `json:"adver-vrid"`
	UseFloatingIp            int                                        `json:"use-floating-ip"`
	FloatingIp               string                                     `json:"floating-ip"`
	AdverVridDefault         int                                        `json:"adver-vrid-default"`
	UseFloatingIpDefaultVrid int                                        `json:"use-floating-ip-default-vrid"`
	FloatingIpDefaultVrid    string                                     `json:"floating-ip-default-vrid"`
}

type InterfaceVeIpv6RouterAdverPrefixList1076 struct {
	Prefix            string `json:"prefix"`
	NotAutonomous     int    `json:"not-autonomous"`
	NotOnLink         int    `json:"not-on-link"`
	PreferredLifetime int    `json:"preferred-lifetime" dval:"604800"`
	ValidLifetime     int    `json:"valid-lifetime" dval:"2592000"`
}

type InterfaceVeIpv6Router1077 struct {
	Ripng InterfaceVeIpv6RouterRipng1078 `json:"ripng"`
	Ospf  InterfaceVeIpv6RouterOspf1079  `json:"ospf"`
	Isis  InterfaceVeIpv6RouterIsis1081  `json:"isis"`
}

type InterfaceVeIpv6RouterRipng1078 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceVeIpv6RouterOspf1079 struct {
	AreaList []InterfaceVeIpv6RouterOspfAreaList1080 `json:"area-list"`
	Uuid     string                                  `json:"uuid"`
}

type InterfaceVeIpv6RouterOspfAreaList1080 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceVeIpv6RouterIsis1081 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceVeIpv6Rip1082 struct {
	SplitHorizonCfg InterfaceVeIpv6RipSplitHorizonCfg1083 `json:"split-horizon-cfg"`
	Uuid            string                                `json:"uuid"`
}

type InterfaceVeIpv6RipSplitHorizonCfg1083 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceVeIpv6Ospf1084 struct {
	NetworkList           []InterfaceVeIpv6OspfNetworkList1085           `json:"network-list"`
	Bfd                   int                                            `json:"bfd"`
	Disable               int                                            `json:"disable"`
	CostCfg               []InterfaceVeIpv6OspfCostCfg1086               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceVeIpv6OspfDeadIntervalCfg1087       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceVeIpv6OspfHelloIntervalCfg1088      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceVeIpv6OspfMtuIgnoreCfg1089          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceVeIpv6OspfNeighborCfg1090           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceVeIpv6OspfPriorityCfg1091           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceVeIpv6OspfRetransmitIntervalCfg1092 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceVeIpv6OspfTransmitDelayCfg1093      `json:"transmit-delay-cfg"`
	Uuid                  string                                         `json:"uuid"`
}

type InterfaceVeIpv6OspfNetworkList1085 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceVeIpv6OspfCostCfg1086 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceVeIpv6OspfDeadIntervalCfg1087 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceVeIpv6OspfHelloIntervalCfg1088 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceVeIpv6OspfMtuIgnoreCfg1089 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceVeIpv6OspfNeighborCfg1090 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceVeIpv6OspfPriorityCfg1091 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceVeIpv6OspfRetransmitIntervalCfg1092 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceVeIpv6OspfTransmitDelayCfg1093 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceVeIsis1094 struct {
	Authentication           InterfaceVeIsisAuthentication1095             `json:"authentication"`
	BfdCfg                   InterfaceVeIsisBfdCfg1099                     `json:"bfd-cfg"`
	CircuitType              string                                        `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceVeIsisCsnpIntervalList1100         `json:"csnp-interval-list"`
	Padding                  int                                           `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceVeIsisHelloIntervalList1101        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceVeIsisHelloIntervalMinimalList1102 `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceVeIsisHelloMultiplierList1103      `json:"hello-multiplier-list"`
	LspInterval              int                                           `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceVeIsisMeshGroup1104                  `json:"mesh-group"`
	MetricList               []InterfaceVeIsisMetricList1105               `json:"metric-list"`
	Network                  string                                        `json:"network"`
	PasswordList             []InterfaceVeIsisPasswordList1106             `json:"password-list"`
	PriorityList             []InterfaceVeIsisPriorityList1107             `json:"priority-list"`
	RetransmitInterval       int                                           `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceVeIsisWideMetricList1108           `json:"wide-metric-list"`
	Uuid                     string                                        `json:"uuid"`
}

type InterfaceVeIsisAuthentication1095 struct {
	SendOnlyList []InterfaceVeIsisAuthenticationSendOnlyList1096 `json:"send-only-list"`
	ModeList     []InterfaceVeIsisAuthenticationModeList1097     `json:"mode-list"`
	KeyChainList []InterfaceVeIsisAuthenticationKeyChainList1098 `json:"key-chain-list"`
}

type InterfaceVeIsisAuthenticationSendOnlyList1096 struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceVeIsisAuthenticationModeList1097 struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceVeIsisAuthenticationKeyChainList1098 struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceVeIsisBfdCfg1099 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceVeIsisCsnpIntervalList1100 struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceVeIsisHelloIntervalList1101 struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceVeIsisHelloIntervalMinimalList1102 struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceVeIsisHelloMultiplierList1103 struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceVeIsisMeshGroup1104 struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceVeIsisMetricList1105 struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceVeIsisPasswordList1106 struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceVeIsisPriorityList1107 struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceVeIsisWideMetricList1108 struct {
	WideMetric int    `json:"wide-metric" dval:"10"`
	Level      string `json:"level"`
}

type InterfaceVeLw4o61109 struct {
	Outside int    `json:"outside"`
	Inside  int    `json:"inside"`
	Uuid    string `json:"uuid"`
}

type InterfaceVeMap1110 struct {
	Inside      int    `json:"inside"`
	Outside     int    `json:"outside"`
	MapTInside  int    `json:"map-t-inside"`
	MapTOutside int    `json:"map-t-outside"`
	Uuid        string `json:"uuid"`
}

type InterfaceVeNptv61111 struct {
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
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
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
