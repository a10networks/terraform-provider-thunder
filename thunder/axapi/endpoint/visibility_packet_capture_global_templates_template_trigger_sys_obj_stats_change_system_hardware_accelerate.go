

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2095 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2096 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"system-hardware-accelerate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2095 struct {
    HwFwdProgErrors int `json:"hw-fwd-prog-errors"`
    HwFwdFlowSinglebitErrors int `json:"hw-fwd-flow-singlebit-errors"`
    HwFwdFlowTagMismatch int `json:"hw-fwd-flow-tag-mismatch"`
    HwFwdFlowSeqMismatch int `json:"hw-fwd-flow-seq-mismatch"`
    HwFwdFlowErrorCount int `json:"hw-fwd-flow-error-count"`
    HwFwdFlowUnalignCount int `json:"hw-fwd-flow-unalign-count"`
    HwFwdFlowUnderflowCount int `json:"hw-fwd-flow-underflow-count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2096 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    HwFwdProgErrors int `json:"hw-fwd-prog-errors"`
    HwFwdFlowSinglebitErrors int `json:"hw-fwd-flow-singlebit-errors"`
    HwFwdFlowTagMismatch int `json:"hw-fwd-flow-tag-mismatch"`
    HwFwdFlowSeqMismatch int `json:"hw-fwd-flow-seq-mismatch"`
    HwFwdFlowErrorCount int `json:"hw-fwd-flow-error-count"`
    HwFwdFlowUnalignCount int `json:"hw-fwd-flow-unalign-count"`
    HwFwdFlowUnderflowCount int `json:"hw-fwd-flow-underflow-count"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/system-hardware-accelerate"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
