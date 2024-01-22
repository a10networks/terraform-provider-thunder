

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosNotificationTemplateCommon struct {
	Inst struct {

    DefaultTemplate []DdosNotificationTemplateCommonDefaultTemplate `json:"default-template"`
    OnBoxGuiNotification string `json:"on-box-gui-notification" dval:"enable"`
    Uuid string `json:"uuid"`

	} `json:"notification-template-common"`
}


type DdosNotificationTemplateCommonDefaultTemplate struct {
    DefaultNotificationTemplate string `json:"default-notification-template"`
}

func (p *DdosNotificationTemplateCommon) GetId() string{
    return "1"
}

func (p *DdosNotificationTemplateCommon) getPath() string{
    return "ddos/notification-template-common"
}

func (p *DdosNotificationTemplateCommon) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplateCommon::Post")
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

func (p *DdosNotificationTemplateCommon) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplateCommon::Get")
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
func (p *DdosNotificationTemplateCommon) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplateCommon::Put")
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

func (p *DdosNotificationTemplateCommon) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplateCommon::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
