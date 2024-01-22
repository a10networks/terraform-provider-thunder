

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneProfilePortRange struct {
	Inst struct {

    IndicatorList []DdosZoneProfilePortRangeIndicatorList `json:"indicator-list"`
    PortRangeEnd int `json:"port-range-end"`
    PortRangeStart int `json:"port-range-start"`
    Protocol string `json:"protocol"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ProfileName string 

	} `json:"port-range"`
}


type DdosZoneProfilePortRangeIndicatorList struct {
    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfilePortRangeIndicatorListSrcThresholdCfg `json:"src-threshold-cfg"`
    ZoneThresholdCfg DdosZoneProfilePortRangeIndicatorListZoneThresholdCfg `json:"zone-threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneProfilePortRangeIndicatorListSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfilePortRangeIndicatorListZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}

func (p *DdosZoneProfilePortRange) GetId() string{
    return strconv.Itoa(p.Inst.PortRangeStart)+"+"+strconv.Itoa(p.Inst.PortRangeEnd)+"+"+p.Inst.Protocol
}

func (p *DdosZoneProfilePortRange) getPath() string{
    return "ddos/zone-profile/" +p.Inst.ProfileName + "/port-range"
}

func (p *DdosZoneProfilePortRange) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortRange::Post")
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

func (p *DdosZoneProfilePortRange) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortRange::Get")
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
func (p *DdosZoneProfilePortRange) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortRange::Put")
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

func (p *DdosZoneProfilePortRange) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePortRange::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
