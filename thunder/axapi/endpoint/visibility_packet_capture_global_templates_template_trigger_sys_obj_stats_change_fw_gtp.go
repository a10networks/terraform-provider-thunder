package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp struct {
	Inst struct {
		TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2130 `json:"trigger-stats-inc"`

		TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2131 `json:"trigger-stats-rate"`

		Uuid string `json:"uuid"`

		Template_name string
	} `json:"fw-gtp"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2130 struct {
	OutOfSessionMemory                int    `json:"out-of-session-memory"`
	BladeOutOfSessionMemory           int    `json:"blade-out-of-session-memory"`
	GtpSmpPathCheckFailed             int    `json:"gtp-smp-path-check-failed"`
	GtpSmpCCheckFailed                int    `json:"gtp-smp-c-check-failed"`
	BladeGtpSmpPathCheckFailed        int    `json:"blade-gtp-smp-path-check-failed"`
	BladeGtpSmpCCheckFailed           int    `json:"blade-gtp-smp-c-check-failed"`
	GtpTunnelRateLimitEntryCreateFail int    `json:"gtp-tunnel-rate-limit-entry-create-fail"`
	GtpUTunnelRateLimitEntryCreateFa  int    `json:"gtp-u-tunnel-rate-limit-entry-create-fa"`
	GtpRateLimitSmpCreateFailure      int    `json:"gtp-rate-limit-smp-create-failure"`
	GtpRateLimitT3CtrCreateFailure    int    `json:"gtp-rate-limit-t3-ctr-create-failure"`
	GtpRateLimitEntryCreateFailure    int    `json:"gtp-rate-limit-entry-create-failure"`
	BladeGtpRateLimitSmpCreateFailure int    `json:"blade-gtp-rate-limit-smp-create-failure"`
	BladeGtpRateLimitT3CtrCreateFail  int    `json:"blade-gtp-rate-limit-t3-ctr-create-fail"`
	BladeGtpRateLimitEntryCreateFailu int    `json:"blade-gtp-rate-limit-entry-create-failu"`
	GtpSmpDecSessCountCheckFailed     int    `json:"gtp-smp-dec-sess-count-check-failed"`
	GtpUSmpCheckFailed                int    `json:"gtp-u-smp-check-failed"`
	GtpInfoExtNotFound                int    `json:"gtp-info-ext-not-found"`
	BladeGtpSmpDecSessCountCheckFail  int    `json:"blade-gtp-smp-dec-sess-count-check-fail"`
	BladeGtpUSmpCheckFailed           int    `json:"blade-gtp-u-smp-check-failed"`
	BladeGtpInfoExtNotFound           int    `json:"blade-gtp-info-ext-not-found"`
	BladeGtpSmpSessionCountCheckFaile int    `json:"blade-gtp-smp-session-count-check-faile"`
	GtpCSmpSigCheckFailed             int    `json:"gtp-c-smp-sig-check-failed"`
	BladeGtpCSmpSigCheckFailed        int    `json:"blade-gtp-c-smp-sig-check-failed"`
	GtpUSmpSigCheckFailed             int    `json:"gtp-u-smp-sig-check-failed"`
	BladeGtpUSmpSigCheckFailed        int    `json:"blade-gtp-u-smp-sig-check-failed"`
	GtpSmpSigCheckFailed              int    `json:"gtp-smp-sig-check-failed"`
	BladeGtpSmpSigCheckFailed         int    `json:"blade-gtp-smp-sig-check-failed"`
	GtpCFailConnCreateSlow            int    `json:"gtp-c-fail-conn-create-slow"`
	GtpUFailConnCreateSlow            int    `json:"gtp-u-fail-conn-create-slow"`
	GtpPathmFailConnCreateSlow        int    `json:"gtp-pathm-fail-conn-create-slow"`
	Uuid                              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2131 struct {
	ThresholdExceededBy               int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                          int    `json:"duration" dval:"60"`
	OutOfSessionMemory                int    `json:"out-of-session-memory"`
	BladeOutOfSessionMemory           int    `json:"blade-out-of-session-memory"`
	GtpSmpPathCheckFailed             int    `json:"gtp-smp-path-check-failed"`
	GtpSmpCCheckFailed                int    `json:"gtp-smp-c-check-failed"`
	BladeGtpSmpPathCheckFailed        int    `json:"blade-gtp-smp-path-check-failed"`
	BladeGtpSmpCCheckFailed           int    `json:"blade-gtp-smp-c-check-failed"`
	GtpTunnelRateLimitEntryCreateFail int    `json:"gtp-tunnel-rate-limit-entry-create-fail"`
	GtpUTunnelRateLimitEntryCreateFa  int    `json:"gtp-u-tunnel-rate-limit-entry-create-fa"`
	GtpRateLimitSmpCreateFailure      int    `json:"gtp-rate-limit-smp-create-failure"`
	GtpRateLimitT3CtrCreateFailure    int    `json:"gtp-rate-limit-t3-ctr-create-failure"`
	GtpRateLimitEntryCreateFailure    int    `json:"gtp-rate-limit-entry-create-failure"`
	BladeGtpRateLimitSmpCreateFailure int    `json:"blade-gtp-rate-limit-smp-create-failure"`
	BladeGtpRateLimitT3CtrCreateFail  int    `json:"blade-gtp-rate-limit-t3-ctr-create-fail"`
	BladeGtpRateLimitEntryCreateFailu int    `json:"blade-gtp-rate-limit-entry-create-failu"`
	GtpSmpDecSessCountCheckFailed     int    `json:"gtp-smp-dec-sess-count-check-failed"`
	GtpUSmpCheckFailed                int    `json:"gtp-u-smp-check-failed"`
	GtpInfoExtNotFound                int    `json:"gtp-info-ext-not-found"`
	BladeGtpSmpDecSessCountCheckFail  int    `json:"blade-gtp-smp-dec-sess-count-check-fail"`
	BladeGtpUSmpCheckFailed           int    `json:"blade-gtp-u-smp-check-failed"`
	BladeGtpInfoExtNotFound           int    `json:"blade-gtp-info-ext-not-found"`
	BladeGtpSmpSessionCountCheckFaile int    `json:"blade-gtp-smp-session-count-check-faile"`
	GtpCSmpSigCheckFailed             int    `json:"gtp-c-smp-sig-check-failed"`
	BladeGtpCSmpSigCheckFailed        int    `json:"blade-gtp-c-smp-sig-check-failed"`
	GtpUSmpSigCheckFailed             int    `json:"gtp-u-smp-sig-check-failed"`
	BladeGtpUSmpSigCheckFailed        int    `json:"blade-gtp-u-smp-sig-check-failed"`
	GtpSmpSigCheckFailed              int    `json:"gtp-smp-sig-check-failed"`
	BladeGtpSmpSigCheckFailed         int    `json:"blade-gtp-smp-sig-check-failed"`
	GtpCFailConnCreateSlow            int    `json:"gtp-c-fail-conn-create-slow"`
	GtpUFailConnCreateSlow            int    `json:"gtp-u-fail-conn-create-slow"`
	GtpPathmFailConnCreateSlow        int    `json:"gtp-pathm-fail-conn-create-slow"`
	Uuid                              string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp) GetId() string {
	return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp) getPath() string {
	return "visibility/packet-capture/global-templates/template/" + p.Inst.Template_name + "/trigger-sys-obj-stats-change/fw-gtp"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
