

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DdosProtectionL4EntriesOper struct {
    
    Oper Cgnv6DdosProtectionL4EntriesOperOper `json:"oper"`

}
type DataCgnv6DdosProtectionL4EntriesOper struct {
    DtCgnv6DdosProtectionL4EntriesOper Cgnv6DdosProtectionL4EntriesOper `json:"l4-entries"`
}


type Cgnv6DdosProtectionL4EntriesOperOper struct {
    DdosL4EntriesList []Cgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList `json:"ddos-l4-entries-list"`
    TotalEntries int `json:"total-entries"`
    All int `json:"all"`
    NatPool string `json:"nat-pool"`
    V4Netmask string `json:"v4-netmask"`
    NotInHardware int `json:"not-in-hardware"`
}


type Cgnv6DdosProtectionL4EntriesOperOperDdosL4EntriesList struct {
    V4Address string `json:"v4-address"`
    L4Protocol string `json:"l4-protocol"`
    Port int `json:"port"`
    InHardware int `json:"in-hardware"`
    HardwareIndex int `json:"hardware-index"`
    Pps int `json:"pps"`
    Expiration int `json:"expiration"`
    IsDeleted int `json:"is-deleted"`
}

func (p *Cgnv6DdosProtectionL4EntriesOper) GetId() string{
    return "1"
}

func (p *Cgnv6DdosProtectionL4EntriesOper) getPath() string{
    return "cgnv6/ddos-protection/l4-entries/oper"
}

func (p *Cgnv6DdosProtectionL4EntriesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6DdosProtectionL4EntriesOper,error) {
logger.Println("Cgnv6DdosProtectionL4EntriesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6DdosProtectionL4EntriesOper
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
