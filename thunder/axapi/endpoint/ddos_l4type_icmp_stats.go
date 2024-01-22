

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL4typeIcmpStats struct {
    
    Stats DdosL4typeIcmpStatsStats `json:"stats"`

}
type DataDdosL4typeIcmpStats struct {
    DtDdosL4typeIcmpStats DdosL4typeIcmpStats `json:"l4type-icmp"`
}


type DdosL4typeIcmpStatsStats struct {
    Rate_type0_exceed int `json:"rate_type0_exceed"`
    Rate_type1_exceed int `json:"rate_type1_exceed"`
    Rate_type2_exceed int `json:"rate_type2_exceed"`
    Type_deny_drop int `json:"type_deny_drop"`
    Icmpv4_rfc_undef_drop int `json:"icmpv4_rfc_undef_drop"`
    Icmpv6_rfc_undef_drop int `json:"icmpv6_rfc_undef_drop"`
    Wildcard_deny_drop int `json:"wildcard_deny_drop"`
    Port_rcvd int `json:"port_rcvd"`
    Port_drop int `json:"port_drop"`
    Port_pkt_sent int `json:"port_pkt_sent"`
    Type int `json:"type"`
    Type_bl int `json:"type_bl"`
    Wildcard int `json:"wildcard"`
    Wildcard_bl int `json:"wildcard_bl"`
    Rate_type0_exceed_drop int `json:"rate_type0_exceed_drop"`
    Rate_type0_exceed_bl int `json:"rate_type0_exceed_bl"`
    Rate_type1_exceed_drop int `json:"rate_type1_exceed_drop"`
    Rate_type1_exceed_bl int `json:"rate_type1_exceed_bl"`
    Rate_type2_exceed_drop int `json:"rate_type2_exceed_drop"`
    Rate_type2_exceed_bl int `json:"rate_type2_exceed_bl"`
    Port_bytes int `json:"port_bytes"`
    Outbound_port_bytes int `json:"outbound_port_bytes"`
    Outbound_port_rcvd int `json:"outbound_port_rcvd"`
    Outbound_port_pkt_sent int `json:"outbound_port_pkt_sent"`
    Port_bytes_sent int `json:"port_bytes_sent"`
    Port_bytes_drop int `json:"port_bytes_drop"`
    Port_src_bl int `json:"port_src_bl"`
    Port_pkt_rate_exceed int `json:"port_pkt_rate_exceed"`
    Port_kbit_rate_exceed int `json:"port_kbit_rate_exceed"`
    Exceed_drop_prate_src int `json:"exceed_drop_prate_src"`
    Exceed_drop_brate_src int `json:"exceed_drop_brate_src"`
    Outbound_port_bytes_sent int `json:"outbound_port_bytes_sent"`
    Outbound_port_drop int `json:"outbound_port_drop"`
    Outbound_port_bytes_drop int `json:"outbound_port_bytes_drop"`
    Exceed_drop_brate_src_pkt int `json:"exceed_drop_brate_src_pkt"`
    Port_kbit_rate_exceed_pkt int `json:"port_kbit_rate_exceed_pkt"`
    Bl int `json:"bl"`
    Src_drop int `json:"src_drop"`
    Frag_rcvd int `json:"frag_rcvd"`
    Frag_drop int `json:"frag_drop"`
    Frag_timeout int `json:"frag_timeout"`
    Src_frag_drop int `json:"src_frag_drop"`
    Sflow_internal_samples_packed int `json:"sflow_internal_samples_packed"`
    Sflow_external_samples_packed int `json:"sflow_external_samples_packed"`
    Sflow_internal_packets_sent int `json:"sflow_internal_packets_sent"`
    Sflow_external_packets_sent int `json:"sflow_external_packets_sent"`
    Exceed_action_tunnel int `json:"exceed_action_tunnel"`
    Dst_hw_drop int `json:"dst_hw_drop"`
    Exceed_action_drop int `json:"exceed_action_drop"`
}

func (p *DdosL4typeIcmpStats) GetId() string{
    return "1"
}

func (p *DdosL4typeIcmpStats) getPath() string{
    return "ddos/l4type-icmp/stats"
}

func (p *DdosL4typeIcmpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL4typeIcmpStats,error) {
logger.Println("DdosL4typeIcmpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL4typeIcmpStats
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
