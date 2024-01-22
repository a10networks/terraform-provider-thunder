

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type BackupStatusOper struct {
    
    Oper BackupStatusOperOper `json:"oper"`

}
type DataBackupStatusOper struct {
    DtBackupStatusOper BackupStatusOper `json:"status"`
}


type BackupStatusOperOper struct {
    Status int `json:"status"`
    Message string `json:"message"`
    Size int `json:"size"`
    Progress int `json:"progress"`
}

func (p *BackupStatusOper) GetId() string{
    return "1"
}

func (p *BackupStatusOper) getPath() string{
    return "backup/status/oper"
}

func (p *BackupStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataBackupStatusOper,error) {
logger.Println("BackupStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataBackupStatusOper
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
