

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkTwampResponderIpStats struct {
    
    Stats NetworkTwampResponderIpStatsStats `json:"stats"`

}
type DataNetworkTwampResponderIpStats struct {
    DtNetworkTwampResponderIpStats NetworkTwampResponderIpStats `json:"ip"`
}


type NetworkTwampResponderIpStatsStats struct {
    Rx_pkts int `json:"rx_pkts"`
    Tx_pkts int `json:"tx_pkts"`
    Rx_drop_not_enabled_v4 int `json:"rx_drop_not_enabled_v4"`
    Rx_acl_drop int `json:"rx_acl_drop"`
    Twamp_hdr_len_err int `json:"twamp_hdr_len_err"`
    No_route_err int `json:"no_route_err"`
    Other_err int `json:"other_err"`
}

func (p *NetworkTwampResponderIpStats) GetId() string{
    return "1"
}

func (p *NetworkTwampResponderIpStats) getPath() string{
    return "network/twamp/responder/ip/stats"
}

func (p *NetworkTwampResponderIpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkTwampResponderIpStats,error) {
logger.Println("NetworkTwampResponderIpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkTwampResponderIpStats
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
