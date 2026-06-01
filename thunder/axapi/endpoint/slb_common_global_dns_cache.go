package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbCommonGlobalDnsCache struct {
	Inst struct {
		ClassList SlbCommonGlobalDnsCacheClassList1500 `json:"class-list"`

		Uuid string `json:"uuid"`
	} `json:"global-dns-cache"`
}

type SlbCommonGlobalDnsCacheClassList1500 struct {
	Name    string                                        `json:"name"`
	Uuid    string                                        `json:"uuid"`
	LidList []SlbCommonGlobalDnsCacheClassListLidList1501 `json:"lid-list"`
}

type SlbCommonGlobalDnsCacheClassListLidList1501 struct {
	Lidnum          int                                            `json:"lidnum"`
	ConnRateLimit   int                                            `json:"conn-rate-limit"`
	Per             int                                            `json:"per"`
	OverLimitAction string                                         `json:"over-limit-action" dval:"drop"`
	Lockout         int                                            `json:"lockout"`
	Log             int                                            `json:"log"`
	LogInterval     int                                            `json:"log-interval"`
	Dns             SlbCommonGlobalDnsCacheClassListLidListDns1502 `json:"dns"`
	Uuid            string                                         `json:"uuid"`
	UserTag         string                                         `json:"user-tag"`
}

type SlbCommonGlobalDnsCacheClassListLidListDns1502 struct {
	CacheAction            string `json:"cache-action" dval:"cache-enable"`
	Ttl                    int    `json:"ttl" dval:"300"`
	Weight                 int    `json:"weight" dval:"1"`
	HonorServerResponseTtl int    `json:"honor-server-response-ttl"`
}

func (p *SlbCommonGlobalDnsCache) GetId() string {
	return "1"
}

func (p *SlbCommonGlobalDnsCache) getPath() string {
	return "slb/common/global-dns-cache"
}

func (p *SlbCommonGlobalDnsCache) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonGlobalDnsCache::Post")
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

func (p *SlbCommonGlobalDnsCache) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonGlobalDnsCache::Get")
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
func (p *SlbCommonGlobalDnsCache) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonGlobalDnsCache::Put")
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

func (p *SlbCommonGlobalDnsCache) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommonGlobalDnsCache::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
