package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type SystemForcedGroupSpeed struct {
	Inst struct {
		Eth01_to_04 int `json:"eth01_to_04"`

		Eth05_to_08 int `json:"eth05_to_08"`

		Eth09_to_12 int `json:"eth09_to_12"`

		Eth13_to_16 int `json:"eth13_to_16"`

		Eth17_to_20 int `json:"eth17_to_20"`

		Eth21_to_24 int `json:"eth21_to_24"`

		Speed string `json:"speed" dval:"10G"`

		Uuid string `json:"uuid"`
	} `json:"forced-group-speed"`
}

func (p *SystemForcedGroupSpeed) GetId() string {
	return strconv.Itoa(p.Inst.Eth01_to_04) + "+" + strconv.Itoa(p.Inst.Eth05_to_08) + "+" + strconv.Itoa(p.Inst.Eth09_to_12) + "+" + strconv.Itoa(p.Inst.Eth13_to_16) + "+" + strconv.Itoa(p.Inst.Eth17_to_20) + "+" + strconv.Itoa(p.Inst.Eth21_to_24)
}

func (p *SystemForcedGroupSpeed) getPath() string {
	return "system/forced-group-speed"
}

func (p *SystemForcedGroupSpeed) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemForcedGroupSpeed::Post")
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

func (p *SystemForcedGroupSpeed) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemForcedGroupSpeed::Get")
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
func (p *SystemForcedGroupSpeed) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemForcedGroupSpeed::Put")
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

func (p *SystemForcedGroupSpeed) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemForcedGroupSpeed::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
