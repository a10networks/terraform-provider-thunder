

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemInusePortListOper struct {
    
    Oper SystemInusePortListOperOper `json:"oper"`

}
type DataSystemInusePortListOper struct {
    DtSystemInusePortListOper SystemInusePortListOper `json:"inuse-port-list"`
}


type SystemInusePortListOperOper struct {
    SystemInusePortList []SystemInusePortListOperOperSystemInusePortList `json:"system-inuse-port-list"`
}


type SystemInusePortListOperOperSystemInusePortList struct {
    PortNum string `json:"Port-Num"`
    Status string `json:"status"`
    MacAddr string `json:"Mac-Addr"`
    Speed string `json:"Speed"`
    Node string `json:"Node"`
    PciAddr string `json:"PCI-Addr"`
}

func (p *SystemInusePortListOper) GetId() string{
    return "1"
}

func (p *SystemInusePortListOper) getPath() string{
    return "system/inuse-port-list/oper"
}

func (p *SystemInusePortListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemInusePortListOper,error) {
logger.Println("SystemInusePortListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemInusePortListOper
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
