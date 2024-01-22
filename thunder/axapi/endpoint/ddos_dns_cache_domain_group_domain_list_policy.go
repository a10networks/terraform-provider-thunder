

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheDomainGroupDomainListPolicy struct {
	Inst struct {

    CacheAllRecords int `json:"cache-all-records"`
    ClientIpv4 string `json:"client-ipv4"`
    ClientIpv6 string `json:"client-ipv6"`
    Force int `json:"force"`
    ManualRefresh string `json:"manual-refresh"`
    Name string `json:"name"`
    OversizeAnswerResponse string `json:"oversize-answer-response" dval:"set-truncate-bit"`
    PacketCapturing DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing146 `json:"packet-capturing"`
    RefreshIntervalHours int `json:"refresh-interval-hours" dval:"4"`
    ResolveCnameRecord int `json:"resolve-cname-record"`
    RespondWithAuthority int `json:"respond-with-authority"`
    ServerIpv4 string `json:"server-ipv4"`
    ServerIpv6 string `json:"server-ipv6"`
    ServerV4Port int `json:"server-v4-port" dval:"53"`
    ServerV6Port int `json:"server-v6-port" dval:"53"`
    TtlOverride int `json:"ttl-override"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"domain-list-policy"`
}


type DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing146 struct {
    RootZoneList []DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList147 `json:"root-zone-list"`
    Uuid string `json:"uuid"`
}


type DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList147 struct {
    RootZone string `json:"root-zone"`
    CaptureConfig string `json:"capture-config"`
    CaptureMode string `json:"capture-mode"`
}

func (p *DdosDnsCacheDomainGroupDomainListPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosDnsCacheDomainGroupDomainListPolicy) getPath() string{
    return "ddos/dns-cache/domain-group/domain-list-policy"
}

func (p *DdosDnsCacheDomainGroupDomainListPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroupDomainListPolicy::Post")
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

func (p *DdosDnsCacheDomainGroupDomainListPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroupDomainListPolicy::Get")
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
func (p *DdosDnsCacheDomainGroupDomainListPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroupDomainListPolicy::Put")
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

func (p *DdosDnsCacheDomainGroupDomainListPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroupDomainListPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
