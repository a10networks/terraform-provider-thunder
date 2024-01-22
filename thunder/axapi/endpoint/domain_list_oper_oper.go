

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DomainListOperOper struct {
    
    Oper DomainListOperOperOper `json:"oper"`

}
type DataDomainListOperOper struct {
    DtDomainListOperOper DomainListOperOper `json:"domain-list-oper"`
}


type DomainListOperOperOper struct {
    DomainList []DomainListOperOperOperDomainList `json:"domain-list"`
    DomainEntries []DomainListOperOperOperDomainEntries `json:"domain-entries"`
    DisplayedCount int `json:"displayed-count"`
    MatchType string `json:"match-type"`
    DomainListNameFilter string `json:"domain-list-name-filter"`
    DomainListType string `json:"domain-list-type"`
    BindingInfo int `json:"binding-info"`
}


type DomainListOperOperOperDomainList struct {
    DomainListName string `json:"domain-list-name"`
    ConfigType string `json:"config-type"`
    TotalEntryNum int `json:"total-entry-num"`
    BindingNum int `json:"binding-num"`
    BindingInfoList []DomainListOperOperOperDomainListBindingInfoList `json:"binding-info-list"`
}


type DomainListOperOperOperDomainListBindingInfoList struct {
    DomainGroupName string `json:"domain-group-name"`
}


type DomainListOperOperOperDomainEntries struct {
    DomainMatchType string `json:"domain-match-type"`
    DomainName string `json:"domain-name"`
    ZoneTransferIp string `json:"zone-transfer-ip"`
    ZoneTransferPort string `json:"zone-transfer-port"`
    ZoneTransferInterval string `json:"zone-transfer-interval"`
}

func (p *DomainListOperOper) GetId() string{
    return "1"
}

func (p *DomainListOperOper) getPath() string{
    return "domain-list-oper/oper"
}

func (p *DomainListOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDomainListOperOper,error) {
logger.Println("DomainListOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDomainListOperOper
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
