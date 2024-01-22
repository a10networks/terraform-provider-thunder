

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwDdosProtectionStats struct {
    
    Stats FwDdosProtectionStatsStats `json:"stats"`

}
type DataFwDdosProtectionStats struct {
    DtFwDdosProtectionStats FwDdosProtectionStats `json:"ddos-protection"`
}


type FwDdosProtectionStatsStats struct {
    Ddos_entries_too_many int `json:"ddos_entries_too_many"`
    Ddos_entry_added int `json:"ddos_entry_added"`
    Ddos_entry_removed int `json:"ddos_entry_removed"`
    Ddos_entry_added_to_bgp int `json:"ddos_entry_added_to_bgp"`
    Ddos_entry_removed_from_bgp int `json:"ddos_entry_removed_from_bgp"`
    Ddos_entry_add_to_bgp_failure int `json:"ddos_entry_add_to_bgp_failure"`
    Ddos_entry_remove_from_bgp_failure int `json:"ddos_entry_remove_from_bgp_failure"`
    Ddos_packet_dropped int `json:"ddos_packet_dropped"`
}

func (p *FwDdosProtectionStats) GetId() string{
    return "1"
}

func (p *FwDdosProtectionStats) getPath() string{
    return "fw/ddos-protection/stats"
}

func (p *FwDdosProtectionStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwDdosProtectionStats,error) {
logger.Println("FwDdosProtectionStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwDdosProtectionStats
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
