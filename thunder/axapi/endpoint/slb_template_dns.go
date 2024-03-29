

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDns struct {
	Inst struct {

    AddPaddingToClient string `json:"add-padding-to-client"`
    CacheRecordServingPolicy string `json:"cache-record-serving-policy"`
    CacheTtlAdjustmentEnable int `json:"cache-ttl-adjustment-enable"`
    ClassList SlbTemplateDnsClassList1424 `json:"class-list"`
    DefaultPolicy string `json:"default-policy" dval:"nocache"`
    DisableDnsTemplate int `json:"disable-dns-template"`
    DisableRaCachedResp int `json:"disable-ra-cached-resp"`
    DisableRpzAttachSoa int `json:"disable-rpz-attach-soa"`
    DnsLogging string `json:"dns-logging"`
    Dns64 SlbTemplateDnsDns641427 `json:"dns64"`
    DnssecServiceGroup string `json:"dnssec-service-group"`
    Drop int `json:"drop"`
    EnableCacheSharing int `json:"enable-cache-sharing"`
    Forward string `json:"forward"`
    InsertIpv4 int `json:"insert-ipv4"`
    InsertIpv6 int `json:"insert-ipv6"`
    LocalDnsResolution SlbTemplateDnsLocalDnsResolution1428 `json:"local-dns-resolution"`
    MaxCacheEntrySize int `json:"max-cache-entry-size" dval:"1024"`
    MaxCacheSize int `json:"max-cache-size"`
    MaxQueryLength int `json:"max-query-length"`
    Name string `json:"name"`
    NegativeDnsCache SlbTemplateDnsNegativeDnsCache1431 `json:"negative-dns-cache"`
    Period int `json:"period"`
    QueryClassFilter SlbTemplateDnsQueryClassFilter1432 `json:"query-class-filter"`
    QueryIdSwitch int `json:"query-id-switch"`
    QueryTypeFilter SlbTemplateDnsQueryTypeFilter1434 `json:"query-type-filter"`
    RecursiveDnsResolution SlbTemplateDnsRecursiveDnsResolution1436 `json:"recursive-dns-resolution"`
    RedirectToTcpPort int `json:"redirect-to-tcp-port"`
    RemoveAaFlag int `json:"remove-aa-flag"`
    RemoveCsubnet int `json:"remove-csubnet"`
    RemovePaddingToServer int `json:"remove-padding-to-server"`
    ResponseRateLimiting SlbTemplateDnsResponseRateLimiting1441 `json:"response-rate-limiting"`
    RpzList []SlbTemplateDnsRpzList `json:"rpz-list"`
    UdpRetransmit SlbTemplateDnsUdpRetransmit1444 `json:"udp-retransmit"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dns"`
}


type SlbTemplateDnsClassList1424 struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    LidList []SlbTemplateDnsClassListLidList1425 `json:"lid-list"`
}


