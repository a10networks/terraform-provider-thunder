

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangeManualMode struct {
	Inst struct {

    CloseSessionsForUnauthSources int `json:"close-sessions-for-unauth-sources"`
    Config string `json:"config"`
    GlidAction string `json:"glid-action"`
    SrcDefaultGlid string `json:"src-default-glid"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZonePortRangeManualModeZoneTemplate `json:"zone-template"`
    ZoneName string 
    Protocol string 
    PortRangeStart string 
    PortRangeEnd string 

	} `json:"manual-mode"`
}


type DdosDstZonePortRangeManualModeZoneTemplate struct {
    Quic string `json:"quic"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
}

func (p *DdosDstZonePortRangeManualMode) GetId() string{
    return p.Inst.Config
}

func (p *DdosDstZonePortRangeManualMode) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port-range/" +p.Inst.PortRangeStart + "+" +p.Inst.PortRangeEnd + "+" +p.Inst.Protocol + "/manual-mode"
}

func (p *DdosDstZonePortRangeManualMode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeManualMode::Post")
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

func (p *DdosDstZonePortRangeManualMode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeManualMode::Get")
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
func (p *DdosDstZonePortRangeManualMode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeManualMode::Put")
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

func (p *DdosDstZonePortRangeManualMode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRangeManualMode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
