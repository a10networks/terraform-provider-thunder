package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 5_2_1-P4_90
type SlbTemplateCipher struct {
	Inst struct {
		CipherCfg   []SlbTemplateCipherCipherCfg   `json:"cipher-cfg"`
		Cipher13Cfg []SlbTemplateCipherCipher13Cfg `json:"cipher13-cfg"`
		Name        string                         `json:"name"`
		UserTag     string                         `json:"user-tag"`
		Uuid        string                         `json:"uuid"`
	} `json:"cipher"`
}

type SlbTemplateCipherCipherCfg struct {
	CipherSuite string `json:"cipher-suite"`
	Priority    int    `json:"priority" dval:"1"`
}

type SlbTemplateCipherCipher13Cfg struct {
	Cipher13Suite string `json:"cipher13-suite"`
	Priority      int    `json:"priority" dval:"1"`
}

func (p *SlbTemplateCipher) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateCipher) getPath() string {
	return "slb/template/cipher"
}

func (p *SlbTemplateCipher) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateCipher::Post")
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

func (p *SlbTemplateCipher) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateCipher::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *SlbTemplateCipher) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateCipher::Put")
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

func (p *SlbTemplateCipher) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateCipher::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
