package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 5_2_1-P4_90
type InterfaceLoopback struct {
	Inst struct {
		Ifnum      int                         `json:"ifnum" dval:"-1"`
		Ip         InterfaceLoopbackIp         `json:"ip"`
		Ipv6       InterfaceLoopbackIpv6       `json:"ipv6"`
		Isis       InterfaceLoopbackIsis       `json:"isis"`
		Name       string                      `json:"name"`
		SnmpServer InterfaceLoopbackSnmpServer `json:"snmp-server"`
		UserTag    string                      `json:"user-tag"`
		Uuid       string                      `json:"uuid"`
	} `json:"loopback"`
}

type InterfaceLoopbackIp struct {
	AddressList []InterfaceLoopbackIpAddressList `json:"address-list"`
	Uuid        string                           `json:"uuid"`
	Router      InterfaceLoopbackIpRouter        `json:"router"`
	Rip         InterfaceLoopbackIpRip           `json:"rip"`
	Ospf        InterfaceLoopbackIpOspf          `json:"ospf"`
}

type InterfaceLoopbackIpAddressList struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceLoopbackIpRouter struct {
	Isis InterfaceLoopbackIpRouterIsis `json:"isis"`
}

type InterfaceLoopbackIpRouterIsis struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpRip struct {
	Authentication  InterfaceLoopbackIpRipAuthentication  `json:"authentication"`
	SendPacket      int                                   `json:"send-packet" dval:"1"`
	ReceivePacket   int                                   `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceLoopbackIpRipSendCfg         `json:"send-cfg"`
	ReceiveCfg      InterfaceLoopbackIpRipReceiveCfg      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceLoopbackIpRipSplitHorizonCfg `json:"split-horizon-cfg"`
	Uuid            string                                `json:"uuid"`
}

type InterfaceLoopbackIpRipAuthentication struct {
	Str      InterfaceLoopbackIpRipAuthenticationStr      `json:"str"`
	Mode     InterfaceLoopbackIpRipAuthenticationMode     `json:"mode"`
	KeyChain InterfaceLoopbackIpRipAuthenticationKeyChain `json:"key-chain"`
}

type InterfaceLoopbackIpRipAuthenticationStr struct {
	String string `json:"string"`
}

type InterfaceLoopbackIpRipAuthenticationMode struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceLoopbackIpRipAuthenticationKeyChain struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceLoopbackIpRipSendCfg struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceLoopbackIpRipReceiveCfg struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceLoopbackIpRipSplitHorizonCfg struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLoopbackIpOspf struct {
	OspfGlobal InterfaceLoopbackIpOspfOspfGlobal   `json:"ospf-global"`
	OspfIpList []InterfaceLoopbackIpOspfOspfIpList `json:"ospf-ip-list"`
}

type InterfaceLoopbackIpOspfOspfGlobal struct {
	AuthenticationCfg  InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg  `json:"authentication-cfg"`
	AuthenticationKey  string                                              `json:"authentication-key"`
	BfdCfg             InterfaceLoopbackIpOspfOspfGlobalBfdCfg             `json:"bfd-cfg"`
	Cost               int                                                 `json:"cost"`
	DatabaseFilterCfg  InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg  `json:"database-filter-cfg"`
	DeadInterval       int                                                 `json:"dead-interval" dval:"40"`
	Disable            string                                              `json:"disable"`
	HelloInterval      int                                                 `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg"`
	Mtu                int                                                 `json:"mtu"`
	MtuIgnore          int                                                 `json:"mtu-ignore"`
	Priority           int                                                 `json:"priority" dval:"1"`
	RetransmitInterval int                                                 `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                 `json:"transmit-delay" dval:"1"`
	Uuid               string                                              `json:"uuid"`
}

type InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceLoopbackIpOspfOspfGlobalBfdCfg struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg struct {
	MessageDigestKey int                                                  `json:"message-digest-key"`
	Md5              InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5 `json:"md5"`
}

type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLoopbackIpOspfOspfIpList struct {
	IpAddr             string                                              `json:"ip-addr"`
	Authentication     int                                                 `json:"authentication"`
	Value              string                                              `json:"value"`
	AuthenticationKey  string                                              `json:"authentication-key"`
	Cost               int                                                 `json:"cost"`
	DatabaseFilter     string                                              `json:"database-filter"`
	Out                int                                                 `json:"out"`
	DeadInterval       int                                                 `json:"dead-interval" dval:"40"`
	HelloInterval      int                                                 `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
	MtuIgnore          int                                                 `json:"mtu-ignore"`
	Priority           int                                                 `json:"priority" dval:"1"`
	RetransmitInterval int                                                 `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                 `json:"transmit-delay" dval:"1"`
	Uuid               string                                              `json:"uuid"`
}

type InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceLoopbackIpv6 struct {
	AddressList []InterfaceLoopbackIpv6AddressList `json:"address-list"`
	Ipv6Enable  int                                `json:"ipv6-enable"`
	Uuid        string                             `json:"uuid"`
	Router      InterfaceLoopbackIpv6Router        `json:"router"`
	Rip         InterfaceLoopbackIpv6Rip           `json:"rip"`
	Ospf        InterfaceLoopbackIpv6Ospf          `json:"ospf"`
}

type InterfaceLoopbackIpv6AddressList struct {
	Ipv6Addr  string `json:"ipv6-addr"`
	Anycast   int    `json:"anycast"`
	LinkLocal int    `json:"link-local"`
}

type InterfaceLoopbackIpv6Router struct {
	Ripng InterfaceLoopbackIpv6RouterRipng `json:"ripng"`
	Ospf  InterfaceLoopbackIpv6RouterOspf  `json:"ospf"`
	Isis  InterfaceLoopbackIpv6RouterIsis  `json:"isis"`
}

type InterfaceLoopbackIpv6RouterRipng struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpv6RouterOspf struct {
	AreaList []InterfaceLoopbackIpv6RouterOspfAreaList `json:"area-list"`
	Uuid     string                                    `json:"uuid"`
}

type InterfaceLoopbackIpv6RouterOspfAreaList struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceLoopbackIpv6RouterIsis struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpv6Rip struct {
	SplitHorizonCfg InterfaceLoopbackIpv6RipSplitHorizonCfg `json:"split-horizon-cfg"`
	Uuid            string                                  `json:"uuid"`
}

type InterfaceLoopbackIpv6RipSplitHorizonCfg struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLoopbackIpv6Ospf struct {
	Bfd                   int                                              `json:"bfd"`
	Disable               int                                              `json:"disable"`
	CostCfg               []InterfaceLoopbackIpv6OspfCostCfg               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceLoopbackIpv6OspfDeadIntervalCfg       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceLoopbackIpv6OspfHelloIntervalCfg      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceLoopbackIpv6OspfMtuIgnoreCfg          `json:"mtu-ignore-cfg"`
	PriorityCfg           []InterfaceLoopbackIpv6OspfPriorityCfg           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceLoopbackIpv6OspfRetransmitIntervalCfg `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceLoopbackIpv6OspfTransmitDelayCfg      `json:"transmit-delay-cfg"`
	Uuid                  string                                           `json:"uuid"`
}

