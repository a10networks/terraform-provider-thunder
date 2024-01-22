

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwAlgSipStats struct {
    
    Stats FwAlgSipStatsStats `json:"stats"`

}
type DataFwAlgSipStats struct {
    DtFwAlgSipStats FwAlgSipStats `json:"sip"`
}


type FwAlgSipStatsStats struct {
    StatRequest int `json:"stat-request"`
    StatResponse int `json:"stat-response"`
    MethodRegister int `json:"method-register"`
    MethodInvite int `json:"method-invite"`
    MethodAck int `json:"method-ack"`
    MethodCancel int `json:"method-cancel"`
    MethodBye int `json:"method-bye"`
    MethodOptions int `json:"method-options"`
    MethodPrack int `json:"method-prack"`
    MethodSubscribe int `json:"method-subscribe"`
    MethodNotify int `json:"method-notify"`
    MethodPublish int `json:"method-publish"`
    MethodInfo int `json:"method-info"`
    MethodRefer int `json:"method-refer"`
    MethodMessage int `json:"method-message"`
    MethodUpdate int `json:"method-update"`
    MethodUnknown int `json:"method-unknown"`
}

func (p *FwAlgSipStats) GetId() string{
    return "1"
}

func (p *FwAlgSipStats) getPath() string{
    return "fw/alg/sip/stats"
}

func (p *FwAlgSipStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwAlgSipStats,error) {
logger.Println("FwAlgSipStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwAlgSipStats
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
