

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryWebReputationInterceptedUrlsOper struct {
    
    Oper WebCategoryWebReputationInterceptedUrlsOperOper `json:"oper"`

}
type DataWebCategoryWebReputationInterceptedUrlsOper struct {
    DtWebCategoryWebReputationInterceptedUrlsOper WebCategoryWebReputationInterceptedUrlsOper `json:"intercepted-urls"`
}


type WebCategoryWebReputationInterceptedUrlsOperOper struct {
    UrlList []WebCategoryWebReputationInterceptedUrlsOperOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryWebReputationInterceptedUrlsOperOperUrlList struct {
    UrlName string `json:"url-name"`
}

func (p *WebCategoryWebReputationInterceptedUrlsOper) GetId() string{
    return "1"
}

func (p *WebCategoryWebReputationInterceptedUrlsOper) getPath() string{
    return "web-category/web-reputation/intercepted-urls/oper"
}

func (p *WebCategoryWebReputationInterceptedUrlsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryWebReputationInterceptedUrlsOper,error) {
logger.Println("WebCategoryWebReputationInterceptedUrlsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryWebReputationInterceptedUrlsOper
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
