

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemPortCount struct {
	Inst struct {

    PortCountAlg int `json:"port-count-alg" dval:"6000"`
    PortCountHm int `json:"port-count-hm" dval:"1024"`
    PortCountKernel int `json:"port-count-kernel" dval:"18000"`
    PortCountLogging int `json:"port-count-logging" dval:"4096"`
    Uuid string `json:"uuid"`

	} `json:"port-count"`
}

func (p *SystemPortCount) GetId() string{
    return "1"
}

func (p *SystemPortCount) getPath() string{
    return "system/port-count"
}

func (p *SystemPortCount) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPortCount::Post")
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

func (p *SystemPortCount) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPortCount::Get")
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
func (p *SystemPortCount) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPortCount::Put")
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

func (p *SystemPortCount) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPortCount::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
