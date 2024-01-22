

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL4IcmpStats struct {
    
    Stats DdosL4IcmpStatsStats `json:"stats"`

}
type DataDdosL4IcmpStats struct {
    DtDdosL4IcmpStats DdosL4IcmpStats `json:"l4-icmp"`
}


type DdosL4IcmpStatsStats struct {
    Icmp_src_drop int `json:"icmp_src_drop"`
    Icmp_dst_drop int `json:"icmp_dst_drop"`
    Icmp_wildcard_bl int `json:"icmp_wildcard_bl"`
    Src_icmp_bl_user_config int `json:"src_icmp_bl_user_config"`
    Icmp_rcvd int `json:"icmp_rcvd"`
    Icmp_echo_rcvd int `json:"icmp_echo_rcvd"`
    Icmp_other_rcvd int `json:"icmp_other_rcvd"`
    Icmp_src_dst_drop int `json:"icmp_src_dst_drop"`
    Icmp_src_dst_bl_drop_user_cfg int `json:"icmp_src_dst_bl_drop_user_cfg"`
    Icmp_type_deny_drop int `json:"icmp_type_deny_drop"`
    Icmp_v4_rfc_undef_drop int `json:"icmp_v4_rfc_undef_drop"`
    Icmp_v6_rfc_undef_drop int `json:"icmp_v6_rfc_undef_drop"`
    Icmp_wildcard_deny_drop int `json:"icmp_wildcard_deny_drop"`
    Icmp_rate_exceed0 int `json:"icmp_rate_exceed0"`
    Icmp_rate_exceed1 int `json:"icmp_rate_exceed1"`
    Icmp_rate_exceed2 int `json:"icmp_rate_exceed2"`
    Icmp_type int `json:"icmp_type"`
    Icmp_type_bl int `json:"icmp_type_bl"`
    Icmp_rate_exceed0_drop int `json:"icmp_rate_exceed0_drop"`
    Icmp_rate_exceed0_bl int `json:"icmp_rate_exceed0_bl"`
    Icmp_rate_exceed1_drop int `json:"icmp_rate_exceed1_drop"`
    Icmp_rate_exceed1_bl int `json:"icmp_rate_exceed1_bl"`
    Icmp_rate_exceed2_drop int `json:"icmp_rate_exceed2_drop"`
    Icmp_rate_exceed2_bl int `json:"icmp_rate_exceed2_bl"`
    Icmp_wildcard int `json:"icmp_wildcard"`
    Icmp_any_exceed int `json:"icmp_any_exceed"`
    Icmp_drop_bl int `json:"icmp_drop_bl"`
    Icmp_frag_rcvd int `json:"icmp_frag_rcvd"`
    Icmp_frag_drop int `json:"icmp_frag_drop"`
    Icmp_total_bytes_rcv int `json:"icmp_total_bytes_rcv"`
    Icmp_total_drop int `json:"icmp_total_drop"`
    Icmp_total_bytes_drop int `json:"icmp_total_bytes_drop"`
}

func (p *DdosL4IcmpStats) GetId() string{
    return "1"
}

func (p *DdosL4IcmpStats) getPath() string{
    return "ddos/l4-icmp/stats"
}

func (p *DdosL4IcmpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL4IcmpStats,error) {
logger.Println("DdosL4IcmpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL4IcmpStats
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
