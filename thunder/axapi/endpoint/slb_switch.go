

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSwitch struct {
	Inst struct {

    SamplingEnable []SlbSwitchSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"switch"`
}


type SlbSwitchSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbSwitch) GetId() string{
    return "1"
}

func (p *SlbSwitch) getPath() string{
    return "slb/switch"
}

func (p *SlbSwitch) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSwitch::Post")
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

func (p *SlbSwitch) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSwitch::Get")
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
func (p *SlbSwitch) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSwitch::Put")
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

func (p *SlbSwitch) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSwitch::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
