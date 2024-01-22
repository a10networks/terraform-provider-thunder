

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbPolicyOper struct {
    
    Name string `json:"name"`
    Oper GslbPolicyOperOper `json:"oper"`

}
type DataGslbPolicyOper struct {
    DtGslbPolicyOper GslbPolicyOper `json:"policy"`
}


type GslbPolicyOperOper struct {
    MetricList []GslbPolicyOperOperMetricList `json:"metric-list"`
}


type GslbPolicyOperOperMetricList struct {
    Type string `json:"type"`
    Order int `json:"order"`
}

func (p *GslbPolicyOper) GetId() string{
    return "1"
}

func (p *GslbPolicyOper) getPath() string{
    
    return "gslb/policy/"+p.Name+"/oper"
}

func (p *GslbPolicyOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbPolicyOper,error) {
logger.Println("GslbPolicyOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbPolicyOper
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
