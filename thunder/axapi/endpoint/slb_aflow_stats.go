

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbAflowStats struct {
    
    Stats SlbAflowStatsStats `json:"stats"`

}
type DataSlbAflowStats struct {
    DtSlbAflowStats SlbAflowStats `json:"aflow"`
}


type SlbAflowStatsStats struct {
    Pause_conn int `json:"pause_conn"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Resume_conn int `json:"resume_conn"`
    Event_resume_conn int `json:"event_resume_conn"`
    Timer_resume_conn int `json:"timer_resume_conn"`
    Try_to_resume_conn int `json:"try_to_resume_conn"`
    Retry_resume_conn int `json:"retry_resume_conn"`
    Error_resume_conn int `json:"error_resume_conn"`
    Open_new_server_conn int `json:"open_new_server_conn"`
    Reuse_server_idle_conn int `json:"reuse_server_idle_conn"`
    Inc_aflow_limit int `json:"inc_aflow_limit"`
}

func (p *SlbAflowStats) GetId() string{
    return "1"
}

func (p *SlbAflowStats) getPath() string{
    return "slb/aflow/stats"
}

func (p *SlbAflowStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbAflowStats,error) {
logger.Println("SlbAflowStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbAflowStats
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
