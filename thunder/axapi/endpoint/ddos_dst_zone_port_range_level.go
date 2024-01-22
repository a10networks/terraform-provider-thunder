

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangeLevel struct {
	Inst struct {

    ApplyExtractedFilters int `json:"apply-extracted-filters"`
    CloseSessionsForUnauthSources int `json:"close-sessions-for-unauth-sources"`
    GlidAction string `json:"glid-action"`
    IndicatorList []DdosDstZonePortRangeLevelIndicatorList `json:"indicator-list"`
    LevelNum string `json:"level-num"`
    SrcDefaultGlid string `json:"src-default-glid"`
    SrcEscalationScore int `json:"src-escalation-score"`
    SrcViolationActions string `json:"src-violation-actions"`
    StartPatternRecognition int `json:"start-pattern-recognition"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneEscalationScore int `json:"zone-escalation-score"`
    ZoneTemplate DdosDstZonePortRangeLevelZoneTemplate `json:"zone-template"`
    ZoneViolationActions string `json:"zone-violation-actions"`
    ZoneName string 
    Protocol string 
    PortRangeStart string 
    PortRangeEnd string 

	} `json:"level"`
}


type DdosDstZonePortRangeLevelIndicatorList struct {
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


type DdosDstZonePortRangeLevelZoneTemplate struct {
    Quic string `json:"quic"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
}

func (p *DdosDstZonePortRangeLevel) GetId() string{
    return p.Inst.LevelNum
}

func (p *DdosDstZonePortRangeLevel) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port-range/" +p.Inst.PortRangeStart + "+" +p.Inst.PortRangeEnd + "+" +p.Inst.Protocol + "/level"
}

func (p *DdosDstZonePortRangeLevel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeLevel::Post")
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

func (p *DdosDstZonePortRangeLevel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeLevel::Get")
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
func (p *DdosDstZonePortRangeLevel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeLevel::Put")
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

func (p *DdosDstZonePortRangeLevel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeLevel::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
