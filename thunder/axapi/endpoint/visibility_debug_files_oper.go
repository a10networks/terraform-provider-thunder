

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityDebugFilesOper struct {
    
    Oper VisibilityDebugFilesOperOper `json:"oper"`

}
type DataVisibilityDebugFilesOper struct {
    DtVisibilityDebugFilesOper VisibilityDebugFilesOper `json:"debug-files"`
}


type VisibilityDebugFilesOperOper struct {
    DebugFileNameList []VisibilityDebugFilesOperOperDebugFileNameList `json:"debug-file-name-list"`
}


type VisibilityDebugFilesOperOperDebugFileNameList struct {
    EntityKey string `json:"entity-key"`
    Uuid string `json:"uuid"`
    FlatOid int `json:"flat-oid"`
    Ipv4Addr string `json:"ipv4-addr"`
    Ipv6Addr string `json:"ipv6-addr"`
    Protocol string `json:"protocol"`
    Port int `json:"port"`
    DebugFileName string `json:"debug-file-name"`
}

func (p *VisibilityDebugFilesOper) GetId() string{
    return "1"
}

func (p *VisibilityDebugFilesOper) getPath() string{
    return "visibility/debug-files/oper"
}

func (p *VisibilityDebugFilesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityDebugFilesOper,error) {
logger.Println("VisibilityDebugFilesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityDebugFilesOper
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
