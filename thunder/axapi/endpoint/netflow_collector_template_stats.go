

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowCollectorTemplateStats struct {
    
    Stats NetflowCollectorTemplateStatsStats `json:"stats"`

}
type DataNetflowCollectorTemplateStats struct {
    DtNetflowCollectorTemplateStats NetflowCollectorTemplateStats `json:"template"`
}


type NetflowCollectorTemplateStatsStats struct {
    V9TemplatesCreated int `json:"v9-templates-created"`
    V9TemplatesDeleted int `json:"v9-templates-deleted"`
    V10TemplatesCreated int `json:"v10-templates-created"`
    V10TemplatesDeleted int `json:"v10-templates-deleted"`
    TemplateDropExceeded int `json:"template-drop-exceeded"`
    TemplateDropOutOfMemory int `json:"template-drop-out-of-memory"`
}

func (p *NetflowCollectorTemplateStats) GetId() string{
    return "1"
}

func (p *NetflowCollectorTemplateStats) getPath() string{
    return "netflow/collector/template/stats"
}

func (p *NetflowCollectorTemplateStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetflowCollectorTemplateStats,error) {
logger.Println("NetflowCollectorTemplateStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetflowCollectorTemplateStats
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
