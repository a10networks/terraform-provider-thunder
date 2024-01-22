

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nat64FullConeSessionOper struct {
    
    Oper Cgnv6Nat64FullConeSessionOperOper `json:"oper"`

}
type DataCgnv6Nat64FullConeSessionOper struct {
    DtCgnv6Nat64FullConeSessionOper Cgnv6Nat64FullConeSessionOper `json:"full-cone-session"`
}


type Cgnv6Nat64FullConeSessionOperOper struct {
    SessionList []Cgnv6Nat64FullConeSessionOperOperSessionList `json:"session-list"`
    SessionCount int `json:"session-count"`
    TotalSessionCount int `json:"total-session-count"`
    AllPartitions int `json:"all-partitions"`
    SharedPartition int `json:"shared-partition"`
    PartitionName string `json:"partition-name"`
    InsideAddrV6 string `json:"inside-addr-v6"`
    InsideAddrV6Start string `json:"inside-addr-v6-start"`
    InsideAddrV6End string `json:"inside-addr-v6-end"`
    NatAddr string `json:"nat-addr"`
    NatAddrStart string `json:"nat-addr-start"`
    NatAddrEnd string `json:"nat-addr-end"`
    InsidePort int `json:"inside-port"`
    NatPort int `json:"nat-port"`
    NatPoolName string `json:"nat-pool-name"`
    PoolShared int `json:"pool-shared"`
    Pcp int `json:"pcp"`
    Graceful int `json:"graceful"`
    DebugSession int `json:"debug-session"`
}


type Cgnv6Nat64FullConeSessionOperOperSessionList struct {
    Protocol string `json:"protocol"`
    InsideAddress string `json:"inside-address"`
    InsidePort int `json:"inside-port"`
    NatAddress string `json:"nat-address"`
    NatPort int `json:"nat-port"`
    Outbound int `json:"outbound"`
    Inbound int `json:"inbound"`
    NatPoolName string `json:"nat-pool-name"`
    Cpu int `json:"cpu"`
    Age string `json:"age"`
    Flags string `json:"flags"`
}

func (p *Cgnv6Nat64FullConeSessionOper) GetId() string{
    return "1"
}

func (p *Cgnv6Nat64FullConeSessionOper) getPath() string{
    return "cgnv6/nat64/full-cone-session/oper"
}

func (p *Cgnv6Nat64FullConeSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Nat64FullConeSessionOper,error) {
logger.Println("Cgnv6Nat64FullConeSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Nat64FullConeSessionOper
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
