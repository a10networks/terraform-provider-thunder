package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDetectionXflowInterfaceSelection struct {
	Inst struct {
		Regex DdosDetectionXflowInterfaceSelectionRegex155 `json:"regex"`

		Type string `json:"type"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"xflow-interface-selection"`
}

type DdosDetectionXflowInterfaceSelectionRegex155 struct {
	RuleList []DdosDetectionXflowInterfaceSelectionRegexRuleList156 `json:"rule-list"`
	Uuid     string                                                 `json:"uuid"`
}

type DdosDetectionXflowInterfaceSelectionRegexRuleList156 struct {
	SingleRegex string `json:"single-regex"`
}

func (p *DdosDetectionXflowInterfaceSelection) GetId() string {
	return p.Inst.Type
}

func (p *DdosDetectionXflowInterfaceSelection) getPath() string {
	return "ddos/detection/xflow-interface-selection"
}

func (p *DdosDetectionXflowInterfaceSelection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionXflowInterfaceSelection::Post")
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

func (p *DdosDetectionXflowInterfaceSelection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionXflowInterfaceSelection::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *DdosDetectionXflowInterfaceSelection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionXflowInterfaceSelection::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *DdosDetectionXflowInterfaceSelection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionXflowInterfaceSelection::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
