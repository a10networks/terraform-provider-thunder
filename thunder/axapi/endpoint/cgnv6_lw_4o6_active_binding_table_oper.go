

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Lw4o6ActiveBindingTableOper struct {
    
    Oper Cgnv6Lw4o6ActiveBindingTableOperOper `json:"oper"`

}
type DataCgnv6Lw4o6ActiveBindingTableOper struct {
    DtCgnv6Lw4o6ActiveBindingTableOper Cgnv6Lw4o6ActiveBindingTableOper `json:"active-binding-table"`
}


type Cgnv6Lw4o6ActiveBindingTableOperOper struct {
    EntryList []Cgnv6Lw4o6ActiveBindingTableOperOperEntryList `json:"entry-list"`
}


type Cgnv6Lw4o6ActiveBindingTableOperOperEntryList struct {
    TunnelAddress string `json:"tunnel-address"`
    NatAddress string `json:"nat-address"`
    PortStart int `json:"port-start"`
    PortEnd int `json:"port-end"`
    FwdMatchCount int `json:"fwd-match-count"`
    RevMatchCount int `json:"rev-match-count"`
    AftrTunnelAddress string `json:"aftr-tunnel-address"`
}

func (p *Cgnv6Lw4o6ActiveBindingTableOper) GetId() string{
    return "1"
}

func (p *Cgnv6Lw4o6ActiveBindingTableOper) getPath() string{
    return "cgnv6/lw-4o6/active-binding-table/oper"
}

func (p *Cgnv6Lw4o6ActiveBindingTableOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Lw4o6ActiveBindingTableOper,error) {
logger.Println("Cgnv6Lw4o6ActiveBindingTableOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Lw4o6ActiveBindingTableOper
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
