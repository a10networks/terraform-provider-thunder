

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6HttpAlgStats struct {
    
    Stats Cgnv6HttpAlgStatsStats `json:"stats"`

}
type DataCgnv6HttpAlgStats struct {
    DtCgnv6HttpAlgStats Cgnv6HttpAlgStats `json:"http-alg"`
}


type Cgnv6HttpAlgStatsStats struct {
    RequestProcessed int `json:"request-processed"`
    RequestInsertMsisdnPerformed int `json:"request-insert-msisdn-performed"`
    RequestInsertClientIpPerformed int `json:"request-insert-client-ip-performed"`
    RequestInsertMsisdnUnavailable int `json:"request-insert-msisdn-unavailable"`
    QueuedSessionTooMany int `json:"queued-session-too-many"`
    RadiusQuerySucceed int `json:"radius-query-succeed"`
    RadiusQueryFailed int `json:"radius-query-failed"`
    RadiusRequstSent int `json:"radius-requst-sent"`
    RadiusRequstDropped int `json:"radius-requst-dropped"`
    RadiusResponseReceived int `json:"radius-response-received"`
    RadiusResponseDropped int `json:"radius-response-dropped"`
    OutOfMemoryDropped int `json:"out-of-memory-dropped"`
    QueueLenExceedDropped int `json:"queue-len-exceed-dropped"`
    OutOfOrderDropped int `json:"out-of-order-dropped"`
    HeaderInsertionFailed int `json:"header-insertion-failed"`
    HeaderRemovalFailed int `json:"header-removal-failed"`
}

func (p *Cgnv6HttpAlgStats) GetId() string{
    return "1"
}

func (p *Cgnv6HttpAlgStats) getPath() string{
    return "cgnv6/http-alg/stats"
}

func (p *Cgnv6HttpAlgStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6HttpAlgStats,error) {
logger.Println("Cgnv6HttpAlgStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6HttpAlgStats
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
