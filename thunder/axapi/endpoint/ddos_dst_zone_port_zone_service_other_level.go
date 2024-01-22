

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherLevel struct {
	Inst struct {

    ApplyExtractedFilters int `json:"apply-extracted-filters"`
    CloseSessionsForUnauthSources int `json:"close-sessions-for-unauth-sources"`
    GlidAction string `json:"glid-action"`
    IndicatorList []DdosDstZonePortZoneServiceOtherLevelIndicatorList `json:"indicator-list"`
    LevelNum string `json:"level-num"`
    SrcDefaultGlid string `json:"src-default-glid"`
    SrcEscalationScore int `json:"src-escalation-score"`
    SrcViolationActions string `json:"src-violation-actions"`
    StartPatternRecognition int `json:"start-pattern-recognition"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneEscalationScore int `json:"zone-escalation-score"`
    ZoneTemplate DdosDstZonePortZoneServiceOtherLevelZoneTemplate `json:"zone-template"`
    ZoneViolationActions string `json:"zone-violation-actions"`
    PortOther string 
    ZoneName string 
    Protocol string 

	} `json:"level"`
}


type DdosDstZonePortZoneServiceOtherLevelIndicatorList struct {
    Type string `json:"type"`
    TcpWindowSize int `json:"tcp-window-size"`
    DataPacketSize int `json:"data-packet-size"`
    Score int `json:"score"`
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
    SrcViolationActions string `json:"src-violation-actions"`
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
    ZoneViolationActions string `json:"zone-violation-actions"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZonePortZoneServiceOtherLevelZoneTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
}

func (p *DdosDstZonePortZoneServiceOtherLevel) GetId() string{
    return p.Inst.LevelNum
}

func (p *DdosDstZonePortZoneServiceOtherLevel) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port/zone-service-other/" +p.Inst.PortOther + "+" +p.Inst.Protocol + "/level"
}

func (p *DdosDstZonePortZoneServiceOtherLevel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherLevel::Post")
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

func (p *DdosDstZonePortZoneServiceOtherLevel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherLevel::Get")
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
func (p *DdosDstZonePortZoneServiceOtherLevel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherLevel::Put")
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

func (p *DdosDstZonePortZoneServiceOtherLevel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherLevel::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
