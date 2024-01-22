

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugReachabilityOper struct {
    
    Oper ScaleoutDebugReachabilityOperOper `json:"oper"`

}
type DataScaleoutDebugReachabilityOper struct {
    DtScaleoutDebugReachabilityOper ScaleoutDebugReachabilityOper `json:"reachability"`
}


type ScaleoutDebugReachabilityOperOper struct {
    Part_name string `json:"part_name"`
    ScaleoutIpList []ScaleoutDebugReachabilityOperOperScaleoutIpList `json:"scaleout-ip-list"`
}


type ScaleoutDebugReachabilityOperOperScaleoutIpList struct {
    Node int `json:"node"`
    Ip_addr string `json:"ip_addr"`
    Mac string `json:"mac"`
    Vnp_id int `json:"vnp_id"`
    Vlan_id int `json:"vlan_id"`
    Netmask int `json:"netmask"`
    Real_port int `json:"real_port"`
    Name string `json:"name"`
}

func (p *ScaleoutDebugReachabilityOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugReachabilityOper) getPath() string{
    return "scaleout/debug/reachability/oper"
}

func (p *ScaleoutDebugReachabilityOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugReachabilityOper,error) {
logger.Println("ScaleoutDebugReachabilityOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugReachabilityOper
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
