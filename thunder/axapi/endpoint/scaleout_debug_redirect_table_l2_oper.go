

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugRedirectTableL2Oper struct {
    
    Oper ScaleoutDebugRedirectTableL2OperOper `json:"oper"`

}
type DataScaleoutDebugRedirectTableL2Oper struct {
    DtScaleoutDebugRedirectTableL2Oper ScaleoutDebugRedirectTableL2Oper `json:"l2"`
}


type ScaleoutDebugRedirectTableL2OperOper struct {
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Redirect_info_list []ScaleoutDebugRedirectTableL2OperOperRedirect_info_list `json:"redirect_info_list"`
    Link_status int `json:"link_status"`
}


type ScaleoutDebugRedirectTableL2OperOperRedirect_info_list struct {
    Node int `json:"node"`
    Valid int `json:"valid"`
    Reachable int `json:"reachable"`
    Intf_num int `json:"intf_num"`
    Vlan int `json:"vlan"`
    Ip_addr_str string `json:"ip_addr_str"`
    Mac_str string `json:"mac_str"`
}

func (p *ScaleoutDebugRedirectTableL2Oper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugRedirectTableL2Oper) getPath() string{
    return "scaleout/debug/redirect-table/l2/oper"
}

func (p *ScaleoutDebugRedirectTableL2Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugRedirectTableL2Oper,error) {
logger.Println("ScaleoutDebugRedirectTableL2Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugRedirectTableL2Oper
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
