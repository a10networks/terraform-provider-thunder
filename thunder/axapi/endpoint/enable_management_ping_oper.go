

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementPingOper struct {
    
    Oper EnableManagementPingOperOper `json:"oper"`

}
type DataEnableManagementPingOper struct {
    DtEnableManagementPingOper EnableManagementPingOper `json:"ping"`
}


type EnableManagementPingOperOper struct {
    PortList []EnableManagementPingOperOperPortList `json:"port-list"`
}


type EnableManagementPingOperOperPortList struct {
    Management int `json:"management"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
    Action string `json:"action"`
    Ipv4Acl string `json:"ipv4-acl"`
    Ipv6Acl string `json:"ipv6-acl"`
}

func (p *EnableManagementPingOper) GetId() string{
    return "1"
}

func (p *EnableManagementPingOper) getPath() string{
    return "enable-management/ping/oper"
}

func (p *EnableManagementPingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEnableManagementPingOper,error) {
logger.Println("EnableManagementPingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEnableManagementPingOper
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
