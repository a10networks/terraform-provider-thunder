

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheConfig struct {
	Inst struct {

    DisableZoneTransferInOperMode int `json:"disable-zone-transfer-in-oper-mode"`
    DisableZoneTransferInWarmUpMode int `json:"disable-zone-transfer-in-warm-up-mode"`
    EnableCacheWarmUpBgpAdvertise int `json:"enable-cache-warm-up-bgp-advertise"`
    MaxConcurrentZoneTransfers DdosDnsCacheConfigMaxConcurrentZoneTransfers145 `json:"max-concurrent-zone-transfers"`
    Uuid string `json:"uuid"`

	} `json:"dns-cache-config"`
}


type DdosDnsCacheConfigMaxConcurrentZoneTransfers145 struct {
    WarmUpMode int `json:"warm-up-mode" dval:"65472"`
    OperationalMode int `json:"operational-mode"`
    Uuid string `json:"uuid"`
}

func (p *DdosDnsCacheConfig) GetId() string{
    return "1"
}

func (p *DdosDnsCacheConfig) getPath() string{
    return "ddos/dns-cache-config"
}

func (p *DdosDnsCacheConfig) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheConfig::Post")
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

func (p *DdosDnsCacheConfig) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheConfig::Get")
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
func (p *DdosDnsCacheConfig) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheConfig::Put")
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

func (p *DdosDnsCacheConfig) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheConfig::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
