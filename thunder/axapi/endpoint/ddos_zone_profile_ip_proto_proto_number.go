

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneProfileIpProtoProtoNumber struct {
	Inst struct {

    IndicatorList []DdosZoneProfileIpProtoProtoNumberIndicatorList `json:"indicator-list"`
    ProtocolNum int `json:"protocol-num"`
    Uuid string `json:"uuid"`
    ProfileName string 

	} `json:"proto-number"`
}


type DdosZoneProfileIpProtoProtoNumberIndicatorList struct {
    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfileIpProtoProtoNumberIndicatorListSrcThresholdCfg `json:"src-threshold-cfg"`
    ZoneThresholdCfg DdosZoneProfileIpProtoProtoNumberIndicatorListZoneThresholdCfg `json:"zone-threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneProfileIpProtoProtoNumberIndicatorListSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfileIpProtoProtoNumberIndicatorListZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}

func (p *DdosZoneProfileIpProtoProtoNumber) GetId() string{
    return strconv.Itoa(p.Inst.ProtocolNum)
}

func (p *DdosZoneProfileIpProtoProtoNumber) getPath() string{
    return "ddos/zone-profile/" +p.Inst.ProfileName + "/ip-proto/proto-number"
}

func (p *DdosZoneProfileIpProtoProtoNumber) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfileIpProtoProtoNumber::Post")
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

func (p *DdosZoneProfileIpProtoProtoNumber) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfileIpProtoProtoNumber::Get")
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
func (p *DdosZoneProfileIpProtoProtoNumber) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfileIpProtoProtoNumber::Put")
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

func (p *DdosZoneProfileIpProtoProtoNumber) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfileIpProtoProtoNumber::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
