package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosTemplateTcpProgressionTrackingMitigationRequestTracking struct {
	Inst struct {
		FirstRequestMaxTime int `json:"first-request-max-time"`

		ProgressionTrackingReqAction string `json:"progression-tracking-req-action" dval:"drop"`

		ProgressionTrackingReqActionListName string `json:"progression-tracking-req-action-list-name"`

		ProgressionTrackingReqEnabled string `json:"progression-tracking-req-enabled"`

		RequestLengthMax int `json:"request-length-max"`

		RequestLengthMin int `json:"request-length-min"`

		RequestResponseModel string `json:"request-response-model" dval:"enable"`

		RequestToResponseMaxTime int `json:"request-to-response-max-time"`

		ResponseLengthMax int `json:"response-length-max"`

		ResponseLengthMin int `json:"response-length-min"`

		ResponseToRequestMaxTime int `json:"response-to-request-max-time"`

		Uuid string `json:"uuid"`

		Violation int `json:"violation"`

		Tcp_name string
	} `json:"request-tracking"`
}

func (p *DdosTemplateTcpProgressionTrackingMitigationRequestTracking) GetId() string {
	return "1"
}

func (p *DdosTemplateTcpProgressionTrackingMitigationRequestTracking) getPath() string {
	return "ddos/template/tcp/" + p.Inst.Tcp_name + "/progression-tracking/mitigation/request-tracking"
}

func (p *DdosTemplateTcpProgressionTrackingMitigationRequestTracking) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingMitigationRequestTracking::Post")
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

func (p *DdosTemplateTcpProgressionTrackingMitigationRequestTracking) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingMitigationRequestTracking::Get")
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
func (p *DdosTemplateTcpProgressionTrackingMitigationRequestTracking) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingMitigationRequestTracking::Put")
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

func (p *DdosTemplateTcpProgressionTrackingMitigationRequestTracking) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateTcpProgressionTrackingMitigationRequestTracking::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
