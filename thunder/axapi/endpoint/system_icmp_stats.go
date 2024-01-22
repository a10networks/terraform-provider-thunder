

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIcmpStats struct {
    
    Stats SystemIcmpStatsStats `json:"stats"`

}
type DataSystemIcmpStats struct {
    DtSystemIcmpStats SystemIcmpStats `json:"icmp"`
}


type SystemIcmpStatsStats struct {
    Num int `json:"num"`
    Inmsgs int `json:"inmsgs"`
    Inerrors int `json:"inerrors"`
    Indestunreachs int `json:"indestunreachs"`
    Intimeexcds int `json:"intimeexcds"`
    Inparmprobs int `json:"inparmprobs"`
    Insrcquenchs int `json:"insrcquenchs"`
    Inredirects int `json:"inredirects"`
    Inechos int `json:"inechos"`
    Inechoreps int `json:"inechoreps"`
    Intimestamps int `json:"intimestamps"`
    Intimestampreps int `json:"intimestampreps"`
    Inaddrmasks int `json:"inaddrmasks"`
    Inaddrmaskreps int `json:"inaddrmaskreps"`
    Outmsgs int `json:"outmsgs"`
    Outerrors int `json:"outerrors"`
    Outdestunreachs int `json:"outdestunreachs"`
    Outtimeexcds int `json:"outtimeexcds"`
    Outparmprobs int `json:"outparmprobs"`
    Outsrcquenchs int `json:"outsrcquenchs"`
    Outredirects int `json:"outredirects"`
    Outechos int `json:"outechos"`
    Outechoreps int `json:"outechoreps"`
    Outtimestamps int `json:"outtimestamps"`
    Outtimestampreps int `json:"outtimestampreps"`
    Outaddrmasks int `json:"outaddrmasks"`
    Outaddrmaskreps int `json:"outaddrmaskreps"`
}

func (p *SystemIcmpStats) GetId() string{
    return "1"
}

func (p *SystemIcmpStats) getPath() string{
    return "system/icmp/stats"
}

func (p *SystemIcmpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemIcmpStats,error) {
logger.Println("SystemIcmpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemIcmpStats
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
