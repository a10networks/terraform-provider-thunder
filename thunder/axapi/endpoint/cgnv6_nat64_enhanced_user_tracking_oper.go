

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nat64EnhancedUserTrackingOper struct {
    
    Oper Cgnv6Nat64EnhancedUserTrackingOperOper `json:"oper"`

}
type DataCgnv6Nat64EnhancedUserTrackingOper struct {
    DtCgnv6Nat64EnhancedUserTrackingOper Cgnv6Nat64EnhancedUserTrackingOper `json:"enhanced-user-tracking"`
}


type Cgnv6Nat64EnhancedUserTrackingOperOper struct {
    UserList []Cgnv6Nat64EnhancedUserTrackingOperOperUserList `json:"user-list"`
    UserCount int `json:"user-count"`
    PrefixFilter string `json:"prefix-filter"`
    InsideAddrV6 string `json:"inside-addr-v6"`
    AllPartitions int `json:"all-partitions"`
    SharedPartition int `json:"shared-partition"`
    PartitionName string `json:"partition-name"`
    NatPoolName string `json:"nat-pool-name"`
    PoolShared int `json:"pool-shared"`
}


type Cgnv6Nat64EnhancedUserTrackingOperOperUserList struct {
    InsideAddress string `json:"inside-address"`
    PrefixLen int `json:"prefix-len"`
    NatAddress string `json:"nat-address"`
    TcpQuota int `json:"tcp-quota"`
    UdpQuota int `json:"udp-quota"`
    IcmpQuota int `json:"icmp-quota"`
    SessionCount int `json:"session-count"`
    TcpPeak int `json:"tcp-peak"`
    UdpPeak int `json:"udp-peak"`
    IcmpPeak int `json:"icmp-peak"`
    SessionPeak int `json:"session-peak"`
    TotalSessionCount int `json:"total-session-count"`
    UplPackets int `json:"upl-packets"`
    UplBytes int `json:"upl-bytes"`
    DwlPackets int `json:"dwl-packets"`
    DwlBytes int `json:"dwl-bytes"`
    NatPoolName string `json:"nat-pool-name"`
}

func (p *Cgnv6Nat64EnhancedUserTrackingOper) GetId() string{
    return "1"
}

func (p *Cgnv6Nat64EnhancedUserTrackingOper) getPath() string{
    return "cgnv6/nat64/enhanced-user-tracking/oper"
}

func (p *Cgnv6Nat64EnhancedUserTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Nat64EnhancedUserTrackingOper,error) {
logger.Println("Cgnv6Nat64EnhancedUserTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Nat64EnhancedUserTrackingOper
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
