

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6FibOper struct {
    
    Oper Ipv6FibOperOper `json:"oper"`

}
type DataIpv6FibOper struct {
    DtIpv6FibOper Ipv6FibOper `json:"fib"`
}


type Ipv6FibOperOper struct {
    Total_routes int `json:"Total_Routes"`
    Total_paths int `json:"Total_Paths"`
    Ipv6Fib []Ipv6FibOperOperIpv6Fib `json:"IPv6-fib"`
}


type Ipv6FibOperOperIpv6Fib struct {
    Prefix string `json:"Prefix"`
    Prefixlen int `json:"PrefixLen"`
    Nexthop string `json:"Nexthop"`
    Interface string `json:"Interface"`
    Distance int `json:"Distance"`
}

func (p *Ipv6FibOper) GetId() string{
    return "1"
}

func (p *Ipv6FibOper) getPath() string{
    return "ipv6/fib/oper"
}

func (p *Ipv6FibOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpv6FibOper,error) {
logger.Println("Ipv6FibOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpv6FibOper
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
