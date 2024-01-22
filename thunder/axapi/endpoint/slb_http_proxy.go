

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHttpProxy struct {
	Inst struct {

    SamplingEnable []SlbHttpProxySamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"http-proxy"`
}


type SlbHttpProxySamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
}

func (p *SlbHttpProxy) GetId() string{
    return "1"
}

func (p *SlbHttpProxy) getPath() string{
    return "slb/http-proxy"
}

func (p *SlbHttpProxy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbHttpProxy::Post")
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

func (p *SlbHttpProxy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbHttpProxy::Get")
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
func (p *SlbHttpProxy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbHttpProxy::Put")
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

func (p *SlbHttpProxy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbHttpProxy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
