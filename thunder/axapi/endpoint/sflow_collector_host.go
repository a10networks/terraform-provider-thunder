package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
	"strconv"
)

// based on ACOS 7_0_2-102
type SflowCollectorHost struct {
	Inst struct {
		CustomizedSetting SflowCollectorHostCustomizedSetting1490 `json:"customized-setting"`

		Name string `json:"name"`

		Port int `json:"port"`

		UseMgmtPort int `json:"use-mgmt-port"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"host"`
}

type SflowCollectorHostCustomizedSetting1490 struct {
	ExportEnable          string `json:"export-enable"`
	PacketSampling        int    `json:"packet-sampling"`
	CounterPolling        int    `json:"counter-polling"`
	A10ProprietaryPolling int    `json:"a10-proprietary-polling"`
	EventNotification     int    `json:"event-notification"`
	Uuid                  string `json:"uuid"`
}

func (p *SflowCollectorHost) GetId() string {
	return url.QueryEscape(p.Inst.Name) + "+" + strconv.Itoa(p.Inst.Port)
}

func (p *SflowCollectorHost) getPath() string {
	return "sflow/collector/host"
}

func (p *SflowCollectorHost) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SflowCollectorHost::Post")
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

func (p *SflowCollectorHost) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SflowCollectorHost::Get")
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
func (p *SflowCollectorHost) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SflowCollectorHost::Put")
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

func (p *SflowCollectorHost) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SflowCollectorHost::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
