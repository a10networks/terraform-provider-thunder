

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpThreatListStats struct {
    
    Stats SystemIpThreatListStatsStats `json:"stats"`

}
type DataSystemIpThreatListStats struct {
    DtSystemIpThreatListStats SystemIpThreatListStats `json:"ip-threat-list"`
}


type SystemIpThreatListStatsStats struct {
    Packet_hit_count_in_sw int `json:"packet_hit_count_in_sw"`
    Packet_hit_count_in_spe int `json:"packet_hit_count_in_spe"`
    Entries_added_in_sw int `json:"entries_added_in_sw"`
    Entries_removed_from_sw int `json:"entries_removed_from_sw"`
    Entries_added_in_spe int `json:"entries_added_in_spe"`
    Entries_removed_from_spe int `json:"entries_removed_from_spe"`
    Error_out_of_memory int `json:"error_out_of_memory"`
    Error_out_of_spe_entries int `json:"error_out_of_spe_entries"`
}

func (p *SystemIpThreatListStats) GetId() string{
    return "1"
}

func (p *SystemIpThreatListStats) getPath() string{
    return "system/ip-threat-list/stats"
}

func (p *SystemIpThreatListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemIpThreatListStats,error) {
logger.Println("SystemIpThreatListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemIpThreatListStats
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
