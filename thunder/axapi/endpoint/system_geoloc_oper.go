

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeolocOper struct {
    
    Oper SystemGeolocOperOper `json:"oper"`

}
type DataSystemGeolocOper struct {
    DtSystemGeolocOper SystemGeolocOper `json:"geoloc"`
}


type SystemGeolocOperOper struct {
    GeolocList []SystemGeolocOperOperGeolocList `json:"geoloc-list"`
    TotalGeolocs int `json:"total-geolocs"`
    GeoName string `json:"geo-name"`
    Filter1 string `json:"filter1"`
    Filter2 string `json:"filter2"`
    Filter3 string `json:"filter3"`
    Filter4 string `json:"filter4"`
    PolName string `json:"pol-name"`
    Iprangestrt string `json:"iprangestrt"`
    Iprangeend string `json:"iprangeend"`
    Ipv6rangestrt string `json:"ipv6rangestrt"`
    Depth int `json:"depth"`
}


type SystemGeolocOperOperGeolocList struct {
    Name string `json:"name"`
    From string `json:"from"`
    Tomask string `json:"tomask"`
    Last string `json:"last"`
    Hits int `json:"hits"`
    Subcnt int `json:"subcnt"`
    Type string `json:"type"`
}

func (p *SystemGeolocOper) GetId() string{
    return "1"
}

func (p *SystemGeolocOper) getPath() string{
    return "system/geoloc/oper"
}

func (p *SystemGeolocOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemGeolocOper,error) {
logger.Println("SystemGeolocOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemGeolocOper
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