type InterfaceLoopbackIpv6OspfCostCfg struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfDeadIntervalCfg struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfHelloIntervalCfg struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfMtuIgnoreCfg struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfPriorityCfg struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfRetransmitIntervalCfg struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfTransmitDelayCfg struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLoopbackIsis struct {
	Authentication           InterfaceLoopbackIsisAuthentication             `json:"authentication"`
	BfdCfg                   InterfaceLoopbackIsisBfdCfg                     `json:"bfd-cfg"`
	CircuitType              string                                          `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceLoopbackIsisCsnpIntervalList         `json:"csnp-interval-list"`
	Padding                  int                                             `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceLoopbackIsisHelloIntervalList        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceLoopbackIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceLoopbackIsisHelloMultiplierList      `json:"hello-multiplier-list"`
	LspInterval              int                                             `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceLoopbackIsisMeshGroup                  `json:"mesh-group"`
	MetricList               []InterfaceLoopbackIsisMetricList               `json:"metric-list"`
	PasswordList             []InterfaceLoopbackIsisPasswordList             `json:"password-list"`
	PriorityList             []InterfaceLoopbackIsisPriorityList             `json:"priority-list"`
	RetransmitInterval       int                                             `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceLoopbackIsisWideMetricList           `json:"wide-metric-list"`
	Uuid                     string                                          `json:"uuid"`
}

type InterfaceLoopbackIsisAuthentication struct {
	SendOnlyList []InterfaceLoopbackIsisAuthenticationSendOnlyList `json:"send-only-list"`
	ModeList     []InterfaceLoopbackIsisAuthenticationModeList     `json:"mode-list"`
	KeyChainList []InterfaceLoopbackIsisAuthenticationKeyChainList `json:"key-chain-list"`
}

type InterfaceLoopbackIsisAuthenticationSendOnlyList struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisAuthenticationModeList struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceLoopbackIsisAuthenticationKeyChainList struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisBfdCfg struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLoopbackIsisCsnpIntervalList struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceLoopbackIsisHelloIntervalList struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceLoopbackIsisHelloIntervalMinimalList struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceLoopbackIsisHelloMultiplierList struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceLoopbackIsisMeshGroup struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceLoopbackIsisMetricList struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceLoopbackIsisPasswordList struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisPriorityList struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisWideMetricList struct {
	WideMetric int    `json:"wide-metric" dval:"10"`
	Level      string `json:"level"`
}

type InterfaceLoopbackSnmpServer struct {
	TrapSource int `json:"trap-source"`
}

func (p *InterfaceLoopback) GetId() string {
	return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceLoopback) getPath() string {
	return "interface/loopback"
}

func (p *InterfaceLoopback) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLoopback::Post")
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

func (p *InterfaceLoopback) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLoopback::Get")
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

func (p *InterfaceLoopback) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLoopback::Put")
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

func (p *InterfaceLoopback) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLoopback::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
