

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateQuicVersionSupported struct {
	Inst struct {

    MalformedCheck DdosZoneTemplateQuicVersionSupportedMalformedCheck315 `json:"malformed-check"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VersionAction string `json:"version-action"`
    VersionActionListName string `json:"version-action-list-name"`
    VersionEnd string `json:"version-end"`
    VersionStart string `json:"version-start"`
    QuicTmplName string 

	} `json:"version-supported"`
}


type DdosZoneTemplateQuicVersionSupportedMalformedCheck315 struct {
    MalformedEnable string `json:"malformed-enable" dval:"enable"`
    MaxSourceCidLength int `json:"max-source-cid-length" dval:"255"`
    MaxDestinationCidLength int `json:"max-destination-cid-length" dval:"255"`
    MalformedCheckActionListName string `json:"malformed-check-action-list-name"`
    MalformedCheckAction string `json:"malformed-check-action"`
    Uuid string `json:"uuid"`
}

func (p *DdosZoneTemplateQuicVersionSupported) GetId() string{
    return p.Inst.VersionStart+"+"+p.Inst.VersionEnd
}

func (p *DdosZoneTemplateQuicVersionSupported) getPath() string{
    return "ddos/zone-template/quic/" +p.Inst.QuicTmplName + "/version-supported"
}

func (p *DdosZoneTemplateQuicVersionSupported) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuicVersionSupported::Post")
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

func (p *DdosZoneTemplateQuicVersionSupported) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuicVersionSupported::Get")
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
func (p *DdosZoneTemplateQuicVersionSupported) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuicVersionSupported::Put")
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

func (p *DdosZoneTemplateQuicVersionSupported) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuicVersionSupported::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
