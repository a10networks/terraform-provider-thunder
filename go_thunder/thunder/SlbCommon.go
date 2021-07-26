package go_thunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type SlbCommon struct {
	UUID SlbCommonInstance `json:"common,omitempty"`
}

type DNSResponseRateLimiting struct {
	UUID            string `json:"uuid,omitempty"`
	MaxTableEntries int    `json:"max-table-entries,omitempty"`
}
type SrcIPList struct {
	Protocol     string `json:"protocol,omitempty"`
	Log          int    `json:"log,omitempty"`
	LockOut      int    `json:"lock-out,omitempty"`
	LimitPeriod  string `json:"limit-period,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	ExceedAction int    `json:"exceed-action,omitempty"`
	Shared       int    `json:"shared,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}
type ConnRateLimit struct {
	Protocol []SrcIPList `json:"src-ip-list,omitempty"`
}
type PacketsPerSecond struct {
	IpdTCP int `json:"ipd-tcp,omitempty"`
	IpdUDP int `json:"ipd-udp,omitempty"`
}
type Logging2 struct {
	IpdLoggingToggle string `json:"ipd-logging-toggle,omitempty"`
}
type DdosProtection struct {
	IpdTCP           PacketsPerSecond `json:"packets-per-second,omitempty"`
	IpdLoggingToggle Logging2         `json:"logging,omitempty"`
	IpdEnableToggle  string           `json:"ipd-enable-toggle,omitempty"`
}
type SlbCommonInstance struct {
	LowLatency                   int                     `json:"low-latency,omitempty"`
	UseMssTab                    int                     `json:"use-mss-tab,omitempty"`
	StatsDataDisable             int                     `json:"stats-data-disable,omitempty"`
	CompressBlockSize            int                     `json:"compress-block-size,omitempty"`
	PlayerIDCheckEnable          int                     `json:"player-id-check-enable,omitempty"`
	AfterDisable                 int                     `json:"after-disable,omitempty"`
	MslTime                      int                     `json:"msl-time,omitempty"`
	GracefulShutdownEnable       int                     `json:"graceful-shutdown-enable,omitempty"`
	BuffThreshHwBuff             int                     `json:"buff-thresh-hw-buff,omitempty"`
	HwSynRr                      int                     `json:"hw-syn-rr,omitempty"`
	Entity                       string                  `json:"entity,omitempty"`
	ResetStaleSession            int                     `json:"reset-stale-session,omitempty"`
	GatewayHealthCheck           int                     `json:"gateway-health-check,omitempty"`
	ScaleOut                     int                     `json:"scale-out,omitempty"`
	GracefulShutdown             int                     `json:"graceful-shutdown,omitempty"`
	RateLimitLogging             int                     `json:"rate-limit-logging,omitempty"`
	FastPathDisable              int                     `json:"fast-path-disable,omitempty"`
	DropIcmpToVipWhenVipDown     int                     `json:"drop-icmp-to-vip-when-vip-down,omitempty"`
	SsliSniHashEnable            int                     `json:"ssli-sni-hash-enable,omitempty"`
	HwCompression                int                     `json:"hw-compression,omitempty"`
	DNSVipStateless              int                     `json:"dns-vip-stateless,omitempty"`
	BuffThreshSysBuffLow         int                     `json:"buff-thresh-sys-buff-low,omitempty"`
	RangeEnd                     int                     `json:"range-end,omitempty"`
	MaxTableEntries              DNSResponseRateLimiting `json:"dns-response-rate-limiting,omitempty"`
	DNSCacheEnable               int                     `json:"dns-cache-enable,omitempty"`
	MaxLocalRate                 int                     `json:"max-local-rate,omitempty"`
	ExcludeDestination           string                  `json:"exclude-destination,omitempty"`
	DNSCacheAge                  int                     `json:"dns-cache-age,omitempty"`
	MaxHTTPHeaderCount           int                     `json:"max-http-header-count,omitempty"`
	L2L3TrunkLbDisable           int                     `json:"l2l3-trunk-lb-disable,omitempty"`
	ResolvePortConflict          int                     `json:"resolve-port-conflict,omitempty"`
	SortRes                      int                     `json:"sort-res,omitempty"`
	SnatGwyForL3                 int                     `json:"snat-gwy-for-l3,omitempty"`
	BuffThreshRelieveThresh      int                     `json:"buff-thresh-relieve-thresh,omitempty"`
	DsrHealthCheckEnable         int                     `json:"dsr-health-check-enable,omitempty"`
	BuffThresh                   int                     `json:"buff-thresh,omitempty"`
	DNSCacheEntrySize            int                     `json:"dns-cache-entry-size,omitempty"`
	LogForResetUnknownConn       int                     `json:"log-for-reset-unknown-conn,omitempty"`
	AutoNatNoIPRefresh           string                  `json:"auto-nat-no-ip-refresh,omitempty"`
	PktRateForResetUnknownConn   int                     `json:"pkt-rate-for-reset-unknown-conn,omitempty"`
	BuffThreshSysBuffHigh        int                     `json:"buff-thresh-sys-buff-high,omitempty"`
	MaxBuffQueuedPerConn         int                     `json:"max-buff-queued-per-conn,omitempty"`
	MaxRemoteRate                int                     `json:"max-remote-rate,omitempty"`
	TTLThreshold                 int                     `json:"ttl-threshold,omitempty"`
	ExtendedStats                int                     `json:"extended-stats,omitempty"`
	EnableL7ReqAcct              int                     `json:"enable-l7-req-acct,omitempty"`
	UUID                         string                  `json:"uuid,omitempty"`
	SnatOnVip                    int                     `json:"snat-on-vip,omitempty"`
	RangeStart                   int                     `json:"range-start,omitempty"`
	HonorServerResponseTTL       int                     `json:"honor-server-response-ttl,omitempty"`
	Interval                     int                     `json:"interval,omitempty"`
	StatelessSgMultiBinding      int                     `json:"stateless-sg-multi-binding,omitempty"`
	DisableAdaptiveResourceCheck int                     `json:"disable-adaptive-resource-check,omitempty"`
	Range                        int                     `json:"range,omitempty"`
	Protocol                     ConnRateLimit           `json:"conn-rate-limit,omitempty"`
	MssTable                     int                     `json:"mss-table,omitempty"`
	Timeout                      int                     `json:"timeout,omitempty"`
	ResponseType                 string                  `json:"response-type,omitempty"`
	IpdTCP                       DdosProtection          `json:"ddos-protection,omitempty"`
	OverridePort                 int                     `json:"override-port,omitempty"`
	NoAutoUpOnAflex              int                     `json:"no-auto-up-on-aflex,omitempty"`
	DisableServerAutoReselect    int                     `json:"disable-server-auto-reselect,omitempty"`
	Software                     int                     `json:"software,omitempty"`
}

func GetSlbCommon(id string, host string) (*SlbCommon, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/common", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbCommon
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GetSlbCommon REQ RES..........................", m)
			err := check_api_status("GetSlbCommon", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostSlbCommon(id string, vc SlbCommon, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/common", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v SlbCommon
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostSlbCommon", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
