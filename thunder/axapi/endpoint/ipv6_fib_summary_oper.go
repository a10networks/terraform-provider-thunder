

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6FibSummaryOper struct {
    
    Oper Ipv6FibSummaryOperOper `json:"oper"`

}
type DataIpv6FibSummaryOper struct {
    DtIpv6FibSummaryOper Ipv6FibSummaryOper `json:"fib-summary"`
}


type Ipv6FibSummaryOperOper struct {
    Connected_routes int `json:"connected_routes"`
    Static_dynamic_routes int `json:"static_dynamic_routes"`
    Total_routes int `json:"total_routes"`
    Static_dynamic_paths int `json:"static_dynamic_paths"`
    Total_paths int `json:"total_paths"`
}

func (p *Ipv6FibSummaryOper) GetId() string{
    return "1"
}

func (p *Ipv6FibSummaryOper) getPath() string{
    return "ipv6/fib-summary/oper"
}

func (p *Ipv6FibSummaryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpv6FibSummaryOper,error) {
logger.Println("Ipv6FibSummaryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpv6FibSummaryOper
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
