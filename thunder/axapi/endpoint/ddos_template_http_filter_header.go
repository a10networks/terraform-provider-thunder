

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateHttpFilterHeader struct {
	Inst struct {

    HttpFilterHeaderBlacklist int `json:"http-filter-header-blacklist"`
    HttpFilterHeaderCountOnly int `json:"http-filter-header-count-only"`
    HttpFilterHeaderRegex string `json:"http-filter-header-regex"`
    HttpFilterHeaderSeq int `json:"http-filter-header-seq"`
    HttpFilterHeaderUnmatched int `json:"http-filter-header-unmatched"`
    HttpFilterHeaderWhitelist int `json:"http-filter-header-whitelist"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    HttpTmplName string 

	} `json:"filter-header"`
}

func (p *DdosTemplateHttpFilterHeader) GetId() string{
    return strconv.Itoa(p.Inst.HttpFilterHeaderSeq)
}

func (p *DdosTemplateHttpFilterHeader) getPath() string{
    return "ddos/template/http/" +p.Inst.HttpTmplName + "/filter-header"
}

func (p *DdosTemplateHttpFilterHeader) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateHttpFilterHeader::Post")
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

func (p *DdosTemplateHttpFilterHeader) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateHttpFilterHeader::Get")
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
func (p *DdosTemplateHttpFilterHeader) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateHttpFilterHeader::Put")
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

func (p *DdosTemplateHttpFilterHeader) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateHttpFilterHeader::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
