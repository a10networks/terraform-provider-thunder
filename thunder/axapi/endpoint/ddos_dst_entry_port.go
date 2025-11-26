package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type DdosDstEntryPort struct {
	Inst struct {
		CaptureConfig DdosDstEntryPortCaptureConfig `json:"capture-config"`

		Deny int `json:"deny"`

		DetectionEnable int `json:"detection-enable"`

		DnsCache string `json:"dns-cache"`

		EnableTopK int `json:"enable-top-k"`

		Glid string `json:"glid"`

		GlidExceedAction DdosDstEntryPortGlidExceedAction `json:"glid-exceed-action"`

		PatternRecognition DdosDstEntryPortPatternRecognition188 `json:"pattern-recognition"`

		PatternRecognitionPuDetails DdosDstEntryPortPatternRecognitionPuDetails189 `json:"pattern-recognition-pu-details"`

		PortInd DdosDstEntryPortPortInd190 `json:"port-ind"`

		PortNum int `json:"port-num"`

		ProgressionTracking DdosDstEntryPortProgressionTracking192 `json:"progression-tracking"`

		Protocol string `json:"protocol"`

		SetCounterBaseVal int `json:"set-counter-base-val"`

		Sflow DdosDstEntryPortSflow `json:"sflow"`

		SignatureExtraction DdosDstEntryPortSignatureExtraction193 `json:"signature-extraction"`

		Template DdosDstEntryPortTemplate `json:"template"`

		TopkNumRecords int `json:"topk-num-records" dval:"20"`

		TopkSortKey string `json:"topk-sort-key" dval:"avg"`

		TopkSources DdosDstEntryPortTopkSources194 `json:"topk-sources"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		DstEntryName string
	} `json:"port"`
}

type DdosDstEntryPortCaptureConfig struct {
	CaptureConfigName string `json:"capture-config-name"`
	CaptureConfigMode string `json:"capture-config-mode"`
}

type DdosDstEntryPortGlidExceedAction struct {
	StatelessEncapActionCfg DdosDstEntryPortGlidExceedActionStatelessEncapActionCfg `json:"stateless-encap-action-cfg"`
}

type DdosDstEntryPortGlidExceedActionStatelessEncapActionCfg struct {
	StatelessEncapAction string `json:"stateless-encap-action"`
	EncapTemplate        string `json:"encap-template"`
}

type DdosDstEntryPortPatternRecognition188 struct {
	Algorithm               string `json:"algorithm"`
	Mode                    string `json:"mode"`
	Sensitivity             string `json:"sensitivity"`
	FilterThreshold         int    `json:"filter-threshold"`
	FilterInactiveThreshold int    `json:"filter-inactive-threshold"`
	Uuid                    string `json:"uuid"`
}

type DdosDstEntryPortPatternRecognitionPuDetails189 struct {
	Uuid string `json:"uuid"`
}

type DdosDstEntryPortPortInd190 struct {
	Uuid           string                                     `json:"uuid"`
	SamplingEnable []DdosDstEntryPortPortIndSamplingEnable191 `json:"sampling-enable"`
}

type DdosDstEntryPortPortIndSamplingEnable191 struct {
	Counters1 string `json:"counters1"`
}

type DdosDstEntryPortProgressionTracking192 struct {
	Uuid string `json:"uuid"`
}

type DdosDstEntryPortSflow struct {
	Polling DdosDstEntryPortSflowPolling `json:"polling"`
}

type DdosDstEntryPortSflowPolling struct {
	SflowPackets int                                  `json:"sflow-packets"`
	SflowTcp     DdosDstEntryPortSflowPollingSflowTcp `json:"sflow-tcp"`
	SflowHttp    int                                  `json:"sflow-http"`
}

type DdosDstEntryPortSflowPollingSflowTcp struct {
	SflowTcpBasic    int `json:"sflow-tcp-basic"`
	SflowTcpStateful int `json:"sflow-tcp-stateful"`
}

type DdosDstEntryPortSignatureExtraction193 struct {
	Algorithm  string `json:"algorithm"`
	ManualMode int    `json:"manual-mode"`
	Uuid       string `json:"uuid"`
}

type DdosDstEntryPortTemplate struct {
	Dns   string `json:"dns"`
	Http  string `json:"http"`
	SslL4 string `json:"ssl-l4"`
	Sip   string `json:"sip"`
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
}

type DdosDstEntryPortTopkSources194 struct {
	Uuid string `json:"uuid"`
}

func (p *DdosDstEntryPort) GetId() string {
	return strconv.Itoa(p.Inst.PortNum) + "+" + p.Inst.Protocol
}

func (p *DdosDstEntryPort) getPath() string {
	return "ddos/dst/entry/" + p.Inst.DstEntryName + "/port"
}

func (p *DdosDstEntryPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstEntryPort::Post")
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

func (p *DdosDstEntryPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstEntryPort::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *DdosDstEntryPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstEntryPort::Put")
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

func (p *DdosDstEntryPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstEntryPort::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
