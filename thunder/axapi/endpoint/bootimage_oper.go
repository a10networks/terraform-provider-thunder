

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type BootimageOper struct {
    
    Oper BootimageOperOper `json:"oper"`

}
type DataBootimageOper struct {
    DtBootimageOper BootimageOper `json:"bootimage"`
}


type BootimageOperOper struct {
    HdPri string `json:"hd-pri"`
    HdSec string `json:"hd-sec"`
    CfPri string `json:"cf-pri"`
    CfSec string `json:"cf-sec"`
    HdDefault string `json:"hd-default"`
    CfDefault string `json:"cf-default"`
}

func (p *BootimageOper) GetId() string{
    return "1"
}

func (p *BootimageOper) getPath() string{
    return "bootimage/oper"
}

func (p *BootimageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataBootimageOper,error) {
logger.Println("BootimageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataBootimageOper
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
