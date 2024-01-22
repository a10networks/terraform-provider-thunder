

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPop3Proxy struct {
	Inst struct {

    SamplingEnable []SlbPop3ProxySamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"pop3-proxy"`
}


type SlbPop3ProxySamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbPop3Proxy) GetId() string{
    return "1"
}

func (p *SlbPop3Proxy) getPath() string{
    return "slb/pop3-proxy"
}

func (p *SlbPop3Proxy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPop3Proxy::Post")
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

func (p *SlbPop3Proxy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPop3Proxy::Get")
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
func (p *SlbPop3Proxy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPop3Proxy::Put")
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

func (p *SlbPop3Proxy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPop3Proxy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
