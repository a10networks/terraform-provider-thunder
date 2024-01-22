

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbDnsPersistentCache struct {
	Inst struct {

    SamplingEnable []SlbDnsPersistentCacheSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"dns-persistent-cache"`
}


type SlbDnsPersistentCacheSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbDnsPersistentCache) GetId() string{
    return "1"
}

func (p *SlbDnsPersistentCache) getPath() string{
    return "slb/dns-persistent-cache"
}

func (p *SlbDnsPersistentCache) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbDnsPersistentCache::Post")
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

func (p *SlbDnsPersistentCache) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbDnsPersistentCache::Get")
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
func (p *SlbDnsPersistentCache) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbDnsPersistentCache::Put")
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

func (p *SlbDnsPersistentCache) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbDnsPersistentCache::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
