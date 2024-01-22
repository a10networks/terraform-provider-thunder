

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneWebGuiProtection struct {
	Inst struct {

    IpProto DdosDstZoneWebGuiProtectionIpProto244 `json:"ip-proto"`
    Port DdosDstZoneWebGuiProtectionPort245 `json:"port"`
    PortRangeList []DdosDstZoneWebGuiProtectionPortRangeList `json:"port-range-list"`
    ZoneName string 

	} `json:"protection"`
}


type DdosDstZoneWebGuiProtectionIpProto244 struct {
    ProtoNameList []DdosDstZoneWebGuiProtectionIpProtoProtoNameList `json:"proto-name-list"`
}


type DdosDstZoneWebGuiProtectionIpProtoProtoNameList struct {
    Protocol string `json:"protocol"`
    Pbe string `json:"pbe"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneWebGuiProtectionPort245 struct {
    ZoneServiceList []DdosDstZoneWebGuiProtectionPortZoneServiceList `json:"zone-service-list"`
    ZoneServiceOtherList []DdosDstZoneWebGuiProtectionPortZoneServiceOtherList `json:"zone-service-other-list"`
}


type DdosDstZoneWebGuiProtectionPortZoneServiceList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Pbe string `json:"pbe"`
    Uuid string `json:"uuid"`
}


type DdosDstZoneWebGuiProtectionPortZoneServiceOtherList struct {
    PortOther string `json:"port-other"`
    Protocol string `json:"protocol"`
    Pbe string `json:"pbe"`
    Uuid string `json:"uuid"`
}


type DdosDstZoneWebGuiProtectionPortRangeList struct {
    PortRangeStart int `json:"port-range-start"`
    PortRangeEnd int `json:"port-range-end"`
    Protocol string `json:"protocol"`
    Pbe string `json:"pbe"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosDstZoneWebGuiProtection) GetId() string{
    return "1"
}

func (p *DdosDstZoneWebGuiProtection) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/web-gui/protection"
}

func (p *DdosDstZoneWebGuiProtection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtection::Post")
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

func (p *DdosDstZoneWebGuiProtection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtection::Get")
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
func (p *DdosDstZoneWebGuiProtection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtection::Put")
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

func (p *DdosDstZoneWebGuiProtection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
