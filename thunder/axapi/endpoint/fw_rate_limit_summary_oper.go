

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwRateLimitSummaryOper struct {
    
    Oper FwRateLimitSummaryOperOper `json:"oper"`

}
type DataFwRateLimitSummaryOper struct {
    DtFwRateLimitSummaryOper FwRateLimitSummaryOper `json:"summary"`
}


type FwRateLimitSummaryOperOper struct {
    Mem_reserved int `json:"mem_reserved"`
    Mem_used int `json:"mem_used"`
    Alloc_failures int `json:"alloc_failures"`
    Total_num_entries int `json:"total_num_entries"`
    Total_entries_scope_aggregate int `json:"total_entries_scope_aggregate"`
    Total_entries_scope_subscriber_ip int `json:"total_entries_scope_subscriber_ip"`
    Total_entries_scope_subscriber_prefix int `json:"total_entries_scope_subscriber_prefix"`
    Total_entries_scope_parent int `json:"total_entries_scope_parent"`
    Total_entries_scope_parent_subscriber_ip int `json:"total_entries_scope_parent_subscriber_ip"`
    Total_entries_scope_parent_subscriberPrefix int `json:"total_entries_scope_parent_subscriber-prefix"`
}

func (p *FwRateLimitSummaryOper) GetId() string{
    return "1"
}

func (p *FwRateLimitSummaryOper) getPath() string{
    return "fw/rate-limit/summary/oper"
}

func (p *FwRateLimitSummaryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwRateLimitSummaryOper,error) {
logger.Println("FwRateLimitSummaryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwRateLimitSummaryOper
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
