package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type VisibilityMonitoredEntity struct {
	Inst struct {
		Detail VisibilityMonitoredEntityDetail2052 `json:"detail"`

		MonTopk VisibilityMonitoredEntityMonTopk2054 `json:"mon-topk"`

		Secondary VisibilityMonitoredEntitySecondary2056 `json:"secondary"`

		Sessions VisibilityMonitoredEntitySessions2059 `json:"sessions"`

		Uuid string `json:"uuid"`
	} `json:"monitored-entity"`
}

type VisibilityMonitoredEntityDetail2052 struct {
	Uuid  string                                   `json:"uuid"`
	Debug VisibilityMonitoredEntityDetailDebug2053 `json:"debug"`
}

type VisibilityMonitoredEntityDetailDebug2053 struct {
	Uuid string `json:"uuid"`
}

type VisibilityMonitoredEntityMonTopk2054 struct {
	Uuid    string                                      `json:"uuid"`
	Sources VisibilityMonitoredEntityMonTopkSources2055 `json:"sources"`
}

type VisibilityMonitoredEntityMonTopkSources2055 struct {
	Uuid string `json:"uuid"`
}

type VisibilityMonitoredEntitySecondary2056 struct {
	MonTopk VisibilityMonitoredEntitySecondaryMonTopk2057 `json:"mon-topk"`
}

type VisibilityMonitoredEntitySecondaryMonTopk2057 struct {
	Uuid    string                                               `json:"uuid"`
	Sources VisibilityMonitoredEntitySecondaryMonTopkSources2058 `json:"sources"`
}

type VisibilityMonitoredEntitySecondaryMonTopkSources2058 struct {
	Uuid string `json:"uuid"`
}

type VisibilityMonitoredEntitySessions2059 struct {
	Uuid string `json:"uuid"`
}

func (p *VisibilityMonitoredEntity) GetId() string {
	return "1"
}

func (p *VisibilityMonitoredEntity) getPath() string {
	return "visibility/monitored-entity"
}

func (p *VisibilityMonitoredEntity) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityMonitoredEntity::Post")
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

func (p *VisibilityMonitoredEntity) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityMonitoredEntity::Get")
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
func (p *VisibilityMonitoredEntity) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityMonitoredEntity::Put")
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

func (p *VisibilityMonitoredEntity) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VisibilityMonitoredEntity::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
