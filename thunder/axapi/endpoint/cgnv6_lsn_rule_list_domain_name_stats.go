

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDomainNameStats struct {
    
    NameDomain string `json:"name-domain"`
    Stats Cgnv6LsnRuleListDomainNameStatsStats `json:"stats"`
    Name string 

}
type DataCgnv6LsnRuleListDomainNameStats struct {
    DtCgnv6LsnRuleListDomainNameStats Cgnv6LsnRuleListDomainNameStats `json:"domain-name"`
}


type Cgnv6LsnRuleListDomainNameStatsStats struct {
}

func (p *Cgnv6LsnRuleListDomainNameStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListDomainNameStats) getPath() string{
    
    return "cgnv6/lsn-rule-list/" +p.Name + "/domain-name/"+p.NameDomain+"/stats"
}

func (p *Cgnv6LsnRuleListDomainNameStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRuleListDomainNameStats,error) {
logger.Println("Cgnv6LsnRuleListDomainNameStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRuleListDomainNameStats
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
