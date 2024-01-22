

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTrunkXauiHwHash struct {
	Inst struct {

    Mode int `json:"mode" dval:"6"`
    Uuid string `json:"uuid"`

	} `json:"trunk-xaui-hw-hash"`
}

func (p *SystemTrunkXauiHwHash) GetId() string{
    return "1"
}

func (p *SystemTrunkXauiHwHash) getPath() string{
    return "system/trunk-xaui-hw-hash"
}

func (p *SystemTrunkXauiHwHash) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTrunkXauiHwHash::Post")
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

func (p *SystemTrunkXauiHwHash) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTrunkXauiHwHash::Get")
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
func (p *SystemTrunkXauiHwHash) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTrunkXauiHwHash::Put")
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

func (p *SystemTrunkXauiHwHash) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTrunkXauiHwHash::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
