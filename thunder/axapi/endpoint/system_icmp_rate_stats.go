

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIcmpRateStats struct {
    
    Stats SystemIcmpRateStatsStats `json:"stats"`

}
type DataSystemIcmpRateStats struct {
    DtSystemIcmpRateStats SystemIcmpRateStats `json:"icmp-rate"`
}


type SystemIcmpRateStatsStats struct {
    Over_limit_drop int `json:"over_limit_drop"`
    Limit_intf_drop int `json:"limit_intf_drop"`
    Limit_vserver_drop int `json:"limit_vserver_drop"`
    Limit_total_drop int `json:"limit_total_drop"`
    Lockup_time_left int `json:"lockup_time_left"`
    Curr_rate int `json:"curr_rate"`
    V6_over_limit_drop int `json:"v6_over_limit_drop"`
    V6_limit_intf_drop int `json:"v6_limit_intf_drop"`
    V6_limit_vserver_drop int `json:"v6_limit_vserver_drop"`
    V6_limit_total_drop int `json:"v6_limit_total_drop"`
    V6_lockup_time_left int `json:"v6_lockup_time_left"`
    V6_curr_rate int `json:"v6_curr_rate"`
}

func (p *SystemIcmpRateStats) GetId() string{
    return "1"
}

func (p *SystemIcmpRateStats) getPath() string{
    return "system/icmp-rate/stats"
}

func (p *SystemIcmpRateStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemIcmpRateStats,error) {
logger.Println("SystemIcmpRateStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemIcmpRateStats
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
