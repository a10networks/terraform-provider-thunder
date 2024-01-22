

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheShardedDomainGroupShardedDomainListPolicy struct {
	Inst struct {

    ClientIpv4 string `json:"client-ipv4"`
    ClientIpv6 string `json:"client-ipv6"`
    Force int `json:"force"`
    ManualRefresh string `json:"manual-refresh"`
    Name string `json:"name"`
    PacketCapturing DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing148 `json:"packet-capturing"`
    RefreshIntervalHours int `json:"refresh-interval-hours" dval:"4"`
    ServerIpv4 string `json:"server-ipv4"`
    ServerIpv6 string `json:"server-ipv6"`
    ServerV4Port int `json:"server-v4-port" dval:"53"`
    ServerV6Port int `json:"server-v6-port" dval:"53"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"sharded-domain-list-policy"`
}


type DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing148 struct {
    RootZoneList []DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList149 `json:"root-zone-list"`
    Uuid string `json:"uuid"`
}


type DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList149 struct {
    RootZone string `json:"root-zone"`
    CaptureConfig string `json:"capture-config"`
    CaptureMode string `json:"capture-mode"`
}

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicy) getPath() string{
    return "ddos/dns-cache/sharded-domain-group/sharded-domain-list-policy"
}

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheShardedDomainGroupShardedDomainListPolicy::Post")
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

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheShardedDomainGroupShardedDomainListPolicy::Get")
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
func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheShardedDomainGroupShardedDomainListPolicy::Put")
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

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheShardedDomainGroupShardedDomainListPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
