

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NgWafStatusOper struct {
    
    Oper NgWafStatusOperOper `json:"oper"`

}
type DataNgWafStatusOper struct {
    DtNgWafStatusOper NgWafStatusOper `json:"status"`
}


type NgWafStatusOperOper struct {
    NgwafVersion string `json:"ngwaf-version"`
    PartitionList []NgWafStatusOperOperPartitionList `json:"partition-list"`
}


type NgWafStatusOperOperPartitionList struct {
    PartitionName string `json:"partition-name"`
    Status string `json:"status"`
    AgentName string `json:"agent-name"`
    AccessKeyId string `json:"access-key-id"`
    SecretAccessKey string `json:"secret-access-key"`
    CacheEntries int `json:"cache-entries"`
    TrackedCustomSignal int `json:"tracked-custom-signal"`
}

func (p *NgWafStatusOper) GetId() string{
    return "1"
}

func (p *NgWafStatusOper) getPath() string{
    return "ng-waf/status/oper"
}

func (p *NgWafStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNgWafStatusOper,error) {
logger.Println("NgWafStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNgWafStatusOper
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
