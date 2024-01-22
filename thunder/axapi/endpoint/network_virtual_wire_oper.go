

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkVirtualWireOper struct {
    
    Oper NetworkVirtualWireOperOper `json:"oper"`
    VirtualWireId int `json:"virtual-wire-id"`

}
type DataNetworkVirtualWireOper struct {
    DtNetworkVirtualWireOper NetworkVirtualWireOper `json:"virtual-wire"`
}


type NetworkVirtualWireOperOper struct {
    Endpoints []NetworkVirtualWireOperOperEndpoints `json:"endpoints"`
    Virtual_wire_status string `json:"virtual_wire_status"`
}


type NetworkVirtualWireOperOperEndpoints struct {
    Endpoint_type string `json:"endpoint_type"`
    Endpoint_intf int `json:"endpoint_intf"`
    Input_packet int `json:"input_packet"`
    Input_byte int `json:"input_byte"`
    Output_packet int `json:"output_packet"`
    Output_byte int `json:"output_byte"`
    Drop_packet int `json:"drop_packet"`
}

func (p *NetworkVirtualWireOper) GetId() string{
    return "1"
}

func (p *NetworkVirtualWireOper) getPath() string{
    
    return "network/virtual-wire/" +strconv.Itoa(p.VirtualWireId)+"/oper"
}

func (p *NetworkVirtualWireOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVirtualWireOper,error) {
logger.Println("NetworkVirtualWireOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkVirtualWireOper
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
