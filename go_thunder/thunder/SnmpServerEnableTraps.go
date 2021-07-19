package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTraps struct {
	Lldp SnmpServerEnableTrapsInstance `json:"traps,omitempty"`
}

type SnmpServerEnableTrapsInstance struct {
	All                           int                       `json:"all,omitempty"`
	Group                         SnmpServerEnableGslb      `json:"gslb,omitempty"`
	Lldp                          int                       `json:"lldp,omitempty"`
	FixedNatPortMappingFileChange SnmpServerEnableLsn       `json:"lsn,omitempty"`
	TrunkPortThreshold            SnmpServerEnableNetwork   `json:"network,omitempty"`
	Bgp                           SnmpServerEnableRouting   `json:"routing,omitempty"`
	ApplicationBufferLimit        SnmpServerEnableSlb       `json:"slb,omitempty"`
	ConnectionResourceEvent       SnmpServerEnableSlbChange `json:"slb-change,omitempty"`
	Linkup                        SnmpServerEnableSnmp      `json:"snmp,omitempty"`
	ServerCertificateError        SnmpServerEnableSsl       `json:"ssl,omitempty"`
	ControlCPUHigh                SnmpServerEnableSystem    `json:"system,omitempty"`
	UUID                          string                    `json:"uuid,omitempty"`
	StateChange                   SnmpServerEnableVcs       `json:"vcs,omitempty"`
	Active                        SnmpServerEnableVrrpA     `json:"vrrp-a,omitempty"`
}

