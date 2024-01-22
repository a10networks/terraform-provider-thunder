

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate struct {
	Inst struct {

    BadReq int `json:"bad-req"`
    Duration int `json:"duration" dval:"60"`
    Forbidden int `json:"forbidden"`
    NoCreds int `json:"no-creds"`
    NotFound int `json:"not-found"`
    ServerError int `json:"server-error"`
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Unauth int `json:"unauth"`
    Unavailable int `json:"unavailable"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate) getPath() string{
    return "visibility/packet-capture/object-templates/aam-auth-relay-hbase-inst-tmpl/" +p.Inst.Name + "/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplTriggerStatsRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
