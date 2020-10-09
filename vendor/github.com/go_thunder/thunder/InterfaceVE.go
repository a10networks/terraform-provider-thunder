package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceVE struct {
	UUID InterfaceVEInstance `json:"ve,omitempty"`
}

type VeMap struct {
	Inside      int    `json:"inside,omitempty"`
	MapTInside  int    `json:"map-t-inside,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	MapTOutside int    `json:"map-t-outside,omitempty"`
	Outside     int    `json:"outside,omitempty"`
}
type VeDomainList struct {
	DomainName string `json:"domain-name,omitempty"`
	BindType   string `json:"bind-type,omitempty"`
	UUID       string `json:"uuid,omitempty"`
}
type VeNptv6 struct {
	DomainName []VeDomainList `json:"domain-list,omitempty"`
}
type VePriorityList struct {
	Priority int    `json:"priority,omitempty"`
	Level    string `json:"level,omitempty"`
}
type VeHelloIntervalMinimalList struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal,omitempty"`
	Level                string `json:"level,omitempty"`
}
type VeMeshGroup struct {
	Value   int `json:"value,omitempty"`
	Blocked int `json:"blocked,omitempty"`
}
type VeSendOnlyList struct {
	SendOnly int    `json:"send-only,omitempty"`
	Level    string `json:"level,omitempty"`
}
type VeModeList struct {
	Mode  string `json:"mode,omitempty"`
	Level string `json:"level,omitempty"`
}
type VeKeyChainList struct {
	KeyChain string `json:"key-chain,omitempty"`
	Level    string `json:"level,omitempty"`
}
type VeAuthentication struct {
	SendOnly []VeSendOnlyList `json:"send-only-list,omitempty"`
	Mode     []VeModeList     `json:"mode-list,omitempty"`
	KeyChain []VeKeyChainList `json:"key-chain-list,omitempty"`
}
type VeCsnpIntervalList struct {
	CsnpInterval int    `json:"csnp-interval,omitempty"`
	Level        string `json:"level,omitempty"`
}
type VePasswordList struct {
	Password string `json:"password,omitempty"`
	Level    string `json:"level,omitempty"`
}
type VeBfdCfg struct {
	Disable int `json:"disable,omitempty"`
	Bfd     int `json:"bfd,omitempty"`
}
type VeWideMetricList struct {
	WideMetric int    `json:"wide-metric,omitempty"`
	Level      string `json:"level,omitempty"`
}
type VeHelloIntervalList struct {
	HelloInterval int    `json:"hello-interval,omitempty"`
	Level         string `json:"level,omitempty"`
}
type VeHelloMultiplierList struct {
	HelloMultiplier int    `json:"hello-multiplier,omitempty"`
	Level           string `json:"level,omitempty"`
}
type VeMetricList struct {
	Metric int    `json:"metric,omitempty"`
	Level  string `json:"level,omitempty"`
}
type VeIsis struct {
	Priority             []VePriorityList             `json:"priority-list,omitempty"`
	Padding              int                          `json:"padding,omitempty"`
	HelloIntervalMinimal []VeHelloIntervalMinimalList `json:"hello-interval-minimal-list,omitempty"`
	Value                VeMeshGroup                  `json:"mesh-group,omitempty"`
	Network              string                       `json:"network,omitempty"`
	SendOnly             VeAuthentication             `json:"authentication,omitempty"`
	CsnpInterval         []VeCsnpIntervalList         `json:"csnp-interval-list,omitempty"`
	RetransmitInterval   int                          `json:"retransmit-interval,omitempty"`
	Password             []VePasswordList             `json:"password-list,omitempty"`
	Disable              VeBfdCfg                     `json:"bfd-cfg,omitempty"`
	WideMetric           []VeWideMetricList           `json:"wide-metric-list,omitempty"`
	HelloInterval        []VeHelloIntervalList        `json:"hello-interval-list,omitempty"`
	CircuitType          string                       `json:"circuit-type,omitempty"`
	HelloMultiplier      []VeHelloMultiplierList      `json:"hello-multiplier-list,omitempty"`
	Metric               []VeMetricList               `json:"metric-list,omitempty"`
	LspInterval          int                          `json:"lsp-interval,omitempty"`
	UUID                 string                       `json:"uuid,omitempty"`
}
type VeIntervalCfg struct {
	Interval   int `json:"interval,omitempty"`
	MinRx      int `json:"min-rx,omitempty"`
	Multiplier int `json:"multiplier,omitempty"`
}
type VeAuthentication1 struct {
	Encrypted string `json:"encrypted,omitempty"`
	Password  string `json:"password,omitempty"`
	Method    string `json:"method,omitempty"`
	KeyID     int    `json:"key-id,omitempty"`
}
type VeBfd struct {
	Interval  VeIntervalCfg     `json:"interval-cfg,omitempty"`
	Encrypted VeAuthentication1 `json:"authentication,omitempty"`
	Echo      int               `json:"echo,omitempty"`
	UUID      string            `json:"uuid,omitempty"`
	Demand    int               `json:"demand,omitempty"`
}
type VeAddressList struct {
	AddressType string `json:"address-type,omitempty"`
	Ipv6Addr    string `json:"ipv6-addr,omitempty"`
}
type VeHelperAddressList struct {
	HelperAddress string `json:"helper-address,omitempty"`
}
type VeStatefulFirewall struct {
	UUID       string `json:"uuid,omitempty"`
	ClassList  string `json:"class-list,omitempty"`
	Inside     int    `json:"inside,omitempty"`
	Outside    int    `json:"outside,omitempty"`
	ACLID      int    `json:"acl-id,omitempty"`
	AccessList int    `json:"access-list,omitempty"`
}
type VeReceiveCfg struct {
	Receive int    `json:"receive,omitempty"`
	Version string `json:"version,omitempty"`
}
type VeSplitHorizonCfg struct {
	State string `json:"state,omitempty"`
}
type VeKeyChain struct {
	KeyChain string `json:"key-chain,omitempty"`
}
type VeMode struct {
	Mode string `json:"mode,omitempty"`
}
type VeStr struct {
	String string `json:"string,omitempty"`
}
type VeAuthentication2 struct {
	KeyChain VeKeyChain `json:"key-chain,omitempty"`
	Mode     VeMode     `json:"mode,omitempty"`
	String   VeStr      `json:"str,omitempty"`
}
type VeSendCfg struct {
	Version string `json:"version,omitempty"`
	Send    int    `json:"send,omitempty"`
}
type VeRip struct {
	Receive       VeReceiveCfg      `json:"receive-cfg,omitempty"`
	UUID          string            `json:"uuid,omitempty"`
	ReceivePacket int               `json:"receive-packet,omitempty"`
	State         VeSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	KeyChain      VeAuthentication2 `json:"authentication,omitempty"`
	Version       VeSendCfg         `json:"send-cfg,omitempty"`
	SendPacket    int               `json:"send-packet,omitempty"`
}
type VeIsis1 struct {
	Tag  string `json:"tag,omitempty"`
	UUID string `json:"uuid,omitempty"`
}
type VeRouter struct {
	Tag VeIsis1 `json:"isis,omitempty"`
}
type VeMd5 struct {
	Md5Value  string `json:"md5-value,omitempty"`
	Encrypted string `json:"encrypted,omitempty"`
}
type VeMessageDigestCfg struct {
	MessageDigestKey int   `json:"message-digest-key,omitempty"`
	Md5Value         VeMd5 `json:"md5,omitempty"`
}
type VeOspfIPList struct {
	DeadInterval       int                  `json:"dead-interval,omitempty"`
	AuthenticationKey  string               `json:"authentication-key,omitempty"`
	UUID               string               `json:"uuid,omitempty"`
	MtuIgnore          int                  `json:"mtu-ignore,omitempty"`
	TransmitDelay      int                  `json:"transmit-delay,omitempty"`
	Value              string               `json:"value,omitempty"`
	Priority           int                  `json:"priority,omitempty"`
	Authentication     int                  `json:"authentication,omitempty"`
	Cost               int                  `json:"cost,omitempty"`
	DatabaseFilter     string               `json:"database-filter,omitempty"`
	HelloInterval      int                  `json:"hello-interval,omitempty"`
	IPAddr             string               `json:"ip-addr,omitempty"`
	RetransmitInterval int                  `json:"retransmit-interval,omitempty"`
	MessageDigestKey   []VeMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	Out                int                  `json:"out,omitempty"`
}
type VeNetwork struct {
	Broadcast         int `json:"broadcast,omitempty"`
	PointToMultipoint int `json:"point-to-multipoint,omitempty"`
	NonBroadcast      int `json:"non-broadcast,omitempty"`
	PointToPoint      int `json:"point-to-point,omitempty"`
	P2MpNbma          int `json:"p2mp-nbma,omitempty"`
}
type VeAuthenticationCfg struct {
	Authentication int    `json:"authentication,omitempty"`
	Value          string `json:"value,omitempty"`
}
type VeDatabaseFilterCfg struct {
	DatabaseFilter string `json:"database-filter,omitempty"`
	Out            int    `json:"out,omitempty"`
}
type VeOspfGlobal struct {
	Cost               int                  `json:"cost,omitempty"`
	DeadInterval       int                  `json:"dead-interval,omitempty"`
	AuthenticationKey  string               `json:"authentication-key,omitempty"`
	Broadcast          VeNetwork            `json:"network,omitempty"`
	MtuIgnore          int                  `json:"mtu-ignore,omitempty"`
	TransmitDelay      int                  `json:"transmit-delay,omitempty"`
	Authentication     VeAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	RetransmitInterval int                  `json:"retransmit-interval,omitempty"`
	Bfd                VeBfdCfg             `json:"bfd-cfg,omitempty"`
	Disable            string               `json:"disable,omitempty"`
	HelloInterval      int                  `json:"hello-interval,omitempty"`
	DatabaseFilter     VeDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	Priority           int                  `json:"priority,omitempty"`
	Mtu                int                  `json:"mtu,omitempty"`
	MessageDigestKey   []VeMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	UUID               string               `json:"uuid,omitempty"`
}
type VeOspf struct {
	DeadInterval []VeOspfIPList `json:"ospf-ip-list,omitempty"`
	Cost         VeOspfGlobal   `json:"ospf-global,omitempty"`
}
type VeIP1 struct {
	UUID                    string                `json:"uuid,omitempty"`
	GenerateMembershipQuery int                   `json:"generate-membership-query,omitempty"`
	AddressType             []VeAddressList       `json:"address-list,omitempty"`
	Inside                  int                   `json:"inside,omitempty"`
	AllowPromiscuousVip     int                   `json:"allow-promiscuous-vip,omitempty"`
	HelperAddress           []VeHelperAddressList `json:"helper-address-list,omitempty"`
	MaxRespTime             int                   `json:"max-resp-time,omitempty"`
	QueryInterval           int                   `json:"query-interval,omitempty"`
	Outside                 int                   `json:"outside,omitempty"`
	Client                  int                   `json:"client,omitempty"`
	ClassList               VeStatefulFirewall    `json:"stateful-firewall,omitempty"`
	Receive                 VeRip                 `json:"rip,omitempty"`
	TTLIgnore               int                   `json:"ttl-ignore,omitempty"`
	Tag                     VeRouter              `json:"router,omitempty"`
	Dhcp                    int                   `json:"dhcp,omitempty"`
	Server                  int                   `json:"server,omitempty"`
	DeadInterval            VeOspf                `json:"ospf,omitempty"`
	SlbPartitionRedirect    int                   `json:"slb-partition-redirect,omitempty"`
}
type VeIcmpv6RateLimit struct {
	LockupPeriodV6 int `json:"lockup-period-v6,omitempty"`
	NormalV6       int `json:"normal-v6,omitempty"`
	LockupV6       int `json:"lockup-v6,omitempty"`
}
type VeSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}
type VeLw4O6 struct {
	Outside int    `json:"outside,omitempty"`
	Inside  int    `json:"inside,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}
type VeRip1 struct {
	State VeSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	UUID  string            `json:"uuid,omitempty"`
}
type VeStatefulFirewall1 struct {
	UUID       string `json:"uuid,omitempty"`
	ClassList  string `json:"class-list,omitempty"`
	ACLName    string `json:"acl-name,omitempty"`
	Inside     int    `json:"inside,omitempty"`
	Outside    int    `json:"outside,omitempty"`
	AccessList int    `json:"access-list,omitempty"`
}
type VeRipng struct {
	UUID string `json:"uuid,omitempty"`
	Rip  int    `json:"rip,omitempty"`
}
type VeAreaList struct {
	AreaIDAddr string `json:"area-id-addr,omitempty"`
	Tag        string `json:"tag,omitempty"`
	InstanceID int    `json:"instance-id,omitempty"`
	AreaIDNum  int    `json:"area-id-num,omitempty"`
}
type VeOspf1 struct {
	AreaIDAddr []VeAreaList `json:"area-list,omitempty"`
	UUID       string       `json:"uuid,omitempty"`
}
type VeRouter1 struct {
	Rip        VeRipng `json:"ripng,omitempty"`
	AreaIDAddr VeOspf1 `json:"ospf,omitempty"`
	Tag        VeIsis1 `json:"isis,omitempty"`
}
type VeCostCfg struct {
	Cost       int `json:"cost,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type VePriorityCfg struct {
	Priority   int `json:"priority,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type VeHelloIntervalCfg struct {
	HelloInterval int `json:"hello-interval,omitempty"`
	InstanceID    int `json:"instance-id,omitempty"`
}
type VeMtuIgnoreCfg struct {
	MtuIgnore  int `json:"mtu-ignore,omitempty"`
	InstanceID int `json:"instance-id,omitempty"`
}
type VeRetransmitIntervalCfg struct {
	RetransmitInterval int `json:"retransmit-interval,omitempty"`
	InstanceID         int `json:"instance-id,omitempty"`
}
type VeTransmitDelayCfg struct {
	TransmitDelay int `json:"transmit-delay,omitempty"`
	InstanceID    int `json:"instance-id,omitempty"`
}
type VeNeighborCfg struct {
	NeighborPriority     int    `json:"neighbor-priority,omitempty"`
	NeigInst             int    `json:"neig-inst,omitempty"`
	NeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	NeighborCost         int    `json:"neighbor-cost,omitempty"`
	Neighbor             string `json:"neighbor,omitempty"`
}
type VeNetworkList struct {
	BroadcastType     string `json:"broadcast-type,omitempty"`
	P2MpNbma          int    `json:"p2mp-nbma,omitempty"`
	NetworkInstanceID int    `json:"network-instance-id,omitempty"`
}
type VeDeadIntervalCfg struct {
	DeadInterval int `json:"dead-interval,omitempty"`
	InstanceID   int `json:"instance-id,omitempty"`
}
type VeOspf2 struct {
	UUID               string                    `json:"uuid,omitempty"`
	Bfd                int                       `json:"bfd,omitempty"`
	Cost               []VeCostCfg               `json:"cost-cfg,omitempty"`
	Priority           []VePriorityCfg           `json:"priority-cfg,omitempty"`
	HelloInterval      []VeHelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	MtuIgnore          []VeMtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	RetransmitInterval []VeRetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	Disable            int                       `json:"disable,omitempty"`
	TransmitDelay      []VeTransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	NeighborPriority   []VeNeighborCfg           `json:"neighbor-cfg,omitempty"`
	BroadcastType      []VeNetworkList           `json:"network-list,omitempty"`
	DeadInterval       []VeDeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
}
type VePrefixList struct {
	NotAutonomous     int    `json:"not-autonomous,omitempty"`
	ValidLifetime     int    `json:"valid-lifetime,omitempty"`
	NotOnLink         int    `json:"not-on-link,omitempty"`
	Prefix            string `json:"prefix,omitempty"`
	PreferredLifetime int    `json:"preferred-lifetime,omitempty"`
}
type VeRouterAdver struct {
	MaxInterval              int            `json:"max-interval,omitempty"`
	DefaultLifetime          int            `json:"default-lifetime,omitempty"`
	ReachableTime            int            `json:"reachable-time,omitempty"`
	OtherConfigAction        string         `json:"other-config-action,omitempty"`
	FloatingIPDefaultVrid    string         `json:"floating-ip-default-vrid,omitempty"`
	ManagedConfigAction      string         `json:"managed-config-action,omitempty"`
	MinInterval              int            `json:"min-interval,omitempty"`
	RateLimit                int            `json:"rate-limit,omitempty"`
	AdverMtuDisable          int            `json:"adver-mtu-disable,omitempty"`
	NotAutonomous            []VePrefixList `json:"prefix-list,omitempty"`
	FloatingIP               string         `json:"floating-ip,omitempty"`
	AdverVrid                int            `json:"adver-vrid,omitempty"`
	UseFloatingIPDefaultVrid int            `json:"use-floating-ip-default-vrid,omitempty"`
	Action                   string         `json:"action,omitempty"`
	AdverVridDefault         int            `json:"adver-vrid-default,omitempty"`
	AdverMtu                 int            `json:"adver-mtu,omitempty"`
	RetransmitTimer          int            `json:"retransmit-timer,omitempty"`
	HopLimit                 int            `json:"hop-limit,omitempty"`
	UseFloatingIP            int            `json:"use-floating-ip,omitempty"`
}
type VeIpv61 struct {
	UUID        string              `json:"uuid,omitempty"`
	Inbound     int                 `json:"inbound,omitempty"`
	AddressType []VeAddressList     `json:"address-list,omitempty"`
	Inside      int                 `json:"inside,omitempty"`
	Ipv6Enable  int                 `json:"ipv6-enable,omitempty"`
	State       VeRip1              `json:"rip,omitempty"`
	Outside     int                 `json:"outside,omitempty"`
	ClassList   VeStatefulFirewall1 `json:"stateful-firewall,omitempty"`
	V6ACLName   string              `json:"v6-acl-name,omitempty"`
	TTLIgnore   int                 `json:"ttl-ignore,omitempty"`
	Rip         VeRouter1           `json:"router,omitempty"`
	Bfd         VeOspf2             `json:"ospf,omitempty"`
	MaxInterval VeRouterAdver       `json:"router-adver,omitempty"`
}
type VeAccessList struct {
	ACLName string `json:"acl-name,omitempty"`
	ACLID   int    `json:"acl-id,omitempty"`
}
type VeIcmpRateLimit struct {
	Lockup       int `json:"lockup,omitempty"`
	LockupPeriod int `json:"lockup-period,omitempty"`
	Normal       int `json:"normal,omitempty"`
}
type VeDdos struct {
	Outside1 int    `json:"outside,omitempty"`
	Inside   int    `json:"inside,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}
type InterfaceVEInstance struct {
	Inside                  VeMap              `json:"map,omitempty"`
	DomainName              VeNptv6            `json:"nptv6,omitempty"`
	Priority                VeIsis             `json:"isis,omitempty"`
	Name                    string             `json:"name,omitempty"`
	TrapSource              int                `json:"trap-source,omitempty"`
	Interval                VeBfd              `json:"bfd,omitempty"`
	GenerateMembershipQuery VeIP1              `json:"ip,omitempty"`
	LockupPeriodV6          VeIcmpv6RateLimit  `json:"icmpv6-rate-limit,omitempty"`
	UserTag                 string             `json:"user-tag,omitempty"`
	Mtu                     int                `json:"mtu,omitempty"`
	Action                  string             `json:"action,omitempty"`
	Ifnum                   int                `json:"ifnum,omitempty"`
	Counters1               []VeSamplingEnable `json:"sampling-enable,omitempty"`
	Outside                 VeLw4O6            `json:"lw-4o6,omitempty"`
	Inbound                 VeIpv61            `json:"ipv6,omitempty"`
	ACLName                 VeAccessList       `json:"access-list,omitempty"`
	L3VlanFwdDisable        int                `json:"l3-vlan-fwd-disable,omitempty"`
	Lockup                  VeIcmpRateLimit    `json:"icmp-rate-limit,omitempty"`
	Outside1                VeDdos             `json:"ddos,omitempty"`
	UUID                    string             `json:"uuid,omitempty"`
}

func PostInterfaceVE(id string, inst InterfaceVE, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceVE")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ve", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVE
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetInterfaceVE(id string, name string, host string) (*InterfaceVE, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceVE")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ve/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVE
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}

func PutInterfaceVE(id string, name string, inst InterfaceVE, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutInterfaceVE")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/interface/ve/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVE
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteInterfaceVE(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteInterfaceVE")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/interface/ve/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVE
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
