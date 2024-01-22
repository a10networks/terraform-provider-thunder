

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneProfile struct {
	Inst struct {

    IpProto DdosZoneProfileIpProto305 `json:"ip-proto"`
    PortList []DdosZoneProfilePortList `json:"port-list"`
    PortRangeList []DdosZoneProfilePortRangeList `json:"port-range-list"`
    ProfileName string `json:"profile-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"zone-profile"`
}


type DdosZoneProfileIpProto305 struct {
    ProtoNumberList []DdosZoneProfileIpProtoProtoNumberList `json:"proto-number-list"`
    ProtoNameList []DdosZoneProfileIpProtoProtoNameList `json:"proto-name-list"`
}


type DdosZoneProfileIpProtoProtoNumberList struct {
    ProtocolNum int `json:"protocol-num"`
    Uuid string `json:"uuid"`
    IndicatorList []DdosZoneProfileIpProtoProtoNumberListIndicatorList `json:"indicator-list"`
}


type DdosZoneProfileIpProtoProtoNumberListIndicatorList struct {
    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfileIpProtoProtoNumberListIndicatorListSrcThresholdCfg `json:"src-threshold-cfg"`
    ZoneThresholdCfg DdosZoneProfileIpProtoProtoNumberListIndicatorListZoneThresholdCfg `json:"zone-threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneProfileIpProtoProtoNumberListIndicatorListSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfileIpProtoProtoNumberListIndicatorListZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}


type DdosZoneProfileIpProtoProtoNameList struct {
    Protocol string `json:"protocol"`
    Uuid string `json:"uuid"`
    IndicatorList []DdosZoneProfileIpProtoProtoNameListIndicatorList `json:"indicator-list"`
}


type DdosZoneProfileIpProtoProtoNameListIndicatorList struct {
    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfileIpProtoProtoNameListIndicatorListSrcThresholdCfg `json:"src-threshold-cfg"`
    ZoneThresholdCfg DdosZoneProfileIpProtoProtoNameListIndicatorListZoneThresholdCfg `json:"zone-threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneProfileIpProtoProtoNameListIndicatorListSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfileIpProtoProtoNameListIndicatorListZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}


type DdosZoneProfilePortList struct {
    PortNum int `json:"port-num"`
    PortProtocol string `json:"port-protocol"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IndicatorList []DdosZoneProfilePortListIndicatorList `json:"indicator-list"`
}


type DdosZoneProfilePortListIndicatorList struct {
    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfilePortListIndicatorListSrcThresholdCfg `json:"src-threshold-cfg"`
    ZoneThresholdCfg DdosZoneProfilePortListIndicatorListZoneThresholdCfg `json:"zone-threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneProfilePortListIndicatorListSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfilePortListIndicatorListZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}


type DdosZoneProfilePortRangeList struct {
    PortRangeStart int `json:"port-range-start"`
    PortRangeEnd int `json:"port-range-end"`
    Protocol string `json:"protocol"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IndicatorList []DdosZoneProfilePortRangeListIndicatorList `json:"indicator-list"`
}


type DdosZoneProfilePortRangeListIndicatorList struct {
    IndicatorName string `json:"indicator-name"`
    SrcThresholdCfg DdosZoneProfilePortRangeListIndicatorListSrcThresholdCfg `json:"src-threshold-cfg"`
    ZoneThresholdCfg DdosZoneProfilePortRangeListIndicatorListZoneThresholdCfg `json:"zone-threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneProfilePortRangeListIndicatorListSrcThresholdCfg struct {
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
}


type DdosZoneProfilePortRangeListIndicatorListZoneThresholdCfg struct {
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
}

func (p *DdosZoneProfile) GetId() string{
    return url.QueryEscape(p.Inst.ProfileName)
}

func (p *DdosZoneProfile) getPath() string{
    return "ddos/zone-profile"
}

func (p *DdosZoneProfile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfile::Post")
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

func (p *DdosZoneProfile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfile::Get")
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
func (p *DdosZoneProfile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfile::Put")
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

func (p *DdosZoneProfile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneProfile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
