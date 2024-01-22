

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbServiceGroupOper struct {
    
    Oper GslbServiceGroupOperOper `json:"oper"`
    ServiceGroupName string `json:"service-group-name"`

}
type DataGslbServiceGroupOper struct {
    DtGslbServiceGroupOper GslbServiceGroupOper `json:"service-group"`
}


type GslbServiceGroupOperOper struct {
    SessionList []GslbServiceGroupOperOperSessionList `json:"session-list"`
    Matched int `json:"matched"`
    TotalSessions int `json:"total-sessions"`
}


type GslbServiceGroupOperOperSessionList struct {
    Client string `json:"client"`
    Best string `json:"best"`
    Mode string `json:"mode"`
    Hits int `json:"hits"`
    LastSecondHits int `json:"last-second-hits"`
    Ttl string `json:"ttl"`
    Update int `json:"update"`
    Aging int `json:"aging"`
}

func (p *GslbServiceGroupOper) GetId() string{
    return "1"
}

func (p *GslbServiceGroupOper) getPath() string{
    
    return "gslb/service-group/"+p.ServiceGroupName+"/oper"
}

func (p *GslbServiceGroupOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbServiceGroupOper,error) {
logger.Println("GslbServiceGroupOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbServiceGroupOper
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
