

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpsGlobal struct {
	Inst struct {

    SamplingEnable []IpsGlobalSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"global"`
}


type IpsGlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *IpsGlobal) GetId() string{
    return "1"
}

func (p *IpsGlobal) getPath() string{
    return "ips/global"
}

func (p *IpsGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpsGlobal::Post")
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

func (p *IpsGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpsGlobal::Get")
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
func (p *IpsGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpsGlobal::Put")
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

func (p *IpsGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpsGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
