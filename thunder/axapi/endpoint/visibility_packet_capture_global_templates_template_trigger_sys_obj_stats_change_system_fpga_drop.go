

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2093 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2094 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"system-fpga-drop"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2093 struct {
    MrxDrop int `json:"mrx-drop"`
    HrxDrop int `json:"hrx-drop"`
    SizDrop int `json:"siz-drop"`
    FcsDrop int `json:"fcs-drop"`
    LandDrop int `json:"land-drop"`
    EmptyFragDrop int `json:"empty-frag-drop"`
    MicFragDrop int `json:"mic-frag-drop"`
    Ipv4OptDrop int `json:"ipv4-opt-drop"`
    Ipv4Frag int `json:"ipv4-frag"`
    BadIpHdrLen int `json:"bad-ip-hdr-len"`
    BadIpFlagsDrop int `json:"bad-ip-flags-drop"`
    BadIpTtlDrop int `json:"bad-ip-ttl-drop"`
    NoIpPayloadDrop int `json:"no-ip-payload-drop"`
    OversizeIpPayload int `json:"oversize-ip-payload"`
    BadIpPayloadLen int `json:"bad-ip-payload-len"`
    BadIpFragOffset int `json:"bad-ip-frag-offset"`
    BadIpChksumDrop int `json:"bad-ip-chksum-drop"`
    IcmpPodDrop int `json:"icmp-pod-drop"`
    TcpBadUrgOffet int `json:"tcp-bad-urg-offet"`
    TcpShortHdr int `json:"tcp-short-hdr"`
    TcpBadIpLen int `json:"tcp-bad-ip-len"`
    TcpNullFlags int `json:"tcp-null-flags"`
    TcpNullScan int `json:"tcp-null-scan"`
    TcpFinSin int `json:"tcp-fin-sin"`
    TcpXmasFlags int `json:"tcp-xmas-flags"`
    TcpXmasScan int `json:"tcp-xmas-scan"`
    TcpSynFrag int `json:"tcp-syn-frag"`
    TcpFragHdr int `json:"tcp-frag-hdr"`
    TcpBadChksum int `json:"tcp-bad-chksum"`
    UdpShortHdr int `json:"udp-short-hdr"`
    UdpBadIpLen int `json:"udp-bad-ip-len"`
    UdpKbFrags int `json:"udp-kb-frags"`
    UdpPortLb int `json:"udp-port-lb"`
    UdpBadChksum int `json:"udp-bad-chksum"`
    RuntIpHdr int `json:"runt-ip-hdr"`
    RuntTcpudpHdr int `json:"runt-tcpudp-hdr"`
    TunMismatch int `json:"tun-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2094 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MrxDrop int `json:"mrx-drop"`
    HrxDrop int `json:"hrx-drop"`
    SizDrop int `json:"siz-drop"`
    FcsDrop int `json:"fcs-drop"`
    LandDrop int `json:"land-drop"`
    EmptyFragDrop int `json:"empty-frag-drop"`
    MicFragDrop int `json:"mic-frag-drop"`
    Ipv4OptDrop int `json:"ipv4-opt-drop"`
    Ipv4Frag int `json:"ipv4-frag"`
    BadIpHdrLen int `json:"bad-ip-hdr-len"`
    BadIpFlagsDrop int `json:"bad-ip-flags-drop"`
    BadIpTtlDrop int `json:"bad-ip-ttl-drop"`
    NoIpPayloadDrop int `json:"no-ip-payload-drop"`
    OversizeIpPayload int `json:"oversize-ip-payload"`
    BadIpPayloadLen int `json:"bad-ip-payload-len"`
    BadIpFragOffset int `json:"bad-ip-frag-offset"`
    BadIpChksumDrop int `json:"bad-ip-chksum-drop"`
    IcmpPodDrop int `json:"icmp-pod-drop"`
    TcpBadUrgOffet int `json:"tcp-bad-urg-offet"`
    TcpShortHdr int `json:"tcp-short-hdr"`
    TcpBadIpLen int `json:"tcp-bad-ip-len"`
    TcpNullFlags int `json:"tcp-null-flags"`
    TcpNullScan int `json:"tcp-null-scan"`
    TcpFinSin int `json:"tcp-fin-sin"`
    TcpXmasFlags int `json:"tcp-xmas-flags"`
    TcpXmasScan int `json:"tcp-xmas-scan"`
    TcpSynFrag int `json:"tcp-syn-frag"`
    TcpFragHdr int `json:"tcp-frag-hdr"`
    TcpBadChksum int `json:"tcp-bad-chksum"`
    UdpShortHdr int `json:"udp-short-hdr"`
    UdpBadIpLen int `json:"udp-bad-ip-len"`
    UdpKbFrags int `json:"udp-kb-frags"`
    UdpPortLb int `json:"udp-port-lb"`
    UdpBadChksum int `json:"udp-bad-chksum"`
    RuntIpHdr int `json:"runt-ip-hdr"`
    RuntTcpudpHdr int `json:"runt-tcpudp-hdr"`
    TunMismatch int `json:"tun-mismatch"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/system-fpga-drop"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
