

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateDns struct {
	Inst struct {

    Action string `json:"action" dval:"drop"`
    AllowQueryClass DdosTemplateDnsAllowQueryClass `json:"allow-query-class"`
    AllowRecordType DdosTemplateDnsAllowRecordType `json:"allow-record-type"`
    DnsAnyCheck int `json:"dns-any-check"`
    DnsAuthCfg DdosTemplateDnsDnsAuthCfg `json:"dns-auth-cfg"`
    DnsRequestRateLimit DdosTemplateDnsDnsRequestRateLimit `json:"dns-request-rate-limit"`
    DomainGroupName string `json:"domain-group-name"`
    DomainGroupRateExceedAction string `json:"domain-group-rate-exceed-action" dval:"drop"`
    DomainGroupRatePerService int `json:"domain-group-rate-per-service"`
    EncapTemplate string `json:"encap-template"`
    FqdnCfg []DdosTemplateDnsFqdnCfg `json:"fqdn-cfg"`
    FqdnLabelCount int `json:"fqdn-label-count"`
    FqdnLabelLenCfg []DdosTemplateDnsFqdnLabelLenCfg `json:"fqdn-label-len-cfg"`
    MalformedQueryCheck DdosTemplateDnsMalformedQueryCheck295 `json:"malformed-query-check"`
    MultiPuThresholdDistribution DdosTemplateDnsMultiPuThresholdDistribution `json:"multi-pu-threshold-distribution"`
    Name string `json:"name"`
    NxdomainCfg DdosTemplateDnsNxdomainCfg `json:"nxdomain-cfg"`
    OnNoMatch string `json:"on-no-match" dval:"deny"`
    QueryRateThresholdForCacheServing int `json:"query-rate-threshold-for-cache-serving"`
    SymtimeoutCfg DdosTemplateDnsSymtimeoutCfg `json:"symtimeout-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dns"`
}


type DdosTemplateDnsAllowQueryClass struct {
    AllowInternetQueryClass int `json:"allow-internet-query-class"`
    AllowCsnetQueryClass int `json:"allow-csnet-query-class"`
    AllowChaosQueryClass int `json:"allow-chaos-query-class"`
    AllowHesiodQueryClass int `json:"allow-hesiod-query-class"`
    AllowNoneQueryClass int `json:"allow-none-query-class"`
    AllowAnyQueryClass int `json:"allow-any-query-class"`
}


type DdosTemplateDnsAllowRecordType struct {
    AllowAType int `json:"allow-a-type"`
    AllowAaaaType int `json:"allow-aaaa-type"`
    AllowCnameType int `json:"allow-cname-type"`
    AllowMxType int `json:"allow-mx-type"`
    AllowNsType int `json:"allow-ns-type"`
    AllowSrvType int `json:"allow-srv-type"`
    RecordNumCfg []DdosTemplateDnsAllowRecordTypeRecordNumCfg `json:"record-num-cfg"`
}


type DdosTemplateDnsAllowRecordTypeRecordNumCfg struct {
    AllowNumType int `json:"allow-num-type"`
}


type DdosTemplateDnsDnsAuthCfg struct {
    DnsAuth int `json:"dns-auth"`
    DnsAuthType string `json:"dns-auth-type"`
    UdpTimeoutValOnly int `json:"udp-timeout-val-only"`
    UdpTimeout int `json:"udp-timeout"`
    MinRetryGap int `json:"min-retry-gap"`
    MinRetryGapInterval string `json:"min-retry-gap-interval" dval:"1sec"`
    WithUdpAuth int `json:"with-udp-auth"`
    ForceTcpTimeout int `json:"force-tcp-timeout"`
    ForceTcpMinRetryGap int `json:"force-tcp-min-retry-gap"`
    ForceTcpIgnoreClientSourcePort int `json:"force-tcp-ignore-client-source-port"`
}


type DdosTemplateDnsDnsRequestRateLimit struct {
    Type DdosTemplateDnsDnsRequestRateLimitType `json:"type"`
}


type DdosTemplateDnsDnsRequestRateLimitType struct {
    ACfg DdosTemplateDnsDnsRequestRateLimitTypeACfg `json:"A-cfg"`
    AaaaCfg DdosTemplateDnsDnsRequestRateLimitTypeAaaaCfg `json:"AAAA-cfg"`
    CnameCfg DdosTemplateDnsDnsRequestRateLimitTypeCnameCfg `json:"CNAME-cfg"`
    MxCfg DdosTemplateDnsDnsRequestRateLimitTypeMxCfg `json:"MX-cfg"`
    NsCfg DdosTemplateDnsDnsRequestRateLimitTypeNsCfg `json:"NS-cfg"`
    SrvCfg DdosTemplateDnsDnsRequestRateLimitTypeSrvCfg `json:"SRV-cfg"`
    DnsTypeCfg []DdosTemplateDnsDnsRequestRateLimitTypeDnsTypeCfg `json:"dns-type-cfg"`
}


type DdosTemplateDnsDnsRequestRateLimitTypeACfg struct {
    A int `json:"A"`
    DnsARate int `json:"dns-a-rate"`
}


type DdosTemplateDnsDnsRequestRateLimitTypeAaaaCfg struct {
    Aaaa int `json:"AAAA"`
    DnsAaaaRate int `json:"dns-aaaa-rate"`
}


type DdosTemplateDnsDnsRequestRateLimitTypeCnameCfg struct {
    Cname int `json:"CNAME"`
    DnsCnameRate int `json:"dns-cname-rate"`
}


type DdosTemplateDnsDnsRequestRateLimitTypeMxCfg struct {
    Mx int `json:"MX"`
    DnsMxRate int `json:"dns-mx-rate"`
}


type DdosTemplateDnsDnsRequestRateLimitTypeNsCfg struct {
    Ns int `json:"NS"`
    DnsNsRate int `json:"dns-ns-rate"`
}


type DdosTemplateDnsDnsRequestRateLimitTypeSrvCfg struct {
    Srv int `json:"SRV"`
    DnsSrvRate int `json:"dns-srv-rate"`
}


type DdosTemplateDnsDnsRequestRateLimitTypeDnsTypeCfg struct {
    DnsRequestType int `json:"dns-request-type"`
    DnsRequestTypeRate int `json:"dns-request-type-rate"`
}


type DdosTemplateDnsFqdnCfg struct {
    DnsFqdnRateLimit int `json:"dns-fqdn-rate-limit"`
    DnsFqdnRate int `json:"dns-fqdn-rate"`
    Per string `json:"per"`
    PerDomainPerSrcIp int `json:"per-domain-per-src-ip"`
    FqdnRateSuffix int `json:"fqdn-rate-suffix"`
    FqdnRateLabelCount int `json:"fqdn-rate-label-count"`
    By string `json:"by"`
    FqdnRateSuffixBy int `json:"fqdn-rate-suffix-by"`
}


type DdosTemplateDnsFqdnLabelLenCfg struct {
    FqdnLabelLength int `json:"fqdn-label-length"`
    LabelLength int `json:"label-length"`
    FqdnLabelSuffix int `json:"fqdn-label-suffix"`
}


type DdosTemplateDnsMalformedQueryCheck295 struct {
    ValidationType string `json:"validation-type"`
    NonQueryOpcodeCheck string `json:"non-query-opcode-check"`
    SkipMultiPacketCheck int `json:"skip-multi-packet-check"`
    Uuid string `json:"uuid"`
}


type DdosTemplateDnsMultiPuThresholdDistribution struct {
    MultiPuThresholdDistributionValue int `json:"multi-pu-threshold-distribution-value"`
    MultiPuThresholdDistributionDisable string `json:"multi-pu-threshold-distribution-disable"`
}


type DdosTemplateDnsNxdomainCfg struct {
    DnsNxdomainRateLimit int `json:"dns-nxdomain-rate-limit"`
    DnsNxdomainRate int `json:"dns-nxdomain-rate"`
    DnsNxdomainRateLimitAction string `json:"dns-nxdomain-rate-limit-action"`
}


type DdosTemplateDnsSymtimeoutCfg struct {
    SymTimeout int `json:"sym-timeout"`
    SymTimeoutValue int `json:"sym-timeout-value"`
}

func (p *DdosTemplateDns) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosTemplateDns) getPath() string{
    return "ddos/template/dns"
}

func (p *DdosTemplateDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateDns::Post")
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

func (p *DdosTemplateDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateDns::Get")
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
func (p *DdosTemplateDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateDns::Put")
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

func (p *DdosTemplateDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
