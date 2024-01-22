

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbMysqlStats struct {
    
    Stats SlbMysqlStatsStats `json:"stats"`

}
type DataSlbMysqlStats struct {
    DtSlbMysqlStats SlbMysqlStats `json:"mysql"`
}


type SlbMysqlStatsStats struct {
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Curr_be_enc int `json:"curr_be_enc"`
    Total_be_enc int `json:"total_be_enc"`
    Curr_fe_enc int `json:"curr_fe_enc"`
    Total_fe_enc int `json:"total_fe_enc"`
    Client_fin int `json:"client_fin"`
    Server_fin int `json:"server_fin"`
    Session_err int `json:"session_err"`
    Queries int `json:"queries"`
    Commands int `json:"commands"`
}

func (p *SlbMysqlStats) GetId() string{
    return "1"
}

func (p *SlbMysqlStats) getPath() string{
    return "slb/mysql/stats"
}

func (p *SlbMysqlStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbMysqlStats,error) {
logger.Println("SlbMysqlStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbMysqlStats
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
