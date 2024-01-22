

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemViewShowMonitorOper struct {
    
    Oper SystemViewShowMonitorOperOper `json:"oper"`

}
type DataSystemViewShowMonitorOper struct {
    DtSystemViewShowMonitorOper SystemViewShowMonitorOper `json:"show-monitor"`
}


type SystemViewShowMonitorOperOper struct {
    DiskValue int `json:"disk-value"`
    MemValue int `json:"mem-value"`
    CtrlCpu int `json:"ctrl-cpu"`
    DataCpu int `json:"data-cpu"`
    BuffValue int `json:"buff-value"`
    BuffDrop int `json:"buff-drop"`
    WarnTemp int `json:"warn-temp"`
    Spm0 int `json:"spm0"`
    Spm1 int `json:"spm1"`
    Spm2 int `json:"spm2"`
    Spm3 int `json:"spm3"`
    Spm4 int `json:"spm4"`
    Smp0 int `json:"smp0"`
    Smp1 int `json:"smp1"`
    Smp2 int `json:"smp2"`
    Smp3 int `json:"smp3"`
    Smp4 int `json:"smp4"`
}

func (p *SystemViewShowMonitorOper) GetId() string{
    return "1"
}

func (p *SystemViewShowMonitorOper) getPath() string{
    return "system-view/show-monitor/oper"
}

func (p *SystemViewShowMonitorOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemViewShowMonitorOper,error) {
logger.Println("SystemViewShowMonitorOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemViewShowMonitorOper
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
