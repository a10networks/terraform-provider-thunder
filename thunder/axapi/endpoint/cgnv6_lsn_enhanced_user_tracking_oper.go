

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnEnhancedUserTrackingOper struct {
    
    Oper Cgnv6LsnEnhancedUserTrackingOperOper `json:"oper"`

}
type DataCgnv6LsnEnhancedUserTrackingOper struct {
    DtCgnv6LsnEnhancedUserTrackingOper Cgnv6LsnEnhancedUserTrackingOper `json:"enhanced-user-tracking"`
}


type Cgnv6LsnEnhancedUserTrackingOperOper struct {
    UserList []Cgnv6LsnEnhancedUserTrackingOperOperUserList `json:"user-list"`
    UserCount int `json:"user-count"`
    AllPartitions int `json:"all-partitions"`
    SharedPartition int `json:"shared-partition"`
    PartitionName string `json:"partition-name"`
    InsideAddr string `json:"inside-addr"`
    NatPoolName string `json:"nat-pool-name"`
    PoolShared int `json:"pool-shared"`
}


type Cgnv6LsnEnhancedUserTrackingOperOperUserList struct {
    InsideAddress string `json:"inside-address"`
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

func (p *Cgnv6LsnEnhancedUserTrackingOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnEnhancedUserTrackingOper) getPath() string{
    return "cgnv6/lsn/enhanced-user-tracking/oper"
}

func (p *Cgnv6LsnEnhancedUserTrackingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnEnhancedUserTrackingOper,error) {
logger.Println("Cgnv6LsnEnhancedUserTrackingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnEnhancedUserTrackingOper
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
