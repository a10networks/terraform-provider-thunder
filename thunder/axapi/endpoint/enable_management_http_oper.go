

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementHttpOper struct {
    
    Oper EnableManagementHttpOperOper `json:"oper"`

}
type DataEnableManagementHttpOper struct {
    DtEnableManagementHttpOper EnableManagementHttpOper `json:"http"`
}


type EnableManagementHttpOperOper struct {
    PortList []EnableManagementHttpOperOperPortList `json:"port-list"`
}


type EnableManagementHttpOperOperPortList struct {
    Management int `json:"management"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
    Action string `json:"action"`
    Ipv4Acl string `json:"ipv4-acl"`
    Ipv6Acl string `json:"ipv6-acl"`
}

func (p *EnableManagementHttpOper) GetId() string{
    return "1"
}

func (p *EnableManagementHttpOper) getPath() string{
    return "enable-management/http/oper"
}

func (p *EnableManagementHttpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEnableManagementHttpOper,error) {
logger.Println("EnableManagementHttpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEnableManagementHttpOper
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
