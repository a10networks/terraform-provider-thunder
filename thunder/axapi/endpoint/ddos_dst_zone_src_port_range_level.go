

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortRangeLevel struct {
	Inst struct {

    IndicatorList []DdosDstZoneSrcPortRangeLevelIndicatorList `json:"indicator-list"`
    LevelNum string `json:"level-num"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    SrcPortRangeEnd string 
    ZoneName string 
    Protocol string 
    SrcPortRangeStart string 

	} `json:"level"`
}


type DdosDstZoneSrcPortRangeLevelIndicatorList struct {
    Type string `json:"type"`
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosDstZoneSrcPortRangeLevel) GetId() string{
    return p.Inst.LevelNum
}

func (p *DdosDstZoneSrcPortRangeLevel) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/src-port-range/" +p.Inst.SrcPortRangeStart + "+" +p.Inst.SrcPortRangeEnd + "+" +p.Inst.Protocol + "/level"
}

func (p *DdosDstZoneSrcPortRangeLevel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRangeLevel::Post")
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

func (p *DdosDstZoneSrcPortRangeLevel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRangeLevel::Get")
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
func (p *DdosDstZoneSrcPortRangeLevel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRangeLevel::Put")
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

func (p *DdosDstZoneSrcPortRangeLevel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRangeLevel::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
