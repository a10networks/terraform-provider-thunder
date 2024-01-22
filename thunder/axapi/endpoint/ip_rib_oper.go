

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpRibOper struct {
    
    Oper IpRibOperOper `json:"oper"`

}
type DataIpRibOper struct {
    DtIpRibOper IpRibOper `json:"rib"`
}


type IpRibOperOper struct {
    Description string `json:"Description"`
    Total int `json:"Total"`
    TotalPaths int `json:"Total-Paths"`
    Limit int `json:"Limit"`
    Ipv4Routes []IpRibOperOperIpv4Routes `json:"IPv4-routes"`
}


type IpRibOperOperIpv4Routes struct {
    Prefix string `json:"Prefix"`
    Prefixlen int `json:"PrefixLen"`
    Type string `json:"Type"`
    Subtype string `json:"Subtype"`
    Distance int `json:"Distance"`
    Metric int `json:"Metric"`
    Nexthop string `json:"Nexthop"`
    Interface string `json:"Interface"`
}

func (p *IpRibOper) GetId() string{
    return "1"
}

func (p *IpRibOper) getPath() string{
    return "ip/rib/oper"
}

func (p *IpRibOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpRibOper,error) {
logger.Println("IpRibOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpRibOper
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
