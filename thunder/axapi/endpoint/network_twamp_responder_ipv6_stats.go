

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkTwampResponderIpv6Stats struct {
    
    Stats NetworkTwampResponderIpv6StatsStats `json:"stats"`

}
type DataNetworkTwampResponderIpv6Stats struct {
    DtNetworkTwampResponderIpv6Stats NetworkTwampResponderIpv6Stats `json:"ipv6"`
}


type NetworkTwampResponderIpv6StatsStats struct {
    Rx_pkts_v6 int `json:"rx_pkts_v6"`
    Tx_pkts_v6 int `json:"tx_pkts_v6"`
    Rx_drop_not_enabled_v6 int `json:"rx_drop_not_enabled_v6"`
    Rx_acl_drop_v6 int `json:"rx_acl_drop_v6"`
    Twamp_hdr_len_err_v6 int `json:"twamp_hdr_len_err_v6"`
    No_route_err_v6 int `json:"no_route_err_v6"`
    Other_err_v6 int `json:"other_err_v6"`
}

func (p *NetworkTwampResponderIpv6Stats) GetId() string{
    return "1"
}

func (p *NetworkTwampResponderIpv6Stats) getPath() string{
    return "network/twamp/responder/ipv6/stats"
}

func (p *NetworkTwampResponderIpv6Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkTwampResponderIpv6Stats,error) {
logger.Println("NetworkTwampResponderIpv6Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkTwampResponderIpv6Stats
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
