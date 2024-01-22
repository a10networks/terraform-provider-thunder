

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneWebGuiProtectionIpProtoProtoName struct {
	Inst struct {

    Pbe string `json:"pbe"`
    Protocol string `json:"protocol"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"proto-name"`
}

func (p *DdosDstZoneWebGuiProtectionIpProtoProtoName) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstZoneWebGuiProtectionIpProtoProtoName) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/web-gui/protection/ip-proto/proto-name"
}

func (p *DdosDstZoneWebGuiProtectionIpProtoProtoName) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionIpProtoProtoName::Post")
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

func (p *DdosDstZoneWebGuiProtectionIpProtoProtoName) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionIpProtoProtoName::Get")
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
func (p *DdosDstZoneWebGuiProtectionIpProtoProtoName) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionIpProtoProtoName::Put")
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

func (p *DdosDstZoneWebGuiProtectionIpProtoProtoName) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionIpProtoProtoName::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
