

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LoggingSourceAddressOper struct {
    
    Oper Cgnv6LoggingSourceAddressOperOper `json:"oper"`

}
type DataCgnv6LoggingSourceAddressOper struct {
    DtCgnv6LoggingSourceAddressOper Cgnv6LoggingSourceAddressOper `json:"source-address"`
}


type Cgnv6LoggingSourceAddressOperOper struct {
    SourceAddressList Cgnv6LoggingSourceAddressOperOperSourceAddressList `json:"source-address-list"`
}


type Cgnv6LoggingSourceAddressOperOperSourceAddressList struct {
    SourceAddress string `json:"source-address"`
    RefCount int `json:"ref-count"`
    TcpTotal int `json:"tcp-total"`
    TcpAllocated int `json:"tcp-allocated"`
    TcpFreed int `json:"tcp-freed"`
    TcpFailed int `json:"tcp-failed"`
    UdpTotal int `json:"udp-total"`
    UdpAllocated int `json:"udp-allocated"`
    UdpFreed int `json:"udp-freed"`
    UdpFailed int `json:"udp-failed"`
}

func (p *Cgnv6LoggingSourceAddressOper) GetId() string{
    return "1"
}

func (p *Cgnv6LoggingSourceAddressOper) getPath() string{
    return "cgnv6/logging/source-address/oper"
}

func (p *Cgnv6LoggingSourceAddressOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LoggingSourceAddressOper,error) {
logger.Println("Cgnv6LoggingSourceAddressOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LoggingSourceAddressOper
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
