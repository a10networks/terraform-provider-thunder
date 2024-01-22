package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCertRevokeStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_cert_revoke_stats`: Statistics for the object ssl-cert-revoke\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslCertRevokeStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ocsp_stapling_response_good": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP stapling response good",
						},
						"ocsp_chain_status_good": {
							Type: schema.TypeInt, Optional: true, Description: "Certificate chain status good",
						},
						"ocsp_chain_status_revoked": {
							Type: schema.TypeInt, Optional: true, Description: "Certificate chain status revoked",
						},
						"ocsp_chain_status_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Certificate chain status unknown",
						},
						"ocsp_request": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP requests",
						},
						"ocsp_response": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP responses",
						},
						"ocsp_connection_error": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP connection error",
						},
						"ocsp_uri_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP URI not found",
						},
						"ocsp_uri_https": {
							Type: schema.TypeInt, Optional: true, Description: "Log OCSP URI https",
						},
						"ocsp_uri_unsupported": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP URI unsupported",
						},
						"ocsp_response_status_good": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP response status good",
						},
						"ocsp_response_status_revoked": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP response status revoked",
						},
						"ocsp_response_status_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP response status unknown",
						},
						"ocsp_cache_status_good": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP cache status good",
						},
						"ocsp_cache_status_revoked": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP cache status revoked",
						},
						"ocsp_cache_miss": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP cache miss",
						},
						"ocsp_cache_expired": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP cache expired",
						},
						"ocsp_other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Log OCSP other errors",
						},
						"ocsp_response_no_nonce": {
							Type: schema.TypeInt, Optional: true, Description: "Log OCSP other errors",
						},
						"ocsp_response_nonce_error": {
							Type: schema.TypeInt, Optional: true, Description: "Log OCSP other errors",
						},
						"crl_request": {
							Type: schema.TypeInt, Optional: true, Description: "CRL requests",
						},
						"crl_response": {
							Type: schema.TypeInt, Optional: true, Description: "CRL responses",
						},
						"crl_connection_error": {
							Type: schema.TypeInt, Optional: true, Description: "CRL connection errors",
						},
						"crl_uri_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "CRL URI not found",
						},
						"crl_uri_https": {
							Type: schema.TypeInt, Optional: true, Description: "CRL URI https",
						},
						"crl_uri_unsupported": {
							Type: schema.TypeInt, Optional: true, Description: "CRL URI unsupported",
						},
						"crl_response_status_good": {
							Type: schema.TypeInt, Optional: true, Description: "CRL response status good",
						},
						"crl_response_status_revoked": {
							Type: schema.TypeInt, Optional: true, Description: "CRL response status revoked",
						},
						"crl_response_status_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "CRL response status unknown",
						},
						"crl_cache_status_good": {
							Type: schema.TypeInt, Optional: true, Description: "CRL cache status good",
						},
						"crl_cache_status_revoked": {
							Type: schema.TypeInt, Optional: true, Description: "CRL cache status revoked",
						},
						"crl_other_error": {
							Type: schema.TypeInt, Optional: true, Description: "CRL other errors",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSslCertRevokeStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCertRevokeStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCertRevokeStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslCertRevokeStatsStats := setObjectSlbSslCertRevokeStatsStats(res)
		d.Set("stats", SlbSslCertRevokeStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslCertRevokeStatsStats(ret edpt.DataSlbSslCertRevokeStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ocsp_stapling_response_good":  ret.DtSlbSslCertRevokeStats.Stats.Ocsp_stapling_response_good,
			"ocsp_chain_status_good":       ret.DtSlbSslCertRevokeStats.Stats.Ocsp_chain_status_good,
			"ocsp_chain_status_revoked":    ret.DtSlbSslCertRevokeStats.Stats.Ocsp_chain_status_revoked,
			"ocsp_chain_status_unknown":    ret.DtSlbSslCertRevokeStats.Stats.Ocsp_chain_status_unknown,
			"ocsp_request":                 ret.DtSlbSslCertRevokeStats.Stats.Ocsp_request,
			"ocsp_response":                ret.DtSlbSslCertRevokeStats.Stats.Ocsp_response,
			"ocsp_connection_error":        ret.DtSlbSslCertRevokeStats.Stats.Ocsp_connection_error,
			"ocsp_uri_not_found":           ret.DtSlbSslCertRevokeStats.Stats.Ocsp_uri_not_found,
			"ocsp_uri_https":               ret.DtSlbSslCertRevokeStats.Stats.Ocsp_uri_https,
			"ocsp_uri_unsupported":         ret.DtSlbSslCertRevokeStats.Stats.Ocsp_uri_unsupported,
			"ocsp_response_status_good":    ret.DtSlbSslCertRevokeStats.Stats.Ocsp_response_status_good,
			"ocsp_response_status_revoked": ret.DtSlbSslCertRevokeStats.Stats.Ocsp_response_status_revoked,
			"ocsp_response_status_unknown": ret.DtSlbSslCertRevokeStats.Stats.Ocsp_response_status_unknown,
			"ocsp_cache_status_good":       ret.DtSlbSslCertRevokeStats.Stats.Ocsp_cache_status_good,
			"ocsp_cache_status_revoked":    ret.DtSlbSslCertRevokeStats.Stats.Ocsp_cache_status_revoked,
			"ocsp_cache_miss":              ret.DtSlbSslCertRevokeStats.Stats.Ocsp_cache_miss,
			"ocsp_cache_expired":           ret.DtSlbSslCertRevokeStats.Stats.Ocsp_cache_expired,
			"ocsp_other_error":             ret.DtSlbSslCertRevokeStats.Stats.Ocsp_other_error,
			"ocsp_response_no_nonce":       ret.DtSlbSslCertRevokeStats.Stats.Ocsp_response_no_nonce,
			"ocsp_response_nonce_error":    ret.DtSlbSslCertRevokeStats.Stats.Ocsp_response_nonce_error,
			"crl_request":                  ret.DtSlbSslCertRevokeStats.Stats.Crl_request,
			"crl_response":                 ret.DtSlbSslCertRevokeStats.Stats.Crl_response,
			"crl_connection_error":         ret.DtSlbSslCertRevokeStats.Stats.Crl_connection_error,
			"crl_uri_not_found":            ret.DtSlbSslCertRevokeStats.Stats.Crl_uri_not_found,
			"crl_uri_https":                ret.DtSlbSslCertRevokeStats.Stats.Crl_uri_https,
			"crl_uri_unsupported":          ret.DtSlbSslCertRevokeStats.Stats.Crl_uri_unsupported,
			"crl_response_status_good":     ret.DtSlbSslCertRevokeStats.Stats.Crl_response_status_good,
			"crl_response_status_revoked":  ret.DtSlbSslCertRevokeStats.Stats.Crl_response_status_revoked,
			"crl_response_status_unknown":  ret.DtSlbSslCertRevokeStats.Stats.Crl_response_status_unknown,
			"crl_cache_status_good":        ret.DtSlbSslCertRevokeStats.Stats.Crl_cache_status_good,
			"crl_cache_status_revoked":     ret.DtSlbSslCertRevokeStats.Stats.Crl_cache_status_revoked,
			"crl_other_error":              ret.DtSlbSslCertRevokeStats.Stats.Crl_other_error,
		},
	}
}

func getObjectSlbSslCertRevokeStatsStats(d []interface{}) edpt.SlbSslCertRevokeStatsStats {

	count1 := len(d)
	var ret edpt.SlbSslCertRevokeStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ocsp_stapling_response_good = in["ocsp_stapling_response_good"].(int)
		ret.Ocsp_chain_status_good = in["ocsp_chain_status_good"].(int)
		ret.Ocsp_chain_status_revoked = in["ocsp_chain_status_revoked"].(int)
		ret.Ocsp_chain_status_unknown = in["ocsp_chain_status_unknown"].(int)
		ret.Ocsp_request = in["ocsp_request"].(int)
		ret.Ocsp_response = in["ocsp_response"].(int)
		ret.Ocsp_connection_error = in["ocsp_connection_error"].(int)
		ret.Ocsp_uri_not_found = in["ocsp_uri_not_found"].(int)
		ret.Ocsp_uri_https = in["ocsp_uri_https"].(int)
		ret.Ocsp_uri_unsupported = in["ocsp_uri_unsupported"].(int)
		ret.Ocsp_response_status_good = in["ocsp_response_status_good"].(int)
		ret.Ocsp_response_status_revoked = in["ocsp_response_status_revoked"].(int)
		ret.Ocsp_response_status_unknown = in["ocsp_response_status_unknown"].(int)
		ret.Ocsp_cache_status_good = in["ocsp_cache_status_good"].(int)
		ret.Ocsp_cache_status_revoked = in["ocsp_cache_status_revoked"].(int)
		ret.Ocsp_cache_miss = in["ocsp_cache_miss"].(int)
		ret.Ocsp_cache_expired = in["ocsp_cache_expired"].(int)
		ret.Ocsp_other_error = in["ocsp_other_error"].(int)
		ret.Ocsp_response_no_nonce = in["ocsp_response_no_nonce"].(int)
		ret.Ocsp_response_nonce_error = in["ocsp_response_nonce_error"].(int)
		ret.Crl_request = in["crl_request"].(int)
		ret.Crl_response = in["crl_response"].(int)
		ret.Crl_connection_error = in["crl_connection_error"].(int)
		ret.Crl_uri_not_found = in["crl_uri_not_found"].(int)
		ret.Crl_uri_https = in["crl_uri_https"].(int)
		ret.Crl_uri_unsupported = in["crl_uri_unsupported"].(int)
		ret.Crl_response_status_good = in["crl_response_status_good"].(int)
		ret.Crl_response_status_revoked = in["crl_response_status_revoked"].(int)
		ret.Crl_response_status_unknown = in["crl_response_status_unknown"].(int)
		ret.Crl_cache_status_good = in["crl_cache_status_good"].(int)
		ret.Crl_cache_status_revoked = in["crl_cache_status_revoked"].(int)
		ret.Crl_other_error = in["crl_other_error"].(int)
	}
	return ret
}

func dataToEndpointSlbSslCertRevokeStats(d *schema.ResourceData) edpt.SlbSslCertRevokeStats {
	var ret edpt.SlbSslCertRevokeStats

	ret.Stats = getObjectSlbSslCertRevokeStatsStats(d.Get("stats").([]interface{}))
	return ret
}
