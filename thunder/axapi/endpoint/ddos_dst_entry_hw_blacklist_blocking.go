

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryHwBlacklistBlocking struct {
	Inst struct {

    DstEnable int `json:"dst-enable"`
    SrcEnable int `json:"src-enable"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"hw-blacklist-blocking"`
}

func (p *DdosDstEntryHwBlacklistBlocking) GetId() string{
    return "1"
}

func (p *DdosDstEntryHwBlacklistBlocking) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/hw-blacklist-blocking"
}

func (p *DdosDstEntryHwBlacklistBlocking) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryHwBlacklistBlocking::Post")
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

func (p *DdosDstEntryHwBlacklistBlocking) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryHwBlacklistBlocking::Get")
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
func (p *DdosDstEntryHwBlacklistBlocking) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryHwBlacklistBlocking::Put")
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

func (p *DdosDstEntryHwBlacklistBlocking) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryHwBlacklistBlocking::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
