

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugRoleOper struct {
    
    Oper ScaleoutDebugRoleOperOper `json:"oper"`

}
type DataScaleoutDebugRoleOper struct {
    DtScaleoutDebugRoleOper ScaleoutDebugRoleOper `json:"role"`
}


type ScaleoutDebugRoleOperOper struct {
    Scaleout_role_action string `json:"scaleout_role_action"`
}

func (p *ScaleoutDebugRoleOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugRoleOper) getPath() string{
    return "scaleout/debug/role/oper"
}

func (p *ScaleoutDebugRoleOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugRoleOper,error) {
logger.Println("ScaleoutDebugRoleOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugRoleOper
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