type SlbTemplateDnsClassListLidList1425 struct {
    Lidnum int `json:"lidnum"`
    ConnRateLimit int `json:"conn-rate-limit"`
    Per int `json:"per"`
    OverLimitAction int `json:"over-limit-action"`
    ActionValue string `json:"action-value"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    LogInterval int `json:"log-interval"`
    Dns SlbTemplateDnsClassListLidListDns1426 `json:"dns"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SlbTemplateDnsClassListLidListDns1426 struct {
    CacheAction string `json:"cache-action" dval:"cache-disable"`
    Ttl int `json:"ttl"`
    Weight int `json:"weight"`
    HonorServerResponseTtl int `json:"honor-server-response-ttl"`
}


type SlbTemplateDnsDns641427 struct {
    Enable int `json:"enable"`
    Cache int `json:"cache"`
    ChangeQuery int `json:"change-query"`
    ParallelQuery int `json:"parallel-query"`
    Retry int `json:"retry" dval:"3"`
    SingleResponseDisable int `json:"single-response-disable"`
    Timeout int `json:"timeout" dval:"1"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsLocalDnsResolution1428 struct {
    HostListCfg []SlbTemplateDnsLocalDnsResolutionHostListCfg1429 `json:"host-list-cfg"`
    LocalResolverCfg []SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1430 `json:"local-resolver-cfg"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsLocalDnsResolutionHostListCfg1429 struct {
    Hostnames string `json:"hostnames"`
}


type SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1430 struct {
    LocalResolver string `json:"local-resolver"`
}


type SlbTemplateDnsNegativeDnsCache1431 struct {
    EnableNegativeDnsCache int `json:"enable-negative-dns-cache"`
    BypassQueryThreshold int `json:"bypass-query-threshold" dval:"100"`
    MaxNegativeCacheTtl int `json:"max-negative-cache-ttl" dval:"7200"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsQueryClassFilter1432 struct {
    QueryClassAction string `json:"query-class-action"`
    QueryClass []SlbTemplateDnsQueryClassFilterQueryClass1433 `json:"query-class"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsQueryClassFilterQueryClass1433 struct {
    StrQueryClass string `json:"str-query-class"`
    NumQueryClass int `json:"num-query-class"`
}


type SlbTemplateDnsQueryTypeFilter1434 struct {
    QueryTypeAction string `json:"query-type-action"`
    QueryType []SlbTemplateDnsQueryTypeFilterQueryType1435 `json:"query-type"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsQueryTypeFilterQueryType1435 struct {
    StrQueryType string `json:"str-query-type"`
    NumQueryType int `json:"num-query-type"`
}


type SlbTemplateDnsRecursiveDnsResolution1436 struct {
    HostListCfg []SlbTemplateDnsRecursiveDnsResolutionHostListCfg1437 `json:"host-list-cfg"`
    CsubnetRetry int `json:"csubnet-retry"`
    NsCacheLookup string `json:"ns-cache-lookup" dval:"enabled"`
    UseServiceGroupResponse string `json:"use-service-group-response" dval:"enabled"`
    Ipv4NatPool string `json:"ipv4-nat-pool"`
    Ipv6NatPool string `json:"ipv6-nat-pool"`
    RetriesPerLevel int `json:"retries-per-level" dval:"6"`
    FullResponse int `json:"full-response"`
    MaxTrials int `json:"max-trials" dval:"255"`
    RequestForPendingResolution string `json:"request-for-pending-resolution" dval:"respond-with-servfail"`
    UdpRetryInterval int `json:"udp-retry-interval" dval:"1"`
    UdpInitialInterval int `json:"udp-initial-interval" dval:"5"`
    UseClientQid int `json:"use-client-qid"`
    DefaultRecursive int `json:"default-recursive"`
    ForceCnameResolution string `json:"force-cname-resolution" dval:"enabled"`
    FastNsSelection string `json:"fast-ns-selection" dval:"enabled"`
    DnssecValidation string `json:"dnssec-validation" dval:"disabled"`
    Uuid string `json:"uuid"`
    LookupOrder SlbTemplateDnsRecursiveDnsResolutionLookupOrder1438 `json:"lookup-order"`
    GatewayHealthCheck SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1440 `json:"gateway-health-check"`
}


type SlbTemplateDnsRecursiveDnsResolutionHostListCfg1437 struct {
    Hostnames string `json:"hostnames"`
}


type SlbTemplateDnsRecursiveDnsResolutionLookupOrder1438 struct {
    QueryType []SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1439 `json:"query-type"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1439 struct {
    StrQueryType string `json:"str-query-type"`
    NumQueryType int `json:"num-query-type"`
    Order string `json:"order"`
}


type SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1440 struct {
    QueryName string `json:"query-name" dval:"a10networks.com"`
    Retry int `json:"retry" dval:"6"`
    Timeout int `json:"timeout" dval:"5"`
    Interval int `json:"interval" dval:"10"`
    UpRetry int `json:"up-retry" dval:"1"`
    RetryMulti int `json:"retry-multi" dval:"1"`
    GwhcNsCacheLookup string `json:"gwhc-ns-cache-lookup" dval:"disabled"`
    StrQueryType string `json:"str-query-type" dval:"A"`
    NumQueryType int `json:"num-query-type"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsResponseRateLimiting1441 struct {
    ResponseRate int `json:"response-rate" dval:"5"`
    FilterResponseRate int `json:"filter-response-rate" dval:"10"`
    SlipRate int `json:"slip-rate"`
    TcRate int `json:"TC-rate"`
    MatchSubnet string `json:"match-subnet" dval:"255.255.255.255"`
    MatchSubnetV6 int `json:"match-subnet-v6" dval:"128"`
    Window int `json:"window" dval:"1"`
    SrcIpOnly int `json:"src-ip-only"`
    EnableLog int `json:"enable-log"`
    Action string `json:"action" dval:"rate-limit"`
    Uuid string `json:"uuid"`
    RrlClassListList []SlbTemplateDnsResponseRateLimitingRrlClassListList1442 `json:"rrl-class-list-list"`
}


type SlbTemplateDnsResponseRateLimitingRrlClassListList1442 struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    LidList []SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1443 `json:"lid-list"`
}


type SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1443 struct {
    Lidnum int `json:"lidnum"`
    LidResponseRate int `json:"lid-response-rate" dval:"5"`
    LidSlipRate int `json:"lid-slip-rate"`
    LidTcRate int `json:"lid-tc-rate"`
    LidMatchSubnet string `json:"lid-match-subnet" dval:"255.255.255.255"`
    LidMatchSubnetV6 int `json:"lid-match-subnet-v6" dval:"128"`
    LidWindow int `json:"lid-window" dval:"1"`
    LidSrcIpOnly int `json:"lid-src-ip-only"`
    LidEnableLog int `json:"lid-enable-log"`
    LidAction string `json:"lid-action" dval:"rate-limit"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SlbTemplateDnsRpzList struct {
    SeqId int `json:"seq-id"`
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    Logging SlbTemplateDnsRpzListLogging `json:"logging"`
}


type SlbTemplateDnsRpzListLogging struct {
    Enable int `json:"enable"`
    RpzAction []SlbTemplateDnsRpzListLoggingRpzAction `json:"rpz-action"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsRpzListLoggingRpzAction struct {
    StrRpzAction string `json:"str-rpz-action"`
}


type SlbTemplateDnsUdpRetransmit1444 struct {
    RetryInterval int `json:"retry-interval" dval:"10"`
    MaxTrials int `json:"max-trials" dval:"3"`
    Uuid string `json:"uuid"`
}

func (p *SlbTemplateDns) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateDns) getPath() string{
    return "slb/template/dns"
}

func (p *SlbTemplateDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDns::Post")
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

func (p *SlbTemplateDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDns::Get")
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
func (p *SlbTemplateDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDns::Put")
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

func (p *SlbTemplateDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
