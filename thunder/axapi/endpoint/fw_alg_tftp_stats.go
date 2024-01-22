

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwAlgTftpStats struct {
    
    Stats FwAlgTftpStatsStats `json:"stats"`

}
type DataFwAlgTftpStats struct {
    DtFwAlgTftpStats FwAlgTftpStats `json:"tftp"`
}


type FwAlgTftpStatsStats struct {
    SessionCreated int `json:"session-created"`
}

func (p *FwAlgTftpStats) GetId() string{
    return "1"
}

func (p *FwAlgTftpStats) getPath() string{
    return "fw/alg/tftp/stats"
}

func (p *FwAlgTftpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwAlgTftpStats,error) {
logger.Println("FwAlgTftpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwAlgTftpStats
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
