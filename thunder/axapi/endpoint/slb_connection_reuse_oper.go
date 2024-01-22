

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbConnectionReuseOper struct {
    
    Oper SlbConnectionReuseOperOper `json:"oper"`

}
type DataSlbConnectionReuseOper struct {
    DtSlbConnectionReuseOper SlbConnectionReuseOper `json:"connection-reuse"`
}


type SlbConnectionReuseOperOper struct {
    ConnectionReuseCpuList []SlbConnectionReuseOperOperConnectionReuseCpuList `json:"connection-reuse-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbConnectionReuseOperOperConnectionReuseCpuList struct {
    Current_open int `json:"current_open"`
    Current_active int `json:"current_active"`
    Nbind int `json:"nbind"`
    Nunbind int `json:"nunbind"`
    Nestab int `json:"nestab"`
    Ntermi int `json:"ntermi"`
    Ntermi_err int `json:"ntermi_err"`
    Delay_unbind int `json:"delay_unbind"`
    Long_resp int `json:"long_resp"`
    Miss_resp int `json:"miss_resp"`
    Unbound_data_rcv int `json:"unbound_data_rcv"`
    Pause_conn int `json:"pause_conn"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Resume_conn int `json:"resume_conn"`
    Not_remove_from_rport int `json:"not_remove_from_rport"`
}

func (p *SlbConnectionReuseOper) GetId() string{
    return "1"
}

func (p *SlbConnectionReuseOper) getPath() string{
    return "slb/connection-reuse/oper"
}

func (p *SlbConnectionReuseOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbConnectionReuseOper,error) {
logger.Println("SlbConnectionReuseOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbConnectionReuseOper
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
