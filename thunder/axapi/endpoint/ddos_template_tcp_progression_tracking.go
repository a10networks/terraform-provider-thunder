package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosTemplateTcpProgressionTracking struct {
	Inst struct {
		IgnoreTlsHandshake int `json:"ignore-TLS-handshake"`

		Mitigation DdosTemplateTcpProgressionTrackingMitigation345 `json:"mitigation"`

		Profiling DdosTemplateTcpProgressionTrackingProfiling350 `json:"profiling"`

		ProgressionTrackingEnabled string `json:"progression-tracking-enabled"`

		Uuid string `json:"uuid"`

		Tcp_name string
	} `json:"progression-tracking"`
}

type DdosTemplateTcpProgressionTrackingMitigation345 struct {
	RequestTracking    DdosTemplateTcpProgressionTrackingMitigationRequestTracking346    `json:"request-tracking"`
	ConnectionTracking DdosTemplateTcpProgressionTrackingMitigationConnectionTracking347 `json:"connection-tracking"`
	TimeWindowTracking DdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking348 `json:"time-window-tracking"`
	SlowAttack         DdosTemplateTcpProgressionTrackingMitigationSlowAttack349         `json:"slow-attack"`
}

type DdosTemplateTcpProgressionTrackingMitigationRequestTracking346 struct {
	ProgressionTrackingReqEnabled        string `json:"progression-tracking-req-enabled"`
	RequestResponseModel                 string `json:"request-response-model" dval:"enable"`
	ResponseLengthMax                    int    `json:"response-length-max"`
	ResponseLengthMin                    int    `json:"response-length-min"`
	RequestLengthMin                     int    `json:"request-length-min"`
	RequestLengthMax                     int    `json:"request-length-max"`
	RequestToResponseMaxTime             int    `json:"request-to-response-max-time"`
	ResponseToRequestMaxTime             int    `json:"response-to-request-max-time"`
	FirstRequestMaxTime                  int    `json:"first-request-max-time"`
	ProgressionTrackingReqActionListName string `json:"progression-tracking-req-action-list-name"`
	Violation                            int    `json:"violation"`
	ProgressionTrackingReqAction         string `json:"progression-tracking-req-action" dval:"drop"`
	Uuid                                 string `json:"uuid"`
}

type DdosTemplateTcpProgressionTrackingMitigationConnectionTracking347 struct {
	ProgressionTrackingConnEnabled        string `json:"progression-tracking-conn-enabled"`
	ConnSentMax                           int    `json:"conn-sent-max"`
	ConnSentMin                           int    `json:"conn-sent-min"`
	ConnRcvdMax                           int    `json:"conn-rcvd-max"`
	ConnRcvdMin                           int    `json:"conn-rcvd-min"`
	ConnRcvdSentRatioMin                  int    `json:"conn-rcvd-sent-ratio-min"`
	ConnRcvdSentRatioMax                  int    `json:"conn-rcvd-sent-ratio-max"`
	ConnDurationMax                       int    `json:"conn-duration-max"`
	ConnDurationMin                       int    `json:"conn-duration-min"`
	ConnViolation                         int    `json:"conn-violation"`
	ProgressionTrackingConnActionListName string `json:"progression-tracking-conn-action-list-name"`
	ProgressionTrackingConnAction         string `json:"progression-tracking-conn-action" dval:"drop"`
	Uuid                                  string `json:"uuid"`
}

type DdosTemplateTcpProgressionTrackingMitigationTimeWindowTracking348 struct {
	ProgressionTrackingWinEnabled            string `json:"progression-tracking-win-enabled"`
	WindowSentMax                            int    `json:"window-sent-max"`
	WindowSentMin                            int    `json:"window-sent-min"`
	WindowRcvdMax                            int    `json:"window-rcvd-max"`
	WindowRcvdMin                            int    `json:"window-rcvd-min"`
	WindowRcvdSentRatioMin                   int    `json:"window-rcvd-sent-ratio-min"`
	WindowRcvdSentRatioMax                   int    `json:"window-rcvd-sent-ratio-max"`
	WindowViolation                          int    `json:"window-violation"`
	ProgressionTrackingWindowsActionListName string `json:"progression-tracking-windows-action-list-name"`
	ProgressionTrackingWindowsAction         string `json:"progression-tracking-windows-action" dval:"drop"`
	Uuid                                     string `json:"uuid"`
}

type DdosTemplateTcpProgressionTrackingMitigationSlowAttack349 struct {
	ResponsePktRateMax                    int    `json:"response-pkt-rate-max"`
	InitResponseMaxTime                   int    `json:"init-response-max-time"`
	InitRequestMaxTime                    int    `json:"init-request-max-time"`
	ProgressionTrackingSlowActionListName string `json:"progression-tracking-slow-action-list-name"`
	ProgressionTrackingSlowAction         string `json:"progression-tracking-slow-action" dval:"drop"`
	Uuid                                  string `json:"uuid"`
}

type DdosTemplateTcpProgressionTrackingProfiling350 struct {
	ProfilingRequestResponseModel int    `json:"profiling-request-response-model"`
	ProfilingConnectionLifeModel  int    `json:"profiling-connection-life-model"`
	ProfilingTimeWindowModel      int    `json:"profiling-time-window-model"`
	Uuid                          string `json:"uuid"`
}

func (p *DdosTemplateTcpProgressionTracking) GetId() string {
	return "1"
}

func (p *DdosTemplateTcpProgressionTracking) getPath() string {
	return "ddos/template/tcp/" + p.Inst.Tcp_name + "/progression-tracking"
}

func (p *DdosTemplateTcpProgressionTracking) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTracking::Post")
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

func (p *DdosTemplateTcpProgressionTracking) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTracking::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *DdosTemplateTcpProgressionTracking) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTracking::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}

func (p *DdosTemplateTcpProgressionTracking) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTracking::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
