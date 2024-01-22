

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosGeoLocationDbOper struct {
    
    Oper DdosGeoLocationDbOperOper `json:"oper"`

}
type DataDdosGeoLocationDbOper struct {
    DtDdosGeoLocationDbOper DdosGeoLocationDbOper `json:"db"`
}


type DdosGeoLocationDbOperOper struct {
    GeoLocationList []DdosGeoLocationDbOperOperGeoLocationList `json:"geo-location-list"`
    TotalGeoLocation int `json:"total-geo-location"`
    Filter string `json:"filter"`
    Ip string `json:"ip"`
    Ipv6 string `json:"ipv6"`
}


type DdosGeoLocationDbOperOperGeoLocationList struct {
    Name string `json:"name"`
    From string `json:"from"`
    Tomask string `json:"tomask"`
    Subcnt int `json:"subcnt"`
    Type string `json:"type"`
}

func (p *DdosGeoLocationDbOper) GetId() string{
    return "1"
}

func (p *DdosGeoLocationDbOper) getPath() string{
    return "ddos/geo-location/db/oper"
}

func (p *DdosGeoLocationDbOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosGeoLocationDbOper,error) {
logger.Println("DdosGeoLocationDbOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosGeoLocationDbOper
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
