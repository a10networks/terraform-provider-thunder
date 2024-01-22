

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceVeMap struct {
	Inst struct {

    Inside int `json:"inside"`
    MapTInside int `json:"map-t-inside"`
    MapTOutside int `json:"map-t-outside"`
    Outside int `json:"outside"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"map"`
}

func (p *InterfaceVeMap) GetId() string{
    return "1"
}

func (p *InterfaceVeMap) getPath() string{
    return "interface/ve/" +p.Inst.Ifnum + "/map"
}

func (p *InterfaceVeMap) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeMap::Post")
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

func (p *InterfaceVeMap) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeMap::Get")
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
func (p *InterfaceVeMap) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeMap::Put")
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

func (p *InterfaceVeMap) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeMap::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
