

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryBypassedUrlsOper struct {
    
    Oper WebCategoryBypassedUrlsOperOper `json:"oper"`

}
type DataWebCategoryBypassedUrlsOper struct {
    DtWebCategoryBypassedUrlsOper WebCategoryBypassedUrlsOper `json:"bypassed-urls"`
}


type WebCategoryBypassedUrlsOperOper struct {
    UrlList []WebCategoryBypassedUrlsOperOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryBypassedUrlsOperOperUrlList struct {
    UrlName string `json:"url-name"`
}

func (p *WebCategoryBypassedUrlsOper) GetId() string{
    return "1"
}

func (p *WebCategoryBypassedUrlsOper) getPath() string{
    return "web-category/bypassed-urls/oper"
}

func (p *WebCategoryBypassedUrlsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryBypassedUrlsOper,error) {
logger.Println("WebCategoryBypassedUrlsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryBypassedUrlsOper
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
