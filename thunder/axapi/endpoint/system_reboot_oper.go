

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemRebootOper struct {
    
    Oper SystemRebootOperOper `json:"oper"`

}
type DataSystemRebootOper struct {
    DtSystemRebootOper SystemRebootOper `json:"reboot"`
}


type SystemRebootOperOper struct {
    BootScheduled int `json:"boot-scheduled"`
    BootNow int `json:"boot-now"`
    EpochTime string `json:"epoch-time"`
    Hour string `json:"hour"`
    Min string `json:"min"`
    Uname string `json:"uname"`
    Hostname string `json:"hostname"`
    Reason string `json:"reason"`
}

func (p *SystemRebootOper) GetId() string{
    return "1"
}

func (p *SystemRebootOper) getPath() string{
    return "system/reboot/oper"
}

func (p *SystemRebootOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemRebootOper,error) {
logger.Println("SystemRebootOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemRebootOper
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
