

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatDisabledConfigOper struct {
    
    Oper Cgnv6FixedNatDisabledConfigOperOper `json:"oper"`

}
type DataCgnv6FixedNatDisabledConfigOper struct {
    DtCgnv6FixedNatDisabledConfigOper Cgnv6FixedNatDisabledConfigOper `json:"disabled-config"`
}


type Cgnv6FixedNatDisabledConfigOperOper struct {
    DisabledConfigList []Cgnv6FixedNatDisabledConfigOperOperDisabledConfigList `json:"disabled-config-list"`
}


type Cgnv6FixedNatDisabledConfigOperOperDisabledConfigList struct {
    InsideStartAddress string `json:"inside-start-address"`
    InsideEndAddress string `json:"inside-end-address"`
    InsideNetmask int `json:"inside-netmask"`
    InsideIpList string `json:"inside-ip-list"`
    Partition string `json:"partition"`
    ActiveUsers int `json:"active-users"`
    ClearSession int `json:"clear-session"`
}

func (p *Cgnv6FixedNatDisabledConfigOper) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatDisabledConfigOper) getPath() string{
    return "cgnv6/fixed-nat/disabled-config/oper"
}

func (p *Cgnv6FixedNatDisabledConfigOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatDisabledConfigOper,error) {
logger.Println("Cgnv6FixedNatDisabledConfigOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatDisabledConfigOper
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
