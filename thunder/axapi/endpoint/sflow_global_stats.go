

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SflowGlobalStats struct {
    
    Stats SflowGlobalStatsStats `json:"stats"`

}
type DataSflowGlobalStats struct {
    DtSflowGlobalStats SflowGlobalStats `json:"global"`
}


type SflowGlobalStatsStats struct {
    TotalPacketSampleRecords int `json:"total-packet-sample-records"`
    TotalCounterSampleRecords int `json:"total-counter-sample-records"`
    TotalSflowPacketsSent int `json:"total-sflow-packets-sent"`
    TotalSflowLocalPacketsSent int `json:"total-sflow-local-packets-sent"`
    TotalSflowPacketsSentMgmt int `json:"total-sflow-packets-sent-mgmt"`
    TotalSflowPacketsDropMgmt int `json:"total-sflow-packets-drop-mgmt"`
}

func (p *SflowGlobalStats) GetId() string{
    return "1"
}

func (p *SflowGlobalStats) getPath() string{
    return "sflow/global/stats"
}

func (p *SflowGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSflowGlobalStats,error) {
logger.Println("SflowGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSflowGlobalStats
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
