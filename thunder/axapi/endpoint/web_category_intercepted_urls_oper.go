

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryInterceptedUrlsOper struct {
    
    Oper WebCategoryInterceptedUrlsOperOper `json:"oper"`

}
type DataWebCategoryInterceptedUrlsOper struct {
    DtWebCategoryInterceptedUrlsOper WebCategoryInterceptedUrlsOper `json:"intercepted-urls"`
}


type WebCategoryInterceptedUrlsOperOper struct {
    UrlList []WebCategoryInterceptedUrlsOperOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryInterceptedUrlsOperOperUrlList struct {
    UrlName string `json:"url-name"`
}

func (p *WebCategoryInterceptedUrlsOper) GetId() string{
    return "1"
}

func (p *WebCategoryInterceptedUrlsOper) getPath() string{
    return "web-category/intercepted-urls/oper"
}

func (p *WebCategoryInterceptedUrlsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryInterceptedUrlsOper,error) {
logger.Println("WebCategoryInterceptedUrlsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryInterceptedUrlsOper
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
