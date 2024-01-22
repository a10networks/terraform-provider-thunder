

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIcmp6 struct {
	Inst struct {

    SamplingEnable []SystemIcmp6SamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"icmp6"`
}


type SystemIcmp6SamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemIcmp6) GetId() string{
    return "1"
}

func (p *SystemIcmp6) getPath() string{
    return "system/icmp6"
}

func (p *SystemIcmp6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIcmp6::Post")
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

func (p *SystemIcmp6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIcmp6::Get")
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
func (p *SystemIcmp6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIcmp6::Put")
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

func (p *SystemIcmp6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIcmp6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
