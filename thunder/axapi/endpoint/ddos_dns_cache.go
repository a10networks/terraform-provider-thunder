package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type DdosDnsCache struct {
	Inst struct {
		AnyQueryActionStr string `json:"any-query-action-str" dval:"respond-refuse"`

		DefaultServingAction string `json:"default-serving-action" dval:"serve-from-cache"`

		DnsLogging string `json:"dns-logging"`

		DomainGroup DdosDnsCacheDomainGroup171 `json:"domain-group"`

		EdnsUdpSize int `json:"edns-udp-size"`

		FqdnManualOverrideActionList []DdosDnsCacheFqdnManualOverrideActionList `json:"fqdn-manual-override-action-list"`

		Name string `json:"name"`

		NegCacheActionFollowQRate int `json:"neg-cache-action-follow-q-rate"`

		NonAuthoritativeZoneQueryActionStr string `json:"non-authoritative-zone-query-action-str" dval:"respond-refuse"`

		RespondServfail int `json:"respond-servfail"`

		SamplingEnable []DdosDnsCacheSamplingEnable `json:"sampling-enable"`

		ShardedDomainGroupList []DdosDnsCacheShardedDomainGroupList `json:"sharded-domain-group-list"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		ZoneDomainLookupMissAction string `json:"zone-domain-lookup-miss-action" dval:"respond-nxdomain"`

		ZoneManualOverrideActionList []DdosDnsCacheZoneManualOverrideActionList `json:"zone-manual-override-action-list"`

		ZoneTransfer DdosDnsCacheZoneTransfer175 `json:"zone-transfer"`
	} `json:"dns-cache"`
}

type DdosDnsCacheDomainGroup171 struct {
	Name                 string                                           `json:"name"`
	Uuid                 string                                           `json:"uuid"`
	DomainListPolicyList []DdosDnsCacheDomainGroupDomainListPolicyList172 `json:"domain-list-policy-list"`
}

type DdosDnsCacheDomainGroupDomainListPolicyList172 struct {
	Name                   string                                                        `json:"name"`
	ServerIpv4             string                                                        `json:"server-ipv4"`
	ServerV4Port           int                                                           `json:"server-v4-port" dval:"53"`
	ClientIpv4             string                                                        `json:"client-ipv4"`
	DnsNotifyEnableIpv4    int                                                           `json:"dns-notify-enable-ipv4"`
	ServerIpv6             string                                                        `json:"server-ipv6"`
	ServerV6Port           int                                                           `json:"server-v6-port" dval:"53"`
	ClientIpv6             string                                                        `json:"client-ipv6"`
	DnsNotifyEnableIpv6    int                                                           `json:"dns-notify-enable-ipv6"`
	RefreshIntervalHours   int                                                           `json:"refresh-interval-hours" dval:"4"`
	TtlOverride            int                                                           `json:"ttl-override"`
	RespondWithAuthority   int                                                           `json:"respond-with-authority"`
	OversizeAnswerResponse string                                                        `json:"oversize-answer-response" dval:"set-truncate-bit"`
	ResolveCnameRecord     int                                                           `json:"resolve-cname-record"`
	ManualRefresh          string                                                        `json:"manual-refresh"`
	Force                  int                                                           `json:"force"`
	CacheAllRecords        int                                                           `json:"cache-all-records"`
	CacheDnssecRecords     int                                                           `json:"cache-dnssec-records"`
	Uuid                   string                                                        `json:"uuid"`
	UserTag                string                                                        `json:"user-tag"`
	PacketCapturing        DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing173 `json:"packet-capturing"`
}

type DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturing173 struct {
	RootZoneList []DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList174 `json:"root-zone-list"`
	Uuid         string                                                                      `json:"uuid"`
}

type DdosDnsCacheDomainGroupDomainListPolicyListPacketCapturingRootZoneList174 struct {
	RootZone      string `json:"root-zone"`
	CaptureConfig string `json:"capture-config"`
	CaptureMode   string `json:"capture-mode"`
}

type DdosDnsCacheFqdnManualOverrideActionList struct {
	FqdnName string `json:"fqdn-name"`
	Action   string `json:"action"`
}

type DdosDnsCacheSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDnsCacheShardedDomainGroupList struct {
	Name                        string                                                          `json:"name"`
	MatchAction                 string                                                          `json:"match-action" dval:"forward"`
	EncapTemplate               string                                                          `json:"encap-template"`
	Uuid                        string                                                          `json:"uuid"`
	UserTag                     string                                                          `json:"user-tag"`
	ShardedDomainListPolicyList []DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyList `json:"sharded-domain-list-policy-list"`
}

type DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyList struct {
	Name                 string                                                                       `json:"name"`
	ServerIpv4           string                                                                       `json:"server-ipv4"`
	ServerV4Port         int                                                                          `json:"server-v4-port" dval:"53"`
	ClientIpv4           string                                                                       `json:"client-ipv4"`
	DnsNotifyEnableIpv4  int                                                                          `json:"dns-notify-enable-ipv4"`
	ServerIpv6           string                                                                       `json:"server-ipv6"`
	ServerV6Port         int                                                                          `json:"server-v6-port" dval:"53"`
	ClientIpv6           string                                                                       `json:"client-ipv6"`
	DnsNotifyEnableIpv6  int                                                                          `json:"dns-notify-enable-ipv6"`
	RefreshIntervalHours int                                                                          `json:"refresh-interval-hours" dval:"4"`
	ManualRefresh        string                                                                       `json:"manual-refresh"`
	Force                int                                                                          `json:"force"`
	Uuid                 string                                                                       `json:"uuid"`
	UserTag              string                                                                       `json:"user-tag"`
	PacketCapturing      DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturing `json:"packet-capturing"`
}

type DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturing struct {
	RootZoneList []DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturingRootZoneList `json:"root-zone-list"`
	Uuid         string                                                                                     `json:"uuid"`
}

type DdosDnsCacheShardedDomainGroupListShardedDomainListPolicyListPacketCapturingRootZoneList struct {
	RootZone      string `json:"root-zone"`
	CaptureConfig string `json:"capture-config"`
	CaptureMode   string `json:"capture-mode"`
}

type DdosDnsCacheZoneManualOverrideActionList struct {
	ZoneName string `json:"zone-name"`
	Action   string `json:"action"`
}

type DdosDnsCacheZoneTransfer175 struct {
	Uuid string `json:"uuid"`
}

func (p *DdosDnsCache) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *DdosDnsCache) getPath() string {
	return "ddos/dns-cache"
}

func (p *DdosDnsCache) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsCache::Post")
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

func (p *DdosDnsCache) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsCache::Get")
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
func (p *DdosDnsCache) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsCache::Put")
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

func (p *DdosDnsCache) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsCache::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
