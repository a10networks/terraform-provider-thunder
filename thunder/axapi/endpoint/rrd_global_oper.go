

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdGlobalOper struct {
    
    Oper RrdGlobalOperOper `json:"oper"`

}
type DataRrdGlobalOper struct {
    DtRrdGlobalOper RrdGlobalOper `json:"global"`
}


type RrdGlobalOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    GlobalData []RrdGlobalOperOperGlobalData `json:"global-data"`
}


type RrdGlobalOperOperGlobalData struct {
    Time int `json:"time"`
    Tcp_est_count int `json:"tcp_est_count"`
    Tcp_half_open int `json:"tcp_half_open"`
    Udp_conn_count int `json:"udp_conn_count"`
    Ip_conn_count int `json:"ip_conn_count"`
    Other_conn_count int `json:"other_conn_count"`
    Snat_tcp_count int `json:"snat_tcp_count"`
    Snat_udp_count int `json:"snat_udp_count"`
    Free_buff_count int `json:"free_buff_count"`
    Curr_free_conn int `json:"curr_free_conn"`
    Conn_get_cnt int `json:"conn_get_cnt"`
    Conn_free_cnt int `json:"conn_free_cnt"`
    Syn_half_open int `json:"syn_half_open"`
    Conn_smp_alloc int `json:"conn_smp_alloc"`
    Conn_smp_free int `json:"conn_smp_free"`
    Conn_smp_aged int `json:"conn_smp_aged"`
}

func (p *RrdGlobalOper) GetId() string{
    return "1"
}

func (p *RrdGlobalOper) getPath() string{
    return "rrd/global/oper"
}

func (p *RrdGlobalOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdGlobalOper,error) {
logger.Println("RrdGlobalOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdGlobalOper
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
