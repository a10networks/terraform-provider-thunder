

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemHrxqStatusOper struct {
    
    Oper SystemHrxqStatusOperOper `json:"oper"`

}
type DataSystemHrxqStatusOper struct {
    DtSystemHrxqStatusOper SystemHrxqStatusOper `json:"hrxq-status"`
}


type SystemHrxqStatusOperOper struct {
    DeepHrxqEnable string `json:"deep-hrxq-enable"`
    Hrxq_num_chunks int `json:"hrxq_num_chunks"`
}

func (p *SystemHrxqStatusOper) GetId() string{
    return "1"
}

func (p *SystemHrxqStatusOper) getPath() string{
    return "system/hrxq-status/oper"
}

func (p *SystemHrxqStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemHrxqStatusOper,error) {
logger.Println("SystemHrxqStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemHrxqStatusOper
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
