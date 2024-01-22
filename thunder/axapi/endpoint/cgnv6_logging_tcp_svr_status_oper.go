

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LoggingTcpSvrStatusOper struct {
    
    Oper Cgnv6LoggingTcpSvrStatusOperOper `json:"oper"`

}
type DataCgnv6LoggingTcpSvrStatusOper struct {
    DtCgnv6LoggingTcpSvrStatusOper Cgnv6LoggingTcpSvrStatusOper `json:"tcp-svr-status"`
}


type Cgnv6LoggingTcpSvrStatusOperOper struct {
    TcpSvrList Cgnv6LoggingTcpSvrStatusOperOperTcpSvrList `json:"tcp-svr-list"`
}


type Cgnv6LoggingTcpSvrStatusOperOperTcpSvrList struct {
    Template string `json:"template"`
    ServerName string `json:"server-name"`
    ServerPort int `json:"server-port"`
    OverallStatus int `json:"overall-status"`
    NumCpus int `json:"num-cpus"`
    Status string `json:"status"`
}

func (p *Cgnv6LoggingTcpSvrStatusOper) GetId() string{
    return "1"
}

func (p *Cgnv6LoggingTcpSvrStatusOper) getPath() string{
    return "cgnv6/logging/tcp-svr-status/oper"
}

func (p *Cgnv6LoggingTcpSvrStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LoggingTcpSvrStatusOper,error) {
logger.Println("Cgnv6LoggingTcpSvrStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LoggingTcpSvrStatusOper
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
