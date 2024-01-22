

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nat64HwAccelerateStats struct {
    
    Stats Cgnv6Nat64HwAccelerateStatsStats `json:"stats"`

}
type DataCgnv6Nat64HwAccelerateStats struct {
    DtCgnv6Nat64HwAccelerateStats Cgnv6Nat64HwAccelerateStats `json:"hw-accelerate"`
}


type Cgnv6Nat64HwAccelerateStatsStats struct {
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

func (p *Cgnv6Nat64HwAccelerateStats) GetId() string{
    return "1"
}

func (p *Cgnv6Nat64HwAccelerateStats) getPath() string{
    return "cgnv6/nat64/hw-accelerate/stats"
}

func (p *Cgnv6Nat64HwAccelerateStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Nat64HwAccelerateStats,error) {
logger.Println("Cgnv6Nat64HwAccelerateStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Nat64HwAccelerateStats
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
