

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6NeighborDynamicOper struct {
    
    Oper Ipv6NeighborDynamicOperOper `json:"oper"`

}
type DataIpv6NeighborDynamicOper struct {
    DtIpv6NeighborDynamicOper Ipv6NeighborDynamicOper `json:"dynamic"`
}


type Ipv6NeighborDynamicOperOper struct {
    V6neighborList []Ipv6NeighborDynamicOperOperV6neighborList `json:"v6neighbor-list"`
}


type Ipv6NeighborDynamicOperOperV6neighborList struct {
    Ipv6Address string `json:"IPV6-Address"`
    MacAddress string `json:"MAC-Address"`
    Type string `json:"Type"`
    Age int `json:"Age"`
    State string `json:"State"`
    Interface string `json:"Interface"`
    Vlan int `json:"Vlan"`
}

func (p *Ipv6NeighborDynamicOper) GetId() string{
    return "1"
}

func (p *Ipv6NeighborDynamicOper) getPath() string{
    return "ipv6/neighbor/dynamic/oper"
}

func (p *Ipv6NeighborDynamicOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpv6NeighborDynamicOper,error) {
logger.Println("Ipv6NeighborDynamicOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpv6NeighborDynamicOper
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
