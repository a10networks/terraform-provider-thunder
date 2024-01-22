

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDomainListNameStats struct {
    
    NameDomainList string `json:"name-domain-list"`
    Stats Cgnv6LsnRuleListDomainListNameStatsStats `json:"stats"`
    Name string 

}
type DataCgnv6LsnRuleListDomainListNameStats struct {
    DtCgnv6LsnRuleListDomainListNameStats Cgnv6LsnRuleListDomainListNameStats `json:"domain-list-name"`
}


type Cgnv6LsnRuleListDomainListNameStatsStats struct {
}

func (p *Cgnv6LsnRuleListDomainListNameStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListDomainListNameStats) getPath() string{
    
    return "cgnv6/lsn-rule-list/" +p.Name + "/domain-list-name/"+p.NameDomainList+"/stats"
}

func (p *Cgnv6LsnRuleListDomainListNameStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRuleListDomainListNameStats,error) {
logger.Println("Cgnv6LsnRuleListDomainListNameStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRuleListDomainListNameStats
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
