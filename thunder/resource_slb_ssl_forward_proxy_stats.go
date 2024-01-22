package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslForwardProxyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_forward_proxy_stats`: Statistics for the object ssl-forward-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslForwardProxyStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_create": {
							Type: schema.TypeInt, Optional: true, Description: "Certificates created",
						},
						"cert_expr": {
							Type: schema.TypeInt, Optional: true, Description: "Certificates expired",
						},
						"cert_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Certificate cache hits",
						},
						"cert_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Certificate cache miss",
						},
						"conn_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "Connections bypassed",
						},
						"conn_inspect": {
							Type: schema.TypeInt, Optional: true, Description: "Connections inspected",
						},
						"bypass_failsafe_ssl_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Failsafe SSL sessions",
						},
						"bypass_sni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass SNI sessions",
						},
						"bypass_client_auth_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Client Auth sessions",
						},
						"failed_in_ssl_handshakes": {
							Type: schema.TypeInt, Optional: true, Description: "Failed in SSL handshakes",
						},
						"failed_in_crypto_operations": {
							Type: schema.TypeInt, Optional: true, Description: "Failed in crypto operations",
						},
						"failed_in_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Failed in TCP",
						},
						"failed_in_certificate_verification": {
							Type: schema.TypeInt, Optional: true, Description: "Failed in Certificate verification",
						},
						"failed_in_certificate_signing": {
							Type: schema.TypeInt, Optional: true, Description: "Failed in Certificate signing",
						},
						"invalid_ocsp_stapling_response": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid OCSP Stapling Response",
						},
						"revoked_ocsp_response": {
							Type: schema.TypeInt, Optional: true, Description: "Revoked OCSP Response",
						},
						"unsupported_ssl_version": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported SSL version",
						},
						"certificates_in_cache": {
							Type: schema.TypeInt, Optional: true, Description: "Certificates in cache",
						},
						"connections_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Connections failed",
						},
						"aflex_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass triggered by aFleX",
						},
						"bypass_cert_subject_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Cert Subject sessions",
						},
						"bypass_cert_issuer_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Cert issuer sessions",
						},
						"bypass_cert_san_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Cert SAN sessions",
						},
						"bypass_no_sni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass NO SNI sessions",
						},
						"reset_no_sni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Reset No SNI sessions",
						},
						"bypass_esni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass ESNI sessions",
						},
						"drop_esni_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Drop ESNI sessions",
						},
						"bypass_username_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Username sessions",
						},
						"bypass_ad_group_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass AD-group sessions",
						},
						"tot_conn_in_buff": {
							Type: schema.TypeInt, Optional: true, Description: "Total buffered async connections",
						},
						"curr_conn_in_buff": {
							Type: schema.TypeInt, Optional: true, Description: "Current buffered async connections",
						},
						"async_conn_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "SSLi Async connections failures due to timeout",
						},
						"async_conn_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SSLi Async connections failures due to limit",
						},
						"cert_in_cache": {
							Type: schema.TypeInt, Optional: true, Description: "Certificates in cache used by HC SSLi App",
						},
						"bypass_client_ip_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Client IP sessions",
						},
						"bypass_server_ip_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Bypass Server IP sessions",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslForwardProxyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslForwardProxyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslForwardProxyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslForwardProxyStatsStats := setObjectSlbSslForwardProxyStatsStats(res)
		d.Set("stats", SlbSslForwardProxyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslForwardProxyStatsStats(ret edpt.DataSlbSslForwardProxyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cert_create":                        ret.DtSlbSslForwardProxyStats.Stats.Cert_create,
			"cert_expr":                          ret.DtSlbSslForwardProxyStats.Stats.Cert_expr,
			"cert_hit":                           ret.DtSlbSslForwardProxyStats.Stats.Cert_hit,
			"cert_miss":                          ret.DtSlbSslForwardProxyStats.Stats.Cert_miss,
			"conn_bypass":                        ret.DtSlbSslForwardProxyStats.Stats.Conn_bypass,
			"conn_inspect":                       ret.DtSlbSslForwardProxyStats.Stats.Conn_inspect,
			"bypass_failsafe_ssl_sessions":       ret.DtSlbSslForwardProxyStats.Stats.BypassFailsafeSslSessions,
			"bypass_sni_sessions":                ret.DtSlbSslForwardProxyStats.Stats.BypassSniSessions,
			"bypass_client_auth_sessions":        ret.DtSlbSslForwardProxyStats.Stats.BypassClientAuthSessions,
			"failed_in_ssl_handshakes":           ret.DtSlbSslForwardProxyStats.Stats.FailedInSslHandshakes,
			"failed_in_crypto_operations":        ret.DtSlbSslForwardProxyStats.Stats.FailedInCryptoOperations,
			"failed_in_tcp":                      ret.DtSlbSslForwardProxyStats.Stats.FailedInTcp,
			"failed_in_certificate_verification": ret.DtSlbSslForwardProxyStats.Stats.FailedInCertificateVerification,
			"failed_in_certificate_signing":      ret.DtSlbSslForwardProxyStats.Stats.FailedInCertificateSigning,
			"invalid_ocsp_stapling_response":     ret.DtSlbSslForwardProxyStats.Stats.InvalidOcspStaplingResponse,
			"revoked_ocsp_response":              ret.DtSlbSslForwardProxyStats.Stats.RevokedOcspResponse,
			"unsupported_ssl_version":            ret.DtSlbSslForwardProxyStats.Stats.UnsupportedSslVersion,
			"certificates_in_cache":              ret.DtSlbSslForwardProxyStats.Stats.CertificatesInCache,
			"connections_failed":                 ret.DtSlbSslForwardProxyStats.Stats.ConnectionsFailed,
			"aflex_bypass":                       ret.DtSlbSslForwardProxyStats.Stats.AflexBypass,
			"bypass_cert_subject_sessions":       ret.DtSlbSslForwardProxyStats.Stats.BypassCertSubjectSessions,
			"bypass_cert_issuer_sessions":        ret.DtSlbSslForwardProxyStats.Stats.BypassCertIssuerSessions,
			"bypass_cert_san_sessions":           ret.DtSlbSslForwardProxyStats.Stats.BypassCertSanSessions,
			"bypass_no_sni_sessions":             ret.DtSlbSslForwardProxyStats.Stats.BypassNoSniSessions,
			"reset_no_sni_sessions":              ret.DtSlbSslForwardProxyStats.Stats.ResetNoSniSessions,
			"bypass_esni_sessions":               ret.DtSlbSslForwardProxyStats.Stats.BypassEsniSessions,
			"drop_esni_sessions":                 ret.DtSlbSslForwardProxyStats.Stats.DropEsniSessions,
			"bypass_username_sessions":           ret.DtSlbSslForwardProxyStats.Stats.BypassUsernameSessions,
			"bypass_ad_group_sessions":           ret.DtSlbSslForwardProxyStats.Stats.BypassAdGroupSessions,
			"tot_conn_in_buff":                   ret.DtSlbSslForwardProxyStats.Stats.Tot_conn_in_buff,
			"curr_conn_in_buff":                  ret.DtSlbSslForwardProxyStats.Stats.Curr_conn_in_buff,
			"async_conn_timeout":                 ret.DtSlbSslForwardProxyStats.Stats.Async_conn_timeout,
			"async_conn_limit_drop":              ret.DtSlbSslForwardProxyStats.Stats.Async_conn_limit_drop,
			"cert_in_cache":                      ret.DtSlbSslForwardProxyStats.Stats.Cert_in_cache,
			"bypass_client_ip_sessions":          ret.DtSlbSslForwardProxyStats.Stats.BypassClientIpSessions,
			"bypass_server_ip_sessions":          ret.DtSlbSslForwardProxyStats.Stats.BypassServerIpSessions,
		},
	}
}

func getObjectSlbSslForwardProxyStatsStats(d []interface{}) edpt.SlbSslForwardProxyStatsStats {

	count1 := len(d)
	var ret edpt.SlbSslForwardProxyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cert_create = in["cert_create"].(int)
		ret.Cert_expr = in["cert_expr"].(int)
		ret.Cert_hit = in["cert_hit"].(int)
		ret.Cert_miss = in["cert_miss"].(int)
		ret.Conn_bypass = in["conn_bypass"].(int)
		ret.Conn_inspect = in["conn_inspect"].(int)
		ret.BypassFailsafeSslSessions = in["bypass_failsafe_ssl_sessions"].(int)
		ret.BypassSniSessions = in["bypass_sni_sessions"].(int)
		ret.BypassClientAuthSessions = in["bypass_client_auth_sessions"].(int)
		ret.FailedInSslHandshakes = in["failed_in_ssl_handshakes"].(int)
		ret.FailedInCryptoOperations = in["failed_in_crypto_operations"].(int)
		ret.FailedInTcp = in["failed_in_tcp"].(int)
		ret.FailedInCertificateVerification = in["failed_in_certificate_verification"].(int)
		ret.FailedInCertificateSigning = in["failed_in_certificate_signing"].(int)
		ret.InvalidOcspStaplingResponse = in["invalid_ocsp_stapling_response"].(int)
		ret.RevokedOcspResponse = in["revoked_ocsp_response"].(int)
		ret.UnsupportedSslVersion = in["unsupported_ssl_version"].(int)
		ret.CertificatesInCache = in["certificates_in_cache"].(int)
		ret.ConnectionsFailed = in["connections_failed"].(int)
		ret.AflexBypass = in["aflex_bypass"].(int)
		ret.BypassCertSubjectSessions = in["bypass_cert_subject_sessions"].(int)
		ret.BypassCertIssuerSessions = in["bypass_cert_issuer_sessions"].(int)
		ret.BypassCertSanSessions = in["bypass_cert_san_sessions"].(int)
		ret.BypassNoSniSessions = in["bypass_no_sni_sessions"].(int)
		ret.ResetNoSniSessions = in["reset_no_sni_sessions"].(int)
		ret.BypassEsniSessions = in["bypass_esni_sessions"].(int)
		ret.DropEsniSessions = in["drop_esni_sessions"].(int)
		ret.BypassUsernameSessions = in["bypass_username_sessions"].(int)
		ret.BypassAdGroupSessions = in["bypass_ad_group_sessions"].(int)
		ret.Tot_conn_in_buff = in["tot_conn_in_buff"].(int)
		ret.Curr_conn_in_buff = in["curr_conn_in_buff"].(int)
		ret.Async_conn_timeout = in["async_conn_timeout"].(int)
		ret.Async_conn_limit_drop = in["async_conn_limit_drop"].(int)
		ret.Cert_in_cache = in["cert_in_cache"].(int)
		ret.BypassClientIpSessions = in["bypass_client_ip_sessions"].(int)
		ret.BypassServerIpSessions = in["bypass_server_ip_sessions"].(int)
	}
	return ret
}

func dataToEndpointSlbSslForwardProxyStats(d *schema.ResourceData) edpt.SlbSslForwardProxyStats {
	var ret edpt.SlbSslForwardProxyStats

	ret.Stats = getObjectSlbSslForwardProxyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
