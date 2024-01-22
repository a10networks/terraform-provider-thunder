

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatHwAccelerateStats struct {
    
    Stats Cgnv6FixedNatHwAccelerateStatsStats `json:"stats"`

}
type DataCgnv6FixedNatHwAccelerateStats struct {
    DtCgnv6FixedNatHwAccelerateStats Cgnv6FixedNatHwAccelerateStats `json:"hw-accelerate"`
}


type Cgnv6FixedNatHwAccelerateStatsStats struct {
    EntryCreate int `json:"entry-create"`
    EntryCreateFailure int `json:"entry-create-failure"`
    EntryCreateFailServerDown int `json:"entry-create-fail-server-down"`
    EntryCreateFailMaxEntry int `json:"entry-create-fail-max-entry"`
    EntryFree int `json:"entry-free"`
    EntryFreeOppEntry int `json:"entry-free-opp-entry"`
    EntryFreeNoHwProg int `json:"entry-free-no-hw-prog"`
    EntryFreeNoConn int `json:"entry-free-no-conn"`
    EntryFreeNoSwEntry int `json:"entry-free-no-sw-entry"`
    EntryCounter int `json:"entry-counter"`
    EntryAgeOut int `json:"entry-age-out"`
    EntryAgeOutIdle int `json:"entry-age-out-idle"`
    EntryAgeOutTcpFin int `json:"entry-age-out-tcp-fin"`
    EntryAgeOutTcpRst int `json:"entry-age-out-tcp-rst"`
    EntryAgeOutInvalidDst int `json:"entry-age-out-invalid-dst"`
    EntryForceHwInvalidate int `json:"entry-force-hw-invalidate"`
    EntryInvalidateServerDown int `json:"entry-invalidate-server-down"`
    TcamCreate int `json:"tcam-create"`
    TcamFree int `json:"tcam-free"`
    TcamCounter int `json:"tcam-counter"`
}

func (p *Cgnv6FixedNatHwAccelerateStats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatHwAccelerateStats) getPath() string{
    return "cgnv6/fixed-nat/hw-accelerate/stats"
}

func (p *Cgnv6FixedNatHwAccelerateStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatHwAccelerateStats,error) {
logger.Println("Cgnv6FixedNatHwAccelerateStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatHwAccelerateStats
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
