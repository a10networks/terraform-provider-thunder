package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc struct {
	Inst struct {
		BladeGtpCSmpSigCheckFailed int `json:"blade-gtp-c-smp-sig-check-failed"`

		BladeGtpInfoExtNotFound int `json:"blade-gtp-info-ext-not-found"`

		BladeGtpRateLimitEntryCreateFailu int `json:"blade-gtp-rate-limit-entry-create-failu"`

		BladeGtpRateLimitSmpCreateFailure int `json:"blade-gtp-rate-limit-smp-create-failure"`

		BladeGtpRateLimitT3CtrCreateFail int `json:"blade-gtp-rate-limit-t3-ctr-create-fail"`

		BladeGtpSmpCCheckFailed int `json:"blade-gtp-smp-c-check-failed"`

		BladeGtpSmpDecSessCountCheckFail int `json:"blade-gtp-smp-dec-sess-count-check-fail"`

		BladeGtpSmpPathCheckFailed int `json:"blade-gtp-smp-path-check-failed"`

		BladeGtpSmpSessionCountCheckFaile int `json:"blade-gtp-smp-session-count-check-faile"`

		BladeGtpSmpSigCheckFailed int `json:"blade-gtp-smp-sig-check-failed"`

		BladeGtpUSmpCheckFailed int `json:"blade-gtp-u-smp-check-failed"`

		BladeGtpUSmpSigCheckFailed int `json:"blade-gtp-u-smp-sig-check-failed"`

		BladeOutOfSessionMemory int `json:"blade-out-of-session-memory"`

		GtpCFailConnCreateSlow int `json:"gtp-c-fail-conn-create-slow"`

		GtpCSmpSigCheckFailed int `json:"gtp-c-smp-sig-check-failed"`

		GtpInfoExtNotFound int `json:"gtp-info-ext-not-found"`

		GtpPathmFailConnCreateSlow int `json:"gtp-pathm-fail-conn-create-slow"`

		GtpRateLimitEntryCreateFailure int `json:"gtp-rate-limit-entry-create-failure"`

		GtpRateLimitSmpCreateFailure int `json:"gtp-rate-limit-smp-create-failure"`

		GtpRateLimitT3CtrCreateFailure int `json:"gtp-rate-limit-t3-ctr-create-failure"`

		GtpSmpCCheckFailed int `json:"gtp-smp-c-check-failed"`

		GtpSmpDecSessCountCheckFailed int `json:"gtp-smp-dec-sess-count-check-failed"`

		GtpSmpPathCheckFailed int `json:"gtp-smp-path-check-failed"`

		GtpSmpSigCheckFailed int `json:"gtp-smp-sig-check-failed"`

		GtpTunnelRateLimitEntryCreateFail int `json:"gtp-tunnel-rate-limit-entry-create-fail"`

		GtpUFailConnCreateSlow int `json:"gtp-u-fail-conn-create-slow"`

		GtpUSmpCheckFailed int `json:"gtp-u-smp-check-failed"`

		GtpUSmpSigCheckFailed int `json:"gtp-u-smp-sig-check-failed"`

		GtpUTunnelRateLimitEntryCreateFa int `json:"gtp-u-tunnel-rate-limit-entry-create-fa"`

		OutOfSessionMemory int `json:"out-of-session-memory"`

		Uuid string `json:"uuid"`

		Template_name string
	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc) GetId() string {
	return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc) getPath() string {
	return "visibility/packet-capture/global-templates/template/" + p.Inst.Template_name + "/trigger-sys-obj-stats-change/fw-gtp/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
