

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc struct {
	Inst struct {

    BadIpChksumDrop int `json:"bad-ip-chksum-drop"`
    BadIpFlagsDrop int `json:"bad-ip-flags-drop"`
    BadIpFragOffset int `json:"bad-ip-frag-offset"`
    BadIpHdrLen int `json:"bad-ip-hdr-len"`
    BadIpPayloadLen int `json:"bad-ip-payload-len"`
    BadIpTtlDrop int `json:"bad-ip-ttl-drop"`
    EmptyFragDrop int `json:"empty-frag-drop"`
    FcsDrop int `json:"fcs-drop"`
    HrxDrop int `json:"hrx-drop"`
    IcmpPodDrop int `json:"icmp-pod-drop"`
    Ipv4Frag int `json:"ipv4-frag"`
    Ipv4OptDrop int `json:"ipv4-opt-drop"`
    LandDrop int `json:"land-drop"`
    MicFragDrop int `json:"mic-frag-drop"`
    MrxDrop int `json:"mrx-drop"`
    NoIpPayloadDrop int `json:"no-ip-payload-drop"`
    OversizeIpPayload int `json:"oversize-ip-payload"`
    RuntIpHdr int `json:"runt-ip-hdr"`
    RuntTcpudpHdr int `json:"runt-tcpudp-hdr"`
    SizDrop int `json:"siz-drop"`
    TcpBadChksum int `json:"tcp-bad-chksum"`
    TcpBadIpLen int `json:"tcp-bad-ip-len"`
    TcpBadUrgOffet int `json:"tcp-bad-urg-offet"`
    TcpFinSin int `json:"tcp-fin-sin"`
    TcpFragHdr int `json:"tcp-frag-hdr"`
    TcpNullFlags int `json:"tcp-null-flags"`
    TcpNullScan int `json:"tcp-null-scan"`
    TcpShortHdr int `json:"tcp-short-hdr"`
    TcpSynFrag int `json:"tcp-syn-frag"`
    TcpXmasFlags int `json:"tcp-xmas-flags"`
    TcpXmasScan int `json:"tcp-xmas-scan"`
    TunMismatch int `json:"tun-mismatch"`
    UdpBadChksum int `json:"udp-bad-chksum"`
    UdpBadIpLen int `json:"udp-bad-ip-len"`
    UdpKbFrags int `json:"udp-kb-frags"`
    UdpPortLb int `json:"udp-port-lb"`
    UdpShortHdr int `json:"udp-short-hdr"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/system-fpga-drop/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
