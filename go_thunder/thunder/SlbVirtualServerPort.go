package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SlbVirtualServerPort struct {
	HaConnMirror SlbVirtualServerPortInstance `json:"port-instance,omitempty"`
}

type SlbVirtualServerPortInstance struct {
	VAclIDSeqNum                                []SlbVirtualServerAclIDList      `json:"acl-id-list,omitempty"`
	AclNameSrcNatPoolShared                     []SlbVirtualServerAclNameList    `json:"acl-name-list,omitempty"`
	Action                                      string                           `json:"action,omitempty"`
	Aflex                                       []SlbVirtualServerAflexScripts   `json:"aflex-scripts,omitempty"`
	AltProtocol1                                string                           `json:"alt-protocol1,omitempty"`
	AltProtocol2                                string                           `json:"alt-protocol2,omitempty"`
	AlternatePort                               int                              `json:"alternate-port,omitempty"`
	AlternatePortNumber                         int                              `json:"alternate-port-number,omitempty"`
	AaaPolicy                                   SlbVirtualServerAuthCfg          `json:"auth-cfg,omitempty"`
	Auto                                        int                              `json:"auto,omitempty"`
	CPUCompute                                  int                              `json:"cpu-compute,omitempty"`
	ClientipStickyNat                           int                              `json:"clientip-sticky-nat,omitempty"`
	ConnLimit                                   int                              `json:"conn-limit,omitempty"`
	DefSelectionIfPrefFailed                    string                           `json:"def-selection-if-pref-failed,omitempty"`
	EnablePlayeridCheck                         int                              `json:"enable-playerid-check,omitempty"`
	EthFwd                                      int                              `json:"eth-fwd,omitempty"`
	EthRev                                      int                              `json:"eth-rev,omitempty"`
	Expand                                      int                              `json:"expand,omitempty"`
	ExtendedStats                               int                              `json:"extended-stats,omitempty"`
	ForceRoutingMode                            int                              `json:"force-routing-mode,omitempty"`
	GslbEnable                                  int                              `json:"gslb-enable,omitempty"`
	GtpSessionLb                                int                              `json:"gtp-session-lb,omitempty"`
	HaConnMirror                                int                              `json:"ha-conn-mirror,omitempty"`
	IPMapList                                   string                           `json:"ip-map-list,omitempty"`
	IgnoreGlobal                                int                              `json:"ignore-global,omitempty"`
	Ipinip                                      int                              `json:"ipinip,omitempty"`
	L7HardwareAssist                            int                              `json:"l7-hardware-assist,omitempty"`
	MemoryCompute                               int                              `json:"memory-compute,omitempty"`
	MessageSwitching                            int                              `json:"message-switching,omitempty"`
	Name                                        string                           `json:"name,omitempty"`
	NoAutoUpOnAflex                             int                              `json:"no-auto-up-on-aflex,omitempty"`
	NoDestNat                                   int                              `json:"no-dest-nat,omitempty"`
	NoLogging                                   int                              `json:"no-logging,omitempty"`
	OnSyn                                       int                              `json:"on-syn,omitempty"`
	OptimizationLevel                           string                           `json:"optimization-level,omitempty"`
	PTemplateSipShared                          int                              `json:"p-template-sip-shared,omitempty"`
	PersistType                                 string                           `json:"persist-type,omitempty"`
	Pool                                        string                           `json:"pool,omitempty"`
	PoolShared                                  string                           `json:"pool-shared,omitempty"`
	PortNumber                                  int                              `json:"port-number,omitempty"`
	PortTranslation                             int                              `json:"port-translation,omitempty"`
	Precedence                                  int                              `json:"precedence,omitempty"`
	Protocol                                    string                           `json:"protocol,omitempty"`
	Range                                       int                              `json:"range,omitempty"`
	Rate                                        int                              `json:"rate,omitempty"`
	RedirectToHTTPS                             int                              `json:"redirect-to-https,omitempty"`
	ReqFail                                     int                              `json:"req-fail,omitempty"`
	Reset                                       int                              `json:"reset,omitempty"`
	ResetOnServerSelectionFail                  int                              `json:"reset-on-server-selection-fail,omitempty"`
	RtpSipCallIDMatch                           int                              `json:"rtp-sip-call-id-match,omitempty"`
	Counters1                                   []SlbVirtualServerSamplingEnable `json:"sampling-enable,omitempty"`
	ScaleoutBucketCount                         int                              `json:"scaleout-bucket-count,omitempty"`
	ScaleoutDeviceGroup                         int                              `json:"scaleout-device-group,omitempty"`
	Secs                                        int                              `json:"secs,omitempty"`
	ServSelFail                                 int                              `json:"serv-sel-fail,omitempty"`
	ServiceGroup                                string                           `json:"service-group,omitempty"`
	SharedPartitionCacheTemplate                int                              `json:"shared-partition-cache-template,omitempty"`
	SharedPartitionClientSslTemplate            int                              `json:"shared-partition-client-ssl-template,omitempty"`
	SharedPartitionConnectionReuseTemplate      int                              `json:"shared-partition-connection-reuse-template,omitempty"`
	SharedPartitionDNSTemplate                  int                              `json:"shared-partition-dns-template,omitempty"`
	SharedPartitionDblbTemplate                 int                              `json:"shared-partition-dblb-template,omitempty"`
	SharedPartitionDiameterTemplate             int                              `json:"shared-partition-diameter-template,omitempty"`
	SharedPartitionDohTemplate                  int                              `json:"shared-partition-doh-template,omitempty"`
	SharedPartitionDynamicServiceTemplate       int                              `json:"shared-partition-dynamic-service-template,omitempty"`
	SharedPartitionExternalServiceTemplate      int                              `json:"shared-partition-external-service-template,omitempty"`
	SharedPartitionFixTemplate                  int                              `json:"shared-partition-fix-template,omitempty"`
	SharedPartitionHTTPPolicyTemplate           int                              `json:"shared-partition-http-policy-template,omitempty"`
	SharedPartitionHTTPTemplate                 int                              `json:"shared-partition-http-template,omitempty"`
	SharedPartitionImapPop3Template             int                              `json:"shared-partition-imap-pop3-template,omitempty"`
	SharedPartitionPersistCookieTemplate        int                              `json:"shared-partition-persist-cookie-template,omitempty"`
	SharedPartitionPersistDestinationIPTemplate int                              `json:"shared-partition-persist-destination-ip-template,omitempty"`
	SharedPartitionPersistSourceIPTemplate      int                              `json:"shared-partition-persist-source-ip-template,omitempty"`
	SharedPartitionPersistSslSidTemplate        int                              `json:"shared-partition-persist-ssl-sid-template,omitempty"`
	SharedPartitionPolicyTemplate               int                              `json:"shared-partition-policy-template,omitempty"`
	SharedPartitionPool                         int                              `json:"shared-partition-pool,omitempty"`
	SharedPartitionSMTPTemplate                 int                              `json:"shared-partition-smtp-template,omitempty"`
	SharedPartitionServerSslTemplate            int                              `json:"shared-partition-server-ssl-template,omitempty"`
	SharedPartitionSmppTemplate                 int                              `json:"shared-partition-smpp-template,omitempty"`
	SharedPartitionTCP                          int                              `json:"shared-partition-tcp,omitempty"`
	SharedPartitionTCPProxyTemplate             int                              `json:"shared-partition-tcp-proxy-template,omitempty"`
	SharedPartitionUDP                          int                              `json:"shared-partition-udp,omitempty"`
	SharedPartitionVirtualPortTemplate          int                              `json:"shared-partition-virtual-port-template,omitempty"`
	SkipRevHash                                 int                              `json:"skip-rev-hash,omitempty"`
	SnatOnVip                                   int                              `json:"snat-on-vip,omitempty"`
	StatsDataAction                             string                           `json:"stats-data-action,omitempty"`
	SubstituteSourceMac                         int                              `json:"substitute-source-mac,omitempty"`
	SupportHTTP2                                int                              `json:"support-http2,omitempty"`
	SynCookie                                   int                              `json:"syn-cookie,omitempty"`
	TemplateCache                               string                           `json:"template-cache,omitempty"`
	TemplateCacheShared                         string                           `json:"template-cache-shared,omitempty"`
	TemplateClientSSH                           string                           `json:"template-client-ssh,omitempty"`
	TemplateClientSsl                           string                           `json:"template-client-ssl,omitempty"`
	TemplateClientSslShared                     string                           `json:"template-client-ssl-shared,omitempty"`
	TemplateConnectionReuse                     string                           `json:"template-connection-reuse,omitempty"`
	TemplateConnectionReuseShared               string                           `json:"template-connection-reuse-shared,omitempty"`
	TemplateDNS                                 string                           `json:"template-dns,omitempty"`
	TemplateDNSShared                           string                           `json:"template-dns-shared,omitempty"`
	TemplateDblb                                string                           `json:"template-dblb,omitempty"`
	TemplateDblbShared                          string                           `json:"template-dblb-shared,omitempty"`
	TemplateDiameter                            string                           `json:"template-diameter,omitempty"`
	TemplateDiameterShared                      string                           `json:"template-diameter-shared,omitempty"`
	TemplateDoh                                 string                           `json:"template-doh,omitempty"`
	TemplateDohShared                           string                           `json:"template-doh-shared,omitempty"`
	TemplateDynamicService                      string                           `json:"template-dynamic-service,omitempty"`
	TemplateDynamicServiceShared                string                           `json:"template-dynamic-service-shared,omitempty"`
	TemplateExternalService                     string                           `json:"template-external-service,omitempty"`
	TemplateExternalServiceShared               string                           `json:"template-external-service-shared,omitempty"`
	TemplateFileInspection                      string                           `json:"template-file-inspection,omitempty"`
	TemplateFix                                 string                           `json:"template-fix,omitempty"`
	TemplateFixShared                           string                           `json:"template-fix-shared,omitempty"`
	TemplateFtp                                 string                           `json:"template-ftp,omitempty"`
	TemplateHTTP                                string                           `json:"template-http,omitempty"`
	TemplateHTTPPolicy                          string                           `json:"template-http-policy,omitempty"`
	TemplateHTTPPolicyShared                    string                           `json:"template-http-policy-shared,omitempty"`
	TemplateHTTPShared                          string                           `json:"template-http-shared,omitempty"`
	TemplateImapPop3                            string                           `json:"template-imap-pop3,omitempty"`
	TemplateImapPop3Shared                      string                           `json:"template-imap-pop3-shared,omitempty"`
	TemplateMqtt                                string                           `json:"template-mqtt,omitempty"`
	TemplatePersistCookie                       string                           `json:"template-persist-cookie,omitempty"`
	TemplatePersistCookieShared                 string                           `json:"template-persist-cookie-shared,omitempty"`
	TemplatePersistDestinationIP                string                           `json:"template-persist-destination-ip,omitempty"`
	TemplatePersistDestinationIPShared          string                           `json:"template-persist-destination-ip-shared,omitempty"`
	TemplatePersistSourceIP                     string                           `json:"template-persist-source-ip,omitempty"`
	TemplatePersistSourceIPShared               string                           `json:"template-persist-source-ip-shared,omitempty"`
	TemplatePersistSslSid                       string                           `json:"template-persist-ssl-sid,omitempty"`
	TemplatePersistSslSidShared                 string                           `json:"template-persist-ssl-sid-shared,omitempty"`
	TemplatePolicy                              string                           `json:"template-policy,omitempty"`
	TemplatePolicyShared                        string                           `json:"template-policy-shared,omitempty"`
	TemplateReqmodIcap                          string                           `json:"template-reqmod-icap,omitempty"`
	TemplateRespmodIcap                         string                           `json:"template-respmod-icap,omitempty"`
	TemplateSMTP                                string                           `json:"template-smtp,omitempty"`
	TemplateSMTPShared                          string                           `json:"template-smtp-shared,omitempty"`
	TemplateScaleout                            string                           `json:"template-scaleout,omitempty"`
	TemplateServerSSH                           string                           `json:"template-server-ssh,omitempty"`
	TemplateServerSsl                           string                           `json:"template-server-ssl,omitempty"`
	TemplateServerSslShared                     string                           `json:"template-server-ssl-shared,omitempty"`
	TemplateSip                                 string                           `json:"template-sip,omitempty"`
	TemplateSipShared                           string                           `json:"template-sip-shared,omitempty"`
	TemplateSmpp                                string                           `json:"template-smpp,omitempty"`
	TemplateSmppShared                          string                           `json:"template-smpp-shared,omitempty"`
	TemplateSsli                                string                           `json:"template-ssli,omitempty"`
	TemplateTCP                                 string                           `json:"template-tcp,omitempty"`
	TemplateTCPProxy                            string                           `json:"template-tcp-proxy,omitempty"`
	TemplateTCPProxyClient                      string                           `json:"template-tcp-proxy-client,omitempty"`
	TemplateTCPProxyServer                      string                           `json:"template-tcp-proxy-server,omitempty"`
	TemplateTCPProxyShared                      string                           `json:"template-tcp-proxy-shared,omitempty"`
	TemplateTCPShared                           string                           `json:"template-tcp-shared,omitempty"`
	TemplateUDP                                 string                           `json:"template-udp,omitempty"`
	TemplateUDPShared                           string                           `json:"template-udp-shared,omitempty"`
	TemplateVirtualPort                         string                           `json:"template-virtual-port,omitempty"`
	TemplateVirtualPortShared                   string                           `json:"template-virtual-port-shared,omitempty"`
	TrunkFwd                                    int                              `json:"trunk-fwd,omitempty"`
	TrunkRev                                    int                              `json:"trunk-rev,omitempty"`
	UUID                                        string                           `json:"uuid,omitempty"`
	UseAlternatePort                            int                              `json:"use-alternate-port,omitempty"`
	UseCgnv6                                    int                              `json:"use-cgnv6,omitempty"`
	UseDefaultIfNoServer                        int                              `json:"use-default-if-no-server,omitempty"`
	UseRcvHopForResp                            int                              `json:"use-rcv-hop-for-resp,omitempty"`
	UserTag                                     string                           `json:"user-tag,omitempty"`
	View                                        int                              `json:"view,omitempty"`
	WafTemplate                                 string                           `json:"waf-template,omitempty"`
	WhenDown                                    int                              `json:"when-down,omitempty"`
	WhenDownProtocol2                           int                              `json:"when-down-protocol2,omitempty"`
}

