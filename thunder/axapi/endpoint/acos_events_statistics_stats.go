

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsStatisticsStats struct {
    
    Stats AcosEventsStatisticsStatsStats `json:"stats"`

}
type DataAcosEventsStatisticsStats struct {
    DtAcosEventsStatisticsStats AcosEventsStatisticsStats `json:"statistics"`
}


type AcosEventsStatisticsStatsStats struct {
    Msg_sent int `json:"msg_sent"`
    Msg_sent_logdb int `json:"msg_sent_logdb"`
    Msg_dropped_format_not_defined int `json:"msg_dropped_format_not_defined"`
    Msg_dropped_malloc_failure int `json:"msg_dropped_malloc_failure"`
    Msg_dropped_no_template int `json:"msg_dropped_no_template"`
    Msg_dropped_selector int `json:"msg_dropped_selector"`
    Msg_dropped_too_long int `json:"msg_dropped_too_long"`
    Msg_dropped_craft_fail int `json:"msg_dropped_craft_fail"`
    Msg_dropped_local_log_ratelimit int `json:"msg_dropped_local_log_ratelimit"`
    Msg_dropped_send_failed int `json:"msg_dropped_send_failed"`
    Msg_dropped_no_active_member int `json:"msg_dropped_no_active_member"`
    Msg_dropped_route_fail int `json:"msg_dropped_route_fail"`
    Msg_dropped_other int `json:"msg_dropped_other"`
    Param_msg_sent_to_hc int `json:"param_msg_sent_to_hc"`
    Param_msg_sent_fail int `json:"param_msg_sent_fail"`
    Param_msg_encode_fail int `json:"param_msg_encode_fail"`
}

func (p *AcosEventsStatisticsStats) GetId() string{
    return "1"
}

func (p *AcosEventsStatisticsStats) getPath() string{
    return "acos-events/statistics/stats"
}

func (p *AcosEventsStatisticsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAcosEventsStatisticsStats,error) {
logger.Println("AcosEventsStatisticsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAcosEventsStatisticsStats
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
