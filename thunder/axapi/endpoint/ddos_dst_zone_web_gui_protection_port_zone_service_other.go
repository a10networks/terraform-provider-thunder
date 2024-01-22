

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneWebGuiProtectionPortZoneServiceOther struct {
	Inst struct {

    Pbe string `json:"pbe"`
    PortOther string `json:"port-other"`
    Protocol string `json:"protocol"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"zone-service-other"`
}

func (p *DdosDstZoneWebGuiProtectionPortZoneServiceOther) GetId() string{
    return p.Inst.PortOther+"+"+p.Inst.Protocol
}

func (p *DdosDstZoneWebGuiProtectionPortZoneServiceOther) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/web-gui/protection/port/zone-service-other"
}

func (p *DdosDstZoneWebGuiProtectionPortZoneServiceOther) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionPortZoneServiceOther::Post")
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

func (p *DdosDstZoneWebGuiProtectionPortZoneServiceOther) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionPortZoneServiceOther::Get")
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
func (p *DdosDstZoneWebGuiProtectionPortZoneServiceOther) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionPortZoneServiceOther::Put")
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

func (p *DdosDstZoneWebGuiProtectionPortZoneServiceOther) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionPortZoneServiceOther::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
