package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SflowPollingDdos struct {
	Inst struct {
		AutoDiscoveredSni int `json:"auto-discovered-sni"`

		DnsCacheZoneStats int `json:"dns-cache-zone-stats"`

		DynEntryStats int `json:"dyn-entry-stats"`

		EnableAnomalyStats int `json:"enable-anomaly-stats"`

		Toggle string `json:"toggle" dval:"disable"`

		Uuid string `json:"uuid"`

		ZoneSession int `json:"zone-session"`
	} `json:"ddos"`
}

func (p *SflowPollingDdos) GetId() string {
	return "1"
}

func (p *SflowPollingDdos) getPath() string {
	return "sflow/polling/ddos"
}

func (p *SflowPollingDdos) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPollingDdos::Post")
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

func (p *SflowPollingDdos) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPollingDdos::Get")
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
func (p *SflowPollingDdos) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPollingDdos::Put")
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

func (p *SflowPollingDdos) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPollingDdos::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
