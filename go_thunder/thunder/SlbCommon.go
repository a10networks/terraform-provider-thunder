package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbCommon struct {
	SlbCommonInstancePortScanDetection SlbCommonInstance `json:"common,omitempty"`
}

type SlbCommonInstance struct {
	SlbCommonInstanceAflexTableEntryAgingInterval                 int                                      `json:"aflex-table-entry-aging-interval,omitempty"`
	SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncEnable SlbCommonInstanceAflexTableEntrySync     `json:"aflex-table-entry-sync,omitempty"`
	SlbCommonInstanceAfterDisable                                 int                                      `json:"after-disable,omitempty"`
	SlbCommonInstanceAllowInGatewayMode                           int                                      `json:"allow-in-gateway-mode,omitempty"`
	SlbCommonInstanceAttackRespCode                               int                                      `json:"attack-resp-code,omitempty"`
	SlbCommonInstanceAutoNatNoIPRefresh                           string                                   `json:"auto-nat-no-ip-refresh,omitempty"`
	SlbCommonInstanceAutoTranslatePort                            int                                      `json:"auto-translate-port,omitempty"`
	SlbCommonInstanceBuffThresh                                   int                                      `json:"buff-thresh,omitempty"`
	SlbCommonInstanceBuffThreshHwBuff                             int                                      `json:"buff-thresh-hw-buff,omitempty"`
	SlbCommonInstanceBuffThreshRelieveThresh                      int                                      `json:"buff-thresh-relieve-thresh,omitempty"`
	SlbCommonInstanceBuffThreshSysBuffHigh                        int                                      `json:"buff-thresh-sys-buff-high,omitempty"`
	SlbCommonInstanceBuffThreshSysBuffLow                         int                                      `json:"buff-thresh-sys-buff-low,omitempty"`
	SlbCommonInstanceCertPinningTTL                               SlbCommonInstanceCertPinning             `json:"cert-pinning,omitempty"`
	SlbCommonInstanceCompressBlockSize                            int                                      `json:"compress-block-size,omitempty"`
	SlbCommonInstanceConnRateLimitSrcIPList                       SlbCommonInstanceConnRateLimit           `json:"conn-rate-limit,omitempty"`
	SlbCommonInstanceDNSCacheAge                                  int                                      `json:"dns-cache-age,omitempty"`
	SlbCommonInstanceDNSCacheEnable                               int                                      `json:"dns-cache-enable,omitempty"`
	SlbCommonInstanceDNSCacheEntrySize                            int                                      `json:"dns-cache-entry-size,omitempty"`
	SlbCommonInstanceDNSResponseRateLimitingMaxTableEntries       SlbCommonInstanceDNSResponseRateLimiting `json:"dns-response-rate-limiting,omitempty"`
	SlbCommonInstanceDNSVipStateless                              int                                      `json:"dns-vip-stateless,omitempty"`
	SlbCommonInstanceDdosPktCountThresh                           int                                      `json:"ddos-pkt-count-thresh,omitempty"`
	SlbCommonInstanceDdosPktSizeThresh                            int                                      `json:"ddos-pkt-size-thresh,omitempty"`
	SlbCommonInstanceDdosProtectionIpdEnableToggle                SlbCommonInstanceDdosProtection          `json:"ddos-protection,omitempty"`
	SlbCommonInstanceDisableAdaptiveResourceCheck                 int                                      `json:"disable-adaptive-resource-check,omitempty"`
	SlbCommonInstanceDisablePersistScoring                        int                                      `json:"disable-persist-scoring,omitempty"`
	SlbCommonInstanceDisablePortMasking                           int                                      `json:"disable-port-masking,omitempty"`
	SlbCommonInstanceDisableServerAutoReselect                    int                                      `json:"disable-server-auto-reselect,omitempty"`
	SlbCommonInstanceDropIcmpToVipWhenVipDown                     int                                      `json:"drop-icmp-to-vip-when-vip-down,omitempty"`
	SlbCommonInstanceDsrHealthCheckEnable                         int                                      `json:"dsr-health-check-enable,omitempty"`
	SlbCommonInstanceEcmpHash                                     string                                   `json:"ecmp-hash,omitempty"`
	SlbCommonInstanceEnableL7ReqAcct                              int                                      `json:"enable-l7-req-acct,omitempty"`
	SlbCommonInstanceEntity                                       string                                   `json:"entity,omitempty"`
	SlbCommonInstanceExcludeDestination                           string                                   `json:"exclude-destination,omitempty"`
	SlbCommonInstanceExtendedStats                                int                                      `json:"extended-stats,omitempty"`
	SlbCommonInstanceFastPathDisable                              int                                      `json:"fast-path-disable,omitempty"`
	SlbCommonInstanceGatewayHealthCheck                           int                                      `json:"gateway-health-check,omitempty"`
	SlbCommonInstanceGracefulShutdown                             int                                      `json:"graceful-shutdown,omitempty"`
	SlbCommonInstanceGracefulShutdownEnable                       int                                      `json:"graceful-shutdown-enable,omitempty"`
	SlbCommonInstanceHTTPFastEnable                               int                                      `json:"http-fast-enable,omitempty"`
	SlbCommonInstanceHealthCheckToAllVip                          int                                      `json:"health-check-to-all-vip,omitempty"`
	SlbCommonInstanceHonorServerResponseTTL                       int                                      `json:"honor-server-response-ttl,omitempty"`
	SlbCommonInstanceHwCompression                                int                                      `json:"hw-compression,omitempty"`
	SlbCommonInstanceHwSynRr                                      int                                      `json:"hw-syn-rr,omitempty"`
	SlbCommonInstanceInterval                                     int                                      `json:"interval,omitempty"`
	SlbCommonInstanceIpv4Offset                                   int                                      `json:"ipv4-offset,omitempty"`
	SlbCommonInstanceL2L3TrunkLbDisable                           int                                      `json:"l2l3-trunk-lb-disable,omitempty"`
	SlbCommonInstanceLogForResetUnknownConn                       int                                      `json:"log-for-reset-unknown-conn,omitempty"`
	SlbCommonInstanceLowLatency                                   int                                      `json:"low-latency,omitempty"`
	SlbCommonInstanceMaxBuffQueuedPerConn                         int                                      `json:"max-buff-queued-per-conn,omitempty"`
	SlbCommonInstanceMaxHTTPHeaderCount                           int                                      `json:"max-http-header-count,omitempty"`
	SlbCommonInstanceMaxLocalRate                                 int                                      `json:"max-local-rate,omitempty"`
	SlbCommonInstanceMaxRemoteRate                                int                                      `json:"max-remote-rate,omitempty"`
	SlbCommonInstanceMslTime                                      int                                      `json:"msl-time,omitempty"`
	SlbCommonInstanceMssTable                                     int                                      `json:"mss-table,omitempty"`
	SlbCommonInstanceN5New                                        int                                      `json:"N5-new,omitempty"`
	SlbCommonInstanceN5Old                                        int                                      `json:"N5-old,omitempty"`
	SlbCommonInstanceNoAutoUpOnAflex                              int                                      `json:"no-auto-up-on-aflex,omitempty"`
	SlbCommonInstanceOneServerConnHmRate                          int                                      `json:"one-server-conn-hm-rate,omitempty"`
	SlbCommonInstanceOverridePort                                 int                                      `json:"override-port,omitempty"`
	SlbCommonInstancePerThrPercent                                int                                      `json:"per-thr-percent,omitempty"`
	SlbCommonInstancePingSweepDetection                           string                                   `json:"ping-sweep-detection,omitempty"`
	SlbCommonInstancePktRateForResetUnknownConn                   int                                      `json:"pkt-rate-for-reset-unknown-conn,omitempty"`
	SlbCommonInstancePlayerIDCheckEnable                          int                                      `json:"player-id-check-enable,omitempty"`
	SlbCommonInstancePortScanDetection                            string                                   `json:"port-scan-detection,omitempty"`
	SlbCommonInstancePreProcessEnable                             int                                      `json:"pre-process-enable,omitempty"`
	SlbCommonInstanceQAT                                          int                                      `json:"QAT,omitempty"`
	SlbCommonInstanceRange                                        int                                      `json:"range,omitempty"`
	SlbCommonInstanceRangeEnd                                     int                                      `json:"range-end,omitempty"`
	SlbCommonInstanceRangeStart                                   int                                      `json:"range-start,omitempty"`
	SlbCommonInstanceRateLimitLogging                             int                                      `json:"rate-limit-logging,omitempty"`
	SlbCommonInstanceResetStaleSession                            int                                      `json:"reset-stale-session,omitempty"`
	SlbCommonInstanceResolvePortConflict                          int                                      `json:"resolve-port-conflict,omitempty"`
	SlbCommonInstanceResponseType                                 string                                   `json:"response-type,omitempty"`
	SlbCommonInstanceScaleOut                                     int                                      `json:"scale-out,omitempty"`
	SlbCommonInstanceScaleOutTrafficMap                           int                                      `json:"scale-out-traffic-map,omitempty"`
	SlbCommonInstanceServiceGroupOnNoDestNatVports                string                                   `json:"service-group-on-no-dest-nat-vports,omitempty"`
	SlbCommonInstanceShowSlbServerLegacyCmd                       int                                      `json:"show-slb-server-legacy-cmd,omitempty"`
	SlbCommonInstanceShowSlbServiceGroupLegacyCmd                 int                                      `json:"show-slb-service-group-legacy-cmd,omitempty"`
	SlbCommonInstanceShowSlbVirtualServerLegacyCmd                int                                      `json:"show-slb-virtual-server-legacy-cmd,omitempty"`
	SlbCommonInstanceSnatGwyForL3                                 int                                      `json:"snat-gwy-for-l3,omitempty"`
	SlbCommonInstanceSnatOnVip                                    int                                      `json:"snat-on-vip,omitempty"`
	SlbCommonInstanceSnatPreserveRange                            SlbCommonInstanceSnatPreserve            `json:"snat-preserve,omitempty"`
	SlbCommonInstanceSoftware                                     int                                      `json:"software,omitempty"`
	SlbCommonInstanceSoftwareTLS13                                int                                      `json:"software-tls13,omitempty"`
	SlbCommonInstanceSoftwareTLS13Offload                         int                                      `json:"software-tls13-offload,omitempty"`
	SlbCommonInstanceSortRes                                      int                                      `json:"sort-res,omitempty"`
	SlbCommonInstanceSsliSniHashEnable                            int                                      `json:"ssli-sni-hash-enable,omitempty"`
	SlbCommonInstanceStatelessSgMultiBinding                      int                                      `json:"stateless-sg-multi-binding,omitempty"`
	SlbCommonInstanceStatsDataDisable                             int                                      `json:"stats-data-disable,omitempty"`
	SlbCommonInstanceSubstituteSourceMac                          int                                      `json:"substitute-source-mac,omitempty"`
	SlbCommonInstanceTTLThreshold                                 int                                      `json:"ttl-threshold,omitempty"`
	SlbCommonInstanceTimeout                                      int                                      `json:"timeout,omitempty"`
	SlbCommonInstanceTrafficMapType                               string                                   `json:"traffic-map-type,omitempty"`
	SlbCommonInstanceUUID                                         string                                   `json:"uuid,omitempty"`
	SlbCommonInstanceUseDefaultSessCount                          int                                      `json:"use-default-sess-count,omitempty"`
	SlbCommonInstanceUseMssTab                                    int                                      `json:"use-mss-tab,omitempty"`
}

