package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 5_2_1-P4_90
type SlbTemplateDns struct {
	Inst struct {
		AddPaddingToClient        string                               `json:"add-padding-to-client"`
		CacheRecordServingPolicy  string                               `json:"cache-record-serving-policy"`
		ClassList                 SlbTemplateDnsClassList              `json:"class-list"`
		DefaultPolicy             string                               `json:"default-policy" dval:"nocache"`
		DisableDnsTemplate        int                                  `json:"disable-dns-template"`
		DisableRaCachedResp       int                                  `json:"disable-ra-cached-resp"`
		DisableRpzAttachSoa       int                                  `json:"disable-rpz-attach-soa"`
		DnsLogging                string                               `json:"dns-logging"`
		DnssecServiceGroup        string                               `json:"dnssec-service-group"`
		Drop                      int                                  `json:"drop"`
		EnableCacheSharing        int                                  `json:"enable-cache-sharing"`
		Forward                   string                               `json:"forward"`
		LocalDnsResolution        SlbTemplateDnsLocalDnsResolution     `json:"local-dns-resolution"`
		MaxCacheEntrySize         int                                  `json:"max-cache-entry-size" dval:"1024"`
		MaxCacheSize              int                                  `json:"max-cache-size"`
		MaxQueryLength            int                                  `json:"max-query-length"`
		Name                      string                               `json:"name"`
		Period                    int                                  `json:"period"`
		QueryClassFilter          SlbTemplateDnsQueryClassFilter       `json:"query-class-filter"`
		QueryIdSwitch             int                                  `json:"query-id-switch"`
		QueryTypeFilter           SlbTemplateDnsQueryTypeFilter        `json:"query-type-filter"`
		RecursiveDnsResolution    SlbTemplateDnsRecursiveDnsResolution `json:"recursive-dns-resolution"`
		RedirectToTcpPort         int                                  `json:"redirect-to-tcp-port"`
		RemoveAaFlag              int                                  `json:"remove-aa-flag"`
		RemoveEdnsCsubnetToServer int                                  `json:"remove-edns-csubnet-to-server"`
		RemovePaddingToServer     int                                  `json:"remove-padding-to-server"`
		ResponseRateLimiting      SlbTemplateDnsResponseRateLimiting   `json:"response-rate-limiting"`
		RpzList                   []SlbTemplateDnsRpzList              `json:"rpz-list"`
		UdpRetransmit             SlbTemplateDnsUdpRetransmit          `json:"udp-retransmit"`
		UserTag                   string                               `json:"user-tag"`
		Uuid                      string                               `json:"uuid"`
	} `json:"dns"`
}

type SlbTemplateDnsClassList struct {
	Name    string                           `json:"name"`
	Uuid    string                           `json:"uuid"`
	LidList []SlbTemplateDnsClassListLidList `json:"lid-list"`
}

type SlbTemplateDnsClassListLidList struct {
	Lidnum          int                               `json:"lidnum"`
	ConnRateLimit   int                               `json:"conn-rate-limit"`
	Per             int                               `json:"per"`
	OverLimitAction int                               `json:"over-limit-action"`
	ActionValue     string                            `json:"action-value"`
	Lockout         int                               `json:"lockout"`
	Log             int                               `json:"log"`
	LogInterval     int                               `json:"log-interval"`
	Dns             SlbTemplateDnsClassListLidListDns `json:"dns"`
	Uuid            string                            `json:"uuid"`
	UserTag         string                            `json:"user-tag"`
}

type SlbTemplateDnsClassListLidListDns struct {
	CacheAction            string `json:"cache-action" dval:"cache-disable"`
	Ttl                    int    `json:"ttl"`
	Weight                 int    `json:"weight"`
	HonorServerResponseTtl int    `json:"honor-server-response-ttl"`
}

type SlbTemplateDnsLocalDnsResolution struct {
	HostListCfg      []SlbTemplateDnsLocalDnsResolutionHostListCfg      `json:"host-list-cfg"`
	LocalResolverCfg []SlbTemplateDnsLocalDnsResolutionLocalResolverCfg `json:"local-resolver-cfg"`
	Uuid             string                                             `json:"uuid"`
}

type SlbTemplateDnsLocalDnsResolutionHostListCfg struct {
	Hostnames string `json:"hostnames"`
}

type SlbTemplateDnsLocalDnsResolutionLocalResolverCfg struct {
	LocalResolver string `json:"local-resolver"`
}

type SlbTemplateDnsQueryClassFilter struct {
	QueryClassAction string                                     `json:"query-class-action"`
	QueryClass       []SlbTemplateDnsQueryClassFilterQueryClass `json:"query-class"`
	Uuid             string                                     `json:"uuid"`
}

