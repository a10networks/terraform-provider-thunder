

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbProtocolEnable struct {
	Inst struct {

    Type string `json:"type"`
    Uuid string `json:"uuid"`

	} `json:"enable"`
}

func (p *GslbProtocolEnable) GetId() string{
    return p.Inst.Type
}

func (p *GslbProtocolEnable) getPath() string{
    return "gslb/protocol/enable"
}

func (p *GslbProtocolEnable) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocolEnable::Post")
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

func (p *GslbProtocolEnable) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocolEnable::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *GslbProtocolEnable) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocolEnable::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *GslbProtocolEnable) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocolEnable::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
