package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P6_74
type GslbPolicy struct {
	Inst struct {
		ActiveRdt                   GslbPolicyActiveRdt         `json:"active-rdt"`
		ActiveServersEnable         int                         `json:"active-servers-enable"`
		ActiveServersFailBreak      int                         `json:"active-servers-fail-break"`
		AdminIpEnable               int                         `json:"admin-ip-enable"`
		AdminIpTopOnly              int                         `json:"admin-ip-top-only"`
		AdminPreference             int                         `json:"admin-preference"`
		AliasAdminPreference        int                         `json:"alias-admin-preference"`
		AmountFirst                 int                         `json:"amount-first"`
		AutoMap                     GslbPolicyAutoMap           `json:"auto-map"`
		BwCostEnable                int                         `json:"bw-cost-enable"`
		BwCostFailBreak             int                         `json:"bw-cost-fail-break"`
		Capacity                    GslbPolicyCapacity          `json:"capacity"`
		ConnectionLoad              GslbPolicyConnectionLoad    `json:"connection-load"`
		Dns                         GslbPolicyDns               `json:"dns"`
		Edns                        GslbPolicyEdns              `json:"edns"`
		GeoLocationList             []GslbPolicyGeoLocationList `json:"geo-location-list"`
		GeoLocationMatch            GslbPolicyGeoLocationMatch  `json:"geo-location-match"`
		Geographic                  int                         `json:"geographic" dval:"1"`
		HealthCheck                 int                         `json:"health-check" dval:"1"`
		HealthCheckPreferenceEnable int                         `json:"health-check-preference-enable"`
		HealthPreferenceTop         int                         `json:"health-preference-top"`
		IpList                      string                      `json:"ip-list"`
		LeastResponse               int                         `json:"least-response"`
		MetricFailBreak             int                         `json:"metric-fail-break"`
		MetricForceCheck            int                         `json:"metric-force-check"`
		MetricOrder                 int                         `json:"metric-order"`
		MetricType                  string                      `json:"metric-type"`
		Name                        string                      `json:"name"`
		NumSessionEnable            int                         `json:"num-session-enable"`
		NumSessionTolerance         int                         `json:"num-session-tolerance" dval:"10"`
		OrderedIpTopOnly            int                         `json:"ordered-ip-top-only"`
		RoundRobin                  int                         `json:"round-robin" dval:"1"`
		UserTag                     string                      `json:"user-tag"`
		Uuid                        string                      `json:"uuid"`
		WeightedAlias               int                         `json:"weighted-alias"`
		WeightedIpEnable            int                         `json:"weighted-ip-enable"`
		WeightedIpTotalHits         int                         `json:"weighted-ip-total-hits"`
		WeightedSiteEnable          int                         `json:"weighted-site-enable"`
		WeightedSiteTotalHits       int                         `json:"weighted-site-total-hits"`
	} `json:"policy"`
}
type GslbPolicyActiveRdt struct {
	Enable         int    `json:"enable"`
	SingleShot     int    `json:"single-shot"`
	Timeout        int    `json:"timeout" dval:"3"`
	Skip           int    `json:"skip" dval:"3"`
	KeepTracking   int    `json:"keep-tracking"`
	IgnoreId       int    `json:"ignore-id"`
	Samples        int    `json:"samples" dval:"5"`
	Tolerance      int    `json:"tolerance" dval:"10"`
	Difference     int    `json:"difference"`
	Limit          int    `json:"limit" dval:"16383"`
	FailBreak      int    `json:"fail-break"`
	Controller     int    `json:"controller"`
	ProtoRdtEnable int    `json:"proto-rdt-enable"`
	Uuid           string `json:"uuid"`
}
type GslbPolicyAutoMap struct {
	Ttl           int    `json:"ttl" dval:"300"`
	ModuleDisable int    `json:"module-disable"`
	All           int    `json:"all"`
	ModuleType    string `json:"module-type"`
	Uuid          string `json:"uuid"`
}
type GslbPolicyCapacity struct {
	CapacityEnable    int    `json:"capacity-enable"`
	Threshold         int    `json:"threshold" dval:"90"`
	CapacityFailBreak int    `json:"capacity-fail-break"`
	Uuid              string `json:"uuid"`
}
type GslbPolicyConnectionLoad struct {
	ConnectionLoadEnable    int    `json:"connection-load-enable"`
	ConnectionLoadFailBreak int    `json:"connection-load-fail-break"`
	ConnectionLoadSamples   int    `json:"connection-load-samples" dval:"5"`
	ConnectionLoadInterval  int    `json:"connection-load-interval" dval:"5"`
	Limit                   int    `json:"limit"`
	ConnectionLoadLimit     int    `json:"connection-load-limit"`
	Uuid                    string `json:"uuid"`
}
type GslbPolicyDns struct {
	Action                  int                                    `json:"action"`
	ActiveOnly              int                                    `json:"active-only"`
	ActiveOnlyFailSafe      int                                    `json:"active-only-fail-safe"`
	DnsAdditionMx           int                                    `json:"dns-addition-mx"`
	DnsAutoMap              int                                    `json:"dns-auto-map"`
	BackupAlias             int                                    `json:"backup-alias"`
	BackupServer            int                                    `json:"backup-server"`
	ExternalIp              int                                    `json:"external-ip" dval:"1"`
	ExternalSoa             int                                    `json:"external-soa"`
	CnameDetect             int                                    `json:"cname-detect" dval:"1"`
	IpReplace               int                                    `json:"ip-replace"`
	GeolocAlias             int                                    `json:"geoloc-alias"`
	GeolocAction            int                                    `json:"geoloc-action"`
	GeolocPolicy            int                                    `json:"geoloc-policy"`
	SelectedOnly            int                                    `json:"selected-only"`
	SelectedOnlyValue       int                                    `json:"selected-only-value"`
	Cache                   int                                    `json:"cache"`
	AgingTime               int                                    `json:"aging-time"`
	Delegation              int                                    `json:"delegation"`
	Hint                    string                                 `json:"hint" dval:"addition"`
	Logging                 string                                 `json:"logging"`
	Template                string                                 `json:"template"`
	Ttl                     int                                    `json:"ttl" dval:"10"`
	UseServerTtl            int                                    `json:"use-server-ttl"`
	Server                  int                                    `json:"server"`
	ServerSrv               int                                    `json:"server-srv"`
	ServerMx                int                                    `json:"server-mx"`
	ServerNaptr             int                                    `json:"server-naptr"`
	ServerAdditionMx        int                                    `json:"server-addition-mx"`
	ServerNs                int                                    `json:"server-ns"`
	ServerAutoNs            int                                    `json:"server-auto-ns"`
	ServerPtr               int                                    `json:"server-ptr"`
	ServerAutoPtr           int                                    `json:"server-auto-ptr"`
	ServerTxt               int                                    `json:"server-txt"`
	ServerCustom            int                                    `json:"server-custom"`
	ServerAny               int                                    `json:"server-any"`
	ServerAnyWithMetric     int                                    `json:"server-any-with-metric"`
	ServerAuthoritative     int                                    `json:"server-authoritative"`
	ServerSec               int                                    `json:"server-sec"`
	ServerNsList            int                                    `json:"server-ns-list"`
	ServerFullList          int                                    `json:"server-full-list"`
	ServerModeOnly          int                                    `json:"server-mode-only"`
	ServerCname             int                                    `json:"server-cname"`
	Ipv6                    []GslbPolicyDnsIpv6                    `json:"ipv6"`
	BlockAction             int                                    `json:"block-action"`
	ActionType              string                                 `json:"action-type"`
	ProxyBlockPortRangeList []GslbPolicyDnsProxyBlockPortRangeList `json:"proxy-block-port-range-list"`
	BlockValue              []GslbPolicyDnsBlockValue              `json:"block-value"`
	BlockType               string                                 `json:"block-type"`
	Sticky                  int                                    `json:"sticky"`
	StickyMask              string                                 `json:"sticky-mask" dval:"/32"`
	StickyIpv6Mask          int                                    `json:"sticky-ipv6-mask" dval:"128"`
	StickyAgingTime         int                                    `json:"sticky-aging-time" dval:"5"`
	DynamicPreference       int                                    `json:"dynamic-preference"`
	DynamicWeight           int                                    `json:"dynamic-weight"`
	Uuid                    string                                 `json:"uuid"`
}
type GslbPolicyDnsIpv6 struct {
	DnsIpv6Option      string `json:"dns-ipv6-option"`
	DnsIpv6MappingType string `json:"dns-ipv6-mapping-type"`
}
type GslbPolicyDnsProxyBlockPortRangeList struct {
	ProxyBlockRangeFrom int `json:"proxy-block-range-from"`
	ProxyBlockRangeTo   int `json:"proxy-block-range-to"`
}
type GslbPolicyDnsBlockValue struct {
	BlockValue int `json:"block-value"`
}
type GslbPolicyEdns struct {
	ClientSubnetGeographic int    `json:"client-subnet-geographic"`
	Uuid                   string `json:"uuid"`
}
type GslbPolicyGeoLocationList struct {
	Name               string                                        `json:"name"`
	IpMultipleFields   []GslbPolicyGeoLocationListIpMultipleFields   `json:"ip-multiple-fields"`
	Ipv6MultipleFields []GslbPolicyGeoLocationListIpv6MultipleFields `json:"ipv6-multiple-fields"`
	Uuid               string                                        `json:"uuid"`
	UserTag            string                                        `json:"user-tag"`
}
type GslbPolicyGeoLocationListIpMultipleFields struct {
	IpSub      string `json:"ip-sub"`
	IpMaskSub  string `json:"ip-mask-sub"`
	IpAddr2Sub string `json:"ip-addr2-sub"`
}
type GslbPolicyGeoLocationListIpv6MultipleFields struct {
	Ipv6Sub      string `json:"ipv6-sub"`
	Ipv6MaskSub  int    `json:"ipv6-mask-sub"`
	Ipv6Addr2Sub string `json:"ipv6-addr2-sub"`
}
type GslbPolicyGeoLocationMatch struct {
	Overlap        int    `json:"overlap"`
	GeoTypeOverlap string `json:"geo-type-overlap"`
	MatchFirst     string `json:"match-first" dval:"global"`
	Uuid           string `json:"uuid"`
}

func (p *GslbPolicy) GetId() string {
	return p.Inst.Name
}
func (p *GslbPolicy) getPath() string {
	return "gslb/policy"
}
func (p *GslbPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicy::Post")
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
func (p *GslbPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicy::Get")
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
func (p *GslbPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicy::Put")
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
func (p *GslbPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("GslbPolicy::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
