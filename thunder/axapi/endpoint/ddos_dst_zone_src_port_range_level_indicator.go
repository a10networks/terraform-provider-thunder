

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortRangeLevelIndicator struct {
	Inst struct {

    Type string `json:"type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdNum int `json:"zone-threshold-num"`
    SrcPortRangeEnd string 
    ZoneName string 
    Protocol string 
    LevelNum string 
    SrcPortRangeStart string 

	} `json:"indicator"`
}

func (p *DdosDstZoneSrcPortRangeLevelIndicator) GetId() string{
    return p.Inst.Type
}

func (p *DdosDstZoneSrcPortRangeLevelIndicator) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/src-port-range/" +p.Inst.SrcPortRangeStart + "+" +p.Inst.SrcPortRangeEnd + "+" +p.Inst.Protocol + "/level/" +p.Inst.LevelNum + "/indicator"
}

func (p *DdosDstZoneSrcPortRangeLevelIndicator) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRangeLevelIndicator::Post")
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

func (p *DdosDstZoneSrcPortRangeLevelIndicator) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRangeLevelIndicator::Get")
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
func (p *DdosDstZoneSrcPortRangeLevelIndicator) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRangeLevelIndicator::Put")
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

func (p *DdosDstZoneSrcPortRangeLevelIndicator) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRangeLevelIndicator::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
