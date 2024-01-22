

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AutomaticUpdateInfoOper struct {
    
    Oper AutomaticUpdateInfoOperOper `json:"oper"`

}
type DataAutomaticUpdateInfoOper struct {
    DtAutomaticUpdateInfoOper AutomaticUpdateInfoOper `json:"info"`
}


type AutomaticUpdateInfoOperOper struct {
    FeatureList []AutomaticUpdateInfoOperOperFeatureList `json:"feature-list"`
}


type AutomaticUpdateInfoOperOperFeatureList struct {
    FeatureName string `json:"feature-name"`
    Version string `json:"version"`
    Schedule string `json:"schedule"`
    Time string `json:"time"`
    LastUpdate string `json:"last-update"`
    NextCheck string `json:"next-check"`
}

func (p *AutomaticUpdateInfoOper) GetId() string{
    return "1"
}

func (p *AutomaticUpdateInfoOper) getPath() string{
    return "automatic-update/info/oper"
}

func (p *AutomaticUpdateInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAutomaticUpdateInfoOper,error) {
logger.Println("AutomaticUpdateInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAutomaticUpdateInfoOper
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
