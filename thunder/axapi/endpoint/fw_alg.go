package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type FwAlg struct {
	Inst struct {
		Dns FwAlgDns417 `json:"dns"`

		Esp FwAlgEsp418 `json:"esp"`

		Ftp FwAlgFtp420 `json:"ftp"`

		Icmp FwAlgIcmp422 `json:"icmp"`

		Pptp FwAlgPptp423 `json:"pptp"`

		Rtsp FwAlgRtsp425 `json:"rtsp"`

		Sctp FwAlgSctp427 `json:"sctp"`

		Sip FwAlgSip428 `json:"sip"`

		Tftp FwAlgTftp430 `json:"tftp"`

		Uuid string `json:"uuid"`
	} `json:"alg"`
}

type FwAlgDns417 struct {
	DefaultPortDisable string `json:"default-port-disable"`
	Uuid               string `json:"uuid"`
}

type FwAlgEsp418 struct {
	DefaultPortDisable string                      `json:"default-port-disable"`
	Uuid               string                      `json:"uuid"`
	SamplingEnable     []FwAlgEspSamplingEnable419 `json:"sampling-enable"`
}

type FwAlgEspSamplingEnable419 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgFtp420 struct {
	DefaultPortDisable string                      `json:"default-port-disable"`
	Uuid               string                      `json:"uuid"`
	SamplingEnable     []FwAlgFtpSamplingEnable421 `json:"sampling-enable"`
}

type FwAlgFtpSamplingEnable421 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgIcmp422 struct {
	Disable string `json:"disable"`
	Uuid    string `json:"uuid"`
}

type FwAlgPptp423 struct {
	DefaultPortDisable string                       `json:"default-port-disable"`
	Uuid               string                       `json:"uuid"`
	SamplingEnable     []FwAlgPptpSamplingEnable424 `json:"sampling-enable"`
}

type FwAlgPptpSamplingEnable424 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgRtsp425 struct {
	DefaultPortDisable string                       `json:"default-port-disable"`
	Uuid               string                       `json:"uuid"`
	SamplingEnable     []FwAlgRtspSamplingEnable426 `json:"sampling-enable"`
}

type FwAlgRtspSamplingEnable426 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgSctp427 struct {
	Action string `json:"action" dval:"enable"`
	Uuid   string `json:"uuid"`
}

type FwAlgSip428 struct {
	DefaultPortDisable string                      `json:"default-port-disable"`
	Uuid               string                      `json:"uuid"`
	SamplingEnable     []FwAlgSipSamplingEnable429 `json:"sampling-enable"`
}

type FwAlgSipSamplingEnable429 struct {
	Counters1 string `json:"counters1"`
}

type FwAlgTftp430 struct {
	DefaultPortDisable string                       `json:"default-port-disable"`
	Uuid               string                       `json:"uuid"`
	SamplingEnable     []FwAlgTftpSamplingEnable431 `json:"sampling-enable"`
}

type FwAlgTftpSamplingEnable431 struct {
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
