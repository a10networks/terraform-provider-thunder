

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeolocationFileErrorInfoOper struct {
    
    Oper SystemGeolocationFileErrorInfoOperOper `json:"oper"`

}
type DataSystemGeolocationFileErrorInfoOper struct {
    DtSystemGeolocationFileErrorInfoOper SystemGeolocationFileErrorInfoOper `json:"error-info"`
}


type SystemGeolocationFileErrorInfoOperOper struct {
    ErrorList []SystemGeolocationFileErrorInfoOperOperErrorList `json:"error-list"`
    Filename string `json:"filename"`
}


type SystemGeolocationFileErrorInfoOperOperErrorList struct {
    Line int `json:"line"`
    Offset int `json:"offset"`
    Error string `json:"error"`
}

func (p *SystemGeolocationFileErrorInfoOper) GetId() string{
    return "1"
}

func (p *SystemGeolocationFileErrorInfoOper) getPath() string{
    return "system/geolocation-file/error-info/oper"
}

func (p *SystemGeolocationFileErrorInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemGeolocationFileErrorInfoOper,error) {
logger.Println("SystemGeolocationFileErrorInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemGeolocationFileErrorInfoOper
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
