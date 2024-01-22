

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneProfileIpProtoProtoName struct {
	Inst struct {

    IndicatorList []DdosZoneProfileIpProtoProtoNameIndicatorList `json:"indicator-list"`
    Protocol string `json:"protocol"`
    Uuid string `json:"uuid"`
    ProfileName string 

	} `json:"proto-name"`
}


type DdosZoneProfileIpProtoProtoNameIndicatorList struct {
    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfileIpProtoProtoNameIndicatorListSrcThresholdCfg `json:"src-threshold-cfg"`
    ZoneThresholdCfg DdosZoneProfileIpProtoProtoNameIndicatorListZoneThresholdCfg `json:"zone-threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneProfileIpProtoProtoNameIndicatorListSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfileIpProtoProtoNameIndicatorListZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}

func (p *DdosZoneProfileIpProtoProtoName) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosZoneProfileIpProtoProtoName) getPath() string{
    return "ddos/zone-profile/" +p.Inst.ProfileName + "/ip-proto/proto-name"
}

func (p *DdosZoneProfileIpProtoProtoName) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfileIpProtoProtoName::Post")
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

func (p *DdosZoneProfileIpProtoProtoName) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfileIpProtoProtoName::Get")
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
func (p *DdosZoneProfileIpProtoProtoName) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfileIpProtoProtoName::Put")
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

func (p *DdosZoneProfileIpProtoProtoName) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfileIpProtoProtoName::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
