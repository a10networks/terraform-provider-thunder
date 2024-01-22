

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ZoneLocalZoneCfg struct {
	Inst struct {

    LocalType int `json:"local-type"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"local-zone-cfg"`
}

func (p *ZoneLocalZoneCfg) GetId() string{
    return "1"
}

func (p *ZoneLocalZoneCfg) getPath() string{
    return "zone/" +p.Inst.Name + "/local-zone-cfg"
}

func (p *ZoneLocalZoneCfg) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneLocalZoneCfg::Post")
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

func (p *ZoneLocalZoneCfg) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneLocalZoneCfg::Get")
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
func (p *ZoneLocalZoneCfg) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneLocalZoneCfg::Put")
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

func (p *ZoneLocalZoneCfg) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneLocalZoneCfg::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
