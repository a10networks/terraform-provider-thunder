

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateIcmpV4TypeOther struct {
	Inst struct {

    Dst DdosZoneTemplateIcmpV4TypeOtherDst `json:"dst"`
    IcmpTypeOtherAction string `json:"icmp-type-other-action"`
    IcmpTypeOtherActionListName string `json:"icmp-type-other-action-list-name"`
    Src DdosZoneTemplateIcmpV4TypeOtherSrc `json:"src"`
    Uuid string `json:"uuid"`
    IcmpTmplName string 

	} `json:"type-other"`
}


type DdosZoneTemplateIcmpV4TypeOtherDst struct {
    DstTypeOtherRate int `json:"dst-type-other-rate"`
    DstTypeOtherRateActionListName string `json:"dst-type-other-rate-action-list-name"`
    DstTypeOtherRateAction string `json:"dst-type-other-rate-action"`
}


type DdosZoneTemplateIcmpV4TypeOtherSrc struct {
    SrcTypeOtherRate int `json:"src-type-other-rate"`
    SrcTypeOtherRateActionListName string `json:"src-type-other-rate-action-list-name"`
    SrcTypeOtherRateAction string `json:"src-type-other-rate-action"`
}

func (p *DdosZoneTemplateIcmpV4TypeOther) GetId() string{
    return "1"
}

func (p *DdosZoneTemplateIcmpV4TypeOther) getPath() string{
    return "ddos/zone-template/icmp-v4/" +p.Inst.IcmpTmplName + "/type-other"
}

func (p *DdosZoneTemplateIcmpV4TypeOther) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV4TypeOther::Post")
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

func (p *DdosZoneTemplateIcmpV4TypeOther) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV4TypeOther::Get")
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
func (p *DdosZoneTemplateIcmpV4TypeOther) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV4TypeOther::Put")
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

func (p *DdosZoneTemplateIcmpV4TypeOther) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV4TypeOther::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
