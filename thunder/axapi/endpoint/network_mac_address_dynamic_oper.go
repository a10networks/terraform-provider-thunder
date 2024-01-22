

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkMacAddressDynamicOper struct {
    
    Oper NetworkMacAddressDynamicOperOper `json:"oper"`

}
type DataNetworkMacAddressDynamicOper struct {
    DtNetworkMacAddressDynamicOper NetworkMacAddressDynamicOper `json:"dynamic"`
}


type NetworkMacAddressDynamicOperOper struct {
    AgeTime int `json:"Age-time"`
    Macoper []NetworkMacAddressDynamicOperOperMacoper `json:"macoper"`
}


type NetworkMacAddressDynamicOperOperMacoper struct {
    MacAddress string `json:"MAC-Address"`
    Port int `json:"Port"`
    Type string `json:"Type"`
    Index int `json:"Index"`
    Vlan int `json:"Vlan"`
    Age int `json:"Age"`
}

func (p *NetworkMacAddressDynamicOper) GetId() string{
    return "1"
}

func (p *NetworkMacAddressDynamicOper) getPath() string{
    return "network/mac-address/dynamic/oper"
}

func (p *NetworkMacAddressDynamicOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkMacAddressDynamicOper,error) {
logger.Println("NetworkMacAddressDynamicOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkMacAddressDynamicOper
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
