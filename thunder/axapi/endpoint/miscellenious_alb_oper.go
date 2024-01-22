

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type MiscelleniousAlbOper struct {
    
    Oper MiscelleniousAlbOperOper `json:"oper"`

}
type DataMiscelleniousAlbOper struct {
    DtMiscelleniousAlbOper MiscelleniousAlbOper `json:"miscellenious-alb"`
}


type MiscelleniousAlbOperOper struct {
    Uptime int `json:"uptime"`
    RebootNum int `json:"reboot-num"`
    CrashCount int `json:"crash-count"`
}

func (p *MiscelleniousAlbOper) GetId() string{
    return "1"
}

func (p *MiscelleniousAlbOper) getPath() string{
    return "miscellenious-alb/oper"
}

func (p *MiscelleniousAlbOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataMiscelleniousAlbOper,error) {
logger.Println("MiscelleniousAlbOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataMiscelleniousAlbOper
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
