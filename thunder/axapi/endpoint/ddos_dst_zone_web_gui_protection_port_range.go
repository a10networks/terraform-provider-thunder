

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneWebGuiProtectionPortRange struct {
	Inst struct {

    Pbe string `json:"pbe"`
    PortRangeEnd int `json:"port-range-end"`
    PortRangeStart int `json:"port-range-start"`
    Protocol string `json:"protocol"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"port-range"`
}

func (p *DdosDstZoneWebGuiProtectionPortRange) GetId() string{
    return strconv.Itoa(p.Inst.PortRangeStart)+"+"+strconv.Itoa(p.Inst.PortRangeEnd)+"+"+p.Inst.Protocol
}

func (p *DdosDstZoneWebGuiProtectionPortRange) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/web-gui/protection/port-range"
}

func (p *DdosDstZoneWebGuiProtectionPortRange) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionPortRange::Post")
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

func (p *DdosDstZoneWebGuiProtectionPortRange) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionPortRange::Get")
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
func (p *DdosDstZoneWebGuiProtectionPortRange) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionPortRange::Put")
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

func (p *DdosDstZoneWebGuiProtectionPortRange) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneWebGuiProtectionPortRange::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
