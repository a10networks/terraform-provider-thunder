

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAlgH323Stats struct {
    
    Stats Cgnv6LsnAlgH323StatsStats `json:"stats"`

}
type DataCgnv6LsnAlgH323Stats struct {
    DtCgnv6LsnAlgH323Stats Cgnv6LsnAlgH323Stats `json:"h323"`
}


type Cgnv6LsnAlgH323StatsStats struct {
    H225rasMessage int `json:"h225ras-message"`
    H225csMessage int `json:"h225cs-message"`
    H245ctlMessage int `json:"h245ctl-message"`
    H245Tunneled int `json:"h245-tunneled"`
    FastStart int `json:"fast-start"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
}

func (p *Cgnv6LsnAlgH323Stats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAlgH323Stats) getPath() string{
    return "cgnv6/lsn/alg/h323/stats"
}

func (p *Cgnv6LsnAlgH323Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnAlgH323Stats,error) {
logger.Println("Cgnv6LsnAlgH323Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnAlgH323Stats
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
