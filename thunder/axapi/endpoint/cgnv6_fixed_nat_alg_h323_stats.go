

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatAlgH323Stats struct {
    
    Stats Cgnv6FixedNatAlgH323StatsStats `json:"stats"`

}
type DataCgnv6FixedNatAlgH323Stats struct {
    DtCgnv6FixedNatAlgH323Stats Cgnv6FixedNatAlgH323Stats `json:"h323"`
}


type Cgnv6FixedNatAlgH323StatsStats struct {
    H225rasMessage int `json:"h225ras-message"`
    H225csMessage int `json:"h225cs-message"`
    H245ctlMessage int `json:"h245ctl-message"`
    H245Tunneled int `json:"h245-tunneled"`
    FastStart int `json:"fast-start"`
}

func (p *Cgnv6FixedNatAlgH323Stats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatAlgH323Stats) getPath() string{
    return "cgnv6/fixed-nat/alg/h323/stats"
}

func (p *Cgnv6FixedNatAlgH323Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatAlgH323Stats,error) {
logger.Println("Cgnv6FixedNatAlgH323Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatAlgH323Stats
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
