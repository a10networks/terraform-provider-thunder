

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceManagementOper struct {
    
    Oper InterfaceManagementOperOper `json:"oper"`

}
type DataInterfaceManagementOper struct {
    DtInterfaceManagementOper InterfaceManagementOper `json:"management"`
}


type InterfaceManagementOperOper struct {
    Interface string `json:"interface"`
    State int `json:"state"`
    Line_protocol string `json:"line_protocol"`
    Link_type string `json:"link_type"`
    Mac string `json:"mac"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv4Mask string `json:"ipv4-mask"`
    Ipv4DefaultGateway string `json:"ipv4-default-gateway"`
    Ipv6Addr string `json:"ipv6-addr"`
    Ipv6Prefix string `json:"ipv6-prefix"`
    Ipv6LinkLocal string `json:"ipv6-link-local"`
    Ipv6LinkLocalPrefix string `json:"ipv6-link-local-prefix"`
    Ipv6DefaultGateway string `json:"ipv6-default-gateway"`
    Speed string `json:"speed"`
    Duplexity string `json:"duplexity"`
    Mtu int `json:"mtu"`
    Flow_control int `json:"flow_control"`
    Ipv4_acl string `json:"ipv4_acl"`
    Ipv6_acl string `json:"ipv6_acl"`
    Dhcp_enabled int `json:"dhcp_enabled"`
}

func (p *InterfaceManagementOper) GetId() string{
    return "1"
}

func (p *InterfaceManagementOper) getPath() string{
    return "interface/management/oper"
}

func (p *InterfaceManagementOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceManagementOper,error) {
logger.Println("InterfaceManagementOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceManagementOper
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