type SnmpServerEnableGslb struct {
	All       int    `json:"all,omitempty"`
	Group     int    `json:"group,omitempty"`
	ServiceIP int    `json:"service-ip,omitempty"`
	Site      int    `json:"site,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	Zone      int    `json:"zone,omitempty"`
}

type SnmpServerEnableLsn struct {
	All                           int    `json:"all,omitempty"`
	FixedNatPortMappingFileChange int    `json:"fixed-nat-port-mapping-file-change,omitempty"`
	MaxIpportThreshold            int    `json:"max-ipport-threshold,omitempty"`
	MaxPortThreshold              int    `json:"max-port-threshold,omitempty"`
	PerIPPortUsageThreshold       int    `json:"per-ip-port-usage-threshold,omitempty"`
	TotalPortUsageThreshold       int    `json:"total-port-usage-threshold,omitempty"`
	TrafficExceeded               int    `json:"traffic-exceeded,omitempty"`
	UUID                          string `json:"uuid,omitempty"`
}

type SnmpServerEnableNetwork struct {
	TrunkPortThreshold int    `json:"trunk-port-threshold,omitempty"`
	UUID               string `json:"uuid,omitempty"`
}

type SnmpServerEnableRouting struct {
	BgpEstablishedNotification SnmpServerEnableBgp  `json:"bgp,omitempty"`
	IsisAuthenticationFailure  SnmpServerEnableIsis `json:"isis,omitempty"`
	OspfLsdbOverflow           SnmpServerEnableOspf `json:"ospf,omitempty"`
}

type SnmpServerEnableSlb struct {
	All                    int    `json:"all,omitempty"`
	ApplicationBufferLimit int    `json:"application-buffer-limit,omitempty"`
	BwRateLimitExceed      int    `json:"bw-rate-limit-exceed,omitempty"`
	BwRateLimitResume      int    `json:"bw-rate-limit-resume,omitempty"`
	GatewayDown            int    `json:"gateway-down,omitempty"`
	GatewayUp              int    `json:"gateway-up,omitempty"`
	ServerConnLimit        int    `json:"server-conn-limit,omitempty"`
	ServerConnResume       int    `json:"server-conn-resume,omitempty"`
	ServerDisabled         int    `json:"server-disabled,omitempty"`
	ServerDown             int    `json:"server-down,omitempty"`
	ServerSelectionFailure int    `json:"server-selection-failure,omitempty"`
	ServerUp               int    `json:"server-up,omitempty"`
	ServiceConnLimit       int    `json:"service-conn-limit,omitempty"`
	ServiceConnResume      int    `json:"service-conn-resume,omitempty"`
	ServiceDown            int    `json:"service-down,omitempty"`
	ServiceGroupDown       int    `json:"service-group-down,omitempty"`
	ServiceGroupMemberDown int    `json:"service-group-member-down,omitempty"`
	ServiceGroupMemberUp   int    `json:"service-group-member-up,omitempty"`
	ServiceGroupUp         int    `json:"service-group-up,omitempty"`
	ServiceUp              int    `json:"service-up,omitempty"`
	UUID                   string `json:"uuid,omitempty"`
	VipConnlimit           int    `json:"vip-connlimit,omitempty"`
	VipConnratelimit       int    `json:"vip-connratelimit,omitempty"`
	VipDown                int    `json:"vip-down,omitempty"`
	VipPortConnlimit       int    `json:"vip-port-connlimit,omitempty"`
	VipPortConnratelimit   int    `json:"vip-port-connratelimit,omitempty"`
	VipPortDown            int    `json:"vip-port-down,omitempty"`
	VipPortUp              int    `json:"vip-port-up,omitempty"`
	VipUp                  int    `json:"vip-up,omitempty"`
}

type SnmpServerEnableSlbChange struct {
	All                     int    `json:"all,omitempty"`
	ConnectionResourceEvent int    `json:"connection-resource-event,omitempty"`
	ResourceUsageWarning    int    `json:"resource-usage-warning,omitempty"`
	Server                  int    `json:"server,omitempty"`
	ServerPort              int    `json:"server-port,omitempty"`
	SslCertChange           int    `json:"ssl-cert-change,omitempty"`
	SslCertExpire           int    `json:"ssl-cert-expire,omitempty"`
	SystemThreshold         int    `json:"system-threshold,omitempty"`
	UUID                    string `json:"uuid,omitempty"`
	Vip                     int    `json:"vip,omitempty"`
	VipPort                 int    `json:"vip-port,omitempty"`
}

type SnmpServerEnableSnmp struct {
	All      int    `json:"all,omitempty"`
	Linkdown int    `json:"linkdown,omitempty"`
	Linkup   int    `json:"linkup,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}

type SnmpServerEnableSsl struct {
	ServerCertificateError int    `json:"server-certificate-error,omitempty"`
	UUID                   string `json:"uuid,omitempty"`
}

type SnmpServerEnableSystem struct {
	All                int    `json:"all,omitempty"`
	ControlCPUHigh     int    `json:"control-cpu-high,omitempty"`
	DataCPUHigh        int    `json:"data-cpu-high,omitempty"`
	Fan                int    `json:"fan,omitempty"`
	FileSysReadOnly    int    `json:"file-sys-read-only,omitempty"`
	HighDiskUse        int    `json:"high-disk-use,omitempty"`
	HighMemoryUse      int    `json:"high-memory-use,omitempty"`
	HighTemp           int    `json:"high-temp,omitempty"`
	LicenseManagement  int    `json:"license-management,omitempty"`
	LowTemp            int    `json:"low-temp,omitempty"`
	PacketDrop         int    `json:"packet-drop,omitempty"`
	Power              int    `json:"power,omitempty"`
	PriDisk            int    `json:"pri-disk,omitempty"`
	Restart            int    `json:"restart,omitempty"`
	SecDisk            int    `json:"sec-disk,omitempty"`
	Shutdown           int    `json:"shutdown,omitempty"`
	SmpResourceEvent   int    `json:"smp-resource-event,omitempty"`
	Start              int    `json:"start,omitempty"`
	SyslogSeverityOne  int    `json:"syslog-severity-one,omitempty"`
	TacacsServerUpDown int    `json:"tacacs-server-up-down,omitempty"`
	UUID               string `json:"uuid,omitempty"`
}

type SnmpServerEnableVcs struct {
	StateChange int    `json:"state-change,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}

type SnmpServerEnableVrrpA struct {
	Active  int    `json:"active,omitempty"`
	All     int    `json:"all,omitempty"`
	Standby int    `json:"standby,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

type SnmpServerEnableBgp struct {
	BgpBackwardTransNotification int    `json:"bgp-backward-trans-notification,omitempty"`
	BgpEstablishedNotification   int    `json:"bgp-established-notification,omitempty"`
	UUID                         string `json:"uuid,omitempty"`
}

type SnmpServerEnableIsis struct {
	IsisAdjacencyChange                  int    `json:"isis-adjacency-change,omitempty"`
	IsisAreaMismatch                     int    `json:"isis-area-mismatch,omitempty"`
	IsisAttemptToExceedMaxSequence       int    `json:"isis-attempt-to-exceed-max-sequence,omitempty"`
	IsisAuthenticationFailure            int    `json:"isis-authentication-failure,omitempty"`
	IsisAuthenticationTypeFailure        int    `json:"isis-authentication-type-failure,omitempty"`
	IsisCorruptedLSPDetected             int    `json:"isis-corrupted-ldp-detected,omitempty"`
	IsisDatabaseOverload                 int    `json:"isis-database-overload,omitempty"`
	IsisIDLenMismatch                    int    `json:"isis-id-len-mismatch,omitempty"`
	IsisLSPTooLargeToPropagate           int    `json:"isis-lsp-too-large-to-propagate,omitempty"`
	IsisManualAddressDrops               int    `json:"isis-manual-address-drops,omitempty"`
	IsisMaxAreaAddressesMismatch         int    `json:"isis-max-area-addresses-mismatch,omitempty"`
	IsisOriginatingLSPBufferSizeMismatch int    `json:"isis-originating-lsp-buffer-size-mismatch,omitempty"`
	IsisOwnLSPPurge                      int    `json:"isis-own-lsp-purge,omitempty"`
	IsisProtocolsSupportedMismatch       int    `json:"isis-protocols-supported-mismatch,omitempty"`
	IsisRejectedAdjacency                int    `json:"isis-rejected-adjacency,omitempty"`
	IsisSequenceNumberSkip               int    `json:"isis-sequence-number-Skip,omitempty"`
	IsisVersionSkew                      int    `json:"isis-version-skew,omitempty"`
	UUID                                 string `json:"uuid,omitempty"`
}

type SnmpServerEnableOspf struct {
	OspfIfAuthFailure           int    `json:"ospf-if-auth-failure,omitempty"`
	OspfIfConfigError           int    `json:"ospf-if-config-error,omitempty"`
	OspfIfRxBadPacket           int    `json:"ospf-if-rx-bad-packet,omitempty"`
	OspfIfStateChange           int    `json:"ospf-if-state-change,omitempty"`
	OspfLsdbApproachingOverflow int    `json:"ospf-lsdb-approaching-overflow,omitempty"`
	OspfLsdbOverflow            int    `json:"ospf-lsdb-overflow,omitempty"`
	OspfMaxAgeLsa               int    `json:"ospf-max-age-lsa,omitempty"`
	OspfNbrStateChange          int    `json:"ospf-nbr-state-change,omitempty"`
	OspfOriginateLsa            int    `json:"ospf-originate-lsa,omitempty"`
	OspfTxRetransmit            int    `json:"ospf-tx-retransmit,omitempty"`
	OspfVirtIfAuthFailure       int    `json:"ospf-virt-if-auth-failure,omitempty"`
	OspfVirtIfConfigError       int    `json:"ospf-virt-if-config-error,omitempty"`
	OspfVirtIfRxBadPacket       int    `json:"ospf-virt-if-rx-bad-packet,omitempty"`
	OspfVirtIfStateChange       int    `json:"ospf-virt-if-state-change,omitempty"`
	OspfVirtIfTxRetransmit      int    `json:"ospf-virt-if-tx-retransmit,omitempty"`
	OspfVirtNbrStateChange      int    `json:"ospf-virt-nbr-state-change,omitempty"`
	UUID                        string `json:"uuid,omitempty"`
}

func PostSnmpServerEnableTraps(id string, inst SnmpServerEnableTraps, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTraps")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTraps
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			check_api_status("PostSnmpServerEnableTraps", data)

		}
	}

}

func GetSnmpServerEnableTraps(id string, host string) (*SnmpServerEnableTraps, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTraps")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTraps
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			check_api_status("GetSnmpServerEnableTraps", data)
			return &m, nil
		}
	}

}
