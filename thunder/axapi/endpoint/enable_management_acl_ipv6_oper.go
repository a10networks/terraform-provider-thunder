

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementAclIpv6Oper struct {
    
    Oper EnableManagementAclIpv6OperOper `json:"oper"`

}
type DataEnableManagementAclIpv6Oper struct {
    DtEnableManagementAclIpv6Oper EnableManagementAclIpv6Oper `json:"acl-ipv6"`
}


type EnableManagementAclIpv6OperOper struct {
    PortList []EnableManagementAclIpv6OperOperPortList `json:"port-list"`
}


type EnableManagementAclIpv6OperOperPortList struct {
    Management int `json:"management"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
    Action string `json:"action"`
}

func (p *EnableManagementAclIpv6Oper) GetId() string{
    return "1"
}

func (p *EnableManagementAclIpv6Oper) getPath() string{
    return "enable-management/acl-ipv6/oper"
}

func (p *EnableManagementAclIpv6Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEnableManagementAclIpv6Oper,error) {
logger.Println("EnableManagementAclIpv6Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEnableManagementAclIpv6Oper
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
