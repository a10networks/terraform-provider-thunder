

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceDriverStatisticOper struct {
    
    Oper InterfaceDriverStatisticOperOper `json:"oper"`

}
type DataInterfaceDriverStatisticOper struct {
    DtInterfaceDriverStatisticOper InterfaceDriverStatisticOper `json:"driver-statistic"`
}


type InterfaceDriverStatisticOperOper struct {
    Interfaces []InterfaceDriverStatisticOperOperInterfaces `json:"interfaces"`
}


type InterfaceDriverStatisticOperOperInterfaces struct {
    IfType string `json:"IF-Type"`
    Port_num int `json:"port_num"`
    Rxpkts64_counts int `json:"rxpkts64_counts"`
    Rxpkts65to127_counts int `json:"rxpkts65to127_counts"`
    Rxpkts128to255_counts int `json:"rxpkts128to255_counts"`
    Rxpkts256to511_counts int `json:"rxpkts256to511_counts"`
    Rxpkts512to1023_counts int `json:"rxpkts512to1023_counts"`
    Rxpkts1024to1518_counts int `json:"rxpkts1024to1518_counts"`
    Rxpkts1519tomax_counts int `json:"rxpkts1519toMax_counts"`
    Txpkts64_counts int `json:"txpkts64_counts"`
    Txpkts65to127_counts int `json:"txpkts65to127_counts"`
    Txpkts128to255_counts int `json:"txpkts128to255_counts"`
    Txpkts256to511_counts int `json:"txpkts256to511_counts"`
    Txpkts512to1023_counts int `json:"txpkts512to1023_counts"`
    Txpkts1024to1518_counts int `json:"txpkts1024to1518_counts"`
    Txpkts1519tomax_counts int `json:"txpkts1519toMax_counts"`
}

func (p *InterfaceDriverStatisticOper) GetId() string{
    return "1"
}

func (p *InterfaceDriverStatisticOper) getPath() string{
    return "interface/driver-statistic/oper"
}

func (p *InterfaceDriverStatisticOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceDriverStatisticOper,error) {
logger.Println("InterfaceDriverStatisticOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceDriverStatisticOper
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
