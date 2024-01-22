

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbCrlSrcipStats struct {
    
    Stats SlbCrlSrcipStatsStats `json:"stats"`

}
type DataSlbCrlSrcipStats struct {
    DtSlbCrlSrcipStats SlbCrlSrcipStats `json:"crl-srcip"`
}


type SlbCrlSrcipStatsStats struct {
    Sessions_alloc int `json:"sessions_alloc"`
    Sessions_freed int `json:"sessions_freed"`
    Out_of_sessions int `json:"out_of_sessions"`
    Too_many_sessions int `json:"too_many_sessions"`
    Called int `json:"called"`
    Permitted int `json:"permitted"`
    Threshold_exceed int `json:"threshold_exceed"`
    Lockout_drop int `json:"lockout_drop"`
    Log_msg_sent int `json:"log_msg_sent"`
}

func (p *SlbCrlSrcipStats) GetId() string{
    return "1"
}

func (p *SlbCrlSrcipStats) getPath() string{
    return "slb/crl-srcip/stats"
}

func (p *SlbCrlSrcipStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbCrlSrcipStats,error) {
logger.Println("SlbCrlSrcipStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbCrlSrcipStats
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
