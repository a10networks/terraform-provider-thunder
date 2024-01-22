

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbDns64Stats struct {
    
    Stats SlbDns64StatsStats `json:"stats"`

}
type DataSlbDns64Stats struct {
    DtSlbDns64Stats SlbDns64Stats `json:"dns64"`
}


type SlbDns64StatsStats struct {
    Query int `json:"query"`
    QueryBadPkt int `json:"query-bad-pkt"`
    QueryChg int `json:"query-chg"`
    QueryParallel int `json:"query-parallel"`
    QueryPassive int `json:"query-passive"`
    Resp int `json:"resp"`
    RespBadPkt int `json:"resp-bad-pkt"`
    RespBadQr int `json:"resp-bad-qr"`
    RespChg int `json:"resp-chg"`
    RespErr int `json:"resp-err"`
    RespEmpty int `json:"resp-empty"`
    RespLocal int `json:"resp-local"`
    Adjust int `json:"adjust"`
    Cache int `json:"cache"`
    Drop int `json:"drop"`
}

func (p *SlbDns64Stats) GetId() string{
    return "1"
}

func (p *SlbDns64Stats) getPath() string{
    return "slb/dns64/stats"
}

func (p *SlbDns64Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbDns64Stats,error) {
logger.Println("SlbDns64Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbDns64Stats
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
