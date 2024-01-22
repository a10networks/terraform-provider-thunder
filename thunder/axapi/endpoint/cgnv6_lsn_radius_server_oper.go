

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRadiusServerOper struct {
    
    Oper Cgnv6LsnRadiusServerOperOper `json:"oper"`

}
type DataCgnv6LsnRadiusServerOper struct {
    DtCgnv6LsnRadiusServerOper Cgnv6LsnRadiusServerOper `json:"server"`
}


type Cgnv6LsnRadiusServerOperOper struct {
    RadiusTableEntriesList []Cgnv6LsnRadiusServerOperOperRadiusTableEntriesList `json:"radius-table-entries-list"`
    TotalEntries int `json:"total-entries"`
    CustomAttrName string `json:"custom-attr-name"`
    CustomAttrValue string `json:"custom-attr-value"`
    StartsWith int `json:"starts-with"`
    CaseInsensitive int `json:"case-insensitive"`
}


type Cgnv6LsnRadiusServerOperOperRadiusTableEntriesList struct {
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

func (p *Cgnv6LsnRadiusServerOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRadiusServerOper) getPath() string{
    return "cgnv6/lsn/radius/server/oper"
}

func (p *Cgnv6LsnRadiusServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRadiusServerOper,error) {
logger.Println("Cgnv6LsnRadiusServerOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRadiusServerOper
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
