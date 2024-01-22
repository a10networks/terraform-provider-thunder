

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallAlgFtpStats struct {
    
    Stats Cgnv6StatefulFirewallAlgFtpStatsStats `json:"stats"`

}
type DataCgnv6StatefulFirewallAlgFtpStats struct {
    DtCgnv6StatefulFirewallAlgFtpStats Cgnv6StatefulFirewallAlgFtpStats `json:"ftp"`
}


type Cgnv6StatefulFirewallAlgFtpStatsStats struct {
    ClientPortRequest int `json:"client-port-request"`
    ClientEprtRequest int `json:"client-eprt-request"`
    ServerPasvReply int `json:"server-pasv-reply"`
    ServerEpsvReply int `json:"server-epsv-reply"`
}

func (p *Cgnv6StatefulFirewallAlgFtpStats) GetId() string{
    return "1"
}

func (p *Cgnv6StatefulFirewallAlgFtpStats) getPath() string{
    return "cgnv6/stateful-firewall/alg/ftp/stats"
}

func (p *Cgnv6StatefulFirewallAlgFtpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6StatefulFirewallAlgFtpStats,error) {
logger.Println("Cgnv6StatefulFirewallAlgFtpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6StatefulFirewallAlgFtpStats
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
