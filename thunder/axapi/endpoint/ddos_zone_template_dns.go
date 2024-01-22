

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateDns struct {
	Inst struct {

    AllowQueryClass DdosZoneTemplateDnsAllowQueryClass `json:"allow-query-class"`
    AllowRecordType DdosZoneTemplateDnsAllowRecordType `json:"allow-record-type"`
    DnsAnyCheck int `json:"dns-any-check"`
    DnsAnyCheckAction string `json:"dns-any-check-action" dval:"drop"`
    DnsAnyCheckActionListName string `json:"dns-any-check-action-list-name"`
    DnsUdpAuthentication DdosZoneTemplateDnsDnsUdpAuthentication `json:"dns-udp-authentication"`
    DomainGroupName string `json:"domain-group-name"`
    Dst DdosZoneTemplateDnsDst `json:"dst"`
    FqdnLabelCountCfg DdosZoneTemplateDnsFqdnLabelCountCfg `json:"fqdn-label-count-cfg"`
    FqdnLabelLenCfg []DdosZoneTemplateDnsFqdnLabelLenCfg `json:"fqdn-label-len-cfg"`
    MalformedQueryCheck DdosZoneTemplateDnsMalformedQueryCheck307 `json:"malformed-query-check"`
    MultiPuThresholdDistribution DdosZoneTemplateDnsMultiPuThresholdDistribution `json:"multi-pu-threshold-distribution"`
    Name string `json:"name"`
    OnNoMatch string `json:"on-no-match" dval:"deny"`
    Src DdosZoneTemplateDnsSrc `json:"src"`
    SymtimeoutCfg DdosZoneTemplateDnsSymtimeoutCfg `json:"symtimeout-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dns"`
}


type DdosZoneTemplateDnsAllowQueryClass struct {
    AllowInternetQueryClass int `json:"allow-internet-query-class"`
    AllowCsnetQueryClass int `json:"allow-csnet-query-class"`
    AllowChaosQueryClass int `json:"allow-chaos-query-class"`
    AllowHesiodQueryClass int `json:"allow-hesiod-query-class"`
    AllowNoneQueryClass int `json:"allow-none-query-class"`
    AllowAnyQueryClass int `json:"allow-any-query-class"`
    AllowQueryClassActionListName string `json:"allow-query-class-action-list-name"`
    AllowQueryClassAction string `json:"allow-query-class-action"`
}


type DdosZoneTemplateDnsAllowRecordType struct {
    AllowAType int `json:"allow-a-type"`
    AllowAaaaType int `json:"allow-aaaa-type"`
    AllowCnameType int `json:"allow-cname-type"`
    AllowMxType int `json:"allow-mx-type"`
    AllowNsType int `json:"allow-ns-type"`
    AllowSrvType int `json:"allow-srv-type"`
    RecordNumCfg []DdosZoneTemplateDnsAllowRecordTypeRecordNumCfg `json:"record-num-cfg"`
    AllowRecordTypeActionListName string `json:"allow-record-type-action-list-name"`
    AllowRecordTypeAction string `json:"allow-record-type-action"`
}


type DdosZoneTemplateDnsAllowRecordTypeRecordNumCfg struct {
    AllowNumType int `json:"allow-num-type"`
}


type DdosZoneTemplateDnsDnsUdpAuthentication struct {
    ForceTcpCfg DdosZoneTemplateDnsDnsUdpAuthenticationForceTcpCfg `json:"force-tcp-cfg"`
    UdpTimeout int `json:"udp-timeout"`
    MinDelay int `json:"min-delay"`
    MinDelayInterval string `json:"min-delay-interval"`
    DnsUdpAuthPassActionListName string `json:"dns-udp-auth-pass-action-list-name"`
    DnsUdpAuthPassAction string `json:"dns-udp-auth-pass-action"`
    DnsUdpAuthFailActionListName string `json:"dns-udp-auth-fail-action-list-name"`
    DnsUdpAuthFailAction string `json:"dns-udp-auth-fail-action"`
}


type DdosZoneTemplateDnsDnsUdpAuthenticationForceTcpCfg struct {
    ForceTcp int `json:"force-tcp"`
    ForceTcpTimeout int `json:"force-tcp-timeout"`
    ForceTcpMinDelay int `json:"force-tcp-min-delay"`
    ForceTcpIgnoreClientSourcePort int `json:"force-tcp-ignore-client-source-port"`
}


type DdosZoneTemplateDnsDst struct {
    RateLimit DdosZoneTemplateDnsDstRateLimit `json:"rate-limit"`
}


type DdosZoneTemplateDnsDstRateLimit struct {
    Fqdn DdosZoneTemplateDnsDstRateLimitFqdn `json:"fqdn"`
    DomainGroupRateExceedAction string `json:"domain-group-rate-exceed-action" dval:"drop"`
    EncapTemplate string `json:"encap-template"`
    DomainGroupRatePerService int `json:"domain-group-rate-per-service"`
    Request DdosZoneTemplateDnsDstRateLimitRequest `json:"request"`
}


type DdosZoneTemplateDnsDstRateLimitFqdn struct {
    DnsFqdnRateCfg []DdosZoneTemplateDnsDstRateLimitFqdnDnsFqdnRateCfg `json:"dns-fqdn-rate-cfg"`
    DnsFqdnRateLimitActionListName string `json:"dns-fqdn-rate-limit-action-list-name"`
    DnsFqdnRateLimitAction string `json:"dns-fqdn-rate-limit-action"`
}


type DdosZoneTemplateDnsDstRateLimitFqdnDnsFqdnRateCfg struct {
    DnsFqdnRate int `json:"dns-fqdn-rate"`
    Per string `json:"per"`
    PerDomainPerSrcIp int `json:"per-domain-per-src-ip"`
    FqdnRateSuffix int `json:"fqdn-rate-suffix"`
    FqdnRateLabelCount int `json:"fqdn-rate-label-count"`
}


type DdosZoneTemplateDnsDstRateLimitRequest struct {
    Type DdosZoneTemplateDnsDstRateLimitRequestType `json:"type"`
    DstDnsRequestRateLimitActionListName string `json:"dst-dns-request-rate-limit-action-list-name"`
    DstDnsRequestRateLimitAction string `json:"dst-dns-request-rate-limit-action"`
}


type DdosZoneTemplateDnsDstRateLimitRequestType struct {
    ACfg DdosZoneTemplateDnsDstRateLimitRequestTypeACfg `json:"A-cfg"`
    AaaaCfg DdosZoneTemplateDnsDstRateLimitRequestTypeAaaaCfg `json:"AAAA-cfg"`
    CnameCfg DdosZoneTemplateDnsDstRateLimitRequestTypeCnameCfg `json:"CNAME-cfg"`
    MxCfg DdosZoneTemplateDnsDstRateLimitRequestTypeMxCfg `json:"MX-cfg"`
    NsCfg DdosZoneTemplateDnsDstRateLimitRequestTypeNsCfg `json:"NS-cfg"`
    SrvCfg DdosZoneTemplateDnsDstRateLimitRequestTypeSrvCfg `json:"SRV-cfg"`
    DnsTypeCfg []DdosZoneTemplateDnsDstRateLimitRequestTypeDnsTypeCfg `json:"dns-type-cfg"`
}


type DdosZoneTemplateDnsDstRateLimitRequestTypeACfg struct {
    A int `json:"A"`
    DnsARate int `json:"dns-a-rate"`
}


type DdosZoneTemplateDnsDstRateLimitRequestTypeAaaaCfg struct {
    Aaaa int `json:"AAAA"`
    DnsAaaaRate int `json:"dns-aaaa-rate"`
}


type DdosZoneTemplateDnsDstRateLimitRequestTypeCnameCfg struct {
    Cname int `json:"CNAME"`
    DnsCnameRate int `json:"dns-cname-rate"`
}


type DdosZoneTemplateDnsDstRateLimitRequestTypeMxCfg struct {
    Mx int `json:"MX"`
    DnsMxRate int `json:"dns-mx-rate"`
}


type DdosZoneTemplateDnsDstRateLimitRequestTypeNsCfg struct {
    Ns int `json:"NS"`
    DnsNsRate int `json:"dns-ns-rate"`
}


type DdosZoneTemplateDnsDstRateLimitRequestTypeSrvCfg struct {
    Srv int `json:"SRV"`
    DnsSrvRate int `json:"dns-srv-rate"`
}


type DdosZoneTemplateDnsDstRateLimitRequestTypeDnsTypeCfg struct {
    DnsRequestType int `json:"dns-request-type"`
    DnsRequestTypeRate int `json:"dns-request-type-rate"`
}


type DdosZoneTemplateDnsFqdnLabelCountCfg struct {
    LabelCount int `json:"label-count"`
    FqdnLabelCountActionListName string `json:"fqdn-label-count-action-list-name"`
    FqdnLabelCountAction string `json:"fqdn-label-count-action"`
}


type DdosZoneTemplateDnsFqdnLabelLenCfg struct {
    LabelLength int `json:"label-length"`
    FqdnLabelSuffix int `json:"fqdn-label-suffix"`
    FqdnLabelLengthActionListName string `json:"fqdn-label-length-action-list-name"`
    FqdnLabelLengthAction string `json:"fqdn-label-length-action"`
}


type DdosZoneTemplateDnsMalformedQueryCheck307 struct {
    ValidationType string `json:"validation-type"`
    NonQueryOpcodeCheck string `json:"non-query-opcode-check"`
    SkipMultiPacketCheck int `json:"skip-multi-packet-check"`
    DnsMalformedQueryActionListName string `json:"dns-malformed-query-action-list-name"`
    DnsMalformedQueryAction string `json:"dns-malformed-query-action"`
    Uuid string `json:"uuid"`
}


type DdosZoneTemplateDnsMultiPuThresholdDistribution struct {
    MultiPuThresholdDistributionValue int `json:"multi-pu-threshold-distribution-value"`
    MultiPuThresholdDistributionDisable string `json:"multi-pu-threshold-distribution-disable"`
}


type DdosZoneTemplateDnsSrc struct {
    RateLimit DdosZoneTemplateDnsSrcRateLimit `json:"rate-limit"`
}


type DdosZoneTemplateDnsSrcRateLimit struct {
    Nxdomain DdosZoneTemplateDnsSrcRateLimitNxdomain `json:"nxdomain"`
    Request DdosZoneTemplateDnsSrcRateLimitRequest `json:"request"`
}


type DdosZoneTemplateDnsSrcRateLimitNxdomain struct {
    DnsNxdomainRate int `json:"dns-nxdomain-rate"`
    DnsNxdomainRateLimitActionListName string `json:"dns-nxdomain-rate-limit-action-list-name"`
    DnsNxdomainRateLimitAction string `json:"dns-nxdomain-rate-limit-action"`
}


type DdosZoneTemplateDnsSrcRateLimitRequest struct {
    Type DdosZoneTemplateDnsSrcRateLimitRequestType `json:"type"`
    SrcDnsRequestRateLimitActionListName string `json:"src-dns-request-rate-limit-action-list-name"`
    SrcDnsRequestRateLimitAction string `json:"src-dns-request-rate-limit-action"`
}


type DdosZoneTemplateDnsSrcRateLimitRequestType struct {
    ACfg DdosZoneTemplateDnsSrcRateLimitRequestTypeACfg `json:"A-cfg"`
    AaaaCfg DdosZoneTemplateDnsSrcRateLimitRequestTypeAaaaCfg `json:"AAAA-cfg"`
    CnameCfg DdosZoneTemplateDnsSrcRateLimitRequestTypeCnameCfg `json:"CNAME-cfg"`
    MxCfg DdosZoneTemplateDnsSrcRateLimitRequestTypeMxCfg `json:"MX-cfg"`
    NsCfg DdosZoneTemplateDnsSrcRateLimitRequestTypeNsCfg `json:"NS-cfg"`
    SrvCfg DdosZoneTemplateDnsSrcRateLimitRequestTypeSrvCfg `json:"SRV-cfg"`
    DnsTypeCfg []DdosZoneTemplateDnsSrcRateLimitRequestTypeDnsTypeCfg `json:"dns-type-cfg"`
}


type DdosZoneTemplateDnsSrcRateLimitRequestTypeACfg struct {
    A int `json:"A"`
    SrcDnsARate int `json:"src-dns-a-rate"`
}


type DdosZoneTemplateDnsSrcRateLimitRequestTypeAaaaCfg struct {
    Aaaa int `json:"AAAA"`
    SrcDnsAaaaRate int `json:"src-dns-aaaa-rate"`
}


type DdosZoneTemplateDnsSrcRateLimitRequestTypeCnameCfg struct {
    Cname int `json:"CNAME"`
    SrcDnsCnameRate int `json:"src-dns-cname-rate"`
}


type DdosZoneTemplateDnsSrcRateLimitRequestTypeMxCfg struct {
    Mx int `json:"MX"`
    SrcDnsMxRate int `json:"src-dns-mx-rate"`
}


type DdosZoneTemplateDnsSrcRateLimitRequestTypeNsCfg struct {
    Ns int `json:"NS"`
    SrcDnsNsRate int `json:"src-dns-ns-rate"`
}


type DdosZoneTemplateDnsSrcRateLimitRequestTypeSrvCfg struct {
    Srv int `json:"SRV"`
    SrcDnsSrvRate int `json:"src-dns-srv-rate"`
}


type DdosZoneTemplateDnsSrcRateLimitRequestTypeDnsTypeCfg struct {
    SrcDnsRequestType int `json:"src-dns-request-type"`
    SrcDnsRequestTypeRate int `json:"src-dns-request-type-rate"`
}


type DdosZoneTemplateDnsSymtimeoutCfg struct {
    SymTimeout int `json:"sym-timeout"`
    SymTimeoutValue int `json:"sym-timeout-value"`
}

func (p *DdosZoneTemplateDns) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosZoneTemplateDns) getPath() string{
    return "ddos/zone-template/dns"
}

func (p *DdosZoneTemplateDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateDns::Post")
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

func (p *DdosZoneTemplateDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateDns::Get")
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
func (p *DdosZoneTemplateDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateDns::Put")
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

func (p *DdosZoneTemplateDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
