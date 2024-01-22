

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nat64UserQuotaSessionOper struct {
    
    Oper Cgnv6Nat64UserQuotaSessionOperOper `json:"oper"`

}
type DataCgnv6Nat64UserQuotaSessionOper struct {
    DtCgnv6Nat64UserQuotaSessionOper Cgnv6Nat64UserQuotaSessionOper `json:"user-quota-session"`
}


type Cgnv6Nat64UserQuotaSessionOperOper struct {
    SessionList []Cgnv6Nat64UserQuotaSessionOperOperSessionList `json:"session-list"`
    SessionCount int `json:"session-count"`
    PrefixFilter string `json:"prefix-filter"`
    InsideAddrV6 string `json:"inside-addr-v6"`
    AllPartitions int `json:"all-partitions"`
    SharedPartition int `json:"shared-partition"`
    PartitionName string `json:"partition-name"`
    NatPoolName string `json:"nat-pool-name"`
    PoolShared int `json:"pool-shared"`
    TopCount int `json:"top-count"`
    TopByTcpUsage int `json:"top-by-tcp-usage"`
    TopByUdpUsage int `json:"top-by-udp-usage"`
    TopByIcmpUsage int `json:"top-by-icmp-usage"`
    TopByAllUsage int `json:"top-by-all-usage"`
    TopSortByCli int `json:"top-sort-by-cli"`
    DisplayDebug string `json:"display-debug"`
    NatAddr string `json:"nat-addr"`
}


type Cgnv6Nat64UserQuotaSessionOperOperSessionList struct {
    InsideAddress string `json:"inside-address"`
    PrefixLen int `json:"prefix-len"`
    NatAddress string `json:"nat-address"`
    IcmpQuota int `json:"icmp-quota"`
    UdpQuota int `json:"udp-quota"`
    TcpQuota int `json:"tcp-quota"`
    SessionCount int `json:"session-count"`
    NatPoolName string `json:"nat-pool-name"`
    LidNumber int `json:"lid-number"`
    Flags string `json:"flags"`
}

func (p *Cgnv6Nat64UserQuotaSessionOper) GetId() string{
    return "1"
}

func (p *Cgnv6Nat64UserQuotaSessionOper) getPath() string{
    return "cgnv6/nat64/user-quota-session/oper"
}

func (p *Cgnv6Nat64UserQuotaSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Nat64UserQuotaSessionOper,error) {
logger.Println("Cgnv6Nat64UserQuotaSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Nat64UserQuotaSessionOper
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
