

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugIpv6RedirectTableL2Oper struct {
    
    Oper ScaleoutDebugIpv6RedirectTableL2OperOper `json:"oper"`

}
type DataScaleoutDebugIpv6RedirectTableL2Oper struct {
    DtScaleoutDebugIpv6RedirectTableL2Oper ScaleoutDebugIpv6RedirectTableL2Oper `json:"l2"`
}


type ScaleoutDebugIpv6RedirectTableL2OperOper struct {
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Redirect_info_list_v6 []ScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6 `json:"redirect_info_list_v6"`
    Link_status int `json:"link_status"`
}


type ScaleoutDebugIpv6RedirectTableL2OperOperRedirect_info_list_v6 struct {
    Node int `json:"node"`
    Valid int `json:"valid"`
    Reachable int `json:"reachable"`
    Intf_num int `json:"intf_num"`
    Vlan int `json:"vlan"`
    Ipv6_addr_str string `json:"ipv6_addr_str"`
    Mac_str string `json:"mac_str"`
}

func (p *ScaleoutDebugIpv6RedirectTableL2Oper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugIpv6RedirectTableL2Oper) getPath() string{
    return "scaleout/debug/ipv6/redirect-table/l2/oper"
}

func (p *ScaleoutDebugIpv6RedirectTableL2Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugIpv6RedirectTableL2Oper,error) {
logger.Println("ScaleoutDebugIpv6RedirectTableL2Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugIpv6RedirectTableL2Oper
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
