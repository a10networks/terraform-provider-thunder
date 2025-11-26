package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type FwTemplateLogging struct {
	Inst struct {
		Custom FwTemplateLoggingCustom `json:"custom"`

		DisableLogByDestination FwTemplateLoggingDisableLogByDestination450 `json:"disable-log-by-destination"`

		EnableLogByDestination FwTemplateLoggingEnableLogByDestination459 `json:"enable-log-by-destination"`

		Facility string `json:"facility" dval:"local0"`

		Format string `json:"format" dval:"cef"`

		IncludeDestFqdn int `json:"include-dest-fqdn"`

		IncludeHttp FwTemplateLoggingIncludeHttp `json:"include-http"`

		IncludeRadiusAttribute FwTemplateLoggingIncludeRadiusAttribute `json:"include-radius-attribute"`

		IncludeYear int `json:"include-year"`

		Log FwTemplateLoggingLog `json:"log"`

		MergedStyle int `json:"merged-style"`

		Name string `json:"name"`

		Resolution string `json:"resolution" dval:"seconds"`

		Rule FwTemplateLoggingRule `json:"rule"`

		ServiceGroup string `json:"service-group"`

		SessionPeriodicLog FwTemplateLoggingSessionPeriodicLog468 `json:"session-periodic-log"`

		Severity string `json:"severity" dval:"informational"`

		SourceAddress FwTemplateLoggingSourceAddress469 `json:"source-address"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"logging"`
}

type FwTemplateLoggingCustom struct {
	CustomMessage FwTemplateLoggingCustomCustomMessage `json:"custom-message"`
}

type FwTemplateLoggingCustomCustomMessage struct {
	CustomSessionCreated string `json:"custom-session-created"`
	CustomSessionDeleted string `json:"custom-session-deleted"`
}

type FwTemplateLoggingDisableLogByDestination450 struct {
	TcpList []FwTemplateLoggingDisableLogByDestinationTcpList451 `json:"tcp-list"`
	UdpList []FwTemplateLoggingDisableLogByDestinationUdpList452 `json:"udp-list"`
	Icmp    int                                                  `json:"icmp"`
	Others  int                                                  `json:"others"`
	Uuid    string                                               `json:"uuid"`
	IpList  []FwTemplateLoggingDisableLogByDestinationIpList453  `json:"ip-list"`
	Ip6List []FwTemplateLoggingDisableLogByDestinationIp6List456 `json:"ip6-list"`
}

type FwTemplateLoggingDisableLogByDestinationTcpList451 struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationUdpList452 struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationIpList453 struct {
	Ipv4Addr string                                                     `json:"ipv4-addr"`
	TcpList  []FwTemplateLoggingDisableLogByDestinationIpListTcpList454 `json:"tcp-list"`
	UdpList  []FwTemplateLoggingDisableLogByDestinationIpListUdpList455 `json:"udp-list"`
	Icmp     int                                                        `json:"icmp"`
	Others   int                                                        `json:"others"`
	Uuid     string                                                     `json:"uuid"`
	UserTag  string                                                     `json:"user-tag"`
}

type FwTemplateLoggingDisableLogByDestinationIpListTcpList454 struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationIpListUdpList455 struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationIp6List456 struct {
	Ipv6Addr string                                                      `json:"ipv6-addr"`
	TcpList  []FwTemplateLoggingDisableLogByDestinationIp6ListTcpList457 `json:"tcp-list"`
	UdpList  []FwTemplateLoggingDisableLogByDestinationIp6ListUdpList458 `json:"udp-list"`
	Icmp     int                                                         `json:"icmp"`
	Others   int                                                         `json:"others"`
	Uuid     string                                                      `json:"uuid"`
	UserTag  string                                                      `json:"user-tag"`
}

type FwTemplateLoggingDisableLogByDestinationIp6ListTcpList457 struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingDisableLogByDestinationIp6ListUdpList458 struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingEnableLogByDestination459 struct {
	TcpList []FwTemplateLoggingEnableLogByDestinationTcpList460 `json:"tcp-list"`
	UdpList []FwTemplateLoggingEnableLogByDestinationUdpList461 `json:"udp-list"`
	Icmp    int                                                 `json:"icmp"`
	Others  int                                                 `json:"others"`
	Uuid    string                                              `json:"uuid"`
	IpList  []FwTemplateLoggingEnableLogByDestinationIpList462  `json:"ip-list"`
	Ip6List []FwTemplateLoggingEnableLogByDestinationIp6List465 `json:"ip6-list"`
}

type FwTemplateLoggingEnableLogByDestinationTcpList460 struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationUdpList461 struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationIpList462 struct {
	Ipv4Addr string                                                    `json:"ipv4-addr"`
	TcpList  []FwTemplateLoggingEnableLogByDestinationIpListTcpList463 `json:"tcp-list"`
	UdpList  []FwTemplateLoggingEnableLogByDestinationIpListUdpList464 `json:"udp-list"`
	Icmp     int                                                       `json:"icmp"`
	Others   int                                                       `json:"others"`
	Uuid     string                                                    `json:"uuid"`
	UserTag  string                                                    `json:"user-tag"`
}

type FwTemplateLoggingEnableLogByDestinationIpListTcpList463 struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationIpListUdpList464 struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationIp6List465 struct {
	Ipv6Addr string                                                     `json:"ipv6-addr"`
	TcpList  []FwTemplateLoggingEnableLogByDestinationIp6ListTcpList466 `json:"tcp-list"`
	UdpList  []FwTemplateLoggingEnableLogByDestinationIp6ListUdpList467 `json:"udp-list"`
	Icmp     int                                                        `json:"icmp"`
	Others   int                                                        `json:"others"`
	Uuid     string                                                     `json:"uuid"`
	UserTag  string                                                     `json:"user-tag"`
}

type FwTemplateLoggingEnableLogByDestinationIp6ListTcpList466 struct {
	TcpPortStart int `json:"tcp-port-start"`
	TcpPortEnd   int `json:"tcp-port-end"`
}

type FwTemplateLoggingEnableLogByDestinationIp6ListUdpList467 struct {
	UdpPortStart int `json:"udp-port-start"`
	UdpPortEnd   int `json:"udp-port-end"`
}

type FwTemplateLoggingIncludeHttp struct {
	HeaderCfg     []FwTemplateLoggingIncludeHttpHeaderCfg `json:"header-cfg"`
	L4SessionInfo int                                     `json:"l4-session-info"`
	Method        int                                     `json:"method"`
	RequestNumber int                                     `json:"request-number"`
	FileExtension int                                     `json:"file-extension"`
}

type FwTemplateLoggingIncludeHttpHeaderCfg struct {
	HttpHeader       string `json:"http-header"`
	MaxLength        int    `json:"max-length" dval:"100"`
	CustomHeaderName string `json:"custom-header-name"`
	CustomMaxLength  int    `json:"custom-max-length" dval:"100"`
}

type FwTemplateLoggingIncludeRadiusAttribute struct {
	AttrCfg             []FwTemplateLoggingIncludeRadiusAttributeAttrCfg `json:"attr-cfg"`
	NoQuote             int                                              `json:"no-quote"`
	FramedIpv6Prefix    int                                              `json:"framed-ipv6-prefix"`
	PrefixLength        string                                           `json:"prefix-length"`
	InsertIfNotExisting int                                              `json:"insert-if-not-existing"`
	ZeroInCustomAttr    int                                              `json:"zero-in-custom-attr"`
}

type FwTemplateLoggingIncludeRadiusAttributeAttrCfg struct {
	Attr      string `json:"attr"`
	AttrEvent string `json:"attr-event"`
}

type FwTemplateLoggingLog struct {
	HttpRequests string `json:"http-requests"`
}

type FwTemplateLoggingRule struct {
	RuleHttpRequests FwTemplateLoggingRuleRuleHttpRequests `json:"rule-http-requests"`
}

type FwTemplateLoggingRuleRuleHttpRequests struct {
	DestPort             []FwTemplateLoggingRuleRuleHttpRequestsDestPort `json:"dest-port"`
	LogEveryHttpRequest  int                                             `json:"log-every-http-request"`
	MaxUrlLen            int                                             `json:"max-url-len" dval:"100"`
	IncludeAllHeaders    int                                             `json:"include-all-headers"`
	DisableSequenceCheck int                                             `json:"disable-sequence-check"`
}

type FwTemplateLoggingRuleRuleHttpRequestsDestPort struct {
	DestPortNumber   int `json:"dest-port-number"`
	IncludeByteCount int `json:"include-byte-count"`
}

type FwTemplateLoggingSessionPeriodicLog468 struct {
	Interval int    `json:"interval"`
	Uuid     string `json:"uuid"`
}

type FwTemplateLoggingSourceAddress469 struct {
	Ip   string `json:"ip"`
	Ipv6 string `json:"ipv6"`
	Uuid string `json:"uuid"`
}

func (p *FwTemplateLogging) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *FwTemplateLogging) getPath() string {
	return "fw/template/logging"
}

func (p *FwTemplateLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLogging::Post")
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

func (p *FwTemplateLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLogging::Get")
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
func (p *FwTemplateLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLogging::Put")
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

func (p *FwTemplateLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwTemplateLogging::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
