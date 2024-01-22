

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallGlobalStats struct {
    
    Stats Cgnv6StatefulFirewallGlobalStatsStats `json:"stats"`

}
type DataCgnv6StatefulFirewallGlobalStats struct {
    DtCgnv6StatefulFirewallGlobalStats Cgnv6StatefulFirewallGlobalStats `json:"global"`
}


type Cgnv6StatefulFirewallGlobalStatsStats struct {
    Tcp_packet_process int `json:"tcp_packet_process"`
    Udp_packet_process int `json:"udp_packet_process"`
    Other_packet_process int `json:"other_packet_process"`
    Packet_inbound_deny int `json:"packet_inbound_deny"`
    Packet_process_failure int `json:"packet_process_failure"`
    Outbound_session_created int `json:"outbound_session_created"`
    Outbound_session_freed int `json:"outbound_session_freed"`
    Inbound_session_created int `json:"inbound_session_created"`
    Inbound_session_freed int `json:"inbound_session_freed"`
    Tcp_session_created int `json:"tcp_session_created"`
    Tcp_session_freed int `json:"tcp_session_freed"`
    Udp_session_created int `json:"udp_session_created"`
    Udp_session_freed int `json:"udp_session_freed"`
    Other_session_created int `json:"other_session_created"`
    Other_session_freed int `json:"other_session_freed"`
    Session_creation_failure int `json:"session_creation_failure"`
    No_fwd_route int `json:"no_fwd_route"`
    No_rev_route int `json:"no_rev_route"`
    Packet_standby_drop int `json:"packet_standby_drop"`
    Tcp_fullcone_created int `json:"tcp_fullcone_created"`
    Tcp_fullcone_freed int `json:"tcp_fullcone_freed"`
    Udp_fullcone_created int `json:"udp_fullcone_created"`
    Udp_fullcone_freed int `json:"udp_fullcone_freed"`
    Fullcone_creation_failure int `json:"fullcone_creation_failure"`
    Eif_process int `json:"eif_process"`
    One_arm_drop int `json:"one_arm_drop"`
    No_class_list_match int `json:"no_class_list_match"`
}

func (p *Cgnv6StatefulFirewallGlobalStats) GetId() string{
    return "1"
}

func (p *Cgnv6StatefulFirewallGlobalStats) getPath() string{
    return "cgnv6/stateful-firewall/global/stats"
}

func (p *Cgnv6StatefulFirewallGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6StatefulFirewallGlobalStats,error) {
logger.Println("Cgnv6StatefulFirewallGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6StatefulFirewallGlobalStats
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
