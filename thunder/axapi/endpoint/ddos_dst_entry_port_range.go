

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortRange struct {
	Inst struct {

    CaptureConfig DdosDstEntryPortRangeCaptureConfig `json:"capture-config"`
    Deny int `json:"deny"`
    DetectionEnable int `json:"detection-enable"`
    EnableTopK int `json:"enable-top-k"`
    Glid string `json:"glid"`
    GlidExceedAction DdosDstEntryPortRangeGlidExceedAction `json:"glid-exceed-action"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    IpFilteringPolicyOper DdosDstEntryPortRangeIpFilteringPolicyOper163 `json:"ip-filtering-policy-oper"`
    PatternRecognition DdosDstEntryPortRangePatternRecognition164 `json:"pattern-recognition"`
    PatternRecognitionPuDetails DdosDstEntryPortRangePatternRecognitionPuDetails165 `json:"pattern-recognition-pu-details"`
    PortInd DdosDstEntryPortRangePortInd166 `json:"port-ind"`
    PortRangeEnd int `json:"port-range-end"`
    PortRangeStart int `json:"port-range-start"`
    ProgressionTracking DdosDstEntryPortRangeProgressionTracking168 `json:"progression-tracking"`
    Protocol string `json:"protocol"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    Sflow DdosDstEntryPortRangeSflow `json:"sflow"`
    Template DdosDstEntryPortRangeTemplate `json:"template"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
    TopkSources DdosDstEntryPortRangeTopkSources169 `json:"topk-sources"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"port-range"`
}


type DdosDstEntryPortRangeCaptureConfig struct {
    CaptureConfigName string `json:"capture-config-name"`
    CaptureConfigMode string `json:"capture-config-mode"`
}


type DdosDstEntryPortRangeGlidExceedAction struct {
    StatelessEncapActionCfg DdosDstEntryPortRangeGlidExceedActionStatelessEncapActionCfg `json:"stateless-encap-action-cfg"`
}


type DdosDstEntryPortRangeGlidExceedActionStatelessEncapActionCfg struct {
    StatelessEncapAction string `json:"stateless-encap-action"`
    EncapTemplate string `json:"encap-template"`
}


type DdosDstEntryPortRangeIpFilteringPolicyOper163 struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortRangePatternRecognition164 struct {
    Algorithm string `json:"algorithm"`
    Mode string `json:"mode"`
    Sensitivity string `json:"sensitivity"`
    FilterThreshold int `json:"filter-threshold"`
    FilterInactiveThreshold int `json:"filter-inactive-threshold"`
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortRangePatternRecognitionPuDetails165 struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortRangePortInd166 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []DdosDstEntryPortRangePortIndSamplingEnable167 `json:"sampling-enable"`
}


type DdosDstEntryPortRangePortIndSamplingEnable167 struct {
    Counters1 string `json:"counters1"`
}


type DdosDstEntryPortRangeProgressionTracking168 struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortRangeSflow struct {
    Polling DdosDstEntryPortRangeSflowPolling `json:"polling"`
}


type DdosDstEntryPortRangeSflowPolling struct {
    SflowPackets int `json:"sflow-packets"`
    SflowTcp DdosDstEntryPortRangeSflowPollingSflowTcp `json:"sflow-tcp"`
    SflowHttp int `json:"sflow-http"`
}


type DdosDstEntryPortRangeSflowPollingSflowTcp struct {
    SflowTcpBasic int `json:"sflow-tcp-basic"`
    SflowTcpStateful int `json:"sflow-tcp-stateful"`
}


type DdosDstEntryPortRangeTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
}


type DdosDstEntryPortRangeTopkSources169 struct {
    Uuid string `json:"uuid"`
}

func (p *DdosDstEntryPortRange) GetId() string{
    return strconv.Itoa(p.Inst.PortRangeStart)+"+"+strconv.Itoa(p.Inst.PortRangeEnd)+"+"+p.Inst.Protocol
}

func (p *DdosDstEntryPortRange) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/port-range"
}

func (p *DdosDstEntryPortRange) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRange::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *DdosDstEntryPortRange) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRange::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *DdosDstEntryPortRange) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRange::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *DdosDstEntryPortRange) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRange::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
