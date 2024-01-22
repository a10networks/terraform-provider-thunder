

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugIpv6ReachabilityOper struct {
    
    Oper ScaleoutDebugIpv6ReachabilityOperOper `json:"oper"`

}
type DataScaleoutDebugIpv6ReachabilityOper struct {
    DtScaleoutDebugIpv6ReachabilityOper ScaleoutDebugIpv6ReachabilityOper `json:"reachability"`
}


type ScaleoutDebugIpv6ReachabilityOperOper struct {
    Part_name string `json:"part_name"`
    ScaleoutIpv6List []ScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List `json:"scaleout-ipv6-list"`
}


type ScaleoutDebugIpv6ReachabilityOperOperScaleoutIpv6List struct {
    Node int `json:"node"`
    Ipv6_addr string `json:"ipv6_addr"`
    Mac string `json:"mac"`
    Vnp_id int `json:"vnp_id"`
    Vlan_id int `json:"vlan_id"`
    Prefix_len int `json:"prefix_len"`
    Real_port int `json:"real_port"`
    Name string `json:"name"`
}

func (p *ScaleoutDebugIpv6ReachabilityOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugIpv6ReachabilityOper) getPath() string{
    return "scaleout/debug/ipv6/reachability/oper"
}

func (p *ScaleoutDebugIpv6ReachabilityOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugIpv6ReachabilityOper,error) {
logger.Println("ScaleoutDebugIpv6ReachabilityOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugIpv6ReachabilityOper
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
