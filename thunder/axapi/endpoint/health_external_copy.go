

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthExternalCopy struct {
	Inst struct {

    DstFile string `json:"dst-file"`
    SrcFile string `json:"src-file"`

	} `json:"copy"`
}

func (p *HealthExternalCopy) GetId() string{
    return "1"
}

func (p *HealthExternalCopy) getPath() string{
    return "health/external/copy"
}

func (p *HealthExternalCopy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthExternalCopy::Post")
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

func (p *HealthExternalCopy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthExternalCopy::Get")
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
func (p *HealthExternalCopy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthExternalCopy::Put")
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

func (p *HealthExternalCopy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthExternalCopy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
