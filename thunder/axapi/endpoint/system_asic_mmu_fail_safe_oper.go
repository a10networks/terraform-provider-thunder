

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemAsicMmuFailSafeOper struct {
    
    Oper SystemAsicMmuFailSafeOperOper `json:"oper"`

}
type DataSystemAsicMmuFailSafeOper struct {
    DtSystemAsicMmuFailSafeOper SystemAsicMmuFailSafeOper `json:"asic-mmu-fail-safe"`
}


type SystemAsicMmuFailSafeOperOper struct {
    Enable_val string `json:"enable_val"`
    Threshold_val int `json:"threshold_val"`
    Period_val int `json:"period_val"`
    Reboot_val string `json:"reboot_val"`
    Parity_unrecov int `json:"parity_unrecov"`
    Parity_error int `json:"parity_error"`
    Parity_ecc int `json:"parity_ecc"`
    Parity_ser int `json:"parity_ser"`
    Parity_start_err int `json:"parity_start_err"`
    Mmu_ser int `json:"mmu_ser"`
    Parity_bst int `json:"parity_bst"`
    Mmu_total_error int `json:"mmu_total_error"`
}

func (p *SystemAsicMmuFailSafeOper) GetId() string{
    return "1"
}

func (p *SystemAsicMmuFailSafeOper) getPath() string{
    return "system/asic-mmu-fail-safe/oper"
}

func (p *SystemAsicMmuFailSafeOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemAsicMmuFailSafeOper,error) {
logger.Println("SystemAsicMmuFailSafeOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemAsicMmuFailSafeOper
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
