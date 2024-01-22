

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryWebReputationOper struct {
    
    BypassedUrls WebCategoryWebReputationOperBypassedUrls `json:"bypassed-urls"`
    InterceptedUrls WebCategoryWebReputationOperInterceptedUrls `json:"intercepted-urls"`
    Oper WebCategoryWebReputationOperOper `json:"oper"`
    Url WebCategoryWebReputationOperUrl `json:"url"`

}
type DataWebCategoryWebReputationOper struct {
    DtWebCategoryWebReputationOper WebCategoryWebReputationOper `json:"web-reputation"`
}


type WebCategoryWebReputationOperBypassedUrls struct {
    Oper WebCategoryWebReputationOperBypassedUrlsOper `json:"oper"`
}


type WebCategoryWebReputationOperBypassedUrlsOper struct {
    UrlList []WebCategoryWebReputationOperBypassedUrlsOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryWebReputationOperBypassedUrlsOperUrlList struct {
    UrlName string `json:"url-name"`
}


type WebCategoryWebReputationOperInterceptedUrls struct {
    Oper WebCategoryWebReputationOperInterceptedUrlsOper `json:"oper"`
}


type WebCategoryWebReputationOperInterceptedUrlsOper struct {
    UrlList []WebCategoryWebReputationOperInterceptedUrlsOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryWebReputationOperInterceptedUrlsOperUrlList struct {
    UrlName string `json:"url-name"`
}


type WebCategoryWebReputationOperOper struct {
    UrlList []WebCategoryWebReputationOperOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryWebReputationOperOperUrlList struct {
    UrlName string `json:"url-name"`
}


type WebCategoryWebReputationOperUrl struct {
    Oper WebCategoryWebReputationOperUrlOper `json:"oper"`
}


type WebCategoryWebReputationOperUrlOper struct {
    ReputationScore string `json:"reputation-score"`
    Name string `json:"name"`
    LocalDbOnly int `json:"local-db-only"`
}

func (p *WebCategoryWebReputationOper) GetId() string{
    return "1"
}

func (p *WebCategoryWebReputationOper) getPath() string{
    return "web-category/web-reputation/oper"
}

func (p *WebCategoryWebReputationOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryWebReputationOper,error) {
logger.Println("WebCategoryWebReputationOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryWebReputationOper
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
