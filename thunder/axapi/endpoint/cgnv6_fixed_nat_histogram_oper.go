

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatHistogramOper struct {
    
    Oper Cgnv6FixedNatHistogramOperOper `json:"oper"`

}
type DataCgnv6FixedNatHistogramOper struct {
    DtCgnv6FixedNatHistogramOper Cgnv6FixedNatHistogramOper `json:"histogram"`
}


type Cgnv6FixedNatHistogramOperOper struct {
    InsideUserV4 string `json:"inside-user-v4"`
    InsideUserV6 string `json:"inside-user-v6"`
    NatAddress string `json:"nat-address"`
    InsideStartIpv4 string `json:"inside-start-ipv4"`
    InsideEndIpv4 string `json:"inside-end-ipv4"`
    InsideStartIpv6 string `json:"inside-start-ipv6"`
    InsideEndIpv6 string `json:"inside-end-ipv6"`
    OutsideStartIp string `json:"outside-start-ip"`
    OutsideEndIp string `json:"outside-end-ip"`
    TotalUsers int `json:"total-users"`
    HistogramList []Cgnv6FixedNatHistogramOperOperHistogramList `json:"histogram-list"`
}


type Cgnv6FixedNatHistogramOperOperHistogramList struct {
    BinStart int `json:"bin-start"`
    BinEnd int `json:"bin-end"`
    ActiveTcpUsers int `json:"active-tcp-users"`
    ActiveUdpUsers int `json:"active-udp-users"`
}

func (p *Cgnv6FixedNatHistogramOper) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatHistogramOper) getPath() string{
    return "cgnv6/fixed-nat/histogram/oper"
}

func (p *Cgnv6FixedNatHistogramOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatHistogramOper,error) {
logger.Println("Cgnv6FixedNatHistogramOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatHistogramOper
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
