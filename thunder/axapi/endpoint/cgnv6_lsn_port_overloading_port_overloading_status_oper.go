

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnPortOverloadingPortOverloadingStatusOper struct {
    
    Oper Cgnv6LsnPortOverloadingPortOverloadingStatusOperOper `json:"oper"`

}
type DataCgnv6LsnPortOverloadingPortOverloadingStatusOper struct {
    DtCgnv6LsnPortOverloadingPortOverloadingStatusOper Cgnv6LsnPortOverloadingPortOverloadingStatusOper `json:"port-overloading-status"`
}


type Cgnv6LsnPortOverloadingPortOverloadingStatusOperOper struct {
    TcpList Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperTcpList `json:"tcp-list"`
    UdpList Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperUdpList `json:"udp-list"`
    Unique string `json:"unique"`
}


type Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperTcpList struct {
    StartPortTcp int `json:"start-port-tcp"`
    EndPortTcp int `json:"end-port-tcp"`
    StatusTcp string `json:"status-tcp"`
}


type Cgnv6LsnPortOverloadingPortOverloadingStatusOperOperUdpList struct {
    StartPortUdp int `json:"start-port-udp"`
    EndPortUdp int `json:"end-port-udp"`
    StatusUdp string `json:"status-udp"`
}

func (p *Cgnv6LsnPortOverloadingPortOverloadingStatusOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnPortOverloadingPortOverloadingStatusOper) getPath() string{
    return "cgnv6/lsn/port-overloading/port-overloading-status/oper"
}

func (p *Cgnv6LsnPortOverloadingPortOverloadingStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnPortOverloadingPortOverloadingStatusOper,error) {
logger.Println("Cgnv6LsnPortOverloadingPortOverloadingStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnPortOverloadingPortOverloadingStatusOper
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
