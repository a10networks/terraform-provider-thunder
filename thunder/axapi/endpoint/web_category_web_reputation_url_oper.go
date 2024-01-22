

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryWebReputationUrlOper struct {
    
    Oper WebCategoryWebReputationUrlOperOper `json:"oper"`

}
type DataWebCategoryWebReputationUrlOper struct {
    DtWebCategoryWebReputationUrlOper WebCategoryWebReputationUrlOper `json:"url"`
}


type WebCategoryWebReputationUrlOperOper struct {
    ReputationScore string `json:"reputation-score"`
    Name string `json:"name"`
    LocalDbOnly int `json:"local-db-only"`
}

func (p *WebCategoryWebReputationUrlOper) GetId() string{
    return "1"
}

func (p *WebCategoryWebReputationUrlOper) getPath() string{
    return "web-category/web-reputation/url/oper"
}

func (p *WebCategoryWebReputationUrlOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryWebReputationUrlOper,error) {
logger.Println("WebCategoryWebReputationUrlOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryWebReputationUrlOper
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
