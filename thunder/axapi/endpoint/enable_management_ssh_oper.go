

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementSshOper struct {
    
    Oper EnableManagementSshOperOper `json:"oper"`

}
type DataEnableManagementSshOper struct {
    DtEnableManagementSshOper EnableManagementSshOper `json:"ssh"`
}


type EnableManagementSshOperOper struct {
    PortList []EnableManagementSshOperOperPortList `json:"port-list"`
}


type EnableManagementSshOperOperPortList struct {
    Management int `json:"management"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
    Action string `json:"action"`
    Ipv4Acl string `json:"ipv4-acl"`
    Ipv6Acl string `json:"ipv6-acl"`
}

func (p *EnableManagementSshOper) GetId() string{
    return "1"
}

func (p *EnableManagementSshOper) getPath() string{
    return "enable-management/ssh/oper"
}

func (p *EnableManagementSshOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEnableManagementSshOper,error) {
logger.Println("EnableManagementSshOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEnableManagementSshOper
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
