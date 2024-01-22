

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementHttpsOper struct {
    
    Oper EnableManagementHttpsOperOper `json:"oper"`

}
type DataEnableManagementHttpsOper struct {
    DtEnableManagementHttpsOper EnableManagementHttpsOper `json:"https"`
}


type EnableManagementHttpsOperOper struct {
    PortList []EnableManagementHttpsOperOperPortList `json:"port-list"`
}


type EnableManagementHttpsOperOperPortList struct {
    Management int `json:"management"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
    Action string `json:"action"`
    Ipv4Acl string `json:"ipv4-acl"`
    Ipv6Acl string `json:"ipv6-acl"`
}

func (p *EnableManagementHttpsOper) GetId() string{
    return "1"
}

func (p *EnableManagementHttpsOper) getPath() string{
    return "enable-management/https/oper"
}

func (p *EnableManagementHttpsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEnableManagementHttpsOper,error) {
logger.Println("EnableManagementHttpsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEnableManagementHttpsOper
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
