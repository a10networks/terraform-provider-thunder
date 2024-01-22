

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpATeMacOper struct {
    
    Oper VrrpATeMacOperOper `json:"oper"`

}
type DataVrrpATeMacOper struct {
    DtVrrpATeMacOper VrrpATeMacOper `json:"te-mac"`
}


type VrrpATeMacOperOper struct {
    InvalidFlag int `json:"invalid-flag"`
    VridMacList []VrrpATeMacOperOperVridMacList `json:"vrid-mac-list"`
}


type VrrpATeMacOperOperVridMacList struct {
    Vrid int `json:"vrid"`
    Mac_address_inside string `json:"mac_address_inside"`
    Mac_address_outside string `json:"mac_address_outside"`
}

func (p *VrrpATeMacOper) GetId() string{
    return "1"
}

func (p *VrrpATeMacOper) getPath() string{
    return "vrrp-a/te-mac/oper"
}

func (p *VrrpATeMacOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpATeMacOper,error) {
logger.Println("VrrpATeMacOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpATeMacOper
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
