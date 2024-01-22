

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemUpgradeStatusOper struct {
    
    Oper SystemUpgradeStatusOperOper `json:"oper"`

}
type DataSystemUpgradeStatusOper struct {
    DtSystemUpgradeStatusOper SystemUpgradeStatusOper `json:"upgrade-status"`
}


type SystemUpgradeStatusOperOper struct {
    Status int `json:"status"`
    Message string `json:"message"`
    RollbackPri string `json:"rollback-pri"`
    RollbackSec string `json:"rollback-sec"`
}

func (p *SystemUpgradeStatusOper) GetId() string{
    return "1"
}

func (p *SystemUpgradeStatusOper) getPath() string{
    return "system/upgrade-status/oper"
}

func (p *SystemUpgradeStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemUpgradeStatusOper,error) {
logger.Println("SystemUpgradeStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemUpgradeStatusOper
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
