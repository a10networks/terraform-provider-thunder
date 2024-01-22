

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysAuditLogOper struct {
    
    Oper SysAuditLogOperOper `json:"oper"`

}
type DataSysAuditLogOper struct {
    DtSysAuditLogOper SysAuditLogOper `json:"sys-audit-log"`
}


type SysAuditLogOperOper struct {
    SystemAuditLog []SysAuditLogOperOperSystemAuditLog `json:"system-audit-log"`
}


type SysAuditLogOperOperSystemAuditLog struct {
    LogAuditData string `json:"log-audit-data"`
    Partitions string `json:"partitions"`
    LogAuditSearch string `json:"log-audit-search"`
}

func (p *SysAuditLogOper) GetId() string{
    return "1"
}

func (p *SysAuditLogOper) getPath() string{
    return "sys-audit-log/oper"
}

func (p *SysAuditLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSysAuditLogOper,error) {
logger.Println("SysAuditLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSysAuditLogOper
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
