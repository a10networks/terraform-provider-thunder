

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateQuicVersionSupportedMalformedCheck struct {
	Inst struct {

    MalformedCheckAction string `json:"malformed-check-action"`
    MalformedCheckActionListName string `json:"malformed-check-action-list-name"`
    MalformedEnable string `json:"malformed-enable"`
    MaxDestinationCidLength int `json:"max-destination-cid-length" dval:"255"`
    MaxSourceCidLength int `json:"max-source-cid-length" dval:"255"`
    Uuid string `json:"uuid"`
    VersionStart string 
    VersionEnd string 
    QuicTmplName string 

	} `json:"malformed-check"`
}

func (p *DdosZoneTemplateQuicVersionSupportedMalformedCheck) GetId() string{
    return "1"
}

func (p *DdosZoneTemplateQuicVersionSupportedMalformedCheck) getPath() string{
    return "ddos/zone-template/quic/" +p.Inst.QuicTmplName + "/version-supported/" +p.Inst.VersionStart + "+" +p.Inst.VersionEnd + "/malformed-check"
}

func (p *DdosZoneTemplateQuicVersionSupportedMalformedCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuicVersionSupportedMalformedCheck::Post")
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

func (p *DdosZoneTemplateQuicVersionSupportedMalformedCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuicVersionSupportedMalformedCheck::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *DdosZoneTemplateQuicVersionSupportedMalformedCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuicVersionSupportedMalformedCheck::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *DdosZoneTemplateQuicVersionSupportedMalformedCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuicVersionSupportedMalformedCheck::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
