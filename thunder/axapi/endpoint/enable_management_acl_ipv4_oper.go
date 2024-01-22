

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementAclIpv4Oper struct {
    
    Oper EnableManagementAclIpv4OperOper `json:"oper"`

}
type DataEnableManagementAclIpv4Oper struct {
    DtEnableManagementAclIpv4Oper EnableManagementAclIpv4Oper `json:"acl-ipv4"`
}


type EnableManagementAclIpv4OperOper struct {
    PortList []EnableManagementAclIpv4OperOperPortList `json:"port-list"`
}


type EnableManagementAclIpv4OperOperPortList struct {
    Management int `json:"management"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
    Action string `json:"action"`
}

func (p *EnableManagementAclIpv4Oper) GetId() string{
    return "1"
}

func (p *EnableManagementAclIpv4Oper) getPath() string{
    return "enable-management/acl-ipv4/oper"
}

func (p *EnableManagementAclIpv4Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEnableManagementAclIpv4Oper,error) {
logger.Println("EnableManagementAclIpv4Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEnableManagementAclIpv4Oper
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
