

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryUrlOper struct {
    
    Oper WebCategoryUrlOperOper `json:"oper"`

}
type DataWebCategoryUrlOper struct {
    DtWebCategoryUrlOper WebCategoryUrlOper `json:"url"`
}


type WebCategoryUrlOperOper struct {
    CategoryList []WebCategoryUrlOperOperCategoryList `json:"category-list"`
    Name string `json:"name"`
    LocalDbOnly int `json:"local-db-only"`
}


type WebCategoryUrlOperOperCategoryList struct {
    Category string `json:"category"`
}

func (p *WebCategoryUrlOper) GetId() string{
    return "1"
}

func (p *WebCategoryUrlOper) getPath() string{
    return "web-category/url/oper"
}

func (p *WebCategoryUrlOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryUrlOper,error) {
logger.Println("WebCategoryUrlOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryUrlOper
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
