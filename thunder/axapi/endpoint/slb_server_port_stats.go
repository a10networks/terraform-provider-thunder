

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbServerPortStats47 struct {
    
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats SlbServerPortStats47Stats `json:"stats"`
    Name string 

}
type DataSlbServerPortStats47 struct {
    DtSlbServerPortStats47 SlbServerPortStats47 `json:"port"`
}


type SlbServerPortStats47Stats struct {
    Curr_conn int `json:"curr_conn"`
    Curr_req int `json:"curr_req"`
    Total_req int `json:"total_req"`
    Total_req_succ int `json:"total_req_succ"`
    Total_fwd_bytes int `json:"total_fwd_bytes"`
    Total_fwd_pkts int `json:"total_fwd_pkts"`
    Total_rev_bytes int `json:"total_rev_bytes"`
    Total_rev_pkts int `json:"total_rev_pkts"`
    Total_conn int `json:"total_conn"`
    Last_total_conn int `json:"last_total_conn"`
    Peak_conn int `json:"peak_conn"`
    Es_resp_200 int `json:"es_resp_200"`
    Es_resp_300 int `json:"es_resp_300"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_other int `json:"es_resp_other"`
    Es_req_count int `json:"es_req_count"`
    Es_resp_count int `json:"es_resp_count"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Total_rev_pkts_inspected int `json:"total_rev_pkts_inspected"`
    Total_rev_pkts_inspected_good_status_code int `json:"total_rev_pkts_inspected_good_status_code"`
    Response_time int `json:"response_time"`
    Fastest_rsp_time int `json:"fastest_rsp_time"`
    Slowest_rsp_time int `json:"slowest_rsp_time"`
    Curr_ssl_conn int `json:"curr_ssl_conn"`
    Total_ssl_conn int `json:"total_ssl_conn"`
    RespCount int `json:"resp-count"`
    Resp1xx int `json:"resp-1xx"`
    Resp2xx int `json:"resp-2xx"`
    Resp3xx int `json:"resp-3xx"`
    Resp4xx int `json:"resp-4xx"`
    Resp5xx int `json:"resp-5xx"`
    RespOther int `json:"resp-other"`
    RespLatency int `json:"resp-latency"`
    Curr_pconn int `json:"curr_pconn"`
}

func (p *SlbServerPortStats47) GetId() string{
    return "1"
}

func (p *SlbServerPortStats47) getPath() string{
    
    return "slb/server/" +p.Name + "/port/" +strconv.Itoa(p.PortNumber)+"+"+p.Protocol+"/stats"
}

func (p *SlbServerPortStats47) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbServerPortStats47,error) {
logger.Println("SlbServerPortStats47::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbServerPortStats47
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
