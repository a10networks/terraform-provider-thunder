package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy struct {
	Inst struct {
		TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2205 `json:"trigger-stats-inc"`

		TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2206 `json:"trigger-stats-rate"`

		Uuid string `json:"uuid"`

		Template_name string
	} `json:"slb-ssl-forward-proxy"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2205 struct {
	FailedInSslHandshakes           int    `json:"failed-in-ssl-handshakes"`
	FailedInCryptoOperations        int    `json:"failed-in-crypto-operations"`
	FailedInTcp                     int    `json:"failed-in-tcp"`
	FailedInCertificateVerification int    `json:"failed-in-certificate-verification"`
	FailedInCertificateSigning      int    `json:"failed-in-certificate-signing"`
	InvalidOcspStaplingResponse     int    `json:"invalid-ocsp-stapling-response"`
	RevokedOcspResponse             int    `json:"revoked-ocsp-response"`
	UnsupportedSslVersion           int    `json:"unsupported-ssl-version"`
	ConnectionsFailed               int    `json:"connections-failed"`
	Uuid                            string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2206 struct {
	ThresholdExceededBy             int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                        int    `json:"duration" dval:"60"`
	FailedInSslHandshakes           int    `json:"failed-in-ssl-handshakes"`
	FailedInCryptoOperations        int    `json:"failed-in-crypto-operations"`
	FailedInTcp                     int    `json:"failed-in-tcp"`
	FailedInCertificateVerification int    `json:"failed-in-certificate-verification"`
	FailedInCertificateSigning      int    `json:"failed-in-certificate-signing"`
	InvalidOcspStaplingResponse     int    `json:"invalid-ocsp-stapling-response"`
	RevokedOcspResponse             int    `json:"revoked-ocsp-response"`
	UnsupportedSslVersion           int    `json:"unsupported-ssl-version"`
	ConnectionsFailed               int    `json:"connections-failed"`
	Uuid                            string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy) GetId() string {
	return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy) getPath() string {
	return "visibility/packet-capture/global-templates/template/" + p.Inst.Template_name + "/trigger-sys-obj-stats-change/slb-ssl-forward-proxy"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
