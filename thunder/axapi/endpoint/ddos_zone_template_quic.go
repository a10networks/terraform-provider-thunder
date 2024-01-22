

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateQuic struct {
	Inst struct {

    FixedBitCheckDisable int `json:"fixed-bit-check-disable"`
    QuicTmplName string `json:"quic-tmpl-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VersionSupportedList []DdosZoneTemplateQuicVersionSupportedList `json:"version-supported-list"`

	} `json:"quic"`
}


type DdosZoneTemplateQuicVersionSupportedList struct {
    VersionStart string `json:"version-start"`
    VersionEnd string `json:"version-end"`
    VersionActionListName string `json:"version-action-list-name"`
    VersionAction string `json:"version-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    MalformedCheck DdosZoneTemplateQuicVersionSupportedListMalformedCheck `json:"malformed-check"`
}


type DdosZoneTemplateQuicVersionSupportedListMalformedCheck struct {
    MalformedEnable string `json:"malformed-enable" dval:"enable"`
    MaxSourceCidLength int `json:"max-source-cid-length" dval:"255"`
    MaxDestinationCidLength int `json:"max-destination-cid-length" dval:"255"`
    MalformedCheckActionListName string `json:"malformed-check-action-list-name"`
    MalformedCheckAction string `json:"malformed-check-action"`
    Uuid string `json:"uuid"`
}

func (p *DdosZoneTemplateQuic) GetId() string{
    return url.QueryEscape(p.Inst.QuicTmplName)
}

func (p *DdosZoneTemplateQuic) getPath() string{
    return "ddos/zone-template/quic"
}

func (p *DdosZoneTemplateQuic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuic::Post")
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

func (p *DdosZoneTemplateQuic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuic::Get")
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
func (p *DdosZoneTemplateQuic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuic::Put")
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

func (p *DdosZoneTemplateQuic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateQuic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
