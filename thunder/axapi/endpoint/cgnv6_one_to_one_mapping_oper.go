

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOneMappingOper struct {
    
    Oper Cgnv6OneToOneMappingOperOper `json:"oper"`

}
type DataCgnv6OneToOneMappingOper struct {
    DtCgnv6OneToOneMappingOper Cgnv6OneToOneMappingOper `json:"mapping"`
}


type Cgnv6OneToOneMappingOperOper struct {
    SessionMappingList []Cgnv6OneToOneMappingOperOperSessionMappingList `json:"session-mapping-list"`
    Total int `json:"total"`
    InsideAddressIpv4 string `json:"inside-address-ipv4"`
    InsideAddressIpv6 string `json:"inside-address-ipv6"`
    AllPartitions int `json:"all-partitions"`
    SharedPartition int `json:"shared-partition"`
    PartitionName string `json:"partition-name"`
    InsideAddrStart string `json:"inside-addr-start"`
    InsideAddrEnd string `json:"inside-addr-end"`
    NatAddrVal string `json:"nat-addr-val"`
    NatAddrStart string `json:"nat-addr-start"`
    NatAddrEnd string `json:"nat-addr-end"`
    PoolName string `json:"pool-name"`
    SharedPoolName string `json:"shared-pool-name"`
}


type Cgnv6OneToOneMappingOperOperSessionMappingList struct {
    InsideIpv4Address string `json:"inside-ipv4-address"`
    InsideIpv6Address string `json:"inside-ipv6-address"`
    NatAddress string `json:"nat-address"`
    Sessions int `json:"sessions"`
    Age string `json:"age"`
    Pool string `json:"pool"`
}

func (p *Cgnv6OneToOneMappingOper) GetId() string{
    return "1"
}

func (p *Cgnv6OneToOneMappingOper) getPath() string{
    return "cgnv6/one-to-one/mapping/oper"
}

func (p *Cgnv6OneToOneMappingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6OneToOneMappingOper,error) {
logger.Println("Cgnv6OneToOneMappingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6OneToOneMappingOper
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
