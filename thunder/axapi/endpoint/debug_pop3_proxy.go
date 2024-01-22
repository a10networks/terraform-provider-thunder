

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugPop3Proxy struct {
	Inst struct {

    Dumy int `json:"dumy"`
    Uuid string `json:"uuid"`

	} `json:"pop3-proxy"`
}

func (p *DebugPop3Proxy) GetId() string{
    return "1"
}

func (p *DebugPop3Proxy) getPath() string{
    return "debug/pop3-proxy"
}

func (p *DebugPop3Proxy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugPop3Proxy::Post")
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

func (p *DebugPop3Proxy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugPop3Proxy::Get")
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
func (p *DebugPop3Proxy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugPop3Proxy::Put")
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

func (p *DebugPop3Proxy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugPop3Proxy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
