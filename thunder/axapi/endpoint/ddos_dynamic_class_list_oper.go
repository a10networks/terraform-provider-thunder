

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDynamicClassListOper struct {
    
    Oper DdosDynamicClassListOperOper `json:"oper"`

}
type DataDdosDynamicClassListOper struct {
    DtDdosDynamicClassListOper DdosDynamicClassListOper `json:"dynamic-class-list"`
}


type DdosDynamicClassListOperOper struct {
    Type string `json:"type"`
    FileOrString string `json:"file-or-string"`
    UserTag string `json:"user-tag"`
    Ipv4TotalSingleIp int `json:"ipv4-total-single-ip"`
    Ipv4TotalSubnet int `json:"ipv4-total-subnet"`
    Ipv6TotalSingleIp int `json:"ipv6-total-single-ip"`
    Ipv6TotalSubnet int `json:"ipv6-total-subnet"`
    DnsTotalEntries int `json:"dns-total-entries"`
    StringTotalEntries int `json:"string-total-entries"`
    AcTotalEntries int `json:"ac-total-entries"`
    GeoLocationTotalEntries int `json:"geo-location-total-entries"`
    Ipv4Entries []DdosDynamicClassListOperOperIpv4Entries `json:"ipv4-entries"`
    Ipv6Entries []DdosDynamicClassListOperOperIpv6Entries `json:"ipv6-entries"`
    DnsEntries []DdosDynamicClassListOperOperDnsEntries `json:"dns-entries"`
    StringEntries []DdosDynamicClassListOperOperStringEntries `json:"string-entries"`
    AcEntries []DdosDynamicClassListOperOperAcEntries `json:"ac-entries"`
    GeoLocationEntries []DdosDynamicClassListOperOperGeoLocationEntries `json:"geo-location-entries"`
}


type DdosDynamicClassListOperOperIpv4Entries struct {
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv4Lid int `json:"ipv4-lid"`
    Ipv4Glid int `json:"ipv4-glid"`
    Ipv4LsnLid int `json:"ipv4-lsn-lid"`
    Ipv4LsnRadiusProfile int `json:"ipv4-lsn-radius-profile"`
    Ipv4HitCount int `json:"ipv4-hit-count"`
    Ipv4Age int `json:"ipv4-age"`
}


type DdosDynamicClassListOperOperIpv6Entries struct {
    Ipv6addr string `json:"ipv6addr"`
    Ipv6Lid int `json:"ipv6-lid"`
    Ipv6Glid int `json:"ipv6-glid"`
    Ipv6LsnLid int `json:"ipv6-lsn-lid"`
    Ipv6LsnRadiusProfile int `json:"ipv6-lsn-radius-profile"`
    Ipv6HitCount int `json:"ipv6-hit-count"`
    Ipv6Age int `json:"ipv6-age"`
}


type DdosDynamicClassListOperOperDnsEntries struct {
    DnsMatchType string `json:"dns-match-type"`
    DnsMatchString string `json:"dns-match-string"`
    DnsLid int `json:"dns-lid"`
    DnsGlid int `json:"dns-glid"`
    DnsHitCount int `json:"dns-hit-count"`
}


type DdosDynamicClassListOperOperStringEntries struct {
    StringKey string `json:"string-key"`
    StringValue string `json:"string-value"`
    StringLid int `json:"string-lid"`
    StringGlid int `json:"string-glid"`
    StringHitCount int `json:"string-hit-count"`
}


type DdosDynamicClassListOperOperAcEntries struct {
    AcMatchType string `json:"ac-match-type"`
    AcMatchString string `json:"ac-match-string"`
    AcMatchValue string `json:"ac-match-value"`
    AcHitCount int `json:"ac-hit-count"`
}


type DdosDynamicClassListOperOperGeoLocationEntries struct {
    GeoLocation string `json:"geo-location"`
}

func (p *DdosDynamicClassListOper) GetId() string{
    return "1"
}

func (p *DdosDynamicClassListOper) getPath() string{
    return "ddos/dynamic-class-list/oper"
}

func (p *DdosDynamicClassListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDynamicClassListOper,error) {
logger.Println("DdosDynamicClassListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDynamicClassListOper
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
