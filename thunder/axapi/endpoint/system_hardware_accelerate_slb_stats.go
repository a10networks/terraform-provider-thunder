

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemHardwareAccelerateSlbStats struct {
    
    Stats SystemHardwareAccelerateSlbStatsStats `json:"stats"`

}
type DataSystemHardwareAccelerateSlbStats struct {
    DtSystemHardwareAccelerateSlbStats SystemHardwareAccelerateSlbStats `json:"slb"`
}


type SystemHardwareAccelerateSlbStatsStats struct {
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

func (p *SystemHardwareAccelerateSlbStats) GetId() string{
    return "1"
}

func (p *SystemHardwareAccelerateSlbStats) getPath() string{
    return "system/hardware-accelerate/slb/stats"
}

func (p *SystemHardwareAccelerateSlbStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemHardwareAccelerateSlbStats,error) {
logger.Println("SystemHardwareAccelerateSlbStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemHardwareAccelerateSlbStats
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
