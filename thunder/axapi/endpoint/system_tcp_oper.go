

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTcpOper struct {
    
    Oper SystemTcpOperOper `json:"oper"`
    RateLimitResetUnknownConn SystemTcpOperRateLimitResetUnknownConn `json:"rate-limit-reset-unknown-conn"`

}
type DataSystemTcpOper struct {
    DtSystemTcpOper SystemTcpOper `json:"tcp"`
}


type SystemTcpOperOper struct {
    TcpCpuList []SystemTcpOperOperTcpCpuList `json:"tcp-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SystemTcpOperOperTcpCpuList struct {
    Currestab int `json:"currestab"`
    Activeopens int `json:"activeopens"`
    Passiveopens int `json:"passiveopens"`
    Attemptfails int `json:"attemptfails"`
    Insegs int `json:"insegs"`
    Outsegs int `json:"outsegs"`
    Retranssegs int `json:"retranssegs"`
    Estabresets int `json:"estabresets"`
    Outrsts int `json:"outrsts"`
    Noroute int `json:"noroute"`
    Tfo_conns int `json:"tfo_conns"`
    Tfo_actives int `json:"tfo_actives"`
    Tfo_denied int `json:"tfo_denied"`
    Inerrs int `json:"inerrs"`
    Sock_alloc int `json:"sock_alloc"`
    Orphan_count int `json:"orphan_count"`
    Mem_alloc int `json:"mem_alloc"`
    Recv_mem int `json:"recv_mem"`
    Send_mem int `json:"send_mem"`
    Currsyssnt int `json:"currsyssnt"`
    Currsynrcv int `json:"currsynrcv"`
    Currfinw1 int `json:"currfinw1"`
    Currfinw2 int `json:"currfinw2"`
    Currtimew int `json:"currtimew"`
    Currclose int `json:"currclose"`
    Currclsw int `json:"currclsw"`
    Currlack int `json:"currlack"`
    Currlstn int `json:"currlstn"`
    Currclsg int `json:"currclsg"`
    Pawsactiverejected int `json:"pawsactiverejected"`
    Syn_rcv_rstack int `json:"syn_rcv_rstack"`
    Syn_rcv_rst int `json:"syn_rcv_rst"`
    Syn_rcv_ack int `json:"syn_rcv_ack"`
    Tcpabortontimeout int `json:"tcpabortontimeout"`
    Ax_rexmit_syn int `json:"ax_rexmit_syn"`
    Exceedmss int `json:"exceedmss"`
}


type SystemTcpOperRateLimitResetUnknownConn struct {
    Oper SystemTcpOperRateLimitResetUnknownConnOper `json:"oper"`
}


type SystemTcpOperRateLimitResetUnknownConnOper struct {
    UnknownConnRateLimit int `json:"unknown-conn-rate-limit"`
    UnknownConnCurrentRate int `json:"unknown-conn-current-rate"`
    UnknownConnRateLimitDrop int `json:"unknown-conn-rate-limit-drop"`
}

func (p *SystemTcpOper) GetId() string{
    return "1"
}

func (p *SystemTcpOper) getPath() string{
    return "system/tcp/oper"
}

func (p *SystemTcpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTcpOper,error) {
logger.Println("SystemTcpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTcpOper
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
