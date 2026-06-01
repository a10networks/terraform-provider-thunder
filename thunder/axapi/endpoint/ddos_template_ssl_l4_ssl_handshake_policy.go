package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosTemplateSslL4SslHandshakePolicy struct {
	Inst struct {
		Action string `json:"action"`

		CipherSuitesLimit int `json:"cipher-suites-limit"`

		ClientExtensionsLimit int `json:"client-extensions-limit"`

		ClienthelloToAppdataTimeout int `json:"clienthello-to-appdata-timeout"`

		FinishedToAppdataTimeout int `json:"finished-to-appdata-timeout"`

		SrcHandshakingConnLimit int `json:"src-handshaking-conn-limit"`

		SslHandshakePolicyActionListName string `json:"ssl-handshake-policy-action-list-name"`

		Uuid string `json:"uuid"`

		SslL4TmplName string
	} `json:"ssl-handshake-policy"`
}

func (p *DdosTemplateSslL4SslHandshakePolicy) GetId() string {
	return "1"
}

func (p *DdosTemplateSslL4SslHandshakePolicy) getPath() string {
	return "ddos/template/ssl-l4/" + p.Inst.SslL4TmplName + "/ssl-handshake-policy"
}

func (p *DdosTemplateSslL4SslHandshakePolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateSslL4SslHandshakePolicy::Post")
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

func (p *DdosTemplateSslL4SslHandshakePolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateSslL4SslHandshakePolicy::Get")
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
func (p *DdosTemplateSslL4SslHandshakePolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateSslL4SslHandshakePolicy::Put")
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

func (p *DdosTemplateSslL4SslHandshakePolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosTemplateSslL4SslHandshakePolicy::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
