

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugNatListOper struct {
    
    Oper ScaleoutDebugNatListOperOper `json:"oper"`

}
type DataScaleoutDebugNatListOper struct {
    DtScaleoutDebugNatListOper ScaleoutDebugNatListOper `json:"nat-list"`
}


type ScaleoutDebugNatListOperOper struct {
    Vnp_id_list []ScaleoutDebugNatListOperOperVnp_id_list `json:"vnp_id_list"`
}


type ScaleoutDebugNatListOperOperVnp_id_list struct {
    Nat_list []ScaleoutDebugNatListOperOperVnp_id_listNat_list `json:"nat_list"`
}


type ScaleoutDebugNatListOperOperVnp_id_listNat_list struct {
    Ip string `json:"ip"`
    Device int `json:"device"`
    Active int `json:"active"`
}

func (p *ScaleoutDebugNatListOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugNatListOper) getPath() string{
    return "scaleout/debug/nat-list/oper"
}

func (p *ScaleoutDebugNatListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugNatListOper,error) {
logger.Println("ScaleoutDebugNatListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugNatListOper
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
