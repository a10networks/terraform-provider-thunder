

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAllSessionsOper struct {
    
    Oper Cgnv6LsnAllSessionsOperOper `json:"oper"`

}
type DataCgnv6LsnAllSessionsOper struct {
    DtCgnv6LsnAllSessionsOper Cgnv6LsnAllSessionsOper `json:"all-sessions"`
}


type Cgnv6LsnAllSessionsOperOper struct {
    Status string `json:"status"`
    NatPoolName string `json:"nat-pool-name"`
}

func (p *Cgnv6LsnAllSessionsOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAllSessionsOper) getPath() string{
    return "cgnv6/lsn/all-sessions/oper"
}

func (p *Cgnv6LsnAllSessionsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnAllSessionsOper,error) {
logger.Println("Cgnv6LsnAllSessionsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnAllSessionsOper
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