type SlbCommonInstanceAflexTableEntrySync struct {
	SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncEnable      int    `json:"aflex-table-entry-sync-enable,omitempty"`
	SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncMaxKeyLen   int    `json:"aflex-table-entry-sync-max-key-len,omitempty"`
	SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncMaxValueLen int    `json:"aflex-table-entry-sync-max-value-len,omitempty"`
	SlbCommonInstanceAflexTableEntrySyncAflexTableEntrySyncMinLifetime int    `json:"aflex-table-entry-sync-min-lifetime,omitempty"`
	SlbCommonInstanceAflexTableEntrySyncUUID                           string `json:"uuid,omitempty"`
}

type SlbCommonInstanceCertPinning struct {
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInEnable SlbCommonInstanceCertPinningCandidateListFeedbackOptIn `json:"candidate-list-feedback-opt-in,omitempty"`
	SlbCommonInstanceCertPinningTTL                              int                                                    `json:"ttl,omitempty"`
	SlbCommonInstanceCertPinningUUID                             string                                                 `json:"uuid,omitempty"`
}

type SlbCommonInstanceConnRateLimit struct {
	SlbCommonInstanceConnRateLimitSrcIPListProtocol []SlbCommonInstanceConnRateLimitSrcIPList `json:"src-ip-list,omitempty"`
}

