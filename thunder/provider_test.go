package thunder

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var TEST_PARTITION = "Common"

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

var testProviders = map[string]terraform.ResourceProvider{
	"thunder": Provider(),
}

func init() {

	testAccProvider = ProviderTest().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"thunder": testAccProvider,
	}

}

func ProviderTest() terraform.ResourceProvider {
	return &schema.Provider{

		ResourcesMap: map[string]*schema.Resource{
			"thunder_virtual_server":                    resourceVirtualServer(),
			"thunder_service_group":                     resourceServiceGroup(),
			"thunder_server":                            resourceServer(),
			"thunder_ethernet":                          resourceEthernet(),
			"thunder_rib_route":                         resourceRibRoute(),
			"thunder_vrrp_vrid":                         resourceVrrpVrid(),
			"thunder_glm":                               resourceGlm(),
			"thunder_import":                            resourceImport(),
			"thunder_vrrp_common":                       resourceVrrpCommon(),
			"thunder_vrrp_peer_group":                   resourceVrrpPeerGroup(),
			"thunder_vrrp_session_sync":                 resourceVrrpSessionSync(),
			"thunder_dns_primary":                       resourceDnsPrimary(),
			"thunder_harmony_controller_profile":        resourceHarmonyControllerProfile(),
			"thunder_overlay_tunnel_options":            resourceOverlayTunnelOptions(),
			"thunder_overlay_tunnel_vtep":               resourceOverlayTunnelVtep(),
			"thunder_partition":                         resourceOverlayTunnelPartition(),
			"thunder_slb_template_tcp":                  resourceSlbTemplateTcp(),
			"thunder_slb_template_udp":                  resourceSlbTemplateUdp(),
			"thunder_slb_template_smpp":                 resourceSlbTemplateSmpp(),
			"thunder_slb_template_fix":                  resourceTemplateFix(),
			"thunder_slb_template_ftp":                  resourceSlbTemplateFTP(),
			"thunder_slb_template_mqtt":                 resourceSlbTemplateMqtt(),
			"thunder_slb_template_http_policy":          resourceSlbTemplateHttpPolicy(),
			"thunder_slb_template_ssli":                 resourceTemplateSSLI(),
			"thunder_slb_template_cipher":               resourceTemplateCipher(),
			"thunder_slb_template_imap_pop3":            resourceTemplateImap_POP3(),
			"thunder_slb_template_dns":                  resourceTemplateDNS(),
			"thunder_slb_template_connection_reuse":     resourceTemplateConnectionReuse(),
			"thunder_slb_template_client_ssh":           resourceTemplateClientSsh(),
			"thunder_slb_template_dblb":                 resourceTemplateDBLB(),
			"thunder_slb_template_port":                 resourceTemplatePort(),
			"thunder_slb_template_smtp":                 resourceSlbTemplateSMTP(),
			"thunder_slb_template_tcp_proxy":            resourceSlbTemplateTcpProxy(),
			"thunder_slb_template_server":               resourceSlbTemplateServer(),
			"thunder_slb_template_virtual_port":         resourceSlbTemplateVirtualPort(),
			"thunder_slb_template_diameter":             resourceSlbTemplateDiameter(),
			"thunder_slb_template_cache":                resourceSlbTemplateCache(),
			"thunder_slb_template_csv":                  resourceSlbTemplateCSV(),
			"thunder_slb_template_reqmod_icap":          resourceSlbTemplateReqmodIcap(),
			"thunder_slb_template_respmod_icap":         resourceSlbTemplateRespmodIcap(),
			"thunder_slb_template_dynamic_service":      resourceSlbTemplateDynamicService(),
			"thunder_slb_template_server_ssh":           resourceSlbTemplateServerSSH(),
			"thunder_slb_template_policy":               resourceSlbTemplatePolicy(),
			"thunder_slb_template_server_ssl":           resourceSlbTemplateServerSSL(),
			"thunder_slb_template_client_ssl":           resourceSlbTemplateClientSSL(),
			"thunder_slb_template_sip":                  resourceSlbTemplateSIP(),
			"thunder_slb_template_http":                 resourceSlbTemplateHTTP(),
			"thunder_slb_template_monitor":              resourceSlbTemplateMonitor(),
			"thunder_slb_template_logging":              resourceSlbTemplateLogging(),
			"thunder_slb_template_snmp":                 resourceSlbTemplateSNMP(),
			"thunder_slb_template_external_service":     resourceSlbTemplateExternalService(),
			"thunder_slb_template_virtual_server":       resourceSlbTemplateVirtualServer(),
			"thunder_slb_aflow":                         resourceSlbAflow(),
			"thunder_slb_common":                        resourceSlbCommon(),
			"thunder_slb_common_conn_rate_limit_src_ip": resourceSlbCommonConnRateLimitSrcIP(),
			"thunder_slb_connection_reuse":              resourceSlbConnectionReuse(),
			"thunder_slb_crl_srcip":                     resourceSlbCrlSrcip(),
			"thunder_slb_dns":                           resourceSlbDNS(),
			"thunder_slb_sip":                           resourceSlbSip(),
			"thunder_slb_smpp":                          resourceSlbSmpp(),
			"thunder_slb_smtp":                          resourceSlbSmtp(),
			"thunder_slb_switch":                        resourceSlbSwitch(),
			"thunder_slb_spdy_proxy":                    resourceSlbSpdyProxy(),
			"thunder_slb_svm_source_nat":                resourceSlbSvmSourceNat(),
			"thunder_slb_sport_rate_limit":              resourceSlbSportRateLimit(),
			"thunder_slb_transparent_acl_template":      resourceSlbTransparentAclTemplate(),
			"thunder_slb_transparent_tcp_template":      resourceSlbTransperentTcpTemplate(),
			"thunder_slb_dns_cache":                     resourceSlbDNSCache(),
			"thunder_slb_dns_response_rate_limiting":    resourceSlbDNSResponseRateLimiting(),
			"thunder_slb_fast_http_proxy":               resourceSlbFastHttpProxy(),
			"thunder_slb_fix":                           resourceSlbFix(),
			"thunder_slb_ftp_ctl":                       resourceSlbFTPCtl(),
			"thunder_slb_ftp_data":                      resourceSlbFTPData(),
			"thunder_slb_ftp_proxy":                     resourceSlbFTPProxy(),
			"thunder_slb_generic_proxy":                 resourceSlbGenericProxy(),
			"thunder_slb_passthrough":                   resourceSlbPassthrough(),
			"thunder_slb_perf":                          resourceSlbPerf(),
			"thunder_slb_persist":                       resourceSlbpersist(),
			"thunder_slb_player_id_global":              resourceSlbPlayerIdGlobal(),
			"thunder_slb_pop3_proxy":                    resourceSlbPop3Proxy(),
			"thunder_slb_proxy":                         resourceSlbProxy(),
			"thunder_slb_rate_limit_log":                resourceSlbRateLimitLog(),
			"thunder_slb_rc_cache_global":               resourceSlbRcCacheGlobal(),
			"thunder_slb_resource_usage":                resourceSlbResourceUsage(),
			"thunder_slb_health_gateway":                resourceSlbHealthGateway(),
			"thunder_slb_health_stat":                   resourceSlbHealthStat(),
			"thunder_slb_http_proxy":                    resourceSlbHttpProxy(),
			"thunder_slb_http2":                         resourceSlbHTTP2(),
			"thunder_slb_hw_compress":                   resourceSlbHwCompress(),
			"thunder_slb_ssl_cert_revoke":               resourceSlbSSLCertRevoke(),
			"thunder_slb_ssl_forward_proxy":             resourceSlbSSLForwardProxy(),
			"thunder_slb_icap_http":                     resourceSlbIcapHTTP(),
			"thunder_slb_icap":                          resourceSlbIcap(),
			"thunder_slb_imapproxy":                     resourceSlbImapproxy(),
			"thunder_slb_l4":                            resourceSlL4(),
			"thunder_slb_l7session":                     resourceSlL7session(),
			"thunder_slb_mlb":                           resourceSlbMlb(),
			"thunder_slb_mssql":                         resourceSlbMssql(),
			"thunder_slb_mysql":                         resourceSlbMysql(),
			"thunder_slb_ssl_expire_check":              resourceSlbSSLExpireCheck(),
			"thunder_ip_dns_primary":                    resourceIpDnsPrimary(),
			"thunder_ip_dns_secondary":                  resourceIpDnsSecondary(),
			"thunder_ip_dns_suffix":                     resourceIpDnsSuffix(),
			"thunder_ip_frag":                           resourceIpFrag(),
			"thunder_ip_icmp":                           resourceIpIcmp(),
			"thunder_ip_nat_alg_pptp":                   resourceIpNatAlgPptp(),
			"thunder_ip_nat_global":                     resourceIpNatGlobal(),
			"thunder_ip_nat_icmp":                       resourceIpNatIcmp(),
			"thunder_ip_prefix_list":                    resourceIpPrefixList(),
			"thunder_ip_reroute":                        resourceIpReroute(),
			"thunder_ip_tcp":                            resourceIpTcp(),
			"thunder_ipv6_frag":                         resourceIpv6Frag(),
			"thunder_ipv6_icmpv6":                       resourceIpv6Icmpv6(),
			"thunder_ipv6_nat_icmpv6":                   resourceIpv6NatIcmpv6(),
			"thunder_ip_route_static_bfd":               resourceIPRouteStaticBfd(),
			"thunder_ip_address":                        resourceIPAddress(),
			"thunder_interface_management":              resourceInterfaceManagement(),
			"thunder_interface_ethernet_bfd":            resourceInterfaceEthernetBFD(),
			"thunder_interface_ethernet_lldp":           resourceInterfaceEthernetLLDP(),
			"thunder_interface_ve_bfd":                  resourceInterfaceVeBFD(),
			"thunder_interface_ethernet_trunk_group":    resourceInterfaceEthernetTrunkGroup(),
			"thunder_interface_ethernet_ipv6":           resourceInterfaceEthernetIPv6(),
			"thunder_interface_ve_ip":                   resourceInterfaceVeIP(),
			"thunder_interface_ve_ipv6":                 resourceInterfaceVeIPv6(),
			"thunder_interface_ethernet":                resourceInterfaceEthernet(),
			"thunder_interface_ve":                      resourceInterfaceVE(),
			"thunder_fw_alg_icmp":                       resourceFwAlgIcmp(),
			"thunder_fw_active_rule_set":                resourceFwActiveRuleSet(),
			"thunder_fw_alg_dns":                        resourceFwAlgDns(),
			"thunder_fw_alg_ftp":                        resourceFwAlgFtp(),
			"thunder_fw_alg_pptp":                       resourceFwAlgPptp(),
			"thunder_fw_alg_rtsp":                       resourceFwAlgRtsp(),
			"thunder_fw_alg_sip":                        resourceFwAlgSip(),
			"thunder_fw_alg_tftp":                       resourceFwAlgTftp(),
			"thunder_fw_apply_changes":                  resourceFwApplyChanges(),
			"thunder_fw_app":                            resourceFwApp(),
			"thunder_fw_global":                         resourceFwGlobal(),
			"thunder_fw_gtp":                            resourceFwGtp(),
			"thunder_fw_gtp_in_gtp_filtering":           resourceFwGtpInGtpFiltering(),
			"thunder_fw_gtp_v0":                         resourceFwGtpV0(),
			"thunder_fw_helper_sessions":                resourceFwHelperSessions(),
			"thunder_fw_limit_entry":                    resourceFwLimitEntry(),
			"thunder_fw_local_log":                      resourceFwLocalLog(),
			"thunder_fw_logging":                        resourceFwLogging(),
			"thunder_fw_radius_server":                  resourceFwRadiusServer(),
			"thunder_fw_resource_usage":                 resourceFwResourceUsage(),
			"thunder_fw_clear_session_filter":           resourceFwClearSessionFilter(),
			"thunder_fw_full_cone_session":              resourceFwFullConeSession(),
			"thunder_fw_server":                         resourceFwServer(),
			"thunder_fw_service_group":                  resourceFwServiceGroup(),
			"thunder_fw_status":                         resourceFwStatus(),
			"thunder_fw_system_status":                  resourceFwSystemStatus(),
			"thunder_fw_tap_monitor":                    resourceFwTapMonitor(),
			"thunder_fw_tcp_mss_clamp":                  resourceFwTcpMssClamp(),
			"thunder_fw_tcp_reset_on_error":             resourceFwTcpResetOnError(),
			"thunder_fw_tcp_rst_close_immediate":        resourceFwTcpRstCloseImmediate(),
			"thunder_fw_tcp_window_check":               resourceFwTcpWindowCheck(),
			"thunder_fw_top_k_rules":                    resourceFwTopKRules(),
			"thunder_fw_urpf":                           resourceFwUrpf(),
			"thunder_fw_vrid":                           resourceFwVrid(),
			"thunder_fw_template_logging":               resourceFwTemplateLogging(),
			"thunder_glm_send":                          resourceGlmSend(),
			"thunder_write_memory":                      resourceWriteMemory(),
			"thunder_hostname":                          resourceHostname(),
			"thunder_ip_nat_pool":                       resourceIpNatPool(),
			"thunder_slb_server_port":                   resourceSlbServerPort(),
			"thunder_slb_virtual_server_port":           resourceSlbVirtualServerPort(),
			"thunder_slb_template_persist_cookie":       resourceSlbTemplatePersistCookie(),
			"thunder_slb_template_persist_source_ip":    resourceSlbTemplatePersistSourceIp(),
		},

		ConfigureFunc: providerConfigureTest,
	}
}

func providerConfigureTest(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Address:  os.Getenv("THUNDER_HOST"),     //hostname
		Username: os.Getenv("THUNDER_USER"),     //username
		Password: os.Getenv("THUNDER_PASSWORD"), //password
	}

	return config.Client()
}

func TestAccProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAcctPreCheck(t *testing.T) {
	if v := os.Getenv("THUNDER_HOST"); v == "" {
		t.Fatal("ADDRESS must be set for acceptance tests")
	}
	if v := os.Getenv("THUNDER_USER"); v == "" {
		t.Fatal("USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("THUNDER_PASSWORD"); v == "" {
		t.Fatal("PASSWORD must be set for acceptance tests")
	}
}