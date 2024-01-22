

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwLimitEntryOper struct {
    
    Oper FwLimitEntryOperOper `json:"oper"`

}
type DataFwLimitEntryOper struct {
    DtFwLimitEntryOper FwLimitEntryOper `json:"limit-entry"`
}


type FwLimitEntryOperOper struct {
    LimitEntryList []FwLimitEntryOperOperLimitEntryList `json:"limit-entry-list"`
    LimitEntryCount int `json:"limit-entry-count"`
    Prefix6 string `json:"prefix6"`
    Prefix4 string `json:"prefix4"`
    PrefixLen6 int `json:"prefix-len6"`
    PrefixLen4 int `json:"prefix-len4"`
}


type FwLimitEntryOperOperLimitEntryList struct {
    Address string `json:"address"`
    PrefixLen int `json:"prefix-len"`
    RuleName string `json:"rule-name"`
    CurrCount int `json:"curr-count"`
    MaxCount int `json:"max-count"`
    Type string `json:"type"`
}

func (p *FwLimitEntryOper) GetId() string{
    return "1"
}

func (p *FwLimitEntryOper) getPath() string{
    return "fw/limit-entry/oper"
}

func (p *FwLimitEntryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwLimitEntryOper,error) {
logger.Println("FwLimitEntryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwLimitEntryOper
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
