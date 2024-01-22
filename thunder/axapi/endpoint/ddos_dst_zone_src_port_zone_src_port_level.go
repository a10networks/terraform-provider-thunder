

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortZoneSrcPortLevel struct {
	Inst struct {

    IndicatorList []DdosDstZoneSrcPortZoneSrcPortLevelIndicatorList `json:"indicator-list"`
    LevelNum string `json:"level-num"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneName string 
    PortNum string 
    Protocol string 

	} `json:"level"`
}


type DdosDstZoneSrcPortZoneSrcPortLevelIndicatorList struct {
    Type string `json:"type"`
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosDstZoneSrcPortZoneSrcPortLevel) GetId() string{
    return p.Inst.LevelNum
}

func (p *DdosDstZoneSrcPortZoneSrcPortLevel) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/src-port/zone-src-port/" +p.Inst.PortNum + "+" +p.Inst.Protocol + "/level"
}

func (p *DdosDstZoneSrcPortZoneSrcPortLevel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPortLevel::Post")
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

func (p *DdosDstZoneSrcPortZoneSrcPortLevel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPortLevel::Get")
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
func (p *DdosDstZoneSrcPortZoneSrcPortLevel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPortLevel::Put")
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

func (p *DdosDstZoneSrcPortZoneSrcPortLevel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPortLevel::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
