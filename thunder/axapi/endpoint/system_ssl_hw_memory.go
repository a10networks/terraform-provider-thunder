package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SystemSslHwMemory struct {
	Inst struct {
		MemBlockCfg []SystemSslHwMemoryMemBlockCfg `json:"mem-block-cfg"`

		Uuid string `json:"uuid"`
	} `json:"ssl-hw-memory"`
}

type SystemSslHwMemoryMemBlockCfg struct {
	MemBlock string `json:"mem-block"`
	Size     int    `json:"size"`
}

func (p *SystemSslHwMemory) GetId() string {
	return "1"
}

func (p *SystemSslHwMemory) getPath() string {
	return "system/ssl-hw-memory"
}

func (p *SystemSslHwMemory) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemSslHwMemory::Post")
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

func (p *SystemSslHwMemory) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemSslHwMemory::Get")
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
func (p *SystemSslHwMemory) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemSslHwMemory::Put")
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

func (p *SystemSslHwMemory) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemSslHwMemory::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
