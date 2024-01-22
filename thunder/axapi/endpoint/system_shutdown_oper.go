

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemShutdownOper struct {
    
    Oper SystemShutdownOperOper `json:"oper"`

}
type DataSystemShutdownOper struct {
    DtSystemShutdownOper SystemShutdownOper `json:"shutdown"`
}


type SystemShutdownOperOper struct {
    BootScheduled int `json:"boot-scheduled"`
    BootNow int `json:"boot-now"`
    EpochTime string `json:"epoch-time"`
    Hour string `json:"hour"`
    Min string `json:"min"`
    Uname string `json:"uname"`
    Hostname string `json:"hostname"`
    Reason string `json:"reason"`
}

func (p *SystemShutdownOper) GetId() string{
    return "1"
}

func (p *SystemShutdownOper) getPath() string{
    return "system/shutdown/oper"
}

func (p *SystemShutdownOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemShutdownOper,error) {
logger.Println("SystemShutdownOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemShutdownOper
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
