

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceSwitchDropReasonOper struct {
    
    Oper InterfaceSwitchDropReasonOperOper `json:"oper"`

}
type DataInterfaceSwitchDropReasonOper struct {
    DtInterfaceSwitchDropReasonOper InterfaceSwitchDropReasonOper `json:"switch-drop-reason"`
}


type InterfaceSwitchDropReasonOperOper struct {
    Interfaces []InterfaceSwitchDropReasonOperOperInterfaces `json:"interfaces"`
}


type InterfaceSwitchDropReasonOperOperInterfaces struct {
    IfType string `json:"IF-Type"`
    Port_num int `json:"port_num"`
    Ifinconddrop int `json:"IfInCondDrop"`
    Ifinvlandrop int `json:"IfInVlanDrop"`
    Ifinbitmapzerodrop int `json:"IfInBitmapZeroDrop"`
    Ifinfilterdrop int `json:"IfInFilterDrop"`
    Ifinportdrop int `json:"IfInPortDrop"`
    Ifinipv4drop int `json:"IfInIpv4Drop"`
    Ifinipv6drop int `json:"IfInIpv6Drop"`
    Ifinmcdrop int `json:"IfInMcDrop"`
    Ifinqueuedrop int `json:"IfInQueueDrop"`
    Ifinerrordrop int `json:"IfInErrorDrop"`
    Ifineventdrop int `json:"IfInEventDrop"`
    Ifinundersizedrop int `json:"IfInUnderSizeDrop"`
    Ifinmtuexceeddrop int `json:"IfInMTUExceedDrop"`
    Ifinfragdrop int `json:"IfInFragDrop"`
    Ifinjabberdrop int `json:"IfInJabberDrop"`
    Ifinfcsdrop int `json:"IfInFCSDrop"`
    Ifoutipv4drop int `json:"IfOutIpv4Drop"`
    Ifoutipv6drop int `json:"IfOutIpv6Drop"`
    Ifouttnldrop int `json:"IfOutTnlDrop"`
    Ifoutmcdrop int `json:"IfOutMcDrop"`
    Ifoutvlandrop int `json:"IfOutVlanDrop"`
    Ifoutconddrop int `json:"IfOutCondDrop"`
    Ifoutparitydrop int `json:"IfOutParityDrop"`
    Ifoutstgdrop int `json:"IfOutStgDrop"`
}

func (p *InterfaceSwitchDropReasonOper) GetId() string{
    return "1"
}

func (p *InterfaceSwitchDropReasonOper) getPath() string{
    return "interface/switch-drop-reason/oper"
}

func (p *InterfaceSwitchDropReasonOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceSwitchDropReasonOper,error) {
logger.Println("InterfaceSwitchDropReasonOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceSwitchDropReasonOper
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
