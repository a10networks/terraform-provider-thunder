

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHwCompressStats struct {
    
    Stats SlbHwCompressStatsStats `json:"stats"`

}
type DataSlbHwCompressStats struct {
    DtSlbHwCompressStats SlbHwCompressStats `json:"hw-compress"`
}


type SlbHwCompressStatsStats struct {
    Request_count int `json:"request_count"`
    Submit_count int `json:"submit_count"`
    Response_count int `json:"response_count"`
    Failure_count int `json:"failure_count"`
    Failure_code int `json:"failure_code"`
    Ring_full_count int `json:"ring_full_count"`
    Max_outstanding_request_count int `json:"max_outstanding_request_count"`
    Max_outstanding_submit_count int `json:"max_outstanding_submit_count"`
}

func (p *SlbHwCompressStats) GetId() string{
    return "1"
}

func (p *SlbHwCompressStats) getPath() string{
    return "slb/hw-compress/stats"
}

func (p *SlbHwCompressStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHwCompressStats,error) {
logger.Println("SlbHwCompressStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHwCompressStats
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
