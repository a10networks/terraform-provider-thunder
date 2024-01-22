

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpMsisdnList struct {
	Inst struct {

    Action string `json:"action" dval:"deny"`
    Name string `json:"name"`
    StrList []TemplateGtpMsisdnListStrList `json:"str-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"msisdn-list"`
}


type TemplateGtpMsisdnListStrList struct {
    Msisdn string `json:"msisdn"`
}

func (p *TemplateGtpMsisdnList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateGtpMsisdnList) getPath() string{
    return "template/gtp/msisdn-list"
}

func (p *TemplateGtpMsisdnList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMsisdnList::Post")
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

func (p *TemplateGtpMsisdnList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMsisdnList::Get")
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
func (p *TemplateGtpMsisdnList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMsisdnList::Put")
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

func (p *TemplateGtpMsisdnList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMsisdnList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
