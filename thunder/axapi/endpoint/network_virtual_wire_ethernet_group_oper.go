

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkVirtualWireEthernetGroupOper struct {
    
    GroupId int `json:"group-id"`
    Oper NetworkVirtualWireEthernetGroupOperOper `json:"oper"`

}
type DataNetworkVirtualWireEthernetGroupOper struct {
    DtNetworkVirtualWireEthernetGroupOper NetworkVirtualWireEthernetGroupOper `json:"virtual-wire-ethernet-group"`
}


type NetworkVirtualWireEthernetGroupOperOper struct {
    Eth_member_status []NetworkVirtualWireEthernetGroupOperOperEth_member_status `json:"eth_member_status"`
    Trunk_member_status []NetworkVirtualWireEthernetGroupOperOperTrunk_member_status `json:"trunk_member_status"`
    Lead_port int `json:"lead_port"`
    Group_status string `json:"group_status"`
}


type NetworkVirtualWireEthernetGroupOperOperEth_member_status struct {
    Eth_member_num int `json:"eth_member_num"`
    Eth_member_status string `json:"eth_member_status"`
}


type NetworkVirtualWireEthernetGroupOperOperTrunk_member_status struct {
    Trunk_member_num int `json:"trunk_member_num"`
    Trunk_member_status string `json:"trunk_member_status"`
}

func (p *NetworkVirtualWireEthernetGroupOper) GetId() string{
    return "1"
}

func (p *NetworkVirtualWireEthernetGroupOper) getPath() string{
    
    return "network/virtual-wire-ethernet-group/" +strconv.Itoa(p.GroupId)+"/oper"
}

func (p *NetworkVirtualWireEthernetGroupOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVirtualWireEthernetGroupOper,error) {
logger.Println("NetworkVirtualWireEthernetGroupOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkVirtualWireEthernetGroupOper
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
