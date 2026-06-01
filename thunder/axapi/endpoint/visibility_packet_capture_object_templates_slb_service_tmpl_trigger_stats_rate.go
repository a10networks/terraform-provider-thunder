package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate struct {
	Inst struct {
		Duration int `json:"duration" dval:"60"`

		Es_resp_300 int `json:"es_resp_300"`

		Es_resp_400 int `json:"es_resp_400"`

		Es_resp_500 int `json:"es_resp_500"`

		Resp3xx int `json:"resp-3xx"`

		Resp4xx int `json:"resp-4xx"`

		Resp5xx int `json:"resp-5xx"`

		ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`

		Uuid string `json:"uuid"`

		Slb_service_tmpl_name string
	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate) GetId() string {
	return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate) getPath() string {
	return "visibility/packet-capture/object-templates/slb-service-tmpl/" + p.Inst.Slb_service_tmpl_name + "/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