type SlbTemplateDnsQueryClassFilterQueryClass struct {
	StrQueryClass string `json:"str-query-class"`
	NumQueryClass int    `json:"num-query-class"`
}

type SlbTemplateDnsQueryTypeFilter struct {
	QueryTypeAction string                                   `json:"query-type-action"`
	QueryType       []SlbTemplateDnsQueryTypeFilterQueryType `json:"query-type"`
	Uuid            string                                   `json:"uuid"`
}

type SlbTemplateDnsQueryTypeFilterQueryType struct {
	StrQueryType string `json:"str-query-type"`
	NumQueryType int    `json:"num-query-type"`
}

type SlbTemplateDnsRecursiveDnsResolution struct {
	HostListCfg             []SlbTemplateDnsRecursiveDnsResolutionHostListCfg `json:"host-list-cfg"`
	NsCacheLookup           string                                            `json:"ns-cache-lookup" dval:"enabled"`
	UseServiceGroupResponse string                                            `json:"use-service-group-response" dval:"enabled"`
	Ipv4NatPool             string                                            `json:"ipv4-nat-pool"`
	Ipv6NatPool             string                                            `json:"ipv6-nat-pool"`
	RetriesPerLevel         int                                               `json:"retries-per-level" dval:"6"`
	FullResponse            int                                               `json:"full-response"`
	MaxTrials               int                                               `json:"max-trials"`
	UseClientQid            int                                               `json:"use-client-qid"`
	Uuid                    string                                            `json:"uuid"`
}

type SlbTemplateDnsRecursiveDnsResolutionHostListCfg struct {
	Hostnames string `json:"hostnames"`
}

type SlbTemplateDnsResponseRateLimiting struct {
	ResponseRate       int                                                  `json:"response-rate" dval:"5"`
	FilterResponseRate int                                                  `json:"filter-response-rate" dval:"10"`
	SlipRate           int                                                  `json:"slip-rate"`
	Window             int                                                  `json:"window" dval:"1"`
	EnableLog          int                                                  `json:"enable-log"`
	Action             string                                               `json:"action" dval:"rate-limit"`
	Uuid               string                                               `json:"uuid"`
	RrlClassListList   []SlbTemplateDnsResponseRateLimitingRrlClassListList `json:"rrl-class-list-list"`
}

type SlbTemplateDnsResponseRateLimitingRrlClassListList struct {
	Name    string                                                      `json:"name"`
	Uuid    string                                                      `json:"uuid"`
	UserTag string                                                      `json:"user-tag"`
	LidList []SlbTemplateDnsResponseRateLimitingRrlClassListListLidList `json:"lid-list"`
}

type SlbTemplateDnsResponseRateLimitingRrlClassListListLidList struct {
	Lidnum          int    `json:"lidnum"`
	LidResponseRate int    `json:"lid-response-rate" dval:"5"`
	LidSlipRate     int    `json:"lid-slip-rate"`
	LidWindow       int    `json:"lid-window" dval:"1"`
	LidEnableLog    int    `json:"lid-enable-log"`
	LidAction       string `json:"lid-action" dval:"rate-limit"`
	Uuid            string `json:"uuid"`
	UserTag         string `json:"user-tag"`
}

type SlbTemplateDnsRpzList struct {
	SeqId   int                          `json:"seq-id"`
	Name    string                       `json:"name"`
	Uuid    string                       `json:"uuid"`
	UserTag string                       `json:"user-tag"`
	Logging SlbTemplateDnsRpzListLogging `json:"logging"`
}

type SlbTemplateDnsRpzListLogging struct {
	Enable    int                                     `json:"enable"`
	RpzAction []SlbTemplateDnsRpzListLoggingRpzAction `json:"rpz-action"`
	Uuid      string                                  `json:"uuid"`
}

type SlbTemplateDnsRpzListLoggingRpzAction struct {
	StrRpzAction string `json:"str-rpz-action"`
}

type SlbTemplateDnsUdpRetransmit struct {
	RetryInterval int    `json:"retry-interval" dval:"10"`
	MaxTrials     int    `json:"max-trials" dval:"3"`
	Uuid          string `json:"uuid"`
}

func (p *SlbTemplateDns) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateDns) getPath() string {
	return "slb/template/dns"
}

func (p *SlbTemplateDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDns::Post")
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

func (p *SlbTemplateDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDns::Get")
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

func (p *SlbTemplateDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDns::Put")
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

func (p *SlbTemplateDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateDns::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
