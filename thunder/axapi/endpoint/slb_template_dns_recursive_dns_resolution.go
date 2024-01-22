

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsRecursiveDnsResolution struct {
	Inst struct {

    CsubnetRetry int `json:"csubnet-retry"`
    DefaultRecursive int `json:"default-recursive"`
    DnssecValidation string `json:"dnssec-validation" dval:"disabled"`
    FastNsSelection string `json:"fast-ns-selection" dval:"enabled"`
    ForceCnameResolution string `json:"force-cname-resolution" dval:"enabled"`
    FullResponse int `json:"full-response"`
    GatewayHealthCheck SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1419 `json:"gateway-health-check"`
    HostListCfg []SlbTemplateDnsRecursiveDnsResolutionHostListCfg `json:"host-list-cfg"`
    Ipv4NatPool string `json:"ipv4-nat-pool"`
    Ipv6NatPool string `json:"ipv6-nat-pool"`
    LookupOrder SlbTemplateDnsRecursiveDnsResolutionLookupOrder1420 `json:"lookup-order"`
    MaxTrials int `json:"max-trials" dval:"255"`
    NsCacheLookup string `json:"ns-cache-lookup" dval:"enabled"`
    RequestForPendingResolution string `json:"request-for-pending-resolution" dval:"respond-with-servfail"`
    RetriesPerLevel int `json:"retries-per-level" dval:"6"`
    UdpInitialInterval int `json:"udp-initial-interval" dval:"5"`
    UdpRetryInterval int `json:"udp-retry-interval" dval:"1"`
    UseClientQid int `json:"use-client-qid"`
    UseServiceGroupResponse string `json:"use-service-group-response" dval:"enabled"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"recursive-dns-resolution"`
}


type SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1419 struct {
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


type SlbTemplateDnsRecursiveDnsResolutionHostListCfg struct {
    Hostnames string `json:"hostnames"`
}


type SlbTemplateDnsRecursiveDnsResolutionLookupOrder1420 struct {
    QueryType []SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1421 `json:"query-type"`
    Uuid string `json:"uuid"`
}


type SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1421 struct {
    StrQueryType string `json:"str-query-type"`
    NumQueryType int `json:"num-query-type"`
    Order string `json:"order"`
}

func (p *SlbTemplateDnsRecursiveDnsResolution) GetId() string{
    return "1"
}

func (p *SlbTemplateDnsRecursiveDnsResolution) getPath() string{
    return "slb/template/dns/" +p.Inst.Name + "/recursive-dns-resolution"
}

func (p *SlbTemplateDnsRecursiveDnsResolution) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolution::Post")
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

func (p *SlbTemplateDnsRecursiveDnsResolution) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolution::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *SlbTemplateDnsRecursiveDnsResolution) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolution::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *SlbTemplateDnsRecursiveDnsResolution) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolution::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
