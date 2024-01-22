

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheShardedDomainGroup struct {
	Inst struct {

    EncapTemplate string `json:"encap-template"`
    MatchAction string `json:"match-action" dval:"forward"`
    Name string `json:"name"`
    ShardedDomainListPolicyList []DdosDnsCacheShardedDomainGroupShardedDomainListPolicyList `json:"sharded-domain-list-policy-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"sharded-domain-group"`
}


type DdosDnsCacheShardedDomainGroupShardedDomainListPolicyList struct {
    Name string `json:"name"`
    ServerIpv4 string `json:"server-ipv4"`
    ServerV4Port int `json:"server-v4-port" dval:"53"`
    ClientIpv4 string `json:"client-ipv4"`
    ServerIpv6 string `json:"server-ipv6"`
    ServerV6Port int `json:"server-v6-port" dval:"53"`
    ClientIpv6 string `json:"client-ipv6"`
    RefreshIntervalHours int `json:"refresh-interval-hours" dval:"4"`
    ManualRefresh string `json:"manual-refresh"`
    Force int `json:"force"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    PacketCapturing DdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturing `json:"packet-capturing"`
}


type DdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturing struct {
    RootZoneList []DdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturingRootZoneList `json:"root-zone-list"`
    Uuid string `json:"uuid"`
}


type DdosDnsCacheShardedDomainGroupShardedDomainListPolicyListPacketCapturingRootZoneList struct {
    RootZone string `json:"root-zone"`
    CaptureConfig string `json:"capture-config"`
    CaptureMode string `json:"capture-mode"`
}

func (p *DdosDnsCacheShardedDomainGroup) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosDnsCacheShardedDomainGroup) getPath() string{
    return "ddos/dns-cache/sharded-domain-group"
}

func (p *DdosDnsCacheShardedDomainGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheShardedDomainGroup::Post")
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

func (p *DdosDnsCacheShardedDomainGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheShardedDomainGroup::Get")
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
func (p *DdosDnsCacheShardedDomainGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheShardedDomainGroup::Put")
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

func (p *DdosDnsCacheShardedDomainGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheShardedDomainGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
