

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListIpStats struct {
    
    Ipv4Addr string `json:"ipv4-addr"`
    Stats Cgnv6LsnRuleListIpStatsStats `json:"stats"`
    Name string 

}
type DataCgnv6LsnRuleListIpStats struct {
    DtCgnv6LsnRuleListIpStats Cgnv6LsnRuleListIpStats `json:"ip"`
}


type Cgnv6LsnRuleListIpStatsStats struct {
}

func (p *Cgnv6LsnRuleListIpStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListIpStats) getPath() string{
    
    return "cgnv6/lsn-rule-list/" +p.Name + "/ip/"+p.Ipv4Addr+"/stats"
}

func (p *Cgnv6LsnRuleListIpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRuleListIpStats,error) {
logger.Println("Cgnv6LsnRuleListIpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRuleListIpStats
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
