

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatAlgSipStats struct {
    
    Stats Cgnv6FixedNatAlgSipStatsStats `json:"stats"`

}
type DataCgnv6FixedNatAlgSipStats struct {
    DtCgnv6FixedNatAlgSipStats Cgnv6FixedNatAlgSipStats `json:"sip"`
}


type Cgnv6FixedNatAlgSipStatsStats struct {
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

func (p *Cgnv6FixedNatAlgSipStats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatAlgSipStats) getPath() string{
    return "cgnv6/fixed-nat/alg/sip/stats"
}

func (p *Cgnv6FixedNatAlgSipStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatAlgSipStats,error) {
logger.Println("Cgnv6FixedNatAlgSipStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatAlgSipStats
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
