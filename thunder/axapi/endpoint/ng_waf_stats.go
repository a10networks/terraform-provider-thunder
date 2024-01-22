

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type NgWafStats struct {
    
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

}
type DataNgWafStats struct {
    DtNgWafStats NgWafStats `json:"stats"`
}

func (p *NgWafStats) GetId() string{
    
    return url.QueryEscape(p.Name)
}

func (p *NgWafStats) getPath() string{
    return "ng-waf/stats"
}

func (p *NgWafStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNgWafStats,error) {
logger.Println("NgWafStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNgWafStats
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
