

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Lw4o6GlobalStats struct {
    
    Stats Cgnv6Lw4o6GlobalStatsStats `json:"stats"`

}
type DataCgnv6Lw4o6GlobalStats struct {
    DtCgnv6Lw4o6GlobalStats Cgnv6Lw4o6GlobalStats `json:"global"`
}


type Cgnv6Lw4o6GlobalStatsStats struct {
    Entry_count int `json:"entry_count"`
    Self_hairpinning_drop int `json:"self_hairpinning_drop"`
    All_hairpinning_drop int `json:"all_hairpinning_drop"`
    No_match_icmpv6_sent int `json:"no_match_icmpv6_sent"`
    No_match_icmp_sent int `json:"no_match_icmp_sent"`
    Icmp_inbound_drop int `json:"icmp_inbound_drop"`
    Fwd_lookup_failed int `json:"fwd_lookup_failed"`
    Rev_lookup_failed int `json:"rev_lookup_failed"`
    Interface_not_configured int `json:"interface_not_configured"`
    No_binding_table_matches_fwd int `json:"no_binding_table_matches_fwd"`
    No_binding_table_matches_rev int `json:"no_binding_table_matches_rev"`
}

func (p *Cgnv6Lw4o6GlobalStats) GetId() string{
    return "1"
}

func (p *Cgnv6Lw4o6GlobalStats) getPath() string{
    return "cgnv6/lw-4o6/global/stats"
}

func (p *Cgnv6Lw4o6GlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Lw4o6GlobalStats,error) {
logger.Println("Cgnv6Lw4o6GlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Lw4o6GlobalStats
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