type SlbVirtualServerAclIDList struct {
	AclID                  int    `json:"acl-id,omitempty"`
	AclIDSeqNum            int    `json:"acl-id-seq-num,omitempty"`
	AclIDSeqNumShared      int    `json:"acl-id-seq-num-shared,omitempty"`
	AclIDShared            int    `json:"acl-id-shared,omitempty"`
	AclIDSrcNatPool        string `json:"acl-id-src-nat-pool,omitempty"`
	AclIDSrcNatPoolShared  string `json:"acl-id-src-nat-pool-shared,omitempty"`
	SharedPartitionPoolID  int    `json:"shared-partition-pool-id,omitempty"`
	VAclIDSeqNum           int    `json:"v-acl-id-seq-num,omitempty"`
	VAclIDSeqNumShared     int    `json:"v-acl-id-seq-num-shared,omitempty"`
	VAclIDSrcNatPool       string `json:"v-acl-id-src-nat-pool,omitempty"`
	VAclIDSrcNatPoolShared string `json:"v-acl-id-src-nat-pool-shared,omitempty"`
	VSharedPartitionPoolID int    `json:"v-shared-partition-pool-id,omitempty"`
}

type SlbVirtualServerAclNameList struct {
	AclName                  string `json:"acl-name,omitempty"`
	AclNameSeqNum            int    `json:"acl-name-seq-num,omitempty"`
	AclNameSeqNumShared      int    `json:"acl-name-seq-num-shared,omitempty"`
	AclNameShared            string `json:"acl-name-shared,omitempty"`
	AclNameSrcNatPool        string `json:"acl-name-src-nat-pool,omitempty"`
	AclNameSrcNatPoolShared  string `json:"acl-name-src-nat-pool-shared,omitempty"`
	SharedPartitionPoolName  int    `json:"shared-partition-pool-name,omitempty"`
	VAclNameSeqNum           int    `json:"v-acl-name-seq-num,omitempty"`
	VAclNameSeqNumShared     int    `json:"v-acl-name-seq-num-shared,omitempty"`
	VAclNameSrcNatPool       string `json:"v-acl-name-src-nat-pool,omitempty"`
	VAclNameSrcNatPoolShared string `json:"v-acl-name-src-nat-pool-shared,omitempty"`
	VSharedPartitionPoolName int    `json:"v-shared-partition-pool-name,omitempty"`
}

