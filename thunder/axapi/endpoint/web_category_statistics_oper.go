

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryStatisticsOper struct {
    
    Oper WebCategoryStatisticsOperOper `json:"oper"`

}
type DataWebCategoryStatisticsOper struct {
    DtWebCategoryStatisticsOper WebCategoryStatisticsOper `json:"statistics"`
}


type WebCategoryStatisticsOperOper struct {
    NumDplaneThreads int `json:"num-dplane-threads"`
    NumLookupThreads int `json:"num-lookup-threads"`
    PerCpuList []WebCategoryStatisticsOperOperPerCpuList `json:"per-cpu-list"`
    TotalReqQueue int `json:"total-req-queue"`
    TotalReqDropped int `json:"total-req-dropped"`
    TotalReqProcessed int `json:"total-req-processed"`
    TotalReqLookupProcessed int `json:"total-req-lookup-processed"`
    ClearCache string `json:"clear-cache"`
}


type WebCategoryStatisticsOperOperPerCpuList struct {
    ReqQueue int `json:"req-queue"`
    ReqDropped int `json:"req-dropped"`
    ReqProcessed int `json:"req-processed"`
    ReqLookupProcessed int `json:"req-lookup-processed"`
}

func (p *WebCategoryStatisticsOper) GetId() string{
    return "1"
}

func (p *WebCategoryStatisticsOper) getPath() string{
    return "web-category/statistics/oper"
}

func (p *WebCategoryStatisticsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryStatisticsOper,error) {
logger.Println("WebCategoryStatisticsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryStatisticsOper
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
