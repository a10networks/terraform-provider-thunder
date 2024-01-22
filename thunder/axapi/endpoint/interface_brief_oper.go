

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceBriefOper struct {
    
    Oper InterfaceBriefOperOper `json:"oper"`

}
type DataInterfaceBriefOper struct {
    DtInterfaceBriefOper InterfaceBriefOper `json:"brief"`
}


type InterfaceBriefOperOper struct {
    Interfaces []InterfaceBriefOperOperInterfaces `json:"interfaces"`
}


type InterfaceBriefOperOperInterfaces struct {
    Port_num string `json:"port_num"`
    Speed string `json:"speed"`
    Duplexity string `json:"duplexity"`
    Trunk_group string `json:"trunk_group"`
    Vlan_info string `json:"vlan_info"`
    Encap_info string `json:"encap_info"`
    State string `json:"state"`
    Mac string `json:"mac"`
    Ipv4_addr_count int `json:"ipv4_addr_count"`
    Ipv4_addr string `json:"ipv4_addr"`
    Ipv4_mask string `json:"ipv4_mask"`
    Ipv6_addr_count int `json:"ipv6_addr_count"`
    Ipv6_addr string `json:"ipv6_addr"`
    Ipv6_prefix string `json:"ipv6_prefix"`
    Intf_name string `json:"intf_name"`
    Unnumbered_oper string `json:"unnumbered_oper"`
}

func (p *InterfaceBriefOper) GetId() string{
    return "1"
}

func (p *InterfaceBriefOper) getPath() string{
    return "interface/brief/oper"
}

func (p *InterfaceBriefOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceBriefOper,error) {
logger.Println("InterfaceBriefOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceBriefOper
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
