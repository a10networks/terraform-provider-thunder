

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpFibOper struct {
    
    Oper IpFibOperOper `json:"oper"`

}
type DataIpFibOper struct {
    DtIpFibOper IpFibOper `json:"fib"`
}


type IpFibOperOper struct {
    Total_routes int `json:"Total_Routes"`
    Total_paths int `json:"Total_Paths"`
    Ipv4Fib []IpFibOperOperIpv4Fib `json:"IPv4-fib"`
}


type IpFibOperOperIpv4Fib struct {
    Prefix string `json:"Prefix"`
    Prefixlen int `json:"PrefixLen"`
    Nexthop string `json:"Nexthop"`
    Interface string `json:"Interface"`
    Distance int `json:"Distance"`
}

func (p *IpFibOper) GetId() string{
    return "1"
}

func (p *IpFibOper) getPath() string{
    return "ip/fib/oper"
}

func (p *IpFibOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpFibOper,error) {
logger.Println("IpFibOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpFibOper
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
