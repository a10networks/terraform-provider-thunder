

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpApnImsiList struct {
	Inst struct {

    Action string `json:"action" dval:"deny"`
    Name string `json:"name"`
    StrList []TemplateGtpApnImsiListStrList `json:"str-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"apn-imsi-list"`
}


type TemplateGtpApnImsiListStrList struct {
    Apn string `json:"apn"`
    SelectionMode string `json:"selection-mode"`
    ImsiSelection string `json:"imsi-selection"`
    Imsi string `json:"imsi"`
}

func (p *TemplateGtpApnImsiList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateGtpApnImsiList) getPath() string{
    return "template/gtp/apn-imsi-list"
}

func (p *TemplateGtpApnImsiList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpApnImsiList::Post")
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

func (p *TemplateGtpApnImsiList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpApnImsiList::Get")
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
func (p *TemplateGtpApnImsiList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpApnImsiList::Put")
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

func (p *TemplateGtpApnImsiList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpApnImsiList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
