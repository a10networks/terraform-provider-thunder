

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDomainIpStats struct {
    
    Stats Cgnv6LsnRuleListDomainIpStatsStats `json:"stats"`
    Name string 

}
type DataCgnv6LsnRuleListDomainIpStats struct {
    DtCgnv6LsnRuleListDomainIpStats Cgnv6LsnRuleListDomainIpStats `json:"domain-ip"`
}


type Cgnv6LsnRuleListDomainIpStatsStats struct {
}

func (p *Cgnv6LsnRuleListDomainIpStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListDomainIpStats) getPath() string{
    
    return "cgnv6/lsn-rule-list/" +p.Name + "/domain-ip/stats"
}

func (p *Cgnv6LsnRuleListDomainIpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRuleListDomainIpStats,error) {
logger.Println("Cgnv6LsnRuleListDomainIpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRuleListDomainIpStats
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
