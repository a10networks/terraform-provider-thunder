package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SflowPollingMgmtSvcAcl struct {
	Inst struct {
		Toggle string `json:"toggle" dval:"disable"`

		Uuid string `json:"uuid"`
	} `json:"mgmt-svc-acl"`
}

func (p *SflowPollingMgmtSvcAcl) GetId() string {
	return "1"
}

func (p *SflowPollingMgmtSvcAcl) getPath() string {
	return "sflow/polling/mgmt-svc-acl"
}

func (p *SflowPollingMgmtSvcAcl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPollingMgmtSvcAcl::Post")
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

func (p *SflowPollingMgmtSvcAcl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPollingMgmtSvcAcl::Get")
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
func (p *SflowPollingMgmtSvcAcl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPollingMgmtSvcAcl::Put")
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

func (p *SflowPollingMgmtSvcAcl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SflowPollingMgmtSvcAcl::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
