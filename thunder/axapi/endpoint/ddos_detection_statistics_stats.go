

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionStatisticsStats struct {
    
    Stats DdosDetectionStatisticsStatsStats `json:"stats"`

}
type DataDdosDetectionStatisticsStats struct {
    DtDdosDetectionStatisticsStats DdosDetectionStatisticsStats `json:"statistics"`
}


type DdosDetectionStatisticsStatsStats struct {
    SflowPacketsReceived int `json:"sflow-packets-received"`
    NotSupportedSflowVer int `json:"not-supported-sflow-ver"`
    NetflowPktsReceived int `json:"netflow-pkts-received"`
    NotSupprotedNflowVer int `json:"not-supproted-nflow-ver"`
    AgentNotFound int `json:"agent-not-found"`
    PktDrop int `json:"pkt-drop"`
    ReportAllocFail int `json:"report-alloc-fail"`
    ReportEnqueueFail int `json:"report-enqueue-fail"`
    SampleProcssed int `json:"sample-procssed"`
    Ip_rcvd int `json:"ip_rcvd"`
    Ipv6_rcvd int `json:"ipv6_rcvd"`
    Frag_rcvd int `json:"frag_rcvd"`
    DstHit int `json:"dst-hit"`
    DstMiss int `json:"dst-miss"`
    DstLearn int `json:"dst-learn"`
    DstAge int `json:"dst-age"`
    DstServiceHit int `json:"dst-service-hit"`
    DstServiceMiss int `json:"dst-service-miss"`
    DstServiceLearn int `json:"dst-service-learn"`
    DstServiceAge int `json:"dst-service-age"`
    SrcHit int `json:"src-hit"`
    SrcMiss int `json:"src-miss"`
    SrcLearn int `json:"src-learn"`
    SrcAge int `json:"src-age"`
    EntryAllocFail int `json:"entry-alloc-fail"`
    GeoLearn int `json:"geo-learn"`
    GeoAge int `json:"geo-age"`
    UnmatchEntryPortZero int `json:"unmatch-entry-port-zero"`
    ObjectAllocOom int `json:"object-alloc-oom"`
    InvalidEvent int `json:"invalid-event"`
    RtbhStartSent int `json:"rtbh-start-sent"`
    RtbhStopSent int `json:"rtbh-stop-sent"`
    RtbhStartFail int `json:"rtbh-start-fail"`
    RtbhStopFail int `json:"rtbh-stop-fail"`
    InvalidProto int `json:"invalid-proto"`
    DstIpLearn int `json:"dst-ip-learn"`
    DstIpAge int `json:"dst-ip-age"`
    NSubnetLearned int `json:"n-subnet-learned"`
    NSubnetAged int `json:"n-subnet-aged"`
    NIpLearned int `json:"n-ip-learned"`
    NIpAged int `json:"n-ip-aged"`
    NServiceLearned int `json:"n-service-learned"`
    NServiceAged int `json:"n-service-aged"`
    NetworkMatchMiss int `json:"network-match-miss"`
    SessionMatchMiss int `json:"session-match-miss"`
    SessionAllocateFail int `json:"session-allocate-fail"`
    SessionLearned int `json:"session-learned"`
    SessionAged int `json:"session-aged"`
    SrcPortHit int `json:"src-port-hit"`
    SrcPortMiss int `json:"src-port-miss"`
    SrcPortLearn int `json:"src-port-learn"`
    SrcPortAge int `json:"src-port-age"`
    NServiceNotFound int `json:"n-service-not-found"`
    NSubnetCreateFail int `json:"n-subnet-create-fail"`
    NIpCreateFail int `json:"n-ip-create-fail"`
    NServiceCreateFail int `json:"n-service-create-fail"`
    DbUnexpectedError int `json:"db-unexpected-error"`
    DbOperFailure int `json:"db-oper-failure"`
    DbOpenFailure int `json:"db-open-failure"`
    DbNSubnetTableCreateFailure int `json:"db-n-subnet-table-create-failure"`
    DbNIpTableCreateFailure int `json:"db-n-ip-table-create-failure"`
    DbNSvcTableCreateFailure int `json:"db-n-svc-table-create-failure"`
    DbNSubnetSaveAttempt int `json:"db-n-subnet-save-attempt"`
    DbNSubnetSaveFailure int `json:"db-n-subnet-save-failure"`
    DbNSubnetRestoreAttempt int `json:"db-n-subnet-restore-attempt"`
    DbNIpSaveAttempt int `json:"db-n-ip-save-attempt"`
    DbNIpSaveFailure int `json:"db-n-ip-save-failure"`
    DbNIpRestoreAttempt int `json:"db-n-ip-restore-attempt"`
    DbNSvcSaveAttempt int `json:"db-n-svc-save-attempt"`
    DbNSvcSaveFailure int `json:"db-n-svc-save-failure"`
    DbNSvcRestoreAttempt int `json:"db-n-svc-restore-attempt"`
    DbNStaticSubnetNotFound int `json:"db-n-static-subnet-not-found"`
    DbNParentEntryNotFound int `json:"db-n-parent-entry-not-found"`
    DbWorkerEnqFailure int `json:"db-worker-enq-failure"`
    DbNSubnetTablePurgeFailure int `json:"db-n-subnet-table-purge-failure"`
    DbNIpTablePurgeFailure int `json:"db-n-ip-table-purge-failure"`
    DbNSvcTablePurgeFailure int `json:"db-n-svc-table-purge-failure"`
}

func (p *DdosDetectionStatisticsStats) GetId() string{
    return "1"
}

func (p *DdosDetectionStatisticsStats) getPath() string{
    return "ddos/detection/statistics/stats"
}

func (p *DdosDetectionStatisticsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDetectionStatisticsStats,error) {
logger.Println("DdosDetectionStatisticsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDetectionStatisticsStats
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
