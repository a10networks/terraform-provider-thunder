package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SflowPolling struct {
	Inst struct {
		A10Proprietary SflowPollingA10Proprietary1493 `json:"a10-proprietary"`

		AcosInfo SflowPollingAcosInfo1494 `json:"acos-info"`

		CpuUsage int `json:"cpu-usage"`

		Ddos SflowPollingDdos1495 `json:"ddos"`

		EthList []SflowPollingEthList `json:"eth-list"`

		EthernetExtList []SflowPollingEthernetExtList `json:"ethernet-ext-list"`

		EthernetList []SflowPollingEthernetList `json:"ethernet-list"`

		Http SflowPollingHttp1496 `json:"http"`

		HttpCounter int `json:"http-counter"`

		MgmtSvcAcl SflowPollingMgmtSvcAcl1497 `json:"mgmt-svc-acl"`

		SystemHealth SflowPollingSystemHealth1498 `json:"system-health"`

		Uuid string `json:"uuid"`

		VeList []SflowPollingVeList `json:"ve-list"`
	} `json:"polling"`
}

type SflowPollingA10Proprietary1493 struct {
	ExportDeprecatedCounters int    `json:"export-deprecated-counters"`
	Uuid                     string `json:"uuid"`
}

type SflowPollingAcosInfo1494 struct {
	Toggle string `json:"toggle" dval:"disable"`
	Uuid   string `json:"uuid"`
}

type SflowPollingDdos1495 struct {
	Toggle             string `json:"toggle" dval:"disable"`
	DnsCacheZoneStats  int    `json:"dns-cache-zone-stats"`
	EnableAnomalyStats int    `json:"enable-anomaly-stats"`
	DynEntryStats      int    `json:"dyn-entry-stats"`
	ZoneSession        int    `json:"zone-session"`
	AutoDiscoveredSni  int    `json:"auto-discovered-sni"`
	Uuid               string `json:"uuid"`
}

type SflowPollingEthList struct {
	EthStart int `json:"eth-start"`
	EthEnd   int `json:"eth-end"`
}

type SflowPollingEthernetExtList struct {
	Start int    `json:"start"`
	Uuid  string `json:"uuid"`
}

type SflowPollingEthernetList struct {
	Start int    `json:"start"`
	Uuid  string `json:"uuid"`
}

type SflowPollingHttp1496 struct {
	Toggle string `json:"toggle" dval:"disable"`
	Uuid   string `json:"uuid"`
}

type SflowPollingMgmtSvcAcl1497 struct {
	Toggle string `json:"toggle" dval:"disable"`
	Uuid   string `json:"uuid"`
}

type SflowPollingSystemHealth1498 struct {
	SystemHealthUsage  string `json:"system-health-usage" dval:"disable"`
	PerControlCpuUsage string `json:"per-control-cpu-usage" dval:"disable"`
	PerDataCpuUsage    string `json:"per-data-cpu-usage" dval:"disable"`
	LicenseStatistics  string `json:"license-statistics" dval:"disable"`
	Uuid               string `json:"uuid"`
}

type SflowPollingVeList struct {
	VeStart int `json:"ve-start"`
	VeEnd   int `json:"ve-end"`
}

func (p *SflowPolling) GetId() string {
	return "1"
}

func (p *SflowPolling) getPath() string {
	return "sflow/polling"
}

func (p *SflowPolling) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPolling::Post")
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

func (p *SflowPolling) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPolling::Get")
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
func (p *SflowPolling) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPolling::Put")
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

func (p *SflowPolling) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPolling::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