type SlbCommonInstanceDNSResponseRateLimiting struct {
	SlbCommonInstanceDNSResponseRateLimitingMaxTableEntries int    `json:"max-table-entries,omitempty"`
	SlbCommonInstanceDNSResponseRateLimitingUUID            string `json:"uuid,omitempty"`
}

type SlbCommonInstanceDdosProtection struct {
	SlbCommonInstanceDdosProtectionIpdEnableToggle         string                                          `json:"ipd-enable-toggle,omitempty"`
	SlbCommonInstanceDdosProtectionLoggingIpdLoggingToggle SlbCommonInstanceDdosProtectionLogging          `json:"logging,omitempty"`
	SlbCommonInstanceDdosProtectionPacketsPerSecondIpdTCP  SlbCommonInstanceDdosProtectionPacketsPerSecond `json:"packets-per-second,omitempty"`
}

type SlbCommonInstanceSnatPreserve struct {
	SlbCommonInstanceSnatPreserveRangePort1 []SlbCommonInstanceSnatPreserveRange `json:"range,omitempty"`
}

type SlbCommonInstanceCertPinningCandidateListFeedbackOptIn struct {
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInDaily       int    `json:"daily,omitempty"`
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInDayTime     string `json:"day-time,omitempty"`
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInEnable      int    `json:"enable,omitempty"`
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInSchedule    int    `json:"schedule,omitempty"`
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInUUID        string `json:"uuid,omitempty"`
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInUseMgmtPort int    `json:"use-mgmt-port,omitempty"`
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInWeekDay     string `json:"week-day,omitempty"`
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInWeekTime    string `json:"week-time,omitempty"`
	SlbCommonInstanceCertPinningCandidateListFeedbackOptInWeekly      int    `json:"weekly,omitempty"`
}

