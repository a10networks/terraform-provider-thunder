

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateSipFilterHeader struct {
	Inst struct {

    SipFilterAction string `json:"sip-filter-action"`
    SipFilterActionListName string `json:"sip-filter-action-list-name"`
    SipFilterHeaderSeq int `json:"sip-filter-header-seq"`
    SipFilterName string `json:"sip-filter-name"`
    SipHeaderCfg DdosZoneTemplateSipFilterHeaderSipHeaderCfg `json:"sip-header-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    SipTmplName string 

	} `json:"filter-header"`
}


type DdosZoneTemplateSipFilterHeaderSipHeaderCfg struct {
    SipFilterHeaderRegex string `json:"sip-filter-header-regex"`
    SipFilterHeaderInverseMatch int `json:"sip-filter-header-inverse-match"`
}

func (p *DdosZoneTemplateSipFilterHeader) GetId() string{
    return url.QueryEscape(p.Inst.SipFilterName)
}

func (p *DdosZoneTemplateSipFilterHeader) getPath() string{
    return "ddos/zone-template/sip/" +p.Inst.SipTmplName + "/filter-header"
}

func (p *DdosZoneTemplateSipFilterHeader) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSipFilterHeader::Post")
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

func (p *DdosZoneTemplateSipFilterHeader) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSipFilterHeader::Get")
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
func (p *DdosZoneTemplateSipFilterHeader) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSipFilterHeader::Put")
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

func (p *DdosZoneTemplateSipFilterHeader) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSipFilterHeader::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
