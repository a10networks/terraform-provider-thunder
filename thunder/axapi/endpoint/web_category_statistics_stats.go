

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryStatisticsStats struct {
    
    Stats WebCategoryStatisticsStatsStats `json:"stats"`

}
type DataWebCategoryStatisticsStats struct {
    DtWebCategoryStatisticsStats WebCategoryStatisticsStats `json:"statistics"`
}


type WebCategoryStatisticsStatsStats struct {
    DbLookup int `json:"db-lookup"`
    CloudCacheLookup int `json:"cloud-cache-lookup"`
    CloudLookup int `json:"cloud-lookup"`
    RtuLookup int `json:"rtu-lookup"`
    LookupLatency int `json:"lookup-latency"`
    DbMem int `json:"db-mem"`
    RtuCacheMem int `json:"rtu-cache-mem"`
    LookupCacheMem int `json:"lookup-cache-mem"`
}

func (p *WebCategoryStatisticsStats) GetId() string{
    return "1"
}

func (p *WebCategoryStatisticsStats) getPath() string{
    return "web-category/statistics/stats"
}

func (p *WebCategoryStatisticsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryStatisticsStats,error) {
logger.Println("WebCategoryStatisticsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryStatisticsStats
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
