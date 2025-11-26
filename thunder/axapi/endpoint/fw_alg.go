package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type FwAlg struct {
	Inst struct {
		Dns FwAlgDns426 `json:"dns"`

		Esp FwAlgEsp427 `json:"esp"`

		Ftp FwAlgFtp429 `json:"ftp"`

		Icmp FwAlgIcmp431 `json:"icmp"`

		Pptp FwAlgPptp432 `json:"pptp"`

		Rtsp FwAlgRtsp434 `json:"rtsp"`

		Sctp FwAlgSctp436 `json:"sctp"`

		Sip FwAlgSip437 `json:"sip"`

		Tftp FwAlgTftp439 `json:"tftp"`

		Uuid string `json:"uuid"`
	} `json:"alg"`
}

type FwAlgDns426 struct {
	DefaultPortDisable string `json:"default-port-disable"`
	Uuid               string `json:"uuid"`
}

type FwAlgEsp427 struct {
	DefaultPortDisable string                      `json:"default-port-disable"`
	Uuid               string                      `json:"uuid"`
	SamplingEnable     []FwAlgEspSamplingEnable428 `json:"sampling-enable"`
}

type FwAlgEspSamplingEnable428 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgFtp429 struct {
	DefaultPortDisable string                      `json:"default-port-disable"`
	Uuid               string                      `json:"uuid"`
	SamplingEnable     []FwAlgFtpSamplingEnable430 `json:"sampling-enable"`
}

type FwAlgFtpSamplingEnable430 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgIcmp431 struct {
	Disable string `json:"disable"`
	Uuid    string `json:"uuid"`
}

type FwAlgPptp432 struct {
	DefaultPortDisable string                       `json:"default-port-disable"`
	Uuid               string                       `json:"uuid"`
	SamplingEnable     []FwAlgPptpSamplingEnable433 `json:"sampling-enable"`
}

type FwAlgPptpSamplingEnable433 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgRtsp434 struct {
	DefaultPortDisable string                       `json:"default-port-disable"`
	Uuid               string                       `json:"uuid"`
	SamplingEnable     []FwAlgRtspSamplingEnable435 `json:"sampling-enable"`
}

type FwAlgRtspSamplingEnable435 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgSctp436 struct {
	Action string `json:"action" dval:"enable"`
	Uuid   string `json:"uuid"`
}

type FwAlgSip437 struct {
	DefaultPortDisable string                      `json:"default-port-disable"`
	Uuid               string                      `json:"uuid"`
	SamplingEnable     []FwAlgSipSamplingEnable438 `json:"sampling-enable"`
}

type FwAlgSipSamplingEnable438 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgTftp439 struct {
	DefaultPortDisable string                       `json:"default-port-disable"`
	Uuid               string                       `json:"uuid"`
	SamplingEnable     []FwAlgTftpSamplingEnable440 `json:"sampling-enable"`
}

type FwAlgTftpSamplingEnable440 struct {
	Counters1 string `json:"counters1"`
}

func (p *FwAlg) GetId() string {
	return "1"
}

func (p *FwAlg) getPath() string {
	return "fw/alg"
}

func (p *FwAlg) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwAlg::Post")
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

func (p *FwAlg) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwAlg::Get")
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
func (p *FwAlg) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FwAlg::Put")
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

func (p *FwAlg) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FwAlg::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
