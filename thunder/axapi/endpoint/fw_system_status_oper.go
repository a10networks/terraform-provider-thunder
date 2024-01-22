

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwSystemStatusOper struct {
    
    Oper FwSystemStatusOperOper `json:"oper"`

}
type DataFwSystemStatusOper struct {
    DtFwSystemStatusOper FwSystemStatusOper `json:"system-status"`
}


type FwSystemStatusOperOper struct {
    DataSessionsUsed int `json:"data-sessions-used"`
    DataSessionsFree int `json:"data-sessions-free"`
    SmpSessionsUsed int `json:"smp-sessions-used"`
    SmpSessionsFree int `json:"smp-sessions-free"`
    RadiusEntriesUsed int `json:"radius-entries-used"`
    RadiusEntriesFree int `json:"radius-entries-free"`
}

func (p *FwSystemStatusOper) GetId() string{
    return "1"
}

func (p *FwSystemStatusOper) getPath() string{
    return "fw/system-status/oper"
}

func (p *FwSystemStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwSystemStatusOper,error) {
logger.Println("FwSystemStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwSystemStatusOper
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
