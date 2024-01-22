

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL4OtherStats struct {
    
    Stats DdosL4OtherStatsStats `json:"stats"`

}
type DataDdosL4OtherStats struct {
    DtDdosL4OtherStats DdosL4OtherStats `json:"l4-other"`
}


type DdosL4OtherStatsStats struct {
    Other_receive int `json:"other_receive"`
    Other_frag_rcvd int `json:"other_frag_rcvd"`
    Other_total_drop int `json:"other_total_drop"`
    Other_dst_drop int `json:"other_dst_drop"`
    Other_frag_drop int `json:"other_frag_drop"`
    Other_src_drop int `json:"other_src_drop"`
    Other_drop_black_user_cfg_src int `json:"other_drop_black_user_cfg_src"`
    Other_src_dst_drop int `json:"other_src_dst_drop"`
    Other_drop_black_user_cfg_src_dst int `json:"other_drop_black_user_cfg_src_dst"`
    Dst_other_filter_match int `json:"dst_other_filter_match"`
    Dst_other_filter_not_match int `json:"dst_other_filter_not_match"`
    Dst_other_filter_action_blacklist int `json:"dst_other_filter_action_blacklist"`
    Dst_other_filter_action_drop int `json:"dst_other_filter_action_drop"`
    Dst_other_filter_action_default_pass int `json:"dst_other_filter_action_default_pass"`
    Dst_other_filter_action_whitelist int `json:"dst_other_filter_action_whitelist"`
    Src_other_filter_match int `json:"src_other_filter_match"`
    Src_other_filter_not_match int `json:"src_other_filter_not_match"`
    Src_other_filter_action_blacklist int `json:"src_other_filter_action_blacklist"`
    Src_other_filter_action_drop int `json:"src_other_filter_action_drop"`
    Src_other_filter_action_default_pass int `json:"src_other_filter_action_default_pass"`
    Src_other_filter_action_whitelist int `json:"src_other_filter_action_whitelist"`
    Src_dst_other_filter_match int `json:"src_dst_other_filter_match"`
    Src_dst_other_filter_not_match int `json:"src_dst_other_filter_not_match"`
    Src_dst_other_filter_action_blacklist int `json:"src_dst_other_filter_action_blacklist"`
    Src_dst_other_filter_action_drop int `json:"src_dst_other_filter_action_drop"`
    Src_dst_other_filter_action_default_pass int `json:"src_dst_other_filter_action_default_pass"`
    Src_dst_other_filter_action_whitelist int `json:"src_dst_other_filter_action_whitelist"`
    Other_any_exceed int `json:"other_any_exceed"`
    Other_drop_bl int `json:"other_drop_bl"`
    Other_total_bytes_rcv int `json:"other_total_bytes_rcv"`
    Other_total_bytes_drop int `json:"other_total_bytes_drop"`
}

func (p *DdosL4OtherStats) GetId() string{
    return "1"
}

func (p *DdosL4OtherStats) getPath() string{
    return "ddos/l4-other/stats"
}

func (p *DdosL4OtherStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL4OtherStats,error) {
logger.Println("DdosL4OtherStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL4OtherStats
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
