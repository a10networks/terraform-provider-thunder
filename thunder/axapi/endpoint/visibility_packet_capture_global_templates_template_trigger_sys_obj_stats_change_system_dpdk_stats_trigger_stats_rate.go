

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate struct {
	Inst struct {

    Duration int `json:"duration" dval:"60"`
    ErrPktDrop int `json:"err-pkt-drop"`
    IoRingDrop int `json:"io-ring-drop"`
    IoRxQueDrop int `json:"io-rx-que-drop"`
    IoTxQueDrop int `json:"io-tx-que-drop"`
    PktDrop int `json:"pkt-drop"`
    PktLnkDownDrop int `json:"pkt-lnk-down-drop"`
    RxAlignErr int `json:"rx-align-err"`
    RxCrcErr int `json:"rx-crc-err"`
    RxCsumOffloadErr int `json:"rx-csum-offload-err"`
    RxErr int `json:"rx-err"`
    RxFrameErr int `json:"rx-frame-err"`
    RxLenErr int `json:"rx-len-err"`
    RxLongLenErr int `json:"rx-long-len-err"`
    RxMissErr int `json:"rx-miss-err"`
    RxNoBuffErr int `json:"rx-no-buff-err"`
    RxOverErr int `json:"rx-over-err"`
    RxShortLenErr int `json:"rx-short-len-err"`
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    TxAbortErr int `json:"tx-abort-err"`
    TxCarrierErr int `json:"tx-carrier-err"`
    TxDrop int `json:"tx-drop"`
    TxErr int `json:"tx-err"`
    TxFifoErr int `json:"tx-fifo-err"`
    TxHbeatErr int `json:"tx-hbeat-err"`
    TxWindowsErr int `json:"tx-windows-err"`
    Uuid string `json:"uuid"`
    WLinkDownDrop int `json:"w-link-down-drop"`
    WRingDrop int `json:"w-ring-drop"`
    WTxQueDrop int `json:"w-tx-que-drop"`
    Name string 

	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/system-dpdk-stats/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
