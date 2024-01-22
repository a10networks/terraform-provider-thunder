

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemRadiusServerOper struct {
    
    Oper SystemRadiusServerOperOper `json:"oper"`

}
type DataSystemRadiusServerOper struct {
    DtSystemRadiusServerOper SystemRadiusServerOper `json:"server"`
}


type SystemRadiusServerOperOper struct {
    RadiusTableEntriesList []SystemRadiusServerOperOperRadiusTableEntriesList `json:"radius-table-entries-list"`
    TotalEntries int `json:"total-entries"`
    CustomAttrName string `json:"custom-attr-name"`
    CustomAttrValue string `json:"custom-attr-value"`
    StartsWith int `json:"starts-with"`
    CaseInsensitive int `json:"case-insensitive"`
}


type SystemRadiusServerOperOperRadiusTableEntriesList struct {
    InsideIp string `json:"inside-ip"`
    InsideIpv6 string `json:"inside-ipv6"`
    PrefixLen int `json:"prefix-len"`
    Msisdn string `json:"msisdn"`
    Imei string `json:"imei"`
    Imsi string `json:"imsi"`
    Custom1AttrValue string `json:"custom1-attr-value"`
    Custom2AttrValue string `json:"custom2-attr-value"`
    Custom3AttrValue string `json:"custom3-attr-value"`
    Custom4AttrValue string `json:"custom4-attr-value"`
    Custom5AttrValue string `json:"custom5-attr-value"`
    Custom6AttrValue string `json:"custom6-attr-value"`
    IsObsolete int `json:"is-obsolete"`
}

func (p *SystemRadiusServerOper) GetId() string{
    return "1"
}

func (p *SystemRadiusServerOper) getPath() string{
    return "system/radius/server/oper"
}

func (p *SystemRadiusServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemRadiusServerOper,error) {
logger.Println("SystemRadiusServerOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemRadiusServerOper
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
