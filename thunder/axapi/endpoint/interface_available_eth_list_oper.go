

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceAvailableEthListOper struct {
    
    Oper InterfaceAvailableEthListOperOper `json:"oper"`

}
type DataInterfaceAvailableEthListOper struct {
    DtInterfaceAvailableEthListOper InterfaceAvailableEthListOper `json:"available-eth-list"`
}


type InterfaceAvailableEthListOperOper struct {
    Tot_num_of_ports int `json:"tot_num_of_ports"`
    IfList []InterfaceAvailableEthListOperOperIfList `json:"if-list"`
}


type InterfaceAvailableEthListOperOperIfList struct {
    IfType string `json:"IF-Type"`
    IfNum int `json:"IF-Num"`
    IfStatus string `json:"IF-Status"`
    State string `json:"state"`
}

func (p *InterfaceAvailableEthListOper) GetId() string{
    return "1"
}

func (p *InterfaceAvailableEthListOper) getPath() string{
    return "interface/available-eth-list/oper"
}

func (p *InterfaceAvailableEthListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceAvailableEthListOper,error) {
logger.Println("InterfaceAvailableEthListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceAvailableEthListOper
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
