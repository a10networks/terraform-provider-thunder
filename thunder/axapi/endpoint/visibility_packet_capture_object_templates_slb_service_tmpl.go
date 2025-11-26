package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl struct {
	Inst struct {
		CaptureConfig string `json:"capture-config"`

		Name string `json:"name"`

		TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsInc2812 `json:"trigger-stats-inc"`

		TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate2813 `json:"trigger-stats-rate"`

		TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsSeverity2814 `json:"trigger-stats-severity"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"slb-service-tmpl"`
}

type VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsInc2812 struct {
	Es_resp_300 int    `json:"es_resp_300"`
	Es_resp_400 int    `json:"es_resp_400"`
	Es_resp_500 int    `json:"es_resp_500"`
	Resp3xx     int    `json:"resp-3xx"`
	Resp4xx     int    `json:"resp-4xx"`
	Resp5xx     int    `json:"resp-5xx"`
	Uuid        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate2813 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Es_resp_300         int    `json:"es_resp_300"`
	Es_resp_400         int    `json:"es_resp_400"`
	Es_resp_500         int    `json:"es_resp_500"`
	Resp3xx             int    `json:"resp-3xx"`
	Resp4xx             int    `json:"resp-4xx"`
	Resp5xx             int    `json:"resp-5xx"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsSeverity2814 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl) getPath() string {
	return "visibility/packet-capture/object-templates/slb-service-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
