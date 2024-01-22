

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DomainGroupOperOper struct {
    
    Oper DomainGroupOperOperOper `json:"oper"`

}
type DataDomainGroupOperOper struct {
    DtDomainGroupOperOper DomainGroupOperOper `json:"domain-group-oper"`
}


type DomainGroupOperOperOper struct {
    DomainGroups []DomainGroupOperOperOperDomainGroups `json:"domain-groups"`
    DomainGroupEntries []DomainGroupOperOperOperDomainGroupEntries `json:"domain-group-entries"`
    DisplayedCount int `json:"displayed-count"`
    DomainGroupNameFilter string `json:"domain-group-name-filter"`
    DomainNameFilter string `json:"domain-name-filter"`
}


type DomainGroupOperOperOperDomainGroups struct {
    DomainGroupName string `json:"domain-group-name"`
    TotalEntryNum int `json:"total-entry-num"`
    BindingNum int `json:"binding-num"`
}


type DomainGroupOperOperOperDomainGroupEntries struct {
    DomainMatchType string `json:"domain-match-type"`
    DomainName string `json:"domain-name"`
    DomainListName string `json:"domain-list-name"`
    Action string `json:"action"`
    HitCount int `json:"hit-count"`
}

func (p *DomainGroupOperOper) GetId() string{
    return "1"
}

func (p *DomainGroupOperOper) getPath() string{
    return "domain-group-oper/oper"
}

func (p *DomainGroupOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDomainGroupOperOper,error) {
logger.Println("DomainGroupOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDomainGroupOperOper
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
