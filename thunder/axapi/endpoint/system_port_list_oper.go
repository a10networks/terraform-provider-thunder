

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemPortListOper struct {
    
    Oper SystemPortListOperOper `json:"oper"`

}
type DataSystemPortListOper struct {
    DtSystemPortListOper SystemPortListOper `json:"port-list"`
}


type SystemPortListOperOper struct {
    SystemPortList []SystemPortListOperOperSystemPortList `json:"system-port-list"`
}


type SystemPortListOperOperSystemPortList struct {
    PortIdx string `json:"Port-IDX"`
    PortNum string `json:"Port-Num"`
    Status string `json:"status"`
    MacAddr string `json:"Mac-Addr"`
    Speed string `json:"Speed"`
    Node string `json:"Node"`
    PciAddr string `json:"PCI-Addr"`
}

func (p *SystemPortListOper) GetId() string{
    return "1"
}

func (p *SystemPortListOper) getPath() string{
    return "system/port-list/oper"
}

func (p *SystemPortListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemPortListOper,error) {
logger.Println("SystemPortListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemPortListOper
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
