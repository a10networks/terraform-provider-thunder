

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallAlgTftpStats struct {
    
    Stats Cgnv6StatefulFirewallAlgTftpStatsStats `json:"stats"`

}
type DataCgnv6StatefulFirewallAlgTftpStats struct {
    DtCgnv6StatefulFirewallAlgTftpStats Cgnv6StatefulFirewallAlgTftpStats `json:"tftp"`
}


type Cgnv6StatefulFirewallAlgTftpStatsStats struct {
    SessionCreated int `json:"session-created"`
}

func (p *Cgnv6StatefulFirewallAlgTftpStats) GetId() string{
    return "1"
}

func (p *Cgnv6StatefulFirewallAlgTftpStats) getPath() string{
    return "cgnv6/stateful-firewall/alg/tftp/stats"
}

func (p *Cgnv6StatefulFirewallAlgTftpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6StatefulFirewallAlgTftpStats,error) {
logger.Println("Cgnv6StatefulFirewallAlgTftpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6StatefulFirewallAlgTftpStats
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
