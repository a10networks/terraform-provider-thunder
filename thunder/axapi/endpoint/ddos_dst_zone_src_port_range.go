

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortRange struct {
	Inst struct {

    CaptureConfig DdosDstZoneSrcPortRangeCaptureConfig `json:"capture-config"`
    DefaultActionList string `json:"default-action-list"`
    Deny int `json:"deny"`
    GlidCfg DdosDstZoneSrcPortRangeGlidCfg `json:"glid-cfg"`
    LevelList []DdosDstZoneSrcPortRangeLevelList `json:"level-list"`
    PortInd DdosDstZoneSrcPortRangePortInd241 `json:"port-ind"`
    Protocol string `json:"protocol"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    SrcPortRangeEnd int `json:"src-port-range-end"`
    SrcPortRangeStart int `json:"src-port-range-start"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZoneSrcPortRangeZoneTemplate `json:"zone-template"`
    ZoneName string 

	} `json:"src-port-range"`
}


type DdosDstZoneSrcPortRangeCaptureConfig struct {
    CaptureConfigName string `json:"capture-config-name"`
    CaptureConfigMode string `json:"capture-config-mode"`
}


type DdosDstZoneSrcPortRangeGlidCfg struct {
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
}


type DdosDstZoneSrcPortRangeLevelList struct {
    LevelNum string `json:"level-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IndicatorList []DdosDstZoneSrcPortRangeLevelListIndicatorList `json:"indicator-list"`
}


type DdosDstZoneSrcPortRangeLevelListIndicatorList struct {
    Type string `json:"type"`
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneSrcPortRangePortInd241 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneSrcPortRangeZoneTemplate struct {
    SrcUdp string `json:"src-udp"`
    SrcTcp string `json:"src-tcp"`
}

func (p *DdosDstZoneSrcPortRange) GetId() string{
    return strconv.Itoa(p.Inst.SrcPortRangeStart)+"+"+strconv.Itoa(p.Inst.SrcPortRangeEnd)+"+"+p.Inst.Protocol
}

func (p *DdosDstZoneSrcPortRange) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/src-port-range"
}

func (p *DdosDstZoneSrcPortRange) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRange::Post")
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

func (p *DdosDstZoneSrcPortRange) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRange::Get")
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
func (p *DdosDstZoneSrcPortRange) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRange::Put")
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

func (p *DdosDstZoneSrcPortRange) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortRange::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
