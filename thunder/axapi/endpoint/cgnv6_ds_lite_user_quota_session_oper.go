

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DsLiteUserQuotaSessionOper struct {
    
    Oper Cgnv6DsLiteUserQuotaSessionOperOper `json:"oper"`

}
type DataCgnv6DsLiteUserQuotaSessionOper struct {
    DtCgnv6DsLiteUserQuotaSessionOper Cgnv6DsLiteUserQuotaSessionOper `json:"user-quota-session"`
}


type Cgnv6DsLiteUserQuotaSessionOperOper struct {
    SessionList []Cgnv6DsLiteUserQuotaSessionOperOperSessionList `json:"session-list"`
    SessionCount int `json:"session-count"`
    AllPartitions int `json:"all-partitions"`
    SharedPartition int `json:"shared-partition"`
    PartitionName string `json:"partition-name"`
    PrefixFilter string `json:"prefix-filter"`
    InsideAddrV6 string `json:"inside-addr-v6"`
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


type Cgnv6DsLiteUserQuotaSessionOperOperSessionList struct {
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

func (p *Cgnv6DsLiteUserQuotaSessionOper) GetId() string{
    return "1"
}

func (p *Cgnv6DsLiteUserQuotaSessionOper) getPath() string{
    return "cgnv6/ds-lite/user-quota-session/oper"
}

func (p *Cgnv6DsLiteUserQuotaSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6DsLiteUserQuotaSessionOper,error) {
logger.Println("Cgnv6DsLiteUserQuotaSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6DsLiteUserQuotaSessionOper
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
