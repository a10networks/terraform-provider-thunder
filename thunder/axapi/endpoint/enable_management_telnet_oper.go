

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementTelnetOper struct {
    
    Oper EnableManagementTelnetOperOper `json:"oper"`

}
type DataEnableManagementTelnetOper struct {
    DtEnableManagementTelnetOper EnableManagementTelnetOper `json:"telnet"`
}


type EnableManagementTelnetOperOper struct {
    PortList []EnableManagementTelnetOperOperPortList `json:"port-list"`
}


type EnableManagementTelnetOperOperPortList struct {
    Management int `json:"management"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
    Action string `json:"action"`
    Ipv4Acl string `json:"ipv4-acl"`
    Ipv6Acl string `json:"ipv6-acl"`
}

func (p *EnableManagementTelnetOper) GetId() string{
    return "1"
}

func (p *EnableManagementTelnetOper) getPath() string{
    return "enable-management/telnet/oper"
}

func (p *EnableManagementTelnetOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEnableManagementTelnetOper,error) {
logger.Println("EnableManagementTelnetOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEnableManagementTelnetOper
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
