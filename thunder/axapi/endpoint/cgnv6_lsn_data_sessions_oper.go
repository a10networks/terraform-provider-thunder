

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnDataSessionsOper struct {
    
    Oper Cgnv6LsnDataSessionsOperOper `json:"oper"`

}
type DataCgnv6LsnDataSessionsOper struct {
    DtCgnv6LsnDataSessionsOper Cgnv6LsnDataSessionsOper `json:"data-sessions"`
}


type Cgnv6LsnDataSessionsOperOper struct {
    Status string `json:"status"`
    InsideAddr string `json:"inside-addr"`
    InsideAddrStart string `json:"inside-addr-start"`
    InsideAddrEnd string `json:"inside-addr-end"`
    NatAddr string `json:"nat-addr"`
    NatAddrStart string `json:"nat-addr-start"`
    NatAddrEnd string `json:"nat-addr-end"`
    NatPort int `json:"nat-port"`
    InsidePort int `json:"inside-port"`
}

func (p *Cgnv6LsnDataSessionsOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnDataSessionsOper) getPath() string{
    return "cgnv6/lsn/data-sessions/oper"
}

func (p *Cgnv6LsnDataSessionsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnDataSessionsOper,error) {
logger.Println("Cgnv6LsnDataSessionsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnDataSessionsOper
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
