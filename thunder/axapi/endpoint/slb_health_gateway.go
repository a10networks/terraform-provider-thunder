

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthGateway struct {
	Inst struct {

    SamplingEnable []SlbHealthGatewaySamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"health-gateway"`
}


type SlbHealthGatewaySamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbHealthGateway) GetId() string{
    return "1"
}

func (p *SlbHealthGateway) getPath() string{
    return "slb/health-gateway"
}

func (p *SlbHealthGateway) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbHealthGateway::Post")
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

func (p *SlbHealthGateway) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbHealthGateway::Get")
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
func (p *SlbHealthGateway) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbHealthGateway::Put")
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

func (p *SlbHealthGateway) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbHealthGateway::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
