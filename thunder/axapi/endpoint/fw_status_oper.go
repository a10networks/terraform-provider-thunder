

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwStatusOper struct {
    
    Oper FwStatusOperOper `json:"oper"`

}
type DataFwStatusOper struct {
    DtFwStatusOper FwStatusOper `json:"status"`
}


type FwStatusOperOper struct {
    CurrentActiveRuleSet string `json:"current-active-rule-set"`
    PreviousSuccessfulCompilationAttempt string `json:"previous-successful-compilation-attempt"`
    PreviousSuccessfulCompilationDuration string `json:"previous-successful-compilation-duration"`
    MostRecentCompilationAttempt string `json:"most-recent-compilation-attempt"`
    MostRecentCompilationStatus string `json:"most-recent-compilation-status"`
    Internal int `json:"internal"`
}

func (p *FwStatusOper) GetId() string{
    return "1"
}

func (p *FwStatusOper) getPath() string{
    return "fw/status/oper"
}

func (p *FwStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwStatusOper,error) {
logger.Println("FwStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwStatusOper
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
