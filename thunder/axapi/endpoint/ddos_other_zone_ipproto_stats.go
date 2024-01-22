

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosOtherZoneIpprotoStats struct {
    
    Stats DdosOtherZoneIpprotoStatsStats `json:"stats"`

}
type DataDdosOtherZoneIpprotoStats struct {
    DtDdosOtherZoneIpprotoStats DdosOtherZoneIpprotoStats `json:"other-zone-ipproto"`
}


type DdosOtherZoneIpprotoStatsStats struct {
    Filter1_match int `json:"filter1_match"`
    Filter2_match int `json:"filter2_match"`
    Filter3_match int `json:"filter3_match"`
    Filter4_match int `json:"filter4_match"`
    Filter5_match int `json:"filter5_match"`
    Filter_none_match int `json:"filter_none_match"`
    Port_rcvd int `json:"port_rcvd"`
    Port_drop int `json:"port_drop"`
    Port_pkt_sent int `json:"port_pkt_sent"`
    Port_pkt_rate_exceed int `json:"port_pkt_rate_exceed"`
    Port_kbit_rate_exceed int `json:"port_kbit_rate_exceed"`
    Filter_auth_fail int `json:"filter_auth_fail"`
    Port_bytes int `json:"port_bytes"`
    Outbound_port_bytes int `json:"outbound_port_bytes"`
    Outbound_port_rcvd int `json:"outbound_port_rcvd"`
    Outbound_port_pkt_sent int `json:"outbound_port_pkt_sent"`
    Port_bytes_sent int `json:"port_bytes_sent"`
    Port_bytes_drop int `json:"port_bytes_drop"`
    Port_src_bl int `json:"port_src_bl"`
    Port_src_escalation int `json:"port_src_escalation"`
    Current_es_level int `json:"current_es_level"`
    Filter_action_blacklist int `json:"filter_action_blacklist"`
    Filter_action_drop int `json:"filter_action_drop"`
    Filter_action_default_pass int `json:"filter_action_default_pass"`
    Filter_action_whitelist int `json:"filter_action_whitelist"`
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
    Filter_total_not_match int `json:"filter_total_not_match"`
    Src_filter_action_blacklist int `json:"src_filter_action_blacklist"`
    Src_filter_action_drop int `json:"src_filter_action_drop"`
    Src_filter_action_default_pass int `json:"src_filter_action_default_pass"`
    Src_filter_action_whitelist int `json:"src_filter_action_whitelist"`
    Secondary_port_pkt_rate_exceed int `json:"secondary_port_pkt_rate_exceed"`
    Secondary_port_kbit_rate_exceed int `json:"secondary_port_kbit_rate_exceed"`
    Secondary_port_kbit_rate_exceed_pkt int `json:"secondary_port_kbit_rate_exceed_pkt"`
    No_policy_class_list_match int `json:"no_policy_class_list_match"`
    Frag_timeout int `json:"frag_timeout"`
    Src_filter1_match int `json:"src_filter1_match"`
    Src_filter2_match int `json:"src_filter2_match"`
    Src_filter3_match int `json:"src_filter3_match"`
    Src_filter4_match int `json:"src_filter4_match"`
    Src_filter5_match int `json:"src_filter5_match"`
    Src_filter_none_match int `json:"src_filter_none_match"`
    Src_filter_total_not_match int `json:"src_filter_total_not_match"`
    Src_filter_auth_fail int `json:"src_filter_auth_fail"`
    Src_frag_drop int `json:"src_frag_drop"`
    Sflow_internal_samples_packed int `json:"sflow_internal_samples_packed"`
    Sflow_external_samples_packed int `json:"sflow_external_samples_packed"`
    Sflow_internal_packets_sent int `json:"sflow_internal_packets_sent"`
    Sflow_external_packets_sent int `json:"sflow_external_packets_sent"`
    Exceed_action_tunnel int `json:"exceed_action_tunnel"`
    Dst_hw_drop int `json:"dst_hw_drop"`
    Secondary_port_hit int `json:"secondary_port_hit"`
    Src_zone_service_entry_learned int `json:"src_zone_service_entry_learned"`
    Src_zone_service_entry_aged int `json:"src_zone_service_entry_aged"`
    Dst_hw_drop_inserted int `json:"dst_hw_drop_inserted"`
    Dst_hw_drop_removed int `json:"dst_hw_drop_removed"`
    Src_hw_drop_inserted int `json:"src_hw_drop_inserted"`
    Src_hw_drop_removed int `json:"src_hw_drop_removed"`
    Exceed_action_drop int `json:"exceed_action_drop"`
    Ew_inbound_port_rcv int `json:"ew_inbound_port_rcv"`
    Ew_inbound_port_drop int `json:"ew_inbound_port_drop"`
    Ew_inbound_port_sent int `json:"ew_inbound_port_sent"`
    Ew_inbound_port_byte_rcv int `json:"ew_inbound_port_byte_rcv"`
    Ew_inbound_port_byte_drop int `json:"ew_inbound_port_byte_drop"`
    Ew_inbound_port_byte_sent int `json:"ew_inbound_port_byte_sent"`
    Ew_outbound_port_rcv int `json:"ew_outbound_port_rcv"`
    Ew_outbound_port_drop int `json:"ew_outbound_port_drop"`
    Ew_outbound_port_sent int `json:"ew_outbound_port_sent"`
    Ew_outbound_port_byte_rcv int `json:"ew_outbound_port_byte_rcv"`
    Ew_outbound_port_byte_sent int `json:"ew_outbound_port_byte_sent"`
    Ew_outbound_port_byte_drop int `json:"ew_outbound_port_byte_drop"`
    No_route_drop int `json:"no_route_drop"`
    Src_hw_drop int `json:"src_hw_drop"`
    Addr_filter_drop int `json:"addr_filter_drop"`
    Addr_filter_bl int `json:"addr_filter_bl"`
    Src_learn_overflow int `json:"src_learn_overflow"`
}

func (p *DdosOtherZoneIpprotoStats) GetId() string{
    return "1"
}

func (p *DdosOtherZoneIpprotoStats) getPath() string{
    return "ddos/other-zone-ipproto/stats"
}

func (p *DdosOtherZoneIpprotoStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosOtherZoneIpprotoStats,error) {
logger.Println("DdosOtherZoneIpprotoStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosOtherZoneIpprotoStats
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
