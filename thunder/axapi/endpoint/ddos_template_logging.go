

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateLogging struct {
	Inst struct {

    EnableActionLogging int `json:"enable-action-logging"`
    LogFormatCef int `json:"log-format-cef"`
    LogFormatCustom string `json:"log-format-custom"`
    LoggingTmplName string `json:"logging-tmpl-name"`
    UseObjName int `json:"use-obj-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"logging"`
}

func (p *DdosTemplateLogging) GetId() string{
    return url.QueryEscape(p.Inst.LoggingTmplName)
}

func (p *DdosTemplateLogging) getPath() string{
    return "ddos/template/logging"
}

func (p *DdosTemplateLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateLogging::Post")
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

func (p *DdosTemplateLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateLogging::Get")
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
func (p *DdosTemplateLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateLogging::Put")
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

func (p *DdosTemplateLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateLogging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
