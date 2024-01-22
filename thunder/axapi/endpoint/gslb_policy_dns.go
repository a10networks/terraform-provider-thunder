

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbPolicyDns struct {
	Inst struct {

    Action int `json:"action"`
    ActionType string `json:"action-type"`
    ActiveOnly int `json:"active-only"`
    ActiveOnlyFailSafe int `json:"active-only-fail-safe"`
    AgingTime int `json:"aging-time"`
    BackupAlias int `json:"backup-alias"`
    BackupServer int `json:"backup-server"`
    BlockAction int `json:"block-action"`
    BlockType string `json:"block-type"`
    BlockValue []GslbPolicyDnsBlockValue `json:"block-value"`
    Cache int `json:"cache"`
    CnameDetect int `json:"cname-detect" dval:"1"`
    Delegation int `json:"delegation"`
    DnsAdditionMx int `json:"dns-addition-mx"`
    DnsAutoMap int `json:"dns-auto-map"`
    DynamicPreference int `json:"dynamic-preference"`
    DynamicWeight int `json:"dynamic-weight"`
    ExternalIp int `json:"external-ip" dval:"1"`
    ExternalSoa int `json:"external-soa"`
    GeolocAction int `json:"geoloc-action"`
    GeolocAlias int `json:"geoloc-alias"`
    GeolocPolicy int `json:"geoloc-policy"`
    Hint string `json:"hint" dval:"addition"`
    IpReplace int `json:"ip-replace"`
    Ipv6 []GslbPolicyDnsIpv6 `json:"ipv6"`
    Logging string `json:"logging"`
    ProxyBlockPortRangeList []GslbPolicyDnsProxyBlockPortRangeList `json:"proxy-block-port-range-list"`
    SelectedOnly int `json:"selected-only"`
    SelectedOnlyValue int `json:"selected-only-value"`
    Server int `json:"server"`
    ServerAdditionMx int `json:"server-addition-mx"`
    ServerAny int `json:"server-any"`
    ServerAnyWithMetric int `json:"server-any-with-metric"`
    ServerAuthoritative int `json:"server-authoritative"`
    ServerAutoNs int `json:"server-auto-ns"`
    ServerAutoPtr int `json:"server-auto-ptr"`
    ServerCaa int `json:"server-caa"`
    ServerCname int `json:"server-cname"`
    ServerCustom int `json:"server-custom"`
    ServerFullList int `json:"server-full-list"`
    ServerModeOnly int `json:"server-mode-only"`
    ServerMx int `json:"server-mx"`
    ServerNaptr int `json:"server-naptr"`
    ServerNs int `json:"server-ns"`
    ServerNsList int `json:"server-ns-list"`
    ServerPtr int `json:"server-ptr"`
    ServerSec int `json:"server-sec"`
    ServerSrv int `json:"server-srv"`
    ServerTxt int `json:"server-txt"`
    Sticky int `json:"sticky"`
    StickyAgingTime int `json:"sticky-aging-time" dval:"5"`
    StickyIpv6Mask int `json:"sticky-ipv6-mask" dval:"128"`
    StickyMask string `json:"sticky-mask" dval:"/32"`
    Template string `json:"template"`
    Ttl int `json:"ttl" dval:"10"`
    UseServerTtl int `json:"use-server-ttl"`
    Uuid string `json:"uuid"`
    ZoneOwnerMode int `json:"zone-owner-mode"`
    Name string 

	} `json:"dns"`
}


type GslbPolicyDnsBlockValue struct {
    BlockValue int `json:"block-value"`
}


type GslbPolicyDnsIpv6 struct {
    DnsIpv6Option string `json:"dns-ipv6-option"`
    DnsIpv6MappingType string `json:"dns-ipv6-mapping-type"`
}


type GslbPolicyDnsProxyBlockPortRangeList struct {
    ProxyBlockRangeFrom int `json:"proxy-block-range-from"`
    ProxyBlockRangeTo int `json:"proxy-block-range-to"`
}

func (p *GslbPolicyDns) GetId() string{
    return "1"
}

func (p *GslbPolicyDns) getPath() string{
    return "gslb/policy/" +p.Inst.Name + "/dns"
}

func (p *GslbPolicyDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyDns::Post")
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

func (p *GslbPolicyDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyDns::Get")
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
func (p *GslbPolicyDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyDns::Put")
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

func (p *GslbPolicyDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbPolicyDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
