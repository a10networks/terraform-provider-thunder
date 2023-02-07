package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
	"strconv"
)

//based on ACOS 5_2_1-P4_81
type SlbVirtualServerPort struct {
	Inst struct {
		AclList                                     []SlbVirtualServerPortAclList        `json:"acl-list"`
		Action                                      string                               `json:"action" dval:"enable"`
		AflexScripts                                []SlbVirtualServerPortAflexScripts   `json:"aflex-scripts"`
		AflexTableEntrySynDisable                   int                                  `json:"aflex-table-entry-syn-disable"`
		AflexTableEntrySynEnable                    int                                  `json:"aflex-table-entry-syn-enable"`
		AltProtocol1                                string                               `json:"alt-protocol1"`
		AltProtocol2                                string                               `json:"alt-protocol2"`
		AlternatePort                               int                                  `json:"alternate-port"`
		AlternatePortNumber                         int                                  `json:"alternate-port-number"`
		AttackDetection                             int                                  `json:"attack-detection"`
		AuthCfg                                     SlbVirtualServerPortAuthCfg          `json:"auth-cfg"`
		Auto                                        int                                  `json:"auto"`
		ClientipStickyNat                           int                                  `json:"clientip-sticky-nat"`
		ConnLimit                                   int                                  `json:"conn-limit" dval:"64000000"`
		CpuCompute                                  int                                  `json:"cpu-compute"`
		DefSelectionIfPrefFailed                    string                               `json:"def-selection-if-pref-failed" dval:"def-selection-if-pref-failed"`
		EnablePlayeridCheck                         int                                  `json:"enable-playerid-check"`
		EnableScaleout                              int                                  `json:"enable-scaleout"`
		EthFwd                                      int                                  `json:"eth-fwd"`
		EthRev                                      int                                  `json:"eth-rev"`
		Expand                                      int                                  `json:"expand"`
		ExtendedStats                               int                                  `json:"extended-stats"`
		ForceRoutingMode                            int                                  `json:"force-routing-mode"`
		GslbEnable                                  int                                  `json:"gslb-enable"`
		GtpSessionLb                                int                                  `json:"gtp-session-lb"`
		HaConnMirror                                int                                  `json:"ha-conn-mirror"`
		IgnoreGlobal                                int                                  `json:"ignore-global"`
		IpMapList                                   string                               `json:"ip-map-list"`
		IpOnlyLb                                    int                                  `json:"ip-only-lb"`
		IpSmartRr                                   int                                  `json:"ip-smart-rr"`
		Ipinip                                      int                                  `json:"ipinip"`
		L7HardwareAssist                            int                                  `json:"l7-hardware-assist"`
		L7ServiceChain                              int                                  `json:"l7-service-chain"`
		MemoryCompute                               int                                  `json:"memory-compute"`
		MessageSwitching                            int                                  `json:"message-switching"`
		Name                                        string                               // outside key
		NoAutoUpOnAflex                             int                                  `json:"no-auto-up-on-aflex"`
		NoDestNat                                   int                                  `json:"no-dest-nat"`
		NoLogging                                   int                                  `json:"no-logging"`
		OnSyn                                       int                                  `json:"on-syn"`
		OneServerConn                               int                                  `json:"one-server-conn"`
		OptimizationLevel                           string                               `json:"optimization-level" dval:"0"`
		PTemplateSipShared                          int                                  `json:"p-template-sip-shared"`
		PacketCaptureTemplate                       string                               `json:"packet-capture-template"`
		PersistType                                 string                               `json:"persist-type"`
		Pool                                        string                               `json:"pool"`
		PoolShared                                  string                               `json:"pool-shared"`
		PortNumber                                  int                                  `json:"port-number" dval:"-1"`
		PortTranslation                             int                                  `json:"port-translation"`
		Precedence                                  int                                  `json:"precedence"`
		Protocol                                    string                               `json:"protocol"`
		ProxyLayer                                  string                               `json:"proxy-layer"`
		Range                                       int                                  `json:"range"`
		Rate                                        int                                  `json:"rate"`
		RedirectToHttps                             int                                  `json:"redirect-to-https"`
		ReplyAcmeChallenge                          int                                  `json:"reply-acme-challenge"`
		ReqFail                                     int                                  `json:"req-fail"`
		Reselection                                 string                               `json:"reselection"`
		Reset                                       int                                  `json:"reset"`
		ResetOnServerSelectionFail                  int                                  `json:"reset-on-server-selection-fail"`
		ResolveWebCatList                           string                               `json:"resolve-web-cat-list"`
		RtpSipCallIdMatch                           int                                  `json:"rtp-sip-call-id-match"`
		SamplingEnable                              []SlbVirtualServerPortSamplingEnable `json:"sampling-enable"`
		ScaleoutBucketCount                         int                                  `json:"scaleout-bucket-count" dval:"32"`
		ScaleoutDeviceGroup                         int                                  `json:"scaleout-device-group"`
		Secs                                        int                                  `json:"secs"`
		ServSelFail                                 int                                  `json:"serv-sel-fail"`
		ServerGroup                                 string                               `json:"server-group"`
		ServiceGroup                                string                               `json:"service-group"`
		SharedPartitionCacheTemplate                int                                  `json:"shared-partition-cache-template"`
		SharedPartitionClientSslTemplate            int                                  `json:"shared-partition-client-ssl-template"`
		SharedPartitionConnectionReuseTemplate      int                                  `json:"shared-partition-connection-reuse-template"`
		SharedPartitionDblbTemplate                 int                                  `json:"shared-partition-dblb-template"`
		SharedPartitionDiameterTemplate             int                                  `json:"shared-partition-diameter-template"`
		SharedPartitionDnsTemplate                  int                                  `json:"shared-partition-dns-template"`
		SharedPartitionDohTemplate                  int                                  `json:"shared-partition-doh-template"`
		SharedPartitionDynamicServiceTemplate       int                                  `json:"shared-partition-dynamic-service-template"`
		SharedPartitionExternalServiceTemplate      int                                  `json:"shared-partition-external-service-template"`
		SharedPartitionFixTemplate                  int                                  `json:"shared-partition-fix-template"`
		SharedPartitionHttpPolicyTemplate           int                                  `json:"shared-partition-http-policy-template"`
		SharedPartitionHttpTemplate                 int                                  `json:"shared-partition-http-template"`
		SharedPartitionImapPop3Template             int                                  `json:"shared-partition-imap-pop3-template"`
		SharedPartitionPersistCookieTemplate        int                                  `json:"shared-partition-persist-cookie-template"`
		SharedPartitionPersistDestinationIpTemplate int                                  `json:"shared-partition-persist-destination-ip-template"`
		SharedPartitionPersistSourceIpTemplate      int                                  `json:"shared-partition-persist-source-ip-template"`
		SharedPartitionPersistSslSidTemplate        int                                  `json:"shared-partition-persist-ssl-sid-template"`
		SharedPartitionPolicyTemplate               int                                  `json:"shared-partition-policy-template"`
		SharedPartitionPool                         int                                  `json:"shared-partition-pool"`
		SharedPartitionServerSslTemplate            int                                  `json:"shared-partition-server-ssl-template"`
		SharedPartitionSmppTemplate                 int                                  `json:"shared-partition-smpp-template"`
		SharedPartitionSmtpTemplate                 int                                  `json:"shared-partition-smtp-template"`
		SharedPartitionTcp                          int                                  `json:"shared-partition-tcp"`
		SharedPartitionTcpProxyTemplate             int                                  `json:"shared-partition-tcp-proxy-template"`
		SharedPartitionUdp                          int                                  `json:"shared-partition-udp"`
		SharedPartitionVirtualPortTemplate          int                                  `json:"shared-partition-virtual-port-template"`
		SkipRevHash                                 int                                  `json:"skip-rev-hash"`
		SnatOnVip                                   int                                  `json:"snat-on-vip"`
		StatsDataAction                             string                               `json:"stats-data-action" dval:"stats-data-enable"`
		SubstituteSourceMac                         int                                  `json:"substitute-source-mac"`
		SupportHttp2                                int                                  `json:"support-http2"`
		SynCookie                                   int                                  `json:"syn-cookie"`
		TemplateCache                               string                               `json:"template-cache"`
		TemplateCacheShared                         string                               `json:"template-cache-shared"`
		TemplateClientSsh                           string                               `json:"template-client-ssh"`
		TemplateClientSsl                           string                               `json:"template-client-ssl"`
		TemplateClientSslShared                     string                               `json:"template-client-ssl-shared"`
		TemplateConnectionReuse                     string                               `json:"template-connection-reuse"`
		TemplateConnectionReuseShared               string                               `json:"template-connection-reuse-shared"`
		TemplateDblb                                string                               `json:"template-dblb"`
		TemplateDblbShared                          string                               `json:"template-dblb-shared"`
		TemplateDiameter                            string                               `json:"template-diameter"`
		TemplateDiameterShared                      string                               `json:"template-diameter-shared"`
		TemplateDns                                 string                               `json:"template-dns"`
		TemplateDnsShared                           string                               `json:"template-dns-shared"`
		TemplateDoh                                 string                               `json:"template-doh"`
		TemplateDohShared                           string                               `json:"template-doh-shared"`
		TemplateDynamicService                      string                               `json:"template-dynamic-service"`
		TemplateDynamicServiceShared                string                               `json:"template-dynamic-service-shared"`
		TemplateExternalService                     string                               `json:"template-external-service"`
		TemplateExternalServiceShared               string                               `json:"template-external-service-shared"`
		TemplateFix                                 string                               `json:"template-fix"`
		TemplateFixShared                           string                               `json:"template-fix-shared"`
		TemplateFtp                                 string                               `json:"template-ftp"`
		TemplateHttp                                string                               `json:"template-http"`
		TemplateHttpPolicy                          string                               `json:"template-http-policy"`
		TemplateHttpPolicyShared                    string                               `json:"template-http-policy-shared"`
		TemplateHttpShared                          string                               `json:"template-http-shared"`
		TemplateImapPop3                            string                               `json:"template-imap-pop3"`
		TemplateImapPop3Shared                      string                               `json:"template-imap-pop3-shared"`
		TemplateMqtt                                string                               `json:"template-mqtt"`
		TemplatePersistCookie                       string                               `json:"template-persist-cookie"`
		TemplatePersistCookieShared                 string                               `json:"template-persist-cookie-shared"`
		TemplatePersistDestinationIp                string                               `json:"template-persist-destination-ip"`
		TemplatePersistDestinationIpShared          string                               `json:"template-persist-destination-ip-shared"`
		TemplatePersistSourceIp                     string                               `json:"template-persist-source-ip"`
		TemplatePersistSourceIpShared               string                               `json:"template-persist-source-ip-shared"`
		TemplatePersistSslSid                       string                               `json:"template-persist-ssl-sid"`
		TemplatePersistSslSidShared                 string                               `json:"template-persist-ssl-sid-shared"`
		TemplatePolicy                              string                               `json:"template-policy"`
		TemplatePolicyShared                        string                               `json:"template-policy-shared"`
		TemplateRamCache                            string                               `json:"template-ram-cache"`
		TemplateReqmodIcap                          string                               `json:"template-reqmod-icap"`
		TemplateRespmodIcap                         string                               `json:"template-respmod-icap"`
		TemplateScaleout                            string                               `json:"template-scaleout"`
		TemplateServerSsh                           string                               `json:"template-server-ssh"`
		TemplateServerSsl                           string                               `json:"template-server-ssl"`
		TemplateServerSslShared                     string                               `json:"template-server-ssl-shared"`
		TemplateSip                                 string                               `json:"template-sip"`
		TemplateSipShared                           string                               `json:"template-sip-shared"`
		TemplateSmpp                                string                               `json:"template-smpp"`
		TemplateSmppShared                          string                               `json:"template-smpp-shared"`
		TemplateSmtp                                string                               `json:"template-smtp"`
		TemplateSmtpShared                          string                               `json:"template-smtp-shared"`
		TemplateSsli                                string                               `json:"template-ssli"`
		TemplateTcp                                 string                               `json:"template-tcp" dval:"default"`
		TemplateTcpProxy                            string                               `json:"template-tcp-proxy"`
		TemplateTcpProxyClient                      string                               `json:"template-tcp-proxy-client"`
		TemplateTcpProxyServer                      string                               `json:"template-tcp-proxy-server"`
		TemplateTcpProxyShared                      string                               `json:"template-tcp-proxy-shared"`
		TemplateTcpShared                           string                               `json:"template-tcp-shared" dval:"default"`
		TemplateUdp                                 string                               `json:"template-udp" dval:"default"`
		TemplateUdpShared                           string                               `json:"template-udp-shared" dval:"default"`
		TemplateVirtualPort                         string                               `json:"template-virtual-port" dval:"default"`
		TemplateVirtualPortShared                   string                               `json:"template-virtual-port-shared"`
		TrunkFwd                                    int                                  `json:"trunk-fwd"`
		TrunkRev                                    int                                  `json:"trunk-rev"`
		UseAlternatePort                            int                                  `json:"use-alternate-port"`
		UseCgnv6                                    int                                  `json:"use-cgnv6"`
		UseDefaultIfNoServer                        int                                  `json:"use-default-if-no-server"`
		UseRcvHopForResp                            int                                  `json:"use-rcv-hop-for-resp"`
		UseRcvHopGroup                              int                                  `json:"use-rcv-hop-group"`
		UserTag                                     string                               `json:"user-tag"`
		Uuid                                        string                               `json:"uuid"`
		View                                        int                                  `json:"view"`
		WafTemplate                                 string                               `json:"waf-template"`
		WhenDown                                    int                                  `json:"when-down"`
		WhenDownProtocol2                           int                                  `json:"when-down-protocol2"`
	} `json:"port"`
}

