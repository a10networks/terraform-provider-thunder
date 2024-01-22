

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneDetectionOutboundDetection struct {
	Inst struct {

    Configuration string `json:"configuration"`
    DiscoveryMethod string `json:"discovery-method"`
    DiscoveryRecord int `json:"discovery-record" dval:"10"`
    EnableTopK []DdosDstZoneDetectionOutboundDetectionEnableTopK `json:"enable-top-k"`
    IndicatorList []DdosDstZoneDetectionOutboundDetectionIndicatorList `json:"indicator-list"`
    Toggle string `json:"toggle" dval:"disable"`
    TopkSourceSubnet DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet187 `json:"topk-source-subnet"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"outbound-detection"`
}


type DdosDstZoneDetectionOutboundDetectionEnableTopK struct {
    TopkType string `json:"topk-type"`
    TopkNetmask int `json:"topk-netmask" dval:"128"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
}


type DdosDstZoneDetectionOutboundDetectionIndicatorList struct {
    Type string `json:"type"`
    TcpWindowSize int `json:"tcp-window-size"`
    DataPacketSize int `json:"data-packet-size"`
    ThresholdNum int `json:"threshold-num"`
    ThresholdLargeNum int `json:"threshold-large-num"`
    ThresholdStr string `json:"threshold-str"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet187 struct {
    Uuid string `json:"uuid"`
}

func (p *DdosDstZoneDetectionOutboundDetection) GetId() string{
    return "1"
}

func (p *DdosDstZoneDetectionOutboundDetection) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/detection/outbound-detection"
}

func (p *DdosDstZoneDetectionOutboundDetection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionOutboundDetection::Post")
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

func (p *DdosDstZoneDetectionOutboundDetection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionOutboundDetection::Get")
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
func (p *DdosDstZoneDetectionOutboundDetection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionOutboundDetection::Put")
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

func (p *DdosDstZoneDetectionOutboundDetection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneDetectionOutboundDetection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
