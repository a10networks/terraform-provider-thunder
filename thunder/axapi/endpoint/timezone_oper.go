

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TimezoneOper struct {
    
    Oper TimezoneOperOper `json:"oper"`

}
type DataTimezoneOper struct {
    DtTimezoneOper TimezoneOper `json:"timezone"`
}


type TimezoneOperOper struct {
    StdName string `json:"std-name"`
    DstName string `json:"dst-name"`
    DenyDst string `json:"deny-dst"`
    Location string `json:"location"`
}

func (p *TimezoneOper) GetId() string{
    return "1"
}

func (p *TimezoneOper) getPath() string{
    return "timezone/oper"
}

func (p *TimezoneOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataTimezoneOper,error) {
logger.Println("TimezoneOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataTimezoneOper
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
