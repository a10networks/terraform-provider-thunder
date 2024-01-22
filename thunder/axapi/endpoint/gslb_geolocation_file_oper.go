

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbGeolocationFileOper struct {
    
    Oper GslbGeolocationFileOperOper `json:"oper"`

}
type DataGslbGeolocationFileOper struct {
    DtGslbGeolocationFileOper GslbGeolocationFileOper `json:"geolocation-file"`
}


type GslbGeolocationFileOperOper struct {
    Geofiles []GslbGeolocationFileOperOperGeofiles `json:"geofiles"`
}


type GslbGeolocationFileOperOperGeofiles struct {
    Filename string `json:"filename"`
    Type string `json:"type"`
    Template string `json:"template"`
    PercentageLoaded int `json:"percentage-loaded"`
    Lines int `json:"lines"`
    Success int `json:"success"`
    ErrorWarning int `json:"error-warning"`
    Comment int `json:"comment"`
}

func (p *GslbGeolocationFileOper) GetId() string{
    return "1"
}

func (p *GslbGeolocationFileOper) getPath() string{
    return "gslb/geolocation-file/oper"
}

func (p *GslbGeolocationFileOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbGeolocationFileOper,error) {
logger.Println("GslbGeolocationFileOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbGeolocationFileOper
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
