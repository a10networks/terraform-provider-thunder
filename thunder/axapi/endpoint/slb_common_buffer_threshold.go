package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// AxAPI cannot handle "operation" and "configuration" fields at the same time. Break `/axapi/v3/slb/common` to 2 endpoints
// based on ACOS 5_2_1-P4_70
type SlbCommonBufferThreshold struct {
	Inst struct {
		BuffThresh              int `json:"buff-thresh"`
		BuffThreshHwBuff        int `json:"buff-thresh-hw-buff"`
		BuffThreshRelieveThresh int `json:"buff-thresh-relieve-thresh"`
		BuffThreshSysBuffHigh   int `json:"buff-thresh-sys-buff-high"`
		BuffThreshSysBuffLow    int `json:"buff-thresh-sys-buff-low"`
	} `json:"common"`
}

func (p *SlbCommonBufferThreshold) GetId() string {
	return "1"
}

func (p *SlbCommonBufferThreshold) getPath() string {
	return "slb/common"
}

func (p *SlbCommonBufferThreshold) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonBufferThreshold::Post")
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

func (p *SlbCommonBufferThreshold) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonBufferThreshold::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *SlbCommonBufferThreshold) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonBufferThreshold::Delete")
	headers := axapi.GenRequestHeader(authToken)
	//Do not delete /axapi/v3/slb/common. Other configurations of /slb/common/ will be earsed too.
	obj := SlbCommonBufferThreshold{}
	obj.Inst.BuffThresh = 0
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}
