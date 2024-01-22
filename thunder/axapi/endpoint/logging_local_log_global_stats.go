

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogGlobalStats struct {
    
    Stats LoggingLocalLogGlobalStatsStats `json:"stats"`

}
type DataLoggingLocalLogGlobalStats struct {
    DtLoggingLocalLogGlobalStats LoggingLocalLogGlobalStats `json:"global"`
}


type LoggingLocalLogGlobalStatsStats struct {
    Enqueue int `json:"enqueue"`
    EnqueueFull int `json:"enqueue-full"`
    EnqueueError int `json:"enqueue-error"`
    Dequeue int `json:"dequeue"`
    DequeueError int `json:"dequeue-error"`
    RawLog int `json:"raw-log"`
    RawLogError int `json:"raw-log-error"`
    LogSummarized int `json:"log-summarized"`
    L1LogSummarized int `json:"l1-log-summarized"`
    L2LogSummarized int `json:"l2-log-summarized"`
    LogSummarizedError int `json:"log-summarized-error"`
    AamDb int `json:"aam-db"`
    EpDb int `json:"ep-db"`
    FwDb int `json:"fw-db"`
    AamTopUserDb int `json:"aam-top-user-db"`
    EpTopUserDb int `json:"ep-top-user-db"`
    EpTopSrcDb int `json:"ep-top-src-db"`
    EpTopDstDb int `json:"ep-top-dst-db"`
    EpTopDomainDb int `json:"ep-top-domain-db"`
    EpTopWebCategoryDb int `json:"ep-top-web-category-db"`
    EpTopHostDb int `json:"ep-top-host-db"`
    FwTopAppDb int `json:"fw-top-app-db"`
    FwTopSrcDb int `json:"fw-top-src-db"`
    FwTopAppSrcDb int `json:"fw-top-app-src-db"`
    FwTopCategoryDb int `json:"fw-top-category-db"`
    DbErro int `json:"db-erro"`
    Query int `json:"query"`
    Response int `json:"response"`
    QueryError int `json:"query-error"`
    FwTopThrDb int `json:"fw-top-thr-db"`
    FwTopThrSrcDb int `json:"fw-top-thr-src-db"`
}

func (p *LoggingLocalLogGlobalStats) GetId() string{
    return "1"
}

func (p *LoggingLocalLogGlobalStats) getPath() string{
    return "logging/local-log/global/stats"
}

func (p *LoggingLocalLogGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataLoggingLocalLogGlobalStats,error) {
logger.Println("LoggingLocalLogGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataLoggingLocalLogGlobalStats
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
