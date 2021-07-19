package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceEthernet struct {
	UUID InterfaceEthernetInstance `json:"ethernet,omitempty"`
}

type EthernetAddressList struct {
	AddressType string `json:"address-type,omitempty"`
	Ipv6Addr    string `json:"ipv6-addr,omitempty"`
}
type EthernetHelperAddressList struct {
	HelperAddress string `json:"helper-address,omitempty"`
}
type EthernetStatefulFirewallId struct {
	UUID       string `json:"uuid,omitempty"`
	ClassList  string `json:"class-list,omitempty"`
	Inside     int    `json:"inside,omitempty"`
	Outside    int    `json:"outside,omitempty"`
	ACLID      int    `json:"acl-id,omitempty"`
	AccessList int    `json:"access-list,omitempty"`
}
type EthernetReceiveCfg struct {
	Receive int    `json:"receive,omitempty"`
	Version string `json:"version,omitempty"`
}
type EthernetSplitHorizonCfg struct {
	State string `json:"state,omitempty"`
}
type EthernetKeyChain struct {
	KeyChain string `json:"key-chain,omitempty"`
}
type EthernetMode struct {
	Mode string `json:"mode,omitempty"`
}
type EthernetStr struct {
	String string `json:"string,omitempty"`
}
type EthernetAuthentication struct {
	KeyChain EthernetKeyChain `json:"key-chain,omitempty"`
	Mode     EthernetMode     `json:"mode,omitempty"`
	String   EthernetStr      `json:"str,omitempty"`
}
type EthernetSendCfg struct {
	Version string `json:"version,omitempty"`
	Send    int    `json:"send,omitempty"`
}
type EthernetRip struct {
	Receive       EthernetReceiveCfg      `json:"receive-cfg,omitempty"`
	UUID          string                  `json:"uuid,omitempty"`
	ReceivePacket int                     `json:"receive-packet,omitempty"`
	State         EthernetSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	KeyChain      EthernetAuthentication  `json:"authentication,omitempty"`
	Version       EthernetSendCfg         `json:"send-cfg,omitempty"`
	SendPacket    int                     `json:"send-packet,omitempty"`
}
type EthernetIsis struct {
	Tag  string `json:"tag,omitempty"`
	UUID string `json:"uuid,omitempty"`
}
type EthernetRouter struct {
	Tag EthernetIsis `json:"isis,omitempty"`
}
type EthernetMd5 struct {
	Md5Value  string `json:"md5-value,omitempty"`
	Encrypted string `json:"encrypted,omitempty"`
}
type EthernetMessageDigestCfg struct {
	MessageDigestKey int         `json:"message-digest-key,omitempty"`
	Md5Value         EthernetMd5 `json:"md5,omitempty"`
}
type EthernetOspfIPList struct {
	DeadInterval       int                        `json:"dead-interval,omitempty"`
	AuthenticationKey  string                     `json:"authentication-key,omitempty"`
	UUID               string                     `json:"uuid,omitempty"`
	MtuIgnore          int                        `json:"mtu-ignore,omitempty"`
	TransmitDelay      int                        `json:"transmit-delay,omitempty"`
	Value              string                     `json:"value,omitempty"`
	Priority           int                        `json:"priority,omitempty"`
	Authentication     int                        `json:"authentication,omitempty"`
	Cost               int                        `json:"cost,omitempty"`
	DatabaseFilter     string                     `json:"database-filter,omitempty"`
	HelloInterval      int                        `json:"hello-interval,omitempty"`
	IPAddr             string                     `json:"ip-addr,omitempty"`
	RetransmitInterval int                        `json:"retransmit-interval,omitempty"`
	MessageDigestKey   []EthernetMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	Out                int                        `json:"out,omitempty"`
}
type EthernetNetwork struct {
	Broadcast         int `json:"broadcast,omitempty"`
	PointToMultipoint int `json:"point-to-multipoint,omitempty"`
	NonBroadcast      int `json:"non-broadcast,omitempty"`
	PointToPoint      int `json:"point-to-point,omitempty"`
	P2MpNbma          int `json:"p2mp-nbma,omitempty"`
}
type EthernetAuthenticationCfg struct {
	Authentication int    `json:"authentication,omitempty"`
	Value          string `json:"value,omitempty"`
}
type EthernetBfdCfg struct {
	Disable int `json:"disable,omitempty"`
	Bfd     int `json:"bfd,omitempty"`
}
type EthernetDatabaseFilterCfg struct {
	DatabaseFilter string `json:"database-filter,omitempty"`
	Out            int    `json:"out,omitempty"`
}
type EthernetOspfGlobal struct {
	Cost               int                        `json:"cost,omitempty"`
	DeadInterval       int                        `json:"dead-interval,omitempty"`
	AuthenticationKey  string                     `json:"authentication-key,omitempty"`
	Broadcast          EthernetNetwork            `json:"network,omitempty"`
	MtuIgnore          int                        `json:"mtu-ignore,omitempty"`
	TransmitDelay      int                        `json:"transmit-delay,omitempty"`
	Authentication     EthernetAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	RetransmitInterval int                        `json:"retransmit-interval,omitempty"`
	Bfd                EthernetBfdCfg             `json:"bfd-cfg,omitempty"`
	Disable            string                     `json:"disable,omitempty"`
	HelloInterval      int                        `json:"hello-interval,omitempty"`
	DatabaseFilter     EthernetDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	Priority           int                        `json:"priority,omitempty"`
	Mtu                int                        `json:"mtu,omitempty"`
	MessageDigestKey   []EthernetMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	UUID               string                     `json:"uuid,omitempty"`
}
type EthernetOspf struct {
	DeadInterval []EthernetOspfIPList `json:"ospf-ip-list,omitempty"`
	Cost         EthernetOspfGlobal   `json:"ospf-global,omitempty"`
}
type EthernetIP struct {
	UUID                    string                      `json:"uuid,omitempty"`
	AddressType             []EthernetAddressList       `json:"address-list,omitempty"`
	GenerateMembershipQuery int                         `json:"generate-membership-query,omitempty"`
	CacheSpoofingPort       int                         `json:"cache-spoofing-port,omitempty"`
	Inside                  int                         `json:"inside,omitempty"`
	AllowPromiscuousVip     int                         `json:"allow-promiscuous-vip,omitempty"`
	Client                  int                         `json:"client,omitempty"`
	MaxRespTime             int                         `json:"max-resp-time,omitempty"`
	QueryInterval           int                         `json:"query-interval,omitempty"`
	Outside                 int                         `json:"outside,omitempty"`
	HelperAddress           []EthernetHelperAddressList `json:"helper-address-list,omitempty"`
	ClassList               EthernetStatefulFirewallId  `json:"stateful-firewall,omitempty"`
	ReceiveCfg              EthernetRip                 `json:"rip,omitempty"`
	TTLIgnore               int                         `json:"ttl-ignore,omitempty"`
	Tag                     EthernetRouter              `json:"router,omitempty"`
	Dhcp                    int                         `json:"dhcp,omitempty"`
	Server                  int                         `json:"server,omitempty"`
	DeadInterval            EthernetOspf                `json:"ospf,omitempty"`
	SlbPartitionRedirect    int                         `json:"slb-partition-redirect,omitempty"`
}
type EthernetDdos struct {
	Outside int    `json:"outside,omitempty"`
	Inside  int    `json:"inside,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}
type EthernetAccessList struct {
	ACLName string `json:"acl-name,omitempty"`
	ACLID   int    `json:"acl-id,omitempty"`
}
type EthernetTxDot1Cfg struct {
	LinkAggregation int `json:"link-aggregation,omitempty"`
	Vlan            int `json:"vlan,omitempty"`
	TxDot1Tlvs      int `json:"tx-dot1-tlvs,omitempty"`
}
type EthernetNotificationCfg struct {
	Notification int `json:"notification,omitempty"`
	NotifEnable  int `json:"notif-enable,omitempty"`
}
type EthernetEnableCfg struct {
	Rx       int `json:"rx,omitempty"`
	Tx       int `json:"tx,omitempty"`
	RtEnable int `json:"rt-enable,omitempty"`
}
type EthernetTxTlvsCfg struct {
	SystemCapabilities int `json:"system-capabilities,omitempty"`
	SystemDescription  int `json:"system-description,omitempty"`
	ManagementAddress  int `json:"management-address,omitempty"`
	TxTlvs             int `json:"tx-tlvs,omitempty"`
	Exclude            int `json:"exclude,omitempty"`
	PortDescription    int `json:"port-description,omitempty"`
	SystemName         int `json:"system-name,omitempty"`
}
type EthernetLldp struct {
	LinkAggregation    EthernetTxDot1Cfg       `json:"tx-dot1-cfg,omitempty"`
	Notification       EthernetNotificationCfg `json:"notification-cfg,omitempty"`
	Rx                 EthernetEnableCfg       `json:"enable-cfg,omitempty"`
	SystemCapabilities EthernetTxTlvsCfg       `json:"tx-tlvs-cfg,omitempty"`
	UUID               string                  `json:"uuid,omitempty"`
}
type EthernetIntervalCfg struct {
	Interval   int `json:"interval,omitempty"`
	MinRx      int `json:"min-rx,omitempty"`
	Multiplier int `json:"multiplier,omitempty"`
}
type EthernetAuthentication2 struct {
	Encrypted string `json:"encrypted,omitempty"`
	Password  string `json:"password,omitempty"`
	Method    string `json:"method,omitempty"`
	KeyID     int    `json:"key-id,omitempty"`
}
type EthernetBfd struct {
	Interval  EthernetIntervalCfg     `json:"interval-cfg,omitempty"`
	Encrypted EthernetAuthentication2 `json:"authentication,omitempty"`
	Echo      int                     `json:"echo,omitempty"`
	UUID      string                  `json:"uuid,omitempty"`
	Demand    int                     `json:"demand,omitempty"`
}
type EthernetMonitorList struct {
	MonitorVlan int    `json:"monitor-vlan,omitempty"`
	Monitor     string `json:"monitor,omitempty"`
	MirrorIndex int    `json:"mirror-index,omitempty"`
}
type EthernetMap struct {
	Inside      int    `json:"inside,omitempty"`
	MapTInside  int    `json:"map-t-inside,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	MapTOutside int    `json:"map-t-outside,omitempty"`
	Outside     int    `json:"outside,omitempty"`
}
type EthernetUdldTimeoutCfg struct {
	Slow int `json:"slow,omitempty"`
	Fast int `json:"fast,omitempty"`
}
type EthernetTrunkGroupList struct {
	UUID         string                 `json:"uuid,omitempty"`
	TrunkNumber  int                    `json:"trunk-number,omitempty"`
	UserTag      string                 `json:"user-tag,omitempty"`
	Slow         EthernetUdldTimeoutCfg `json:"udld-timeout-cfg,omitempty"`
	Mode         string                 `json:"mode,omitempty"`
	Timeout      string                 `json:"timeout,omitempty"`
	Type         string                 `json:"type,omitempty"`
	AdminKey     int                    `json:"admin-key,omitempty"`
	PortPriority int                    `json:"port-priority,omitempty"`
}
type EthernetDomainList struct {
	DomainName string `json:"domain-name,omitempty"`
	BindType   string `json:"bind-type,omitempty"`
	UUID       string `json:"uuid,omitempty"`
}
type EthernetNptv6 struct {
	DomainName []EthernetDomainList `json:"domain-list,omitempty"`
}
type EthernetPriorityList struct {
	Priority int    `json:"priority,omitempty"`
	Level    string `json:"level,omitempty"`
}
type EthernetHelloIntervalMinimalList struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal,omitempty"`
	Level                string `json:"level,omitempty"`
}
type EthernetMeshGroup struct {
	Value   int `json:"value,omitempty"`
	Blocked int `json:"blocked,omitempty"`
}
type EthernetSendOnlyList struct {
	SendOnly int    `json:"send-only,omitempty"`
	Level    string `json:"level,omitempty"`
}
type EthernetModeList struct {
	Mode  string `json:"mode,omitempty"`
	Level string `json:"level,omitempty"`
}
type EthernetKeyChainList struct {
	KeyChain string `json:"key-chain,omitempty"`
	Level    string `json:"level,omitempty"`
}
type EthernetAuthentication3 struct {
	SendOnly []EthernetSendOnlyList `json:"send-only-list,omitempty"`
	Mode     []EthernetModeList     `json:"mode-list,omitempty"`
	KeyChain []EthernetKeyChainList `json:"key-chain-list,omitempty"`
}
type EthernetCsnpIntervalList struct {
	CsnpInterval int    `json:"csnp-interval,omitempty"`
	Level        string `json:"level,omitempty"`
}
type EthernetPasswordList struct {
	Password string `json:"password,omitempty"`
	Level    string `json:"level,omitempty"`
}
type EthernetWideMetricList struct {
	WideMetric int    `json:"wide-metric,omitempty"`
	Level      string `json:"level,omitempty"`
}
type EthernetHelloIntervalList struct {
	HelloInterval int    `json:"hello-interval,omitempty"`
	Level         string `json:"level,omitempty"`
}
type EthernetHelloMultiplierList struct {
	HelloMultiplier int    `json:"hello-multiplier,omitempty"`
	Level           string `json:"level,omitempty"`
}
type EthernetMetricList struct {
	Metric int    `json:"metric,omitempty"`
	Level  string `json:"level,omitempty"`
}
type EthernetIsis2 struct {
	Priority             []EthernetPriorityList             `json:"priority-list,omitempty"`
	Padding              int                                `json:"padding,omitempty"`
	HelloIntervalMinimal []EthernetHelloIntervalMinimalList `json:"hello-interval-minimal-list,omitempty"`
	Value                EthernetMeshGroup                  `json:"mesh-group,omitempty"`
	Network              string                             `json:"network,omitempty"`
	SendOnly             EthernetAuthentication3            `json:"authentication,omitempty"`
	CsnpInterval         []EthernetCsnpIntervalList         `json:"csnp-interval-list,omitempty"`
	RetransmitInterval   int                                `json:"retransmit-interval,omitempty"`
	Password             []EthernetPasswordList             `json:"password-list,omitempty"`
	Disable              EthernetBfdCfg                     `json:"bfd-cfg,omitempty"`
	WideMetric           []EthernetWideMetricList           `json:"wide-metric-list,omitempty"`
	HelloInterval        []EthernetHelloIntervalList        `json:"hello-interval-list,omitempty"`
	CircuitType          string                             `json:"circuit-type,omitempty"`
	HelloMultiplier      []EthernetHelloMultiplierList      `json:"hello-multiplier-list,omitempty"`
	Metric               []EthernetMetricList               `json:"metric-list,omitempty"`
	LspInterval          int                                `json:"lsp-interval,omitempty"`
	UUID                 string                             `json:"uuid,omitempty"`
}
type EthernetIcmpv6RateLimit struct {
	LockupPeriodV6 int `json:"lockup-period-v6,omitempty"`
	NormalV6       int `json:"normal-v6,omitempty"`
	LockupV6       int `json:"lockup-v6,omitempty"`
}
type EthernetRip2 struct {
	State EthernetSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	UUID  string                  `json:"uuid,omitempty"`
}
type EthernetStatefulFirewall struct {
	UUID       string `json:"uuid,omitempty"`
	ClassList  string `json:"class-list,omitempty"`
	ACLName    string `json:"acl-name,omitempty"`
	Inside     int    `json:"inside,omitempty"`
	Outside    int    `json:"outside,omitempty"`
	AccessList int    `json:"access-list,omitempty"`
}
type EthernetRipng struct {
	UUID string `json:"uuid,omitempty"`
	Rip  int    `json:"rip,omitempty"`
}
type EthernetAreaList struct {
	AreaIDAddr string `json:"area-id-addr,omitempty"`
	Tag        string `json:"tag,omitempty"`
	InstanceID int    `json:"instance-id,omitempty"`
	AreaIDNum  int    `json:"area-id-num,omitempty"`
}
type EthernetOspf2 struct {
	AreaIDAddr []EthernetAreaList `json:"area-list,omitempty"`
	UUID       string             `json:"uuid,omitempty"`
}

