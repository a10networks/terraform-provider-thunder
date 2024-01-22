

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatUserQuotaOper struct {
    
    Oper Cgnv6FixedNatUserQuotaOperOper `json:"oper"`

}
type DataCgnv6FixedNatUserQuotaOper struct {
    DtCgnv6FixedNatUserQuotaOper Cgnv6FixedNatUserQuotaOper `json:"user-quota"`
}


type Cgnv6FixedNatUserQuotaOperOper struct {
    QuotaUsageList []Cgnv6FixedNatUserQuotaOperOperQuotaUsageList `json:"quota-usage-list"`
    Partition string `json:"partition"`
    InsideUserV4 string `json:"inside-user-v4"`
    InsideUserV6 string `json:"inside-user-v6"`
}


type Cgnv6FixedNatUserQuotaOperOperQuotaUsageList struct {
    InsideUser string `json:"inside-user"`
    NatAddress string `json:"nat-address"`
    SessionQuotaUsed int `json:"session-quota-used"`
    TcpPortsUsed int `json:"tcp-ports-used"`
    TcpPortsAvailable int `json:"tcp-ports-available"`
    UdpPortsUsed int `json:"udp-ports-used"`
    UdpPortsAvailable int `json:"udp-ports-available"`
    IcmpResourcesUsed int `json:"icmp-resources-used"`
    IcmpResourcesAvailable int `json:"icmp-resources-available"`
}

func (p *Cgnv6FixedNatUserQuotaOper) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatUserQuotaOper) getPath() string{
    return "cgnv6/fixed-nat/user-quota/oper"
}

func (p *Cgnv6FixedNatUserQuotaOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatUserQuotaOper,error) {
logger.Println("Cgnv6FixedNatUserQuotaOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatUserQuotaOper
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
