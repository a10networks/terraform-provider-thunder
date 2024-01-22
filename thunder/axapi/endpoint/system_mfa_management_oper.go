

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemMfaManagementOper struct {
    
    Oper SystemMfaManagementOperOper `json:"oper"`

}
type DataSystemMfaManagementOper struct {
    DtSystemMfaManagementOper SystemMfaManagementOper `json:"mfa-management"`
}


type SystemMfaManagementOperOper struct {
    Enable int `json:"enable"`
}

func (p *SystemMfaManagementOper) GetId() string{
    return "1"
}

func (p *SystemMfaManagementOper) getPath() string{
    return "system/mfa-management/oper"
}

func (p *SystemMfaManagementOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemMfaManagementOper,error) {
logger.Println("SystemMfaManagementOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemMfaManagementOper
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
