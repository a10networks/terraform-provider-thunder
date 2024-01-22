

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortZoneSrcPortOther struct {
	Inst struct {

    DefaultActionList string `json:"default-action-list"`
    Deny int `json:"deny"`
    GlidCfg DdosDstZoneSrcPortZoneSrcPortOtherGlidCfg `json:"glid-cfg"`
    LevelList []DdosDstZoneSrcPortZoneSrcPortOtherLevelList `json:"level-list"`
    PortInd DdosDstZoneSrcPortZoneSrcPortOtherPortInd242 `json:"port-ind"`
    PortOther string `json:"port-other"`
    Protocol string `json:"protocol"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZoneSrcPortZoneSrcPortOtherZoneTemplate `json:"zone-template"`
    ZoneName string 

	} `json:"zone-src-port-other"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherGlidCfg struct {
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherLevelList struct {
    LevelNum string `json:"level-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IndicatorList []DdosDstZoneSrcPortZoneSrcPortOtherLevelListIndicatorList `json:"indicator-list"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherLevelListIndicatorList struct {
    Type string `json:"type"`
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherPortInd242 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherZoneTemplate struct {
    SrcUdp string `json:"src-udp"`
    SrcTcp string `json:"src-tcp"`
}

func (p *DdosDstZoneSrcPortZoneSrcPortOther) GetId() string{
    return p.Inst.PortOther+"+"+p.Inst.Protocol
}

func (p *DdosDstZoneSrcPortZoneSrcPortOther) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/src-port/zone-src-port-other"
}

func (p *DdosDstZoneSrcPortZoneSrcPortOther) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPortOther::Post")
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

func (p *DdosDstZoneSrcPortZoneSrcPortOther) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPortOther::Get")
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
func (p *DdosDstZoneSrcPortZoneSrcPortOther) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPortOther::Put")
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

func (p *DdosDstZoneSrcPortZoneSrcPortOther) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPortOther::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
