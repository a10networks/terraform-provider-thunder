package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type SlbTemplateDns struct {
	Inst struct {
		AddPaddingToClient string `json:"add-padding-to-client"`

		CacheHitcountEnable int `json:"cache-hitcount-enable"`

		CacheRecordServingPolicy string `json:"cache-record-serving-policy"`

		CacheTtlAdjustmentEnable int `json:"cache-ttl-adjustment-enable"`

		CategoryLookupBypass string `json:"category-lookup-bypass"`

		CategoryLookupList []SlbTemplateDnsCategoryLookupList `json:"category-lookup-list"`

		CategoryLookupOnlineLookup int `json:"category-lookup-online-lookup"`

		ClassList SlbTemplateDnsClassList1524 `json:"class-list"`

		DefaultPolicy string `json:"default-policy" dval:"nocache"`

		DisableDnsTemplate int `json:"disable-dns-template"`

		DisableRaCachedResp int `json:"disable-ra-cached-resp"`

		DisableRpzAttachSoa int `json:"disable-rpz-attach-soa"`

		DnsCookieCachePolicy string `json:"dns-cookie-cache-policy"`

		DnsLogging string `json:"dns-logging"`

		Dns64 SlbTemplateDnsDns641527 `json:"dns64"`

		DnssecServiceGroup string `json:"dnssec-service-group"`

		Drop int `json:"drop"`

		EnableCacheSharing int `json:"enable-cache-sharing"`

		Forward string `json:"forward"`

		InsertIpv4 int `json:"insert-ipv4"`

		InsertIpv6 int `json:"insert-ipv6"`

		LabelCountFilter SlbTemplateDnsLabelCountFilter1528 `json:"label-count-filter"`

		LabelLengthFilter SlbTemplateDnsLabelLengthFilter1529 `json:"label-length-filter"`

		LocalDnsResolution SlbTemplateDnsLocalDnsResolution1531 `json:"local-dns-resolution"`

		MaxCacheEntrySize int `json:"max-cache-entry-size" dval:"1024"`

		MaxCacheSize int `json:"max-cache-size"`

		MaxQueryLength int `json:"max-query-length"`

		MaxUdpSize int `json:"max-udp-size"`

		Name string `json:"name"`

		NegativeDnsCache SlbTemplateDnsNegativeDnsCache1534 `json:"negative-dns-cache"`

		Period int `json:"period"`

		QpsLogHigh int `json:"qps-log-high"`

		QpsLogLow int `json:"qps-log-low"`

		QpsThresholdLog int `json:"qps-threshold-log"`

		QueryClassFilter SlbTemplateDnsQueryClassFilter1535 `json:"query-class-filter"`

		QueryIdSwitch int `json:"query-id-switch"`

		QueryTypeFilter SlbTemplateDnsQueryTypeFilter1537 `json:"query-type-filter"`

		RecursiveDnsResolution SlbTemplateDnsRecursiveDnsResolution1539 `json:"recursive-dns-resolution"`

		RedirectToTcpPort int `json:"redirect-to-tcp-port"`

		RemoveAaFlag int `json:"remove-aa-flag"`

		RemoveCsubnet int `json:"remove-csubnet"`

		RemovePaddingToServer int `json:"remove-padding-to-server"`

		ResponseRateLimiting SlbTemplateDnsResponseRateLimiting1544 `json:"response-rate-limiting"`

		RpzList []SlbTemplateDnsRpzList `json:"rpz-list"`

		TldFilterLogEnable int `json:"tld-filter-log-enable"`

		TldFilterWhiteList string `json:"tld-filter-white-list"`

		UdpRetransmit SlbTemplateDnsUdpRetransmit1547 `json:"udp-retransmit"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"dns"`
}

type SlbTemplateDnsCategoryLookupList struct {
	CategoryName    string `json:"category-name"`
	Permit          int    `json:"permit"`
	Drop            int    `json:"drop"`
	Respond         int    `json:"respond"`
	RespondNxdomain int    `json:"respond-nxdomain"`
	RespondIpAddr   string `json:"respond-ip-addr"`
	RespondIpv6Addr string `json:"respond-ipv6-addr"`
	RespondCnameStr string `json:"respond-cname-str"`
	ResponseTtl     int    `json:"response-ttl" dval:"300"`
	Uuid            string `json:"uuid"`
}

type SlbTemplateDnsClassList1524 struct {
	Name    string                               `json:"name"`
	Uuid    string                               `json:"uuid"`
	LidList []SlbTemplateDnsClassListLidList1525 `json:"lid-list"`
}

type SlbTemplateDnsClassListLidList1525 struct {
	Lidnum          int                                   `json:"lidnum"`
	ConnRateLimit   int                                   `json:"conn-rate-limit"`
	Per             int                                   `json:"per"`
	OverLimitAction int                                   `json:"over-limit-action"`
	ActionValue     string                                `json:"action-value"`
	Lockout         int                                   `json:"lockout"`
	Log             int                                   `json:"log"`
	LogInterval     int                                   `json:"log-interval"`
	Dns             SlbTemplateDnsClassListLidListDns1526 `json:"dns"`
	Uuid            string                                `json:"uuid"`
	UserTag         string                                `json:"user-tag"`
}

type SlbTemplateDnsClassListLidListDns1526 struct {
	CacheAction            string `json:"cache-action" dval:"cache-disable"`
	Ttl                    int    `json:"ttl"`
	Weight                 int    `json:"weight"`
	HonorServerResponseTtl int    `json:"honor-server-response-ttl"`
}

type SlbTemplateDnsDns641527 struct {
	Enable                int    `json:"enable"`
	Cache                 int    `json:"cache"`
	ChangeQuery           int    `json:"change-query"`
	ParallelQuery         int    `json:"parallel-query"`
	Retry                 int    `json:"retry" dval:"3"`
	SingleResponseDisable int    `json:"single-response-disable"`
	Timeout               int    `json:"timeout" dval:"1"`
	Uuid                  string `json:"uuid"`
}

type SlbTemplateDnsLabelCountFilter1528 struct {
	DropLogEnable          int    `json:"drop-log-enable"`
	LabelCountFilterAction string `json:"label-count-filter-action" dval:"drop"`
	MinFqdnLabelCount      int    `json:"min-fqdn-label-count"`
	MaxFqdnLabelCount      int    `json:"max-fqdn-label-count"`
	Uuid                   string `json:"uuid"`
}

type SlbTemplateDnsLabelLengthFilter1529 struct {
	DropLogEnable           int                                                  `json:"drop-log-enable"`
	LabelLengthFilterAction string                                               `json:"label-length-filter-action" dval:"drop"`
	FqdnLabelLength         []SlbTemplateDnsLabelLengthFilterFqdnLabelLength1530 `json:"fqdn-label-length"`
	Uuid                    string                                               `json:"uuid"`
}

type SlbTemplateDnsLabelLengthFilterFqdnLabelLength1530 struct {
	Length int `json:"length"`
	Suffix int `json:"suffix"`
}

type SlbTemplateDnsLocalDnsResolution1531 struct {
	HostListCfg      []SlbTemplateDnsLocalDnsResolutionHostListCfg1532      `json:"host-list-cfg"`
	LocalResolverCfg []SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1533 `json:"local-resolver-cfg"`
	Uuid             string                                                 `json:"uuid"`
}

type SlbTemplateDnsLocalDnsResolutionHostListCfg1532 struct {
	Hostnames string `json:"hostnames"`
}

type SlbTemplateDnsLocalDnsResolutionLocalResolverCfg1533 struct {
	LocalResolver string `json:"local-resolver"`
}

type SlbTemplateDnsNegativeDnsCache1534 struct {
	EnableNegativeDnsCache int    `json:"enable-negative-dns-cache"`
	BypassQueryThreshold   int    `json:"bypass-query-threshold" dval:"100"`
	MaxNegativeCacheTtl    int    `json:"max-negative-cache-ttl" dval:"7200"`
	CacheNonValid          int    `json:"cache-non-valid"`
	Uuid                   string `json:"uuid"`
}

type SlbTemplateDnsQueryClassFilter1535 struct {
	QueryClassAction string                                         `json:"query-class-action"`
	QueryClass       []SlbTemplateDnsQueryClassFilterQueryClass1536 `json:"query-class"`
	Uuid             string                                         `json:"uuid"`
}

type SlbTemplateDnsQueryClassFilterQueryClass1536 struct {
	StrQueryClass string `json:"str-query-class"`
	NumQueryClass int    `json:"num-query-class"`
}

type SlbTemplateDnsQueryTypeFilter1537 struct {
	QueryTypeAction string                                       `json:"query-type-action"`
	QueryType       []SlbTemplateDnsQueryTypeFilterQueryType1538 `json:"query-type"`
	Uuid            string                                       `json:"uuid"`
}

type SlbTemplateDnsQueryTypeFilterQueryType1538 struct {
	StrQueryType string `json:"str-query-type"`
	NumQueryType int    `json:"num-query-type"`
}

type SlbTemplateDnsRecursiveDnsResolution1539 struct {
	HostListCfg                    []SlbTemplateDnsRecursiveDnsResolutionHostListCfg1540      `json:"host-list-cfg"`
	CsubnetRetry                   int                                                        `json:"csubnet-retry"`
	NsCacheLookup                  string                                                     `json:"ns-cache-lookup" dval:"enabled"`
	NsLongestMatch                 string                                                     `json:"ns-longest-match" dval:"enabled"`
	UseServiceGroupResponse        string                                                     `json:"use-service-group-response" dval:"enabled"`
	Ipv4NatPool                    string                                                     `json:"ipv4-nat-pool"`
	Ipv6NatPool                    string                                                     `json:"ipv6-nat-pool"`
	RetriesPerLevel                int                                                        `json:"retries-per-level" dval:"6"`
	ParallelQueries                int                                                        `json:"parallel-queries" dval:"1"`
	FullResponse                   int                                                        `json:"full-response"`
	MaxTrials                      int                                                        `json:"max-trials" dval:"255"`
	RequestForPendingResolution    string                                                     `json:"request-for-pending-resolution" dval:"respond-with-servfail"`
	UdpRetryInterval               int                                                        `json:"udp-retry-interval" dval:"1"`
	UdpInitialInterval             int                                                        `json:"udp-initial-interval" dval:"5"`
	UseClientQid                   int                                                        `json:"use-client-qid"`
	DefaultRecursive               int                                                        `json:"default-recursive"`
	ForceCnameResolution           string                                                     `json:"force-cname-resolution" dval:"enabled"`
	FastNsSelection                string                                                     `json:"fast-ns-selection" dval:"enabled"`
	DnssecValidation               string                                                     `json:"dnssec-validation" dval:"disabled"`
	EdnsUdpSize                    int                                                        `json:"edns-udp-size" dval:"4096"`
	MaxSignatureValidationAttempts int                                                        `json:"max-signature-validation-attempts"`
	MaxSignatureValidationFailures int                                                        `json:"max-signature-validation-failures"`
	MaxKeyDigestValidationFailures int                                                        `json:"max-key-digest-validation-failures"`
	Uuid                           string                                                     `json:"uuid"`
	LookupOrder                    SlbTemplateDnsRecursiveDnsResolutionLookupOrder1541        `json:"lookup-order"`
	GatewayHealthCheck             SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1543 `json:"gateway-health-check"`
}

type SlbTemplateDnsRecursiveDnsResolutionHostListCfg1540 struct {
	Hostnames string `json:"hostnames"`
}

type SlbTemplateDnsRecursiveDnsResolutionLookupOrder1541 struct {
	QueryType []SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1542 `json:"query-type"`
	Uuid      string                                                         `json:"uuid"`
}

type SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1542 struct {
	StrQueryType string `json:"str-query-type"`
	NumQueryType int    `json:"num-query-type"`
	Order        string `json:"order"`
}

type SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1543 struct {
	QueryName         string `json:"query-name" dval:"a10networks.com"`
	Retry             int    `json:"retry" dval:"6"`
	Timeout           int    `json:"timeout" dval:"5"`
	Interval          int    `json:"interval" dval:"10"`
	UpRetry           int    `json:"up-retry" dval:"1"`
	RetryMulti        int    `json:"retry-multi" dval:"1"`
	GwhcNsCacheLookup string `json:"gwhc-ns-cache-lookup" dval:"disabled"`
	StrQueryType      string `json:"str-query-type" dval:"A"`
	NumQueryType      int    `json:"num-query-type"`
	Uuid              string `json:"uuid"`
}

type SlbTemplateDnsResponseRateLimiting1544 struct {
	ResponseRate       int                                                      `json:"response-rate" dval:"5"`
	NxResponseRate     int                                                      `json:"nx-response-rate" dval:"5"`
	FilterResponseRate int                                                      `json:"filter-response-rate" dval:"10"`
	SlipRate           int                                                      `json:"slip-rate"`
	TcRate             int                                                      `json:"TC-rate"`
	MatchSubnet        string                                                   `json:"match-subnet" dval:"255.255.255.255"`
	MatchSubnetV6      int                                                      `json:"match-subnet-v6" dval:"128"`
	Window             int                                                      `json:"window" dval:"1"`
	SrcIpOnly          int                                                      `json:"src-ip-only"`
	EnableLog          int                                                      `json:"enable-log"`
	Action             string                                                   `json:"action" dval:"rate-limit"`
	Uuid               string                                                   `json:"uuid"`
	RrlClassListList   []SlbTemplateDnsResponseRateLimitingRrlClassListList1545 `json:"rrl-class-list-list"`
}

type SlbTemplateDnsResponseRateLimitingRrlClassListList1545 struct {
	Name    string                                                          `json:"name"`
	Uuid    string                                                          `json:"uuid"`
	UserTag string                                                          `json:"user-tag"`
	LidList []SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1546 `json:"lid-list"`
}

type SlbTemplateDnsResponseRateLimitingRrlClassListListLidList1546 struct {
	Lidnum            int    `json:"lidnum"`
	LidResponseRate   int    `json:"lid-response-rate" dval:"5"`
	LidSlipRate       int    `json:"lid-slip-rate"`
	LidNxResponseRate int    `json:"lid-nx-response-rate" dval:"5"`
	LidTcRate         int    `json:"lid-tc-rate"`
	LidMatchSubnet    string `json:"lid-match-subnet" dval:"255.255.255.255"`
	LidMatchSubnetV6  int    `json:"lid-match-subnet-v6" dval:"128"`
	LidWindow         int    `json:"lid-window" dval:"1"`
	LidSrcIpOnly      int    `json:"lid-src-ip-only"`
	LidEnableLog      int    `json:"lid-enable-log"`
	LidAction         string `json:"lid-action" dval:"rate-limit"`
	Uuid              string `json:"uuid"`
	UserTag           string `json:"user-tag"`
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

type SlbTemplateDnsUdpRetransmit1547 struct {
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
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
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
