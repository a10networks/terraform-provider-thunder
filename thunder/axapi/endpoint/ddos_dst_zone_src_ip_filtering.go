package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDstZoneSrcIpFiltering struct {
	Inst struct {
		Name string `json:"name"`

		PerClassListHitTracking int `json:"per-class-list-hit-tracking"`

		PerServiceTracking int `json:"per-service-tracking"`

		SamplingEnable []DdosDstZoneSrcIpFilteringSamplingEnable `json:"sampling-enable"`

		Uuid string `json:"uuid"`

		ZoneName string
	} `json:"src-ip-filtering"`
}

type DdosDstZoneSrcIpFilteringSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *DdosDstZoneSrcIpFiltering) GetId() string {
	return "1"
}

func (p *DdosDstZoneSrcIpFiltering) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/src-ip-filtering"
}

func (p *DdosDstZoneSrcIpFiltering) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZoneSrcIpFiltering::Post")
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

func (p *DdosDstZoneSrcIpFiltering) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZoneSrcIpFiltering::Get")
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
func (p *DdosDstZoneSrcIpFiltering) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZoneSrcIpFiltering::Put")
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

func (p *DdosDstZoneSrcIpFiltering) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZoneSrcIpFiltering::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
