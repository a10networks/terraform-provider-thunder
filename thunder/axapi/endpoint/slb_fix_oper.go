

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbFixOper struct {
    
    Oper SlbFixOperOper `json:"oper"`

}
type DataSlbFixOper struct {
    DtSlbFixOper SlbFixOper `json:"fix"`
}


type SlbFixOperOper struct {
    FixCpuList []SlbFixOperOperFixCpuList `json:"fix-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbFixOperOperFixCpuList struct {
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Svrsel_fail int `json:"svrsel_fail"`
    Noroute int `json:"noroute"`
    Snat_fail int `json:"snat_fail"`
    Client_err int `json:"client_err"`
    Server_err int `json:"server_err"`
    Insert_clientip int `json:"insert_clientip"`
    Default_switching int `json:"default_switching"`
    Sender_switching int `json:"sender_switching"`
    Target_switching int `json:"target_switching"`
    Client_tls_conn int `json:"client_tls_conn"`
    Server_tls_conn int `json:"server_tls_conn"`
}

func (p *SlbFixOper) GetId() string{
    return "1"
}

func (p *SlbFixOper) getPath() string{
    return "slb/fix/oper"
}

func (p *SlbFixOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbFixOper,error) {
logger.Println("SlbFixOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbFixOper
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
