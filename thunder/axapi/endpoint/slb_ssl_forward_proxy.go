

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslForwardProxy struct {
	Inst struct {

    SamplingEnable []SlbSslForwardProxySamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"ssl-forward-proxy"`
}


type SlbSslForwardProxySamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbSslForwardProxy) GetId() string{
    return "1"
}

func (p *SlbSslForwardProxy) getPath() string{
    return "slb/ssl-forward-proxy"
}

func (p *SlbSslForwardProxy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslForwardProxy::Post")
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

func (p *SlbSslForwardProxy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslForwardProxy::Get")
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
func (p *SlbSslForwardProxy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslForwardProxy::Put")
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

func (p *SlbSslForwardProxy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslForwardProxy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
