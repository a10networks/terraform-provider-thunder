

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnSystemStatusOper struct {
    
    Oper Cgnv6LsnSystemStatusOperOper `json:"oper"`

}
type DataCgnv6LsnSystemStatusOper struct {
    DtCgnv6LsnSystemStatusOper Cgnv6LsnSystemStatusOper `json:"system-status"`
}


type Cgnv6LsnSystemStatusOperOper struct {
    LsnCps int `json:"lsn-cps"`
    DataSessionsUsed int `json:"data-sessions-used"`
    DataSessionsFree int `json:"data-sessions-free"`
    SmpSessionsUsed int `json:"smp-sessions-used"`
    SmpSessionsFree int `json:"smp-sessions-free"`
    TcpNatPortsUsed int `json:"tcp-nat-ports-used"`
    TcpNatPortsFree int `json:"tcp-nat-ports-free"`
    UdpNatPortsUsed int `json:"udp-nat-ports-used"`
    UdpNatPortsFree int `json:"udp-nat-ports-free"`
    RadiusEntriesUsed int `json:"radius-entries-used"`
    RadiusEntriesFree int `json:"radius-entries-free"`
}

func (p *Cgnv6LsnSystemStatusOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnSystemStatusOper) getPath() string{
    return "cgnv6/lsn/system-status/oper"
}

func (p *Cgnv6LsnSystemStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnSystemStatusOper,error) {
logger.Println("Cgnv6LsnSystemStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnSystemStatusOper
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
