

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6SctpRateLimitEntriesOper struct {
    
    Oper Cgnv6SctpRateLimitEntriesOperOper `json:"oper"`

}
type DataCgnv6SctpRateLimitEntriesOper struct {
    DtCgnv6SctpRateLimitEntriesOper Cgnv6SctpRateLimitEntriesOper `json:"rate-limit-entries"`
}


type Cgnv6SctpRateLimitEntriesOperOper struct {
    RateLimitEntriesList []Cgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList `json:"rate-limit-entries-list"`
    EntryCount int `json:"entry-count"`
}


type Cgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList struct {
    Address string `json:"address"`
    Direction string `json:"direction"`
    Pps int `json:"pps"`
    RateLimit int `json:"rate-limit"`
}

func (p *Cgnv6SctpRateLimitEntriesOper) GetId() string{
    return "1"
}

func (p *Cgnv6SctpRateLimitEntriesOper) getPath() string{
    return "cgnv6/sctp/rate-limit-entries/oper"
}

func (p *Cgnv6SctpRateLimitEntriesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6SctpRateLimitEntriesOper,error) {
logger.Println("Cgnv6SctpRateLimitEntriesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6SctpRateLimitEntriesOper
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
