

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemViewShowBackupOper struct {
    
    Oper SystemViewShowBackupOperOper `json:"oper"`

}
type DataSystemViewShowBackupOper struct {
    DtSystemViewShowBackupOper SystemViewShowBackupOper `json:"show-backup"`
}


type SystemViewShowBackupOperOper struct {
    BackupShow1 string `json:"backup-show-1"`
    BackupShow2 string `json:"backup-show-2"`
    BackupShow3 string `json:"backup-show-3"`
}

func (p *SystemViewShowBackupOper) GetId() string{
    return "1"
}

func (p *SystemViewShowBackupOper) getPath() string{
    return "system-view/show-backup/oper"
}

func (p *SystemViewShowBackupOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemViewShowBackupOper,error) {
logger.Println("SystemViewShowBackupOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemViewShowBackupOper
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
