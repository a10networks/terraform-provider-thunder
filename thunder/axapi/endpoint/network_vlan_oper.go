

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkVlanOper struct {
    
    Oper NetworkVlanOperOper `json:"oper"`
    VlanNum int `json:"vlan-num"`

}
type DataNetworkVlanOper struct {
    DtNetworkVlanOper NetworkVlanOper `json:"vlan"`
}


type NetworkVlanOperOper struct {
    Vlan_name string `json:"vlan_name"`
    Ve_num int `json:"ve_num"`
    Is_shared_vlan string `json:"is_shared_vlan"`
    Un_tagg_eth_ports NetworkVlanOperOperUn_tagg_eth_ports `json:"un_tagg_eth_ports"`
    Tagg_eth_ports NetworkVlanOperOperTagg_eth_ports `json:"tagg_eth_ports"`
    Un_tagg_logical_ports NetworkVlanOperOperUn_tagg_logical_ports `json:"un_tagg_logical_ports"`
    Tagg_logical_ports NetworkVlanOperOperTagg_logical_ports `json:"tagg_logical_ports"`
    SpanTree string `json:"span-tree"`
}


type NetworkVlanOperOperUn_tagg_eth_ports struct {
    Ports int `json:"ports"`
}


type NetworkVlanOperOperTagg_eth_ports struct {
    Ports int `json:"ports"`
}


type NetworkVlanOperOperUn_tagg_logical_ports struct {
    Ports int `json:"ports"`
}


type NetworkVlanOperOperTagg_logical_ports struct {
    Ports int `json:"ports"`
}

func (p *NetworkVlanOper) GetId() string{
    return "1"
}

func (p *NetworkVlanOper) getPath() string{
    
    return "network/vlan/" +strconv.Itoa(p.VlanNum)+"/oper"
}

func (p *NetworkVlanOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVlanOper,error) {
logger.Println("NetworkVlanOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkVlanOper
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
