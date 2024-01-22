

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type OverlayMgmtInfo struct {
	Inst struct {

    Appstring string `json:"appstring"`
    Plugin_name string `json:"plugin_name"`
    Uuid string `json:"uuid"`

	} `json:"overlay-mgmt-info"`
}

func (p *OverlayMgmtInfo) GetId() string{
    return p.Inst.Plugin_name+"+"+p.Inst.Appstring
}

func (p *OverlayMgmtInfo) getPath() string{
    return "overlay-mgmt-info"
}

func (p *OverlayMgmtInfo) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayMgmtInfo::Post")
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

func (p *OverlayMgmtInfo) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayMgmtInfo::Get")
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
func (p *OverlayMgmtInfo) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayMgmtInfo::Put")
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

func (p *OverlayMgmtInfo) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayMgmtInfo::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
