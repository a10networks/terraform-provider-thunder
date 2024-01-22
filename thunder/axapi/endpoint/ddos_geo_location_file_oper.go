

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosGeoLocationFileOper struct {
    
    Oper DdosGeoLocationFileOperOper `json:"oper"`

}
type DataDdosGeoLocationFileOper struct {
    DtDdosGeoLocationFileOper DdosGeoLocationFileOper `json:"file"`
}


type DdosGeoLocationFileOperOper struct {
    FileList []DdosGeoLocationFileOperOperFileList `json:"file-list"`
}


type DdosGeoLocationFileOperOperFileList struct {
    Filename string `json:"filename"`
    Type string `json:"type"`
    Lines int `json:"lines"`
    Success int `json:"success"`
    ErrorWarning int `json:"error-warning"`
    ErrorList []DdosGeoLocationFileOperOperFileListErrorList `json:"error-list"`
}


type DdosGeoLocationFileOperOperFileListErrorList struct {
    ErrorLine int `json:"error-line"`
    ErrorInformation string `json:"error-information"`
}

func (p *DdosGeoLocationFileOper) GetId() string{
    return "1"
}

func (p *DdosGeoLocationFileOper) getPath() string{
    return "ddos/geo-location/file/oper"
}

func (p *DdosGeoLocationFileOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosGeoLocationFileOper,error) {
logger.Println("DdosGeoLocationFileOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosGeoLocationFileOper
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
