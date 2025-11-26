package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosZoneTemplateQuicActionOnInitial struct {
	Inst struct {
		RetryTokenAuthenticationDynamic string `json:"retry-token-authentication-dynamic"`

		RetryTokenAuthenticationStatic string `json:"retry-token-authentication-static"`

		RetryUnauthenticatedInitialDecrypt int `json:"retry-unauthenticated-initial-decrypt"`

		ScidLength int `json:"scid-length" dval:"20"`

		UnauthShortHdrAction string `json:"unauth-short-hdr-action"`

		Uuid string `json:"uuid"`

		QuicTmplName string
	} `json:"action-on-initial"`
}

func (p *DdosZoneTemplateQuicActionOnInitial) GetId() string {
	return "1"
}

func (p *DdosZoneTemplateQuicActionOnInitial) getPath() string {
	return "ddos/zone-template/quic/" + p.Inst.QuicTmplName + "/action-on-initial"
}

func (p *DdosZoneTemplateQuicActionOnInitial) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateQuicActionOnInitial::Post")
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

func (p *DdosZoneTemplateQuicActionOnInitial) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateQuicActionOnInitial::Get")
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
func (p *DdosZoneTemplateQuicActionOnInitial) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateQuicActionOnInitial::Put")
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

func (p *DdosZoneTemplateQuicActionOnInitial) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosZoneTemplateQuicActionOnInitial::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
