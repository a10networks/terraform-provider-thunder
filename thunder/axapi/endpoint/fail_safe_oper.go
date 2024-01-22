

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FailSafeOper struct {
    
    Config FailSafeOperConfig `json:"config"`
    Oper FailSafeOperOper `json:"oper"`

}
type DataFailSafeOper struct {
    DtFailSafeOper FailSafeOper `json:"fail-safe"`
}


type FailSafeOperConfig struct {
    Oper FailSafeOperConfigOper `json:"oper"`
}


type FailSafeOperConfigOper struct {
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


type FailSafeOperOper struct {
    Free_session_memory int `json:"free_session_memory"`
    Total_session_memory int `json:"total_session_memory"`
    Sess_mem_recovery_threshold int `json:"sess_mem_recovery_threshold"`
    Total_fpga_buffers int `json:"total_fpga_buffers"`
    Avail_fpga_buff_domain1 int `json:"avail_fpga_buff_domain1"`
    Avail_fpga_buff_domain2 int `json:"avail_fpga_buff_domain2"`
    Total_free_fpga_buff int `json:"total_free_fpga_buff"`
    Free_fpga_buffers int `json:"free_fpga_buffers"`
    Fpga_buff_recovery_threshold int `json:"fpga_buff_recovery_threshold"`
    Total_system_memory int `json:"total_system_memory"`
    Fpga_stats_num_cntrs int `json:"fpga_stats_num_cntrs"`
    Fpga_stats_iochan []FailSafeOperOperFpga_stats_iochan `json:"fpga_stats_iochan"`
}


type FailSafeOperOperFpga_stats_iochan struct {
    Fpga_stats_iochan_id int `json:"fpga_stats_iochan_id"`
    Fpga_stats_iochan_tx int `json:"fpga_stats_iochan_tx"`
    Fpga_stats_iochan_rx int `json:"fpga_stats_iochan_rx"`
}

func (p *FailSafeOper) GetId() string{
    return "1"
}

func (p *FailSafeOper) getPath() string{
    return "fail-safe/oper"
}

func (p *FailSafeOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFailSafeOper,error) {
logger.Println("FailSafeOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFailSafeOper
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
