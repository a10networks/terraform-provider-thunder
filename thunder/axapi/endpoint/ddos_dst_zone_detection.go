

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneDetection struct {
	Inst struct {

    Notification DdosDstZoneDetectionNotification188 `json:"notification"`
    OutboundDetection DdosDstZoneDetectionOutboundDetection190 `json:"outbound-detection"`
    PacketAnomalyDetection DdosDstZoneDetectionPacketAnomalyDetection194 `json:"packet-anomaly-detection"`
    ServiceDiscovery DdosDstZoneDetectionServiceDiscovery196 `json:"service-discovery"`
    Settings string `json:"settings"`
    Toggle string `json:"toggle" dval:"enable"`
    Uuid string `json:"uuid"`
    VictimIpDetection DdosDstZoneDetectionVictimIpDetection197 `json:"victim-ip-detection"`
    ZoneName string 

	} `json:"detection"`
}


type DdosDstZoneDetectionNotification188 struct {
    Configuration string `json:"configuration"`
    Notification []DdosDstZoneDetectionNotificationNotification189 `json:"notification"`
    Uuid string `json:"uuid"`
}


type DdosDstZoneDetectionNotificationNotification189 struct {
    NotificationTemplateName string `json:"notification-template-name"`
}


type DdosDstZoneDetectionOutboundDetection190 struct {
    Configuration string `json:"configuration"`
    Toggle string `json:"toggle" dval:"disable"`
    DiscoveryMethod string `json:"discovery-method"`
    DiscoveryRecord int `json:"discovery-record" dval:"10"`
    EnableTopK []DdosDstZoneDetectionOutboundDetectionEnableTopK191 `json:"enable-top-k"`
    Uuid string `json:"uuid"`
    IndicatorList []DdosDstZoneDetectionOutboundDetectionIndicatorList192 `json:"indicator-list"`
    TopkSourceSubnet DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet193 `json:"topk-source-subnet"`
}


type DdosDstZoneDetectionOutboundDetectionEnableTopK191 struct {
    TopkType string `json:"topk-type"`
    TopkNetmask int `json:"topk-netmask" dval:"128"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
}


type DdosDstZoneDetectionOutboundDetectionIndicatorList192 struct {
    Type string `json:"type"`
    TcpWindowSize int `json:"tcp-window-size"`
    DataPacketSize int `json:"data-packet-size"`
    ThresholdNum int `json:"threshold-num"`
    ThresholdLargeNum int `json:"threshold-large-num"`
    ThresholdStr string `json:"threshold-str"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet193 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneDetectionPacketAnomalyDetection194 struct {
    Configuration string `json:"configuration"`
    Toggle string `json:"toggle" dval:"enable"`
    Uuid string `json:"uuid"`
    IndicatorList []DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList195 `json:"indicator-list"`
}


type DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList195 struct {
    Type string `json:"type"`
    ThresholdNum int `json:"threshold-num" dval:"100"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneDetectionServiceDiscovery196 struct {
    Configuration string `json:"configuration"`
    Toggle string `json:"toggle" dval:"disable"`
    PktRateThreshold int `json:"pkt-rate-threshold" dval:"10"`
    Uuid string `json:"uuid"`
}


type DdosDstZoneDetectionVictimIpDetection197 struct {
    Configuration string `json:"configuration"`
    Toggle string `json:"toggle" dval:"disable"`
    HistogramToggle string `json:"histogram-toggle" dval:"histogram-disable"`
    Uuid string `json:"uuid"`
    IndicatorList []DdosDstZoneDetectionVictimIpDetectionIndicatorList198 `json:"indicator-list"`
}


type DdosDstZoneDetectionVictimIpDetectionIndicatorList198 struct {
    Type string `json:"type"`
    IpThresholdNum int `json:"ip-threshold-num"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosDstZoneDetection) GetId() string{
    return "1"
}

func (p *DdosDstZoneDetection) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/detection"
}

func (p *DdosDstZoneDetection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetection::Post")
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

func (p *DdosDstZoneDetection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetection::Get")
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
func (p *DdosDstZoneDetection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetection::Put")
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

func (p *DdosDstZoneDetection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
