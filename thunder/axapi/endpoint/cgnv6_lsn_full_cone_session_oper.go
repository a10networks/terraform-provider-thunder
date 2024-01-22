

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnFullConeSessionOper struct {
    
    Oper Cgnv6LsnFullConeSessionOperOper `json:"oper"`

}
type DataCgnv6LsnFullConeSessionOper struct {
    DtCgnv6LsnFullConeSessionOper Cgnv6LsnFullConeSessionOper `json:"full-cone-session"`
}


type Cgnv6LsnFullConeSessionOperOper struct {
    SessionList []Cgnv6LsnFullConeSessionOperOperSessionList `json:"session-list"`
    SessionCount int `json:"session-count"`
    TotalSessionCount int `json:"total-session-count"`
    AllPartitions int `json:"all-partitions"`
    SharedPartition int `json:"shared-partition"`
    PartitionName string `json:"partition-name"`
    InsideAddr string `json:"inside-addr"`
    InsideAddrStart string `json:"inside-addr-start"`
    InsideAddrEnd string `json:"inside-addr-end"`
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


type Cgnv6LsnFullConeSessionOperOperSessionList struct {
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

func (p *Cgnv6LsnFullConeSessionOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnFullConeSessionOper) getPath() string{
    return "cgnv6/lsn/full-cone-session/oper"
}

func (p *Cgnv6LsnFullConeSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnFullConeSessionOper,error) {
logger.Println("Cgnv6LsnFullConeSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnFullConeSessionOper
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
