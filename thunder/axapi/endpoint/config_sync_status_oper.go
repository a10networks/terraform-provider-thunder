

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ConfigSyncStatusOper struct {
    
    Oper ConfigSyncStatusOperOper `json:"oper"`

}
type DataConfigSyncStatusOper struct {
    DtConfigSyncStatusOper ConfigSyncStatusOper `json:"config-sync-status"`
}


type ConfigSyncStatusOperOper struct {
    ConfigSyncList []ConfigSyncStatusOperOperConfigSyncList `json:"config-sync-list"`
}


type ConfigSyncStatusOperOperConfigSyncList struct {
    PartitionName string `json:"partition-name"`
    RunSyncStatus string `json:"run-sync-status"`
    StartupSyncStatus string `json:"startup-sync-status"`
}

func (p *ConfigSyncStatusOper) GetId() string{
    return "1"
}

func (p *ConfigSyncStatusOper) getPath() string{
    return "config-sync-status/oper"
}

func (p *ConfigSyncStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataConfigSyncStatusOper,error) {
logger.Println("ConfigSyncStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataConfigSyncStatusOper
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
