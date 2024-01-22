

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPlayerIdGlobalStats struct {
    
    Stats SlbPlayerIdGlobalStatsStats `json:"stats"`

}
type DataSlbPlayerIdGlobalStats struct {
    DtSlbPlayerIdGlobalStats SlbPlayerIdGlobalStats `json:"player-id-global"`
}


type SlbPlayerIdGlobalStatsStats struct {
    Total_playerids_created int `json:"total_playerids_created"`
    Total_playerids_deleted int `json:"total_playerids_deleted"`
    Total_abs_max_age_outs int `json:"total_abs_max_age_outs"`
    Total_pkt_activity_age_outs int `json:"total_pkt_activity_age_outs"`
    Total_invalid_playerid_pkts int `json:"total_invalid_playerid_pkts"`
    Total_invalid_playerid_drops int `json:"total_invalid_playerid_drops"`
    Total_valid_playerid_pkts int `json:"total_valid_playerid_pkts"`
}

func (p *SlbPlayerIdGlobalStats) GetId() string{
    return "1"
}

func (p *SlbPlayerIdGlobalStats) getPath() string{
    return "slb/player-id-global/stats"
}

func (p *SlbPlayerIdGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbPlayerIdGlobalStats,error) {
logger.Println("SlbPlayerIdGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbPlayerIdGlobalStats
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
