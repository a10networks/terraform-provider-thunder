

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementSnmpOper struct {
    
    Oper EnableManagementSnmpOperOper `json:"oper"`

}
type DataEnableManagementSnmpOper struct {
    DtEnableManagementSnmpOper EnableManagementSnmpOper `json:"snmp"`
}


type EnableManagementSnmpOperOper struct {
    PortList []EnableManagementSnmpOperOperPortList `json:"port-list"`
}


type EnableManagementSnmpOperOperPortList struct {
    Management int `json:"management"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
    Action string `json:"action"`
    Ipv4Acl string `json:"ipv4-acl"`
    Ipv6Acl string `json:"ipv6-acl"`
}

func (p *EnableManagementSnmpOper) GetId() string{
    return "1"
}

func (p *EnableManagementSnmpOper) getPath() string{
    return "enable-management/snmp/oper"
}

func (p *EnableManagementSnmpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEnableManagementSnmpOper,error) {
logger.Println("EnableManagementSnmpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEnableManagementSnmpOper
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
