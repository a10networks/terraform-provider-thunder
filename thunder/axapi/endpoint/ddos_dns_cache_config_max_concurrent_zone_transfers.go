

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheConfigMaxConcurrentZoneTransfers struct {
	Inst struct {

    OperationalMode int `json:"operational-mode"`
    Uuid string `json:"uuid"`
    WarmUpMode int `json:"warm-up-mode" dval:"65472"`

	} `json:"max-concurrent-zone-transfers"`
}

func (p *DdosDnsCacheConfigMaxConcurrentZoneTransfers) GetId() string{
    return "1"
}

func (p *DdosDnsCacheConfigMaxConcurrentZoneTransfers) getPath() string{
    return "ddos/dns-cache-config/max-concurrent-zone-transfers"
}

func (p *DdosDnsCacheConfigMaxConcurrentZoneTransfers) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheConfigMaxConcurrentZoneTransfers::Post")
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

func (p *DdosDnsCacheConfigMaxConcurrentZoneTransfers) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheConfigMaxConcurrentZoneTransfers::Get")
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
func (p *DdosDnsCacheConfigMaxConcurrentZoneTransfers) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheConfigMaxConcurrentZoneTransfers::Put")
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

func (p *DdosDnsCacheConfigMaxConcurrentZoneTransfers) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheConfigMaxConcurrentZoneTransfers::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
