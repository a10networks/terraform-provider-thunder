

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTcpSynPerSecOper struct {
    
    Oper SystemTcpSynPerSecOperOper `json:"oper"`

}
type DataSystemTcpSynPerSecOper struct {
    DtSystemTcpSynPerSecOper SystemTcpSynPerSecOper `json:"tcp-syn-per-sec"`
}


type SystemTcpSynPerSecOperOper struct {
    TcpSynPerSec int `json:"tcp-syn-per-sec"`
}

func (p *SystemTcpSynPerSecOper) GetId() string{
    return "1"
}

func (p *SystemTcpSynPerSecOper) getPath() string{
    return "system/tcp-syn-per-sec/oper"
}

func (p *SystemTcpSynPerSecOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTcpSynPerSecOper,error) {
logger.Println("SystemTcpSynPerSecOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTcpSynPerSecOper
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
