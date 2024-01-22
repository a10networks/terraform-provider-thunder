

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6RibOper struct {
    
    Oper Ipv6RibOperOper `json:"oper"`

}
type DataIpv6RibOper struct {
    DtIpv6RibOper Ipv6RibOper `json:"rib"`
}


type Ipv6RibOperOper struct {
    Description string `json:"Description"`
    Total int `json:"Total"`
    TotalPaths int `json:"Total-Paths"`
    Limit int `json:"Limit"`
    Ipv6Routes []Ipv6RibOperOperIpv6Routes `json:"IPv6-routes"`
}


type Ipv6RibOperOperIpv6Routes struct {
    Prefix string `json:"Prefix"`
    Prefixlen int `json:"PrefixLen"`
    Type string `json:"Type"`
    Subtype string `json:"Subtype"`
    Distance int `json:"Distance"`
    Metric int `json:"Metric"`
    Nexthop string `json:"Nexthop"`
    Interface string `json:"Interface"`
}

func (p *Ipv6RibOper) GetId() string{
    return "1"
}

func (p *Ipv6RibOper) getPath() string{
    return "ipv6/rib/oper"
}

func (p *Ipv6RibOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpv6RibOper,error) {
logger.Println("Ipv6RibOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpv6RibOper
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
