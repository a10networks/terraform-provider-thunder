

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryOper struct {
    
    BypassedUrls WebCategoryOperBypassedUrls `json:"bypassed-urls"`
    InterceptedUrls WebCategoryOperInterceptedUrls `json:"intercepted-urls"`
    License WebCategoryOperLicense `json:"license"`
    Oper WebCategoryOperOper `json:"oper"`
    Statistics WebCategoryOperStatistics `json:"statistics"`
    Url WebCategoryOperUrl `json:"url"`
    WebReputation WebCategoryOperWebReputation `json:"web-reputation"`

}
type DataWebCategoryOper struct {
    DtWebCategoryOper WebCategoryOper `json:"web-category"`
}


type WebCategoryOperBypassedUrls struct {
    Oper WebCategoryOperBypassedUrlsOper `json:"oper"`
}


type WebCategoryOperBypassedUrlsOper struct {
    UrlList []WebCategoryOperBypassedUrlsOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryOperBypassedUrlsOperUrlList struct {
    UrlName string `json:"url-name"`
}


type WebCategoryOperInterceptedUrls struct {
    Oper WebCategoryOperInterceptedUrlsOper `json:"oper"`
}


type WebCategoryOperInterceptedUrlsOper struct {
    UrlList []WebCategoryOperInterceptedUrlsOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryOperInterceptedUrlsOperUrlList struct {
    UrlName string `json:"url-name"`
}


type WebCategoryOperLicense struct {
    Oper WebCategoryOperLicenseOper `json:"oper"`
}


type WebCategoryOperLicenseOper struct {
    ModuleStatus string `json:"module-status"`
    LicenseStatus string `json:"license-status"`
    LicenseType string `json:"license-type"`
    LicenseExpiry string `json:"license-expiry"`
    RemainingPeriod string `json:"remaining-period"`
    IsGrace string `json:"is-grace"`
    GracePeriod string `json:"grace-period"`
    SerialNumber string `json:"serial-number"`
}


type WebCategoryOperOper struct {
    WebCatVersion string `json:"web-cat-version"`
    WebCatDatabaseName string `json:"web-cat-database-name"`
    WebCatDatabaseStatus string `json:"web-cat-database-status"`
    WebCatDatabaseSize string `json:"web-cat-database-size"`
    WebCatDatabaseVersion int `json:"web-cat-database-version"`
    WebCatLastUpdateTime string `json:"web-cat-last-update-time"`
    WebCatNextUpdateTime string `json:"web-cat-next-update-time"`
    WebCatConnectionStatus string `json:"web-cat-connection-status"`
    WebCatFailureReason string `json:"web-cat-failure-reason"`
    WebCatLastSuccessfulConnection string `json:"web-cat-last-successful-connection"`
}


type WebCategoryOperStatistics struct {
    Oper WebCategoryOperStatisticsOper `json:"oper"`
}


type WebCategoryOperStatisticsOper struct {
    NumDplaneThreads int `json:"num-dplane-threads"`
    NumLookupThreads int `json:"num-lookup-threads"`
    PerCpuList []WebCategoryOperStatisticsOperPerCpuList `json:"per-cpu-list"`
    TotalReqQueue int `json:"total-req-queue"`
    TotalReqDropped int `json:"total-req-dropped"`
    TotalReqProcessed int `json:"total-req-processed"`
    TotalReqLookupProcessed int `json:"total-req-lookup-processed"`
    ClearCache string `json:"clear-cache"`
}


type WebCategoryOperStatisticsOperPerCpuList struct {
    ReqQueue int `json:"req-queue"`
    ReqDropped int `json:"req-dropped"`
    ReqProcessed int `json:"req-processed"`
    ReqLookupProcessed int `json:"req-lookup-processed"`
}


type WebCategoryOperUrl struct {
    Oper WebCategoryOperUrlOper `json:"oper"`
}


type WebCategoryOperUrlOper struct {
    CategoryList []WebCategoryOperUrlOperCategoryList `json:"category-list"`
    Name string `json:"name"`
    LocalDbOnly int `json:"local-db-only"`
}


type WebCategoryOperUrlOperCategoryList struct {
    Category string `json:"category"`
}


type WebCategoryOperWebReputation struct {
    Oper WebCategoryOperWebReputationOper `json:"oper"`
    InterceptedUrls WebCategoryOperWebReputationInterceptedUrls `json:"intercepted-urls"`
    BypassedUrls WebCategoryOperWebReputationBypassedUrls `json:"bypassed-urls"`
    Url WebCategoryOperWebReputationUrl `json:"url"`
}


type WebCategoryOperWebReputationOper struct {
    UrlList []WebCategoryOperWebReputationOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryOperWebReputationOperUrlList struct {
    UrlName string `json:"url-name"`
}


type WebCategoryOperWebReputationInterceptedUrls struct {
    Oper WebCategoryOperWebReputationInterceptedUrlsOper `json:"oper"`
}


type WebCategoryOperWebReputationInterceptedUrlsOper struct {
    UrlList []WebCategoryOperWebReputationInterceptedUrlsOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryOperWebReputationInterceptedUrlsOperUrlList struct {
    UrlName string `json:"url-name"`
}


type WebCategoryOperWebReputationBypassedUrls struct {
    Oper WebCategoryOperWebReputationBypassedUrlsOper `json:"oper"`
}


type WebCategoryOperWebReputationBypassedUrlsOper struct {
    UrlList []WebCategoryOperWebReputationBypassedUrlsOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryOperWebReputationBypassedUrlsOperUrlList struct {
    UrlName string `json:"url-name"`
}


type WebCategoryOperWebReputationUrl struct {
    Oper WebCategoryOperWebReputationUrlOper `json:"oper"`
}


type WebCategoryOperWebReputationUrlOper struct {
    ReputationScore string `json:"reputation-score"`
    Name string `json:"name"`
    LocalDbOnly int `json:"local-db-only"`
}

func (p *WebCategoryOper) GetId() string{
    return "1"
}

func (p *WebCategoryOper) getPath() string{
    return "web-category/oper"
}

func (p *WebCategoryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryOper,error) {
logger.Println("WebCategoryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryOper
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