type SlbCommonInstanceConnRateLimitSrcIPList struct {
	SlbCommonInstanceConnRateLimitSrcIPListExceedAction int    `json:"exceed-action,omitempty"`
	SlbCommonInstanceConnRateLimitSrcIPListLimit        int    `json:"limit,omitempty"`
	SlbCommonInstanceConnRateLimitSrcIPListLimitPeriod  string `json:"limit-period,omitempty"`
	SlbCommonInstanceConnRateLimitSrcIPListLockOut      int    `json:"lock-out,omitempty"`
	SlbCommonInstanceConnRateLimitSrcIPListLog          int    `json:"log,omitempty"`
	SlbCommonInstanceConnRateLimitSrcIPListProtocol     string `json:"protocol,omitempty"`
	SlbCommonInstanceConnRateLimitSrcIPListShared       int    `json:"shared,omitempty"`
	SlbCommonInstanceConnRateLimitSrcIPListUUID         string `json:"uuid,omitempty"`
}

type SlbCommonInstanceDdosProtectionLogging struct {
	SlbCommonInstanceDdosProtectionLoggingIpdLoggingToggle string `json:"ipd-logging-toggle,omitempty"`
}

type SlbCommonInstanceDdosProtectionPacketsPerSecond struct {
	SlbCommonInstanceDdosProtectionPacketsPerSecondIpdTCP int `json:"ipd-tcp,omitempty"`
	SlbCommonInstanceDdosProtectionPacketsPerSecondIpdUDP int `json:"ipd-udp,omitempty"`
}

type SlbCommonInstanceSnatPreserveRange struct {
	SlbCommonInstanceSnatPreserveRangePort1 int `json:"port1,omitempty"`
	SlbCommonInstanceSnatPreserveRangePort2 int `json:"port2,omitempty"`
}

func PostSlbCommon(id string, inst SlbCommon, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbCommon")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/common", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbCommon
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbCommon", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbCommon(id string, host string) (*SlbCommon, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbCommon")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/common", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbCommon
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbCommon", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
