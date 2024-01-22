

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcEntryHwBlacklistBlocking struct {
	Inst struct {

    SrcEnable int `json:"src-enable"`
    Uuid string `json:"uuid"`
    SrcEntryName string 

	} `json:"hw-blacklist-blocking"`
}

func (p *DdosSrcEntryHwBlacklistBlocking) GetId() string{
    return "1"
}

func (p *DdosSrcEntryHwBlacklistBlocking) getPath() string{
    return "ddos/src/entry/" +p.Inst.SrcEntryName + "/hw-blacklist-blocking"
}

func (p *DdosSrcEntryHwBlacklistBlocking) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntryHwBlacklistBlocking::Post")
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

func (p *DdosSrcEntryHwBlacklistBlocking) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntryHwBlacklistBlocking::Get")
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
func (p *DdosSrcEntryHwBlacklistBlocking) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntryHwBlacklistBlocking::Put")
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

func (p *DdosSrcEntryHwBlacklistBlocking) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcEntryHwBlacklistBlocking::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
