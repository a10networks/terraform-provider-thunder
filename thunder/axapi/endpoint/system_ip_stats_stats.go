

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpStatsStats struct {
    
    Stats SystemIpStatsStatsStats `json:"stats"`

}
type DataSystemIpStatsStats struct {
    DtSystemIpStatsStats SystemIpStatsStats `json:"ip-stats"`
}


type SystemIpStatsStatsStats struct {
    Inreceives int `json:"inreceives"`
    Inhdrerrors int `json:"inhdrerrors"`
    Intoobigerrors int `json:"intoobigerrors"`
    Innoroutes int `json:"innoroutes"`
    Inaddrerrors int `json:"inaddrerrors"`
    Inunknownprotos int `json:"inunknownprotos"`
    Intruncatedpkts int `json:"intruncatedpkts"`
    Indiscards int `json:"indiscards"`
    Indelivers int `json:"indelivers"`
    Outforwdatagrams int `json:"outforwdatagrams"`
    Outrequests int `json:"outrequests"`
    Outdiscards int `json:"outdiscards"`
    Outnoroutes int `json:"outnoroutes"`
    Reasmtimeout int `json:"reasmtimeout"`
    Reasmreqds int `json:"reasmreqds"`
    Reasmoks int `json:"reasmoks"`
    Reasmfails int `json:"reasmfails"`
    Fragoks int `json:"fragoks"`
    Fragfails int `json:"fragfails"`
    Fragcreates int `json:"fragcreates"`
    Inmcastpkts int `json:"inmcastpkts"`
    Outmcastpkts int `json:"outmcastpkts"`
}

func (p *SystemIpStatsStats) GetId() string{
    return "1"
}

func (p *SystemIpStatsStats) getPath() string{
    return "system/ip-stats/stats"
}

func (p *SystemIpStatsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemIpStatsStats,error) {
logger.Println("SystemIpStatsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemIpStatsStats
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
