

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemPortInfoOper struct {
    
    Oper SystemPortInfoOperOper `json:"oper"`

}
type DataSystemPortInfoOper struct {
    DtSystemPortInfoOper SystemPortInfoOper `json:"port-info"`
}


type SystemPortInfoOperOper struct {
    SystemPortInfo []SystemPortInfoOperOperSystemPortInfo `json:"system-port-info"`
}


type SystemPortInfoOperOperSystemPortInfo struct {
    PortNum string `json:"Port-Num"`
    PciAddr string `json:"Pci-Addr"`
    Dev string `json:"Dev"`
    Info string `json:"Info"`
}

func (p *SystemPortInfoOper) GetId() string{
    return "1"
}

func (p *SystemPortInfoOper) getPath() string{
    return "system/port-info/oper"
}

func (p *SystemPortInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemPortInfoOper,error) {
logger.Println("SystemPortInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemPortInfoOper
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
