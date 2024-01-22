

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkVirtualWireGlobalOper struct {
    
    Oper NetworkVirtualWireGlobalOperOper `json:"oper"`

}
type DataNetworkVirtualWireGlobalOper struct {
    DtNetworkVirtualWireGlobalOper NetworkVirtualWireGlobalOper `json:"virtual-wire-global"`
}


type NetworkVirtualWireGlobalOperOper struct {
    Vlan_group []NetworkVirtualWireGlobalOperOperVlan_group `json:"vlan_group"`
    Vlan_set []NetworkVirtualWireGlobalOperOperVlan_set `json:"vlan_set"`
}


type NetworkVirtualWireGlobalOperOperVlan_group struct {
    Group_id int `json:"group_id"`
    Active_member int `json:"active_member"`
}


type NetworkVirtualWireGlobalOperOperVlan_set struct {
    Set_id int `json:"set_id"`
    Active_pair int `json:"active_pair"`
}

func (p *NetworkVirtualWireGlobalOper) GetId() string{
    return "1"
}

func (p *NetworkVirtualWireGlobalOper) getPath() string{
    return "network/virtual-wire-global/oper"
}

func (p *NetworkVirtualWireGlobalOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVirtualWireGlobalOper,error) {
logger.Println("NetworkVirtualWireGlobalOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkVirtualWireGlobalOper
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
