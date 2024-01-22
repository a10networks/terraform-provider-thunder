

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryReputationScopeStats struct {
    
    Name string `json:"name"`
    Stats WebCategoryReputationScopeStatsStats `json:"stats"`

}
type DataWebCategoryReputationScopeStats struct {
    DtWebCategoryReputationScopeStats WebCategoryReputationScopeStats `json:"reputation-scope"`
}


type WebCategoryReputationScopeStatsStats struct {
    Trustworthy int `json:"trustworthy"`
    LowRisk int `json:"low-risk"`
    ModerateRisk int `json:"moderate-risk"`
    Suspicious int `json:"suspicious"`
    Malicious int `json:"malicious"`
}

func (p *WebCategoryReputationScopeStats) GetId() string{
    return "1"
}

func (p *WebCategoryReputationScopeStats) getPath() string{
    
    return "web-category/reputation-scope/"+p.Name+"/stats"
}

func (p *WebCategoryReputationScopeStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryReputationScopeStats,error) {
logger.Println("WebCategoryReputationScopeStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryReputationScopeStats
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
