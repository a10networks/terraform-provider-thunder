

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityReportingTemplateNotificationDebugOper struct {
    
    Oper VisibilityReportingTemplateNotificationDebugOperOper `json:"oper"`

}
type DataVisibilityReportingTemplateNotificationDebugOper struct {
    DtVisibilityReportingTemplateNotificationDebugOper VisibilityReportingTemplateNotificationDebugOper `json:"debug"`
}


type VisibilityReportingTemplateNotificationDebugOperOper struct {
    DebugList []VisibilityReportingTemplateNotificationDebugOperOperDebugList `json:"debug-list"`
    Log int `json:"log"`
}


type VisibilityReportingTemplateNotificationDebugOperOperDebugList struct {
    DebugLog string `json:"debug-log"`
}

func (p *VisibilityReportingTemplateNotificationDebugOper) GetId() string{
    return "1"
}

func (p *VisibilityReportingTemplateNotificationDebugOper) getPath() string{
    return "visibility/reporting/template/notification/debug/oper"
}

func (p *VisibilityReportingTemplateNotificationDebugOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityReportingTemplateNotificationDebugOper,error) {
logger.Println("VisibilityReportingTemplateNotificationDebugOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityReportingTemplateNotificationDebugOper
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
