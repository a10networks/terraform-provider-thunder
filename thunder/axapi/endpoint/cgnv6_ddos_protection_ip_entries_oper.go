

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DdosProtectionIpEntriesOper struct {
    
    Oper Cgnv6DdosProtectionIpEntriesOperOper `json:"oper"`

}
type DataCgnv6DdosProtectionIpEntriesOper struct {
    DtCgnv6DdosProtectionIpEntriesOper Cgnv6DdosProtectionIpEntriesOper `json:"ip-entries"`
}


type Cgnv6DdosProtectionIpEntriesOperOper struct {
    DdosIpEntriesList []Cgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList `json:"ddos-ip-entries-list"`
    TotalEntries int `json:"total-entries"`
    All int `json:"all"`
    NatPool string `json:"nat-pool"`
    V4Netmask string `json:"v4-netmask"`
}


type Cgnv6DdosProtectionIpEntriesOperOperDdosIpEntriesList struct {
    V4Address string `json:"v4-address"`
    Expiration int `json:"expiration"`
    Hints string `json:"hints"`
    L4EntriesCount int `json:"l4-entries-count"`
    SwReceivePps int `json:"sw-receive-pps"`
    SwL4DropPps int `json:"sw-l4-drop-pps"`
    HwL4DropPps int `json:"hw-l4-drop-pps"`
    SwL3DropPps int `json:"sw-l3-drop-pps"`
    HwL3DropPps int `json:"hw-l3-drop-pps"`
    TotalPps int `json:"total-pps"`
    InBlacklist int `json:"in-blacklist"`
    InHardware int `json:"in-hardware"`
    HardwareIndex int `json:"hardware-index"`
}

func (p *Cgnv6DdosProtectionIpEntriesOper) GetId() string{
    return "1"
}

func (p *Cgnv6DdosProtectionIpEntriesOper) getPath() string{
    return "cgnv6/ddos-protection/ip-entries/oper"
}

func (p *Cgnv6DdosProtectionIpEntriesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6DdosProtectionIpEntriesOper,error) {
logger.Println("Cgnv6DdosProtectionIpEntriesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6DdosProtectionIpEntriesOper
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
