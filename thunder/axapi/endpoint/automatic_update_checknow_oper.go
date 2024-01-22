

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AutomaticUpdateChecknowOper struct {
    
    Oper AutomaticUpdateChecknowOperOper `json:"oper"`

}
type DataAutomaticUpdateChecknowOper struct {
    DtAutomaticUpdateChecknowOper AutomaticUpdateChecknowOper `json:"checknow"`
}


type AutomaticUpdateChecknowOperOper struct {
    FeatureName string `json:"feature-name"`
    Result string `json:"result"`
}

func (p *AutomaticUpdateChecknowOper) GetId() string{
    return "1"
}

func (p *AutomaticUpdateChecknowOper) getPath() string{
    return "automatic-update/checknow/oper"
}

func (p *AutomaticUpdateChecknowOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAutomaticUpdateChecknowOper,error) {
logger.Println("AutomaticUpdateChecknowOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAutomaticUpdateChecknowOper
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
