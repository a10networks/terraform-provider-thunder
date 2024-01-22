

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwResourceUsageOper struct {
    
    Oper FwResourceUsageOperOper `json:"oper"`

}
type DataFwResourceUsageOper struct {
    DtFwResourceUsageOper FwResourceUsageOper `json:"resource-usage"`
}


type FwResourceUsageOperOper struct {
    FwObjectCurrentCount int `json:"fw-object-current-count"`
    FwObjectTotalCount int `json:"fw-object-total-count"`
    FwObjectGroupCurrentCount int `json:"fw-object-group-current-count"`
    FwObjectGroupTotalCount int `json:"fw-object-group-total-count"`
    FwRuleSetCurrentCount int `json:"fw-rule-set-current-count"`
    FwRuleSetTotalCount int `json:"fw-rule-set-total-count"`
    FwRuleCurrentCount int `json:"fw-rule-current-count"`
    FwRuleTotalCount int `json:"fw-rule-total-count"`
    FwZoneCurrentCount int `json:"fw-zone-current-count"`
    FwZoneTotalCount int `json:"fw-zone-total-count"`
    FwIpRangeCurrentCount int `json:"fw-ip-range-current-count"`
    FwIpRangeTotalCount int `json:"fw-ip-range-total-count"`
    FwHelperSessionsCurrentCount int `json:"fw-helper-sessions-current-count"`
    FwHelperSessionsTotalCount int `json:"fw-helper-sessions-total-count"`
    RadiusTableCurrentCount int `json:"radius-table-current-count"`
    RadiusTableTotalCount int `json:"radius-table-total-count"`
    ClausePerObjGrpCurrentCount string `json:"clause-per-obj-grp-current-count"`
    ClausePerObjGrpTotalCount int `json:"clause-per-obj-grp-total-count"`
    Object int `json:"object"`
    ObjectGroup int `json:"object-group"`
    RuleSet int `json:"rule-set"`
    Rule int `json:"rule"`
    Zone int `json:"zone"`
    IpRange int `json:"ip-range"`
    HelperSessions int `json:"helper-sessions"`
    RadiusTableSize int `json:"radius-table-size"`
    ClausePerObjGrp int `json:"clause-per-obj-grp"`
}

func (p *FwResourceUsageOper) GetId() string{
    return "1"
}

func (p *FwResourceUsageOper) getPath() string{
    return "fw/resource-usage/oper"
}

func (p *FwResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwResourceUsageOper,error) {
logger.Println("FwResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwResourceUsageOper
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
