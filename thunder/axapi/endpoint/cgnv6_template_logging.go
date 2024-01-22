

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplateLogging struct {
	Inst struct {

    BatchedLoggingDisable int `json:"batched-logging-disable"`
    Custom Cgnv6TemplateLoggingCustom `json:"custom"`
    DisableLogByDestination Cgnv6TemplateLoggingDisableLogByDestination115 `json:"disable-log-by-destination"`
    Facility string `json:"facility" dval:"local0"`
    Format string `json:"format" dval:"default"`
    IncludeDestination int `json:"include-destination"`
    IncludeHttp Cgnv6TemplateLoggingIncludeHttp `json:"include-http"`
    IncludeInsideUserMac int `json:"include-inside-user-mac"`
    IncludePartitionName int `json:"include-partition-name"`
    IncludePortBlockAccount int `json:"include-port-block-account"`
    IncludeRadiusAttribute Cgnv6TemplateLoggingIncludeRadiusAttribute `json:"include-radius-attribute"`
    IncludeSessionByteCount int `json:"include-session-byte-count"`
    Log Cgnv6TemplateLoggingLog `json:"log"`
    LogReceiver Cgnv6TemplateLoggingLogReceiver `json:"log-receiver"`
    Name string `json:"name"`
    Resolution string `json:"resolution" dval:"seconds"`
    RfcCustom Cgnv6TemplateLoggingRfcCustom `json:"rfc-custom"`
    Rule Cgnv6TemplateLoggingRule `json:"rule"`
    ServiceGroup string `json:"service-group"`
    Severity Cgnv6TemplateLoggingSeverity `json:"severity"`
    Shared int `json:"shared"`
    SourceAddress Cgnv6TemplateLoggingSourceAddress124 `json:"source-address"`
    SourcePort Cgnv6TemplateLoggingSourcePort `json:"source-port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"logging"`
}


type Cgnv6TemplateLoggingCustom struct {
    CustomHeader string `json:"custom-header"`
    CustomMessage Cgnv6TemplateLoggingCustomCustomMessage `json:"custom-message"`
    CustomTimeStampFormat string `json:"custom-time-stamp-format"`
}


type Cgnv6TemplateLoggingCustomCustomMessage struct {
    CustomDhcpv6MapPrefixAssigned string `json:"custom-dhcpv6-map-prefix-assigned"`
    CustomDhcpv6MapPrefixReleased string `json:"custom-dhcpv6-map-prefix-released"`
    CustomDhcpv6MapPrefixRenewed string `json:"custom-dhcpv6-map-prefix-renewed"`
    CustomFixedNatAllocated string `json:"custom-fixed-nat-allocated"`
    CustomFixedNatInterimUpdate string `json:"custom-fixed-nat-interim-update"`
    CustomFixedNatFreed string `json:"custom-fixed-nat-freed"`
    CustomHttpRequestGot string `json:"custom-http-request-got"`
    CustomPortAllocated string `json:"custom-port-allocated"`
    CustomPortBatchAllocated string `json:"custom-port-batch-allocated"`
    CustomPortBatchFreed string `json:"custom-port-batch-freed"`
    CustomPortBatchV2Allocated string `json:"custom-port-batch-v2-allocated"`
    CustomPortBatchV2Freed string `json:"custom-port-batch-v2-freed"`
    CustomPortBatchV2InterimUpdate string `json:"custom-port-batch-v2-interim-update"`
    CustomPortFreed string `json:"custom-port-freed"`
    CustomSessionCreated string `json:"custom-session-created"`
    CustomSessionDeleted string `json:"custom-session-deleted"`
}


type Cgnv6TemplateLoggingDisableLogByDestination115 struct {
    TcpList []Cgnv6TemplateLoggingDisableLogByDestinationTcpList116 `json:"tcp-list"`
    UdpList []Cgnv6TemplateLoggingDisableLogByDestinationUdpList117 `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    IpList []Cgnv6TemplateLoggingDisableLogByDestinationIpList118 `json:"ip-list"`
    Ip6List []Cgnv6TemplateLoggingDisableLogByDestinationIp6List121 `json:"ip6-list"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationTcpList116 struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationUdpList117 struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIpList118 struct {
    Ipv4Addr string `json:"ipv4-addr"`
    TcpList []Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList119 `json:"tcp-list"`
    UdpList []Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList120 `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList119 struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList120 struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIp6List121 struct {
    Ipv6Addr string `json:"ipv6-addr"`
    TcpList []Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList122 `json:"tcp-list"`
    UdpList []Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList123 `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList122 struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList123 struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type Cgnv6TemplateLoggingIncludeHttp struct {
    HeaderCfg []Cgnv6TemplateLoggingIncludeHttpHeaderCfg `json:"header-cfg"`
    L4SessionInfo int `json:"l4-session-info"`
    Method int `json:"method"`
    RequestNumber int `json:"request-number"`
    FileExtension int `json:"file-extension"`
}


type Cgnv6TemplateLoggingIncludeHttpHeaderCfg struct {
    HttpHeader string `json:"http-header"`
    MaxLength int `json:"max-length" dval:"100"`
    CustomHeaderName string `json:"custom-header-name"`
    CustomMaxLength int `json:"custom-max-length" dval:"100"`
}


type Cgnv6TemplateLoggingIncludeRadiusAttribute struct {
    AttrCfg []Cgnv6TemplateLoggingIncludeRadiusAttributeAttrCfg `json:"attr-cfg"`
    NoQuote int `json:"no-quote"`
    InsertIfNotExisting int `json:"insert-if-not-existing"`
    ZeroInCustomAttr int `json:"zero-in-custom-attr"`
    FramedIpv6Prefix int `json:"framed-ipv6-prefix"`
    PrefixLength string `json:"prefix-length"`
}


type Cgnv6TemplateLoggingIncludeRadiusAttributeAttrCfg struct {
    Attr string `json:"attr"`
    AttrEvent string `json:"attr-event"`
}


type Cgnv6TemplateLoggingLog struct {
    FixedNat Cgnv6TemplateLoggingLogFixedNat `json:"fixed-nat"`
    OneToOneNat Cgnv6TemplateLoggingLogOneToOneNat `json:"one-to-one-nat"`
    MapDhcpv6 Cgnv6TemplateLoggingLogMapDhcpv6 `json:"map-dhcpv6"`
    HttpRequests string `json:"http-requests"`
    PortMappings string `json:"port-mappings" dval:"both"`
    PortOverloading int `json:"port-overloading"`
    UserData int `json:"user-data"`
    Sessions int `json:"sessions"`
    MergedStyle int `json:"merged-style"`
}


type Cgnv6TemplateLoggingLogFixedNat struct {
    FixedNatHttpRequests string `json:"fixed-nat-http-requests"`
    FixedNatPortMappings string `json:"fixed-nat-port-mappings"`
    FixedNatSessions int `json:"fixed-nat-sessions"`
    FixedNatMergedStyle int `json:"fixed-nat-merged-style"`
    UserPorts Cgnv6TemplateLoggingLogFixedNatUserPorts `json:"user-ports"`
}


type Cgnv6TemplateLoggingLogFixedNatUserPorts struct {
    UserPorts int `json:"user-ports"`
    Days int `json:"days"`
    StartTime string `json:"start-time"`
}


type Cgnv6TemplateLoggingLogOneToOneNat struct {
    OneToOneNatSessions int `json:"one-to-one-nat-sessions"`
    OneToOneMergedStyle int `json:"one-to-one-merged-style"`
}


type Cgnv6TemplateLoggingLogMapDhcpv6 struct {
    MapDhcpv6PrefixAll int `json:"map-dhcpv6-prefix-all"`
    MapDhcpv6MsgType []Cgnv6TemplateLoggingLogMapDhcpv6MapDhcpv6MsgType `json:"map-dhcpv6-msg-type"`
}


type Cgnv6TemplateLoggingLogMapDhcpv6MapDhcpv6MsgType struct {
    MapDhcpv6MsgType string `json:"map-dhcpv6-msg-type"`
}


type Cgnv6TemplateLoggingLogReceiver struct {
    Radius int `json:"radius"`
    SecretString string `json:"secret-string"`
    Encrypted string `json:"encrypted"`
}


type Cgnv6TemplateLoggingRfcCustom struct {
    Header Cgnv6TemplateLoggingRfcCustomHeader `json:"header"`
    Message Cgnv6TemplateLoggingRfcCustomMessage `json:"message"`
}


type Cgnv6TemplateLoggingRfcCustomHeader struct {
    UseAlternateTimestamp int `json:"use-alternate-timestamp"`
}


type Cgnv6TemplateLoggingRfcCustomMessage struct {
    Ipv6Tech []Cgnv6TemplateLoggingRfcCustomMessageIpv6Tech `json:"ipv6-tech"`
    Dhcpv6MapPrefixAssigned string `json:"dhcpv6-map-prefix-assigned"`
    Dhcpv6MapPrefixReleased string `json:"dhcpv6-map-prefix-released"`
    Dhcpv6MapPrefixRenewed string `json:"dhcpv6-map-prefix-renewed"`
    HttpRequestGot string `json:"http-request-got"`
    SessionCreated string `json:"session-created"`
    SessionDeleted string `json:"session-deleted"`
}


type Cgnv6TemplateLoggingRfcCustomMessageIpv6Tech struct {
    TechType string `json:"tech-type"`
    FixedNatAllocated string `json:"fixed-nat-allocated"`
    FixedNatFreed string `json:"fixed-nat-freed"`
    PortAllocated string `json:"port-allocated"`
    PortFreed string `json:"port-freed"`
    PortBatchAllocated string `json:"port-batch-allocated"`
    PortBatchFreed string `json:"port-batch-freed"`
    PortBatchV2Allocated string `json:"port-batch-v2-allocated"`
    PortBatchV2Freed string `json:"port-batch-v2-freed"`
}


type Cgnv6TemplateLoggingRule struct {
    RuleHttpRequests Cgnv6TemplateLoggingRuleRuleHttpRequests `json:"rule-http-requests"`
    InterimUpdateInterval int `json:"interim-update-interval"`
}


type Cgnv6TemplateLoggingRuleRuleHttpRequests struct {
    DestPort []Cgnv6TemplateLoggingRuleRuleHttpRequestsDestPort `json:"dest-port"`
    LogEveryHttpRequest int `json:"log-every-http-request"`
    MaxUrlLen int `json:"max-url-len" dval:"128"`
    IncludeAllHeaders int `json:"include-all-headers"`
    DisableSequenceCheck int `json:"disable-sequence-check"`
}


type Cgnv6TemplateLoggingRuleRuleHttpRequestsDestPort struct {
    DestPortNumber int `json:"dest-port-number"`
    IncludeByteCount int `json:"include-byte-count"`
}


type Cgnv6TemplateLoggingSeverity struct {
    SeverityString string `json:"severity-string" dval:"debug"`
    SeverityVal int `json:"severity-val" dval:"7"`
}


type Cgnv6TemplateLoggingSourceAddress124 struct {
    Ip string `json:"ip"`
    Ipv6 string `json:"ipv6"`
    Uuid string `json:"uuid"`
}


type Cgnv6TemplateLoggingSourcePort struct {
    SourcePortNum int `json:"source-port-num" dval:"514"`
    Any int `json:"any"`
}

func (p *Cgnv6TemplateLogging) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6TemplateLogging) getPath() string{
    return "cgnv6/template/logging"
}

func (p *Cgnv6TemplateLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLogging::Post")
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

func (p *Cgnv6TemplateLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLogging::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *Cgnv6TemplateLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLogging::Put")
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

func (p *Cgnv6TemplateLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLogging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