type SlbVirtualServerPortAclList struct {
	AclId                    int    `json:"acl-id"`
	AclName                  string `json:"acl-name"`
	AclIdShared              int    `json:"acl-id-shared"`
	AclNameShared            string `json:"acl-name-shared"`
	AclIdSrcNatPool          string `json:"acl-id-src-nat-pool"`
	AclIdSeqNum              int    `json:"acl-id-seq-num"`
	SharedPartitionPoolId    int    `json:"shared-partition-pool-id"`
	AclIdSrcNatPoolShared    string `json:"acl-id-src-nat-pool-shared"`
	AclIdSeqNumShared        int    `json:"acl-id-seq-num-shared"`
	VAclIdSrcNatPool         string `json:"v-acl-id-src-nat-pool"`
	VAclIdSeqNum             int    `json:"v-acl-id-seq-num"`
	VSharedPartitionPoolId   int    `json:"v-shared-partition-pool-id"`
	VAclIdSrcNatPoolShared   string `json:"v-acl-id-src-nat-pool-shared"`
	VAclIdSeqNumShared       int    `json:"v-acl-id-seq-num-shared"`
	AclNameSrcNatPool        string `json:"acl-name-src-nat-pool"`
	AclNameSeqNum            int    `json:"acl-name-seq-num"`
	SharedPartitionPoolName  int    `json:"shared-partition-pool-name"`
	AclNameSrcNatPoolShared  string `json:"acl-name-src-nat-pool-shared"`
	AclNameSeqNumShared      int    `json:"acl-name-seq-num-shared"`
	VAclNameSrcNatPool       string `json:"v-acl-name-src-nat-pool"`
	VAclNameSeqNum           int    `json:"v-acl-name-seq-num"`
	VSharedPartitionPoolName int    `json:"v-shared-partition-pool-name"`
	VAclNameSrcNatPoolShared string `json:"v-acl-name-src-nat-pool-shared"`
	VAclNameSeqNumShared     int    `json:"v-acl-name-seq-num-shared"`
}

type SlbVirtualServerPortAflexScripts struct {
	Aflex       string `json:"aflex"`
	AflexShared string `json:"aflex-shared"`
}

type SlbVirtualServerPortAuthCfg struct {
	AaaPolicy string `json:"aaa-policy"`
}

type SlbVirtualServerPortSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *SlbVirtualServerPort) GetId() string {
	return strconv.Itoa(p.Inst.PortNumber) + "+" + p.Inst.Protocol
}

func (p *SlbVirtualServerPort) getPath() string {
	return "slb/virtual-server/" + url.QueryEscape(p.Inst.Name) + "/port"
}

func (p *SlbVirtualServerPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbVirtualServerPort::Post")
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

func (p *SlbVirtualServerPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbVirtualServerPort::Get")
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

func (p *SlbVirtualServerPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbVirtualServerPort::Put")
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

func (p *SlbVirtualServerPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbVirtualServerPort::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
