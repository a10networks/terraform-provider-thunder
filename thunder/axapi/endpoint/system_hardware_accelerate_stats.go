

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemHardwareAccelerateStats struct {
    
    Slb SystemHardwareAccelerateStatsSlb `json:"slb"`
    Stats SystemHardwareAccelerateStatsStats `json:"stats"`

}
type DataSystemHardwareAccelerateStats struct {
    DtSystemHardwareAccelerateStats SystemHardwareAccelerateStats `json:"hardware-accelerate"`
}


type SystemHardwareAccelerateStatsSlb struct {
    Stats SystemHardwareAccelerateStatsSlbStats `json:"stats"`
}


type SystemHardwareAccelerateStatsSlbStats struct {
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


type SystemHardwareAccelerateStatsStats struct {
    HitCounts int `json:"hit-counts"`
    HitIndex int `json:"hit-index"`
    Ipv4ForwardCounts int `json:"ipv4-forward-counts"`
    Ipv6ForwardCounts int `json:"ipv6-forward-counts"`
    HwFwdModuleStatus int `json:"hw-fwd-module-status"`
    HwFwdProgReqs int `json:"hw-fwd-prog-reqs"`
    HwFwdProgErrors int `json:"hw-fwd-prog-errors"`
    HwFwdFlowSinglebitErrors int `json:"hw-fwd-flow-singlebit-errors"`
    HwFwdFlowTagMismatch int `json:"hw-fwd-flow-tag-mismatch"`
    HwFwdFlowSeqMismatch int `json:"hw-fwd-flow-seq-mismatch"`
    HwFwdAgeoutDropCount int `json:"hw-fwd-ageout-drop-count"`
    HwFwdInvalidationDrop int `json:"hw-fwd-invalidation-drop"`
    HwFwdFlowHitIndex int `json:"hw-fwd-flow-hit-index"`
    HwFwdFlowReasonFlags int `json:"hw-fwd-flow-reason-flags"`
    HwFwdFlowDropCount int `json:"hw-fwd-flow-drop-count"`
    HwFwdFlowErrorCount int `json:"hw-fwd-flow-error-count"`
    HwFwdFlowUnalignCount int `json:"hw-fwd-flow-unalign-count"`
    HwFwdFlowUnderflowCount int `json:"hw-fwd-flow-underflow-count"`
    HwFwdFlowTxFullDrop int `json:"hw-fwd-flow-tx-full-drop"`
    HwFwdFlowQdrFullDrop int `json:"hw-fwd-flow-qdr-full-drop"`
    HwFwdPhyportMismatchDrop int `json:"hw-fwd-phyport-mismatch-drop"`
    HwFwdVlanidMismatchDrop int `json:"hw-fwd-vlanid-mismatch-drop"`
    HwFwdVmidDrop int `json:"hw-fwd-vmid-drop"`
    HwFwdProtocolMismatchDrop int `json:"hw-fwd-protocol-mismatch-drop"`
    HwFwdAvailIpv4Entry int `json:"hw-fwd-avail-ipv4-entry"`
    HwFwdAvailIpv6Entry int `json:"hw-fwd-avail-ipv6-entry"`
    HwFwdRateDropCount int `json:"hw-fwd-rate-drop-count"`
    HwFwdNormalAgeoutRcvd int `json:"hw-fwd-normal-ageout-rcvd"`
    HwFwdTcpFinAgeoutRcvd int `json:"hw-fwd-tcp-fin-ageout-rcvd"`
    HwFwdTcpRstAgeoutRcvd int `json:"hw-fwd-tcp-rst-ageout-rcvd"`
    HwFwdLookupFailRcvd int `json:"hw-fwd-lookup-fail-rcvd"`
    HwFwdStatsUpdateRcvd int `json:"hw-fwd-stats-update-rcvd"`
    HwFwdFlowSflowCount int `json:"hw-fwd-flow-sflow-count"`
}

func (p *SystemHardwareAccelerateStats) GetId() string{
    return "1"
}

func (p *SystemHardwareAccelerateStats) getPath() string{
    return "system/hardware-accelerate/stats"
}

func (p *SystemHardwareAccelerateStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemHardwareAccelerateStats,error) {
logger.Println("SystemHardwareAccelerateStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemHardwareAccelerateStats
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
