

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemViewHotfixOper struct {
    
    Oper SystemViewHotfixOperOper `json:"oper"`

}
type DataSystemViewHotfixOper struct {
    DtSystemViewHotfixOper SystemViewHotfixOper `json:"hotfix"`
}


type SystemViewHotfixOperOper struct {
    Applied string `json:"applied"`
    Copied string `json:"copied"`
}

func (p *SystemViewHotfixOper) GetId() string{
    return "1"
}

func (p *SystemViewHotfixOper) getPath() string{
    return "system-view/hotfix/oper"
}

func (p *SystemViewHotfixOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemViewHotfixOper,error) {
logger.Println("SystemViewHotfixOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemViewHotfixOper
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
