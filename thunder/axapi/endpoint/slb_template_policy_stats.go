

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyStats struct {
    
    Name string `json:"name"`
    Stats SlbTemplatePolicyStatsStats `json:"stats"`

}
type DataSlbTemplatePolicyStats struct {
    DtSlbTemplatePolicyStats SlbTemplatePolicyStats `json:"policy"`
}


type SlbTemplatePolicyStatsStats struct {
    FwdPolicyDnsUnresolved int `json:"fwd-policy-dns-unresolved"`
    FwdPolicyDnsOutstanding int `json:"fwd-policy-dns-outstanding"`
    FwdPolicySnatFail int `json:"fwd-policy-snat-fail"`
    FwdPolicyHits int `json:"fwd-policy-hits"`
    FwdPolicyForwardToInternet int `json:"fwd-policy-forward-to-internet"`
    FwdPolicyForwardToServiceGroup int `json:"fwd-policy-forward-to-service-group"`
    FwdPolicyForwardToProxy int `json:"fwd-policy-forward-to-proxy"`
    FwdPolicyPolicyDrop int `json:"fwd-policy-policy-drop"`
    FwdPolicySourceMatchNotFound int `json:"fwd-policy-source-match-not-found"`
    ExpClientHelloNotFound int `json:"exp-client-hello-not-found"`
}

func (p *SlbTemplatePolicyStats) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyStats) getPath() string{
    
    return "slb/template/policy/"+p.Name+"/stats"
}

func (p *SlbTemplatePolicyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyStats,error) {
logger.Println("SlbTemplatePolicyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplatePolicyStats
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
