

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateIcmpV6Filter struct {
	Inst struct {

    ByteOffsetFilter string `json:"byte-offset-filter"`
    IcmpFilterAction string `json:"icmp-filter-action" dval:"drop"`
    IcmpFilterActionListName string `json:"icmp-filter-action-list-name"`
    IcmpFilterInverseMatch int `json:"icmp-filter-inverse-match"`
    IcmpFilterName string `json:"icmp-filter-name"`
    IcmpFilterRegex string `json:"icmp-filter-regex"`
    IcmpFilterSeq int `json:"icmp-filter-seq"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    IcmpTmplName string 

	} `json:"filter"`
}

func (p *DdosZoneTemplateIcmpV6Filter) GetId() string{
    return url.QueryEscape(p.Inst.IcmpFilterName)
}

func (p *DdosZoneTemplateIcmpV6Filter) getPath() string{
    return "ddos/zone-template/icmp-v6/" +p.Inst.IcmpTmplName + "/filter"
}

func (p *DdosZoneTemplateIcmpV6Filter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6Filter::Post")
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

func (p *DdosZoneTemplateIcmpV6Filter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6Filter::Get")
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
func (p *DdosZoneTemplateIcmpV6Filter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6Filter::Put")
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

func (p *DdosZoneTemplateIcmpV6Filter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6Filter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
