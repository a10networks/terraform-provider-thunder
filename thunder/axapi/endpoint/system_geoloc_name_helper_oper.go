

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeolocNameHelperOper struct {
    
    Oper SystemGeolocNameHelperOperOper `json:"oper"`

}
type DataSystemGeolocNameHelperOper struct {
    DtSystemGeolocNameHelperOper SystemGeolocNameHelperOper `json:"geoloc-name-helper"`
}


type SystemGeolocNameHelperOperOper struct {
    GeolocCandidateList []SystemGeolocNameHelperOperOperGeolocCandidateList `json:"geoloc-candidate-list"`
    Geoloc string `json:"geoloc"`
}


type SystemGeolocNameHelperOperOperGeolocCandidateList struct {
    GeolocName string `json:"geoloc-name"`
    HasSubregion int `json:"has-subregion"`
}

func (p *SystemGeolocNameHelperOper) GetId() string{
    return "1"
}

func (p *SystemGeolocNameHelperOper) getPath() string{
    return "system/geoloc-name-helper/oper"
}

func (p *SystemGeolocNameHelperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemGeolocNameHelperOper,error) {
logger.Println("SystemGeolocNameHelperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemGeolocNameHelperOper
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
