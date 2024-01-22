

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkVlanGlobalOper struct {
    
    Oper NetworkVlanGlobalOperOper `json:"oper"`

}
type DataNetworkVlanGlobalOper struct {
    DtNetworkVlanGlobalOper NetworkVlanGlobalOper `json:"vlan-global"`
}


type NetworkVlanGlobalOperOper struct {
    VlanTransList NetworkVlanGlobalOperOperVlanTransList `json:"vlan-trans-list"`
}


type NetworkVlanGlobalOperOperVlanTransList struct {
    Vlan int `json:"vlan"`
}

func (p *NetworkVlanGlobalOper) GetId() string{
    return "1"
}

func (p *NetworkVlanGlobalOper) getPath() string{
    return "network/vlan-global/oper"
}

func (p *NetworkVlanGlobalOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkVlanGlobalOper,error) {
logger.Println("NetworkVlanGlobalOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkVlanGlobalOper
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
