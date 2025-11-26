package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type InterfaceLoopback struct {
	Inst struct {
		Ifnum int `json:"ifnum"`

		Ip InterfaceLoopbackIp774 `json:"ip"`

		Ipv6 InterfaceLoopbackIpv6795 `json:"ipv6"`

		Isis InterfaceLoopbackIsis812 `json:"isis"`

		Name string `json:"name"`

		SnmpServer InterfaceLoopbackSnmpServer `json:"snmp-server"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"loopback"`
}

type InterfaceLoopbackIp774 struct {
	AddressList []InterfaceLoopbackIpAddressList775 `json:"address-list"`
	Uuid        string                              `json:"uuid"`
	Router      InterfaceLoopbackIpRouter776        `json:"router"`
	Rip         InterfaceLoopbackIpRip778           `json:"rip"`
	Ospf        InterfaceLoopbackIpOspf786          `json:"ospf"`
}

type InterfaceLoopbackIpAddressList775 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceLoopbackIpRouter776 struct {
	Isis InterfaceLoopbackIpRouterIsis777 `json:"isis"`
}

type InterfaceLoopbackIpRouterIsis777 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpRip778 struct {
	Authentication  InterfaceLoopbackIpRipAuthentication779  `json:"authentication"`
	SendPacket      int                                      `json:"send-packet" dval:"1"`
	ReceivePacket   int                                      `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceLoopbackIpRipSendCfg783         `json:"send-cfg"`
	ReceiveCfg      InterfaceLoopbackIpRipReceiveCfg784      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceLoopbackIpRipSplitHorizonCfg785 `json:"split-horizon-cfg"`
	Uuid            string                                   `json:"uuid"`
}

type InterfaceLoopbackIpRipAuthentication779 struct {
	Str      InterfaceLoopbackIpRipAuthenticationStr780      `json:"str"`
	Mode     InterfaceLoopbackIpRipAuthenticationMode781     `json:"mode"`
	KeyChain InterfaceLoopbackIpRipAuthenticationKeyChain782 `json:"key-chain"`
}

type InterfaceLoopbackIpRipAuthenticationStr780 struct {
	String string `json:"string"`
}

type InterfaceLoopbackIpRipAuthenticationMode781 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceLoopbackIpRipAuthenticationKeyChain782 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceLoopbackIpRipSendCfg783 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceLoopbackIpRipReceiveCfg784 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceLoopbackIpRipSplitHorizonCfg785 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLoopbackIpOspf786 struct {
	OspfGlobal InterfaceLoopbackIpOspfOspfGlobal787   `json:"ospf-global"`
	OspfIpList []InterfaceLoopbackIpOspfOspfIpList793 `json:"ospf-ip-list"`
}

type InterfaceLoopbackIpOspfOspfGlobal787 struct {
	AuthenticationCfg  InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg788  `json:"authentication-cfg"`
	AuthenticationKey  string                                                 `json:"authentication-key"`
	BfdCfg             InterfaceLoopbackIpOspfOspfGlobalBfdCfg789             `json:"bfd-cfg"`
	Cost               int                                                    `json:"cost"`
	DatabaseFilterCfg  InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg790  `json:"database-filter-cfg"`
	DeadInterval       int                                                    `json:"dead-interval" dval:"40"`
	Disable            string                                                 `json:"disable"`
	HelloInterval      int                                                    `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg791 `json:"message-digest-cfg"`
	Mtu                int                                                    `json:"mtu"`
	MtuIgnore          int                                                    `json:"mtu-ignore"`
	Priority           int                                                    `json:"priority" dval:"1"`
	RetransmitInterval int                                                    `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                    `json:"transmit-delay" dval:"1"`
	Uuid               string                                                 `json:"uuid"`
}

type InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg788 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceLoopbackIpOspfOspfGlobalBfdCfg789 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg790 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg791 struct {
	MessageDigestKey int                                                     `json:"message-digest-key"`
	Md5              InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5792 `json:"md5"`
}

type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5792 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLoopbackIpOspfOspfIpList793 struct {
	IpAddr             string                                                 `json:"ip-addr"`
	Authentication     int                                                    `json:"authentication"`
	Value              string                                                 `json:"value"`
	AuthenticationKey  string                                                 `json:"authentication-key"`
	Cost               int                                                    `json:"cost"`
	DatabaseFilter     string                                                 `json:"database-filter"`
	Out                int                                                    `json:"out"`
	DeadInterval       int                                                    `json:"dead-interval" dval:"40"`
	HelloInterval      int                                                    `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg794 `json:"message-digest-cfg"`
	MtuIgnore          int                                                    `json:"mtu-ignore"`
	Priority           int                                                    `json:"priority" dval:"1"`
	RetransmitInterval int                                                    `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                    `json:"transmit-delay" dval:"1"`
	Uuid               string                                                 `json:"uuid"`
}

type InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg794 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceLoopbackIpv6795 struct {
	AddressList []InterfaceLoopbackIpv6AddressList796 `json:"address-list"`
	Ipv6Enable  int                                   `json:"ipv6-enable"`
	Uuid        string                                `json:"uuid"`
	Router      InterfaceLoopbackIpv6Router797        `json:"router"`
	Rip         InterfaceLoopbackIpv6Rip802           `json:"rip"`
	Ospf        InterfaceLoopbackIpv6Ospf804          `json:"ospf"`
}

type InterfaceLoopbackIpv6AddressList796 struct {
	Ipv6Addr  string `json:"ipv6-addr"`
	Anycast   int    `json:"anycast"`
	LinkLocal int    `json:"link-local"`
}

type InterfaceLoopbackIpv6Router797 struct {
	Ripng InterfaceLoopbackIpv6RouterRipng798 `json:"ripng"`
	Ospf  InterfaceLoopbackIpv6RouterOspf799  `json:"ospf"`
	Isis  InterfaceLoopbackIpv6RouterIsis801  `json:"isis"`
}

type InterfaceLoopbackIpv6RouterRipng798 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpv6RouterOspf799 struct {
	AreaList []InterfaceLoopbackIpv6RouterOspfAreaList800 `json:"area-list"`
	Uuid     string                                       `json:"uuid"`
}

type InterfaceLoopbackIpv6RouterOspfAreaList800 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceLoopbackIpv6RouterIsis801 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpv6Rip802 struct {
	SplitHorizonCfg InterfaceLoopbackIpv6RipSplitHorizonCfg803 `json:"split-horizon-cfg"`
	Uuid            string                                     `json:"uuid"`
}

type InterfaceLoopbackIpv6RipSplitHorizonCfg803 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLoopbackIpv6Ospf804 struct {
	Bfd                   int                                                 `json:"bfd"`
	Disable               int                                                 `json:"disable"`
	CostCfg               []InterfaceLoopbackIpv6OspfCostCfg805               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceLoopbackIpv6OspfDeadIntervalCfg806       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceLoopbackIpv6OspfHelloIntervalCfg807      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceLoopbackIpv6OspfMtuIgnoreCfg808          `json:"mtu-ignore-cfg"`
	PriorityCfg           []InterfaceLoopbackIpv6OspfPriorityCfg809           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceLoopbackIpv6OspfRetransmitIntervalCfg810 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceLoopbackIpv6OspfTransmitDelayCfg811      `json:"transmit-delay-cfg"`
	Uuid                  string                                              `json:"uuid"`
}

type InterfaceLoopbackIpv6OspfCostCfg805 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfDeadIntervalCfg806 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfHelloIntervalCfg807 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfMtuIgnoreCfg808 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfPriorityCfg809 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfRetransmitIntervalCfg810 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfTransmitDelayCfg811 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLoopbackIsis812 struct {
	Authentication           InterfaceLoopbackIsisAuthentication813             `json:"authentication"`
	BfdCfg                   InterfaceLoopbackIsisBfdCfg817                     `json:"bfd-cfg"`
	CircuitType              string                                             `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceLoopbackIsisCsnpIntervalList818         `json:"csnp-interval-list"`
	Padding                  int                                                `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceLoopbackIsisHelloIntervalList819        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceLoopbackIsisHelloIntervalMinimalList820 `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceLoopbackIsisHelloMultiplierList821      `json:"hello-multiplier-list"`
	LspInterval              int                                                `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceLoopbackIsisMeshGroup822                  `json:"mesh-group"`
	MetricList               []InterfaceLoopbackIsisMetricList823               `json:"metric-list"`
	PasswordList             []InterfaceLoopbackIsisPasswordList824             `json:"password-list"`
	PriorityList             []InterfaceLoopbackIsisPriorityList825             `json:"priority-list"`
	RetransmitInterval       int                                                `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceLoopbackIsisWideMetricList826           `json:"wide-metric-list"`
	Uuid                     string                                             `json:"uuid"`
}

type InterfaceLoopbackIsisAuthentication813 struct {
	SendOnlyList []InterfaceLoopbackIsisAuthenticationSendOnlyList814 `json:"send-only-list"`
	ModeList     []InterfaceLoopbackIsisAuthenticationModeList815     `json:"mode-list"`
	KeyChainList []InterfaceLoopbackIsisAuthenticationKeyChainList816 `json:"key-chain-list"`
}

type InterfaceLoopbackIsisAuthenticationSendOnlyList814 struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisAuthenticationModeList815 struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceLoopbackIsisAuthenticationKeyChainList816 struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisBfdCfg817 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLoopbackIsisCsnpIntervalList818 struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceLoopbackIsisHelloIntervalList819 struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceLoopbackIsisHelloIntervalMinimalList820 struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceLoopbackIsisHelloMultiplierList821 struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceLoopbackIsisMeshGroup822 struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceLoopbackIsisMetricList823 struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceLoopbackIsisPasswordList824 struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisPriorityList825 struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisWideMetricList826 struct {
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
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
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
