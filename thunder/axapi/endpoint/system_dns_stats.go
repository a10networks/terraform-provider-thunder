

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemDnsStats struct {
    
    Stats SystemDnsStatsStats `json:"stats"`

}
type DataSystemDnsStats struct {
    DtSystemDnsStats SystemDnsStats `json:"dns"`
}


type SystemDnsStatsStats struct {
    Slb_req int `json:"slb_req"`
    Slb_resp int `json:"slb_resp"`
    Slb_no_resp int `json:"slb_no_resp"`
    Slb_req_rexmit int `json:"slb_req_rexmit"`
    Slb_resp_no_match int `json:"slb_resp_no_match"`
    Slb_no_resource int `json:"slb_no_resource"`
    Nat_req int `json:"nat_req"`
    Nat_resp int `json:"nat_resp"`
    Nat_no_resp int `json:"nat_no_resp"`
    Nat_req_rexmit int `json:"nat_req_rexmit"`
    Nat_resp_no_match int `json:"nat_resp_no_match"`
    Nat_no_resource int `json:"nat_no_resource"`
    Nat_xid_reused int `json:"nat_xid_reused"`
    Filter_type_drop int `json:"filter_type_drop"`
    Filter_class_drop int `json:"filter_class_drop"`
    Filter_type_any_drop int `json:"filter_type_any_drop"`
    Slb_dns_client_ssl_succ int `json:"slb_dns_client_ssl_succ"`
    Slb_dns_server_ssl_succ int `json:"slb_dns_server_ssl_succ"`
    Slb_dns_udp_conn int `json:"slb_dns_udp_conn"`
    Slb_dns_udp_conn_succ int `json:"slb_dns_udp_conn_succ"`
    Slb_dns_padding_to_server_removed int `json:"slb_dns_padding_to_server_removed"`
    Slb_dns_padding_to_client_added int `json:"slb_dns_padding_to_client_added"`
    Slb_dns_edns_subnet_to_server_removed int `json:"slb_dns_edns_subnet_to_server_removed"`
    Slb_dns_udp_retransmit int `json:"slb_dns_udp_retransmit"`
    Slb_dns_udp_retransmit_fail int `json:"slb_dns_udp_retransmit_fail"`
    Rpz_action_drop int `json:"rpz_action_drop"`
    Rpz_action_pass_thru int `json:"rpz_action_pass_thru"`
    Rpz_action_tcp_only int `json:"rpz_action_tcp_only"`
    Rpz_action_nxdomain int `json:"rpz_action_nxdomain"`
    Rpz_action_nodata int `json:"rpz_action_nodata"`
    Rpz_action_local_data int `json:"rpz_action_local_data"`
    Slb_drop int `json:"slb_drop"`
    Nat_slb_drop int `json:"nat_slb_drop"`
    Invalid_q_len_to_udp int `json:"invalid_q_len_to_udp"`
    Slb_dns_edns_ecs_received int `json:"slb_dns_edns_ecs_received"`
    Slb_dns_edns_ecs_inserted int `json:"slb_dns_edns_ecs_inserted"`
}

func (p *SystemDnsStats) GetId() string{
    return "1"
}

func (p *SystemDnsStats) getPath() string{
    return "system/dns/stats"
}

func (p *SystemDnsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemDnsStats,error) {
logger.Println("SystemDnsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemDnsStats
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
