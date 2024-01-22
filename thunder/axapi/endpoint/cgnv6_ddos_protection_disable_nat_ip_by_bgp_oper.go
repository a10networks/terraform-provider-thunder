

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DdosProtectionDisableNatIpByBgpOper struct {
    
    Oper Cgnv6DdosProtectionDisableNatIpByBgpOperOper `json:"oper"`

}
type DataCgnv6DdosProtectionDisableNatIpByBgpOper struct {
    DtCgnv6DdosProtectionDisableNatIpByBgpOper Cgnv6DdosProtectionDisableNatIpByBgpOper `json:"disable-nat-ip-by-bgp"`
}


type Cgnv6DdosProtectionDisableNatIpByBgpOperOper struct {
    DdosDisabledByBgpEntriesList []Cgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList `json:"ddos-disabled-by-bgp-entries-list"`
    TotalEntries int `json:"total-entries"`
    All int `json:"all"`
    V4Address string `json:"v4-address"`
}


type Cgnv6DdosProtectionDisableNatIpByBgpOperOperDdosDisabledByBgpEntriesList struct {
    V4Address string `json:"v4-address"`
    NatPoolName string `json:"nat-pool-name"`
}

func (p *Cgnv6DdosProtectionDisableNatIpByBgpOper) GetId() string{
    return "1"
}

func (p *Cgnv6DdosProtectionDisableNatIpByBgpOper) getPath() string{
    return "cgnv6/ddos-protection/disable-nat-ip-by-bgp/oper"
}

func (p *Cgnv6DdosProtectionDisableNatIpByBgpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6DdosProtectionDisableNatIpByBgpOper,error) {
logger.Println("Cgnv6DdosProtectionDisableNatIpByBgpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6DdosProtectionDisableNatIpByBgpOper
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
