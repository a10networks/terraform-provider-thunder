

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FailSafeConfigOper struct {
    
    Oper FailSafeConfigOperOper `json:"oper"`

}
type DataFailSafeConfigOper struct {
    DtFailSafeConfigOper FailSafeConfigOper `json:"config"`
}


type FailSafeConfigOperOper struct {
    Sw_error_mon string `json:"sw_error_mon"`
    Hw_error_mon string `json:"hw_error_mon"`
    Sw_recovery_timeout string `json:"sw_recovery_timeout"`
    Hw_recovery_timeout string `json:"hw_recovery_timeout"`
    Dataplane_recovery_timeout string `json:"dataplane_recovery_timeout"`
    Fpga_mon_enable string `json:"fpga_mon_enable"`
    Fpga_mon_forced_reboot string `json:"fpga_mon_forced_reboot"`
    Fpga_mon_interval string `json:"fpga_mon_interval"`
    Fpga_mon_threshold string `json:"fpga_mon_threshold"`
    Mem_mon string `json:"mem_mon"`
}

func (p *FailSafeConfigOper) GetId() string{
    return "1"
}

func (p *FailSafeConfigOper) getPath() string{
    return "fail-safe/config/oper"
}

func (p *FailSafeConfigOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFailSafeConfigOper,error) {
logger.Println("FailSafeConfigOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFailSafeConfigOper
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
