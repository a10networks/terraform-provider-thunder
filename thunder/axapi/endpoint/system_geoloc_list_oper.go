

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeolocListOper struct {
    
    Name string `json:"name"`
    Oper SystemGeolocListOperOper `json:"oper"`

}
type DataSystemGeolocListOper struct {
    DtSystemGeolocListOper SystemGeolocListOper `json:"geoloc-list"`
}


type SystemGeolocListOperOper struct {
    GeolocList []SystemGeolocListOperOperGeolocList `json:"geoloc-list"`
}


type SystemGeolocListOperOperGeolocList struct {
    Type string `json:"type"`
    GeolocName string `json:"geoloc-name"`
    HitCount int `json:"hit-count"`
    Active int `json:"active"`
}

func (p *SystemGeolocListOper) GetId() string{
    return "1"
}

func (p *SystemGeolocListOper) getPath() string{
    
    return "system/geoloc-list/"+p.Name+"/oper"
}

func (p *SystemGeolocListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemGeolocListOper,error) {
logger.Println("SystemGeolocListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemGeolocListOper
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
