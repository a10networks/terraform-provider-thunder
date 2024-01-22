

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugNatMapOper struct {
    
    Oper ScaleoutDebugNatMapOperOper `json:"oper"`

}
type DataScaleoutDebugNatMapOper struct {
    DtScaleoutDebugNatMapOper ScaleoutDebugNatMapOper `json:"nat-map"`
}


type ScaleoutDebugNatMapOperOper struct {
    DeviceId int `json:"device-id"`
    Vnp_id_list []ScaleoutDebugNatMapOperOperVnp_id_list `json:"vnp_id_list"`
}


type ScaleoutDebugNatMapOperOperVnp_id_list struct {
    Nat_map_list []ScaleoutDebugNatMapOperOperVnp_id_listNat_map_list `json:"nat_map_list"`
}


type ScaleoutDebugNatMapOperOperVnp_id_listNat_map_list struct {
    Vnp_id int `json:"vnp_id"`
    Ip string `json:"ip"`
    Owner int `json:"owner"`
    Active int `json:"active"`
}

func (p *ScaleoutDebugNatMapOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugNatMapOper) getPath() string{
    return "scaleout/debug/nat-map/oper"
}

func (p *ScaleoutDebugNatMapOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugNatMapOper,error) {
logger.Println("ScaleoutDebugNatMapOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugNatMapOper
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
