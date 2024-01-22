

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIpsecSaStats struct {
    
    SamplingEnable []VpnIpsecSaStatsSamplingEnable `json:"sampling-enable"`

}
type DataVpnIpsecSaStats struct {
    DtVpnIpsecSaStats VpnIpsecSaStats `json:"ipsec-sa-stats"`
}


type VpnIpsecSaStatsSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *VpnIpsecSaStats) GetId() string{
    
    
    return p.SamplingEnable[0].Counters1
}

func (p *VpnIpsecSaStats) getPath() string{
    return "vpn/ipsec-sa-stats"
}

func (p *VpnIpsecSaStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIpsecSaStats,error) {
logger.Println("VpnIpsecSaStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIpsecSaStats
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
