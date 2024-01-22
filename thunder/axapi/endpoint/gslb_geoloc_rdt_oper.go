

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbGeolocRdtOper struct {
    
    Oper GslbGeolocRdtOperOper `json:"oper"`

}
type DataGslbGeolocRdtOper struct {
    DtGslbGeolocRdtOper GslbGeolocRdtOper `json:"geoloc-rdt"`
}


type GslbGeolocRdtOperOper struct {
    GeolocRdtList []GslbGeolocRdtOperOperGeolocRdtList `json:"geoloc-rdt-list"`
    TotalRdt int `json:"total-rdt"`
    GeoName string `json:"geo-name"`
    SiteName string `json:"site-name"`
    ActiveStatus string `json:"active-status"`
}


type GslbGeolocRdtOperOperGeolocRdtList struct {
    GlName string `json:"gl-name"`
    SiteName string `json:"site-name"`
    Type string `json:"type"`
    Rdt int `json:"rdt"`
    Age int `json:"age"`
}

func (p *GslbGeolocRdtOper) GetId() string{
    return "1"
}

func (p *GslbGeolocRdtOper) getPath() string{
    return "gslb/geoloc-rdt/oper"
}

func (p *GslbGeolocRdtOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbGeolocRdtOper,error) {
logger.Println("GslbGeolocRdtOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbGeolocRdtOper
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
