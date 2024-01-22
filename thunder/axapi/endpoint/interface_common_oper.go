

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceCommonOper struct {
    
    Oper InterfaceCommonOperOper `json:"oper"`

}
type DataInterfaceCommonOper struct {
    DtInterfaceCommonOper InterfaceCommonOper `json:"common"`
}


type InterfaceCommonOperOper struct {
    Interfaces []InterfaceCommonOperOperInterfaces `json:"interfaces"`
    Time string `json:"time"`
    Interval int `json:"interval"`
    Vnp_id int `json:"vnp_id"`
    Tot_num_phy_intf int `json:"tot_num_phy_intf"`
}


type InterfaceCommonOperOperInterfaces struct {
    PortNum string `json:"port-num"`
    PortType string `json:"port-type"`
    Is_part_default_vlan string `json:"is_part_default_vlan"`
    SpanTree string `json:"span-tree"`
    RatePktSent int `json:"rate-pkt-sent"`
    RateByteSent int `json:"rate-byte-sent"`
    RatePktRcvd int `json:"rate-pkt-rcvd"`
    RateByteRcvd int `json:"rate-byte-rcvd"`
    InputUtilization int `json:"input-utilization"`
    OutputUtilization int `json:"output-utilization"`
    Type_vendor_part string `json:"type_vendor_part"`
    Total_lane int `json:"total_lane"`
    Transceivers_info []InterfaceCommonOperOperInterfacesTransceivers_info `json:"transceivers_info"`
}


type InterfaceCommonOperOperInterfacesTransceivers_info struct {
    TransceiverType string `json:"transceiver-type"`
    Lane int `json:"lane"`
    Curr string `json:"curr"`
    Hi_alarm string `json:"hi_alarm"`
    Hi_warn string `json:"hi_warn"`
    Lo_warn string `json:"lo_warn"`
    Lo_alarm string `json:"lo_alarm"`
}

func (p *InterfaceCommonOper) GetId() string{
    return "1"
}

func (p *InterfaceCommonOper) getPath() string{
    return "interface/common/oper"
}

func (p *InterfaceCommonOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceCommonOper,error) {
logger.Println("InterfaceCommonOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceCommonOper
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