type SlbVirtualServerAflexScripts struct {
	Aflex       string `json:"aflex,omitempty"`
	AflexShared string `json:"aflex-shared,omitempty"`
}

type SlbVirtualServerAuthCfg struct {
	AaaPolicy string `json:"aaa-policy,omitempty"`
}

type SlbVirtualServerSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostSlbVirtualServerPort(id string, name string, inst SlbVirtualServerPort, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbVirtualServerPort")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/virtual-server/"+name+"/port", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbVirtualServerPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbVirtualServerPort REQ RES..........................", m)
			err := check_api_status("PostSlbVirtualServerPort", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbVirtualServerPort(id string, name1 string, name2 string, name3 string, host string) (*SlbVirtualServerPort, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbVirtualServerPort")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/virtual-server/"+name1+"/port/"+name2+"+"+name3, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbVirtualServerPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbVirtualServerPort REQ RES..........................", m)
			err := check_api_status("GetSlbVirtualServerPort", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbVirtualServerPort(id string, name1 string, name2 string, name3 string, inst SlbVirtualServerPort, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbVirtualServerPort")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/virtual-server/"+name1+"/port/"+name2+"+"+name3, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbVirtualServerPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbVirtualServerPort REQ RES..........................", m)
			err := check_api_status("PutSlbVirtualServerPort", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbVirtualServerPort(id string, name1 string, name2 string, name3 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbVirtualServerPort")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/virtual-server/"+name1+"/port/"+name2+"+"+name3, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbVirtualServerPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
