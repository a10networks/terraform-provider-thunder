

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2015 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2016 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"fw-tcp-syn-cookie"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2015 struct {
    Verification_failed int `json:"verification_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2016 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Verification_failed int `json:"verification_failed"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/fw-tcp-syn-cookie"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
