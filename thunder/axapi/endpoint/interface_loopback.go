package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type InterfaceLoopback struct {
	Inst struct {
		Ifnum int `json:"ifnum"`

		Ip InterfaceLoopbackIp765 `json:"ip"`

		Ipv6 InterfaceLoopbackIpv6786 `json:"ipv6"`

		Isis InterfaceLoopbackIsis803 `json:"isis"`

		Name string `json:"name"`

		SnmpServer InterfaceLoopbackSnmpServer `json:"snmp-server"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"loopback"`
}

type InterfaceLoopbackIp765 struct {
	AddressList []InterfaceLoopbackIpAddressList766 `json:"address-list"`
	Uuid        string                              `json:"uuid"`
	Router      InterfaceLoopbackIpRouter767        `json:"router"`
	Rip         InterfaceLoopbackIpRip769           `json:"rip"`
	Ospf        InterfaceLoopbackIpOspf777          `json:"ospf"`
}

type InterfaceLoopbackIpAddressList766 struct {
	Ipv4Address string `json:"ipv4-address"`
	Ipv4Netmask string `json:"ipv4-netmask"`
}

type InterfaceLoopbackIpRouter767 struct {
	Isis InterfaceLoopbackIpRouterIsis768 `json:"isis"`
}

type InterfaceLoopbackIpRouterIsis768 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpRip769 struct {
	Authentication  InterfaceLoopbackIpRipAuthentication770  `json:"authentication"`
	SendPacket      int                                      `json:"send-packet" dval:"1"`
	ReceivePacket   int                                      `json:"receive-packet" dval:"1"`
	SendCfg         InterfaceLoopbackIpRipSendCfg774         `json:"send-cfg"`
	ReceiveCfg      InterfaceLoopbackIpRipReceiveCfg775      `json:"receive-cfg"`
	SplitHorizonCfg InterfaceLoopbackIpRipSplitHorizonCfg776 `json:"split-horizon-cfg"`
	Uuid            string                                   `json:"uuid"`
}

type InterfaceLoopbackIpRipAuthentication770 struct {
	Str      InterfaceLoopbackIpRipAuthenticationStr771      `json:"str"`
	Mode     InterfaceLoopbackIpRipAuthenticationMode772     `json:"mode"`
	KeyChain InterfaceLoopbackIpRipAuthenticationKeyChain773 `json:"key-chain"`
}

type InterfaceLoopbackIpRipAuthenticationStr771 struct {
	String string `json:"string"`
}

type InterfaceLoopbackIpRipAuthenticationMode772 struct {
	Mode string `json:"mode" dval:"text"`
}

type InterfaceLoopbackIpRipAuthenticationKeyChain773 struct {
	KeyChain string `json:"key-chain"`
}

type InterfaceLoopbackIpRipSendCfg774 struct {
	Send    int    `json:"send"`
	Version string `json:"version"`
}

type InterfaceLoopbackIpRipReceiveCfg775 struct {
	Receive int    `json:"receive"`
	Version string `json:"version"`
}

type InterfaceLoopbackIpRipSplitHorizonCfg776 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLoopbackIpOspf777 struct {
	OspfGlobal InterfaceLoopbackIpOspfOspfGlobal778   `json:"ospf-global"`
	OspfIpList []InterfaceLoopbackIpOspfOspfIpList784 `json:"ospf-ip-list"`
}

type InterfaceLoopbackIpOspfOspfGlobal778 struct {
	AuthenticationCfg  InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg779  `json:"authentication-cfg"`
	AuthenticationKey  string                                                 `json:"authentication-key"`
	BfdCfg             InterfaceLoopbackIpOspfOspfGlobalBfdCfg780             `json:"bfd-cfg"`
	Cost               int                                                    `json:"cost"`
	DatabaseFilterCfg  InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg781  `json:"database-filter-cfg"`
	DeadInterval       int                                                    `json:"dead-interval" dval:"40"`
	Disable            string                                                 `json:"disable"`
	HelloInterval      int                                                    `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg782 `json:"message-digest-cfg"`
	Mtu                int                                                    `json:"mtu"`
	MtuIgnore          int                                                    `json:"mtu-ignore"`
	Priority           int                                                    `json:"priority" dval:"1"`
	RetransmitInterval int                                                    `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                    `json:"transmit-delay" dval:"1"`
	Uuid               string                                                 `json:"uuid"`
}

type InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg779 struct {
	Authentication int    `json:"authentication"`
	Value          string `json:"value"`
}

type InterfaceLoopbackIpOspfOspfGlobalBfdCfg780 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg781 struct {
	DatabaseFilter string `json:"database-filter"`
	Out            int    `json:"out"`
}

type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg782 struct {
	MessageDigestKey int                                                     `json:"message-digest-key"`
	Md5              InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5783 `json:"md5"`
}

type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5783 struct {
	Md5Value  string `json:"md5-value"`
	Encrypted string `json:"encrypted"`
}

type InterfaceLoopbackIpOspfOspfIpList784 struct {
	IpAddr             string                                                 `json:"ip-addr"`
	Authentication     int                                                    `json:"authentication"`
	Value              string                                                 `json:"value"`
	AuthenticationKey  string                                                 `json:"authentication-key"`
	Cost               int                                                    `json:"cost"`
	DatabaseFilter     string                                                 `json:"database-filter"`
	Out                int                                                    `json:"out"`
	DeadInterval       int                                                    `json:"dead-interval" dval:"40"`
	HelloInterval      int                                                    `json:"hello-interval" dval:"10"`
	MessageDigestCfg   []InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg785 `json:"message-digest-cfg"`
	MtuIgnore          int                                                    `json:"mtu-ignore"`
	Priority           int                                                    `json:"priority" dval:"1"`
	RetransmitInterval int                                                    `json:"retransmit-interval" dval:"5"`
	TransmitDelay      int                                                    `json:"transmit-delay" dval:"1"`
	Uuid               string                                                 `json:"uuid"`
}

type InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg785 struct {
	MessageDigestKey int    `json:"message-digest-key"`
	Md5Value         string `json:"md5-value"`
	Encrypted        string `json:"encrypted"`
}

type InterfaceLoopbackIpv6786 struct {
	AddressList []InterfaceLoopbackIpv6AddressList787 `json:"address-list"`
	Ipv6Enable  int                                   `json:"ipv6-enable"`
	Uuid        string                                `json:"uuid"`
	Router      InterfaceLoopbackIpv6Router788        `json:"router"`
	Rip         InterfaceLoopbackIpv6Rip793           `json:"rip"`
	Ospf        InterfaceLoopbackIpv6Ospf795          `json:"ospf"`
}

type InterfaceLoopbackIpv6AddressList787 struct {
	Ipv6Addr  string `json:"ipv6-addr"`
	Anycast   int    `json:"anycast"`
	LinkLocal int    `json:"link-local"`
}

type InterfaceLoopbackIpv6Router788 struct {
	Ripng InterfaceLoopbackIpv6RouterRipng789 `json:"ripng"`
	Ospf  InterfaceLoopbackIpv6RouterOspf790  `json:"ospf"`
	Isis  InterfaceLoopbackIpv6RouterIsis792  `json:"isis"`
}

type InterfaceLoopbackIpv6RouterRipng789 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpv6RouterOspf790 struct {
	AreaList []InterfaceLoopbackIpv6RouterOspfAreaList791 `json:"area-list"`
	Uuid     string                                       `json:"uuid"`
}

type InterfaceLoopbackIpv6RouterOspfAreaList791 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceLoopbackIpv6RouterIsis792 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpv6Rip793 struct {
	SplitHorizonCfg InterfaceLoopbackIpv6RipSplitHorizonCfg794 `json:"split-horizon-cfg"`
	Uuid            string                                     `json:"uuid"`
}

type InterfaceLoopbackIpv6RipSplitHorizonCfg794 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLoopbackIpv6Ospf795 struct {
	Bfd                   int                                                 `json:"bfd"`
	Disable               int                                                 `json:"disable"`
	CostCfg               []InterfaceLoopbackIpv6OspfCostCfg796               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceLoopbackIpv6OspfDeadIntervalCfg797       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceLoopbackIpv6OspfHelloIntervalCfg798      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceLoopbackIpv6OspfMtuIgnoreCfg799          `json:"mtu-ignore-cfg"`
	PriorityCfg           []InterfaceLoopbackIpv6OspfPriorityCfg800           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceLoopbackIpv6OspfRetransmitIntervalCfg801 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceLoopbackIpv6OspfTransmitDelayCfg802      `json:"transmit-delay-cfg"`
	Uuid                  string                                              `json:"uuid"`
}

type InterfaceLoopbackIpv6OspfCostCfg796 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfDeadIntervalCfg797 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfHelloIntervalCfg798 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfMtuIgnoreCfg799 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfPriorityCfg800 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfRetransmitIntervalCfg801 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfTransmitDelayCfg802 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLoopbackIsis803 struct {
	Authentication           InterfaceLoopbackIsisAuthentication804             `json:"authentication"`
	BfdCfg                   InterfaceLoopbackIsisBfdCfg808                     `json:"bfd-cfg"`
	CircuitType              string                                             `json:"circuit-type" dval:"level-1-2"`
	CsnpIntervalList         []InterfaceLoopbackIsisCsnpIntervalList809         `json:"csnp-interval-list"`
	Padding                  int                                                `json:"padding" dval:"1"`
	HelloIntervalList        []InterfaceLoopbackIsisHelloIntervalList810        `json:"hello-interval-list"`
	HelloIntervalMinimalList []InterfaceLoopbackIsisHelloIntervalMinimalList811 `json:"hello-interval-minimal-list"`
	HelloMultiplierList      []InterfaceLoopbackIsisHelloMultiplierList812      `json:"hello-multiplier-list"`
	LspInterval              int                                                `json:"lsp-interval" dval:"33"`
	MeshGroup                InterfaceLoopbackIsisMeshGroup813                  `json:"mesh-group"`
	MetricList               []InterfaceLoopbackIsisMetricList814               `json:"metric-list"`
	PasswordList             []InterfaceLoopbackIsisPasswordList815             `json:"password-list"`
	PriorityList             []InterfaceLoopbackIsisPriorityList816             `json:"priority-list"`
	RetransmitInterval       int                                                `json:"retransmit-interval" dval:"5"`
	WideMetricList           []InterfaceLoopbackIsisWideMetricList817           `json:"wide-metric-list"`
	Uuid                     string                                             `json:"uuid"`
}

type InterfaceLoopbackIsisAuthentication804 struct {
	SendOnlyList []InterfaceLoopbackIsisAuthenticationSendOnlyList805 `json:"send-only-list"`
	ModeList     []InterfaceLoopbackIsisAuthenticationModeList806     `json:"mode-list"`
	KeyChainList []InterfaceLoopbackIsisAuthenticationKeyChainList807 `json:"key-chain-list"`
}

type InterfaceLoopbackIsisAuthenticationSendOnlyList805 struct {
	SendOnly int    `json:"send-only"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisAuthenticationModeList806 struct {
	Mode  string `json:"mode"`
	Level string `json:"level"`
}

type InterfaceLoopbackIsisAuthenticationKeyChainList807 struct {
	KeyChain string `json:"key-chain"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisBfdCfg808 struct {
	Bfd     int `json:"bfd"`
	Disable int `json:"disable"`
}

type InterfaceLoopbackIsisCsnpIntervalList809 struct {
	CsnpInterval int    `json:"csnp-interval" dval:"10"`
	Level        string `json:"level"`
}

type InterfaceLoopbackIsisHelloIntervalList810 struct {
	HelloInterval int    `json:"hello-interval" dval:"10"`
	Level         string `json:"level"`
}

type InterfaceLoopbackIsisHelloIntervalMinimalList811 struct {
	HelloIntervalMinimal int    `json:"hello-interval-minimal"`
	Level                string `json:"level"`
}

type InterfaceLoopbackIsisHelloMultiplierList812 struct {
	HelloMultiplier int    `json:"hello-multiplier" dval:"3"`
	Level           string `json:"level"`
}

type InterfaceLoopbackIsisMeshGroup813 struct {
	Value   int `json:"value"`
	Blocked int `json:"blocked"`
}

type InterfaceLoopbackIsisMetricList814 struct {
	Metric int    `json:"metric" dval:"10"`
	Level  string `json:"level"`
}

type InterfaceLoopbackIsisPasswordList815 struct {
	Password string `json:"password"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisPriorityList816 struct {
	Priority int    `json:"priority" dval:"64"`
	Level    string `json:"level"`
}

type InterfaceLoopbackIsisWideMetricList817 struct {
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
