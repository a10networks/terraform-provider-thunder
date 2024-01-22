

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnPerformanceOper struct {
    
    Oper Cgnv6LsnPerformanceOperOper `json:"oper"`

}
type DataCgnv6LsnPerformanceOper struct {
    DtCgnv6LsnPerformanceOper Cgnv6LsnPerformanceOper `json:"performance"`
}


type Cgnv6LsnPerformanceOperOper struct {
    DataSessions int `json:"data-sessions"`
    FullConeSessions int `json:"full-cone-sessions"`
    UserQuotas int `json:"user-quotas"`
}

func (p *Cgnv6LsnPerformanceOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnPerformanceOper) getPath() string{
    return "cgnv6/lsn/performance/oper"
}

func (p *Cgnv6LsnPerformanceOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnPerformanceOper,error) {
logger.Println("Cgnv6LsnPerformanceOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnPerformanceOper
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
