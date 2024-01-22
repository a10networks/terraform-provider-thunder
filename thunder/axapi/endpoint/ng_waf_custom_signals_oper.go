

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NgWafCustomSignalsOper struct {
    
    Oper NgWafCustomSignalsOperOper `json:"oper"`

}
type DataNgWafCustomSignalsOper struct {
    DtNgWafCustomSignalsOper NgWafCustomSignalsOper `json:"custom-signals"`
}


type NgWafCustomSignalsOperOper struct {
    SignalList []NgWafCustomSignalsOperOperSignalList `json:"signal-list"`
}


type NgWafCustomSignalsOperOperSignalList struct {
    Signal string `json:"signal"`
}

func (p *NgWafCustomSignalsOper) GetId() string{
    return "1"
}

func (p *NgWafCustomSignalsOper) getPath() string{
    return "ng-waf/custom-signals/oper"
}

func (p *NgWafCustomSignalsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNgWafCustomSignalsOper,error) {
logger.Println("NgWafCustomSignalsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNgWafCustomSignalsOper
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
