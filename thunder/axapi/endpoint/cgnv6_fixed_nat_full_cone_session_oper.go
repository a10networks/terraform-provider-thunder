

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatFullConeSessionOper struct {
    
    Oper Cgnv6FixedNatFullConeSessionOperOper `json:"oper"`

}
type DataCgnv6FixedNatFullConeSessionOper struct {
    DtCgnv6FixedNatFullConeSessionOper Cgnv6FixedNatFullConeSessionOper `json:"full-cone-session"`
}


type Cgnv6FixedNatFullConeSessionOperOper struct {
    SessionList []Cgnv6FixedNatFullConeSessionOperOperSessionList `json:"session-list"`
    SessionCount int `json:"session-count"`
    Nat44TotalSessionCount int `json:"nat44-total-session-count"`
    Nat64TotalSessionCount int `json:"nat64-total-session-count"`
    DsliteTotalSessionCount int `json:"dslite-total-session-count"`
    SessionType string `json:"session-type"`
    AllPartitions int `json:"all-partitions"`
    SharedPartition int `json:"shared-partition"`
    PartitionName string `json:"partition-name"`
    InsideAddr string `json:"inside-addr"`
    InsideAddrStart string `json:"inside-addr-start"`
    InsideAddrEnd string `json:"inside-addr-end"`
    InsideAddrV6 string `json:"inside-addr-v6"`
    InsideAddrV6Start string `json:"inside-addr-v6-start"`
    InsideAddrV6End string `json:"inside-addr-v6-end"`
    NatAddr string `json:"nat-addr"`
    NatAddrStart string `json:"nat-addr-start"`
    NatAddrEnd string `json:"nat-addr-end"`
    InsidePort int `json:"inside-port"`
    NatPort int `json:"nat-port"`
    Pcp int `json:"pcp"`
    Graceful int `json:"graceful"`
    DebugSession int `json:"debug-session"`
}


type Cgnv6FixedNatFullConeSessionOperOperSessionList struct {
    Protocol string `json:"protocol"`
    InsideV6Address string `json:"inside-v6-address"`
    InsideAddress string `json:"inside-address"`
    InsidePort int `json:"inside-port"`
    NatAddress string `json:"nat-address"`
    NatPort int `json:"nat-port"`
    Eim int `json:"EIM"`
    Eif int `json:"EIF"`
    Cpu int `json:"cpu"`
    Age string `json:"age"`
    Flags string `json:"flags"`
}

func (p *Cgnv6FixedNatFullConeSessionOper) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatFullConeSessionOper) getPath() string{
    return "cgnv6/fixed-nat/full-cone-session/oper"
}

func (p *Cgnv6FixedNatFullConeSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatFullConeSessionOper,error) {
logger.Println("Cgnv6FixedNatFullConeSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatFullConeSessionOper
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
