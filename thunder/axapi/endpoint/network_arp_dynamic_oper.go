

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkArpDynamicOper struct {
    
    Oper NetworkArpDynamicOperOper `json:"oper"`

}
type DataNetworkArpDynamicOper struct {
    DtNetworkArpDynamicOper NetworkArpDynamicOper `json:"dynamic"`
}


type NetworkArpDynamicOperOper struct {
    ArpList []NetworkArpDynamicOperOperArpList `json:"arp-list"`
}


type NetworkArpDynamicOperOperArpList struct {
    IpAddress string `json:"IP-Address"`
    MacAddress string `json:"MAC-Address"`
    Type string `json:"Type"`
    Age int `json:"Age"`
    Interface string `json:"Interface"`
    Vlan int `json:"Vlan"`
}

func (p *NetworkArpDynamicOper) GetId() string{
    return "1"
}

func (p *NetworkArpDynamicOper) getPath() string{
    return "network/arp/dynamic/oper"
}

func (p *NetworkArpDynamicOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkArpDynamicOper,error) {
logger.Println("NetworkArpDynamicOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkArpDynamicOper
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
