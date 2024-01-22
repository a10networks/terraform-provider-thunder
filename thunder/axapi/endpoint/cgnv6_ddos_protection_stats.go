

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DdosProtectionStats struct {
    
    Stats Cgnv6DdosProtectionStatsStats `json:"stats"`

}
type DataCgnv6DdosProtectionStats struct {
    DtCgnv6DdosProtectionStats Cgnv6DdosProtectionStats `json:"ddos-protection"`
}


type Cgnv6DdosProtectionStatsStats struct {
    L3_entry_added int `json:"l3_entry_added"`
    L3_entry_deleted int `json:"l3_entry_deleted"`
    L3_entry_added_to_bgp int `json:"l3_entry_added_to_bgp"`
    L3_entry_removed_from_bgp int `json:"l3_entry_removed_from_bgp"`
    L3_entry_added_to_hw int `json:"l3_entry_added_to_hw"`
    L3_entry_removed_from_hw int `json:"l3_entry_removed_from_hw"`
    L3_entry_too_many int `json:"l3_entry_too_many"`
    L3_entry_match_drop int `json:"l3_entry_match_drop"`
    L3_entry_match_drop_hw int `json:"l3_entry_match_drop_hw"`
    L3_entry_drop_max_hw_exceeded int `json:"l3_entry_drop_max_hw_exceeded"`
    L4_entry_added int `json:"l4_entry_added"`
    L4_entry_deleted int `json:"l4_entry_deleted"`
    L4_entry_added_to_hw int `json:"l4_entry_added_to_hw"`
    L4_entry_removed_from_hw int `json:"l4_entry_removed_from_hw"`
    L4_hw_out_of_entries int `json:"l4_hw_out_of_entries"`
    L4_entry_match_drop int `json:"l4_entry_match_drop"`
    L4_entry_match_drop_hw int `json:"l4_entry_match_drop_hw"`
    L4_entry_drop_max_hw_exceeded int `json:"l4_entry_drop_max_hw_exceeded"`
    L4_entry_list_alloc int `json:"l4_entry_list_alloc"`
    L4_entry_list_free int `json:"l4_entry_list_free"`
    L4_entry_list_alloc_failure int `json:"l4_entry_list_alloc_failure"`
    Ip_node_alloc int `json:"ip_node_alloc"`
    Ip_node_free int `json:"ip_node_free"`
    Ip_node_alloc_failure int `json:"ip_node_alloc_failure"`
    Ip_port_block_alloc int `json:"ip_port_block_alloc"`
    Ip_port_block_free int `json:"ip_port_block_free"`
    Ip_port_block_alloc_failure int `json:"ip_port_block_alloc_failure"`
    Ip_other_block_alloc int `json:"ip_other_block_alloc"`
    Ip_other_block_free int `json:"ip_other_block_free"`
    Ip_other_block_alloc_failure int `json:"ip_other_block_alloc_failure"`
    Entry_added_shadow int `json:"entry_added_shadow"`
    Entry_invalidated int `json:"entry_invalidated"`
    L3_entry_add_to_bgp_failure int `json:"l3_entry_add_to_bgp_failure"`
    L3_entry_remove_from_bgp_failure int `json:"l3_entry_remove_from_bgp_failure"`
    L3_entry_add_to_hw_failure int `json:"l3_entry_add_to_hw_failure"`
    Syn_cookie_syn_ack_sent int `json:"syn_cookie_syn_ack_sent"`
    Syn_cookie_verification_passed int `json:"syn_cookie_verification_passed"`
    Syn_cookie_verification_failed int `json:"syn_cookie_verification_failed"`
}

func (p *Cgnv6DdosProtectionStats) GetId() string{
    return "1"
}

func (p *Cgnv6DdosProtectionStats) getPath() string{
    return "cgnv6/ddos-protection/stats"
}

func (p *Cgnv6DdosProtectionStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6DdosProtectionStats,error) {
logger.Println("Cgnv6DdosProtectionStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6DdosProtectionStats
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
