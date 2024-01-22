

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTrunkHwHash struct {
	Inst struct {

    Mode int `json:"mode" dval:"6"`
    Uuid string `json:"uuid"`

	} `json:"trunk-hw-hash"`
}

func (p *SystemTrunkHwHash) GetId() string{
    return "1"
}

func (p *SystemTrunkHwHash) getPath() string{
    return "system/trunk-hw-hash"
}

func (p *SystemTrunkHwHash) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTrunkHwHash::Post")
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

func (p *SystemTrunkHwHash) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTrunkHwHash::Get")
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
func (p *SystemTrunkHwHash) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTrunkHwHash::Put")
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

func (p *SystemTrunkHwHash) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTrunkHwHash::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
