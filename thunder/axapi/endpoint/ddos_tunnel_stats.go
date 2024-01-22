

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTunnelStats struct {
    
    Stats DdosTunnelStatsStats `json:"stats"`

}
type DataDdosTunnelStats struct {
    DtDdosTunnelStats DdosTunnelStats `json:"tunnel"`
}


type DdosTunnelStatsStats struct {
    Ip_tunnel_rcvd int `json:"ip_tunnel_rcvd"`
    Ip_tunnel_rate_limit_inner int `json:"ip_tunnel_rate_limit_inner"`
    Ip_tunnel_encap int `json:"ip_tunnel_encap"`
    Ip_tunnel_encap_fail int `json:"ip_tunnel_encap_fail"`
    Ip_tunnel_decap int `json:"ip_tunnel_decap"`
    Ip_tunnel_decap_fail int `json:"ip_tunnel_decap_fail"`
    Ip6_tunnel_rcvd int `json:"ip6_tunnel_rcvd"`
    Ip6_tunnel_rate_limit_inner int `json:"ip6_tunnel_rate_limit_inner"`
    Ip6_tunnel_encap int `json:"ip6_tunnel_encap"`
    Ip6_tunnel_encap_fail int `json:"ip6_tunnel_encap_fail"`
    Ip6_tunnel_decap int `json:"ip6_tunnel_decap"`
    Ip6_tunnel_decap_fail int `json:"ip6_tunnel_decap_fail"`
    Ip_gre_tunnel_rcvd int `json:"ip_gre_tunnel_rcvd"`
    Ip_gre_tunnel_rate_limit_inner int `json:"ip_gre_tunnel_rate_limit_inner"`
    Ip_gre_tunnel_encap int `json:"ip_gre_tunnel_encap"`
    Ip_gre_tunnel_encap_fail int `json:"ip_gre_tunnel_encap_fail"`
    Ip_gre_tunnel_decap int `json:"ip_gre_tunnel_decap"`
    Ip_gre_tunnel_decap_fail int `json:"ip_gre_tunnel_decap_fail"`
    Ip_gre_tunnel_encap_key int `json:"ip_gre_tunnel_encap_key"`
    Ip_gre_tunnel_decap_key int `json:"ip_gre_tunnel_decap_key"`
    Ip_gre_tunnel_decap_drop int `json:"ip_gre_tunnel_decap_drop"`
    Ip6_gre_tunnel_rcvd int `json:"ip6_gre_tunnel_rcvd"`
    Ip6_gre_tunnel_rate_limit_inner int `json:"ip6_gre_tunnel_rate_limit_inner"`
    Ip6_gre_tunnel_encap int `json:"ip6_gre_tunnel_encap"`
    Ip6_gre_tunnel_encap_fail int `json:"ip6_gre_tunnel_encap_fail"`
    Ip6_gre_tunnel_decap int `json:"ip6_gre_tunnel_decap"`
    Ip6_gre_tunnel_decap_fail int `json:"ip6_gre_tunnel_decap_fail"`
    Ip6_gre_tunnel_encap_key int `json:"ip6_gre_tunnel_encap_key"`
    Ip6_gre_tunnel_decap_key int `json:"ip6_gre_tunnel_decap_key"`
    Ip6_gre_tunnel_decap_drop int `json:"ip6_gre_tunnel_decap_drop"`
    Ip_esp_tunnel_rcvd int `json:"ip_esp_tunnel_rcvd"`
    Ip_esp_tunnel_inspect int `json:"ip_esp_tunnel_inspect"`
    Ip_esp_tunnel_inspect_fail int `json:"ip_esp_tunnel_inspect_fail"`
    Ip6_esp_tunnel_rcvd int `json:"ip6_esp_tunnel_rcvd"`
    Ip6_esp_tunnel_inspect int `json:"ip6_esp_tunnel_inspect"`
    Ip6_esp_tunnel_inspect_fail int `json:"ip6_esp_tunnel_inspect_fail"`
    Ip_vxlan_tunnel_rcvd int `json:"ip_vxlan_tunnel_rcvd"`
    Ip_vxlan_tunnel_invalid_vni int `json:"ip_vxlan_tunnel_invalid_vni"`
    Ip_vxlan_tunnel_decap int `json:"ip_vxlan_tunnel_decap"`
    Ip_vxlan_tunnel_decap_err int `json:"ip_vxlan_tunnel_decap_err"`
    Ip_gre_tunnel_keepalive_rcvd int `json:"ip_gre_tunnel_keepalive_rcvd"`
    Jumbo_in_tunnel_drop int `json:"jumbo_in_tunnel_drop"`
}

func (p *DdosTunnelStats) GetId() string{
    return "1"
}

func (p *DdosTunnelStats) getPath() string{
    return "ddos/tunnel/stats"
}

func (p *DdosTunnelStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosTunnelStats,error) {
logger.Println("DdosTunnelStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosTunnelStats
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
