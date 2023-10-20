package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P4_81
type VrrpACommon struct {
	Inst struct {
		Action                     string                        `json:"action" dval:"disable"`
		ArpRetry                   int                           `json:"arp-retry" dval:"4"`
		DeadTimer                  int                           `json:"dead-timer" dval:"5"`
		DeviceId                   int                           `json:"device-id"`
		DisableDefaultVrid         int                           `json:"disable-default-vrid"`
		EnableSyncSessionSeqNumber int                           `json:"enable-sync-session-seq-number"`
		ForwardL4PacketOnStandby   int                           `json:"forward-l4-packet-on-standby"`
		GetReadyTime               int                           `json:"get-ready-time" dval:"60"`
		HelloInterval              int                           `json:"hello-interval" dval:"2"`
		HostidAppendToVrid         VrrpACommonHostidAppendToVrid `json:"hostid-append-to-vrid"`
		InlineModeCfg              VrrpACommonInlineModeCfg      `json:"inline-mode-cfg"`
		PreemptionDelay            int                           `json:"preemption-delay" dval:"60"`
		RestartTime                int                           `json:"restart-time" dval:"20"`
		SetId                      int                           `json:"set-id"`
		TrackEventDelay            int                           `json:"track-event-delay" dval:"30"`
		Uuid                       string                        `json:"uuid"`
	} `json:"common"`
}

type VrrpACommonHostidAppendToVrid struct {
	HostidAppendToVridDefault int `json:"hostid-append-to-vrid-default"`
	HostidAppendToVridValue   int `json:"hostid-append-to-vrid-value"`
}

type VrrpACommonInlineModeCfg struct {
	InlineMode     int `json:"inline-mode"`
	PreferredPort  int `json:"preferred-port"`
	PreferredTrunk int `json:"preferred-trunk"`
}

func (p *VrrpACommon) GetId() string {
	return "1"
}

func (p *VrrpACommon) getPath() string {
	return "vrrp-a/common"
}

func (p *VrrpACommon) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VrrpACommon::Post")
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

func (p *VrrpACommon) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VrrpACommon::Get")
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

func (p *VrrpACommon) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VrrpACommon::Put")
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

func (p *VrrpACommon) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VrrpACommon::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
