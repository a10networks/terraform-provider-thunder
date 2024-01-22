

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpASetidMonitorOper struct {
    
    Oper VrrpASetidMonitorOperOper `json:"oper"`

}
type DataVrrpASetidMonitorOper struct {
    DtVrrpASetidMonitorOper VrrpASetidMonitorOper `json:"setid-monitor"`
}


type VrrpASetidMonitorOperOper struct {
    SetId int `json:"set-id"`
    SetidList []VrrpASetidMonitorOperOperSetidList `json:"setid-list"`
}


type VrrpASetidMonitorOperOperSetidList struct {
    Interface_type string `json:"interface_type"`
    Lif int `json:"lif"`
    IpList []VrrpASetidMonitorOperOperSetidListIpList `json:"ip-list"`
}


type VrrpASetidMonitorOperOperSetidListIpList struct {
    DeviceId int `json:"device-id"`
    Ipaddress string `json:"ipaddress"`
    Ipv6 string `json:"ipv6"`
    Mac string `json:"mac"`
    Vlan int `json:"vlan"`
    Last_observed int `json:"last_observed"`
}

func (p *VrrpASetidMonitorOper) GetId() string{
    return "1"
}

func (p *VrrpASetidMonitorOper) getPath() string{
    return "vrrp-a/setid-monitor/oper"
}

func (p *VrrpASetidMonitorOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpASetidMonitorOper,error) {
logger.Println("VrrpASetidMonitorOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpASetidMonitorOper
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
