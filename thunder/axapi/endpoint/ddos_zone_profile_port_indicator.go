

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneProfilePortIndicator struct {
	Inst struct {

    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfilePortIndicatorSrcThresholdCfg `json:"src-threshold-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneThresholdCfg DdosZoneProfilePortIndicatorZoneThresholdCfg `json:"zone-threshold-cfg"`
    PortNum string 
    PortProtocol string 
    ProfileName string 

	} `json:"indicator"`
}


type DdosZoneProfilePortIndicatorSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfilePortIndicatorZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}

func (p *DdosZoneProfilePortIndicator) GetId() string{
    return p.Inst.IndicatorName
}

func (p *DdosZoneProfilePortIndicator) getPath() string{
    return "ddos/zone-profile/" +p.Inst.ProfileName + "/port/" +p.Inst.PortNum + "+" +p.Inst.PortProtocol + "/indicator"
}

func (p *DdosZoneProfilePortIndicator) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortIndicator::Post")
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

func (p *DdosZoneProfilePortIndicator) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortIndicator::Get")
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
func (p *DdosZoneProfilePortIndicator) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortIndicator::Put")
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

func (p *DdosZoneProfilePortIndicator) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortIndicator::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
