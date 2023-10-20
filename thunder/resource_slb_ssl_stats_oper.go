package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslStatsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_stats_oper`: Operational Status for the object ssl-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslStatsOperRead,
		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_module_type": {
							Type: schema.TypeString, Optional: true, Description: "SSL module",
						},
						"config_module_type": {
							Type: schema.TypeString, Optional: true, Description: "Config module",
						},
						"thales_hsm_status": {
							Type: schema.TypeString, Optional: true, Description: "Thales HSM Status",
						},
						"ssl_modules_count": {
							Type: schema.TypeInt, Optional: true, Description: "Number of SSL modules",
						},
						"current_clientside_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Current clientside SSL connections",
						},
						"total_clientside_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total clientside SSL connections",
						},
						"current_serverside_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Current serverside SSL connections",
						},
						"total_serverside_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total serverside SSL connections",
						},
						"non_ssl_bypass_connections": {
							Type: schema.TypeInt, Optional: true, Description: "NON SSL bypass connections",
						},
						"total_reuse_client_ssl": {
							Type: schema.TypeInt, Optional: true, Description: "Total times of stateful session reuse in client ssl",
						},
						"total_reuse_server_ssl": {
							Type: schema.TypeInt, Optional: true, Description: "Total times of stateful session reuse in server ssl",
						},
						"total_reuse_session_ticket_client_ssl": {
							Type: schema.TypeInt, Optional: true, Description: "Total times of stateless session reuse in client ssl",
						},
						"total_reuse_session_ticket_server_ssl": {
							Type: schema.TypeInt, Optional: true, Description: "Total times of stateless session reuse in server ssl",
						},
						"total_clientside_early_data_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total clientside early data connections",
						},
						"total_serverside_early_data_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total serverside early data connections",
						},
						"total_clientside_failed_early_data_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total clientside failed early data connections",
						},
						"total_serverside_failed_early_data_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total serverside failed early data connections",
						},
						"failed_handshakes": {
							Type: schema.TypeInt, Optional: true, Description: "Failed SSL handshakes",
						},
						"failed_crypto": {
							Type: schema.TypeInt, Optional: true, Description: "Failed crypto operations",
						},
						"ssl_memory_usage": {
							Type: schema.TypeInt, Optional: true, Description: "SSL memory usage",
						},
						"server_cert_errors": {
							Type: schema.TypeInt, Optional: true, Description: "SSL server certificate errors",
						},
						"client_cert_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "SSL client certificate authorization failed",
						},
						"ca_verification_failures": {
							Type: schema.TypeInt, Optional: true, Description: "SSL fail CA verification",
						},
						"hw_context_total": {
							Type: schema.TypeInt, Optional: true, Description: "HW Context Memory Total Count",
						},
						"hw_context_usage": {
							Type: schema.TypeInt, Optional: true, Description: "HW Context Memory In Use",
						},
						"hw_context_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "HW Context Memory alloc failed",
						},
						"hw_ring_full": {
							Type: schema.TypeInt, Optional: true, Description: "HW ring full",
						},
						"record_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "Record too big",
						},
						"clientssl_context_malloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Total client ssl context malloc failures",
						},
						"max_ssl_contexts": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum SSL contexts",
						},
						"curr_ssl_contexts": {
							Type: schema.TypeInt, Optional: true, Description: "Current SSL contexts in use",
						},
						"static_ssl_contexts": {
							Type: schema.TypeInt, Optional: true, Description: "Current Static SSL contexts",
						},
						"dynamic_ssl_contexts": {
							Type: schema.TypeInt, Optional: true, Description: "Current Dynamic SSL contexts",
						},
						"bypass_failsafe_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass failsafe SSL sessions",
						},
						"bypass_usename_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Username sessions",
						},
						"bypass_ad_group_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass AD-Group sessions",
						},
						"bypass_sni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass SNI sessions",
						},
						"bypass_cert_subject_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Cert Subject sessions",
						},
						"bypass_cert_issuer_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Cert Issuer sessions",
						},
						"bypass_cert_san_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Cert SAN sessions",
						},
						"bypass_no_sni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass No SNI sessions",
						},
						"reset_no_sni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Reset NO SNI sessions",
						},
						"bypass_esni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass ESNI sessions",
						},
						"drop_esni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Drop ESNI sessions",
						},
						"bypass_client_auth_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Client Auth sessions",
						},
						"ssl_handshake_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Failed in SSL handshakes",
						},
						"cryptio_operations": {
							Type: schema.TypeInt, Optional: true, Description: "Crypto operations",
						},
						"tcp_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed in TCP",
						},
						"certificate_verification_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed in Certificate Verification",
						},
						"certificate_signing_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Failed in Certificate Signing",
						},
						"invalid_ocsp_stapling_response": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid OCSP Stapling Response",
						},
						"revoked_ocsp_response": {
							Type: schema.TypeInt, Optional: true, Description: "Revoked OCSP Response",
						},
						"unsupported_ssl_version": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported SSL Version",
						},
						"client_ssl_fatal_alert": {
							Type: schema.TypeInt, Optional: true, Description: "Cert Fetch, SSL Fatal Alert",
						},
						"client_ssl_fin_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Cert Fetch, TCP FIN/RST",
						},
						"ssl_session_fin_reset": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Session, TCP FIN/RST",
						},
						"server_ssl_fatal_alert": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL, SSL Fatal alert",
						},
						"server_ssl_fin_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL, TCP FIN/RST",
						},
						"client_ssl_internal_error": {
							Type: schema.TypeInt, Optional: true, Description: "Client SSL Internal Error",
						},
						"client_ssl_unknown_error": {
							Type: schema.TypeInt, Optional: true, Description: "Client SSL Unknown Error",
						},
						"server_ssl_internal_error": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL Internal Error",
						},
						"server_ssl_unknown_error": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL Unknown Error",
						},
						"fp_fail_handshake": {
							Type: schema.TypeInt, Optional: true, Description: "Cert fetch, Fatal alert",
						},
						"fp_fail_tcp_error": {
							Type: schema.TypeInt, Optional: true, Description: "Cert fetch, TCP FIN/RST",
						},
						"fp_fail_verify_cert": {
							Type: schema.TypeInt, Optional: true, Description: "Cert fetch, validation error",
						},
						"fp_fail_aflex": {
							Type: schema.TypeInt, Optional: true, Description: "aFleX fail",
						},
						"util_enable_status": {
							Type: schema.TypeString, Optional: true, Description: "current module percentage enable Status",
						},
						"ssl_module_stats": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssl_modules_index": {
										Type: schema.TypeInt, Optional: true, Description: "SSL module index",
									},
									"total_enabled_crypto_engines": {
										Type: schema.TypeInt, Optional: true, Description: "Number of enabled crypto engines",
									},
									"total_available_crypto_engines": {
										Type: schema.TypeInt, Optional: true, Description: "Number of available crypto engines",
									},
									"requests_handled": {
										Type: schema.TypeInt, Optional: true, Description: "Number of requests handled",
									},
									"util_percentage": {
										Type: schema.TypeInt, Optional: true, Description: "current module percentage per sec",
									},
									"1sec_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Average module percentage per sec",
									},
									"5sec_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Average module percentage per 5 sec",
									},
									"10sec_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Average module percentage per 10 sec",
									},
									"30sec_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Average module percentage per 30 sec",
									},
									"60sec_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Average module percentage per 60 sec",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslStatsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslStatsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslStatsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		items := setObjectSlbSslStatsOperOper(res)
		d.SetId(obj.GetId())
		d.Set("oper", items)

		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslStatsOperOper(res edpt.SlbSslStatss) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ssl_module_type"] = res.DataSlbSslStats.Oper.SslModuleType
	in["config_module_type"] = res.DataSlbSslStats.Oper.ConfigModuleType
	in["thales_hsm_status"] = res.DataSlbSslStats.Oper.ThalesHsmStatus
	in["ssl_modules_count"] = res.DataSlbSslStats.Oper.SslModulesCount
	in["current_clientside_connections"] = res.DataSlbSslStats.Oper.CurrentClientsideConnections
	in["total_clientside_connections"] = res.DataSlbSslStats.Oper.TotalClientsideConnections
	in["current_serverside_connections"] = res.DataSlbSslStats.Oper.CurrentServersideConnections
	in["total_serverside_connections"] = res.DataSlbSslStats.Oper.TotalServersideConnections
	in["non_ssl_bypass_connections"] = res.DataSlbSslStats.Oper.NonSslBypassConnections
	in["total_reuse_client_ssl"] = res.DataSlbSslStats.Oper.TotalReuseClientSsl
	in["total_reuse_server_ssl"] = res.DataSlbSslStats.Oper.TotalReuseServerSsl
	in["total_reuse_session_ticket_client_ssl"] = res.DataSlbSslStats.Oper.TotalReuseSessionTicketClientSsl
	in["total_reuse_session_ticket_server_ssl"] = res.DataSlbSslStats.Oper.TotalReuseSessionTicketServerSsl
	in["total_clientside_early_data_connections"] = res.DataSlbSslStats.Oper.TotalClientsideEarlyDataConnections
	in["total_serverside_early_data_connections"] = res.DataSlbSslStats.Oper.TotalServersideEarlyDataConnections
	in["total_clientside_failed_early_data_connections"] = res.DataSlbSslStats.Oper.TotalClientsideFailedEarlyDataConnections
	in["total_serverside_failed_early_data_connections"] = res.DataSlbSslStats.Oper.TotalServersideFailedEarlyDataConnections
	in["failed_handshakes"] = res.DataSlbSslStats.Oper.FailedHandshakes
	in["failed_crypto"] = res.DataSlbSslStats.Oper.FailedCrypto
	in["ssl_memory_usage"] = res.DataSlbSslStats.Oper.SslMemoryUsage
	in["server_cert_errors"] = res.DataSlbSslStats.Oper.ServerCertErrors
	in["client_cert_auth_fail"] = res.DataSlbSslStats.Oper.ClientCertAuthFail
	in["ca_verification_failures"] = res.DataSlbSslStats.Oper.CaVerificationFailures
	in["hw_context_total"] = res.DataSlbSslStats.Oper.HwContextTotal
	in["hw_context_usage"] = res.DataSlbSslStats.Oper.HwContextUsage
	in["hw_context_alloc_fail"] = res.DataSlbSslStats.Oper.HwContextAllocFail
	in["hw_ring_full"] = res.DataSlbSslStats.Oper.HwRingFull
	in["record_too_big"] = res.DataSlbSslStats.Oper.RecordTooBig
	in["clientssl_context_malloc_fail"] = res.DataSlbSslStats.Oper.ClientsslContextMallocFail
	in["max_ssl_contexts"] = res.DataSlbSslStats.Oper.MaxSslContexts
	in["curr_ssl_contexts"] = res.DataSlbSslStats.Oper.CurrSslContexts
	in["static_ssl_contexts"] = res.DataSlbSslStats.Oper.StaticSslContexts
	in["dynamic_ssl_contexts"] = res.DataSlbSslStats.Oper.DynamicSslContexts
	in["bypass_failsafe_sessions"] = res.DataSlbSslStats.Oper.BypassFailsafeSessions
	in["bypass_usename_sessions"] = res.DataSlbSslStats.Oper.BypassUsenameSessions
	in["bypass_ad_group_sessions"] = res.DataSlbSslStats.Oper.BypassAdGroupSessions
	in["bypass_sni_sessions"] = res.DataSlbSslStats.Oper.BypassSniSessions
	in["bypass_cert_subject_sessions"] = res.DataSlbSslStats.Oper.BypassCertSubjectSessions
	in["bypass_cert_issuer_sessions"] = res.DataSlbSslStats.Oper.BypassCertIssuerSessions
	in["bypass_cert_san_sessions"] = res.DataSlbSslStats.Oper.BypassCertSanSessions
	in["bypass_no_sni_sessions"] = res.DataSlbSslStats.Oper.BypassNoSniSessions
	in["reset_no_sni_sessions"] = res.DataSlbSslStats.Oper.ResetNoSniSessions
	in["bypass_esni_sessions"] = res.DataSlbSslStats.Oper.BypassEsniSessions
	in["drop_esni_sessions"] = res.DataSlbSslStats.Oper.DropEsniSessions
	in["bypass_client_auth_sessions"] = res.DataSlbSslStats.Oper.BypassClientAuthSessions
	in["ssl_handshake_failure"] = res.DataSlbSslStats.Oper.SslHandshakeFailure
	in["cryptio_operations"] = res.DataSlbSslStats.Oper.CryptioOperations
	in["tcp_failures"] = res.DataSlbSslStats.Oper.TcpFailures
	in["certificate_verification_failures"] = res.DataSlbSslStats.Oper.CertificateVerificationFailures
	in["certificate_signing_failures"] = res.DataSlbSslStats.Oper.CertificateSigningFailures
	in["invalid_ocsp_stapling_response"] = res.DataSlbSslStats.Oper.InvalidOcspStaplingResponse
	in["revoked_ocsp_response"] = res.DataSlbSslStats.Oper.RevokedOcspResponse
	in["unsupported_ssl_version"] = res.DataSlbSslStats.Oper.UnsupportedSslVersion
	in["client_ssl_fatal_alert"] = res.DataSlbSslStats.Oper.ClientSslFatalAlert
	in["client_ssl_fin_reset"] = res.DataSlbSslStats.Oper.ClientSslFinReset
	in["ssl_session_fin_reset"] = res.DataSlbSslStats.Oper.SslSessionFinReset
	in["server_ssl_fatal_alert"] = res.DataSlbSslStats.Oper.ServerSslFatalAlert
	in["server_ssl_fin_reset"] = res.DataSlbSslStats.Oper.ServerSslFinReset
	in["client_ssl_internal_error"] = res.DataSlbSslStats.Oper.ClientSslInternalError
	in["client_ssl_unknown_error"] = res.DataSlbSslStats.Oper.ClientSslUnknownError
	in["server_ssl_internal_error"] = res.DataSlbSslStats.Oper.ServerSslInternalError
	in["server_ssl_unknown_error"] = res.DataSlbSslStats.Oper.ServerSslUnknownError
	in["fp_fail_handshake"] = res.DataSlbSslStats.Oper.FpFailHandshake
	in["fp_fail_tcp_error"] = res.DataSlbSslStats.Oper.FpFailTcpError
	in["fp_fail_verify_cert"] = res.DataSlbSslStats.Oper.FpFailVerifyCert
	in["fp_fail_aflex"] = res.DataSlbSslStats.Oper.FpFailAflex
	in["util_enable_status"] = res.DataSlbSslStats.Oper.UtilEnableStatus
	in["ssl_module_stats"] = setSliceSlbSslStatsOperOperSslModuleStats(res.DataSlbSslStats.Oper.SslModuleStats)

	result = append(result, in)
	return result
}

func setSliceSlbSslStatsOperOperSslModuleStats(d []edpt.SlbSslStatsOperOperSslModuleStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ssl_modules_index"] = item.SslModulesIndex
		in["total_enabled_crypto_engines"] = item.TotalEnabledCryptoEngines
		in["total_available_crypto_engines"] = item.TotalAvailableCryptoEngines
		in["requests_handled"] = item.RequestsHandled
		in["util_percentage"] = item.UtilPercentage
		in["1sec_rate"] = item.Sec1Rate
		in["5sec_rate"] = item.Sec5Rate
		in["10sec_rate"] = item.Sec10Rate
		in["30sec_rate"] = item.Sec30Rate
		in["60sec_rate"] = item.Sec60Rate

		result = append(result, in)
	}
	return result
}

func getObjectSlbSslStatsOperOper(d []interface{}) edpt.SlbSslStatsOperOper {
	count := len(d)
	var ret edpt.SlbSslStatsOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.SslModuleType = in["ssl_module_type"].(string)
		ret.ConfigModuleType = in["config_module_type"].(string)
		ret.ThalesHsmStatus = in["thales_hsm_status"].(string)
		ret.SslModulesCount = in["ssl_modules_count"].(int)
		ret.CurrentClientsideConnections = in["current_clientside_connections"].(int)
		ret.TotalClientsideConnections = in["total_clientside_connections"].(int)
		ret.CurrentServersideConnections = in["current_serverside_connections"].(int)
		ret.TotalServersideConnections = in["total_serverside_connections"].(int)
		ret.NonSslBypassConnections = in["non_ssl_bypass_connections"].(int)
		ret.TotalReuseClientSsl = in["total_reuse_client_ssl"].(int)
		ret.TotalReuseServerSsl = in["total_reuse_server_ssl"].(int)
		ret.TotalReuseSessionTicketClientSsl = in["total_reuse_session_ticket_client_ssl"].(int)
		ret.TotalReuseSessionTicketServerSsl = in["total_reuse_session_ticket_server_ssl"].(int)
		ret.TotalClientsideEarlyDataConnections = in["total_clientside_early_data_connections"].(int)
		ret.TotalServersideEarlyDataConnections = in["total_serverside_early_data_connections"].(int)
		ret.TotalClientsideFailedEarlyDataConnections = in["total_clientside_failed_early_data_connections"].(int)
		ret.TotalServersideFailedEarlyDataConnections = in["total_serverside_failed_early_data_connections"].(int)
		ret.FailedHandshakes = in["failed_handshakes"].(int)
		ret.FailedCrypto = in["failed_crypto"].(int)
		ret.SslMemoryUsage = in["ssl_memory_usage"].(int)
		ret.ServerCertErrors = in["server_cert_errors"].(int)
		ret.ClientCertAuthFail = in["client_cert_auth_fail"].(int)
		ret.CaVerificationFailures = in["ca_verification_failures"].(int)
		ret.HwContextTotal = in["hw_context_total"].(int)
		ret.HwContextUsage = in["hw_context_usage"].(int)
		ret.HwContextAllocFail = in["hw_context_alloc_fail"].(int)
		ret.HwRingFull = in["hw_ring_full"].(int)
		ret.RecordTooBig = in["record_too_big"].(int)
		ret.ClientsslContextMallocFail = in["clientssl_context_malloc_fail"].(int)
		ret.MaxSslContexts = in["max_ssl_contexts"].(int)
		ret.CurrSslContexts = in["curr_ssl_contexts"].(int)
		ret.StaticSslContexts = in["static_ssl_contexts"].(int)
		ret.DynamicSslContexts = in["dynamic_ssl_contexts"].(int)
		ret.BypassFailsafeSessions = in["bypass_failsafe_sessions"].(int)
		ret.BypassUsenameSessions = in["bypass_usename_sessions"].(int)
		ret.BypassAdGroupSessions = in["bypass_ad_group_sessions"].(int)
		ret.BypassSniSessions = in["bypass_sni_sessions"].(int)
		ret.BypassCertSubjectSessions = in["bypass_cert_subject_sessions"].(int)
		ret.BypassCertIssuerSessions = in["bypass_cert_issuer_sessions"].(int)
		ret.BypassCertSanSessions = in["bypass_cert_san_sessions"].(int)
		ret.BypassNoSniSessions = in["bypass_no_sni_sessions"].(int)
		ret.ResetNoSniSessions = in["reset_no_sni_sessions"].(int)
		ret.BypassEsniSessions = in["bypass_esni_sessions"].(int)
		ret.DropEsniSessions = in["drop_esni_sessions"].(int)
		ret.BypassClientAuthSessions = in["bypass_client_auth_sessions"].(int)
		ret.SslHandshakeFailure = in["ssl_handshake_failure"].(int)
		ret.CryptioOperations = in["cryptio_operations"].(int)
		ret.TcpFailures = in["tcp_failures"].(int)
		ret.CertificateVerificationFailures = in["certificate_verification_failures"].(int)
		ret.CertificateSigningFailures = in["certificate_signing_failures"].(int)
		ret.InvalidOcspStaplingResponse = in["invalid_ocsp_stapling_response"].(int)
		ret.RevokedOcspResponse = in["revoked_ocsp_response"].(int)
		ret.UnsupportedSslVersion = in["unsupported_ssl_version"].(int)
		ret.ClientSslFatalAlert = in["client_ssl_fatal_alert"].(int)
		ret.ClientSslFinReset = in["client_ssl_fin_reset"].(int)
		ret.SslSessionFinReset = in["ssl_session_fin_reset"].(int)
		ret.ServerSslFatalAlert = in["server_ssl_fatal_alert"].(int)
		ret.ServerSslFinReset = in["server_ssl_fin_reset"].(int)
		ret.ClientSslInternalError = in["client_ssl_internal_error"].(int)
		ret.ClientSslUnknownError = in["client_ssl_unknown_error"].(int)
		ret.ServerSslInternalError = in["server_ssl_internal_error"].(int)
		ret.ServerSslUnknownError = in["server_ssl_unknown_error"].(int)
		ret.FpFailHandshake = in["fp_fail_handshake"].(int)
		ret.FpFailTcpError = in["fp_fail_tcp_error"].(int)
		ret.FpFailVerifyCert = in["fp_fail_verify_cert"].(int)
		ret.FpFailAflex = in["fp_fail_aflex"].(int)
		ret.UtilEnableStatus = in["util_enable_status"].(string)
		ret.SslModuleStats = getSliceSlbSslStatsOperOperSslModuleStats(in["ssl_module_stats"].([]interface{}))
	}
	return ret
}

func getSliceSlbSslStatsOperOperSslModuleStats(d []interface{}) []edpt.SlbSslStatsOperOperSslModuleStats {
	count := len(d)
	ret := make([]edpt.SlbSslStatsOperOperSslModuleStats, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslStatsOperOperSslModuleStats
		oi.SslModulesIndex = in["ssl_modules_index"].(int)
		oi.TotalEnabledCryptoEngines = in["total_enabled_crypto_engines"].(int)
		oi.TotalAvailableCryptoEngines = in["total_available_crypto_engines"].(int)
		oi.RequestsHandled = in["requests_handled"].(int)
		oi.UtilPercentage = in["util_percentage"].(int)
		oi.Sec1Rate = in["1sec_rate"].(int)
		oi.Sec5Rate = in["5sec_rate"].(int)
		oi.Sec10Rate = in["10sec_rate"].(int)
		oi.Sec30Rate = in["30sec_rate"].(int)
		oi.Sec60Rate = in["60sec_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslStatsOper(d *schema.ResourceData) edpt.SlbSslStatsOper {
	var ret edpt.SlbSslStatsOper
	ret.Oper = getObjectSlbSslStatsOperOper(d.Get("oper").([]interface{}))
	return ret
}
