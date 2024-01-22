

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortZoneSrcPort struct {
	Inst struct {

    DefaultActionList string `json:"default-action-list"`
    Deny int `json:"deny"`
    GlidCfg DdosDstZoneSrcPortZoneSrcPortGlidCfg `json:"glid-cfg"`
    LevelList []DdosDstZoneSrcPortZoneSrcPortLevelList `json:"level-list"`
    OutboundSrcTracking string `json:"outbound-src-tracking" dval:"disable"`
    PortInd DdosDstZoneSrcPortZoneSrcPortPortInd243 `json:"port-ind"`
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZoneSrcPortZoneSrcPortZoneTemplate `json:"zone-template"`
    ZoneName string 

	} `json:"zone-src-port"`
}


type DdosDstZoneSrcPortZoneSrcPortGlidCfg struct {
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
}


type DdosDstZoneSrcPortZoneSrcPortLevelList struct {
    LevelNum string `json:"level-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IndicatorList []DdosDstZoneSrcPortZoneSrcPortLevelListIndicatorList `json:"indicator-list"`
}


type DdosDstZoneSrcPortZoneSrcPortLevelListIndicatorList struct {
    Type string `json:"type"`
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneSrcPortZoneSrcPortPortInd243 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneSrcPortZoneSrcPortZoneTemplate struct {
    SrcUdp string `json:"src-udp"`
    SrcTcp string `json:"src-tcp"`
    SrcDns string `json:"src-dns"`
}

func (p *DdosDstZoneSrcPortZoneSrcPort) GetId() string{
    return strconv.Itoa(p.Inst.PortNum)+"+"+p.Inst.Protocol
}

func (p *DdosDstZoneSrcPortZoneSrcPort) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/src-port/zone-src-port"
}

func (p *DdosDstZoneSrcPortZoneSrcPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPort::Post")
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

func (p *DdosDstZoneSrcPortZoneSrcPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPort::Get")
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
func (p *DdosDstZoneSrcPortZoneSrcPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPort::Put")
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

func (p *DdosDstZoneSrcPortZoneSrcPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneSrcPortZoneSrcPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
