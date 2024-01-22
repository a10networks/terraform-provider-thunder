

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemPbslbStats struct {
    
    Stats SystemPbslbStatsStats `json:"stats"`

}
type DataSystemPbslbStats struct {
    DtSystemPbslbStats SystemPbslbStats `json:"pbslb"`
}


type SystemPbslbStatsStats struct {
    Curr_entries int `json:"curr_entries"`
    Total_v4_entries_created int `json:"total_v4_entries_created"`
    Total_v4_entries_freed int `json:"total_v4_entries_freed"`
    Total_v6_entries_created int `json:"total_v6_entries_created"`
    Total_v6_entries_freed int `json:"total_v6_entries_freed"`
    Total_domain_entries_created int `json:"total_domain_entries_created"`
    Total_domain_entries_freed int `json:"total_domain_entries_freed"`
    Total_direct_action_entries_created int `json:"total_direct_action_entries_created"`
    Total_direct_action_entries_freed int `json:"total_direct_action_entries_freed"`
    Curr_entries_target_global int `json:"curr_entries_target_global"`
    Curr_entries_target_vserver int `json:"curr_entries_target_vserver"`
    Curr_entries_target_vport int `json:"curr_entries_target_vport"`
    Curr_entries_target_loc int `json:"curr_entries_target_LOC"`
    Curr_entries_target_rserver int `json:"curr_entries_target_rserver"`
    Curr_entries_target_rport int `json:"curr_entries_target_rport"`
    Curr_entries_target_service int `json:"curr_entries_target_service"`
    Curr_entries_stats int `json:"curr_entries_stats"`
}

func (p *SystemPbslbStats) GetId() string{
    return "1"
}

func (p *SystemPbslbStats) getPath() string{
    return "system/pbslb/stats"
}

func (p *SystemPbslbStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemPbslbStats,error) {
logger.Println("SystemPbslbStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemPbslbStats
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
