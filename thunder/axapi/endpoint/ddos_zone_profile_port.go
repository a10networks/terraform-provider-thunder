

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneProfilePort struct {
	Inst struct {

    IndicatorList []DdosZoneProfilePortIndicatorList `json:"indicator-list"`
    PortNum int `json:"port-num"`
    PortProtocol string `json:"port-protocol"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ProfileName string 

	} `json:"port"`
}


type DdosZoneProfilePortIndicatorList struct {
    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfilePortIndicatorListSrcThresholdCfg `json:"src-threshold-cfg"`
    ZoneThresholdCfg DdosZoneProfilePortIndicatorListZoneThresholdCfg `json:"zone-threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneProfilePortIndicatorListSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfilePortIndicatorListZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}

func (p *DdosZoneProfilePort) GetId() string{
    return strconv.Itoa(p.Inst.PortNum)+"+"+p.Inst.PortProtocol
}

func (p *DdosZoneProfilePort) getPath() string{
    return "ddos/zone-profile/" +p.Inst.ProfileName + "/port"
}

func (p *DdosZoneProfilePort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePort::Post")
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

func (p *DdosZoneProfilePort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePort::Get")
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
func (p *DdosZoneProfilePort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePort::Put")
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

func (p *DdosZoneProfilePort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfilePort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
