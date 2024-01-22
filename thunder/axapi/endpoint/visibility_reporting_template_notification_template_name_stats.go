

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityReportingTemplateNotificationTemplateNameStats struct {
    
    Name string `json:"name"`
    Stats VisibilityReportingTemplateNotificationTemplateNameStatsStats `json:"stats"`

}
type DataVisibilityReportingTemplateNotificationTemplateNameStats struct {
    DtVisibilityReportingTemplateNotificationTemplateNameStats VisibilityReportingTemplateNotificationTemplateNameStats `json:"template-name"`
}


type VisibilityReportingTemplateNotificationTemplateNameStatsStats struct {
    Sent_successful int `json:"sent_successful"`
    Send_fail int `json:"send_fail"`
    Response_fail int `json:"response_fail"`
}

func (p *VisibilityReportingTemplateNotificationTemplateNameStats) GetId() string{
    return "1"
}

func (p *VisibilityReportingTemplateNotificationTemplateNameStats) getPath() string{
    
    return "visibility/reporting/template/notification/template-name/"+p.Name+"/stats"
}

func (p *VisibilityReportingTemplateNotificationTemplateNameStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityReportingTemplateNotificationTemplateNameStats,error) {
logger.Println("VisibilityReportingTemplateNotificationTemplateNameStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityReportingTemplateNotificationTemplateNameStats
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
