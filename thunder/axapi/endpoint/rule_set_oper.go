

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetOper struct {
    
    Application RuleSetOperApplication `json:"application"`
    Name string `json:"name"`
    Oper RuleSetOperOper `json:"oper"`
    RuleList []RuleSetOperRuleList `json:"rule-list"`
    RulesByZone RuleSetOperRulesByZone `json:"rules-by-zone"`
    TrackAppRuleList RuleSetOperTrackAppRuleList `json:"track-app-rule-list"`

}
type DataRuleSetOper struct {
    DtRuleSetOper RuleSetOper `json:"rule-set"`
}


type RuleSetOperApplication struct {
    Oper RuleSetOperApplicationOper `json:"oper"`
}


type RuleSetOperApplicationOper struct {
    CategoryStat string `json:"category-stat"`
    AppStat string `json:"app-stat"`
    Protocol int `json:"protocol"`
    Rule string `json:"rule"`
    RuleSetOnly int `json:"rule-set-only"`
    RuleList []RuleSetOperApplicationOperRuleList `json:"rule-list"`
}


type RuleSetOperApplicationOperRuleList struct {
    Name string `json:"name"`
    StatList []RuleSetOperApplicationOperRuleListStatList `json:"stat-list"`
}


type RuleSetOperApplicationOperRuleListStatList struct {
    Name string `json:"name"`
    Category string `json:"category"`
    Type string `json:"type"`
    Conns int `json:"conns"`
    Bytes int `json:"bytes"`
}


type RuleSetOperOper struct {
    PolicyStatus string `json:"policy-status"`
    PolicyUnmatchedDrop int `json:"policy-unmatched-drop"`
    PolicyPermit int `json:"policy-permit"`
    PolicyDeny int `json:"policy-deny"`
    PolicyReset int `json:"policy-reset"`
    PolicyRuleCount int `json:"policy-rule-count"`
    RuleStats []RuleSetOperOperRuleStats `json:"rule-stats"`
    TotalHit int `json:"total-hit"`
    TotalPermitBytes int `json:"total-permit-bytes"`
    TotalDenyBytes int `json:"total-deny-bytes"`
    TotalResetBytes int `json:"total-reset-bytes"`
    TotalBytes int `json:"total-bytes"`
    TotalPermitPackets int `json:"total-permit-packets"`
    TotalDenyPackets int `json:"total-deny-packets"`
    TotalResetPackets int `json:"total-reset-packets"`
    TotalPackets int `json:"total-packets"`
    TotalActiveTcp int `json:"total-active-tcp"`
    TotalActiveUdp int `json:"total-active-udp"`
    TotalActiveIcmp int `json:"total-active-icmp"`
    TotalActiveOthers int `json:"total-active-others"`
    ShowTotalStats string `json:"show-total-stats"`
    TopnRules string `json:"topn-rules"`
}


type RuleSetOperOperRuleStats struct {
    RuleName string `json:"rule-name"`
    RuleHitcount int `json:"rule-hitcount"`
    RuleAction string `json:"rule-action"`
    RuleStatus string `json:"rule-status"`
}


type RuleSetOperRuleList struct {
    Name string `json:"name"`
    Oper RuleSetOperRuleListOper `json:"oper"`
}


type RuleSetOperRuleListOper struct {
    Hitcount int `json:"hitcount"`
    LastHitcountTime string `json:"last-hitcount-time"`
    Action string `json:"action"`
    Status string `json:"status"`
    Permitbytes int `json:"permitbytes"`
    Denybytes int `json:"denybytes"`
    Resetbytes int `json:"resetbytes"`
    Totalbytes int `json:"totalbytes"`
    Permitpackets int `json:"permitpackets"`
    Denypackets int `json:"denypackets"`
    Resetpackets int `json:"resetpackets"`
    Totalpackets int `json:"totalpackets"`
    Activesessiontcp int `json:"activesessiontcp"`
    Activesessionudp int `json:"activesessionudp"`
    Activesessionicmp int `json:"activesessionicmp"`
    Activesessionsctp int `json:"activesessionsctp"`
    Activesessionother int `json:"activesessionother"`
    Activesessiontotal int `json:"activesessiontotal"`
    Sessiontcp int `json:"sessiontcp"`
    Sessionudp int `json:"sessionudp"`
    Sessionicmp int `json:"sessionicmp"`
    Sessionsctp int `json:"sessionsctp"`
    Sessionother int `json:"sessionother"`
    Sessiontotal int `json:"sessiontotal"`
    Ratelimitdrops int `json:"ratelimitdrops"`
}


type RuleSetOperRulesByZone struct {
    Oper RuleSetOperRulesByZoneOper `json:"oper"`
}


type RuleSetOperRulesByZoneOper struct {
    GroupList []RuleSetOperRulesByZoneOperGroupList `json:"group-list"`
}


type RuleSetOperRulesByZoneOperGroupList struct {
    From string `json:"from"`
    To string `json:"to"`
    RuleList []RuleSetOperRulesByZoneOperGroupListRuleList `json:"rule-list"`
}


type RuleSetOperRulesByZoneOperGroupListRuleList struct {
    Name string `json:"name"`
    Action string `json:"action"`
    SourceList []RuleSetOperRulesByZoneOperGroupListRuleListSourceList `json:"source-list"`
    DestList []RuleSetOperRulesByZoneOperGroupListRuleListDestList `json:"dest-list"`
    ServiceList []RuleSetOperRulesByZoneOperGroupListRuleListServiceList `json:"service-list"`
    DscpList []RuleSetOperRulesByZoneOperGroupListRuleListDscpList `json:"dscp-list"`
}


type RuleSetOperRulesByZoneOperGroupListRuleListSourceList struct {
    Source string `json:"source"`
}


type RuleSetOperRulesByZoneOperGroupListRuleListDestList struct {
    Dest string `json:"dest"`
}


type RuleSetOperRulesByZoneOperGroupListRuleListServiceList struct {
    Service string `json:"service"`
}


type RuleSetOperRulesByZoneOperGroupListRuleListDscpList struct {
    Dscp string `json:"dscp"`
}


type RuleSetOperTrackAppRuleList struct {
    Oper RuleSetOperTrackAppRuleListOper `json:"oper"`
}


type RuleSetOperTrackAppRuleListOper struct {
    RuleList []RuleSetOperTrackAppRuleListOperRuleList `json:"rule-list"`
}


type RuleSetOperTrackAppRuleListOperRuleList struct {
    Name string `json:"name"`
}

func (p *RuleSetOper) GetId() string{
    return "1"
}

func (p *RuleSetOper) getPath() string{
    
    return "rule-set/"+p.Name+"/oper"
}

func (p *RuleSetOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetOper,error) {
logger.Println("RuleSetOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRuleSetOper
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
