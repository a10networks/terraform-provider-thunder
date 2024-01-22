

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugHashTableOper struct {
    
    Oper ScaleoutDebugHashTableOperOper `json:"oper"`

}
type DataScaleoutDebugHashTableOper struct {
    DtScaleoutDebugHashTableOper ScaleoutDebugHashTableOper `json:"hash-table"`
}


type ScaleoutDebugHashTableOperOper struct {
    Ip int `json:"ip"`
    Mac int `json:"mac"`
    Hash_list []ScaleoutDebugHashTableOperOperHash_list `json:"hash_list"`
}


type ScaleoutDebugHashTableOperOperHash_list struct {
    Hash int `json:"hash"`
    Node int `json:"node"`
    So_vnp_id int `json:"so_vnp_id"`
    So_ip string `json:"so_ip"`
    So_mac string `json:"so_mac"`
    Ref_count int `json:"ref_count"`
}

func (p *ScaleoutDebugHashTableOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugHashTableOper) getPath() string{
    return "scaleout/debug/hash-table/oper"
}

func (p *ScaleoutDebugHashTableOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugHashTableOper,error) {
logger.Println("ScaleoutDebugHashTableOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugHashTableOper
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
