package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type SlbCommonGlobalDnsCacheClassListLid struct {
	Inst struct {
		ConnRateLimit int `json:"conn-rate-limit"`

		Dns SlbCommonGlobalDnsCacheClassListLidDns `json:"dns"`

		Lidnum int `json:"lidnum"`

		Lockout int `json:"lockout"`

		Log int `json:"log"`

		LogInterval int `json:"log-interval"`

		OverLimitAction string `json:"over-limit-action" dval:"drop"`

		Per int `json:"per"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"lid"`
}

type SlbCommonGlobalDnsCacheClassListLidDns struct {
	CacheAction            string `json:"cache-action" dval:"cache-enable"`
	Ttl                    int    `json:"ttl" dval:"300"`
	Weight                 int    `json:"weight" dval:"1"`
	HonorServerResponseTtl int    `json:"honor-server-response-ttl"`
}

func (p *SlbCommonGlobalDnsCacheClassListLid) GetId() string {
	return strconv.Itoa(p.Inst.Lidnum)
}

func (p *SlbCommonGlobalDnsCacheClassListLid) getPath() string {
	return "slb/common/global-dns-cache/class-list/lid"
}

func (p *SlbCommonGlobalDnsCacheClassListLid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonGlobalDnsCacheClassListLid::Post")
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

func (p *SlbCommonGlobalDnsCacheClassListLid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonGlobalDnsCacheClassListLid::Get")
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
func (p *SlbCommonGlobalDnsCacheClassListLid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonGlobalDnsCacheClassListLid::Put")
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

func (p *SlbCommonGlobalDnsCacheClassListLid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonGlobalDnsCacheClassListLid::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
