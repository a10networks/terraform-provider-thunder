

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosNotificationTemplateDebugLogOper struct {
    
    Oper DdosNotificationTemplateDebugLogOperOper `json:"oper"`

}
type DataDdosNotificationTemplateDebugLogOper struct {
    DtDdosNotificationTemplateDebugLogOper DdosNotificationTemplateDebugLogOper `json:"notification-template-debug-log"`
}


type DdosNotificationTemplateDebugLogOperOper struct {
    NotificationTemplateDebugList []DdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList `json:"notification-template-debug-list"`
}


type DdosNotificationTemplateDebugLogOperOperNotificationTemplateDebugList struct {
    NotificationTemplateDebugLog string `json:"notification-template-debug-log"`
}

func (p *DdosNotificationTemplateDebugLogOper) GetId() string{
    return "1"
}

func (p *DdosNotificationTemplateDebugLogOper) getPath() string{
    return "ddos/notification-template-debug-log/oper"
}

func (p *DdosNotificationTemplateDebugLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosNotificationTemplateDebugLogOper,error) {
logger.Println("DdosNotificationTemplateDebugLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosNotificationTemplateDebugLogOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
