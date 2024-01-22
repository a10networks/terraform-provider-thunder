

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemViewChassisDetailOper struct {
    
    Oper SystemViewChassisDetailOperOper `json:"oper"`

}
type DataSystemViewChassisDetailOper struct {
    DtSystemViewChassisDetailOper SystemViewChassisDetailOper `json:"chassis-detail"`
}


type SystemViewChassisDetailOperOper struct {
    PlatformName string `json:"platform-name"`
    SingleBrdModeForced string `json:"single-brd-mode-forced"`
    SingleBrdModeFallback string `json:"single-brd-mode-fallback"`
    Is_now_single_board string `json:"is_now_single_board"`
}

func (p *SystemViewChassisDetailOper) GetId() string{
    return "1"
}

func (p *SystemViewChassisDetailOper) getPath() string{
    return "system-view/chassis-detail/oper"
}

func (p *SystemViewChassisDetailOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemViewChassisDetailOper,error) {
logger.Println("SystemViewChassisDetailOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemViewChassisDetailOper
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
