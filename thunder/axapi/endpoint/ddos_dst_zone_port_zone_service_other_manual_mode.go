

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherManualMode struct {
	Inst struct {

    CloseSessionsForUnauthSources int `json:"close-sessions-for-unauth-sources"`
    Config string `json:"config"`
    GlidAction string `json:"glid-action"`
    SrcDefaultGlid string `json:"src-default-glid"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZonePortZoneServiceOtherManualModeZoneTemplate `json:"zone-template"`
    PortOther string 
    ZoneName string 
    Protocol string 

	} `json:"manual-mode"`
}


type DdosDstZonePortZoneServiceOtherManualModeZoneTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
}

func (p *DdosDstZonePortZoneServiceOtherManualMode) GetId() string{
    return p.Inst.Config
}

func (p *DdosDstZonePortZoneServiceOtherManualMode) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port/zone-service-other/" +p.Inst.PortOther + "+" +p.Inst.Protocol + "/manual-mode"
}

func (p *DdosDstZonePortZoneServiceOtherManualMode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherManualMode::Post")
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

func (p *DdosDstZonePortZoneServiceOtherManualMode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherManualMode::Get")
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
func (p *DdosDstZonePortZoneServiceOtherManualMode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherManualMode::Put")
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

func (p *DdosDstZonePortZoneServiceOtherManualMode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherManualMode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
