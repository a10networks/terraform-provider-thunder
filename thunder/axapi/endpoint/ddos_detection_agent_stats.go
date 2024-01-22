

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionAgentStats struct {
    
    AgentName string `json:"agent-name"`
    Stats DdosDetectionAgentStatsStats `json:"stats"`

}
type DataDdosDetectionAgentStats struct {
    DtDdosDetectionAgentStats DdosDetectionAgentStats `json:"agent"`
}


type DdosDetectionAgentStatsStats struct {
    SflowPacketsReceived int `json:"sflow-packets-received"`
    SflowSamplesReceived int `json:"sflow-samples-received"`
    SflowSamplesBadLen int `json:"sflow-samples-bad-len"`
    SflowSamplesNonStd int `json:"sflow-samples-non-std"`
    SflowSamplesSkipped int `json:"sflow-samples-skipped"`
    SflowSampleRecordBadLen int `json:"sflow-sample-record-bad-len"`
    SflowSamplesSentForDetection int `json:"sflow-samples-sent-for-detection"`
    SflowSampleRecordInvalidLayer2 int `json:"sflow-sample-record-invalid-layer2"`
    SflowSampleIpv6HdrParseFail int `json:"sflow-sample-ipv6-hdr-parse-fail"`
    SflowDisabled int `json:"sflow-disabled"`
    NetflowDisabled int `json:"netflow-disabled"`
    NetflowV5PacketsReceived int `json:"netflow-v5-packets-received"`
    NetflowV5SamplesReceived int `json:"netflow-v5-samples-received"`
    NetflowV5SamplesSentForDetection int `json:"netflow-v5-samples-sent-for-detection"`
    NetflowV5SampleRecordsBadLen int `json:"netflow-v5-sample-records-bad-len"`
    NetflowV5MaxRecordsExceed int `json:"netflow-v5-max-records-exceed"`
    NetflowV9PacketsReceived int `json:"netflow-v9-packets-received"`
    NetflowV9SamplesReceived int `json:"netflow-v9-samples-received"`
    NetflowV9SamplesSentForDetection int `json:"netflow-v9-samples-sent-for-detection"`
    NetflowV9SampleRecordsBadLen int `json:"netflow-v9-sample-records-bad-len"`
    NetflowV9SampleFlowsetBadPadding int `json:"netflow-v9-sample-flowset-bad-padding"`
    NetflowV9MaxRecordsExceed int `json:"netflow-v9-max-records-exceed"`
    NetflowV9TemplateNotFound int `json:"netflow-v9-template-not-found"`
    NetflowV10PacketsReceived int `json:"netflow-v10-packets-received"`
    NetflowV10SamplesReceived int `json:"netflow-v10-samples-received"`
    NetflowV10SamplesSentForDetection int `json:"netflow-v10-samples-sent-for-detection"`
    NetflowV10SampleRecordsBadLen int `json:"netflow-v10-sample-records-bad-len"`
    NetflowV10MaxRecordsExceed int `json:"netflow-v10-max-records-exceed"`
    NetflowTcpSampleReceived int `json:"netflow-tcp-sample-received"`
    NetflowUdpSampleReceived int `json:"netflow-udp-sample-received"`
    NetflowIcmpSampleReceived int `json:"netflow-icmp-sample-received"`
    NetflowOtherSampleReceived int `json:"netflow-other-sample-received"`
    NetflowRecordCopyOomError int `json:"netflow-record-copy-oom-error"`
    NetflowRecordRseInvalid int `json:"netflow-record-rse-invalid"`
    NetflowSampleFlowDurError int `json:"netflow-sample-flow-dur-error"`
    FlowDstEntryMiss int `json:"flow-dst-entry-miss"`
    FlowIpProtoOrPortMiss int `json:"flow-ip-proto-or-port-miss"`
    FlowDetectionMsgqFull int `json:"flow-detection-msgq-full"`
    FlowNetworkEntryMiss int `json:"flow-network-entry-miss"`
}

func (p *DdosDetectionAgentStats) GetId() string{
    return "1"
}

func (p *DdosDetectionAgentStats) getPath() string{
    
    return "ddos/detection/agent/"+p.AgentName+"/stats"
}

func (p *DdosDetectionAgentStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDetectionAgentStats,error) {
logger.Println("DdosDetectionAgentStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDetectionAgentStats
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
