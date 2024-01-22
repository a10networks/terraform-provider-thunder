

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6ScaleoutAddressMappingOper struct {
    
    Oper Cgnv6ScaleoutAddressMappingOperOper `json:"oper"`

}
type DataCgnv6ScaleoutAddressMappingOper struct {
    DtCgnv6ScaleoutAddressMappingOper Cgnv6ScaleoutAddressMappingOper `json:"address-mapping"`
}


type Cgnv6ScaleoutAddressMappingOperOper struct {
    UserGroup int `json:"user-group"`
    ActiveNode int `json:"active-node"`
    StandbyNode int `json:"standby-node"`
    IpList []Cgnv6ScaleoutAddressMappingOperOperIpList `json:"ip-list"`
    Ip string `json:"ip"`
    Ipv6 string `json:"ipv6"`
    NatIp string `json:"nat-ip"`
    NatPool string `json:"nat-pool"`
}


type Cgnv6ScaleoutAddressMappingOperOperIpList struct {
    StartNatIp string `json:"start-nat-ip"`
    EndNatIp string `json:"end-nat-ip"`
}

func (p *Cgnv6ScaleoutAddressMappingOper) GetId() string{
    return "1"
}

func (p *Cgnv6ScaleoutAddressMappingOper) getPath() string{
    return "cgnv6/scaleout/address-mapping/oper"
}

func (p *Cgnv6ScaleoutAddressMappingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6ScaleoutAddressMappingOper,error) {
logger.Println("Cgnv6ScaleoutAddressMappingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6ScaleoutAddressMappingOper
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
