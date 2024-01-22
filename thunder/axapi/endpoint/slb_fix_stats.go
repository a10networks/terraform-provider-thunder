

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbFixStats struct {
    
    Stats SlbFixStatsStats `json:"stats"`

}
type DataSlbFixStats struct {
    DtSlbFixStats SlbFixStats `json:"fix"`
}


type SlbFixStatsStats struct {
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Svrsel_fail int `json:"svrsel_fail"`
    Noroute int `json:"noroute"`
    Snat_fail int `json:"snat_fail"`
    Client_err int `json:"client_err"`
    Server_err int `json:"server_err"`
    Insert_clientip int `json:"insert_clientip"`
    Default_switching int `json:"default_switching"`
    Sender_switching int `json:"sender_switching"`
    Target_switching int `json:"target_switching"`
    Client_tls_conn int `json:"client_tls_conn"`
    Server_tls_conn int `json:"server_tls_conn"`
}

func (p *SlbFixStats) GetId() string{
    return "1"
}

func (p *SlbFixStats) getPath() string{
    return "slb/fix/stats"
}

func (p *SlbFixStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbFixStats,error) {
logger.Println("SlbFixStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbFixStats
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
