

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DsLiteFullConeSessionOper struct {
    
    Oper Cgnv6DsLiteFullConeSessionOperOper `json:"oper"`

}
type DataCgnv6DsLiteFullConeSessionOper struct {
    DtCgnv6DsLiteFullConeSessionOper Cgnv6DsLiteFullConeSessionOper `json:"full-cone-session"`
}


type Cgnv6DsLiteFullConeSessionOperOper struct {
    SessionList []Cgnv6DsLiteFullConeSessionOperOperSessionList `json:"session-list"`
    SessionCount int `json:"session-count"`
    TotalSessionCount int `json:"total-session-count"`
    AllPartitions int `json:"all-partitions"`
    SharedPartition int `json:"shared-partition"`
    PartitionName string `json:"partition-name"`
    InsideAddrV6 string `json:"inside-addr-v6"`
    InsideAddrV6Start string `json:"inside-addr-v6-start"`
    InsideAddrV6End string `json:"inside-addr-v6-end"`
    NatPoolName string `json:"nat-pool-name"`
    PoolShared int `json:"pool-shared"`
    Pcp int `json:"pcp"`
    DebugSession int `json:"debug-session"`
    InsideAddr string `json:"inside-addr"`
    InsideAddrStart string `json:"inside-addr-start"`
    InsideAddrEnd string `json:"inside-addr-end"`
    NatAddr string `json:"nat-addr"`
    NatAddrStart string `json:"nat-addr-start"`
    NatAddrEnd string `json:"nat-addr-end"`
    InsidePort int `json:"inside-port"`
    Graceful int `json:"graceful"`
    NatPort int `json:"nat-port"`
}


type Cgnv6DsLiteFullConeSessionOperOperSessionList struct {
    Protocol string `json:"protocol"`
    InsideV6Address string `json:"inside-v6-address"`
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

func (p *Cgnv6DsLiteFullConeSessionOper) GetId() string{
    return "1"
}

func (p *Cgnv6DsLiteFullConeSessionOper) getPath() string{
    return "cgnv6/ds-lite/full-cone-session/oper"
}

func (p *Cgnv6DsLiteFullConeSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6DsLiteFullConeSessionOper,error) {
logger.Println("Cgnv6DsLiteFullConeSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6DsLiteFullConeSessionOper
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
