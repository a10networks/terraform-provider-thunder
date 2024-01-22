

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheDomainGroup struct {
	Inst struct {

    DomainListPolicyList []DdosDnsCacheDomainGroupDomainListPolicyList `json:"domain-list-policy-list"`
    Name string `json:"name"`
    Uuid string `json:"uuid"`

	} `json:"domain-group"`
}


type DdosDnsCacheDomainGroupDomainListPolicyList struct {
    Name string `json:"name"`
    ServerIpv4 string `json:"server-ipv4"`
    ServerV4Port int `json:"server-v4-port" dval:"53"`
    ClientIpv4 string `json:"client-ipv4"`
    ServerIpv6 string `json:"server-ipv6"`
    ServerV6Port int `json:"server-v6-port" dval:"53"`
    ClientIpv6 string `json:"client-ipv6"`
    RefreshIntervalHours int `json:"refresh-interval-hours" dval:"4"`
    TtlOverride int `json:"ttl-override"`
    RespondWithAuthority int `json:"respond-with-authority"`
    OversizeAnswerResponse string `json:"oversize-answer-response" dval:"set-truncate-bit"`
    ResolveCnameRecord int `json:"resolve-cname-record"`
    ManualRefresh string `json:"manual-refresh"`
    Force int `json:"force"`
    CacheAllRecords int `json:"cache-all-records"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    PacketCapturing DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing `json:"packet-capturing"`
}


type DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing struct {
    RootZoneList []DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList `json:"root-zone-list"`
    Uuid string `json:"uuid"`
}


type DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList struct {
    RootZone string `json:"root-zone"`
    CaptureConfig string `json:"capture-config"`
    CaptureMode string `json:"capture-mode"`
}

func (p *DdosDnsCacheDomainGroup) GetId() string{
    return "1"
}

func (p *DdosDnsCacheDomainGroup) getPath() string{
    return "ddos/dns-cache/"+p.Inst.Name+"/domain-group"
}

func (p *DdosDnsCacheDomainGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroup::Post")
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

func (p *DdosDnsCacheDomainGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroup::Get")
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
func (p *DdosDnsCacheDomainGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroup::Put")
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

func (p *DdosDnsCacheDomainGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
