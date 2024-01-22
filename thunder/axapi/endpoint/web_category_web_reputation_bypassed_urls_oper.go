

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryWebReputationBypassedUrlsOper struct {
    
    Oper WebCategoryWebReputationBypassedUrlsOperOper `json:"oper"`

}
type DataWebCategoryWebReputationBypassedUrlsOper struct {
    DtWebCategoryWebReputationBypassedUrlsOper WebCategoryWebReputationBypassedUrlsOper `json:"bypassed-urls"`
}


type WebCategoryWebReputationBypassedUrlsOperOper struct {
    UrlList []WebCategoryWebReputationBypassedUrlsOperOperUrlList `json:"url-list"`
    NumberOfUrls int `json:"number-of-urls"`
    AllUrls string `json:"all-urls"`
    UrlName string `json:"url-name"`
}


type WebCategoryWebReputationBypassedUrlsOperOperUrlList struct {
    UrlName string `json:"url-name"`
}

func (p *WebCategoryWebReputationBypassedUrlsOper) GetId() string{
    return "1"
}

func (p *WebCategoryWebReputationBypassedUrlsOper) getPath() string{
    return "web-category/web-reputation/bypassed-urls/oper"
}

func (p *WebCategoryWebReputationBypassedUrlsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryWebReputationBypassedUrlsOper,error) {
logger.Println("WebCategoryWebReputationBypassedUrlsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryWebReputationBypassedUrlsOper
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
