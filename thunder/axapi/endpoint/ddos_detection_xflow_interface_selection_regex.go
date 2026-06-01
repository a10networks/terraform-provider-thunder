package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDetectionXflowInterfaceSelectionRegex struct {
	Inst struct {
		RuleList []DdosDetectionXflowInterfaceSelectionRegexRuleList `json:"rule-list"`

		Uuid string `json:"uuid"`

		Type string
	} `json:"regex"`
}

type DdosDetectionXflowInterfaceSelectionRegexRuleList struct {
	SingleRegex string `json:"single-regex"`
}

func (p *DdosDetectionXflowInterfaceSelectionRegex) GetId() string {
	return "1"
}

func (p *DdosDetectionXflowInterfaceSelectionRegex) getPath() string {
	return "ddos/detection/xflow-interface-selection/" + p.Inst.Type + "/regex"
}

func (p *DdosDetectionXflowInterfaceSelectionRegex) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionXflowInterfaceSelectionRegex::Post")
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

func (p *DdosDetectionXflowInterfaceSelectionRegex) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionXflowInterfaceSelectionRegex::Get")
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
func (p *DdosDetectionXflowInterfaceSelectionRegex) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionXflowInterfaceSelectionRegex::Put")
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

func (p *DdosDetectionXflowInterfaceSelectionRegex) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionXflowInterfaceSelectionRegex::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
