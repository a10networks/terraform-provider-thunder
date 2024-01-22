

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGuestFileOper struct {
    
    Oper SystemGuestFileOperOper `json:"oper"`

}
type DataSystemGuestFileOper struct {
    DtSystemGuestFileOper SystemGuestFileOper `json:"guest-file"`
}


type SystemGuestFileOperOper struct {
    FileList []SystemGuestFileOperOperFileList `json:"file-list"`
}


type SystemGuestFileOperOperFileList struct {
    FileName string `json:"File-Name"`
    Size int `json:"Size"`
    UpdateTime string `json:"update-time"`
}

func (p *SystemGuestFileOper) GetId() string{
    return "1"
}

func (p *SystemGuestFileOper) getPath() string{
    return "system/guest-file/oper"
}

func (p *SystemGuestFileOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemGuestFileOper,error) {
logger.Println("SystemGuestFileOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemGuestFileOper
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