type EthernetRouter2 struct {
	UUID       EthernetRipng `json:"ripng,omitempty"`
	AreaIDAddr EthernetOspf2 `json:"ospf,omitempty"`
	Tag        EthernetIsis  `json:"isis,omitempty"`
}
type EthernetAccessListCfg struct {
	Inbound   int    `json:"inbound,omitempty"`
	V6ACLName string `json:"v6-acl-name,omitempty"`
}
type EthernetCostCfg struct {
	Cost       int `json:"cost,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type EthernetPriorityCfg struct {
	Priority   int `json:"priority,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type EthernetHelloIntervalCfg struct {
	HelloInterval int `json:"hello-interval,omitempty"`
	InstanceID    int `json:"instance-id,omitempty"`
}
type EthernetMtuIgnoreCfg struct {
	MtuIgnore  int `json:"mtu-ignore,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type EthernetRetransmitIntervalCfg struct {
	RetransmitInterval int `json:"retransmit-interval,omitempty"`
	InstanceID         int `json:"instance-id,omitempty"`
}
type EthernetTransmitDelayCfg struct {
	TransmitDelay int `json:"transmit-delay,omitempty"`
	InstanceID    int `json:"instance-id,omitempty"`
}
type EthernetNeighborCfg struct {
	NeighborPriority     int    `json:"neighbor-priority,omitempty"`
	NeigInst             int    `json:"neig-inst,omitempty"`
	NeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	NeighborCost         int    `json:"neighbor-cost,omitempty"`
	Neighbor             string `json:"neighbor,omitempty"`
}
type EthernetNetworkList struct {
	BroadcastType     string `json:"broadcast-type,omitempty"`
	P2MpNbma          int    `json:"p2mp-nbma,omitempty"`
	NetworkInstanceID int    `json:"network-instance-id,omitempty"`
}
type EthernetDeadIntervalCfg struct {
	DeadInterval int `json:"dead-interval,omitempty"`
	InstanceID   int `json:"instance-id,omitempty"`
}
type EthernetOspf3 struct {
	UUID               string                          `json:"uuid,omitempty"`
	Bfd                int                             `json:"bfd,omitempty"`
	Cost               []EthernetCostCfg               `json:"cost-cfg,omitempty"`
	Priority           []EthernetPriorityCfg           `json:"priority-cfg,omitempty"`
	HelloInterval      []EthernetHelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	MtuIgnore          []EthernetMtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	RetransmitInterval []EthernetRetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	Disable            int                             `json:"disable,omitempty"`
	TransmitDelay      []EthernetTransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	NeighborPriority   []EthernetNeighborCfg           `json:"neighbor-cfg,omitempty"`
	BroadcastType      []EthernetNetworkList           `json:"network-list,omitempty"`
	DeadInterval       []EthernetDeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
}
type EthernetPrefixList struct {
	NotAutonomous     int    `json:"not-autonomous,omitempty"`
	ValidLifetime     int    `json:"valid-lifetime,omitempty"`
	NotOnLink         int    `json:"not-on-link,omitempty"`
	Prefix            string `json:"prefix,omitempty"`
	PreferredLifetime int    `json:"preferred-lifetime,omitempty"`
}
type EthernetRouterAdver struct {
	MaxInterval              int                  `json:"max-interval,omitempty"`
	DefaultLifetime          int                  `json:"default-lifetime,omitempty"`
	ReachableTime            int                  `json:"reachable-time,omitempty"`
	OtherConfigAction        string               `json:"other-config-action,omitempty"`
	FloatingIPDefaultVrid    string               `json:"floating-ip-default-vrid,omitempty"`
	ManagedConfigAction      string               `json:"managed-config-action,omitempty"`
	MinInterval              int                  `json:"min-interval,omitempty"`
	RateLimit                int                  `json:"rate-limit,omitempty"`
	AdverMtuDisable          int                  `json:"adver-mtu-disable,omitempty"`
	NotAutonomous            []EthernetPrefixList `json:"prefix-list,omitempty"`
	FloatingIP               string               `json:"floating-ip,omitempty"`
	AdverVrid                int                  `json:"adver-vrid,omitempty"`
	UseFloatingIPDefaultVrid int                  `json:"use-floating-ip-default-vrid,omitempty"`
	Action                   string               `json:"action,omitempty"`
	AdverVridDefault         int                  `json:"adver-vrid-default,omitempty"`
	AdverMtu                 int                  `json:"adver-mtu,omitempty"`
	RetransmitTimer          int                  `json:"retransmit-timer,omitempty"`
	HopLimit                 int                  `json:"hop-limit,omitempty"`
	UseFloatingIP            int                  `json:"use-floating-ip,omitempty"`
}
type EthernetIpv6 struct {
	UUID        string                   `json:"uuid,omitempty"`
	AddressType []EthernetAddressList    `json:"address-list,omitempty"`
	Inside      int                      `json:"inside,omitempty"`
	Ipv6Enable  int                      `json:"ipv6-enable,omitempty"`
	State       EthernetRip2             `json:"rip,omitempty"`
	Outside     int                      `json:"outside,omitempty"`
	ClassList   EthernetStatefulFirewall `json:"stateful-firewall,omitempty"`
	TTLIgnore   int                      `json:"ttl-ignore,omitempty"`
	AreaIDAddr  EthernetRouter2          `json:"router,omitempty"`
	Inbound     EthernetAccessListCfg    `json:"access-list-cfg,omitempty"`
	Bfd         EthernetOspf3            `json:"ospf,omitempty"`
	MaxInterval EthernetRouterAdver      `json:"router-adver,omitempty"`
}
type EthernetSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}
type EthernetLw4O6 struct {
	Outside1 int    `json:"outside,omitempty"`
	Inside   int    `json:"inside,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}
type EthernetIcmpRateLimit struct {
	Lockup       int `json:"lockup,omitempty"`
	LockupPeriod int `json:"lockup-period,omitempty"`
	Normal       int `json:"normal,omitempty"`
}
type InterfaceEthernetInstance struct {
	FecForcedOn             int                      `json:"fec-forced-on,omitempty"`
	TrapSource              int                      `json:"trap-source,omitempty"`
	AddressType             EthernetIP               `json:"ip,omitempty"`
	Outside                 EthernetDdos             `json:"ddos,omitempty"`
	L3VlanFwdDisable        int                      `json:"l3-vlan-fwd-disable,omitempty"`
	ACLName                 EthernetAccessList       `json:"access-list,omitempty"`
	Speed                   string                   `json:"speed,omitempty"`
	SpeedForced40G          int                      `json:"speed-forced-40g,omitempty"`
	LinkAggregation         EthernetLldp             `json:"lldp,omitempty"`
	UUID                    string                   `json:"uuid,omitempty"`
	Interval                EthernetBfd              `json:"bfd,omitempty"`
	MediaTypeCopper         int                      `json:"media-type-copper,omitempty"`
	Ifnum                   int                      `json:"ifnum,omitempty"`
	RemoveVlanTag           int                      `json:"remove-vlan-tag,omitempty"`
	MonitorVlan             []EthernetMonitorList    `json:"monitor-list,omitempty"`
	CPUProcess              int                      `json:"cpu-process,omitempty"`
	AutoNegEnable           int                      `json:"auto-neg-enable,omitempty"`
	Inside                  EthernetMap              `json:"map,omitempty"`
	TrafficDistributionMode string                   `json:"traffic-distribution-mode,omitempty"`
	TrunkNumber             []EthernetTrunkGroupList `json:"trunk-group-list,omitempty"`
	Nptv6                   EthernetNptv6            `json:"nptv6,omitempty"`
	CPUProcessDir           string                   `json:"cpu-process-dir,omitempty"`
	Priority                EthernetIsis2            `json:"isis,omitempty"`
	Name                    string                   `json:"name,omitempty"`
	Duplexity               string                   `json:"duplexity,omitempty"`
	LockupPeriodV6          EthernetIcmpv6RateLimit  `json:"icmpv6-rate-limit,omitempty"`
	UserTag                 string                   `json:"user-tag,omitempty"`
	Mtu                     int                      `json:"mtu,omitempty"`
	Ipv6Enable              EthernetIpv6             `json:"ipv6,omitempty"`
	Counters1               []EthernetSamplingEnable `json:"sampling-enable,omitempty"`
	LoadInterval            int                      `json:"load-interval,omitempty"`
	Outside1                EthernetLw4O6            `json:"lw-4o6,omitempty"`
	Action                  string                   `json:"action,omitempty"`
	FecForcedOff            int                      `json:"fec-forced-off,omitempty"`
	Lockup                  EthernetIcmpRateLimit    `json:"icmp-rate-limit,omitempty"`
	FlowControl             int                      `json:"flow-control,omitempty"`
}

func PostInterfaceEthernet(id string, inst InterfaceEthernet, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceEthernet")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernet
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostInterfaceEthernet REQ RES..........................", m)
			check_api_status("PostInterfaceEthernet", data)

		}
	}

}

func GetInterfaceEthernet(id string, name string, host string) (*InterfaceEthernet, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceEthernet")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernet
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetInterfaceEthernet REQ RES..........................", m)
			check_api_status("GetInterfaceEthernet", data)
			return &m, nil
		}
	}

}

func PutInterfaceEthernet(id string, name string, inst InterfaceEthernet, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutInterfaceEthernet")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/interface/ethernet/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernet
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutInterfaceEthernet REQ RES..........................", m)
			check_api_status("PutInterfaceEthernet", data)

		}
	}

}

func DeleteInterfaceEthernet(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteInterfaceEthernet")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/interface/ethernet/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernet
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
