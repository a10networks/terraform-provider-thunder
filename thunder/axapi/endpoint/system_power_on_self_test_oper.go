

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemPowerOnSelfTestOper struct {
    
    Oper SystemPowerOnSelfTestOperOper `json:"oper"`

}
type DataSystemPowerOnSelfTestOper struct {
    DtSystemPowerOnSelfTestOper SystemPowerOnSelfTestOper `json:"power-on-self-test"`
}


type SystemPowerOnSelfTestOperOper struct {
    PowerOnLog []SystemPowerOnSelfTestOperOperPowerOnLog `json:"power-on-log"`
}


type SystemPowerOnSelfTestOperOperPowerOnLog struct {
    DlogData string `json:"dlog-data"`
    DlogDataSearch string `json:"dlog-data-search"`
}

func (p *SystemPowerOnSelfTestOper) GetId() string{
    return "1"
}

func (p *SystemPowerOnSelfTestOper) getPath() string{
    return "system/power-on-self-test/oper"
}

func (p *SystemPowerOnSelfTestOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemPowerOnSelfTestOper,error) {
logger.Println("SystemPowerOnSelfTestOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemPowerOnSelfTestOper
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
