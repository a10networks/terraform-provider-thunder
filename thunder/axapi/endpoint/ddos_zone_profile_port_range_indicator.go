

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneProfilePortRangeIndicator struct {
	Inst struct {

    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfilePortRangeIndicatorSrcThresholdCfg `json:"src-threshold-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneThresholdCfg DdosZoneProfilePortRangeIndicatorZoneThresholdCfg `json:"zone-threshold-cfg"`
    Protocol string 
    ProfileName string 
    PortRangeStart string 
    PortRangeEnd string 

	} `json:"indicator"`
}


type DdosZoneProfilePortRangeIndicatorSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfilePortRangeIndicatorZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}

func (p *DdosZoneProfilePortRangeIndicator) GetId() string{
    return p.Inst.IndicatorName
}

func (p *DdosZoneProfilePortRangeIndicator) getPath() string{
    return "ddos/zone-profile/" +p.Inst.ProfileName + "/port-range/" +p.Inst.PortRangeStart + "+" +p.Inst.PortRangeEnd + "+" +p.Inst.Protocol + "/indicator"
}

func (p *DdosZoneProfilePortRangeIndicator) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortRangeIndicator::Post")
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

func (p *DdosZoneProfilePortRangeIndicator) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortRangeIndicator::Get")
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
func (p *DdosZoneProfilePortRangeIndicator) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortRangeIndicator::Put")
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

func (p *DdosZoneProfilePortRangeIndicator) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortRangeIndicator::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
