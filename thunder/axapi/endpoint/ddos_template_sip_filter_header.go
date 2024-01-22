

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateSipFilterHeader struct {
	Inst struct {

    SipFilterHeaderBlacklist int `json:"sip-filter-header-blacklist"`
    SipFilterHeaderCountOnly int `json:"sip-filter-header-count-only"`
    SipFilterHeaderRegex string `json:"sip-filter-header-regex"`
    SipFilterHeaderSeq int `json:"sip-filter-header-seq"`
    SipFilterHeaderUnmatched int `json:"sip-filter-header-unmatched"`
    SipFilterHeaderWhitelist int `json:"sip-filter-header-whitelist"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    SipTmplName string 

	} `json:"filter-header"`
}

func (p *DdosTemplateSipFilterHeader) GetId() string{
    return strconv.Itoa(p.Inst.SipFilterHeaderSeq)
}

func (p *DdosTemplateSipFilterHeader) getPath() string{
    return "ddos/template/sip/" +p.Inst.SipTmplName + "/filter-header"
}

func (p *DdosTemplateSipFilterHeader) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSipFilterHeader::Post")
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

func (p *DdosTemplateSipFilterHeader) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSipFilterHeader::Get")
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
func (p *DdosTemplateSipFilterHeader) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSipFilterHeader::Put")
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

func (p *DdosTemplateSipFilterHeader) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSipFilterHeader::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
