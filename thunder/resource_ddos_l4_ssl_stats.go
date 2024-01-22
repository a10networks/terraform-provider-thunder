package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL4SslStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l4_ssl_stats`: Statistics for the object l4-ssl\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL4SslStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_l4_policy_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Reset",
						},
						"ssl_l4_policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Dropped",
						},
						"ssl_l4_drop_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Dropped",
						},
						"ssl_l4_er_condition": {
							Type: schema.TypeInt, Optional: true, Description: "Error Condition",
						},
						"ssl_l4_processed": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Processed",
						},
						"ssl_l4_new_syn": {
							Type: schema.TypeInt, Optional: true, Description: "New TCP SYN",
						},
						"ssl_l4_is_ssl3": {
							Type: schema.TypeInt, Optional: true, Description: "SSL v3",
						},
						"ssl_l4_is_tls1_1": {
							Type: schema.TypeInt, Optional: true, Description: "TLS v1.0",
						},
						"ssl_l4_is_tls1_2": {
							Type: schema.TypeInt, Optional: true, Description: "TLS v1.1",
						},
						"ssl_l4_is_tls1_3": {
							Type: schema.TypeInt, Optional: true, Description: "TLS v1.2",
						},
						"ssl_l4_is_renegotiation": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Renegotiation",
						},
						"ssl_l4_renegotiation_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Renegotiation Exceeded",
						},
						"ssl_l4_is_dst_req_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate Exceed Dropped",
						},
						"ssl_l4_is_src_req_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate Exceed Dropped",
						},
						"ssl_l4_do_auth_handshake": {
							Type: schema.TypeInt, Optional: true, Description: "Do Auth Handshake",
						},
						"ssl_l4_reset_for_handshake": {
							Type: schema.TypeInt, Optional: true, Description: "Reset While Others in Handshake",
						},
						"ssl_l4_handshake_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Handshake Timeout",
						},
						"ssl_l4_auth_handshake_ok": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Handshake Success",
						},
						"ssl_l4_auth_handshake_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Handshake Blacklisted",
						},
						"ssl_l4_auth_handshake_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Handshake Failed",
						},
						"ssl_l4_renegotiation_incomplete": {
							Type: schema.TypeInt, Optional: true, Description: "Renegotiation Incomplete",
						},
						"ssl_l4_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Auth Dropped",
						},
						"ssl_l4_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Auth Responded",
						},
						"ssl_non_tls": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Non-TLS Dropped",
						},
						"ssl_header_invalid_type": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Header Invalid Type",
						},
						"ssl_header_bad_ver": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Header Bad Version",
						},
						"ssl_header_bad_len": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Header Bad Length",
						},
						"ssl_bad_header_forw": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Traffic Check Bad Header Forwarded",
						},
						"ssl_bad_header_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SSL Traffic Check Bad Header Dropped",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL4SslStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL4SslStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL4SslStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL4SslStatsStats := setObjectDdosL4SslStatsStats(res)
		d.Set("stats", DdosL4SslStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL4SslStatsStats(ret edpt.DataDdosL4SslStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ssl_l4_policy_reset":             ret.DtDdosL4SslStats.Stats.Ssl_l4_policy_reset,
			"ssl_l4_policy_drop":              ret.DtDdosL4SslStats.Stats.Ssl_l4_policy_drop,
			"ssl_l4_drop_packet":              ret.DtDdosL4SslStats.Stats.Ssl_l4_drop_packet,
			"ssl_l4_er_condition":             ret.DtDdosL4SslStats.Stats.Ssl_l4_er_condition,
			"ssl_l4_processed":                ret.DtDdosL4SslStats.Stats.Ssl_l4_processed,
			"ssl_l4_new_syn":                  ret.DtDdosL4SslStats.Stats.Ssl_l4_new_syn,
			"ssl_l4_is_ssl3":                  ret.DtDdosL4SslStats.Stats.Ssl_l4_is_ssl3,
			"ssl_l4_is_tls1_1":                ret.DtDdosL4SslStats.Stats.Ssl_l4_is_tls1_1,
			"ssl_l4_is_tls1_2":                ret.DtDdosL4SslStats.Stats.Ssl_l4_is_tls1_2,
			"ssl_l4_is_tls1_3":                ret.DtDdosL4SslStats.Stats.Ssl_l4_is_tls1_3,
			"ssl_l4_is_renegotiation":         ret.DtDdosL4SslStats.Stats.Ssl_l4_is_renegotiation,
			"ssl_l4_renegotiation_exceed":     ret.DtDdosL4SslStats.Stats.Ssl_l4_renegotiation_exceed,
			"ssl_l4_is_dst_req_rate_exceed":   ret.DtDdosL4SslStats.Stats.Ssl_l4_is_dst_req_rate_exceed,
			"ssl_l4_is_src_req_rate_exceed":   ret.DtDdosL4SslStats.Stats.Ssl_l4_is_src_req_rate_exceed,
			"ssl_l4_do_auth_handshake":        ret.DtDdosL4SslStats.Stats.Ssl_l4_do_auth_handshake,
			"ssl_l4_reset_for_handshake":      ret.DtDdosL4SslStats.Stats.Ssl_l4_reset_for_handshake,
			"ssl_l4_handshake_timeout":        ret.DtDdosL4SslStats.Stats.Ssl_l4_handshake_timeout,
			"ssl_l4_auth_handshake_ok":        ret.DtDdosL4SslStats.Stats.Ssl_l4_auth_handshake_ok,
			"ssl_l4_auth_handshake_bl":        ret.DtDdosL4SslStats.Stats.Ssl_l4_auth_handshake_bl,
			"ssl_l4_auth_handshake_fail":      ret.DtDdosL4SslStats.Stats.Ssl_l4_auth_handshake_fail,
			"ssl_l4_renegotiation_incomplete": ret.DtDdosL4SslStats.Stats.Ssl_l4_renegotiation_incomplete,
			"ssl_l4_auth_drop":                ret.DtDdosL4SslStats.Stats.Ssl_l4_auth_drop,
			"ssl_l4_auth_resp":                ret.DtDdosL4SslStats.Stats.Ssl_l4_auth_resp,
			"ssl_non_tls":                     ret.DtDdosL4SslStats.Stats.Ssl_non_tls,
			"ssl_header_invalid_type":         ret.DtDdosL4SslStats.Stats.Ssl_header_invalid_type,
			"ssl_header_bad_ver":              ret.DtDdosL4SslStats.Stats.Ssl_header_bad_ver,
			"ssl_header_bad_len":              ret.DtDdosL4SslStats.Stats.Ssl_header_bad_len,
			"ssl_bad_header_forw":             ret.DtDdosL4SslStats.Stats.Ssl_bad_header_forw,
			"ssl_bad_header_drop":             ret.DtDdosL4SslStats.Stats.Ssl_bad_header_drop,
		},
	}
}

func getObjectDdosL4SslStatsStats(d []interface{}) edpt.DdosL4SslStatsStats {

	count1 := len(d)
	var ret edpt.DdosL4SslStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ssl_l4_policy_reset = in["ssl_l4_policy_reset"].(int)
		ret.Ssl_l4_policy_drop = in["ssl_l4_policy_drop"].(int)
		ret.Ssl_l4_drop_packet = in["ssl_l4_drop_packet"].(int)
		ret.Ssl_l4_er_condition = in["ssl_l4_er_condition"].(int)
		ret.Ssl_l4_processed = in["ssl_l4_processed"].(int)
		ret.Ssl_l4_new_syn = in["ssl_l4_new_syn"].(int)
		ret.Ssl_l4_is_ssl3 = in["ssl_l4_is_ssl3"].(int)
		ret.Ssl_l4_is_tls1_1 = in["ssl_l4_is_tls1_1"].(int)
		ret.Ssl_l4_is_tls1_2 = in["ssl_l4_is_tls1_2"].(int)
		ret.Ssl_l4_is_tls1_3 = in["ssl_l4_is_tls1_3"].(int)
		ret.Ssl_l4_is_renegotiation = in["ssl_l4_is_renegotiation"].(int)
		ret.Ssl_l4_renegotiation_exceed = in["ssl_l4_renegotiation_exceed"].(int)
		ret.Ssl_l4_is_dst_req_rate_exceed = in["ssl_l4_is_dst_req_rate_exceed"].(int)
		ret.Ssl_l4_is_src_req_rate_exceed = in["ssl_l4_is_src_req_rate_exceed"].(int)
		ret.Ssl_l4_do_auth_handshake = in["ssl_l4_do_auth_handshake"].(int)
		ret.Ssl_l4_reset_for_handshake = in["ssl_l4_reset_for_handshake"].(int)
		ret.Ssl_l4_handshake_timeout = in["ssl_l4_handshake_timeout"].(int)
		ret.Ssl_l4_auth_handshake_ok = in["ssl_l4_auth_handshake_ok"].(int)
		ret.Ssl_l4_auth_handshake_bl = in["ssl_l4_auth_handshake_bl"].(int)
		ret.Ssl_l4_auth_handshake_fail = in["ssl_l4_auth_handshake_fail"].(int)
		ret.Ssl_l4_renegotiation_incomplete = in["ssl_l4_renegotiation_incomplete"].(int)
		ret.Ssl_l4_auth_drop = in["ssl_l4_auth_drop"].(int)
		ret.Ssl_l4_auth_resp = in["ssl_l4_auth_resp"].(int)
		ret.Ssl_non_tls = in["ssl_non_tls"].(int)
		ret.Ssl_header_invalid_type = in["ssl_header_invalid_type"].(int)
		ret.Ssl_header_bad_ver = in["ssl_header_bad_ver"].(int)
		ret.Ssl_header_bad_len = in["ssl_header_bad_len"].(int)
		ret.Ssl_bad_header_forw = in["ssl_bad_header_forw"].(int)
		ret.Ssl_bad_header_drop = in["ssl_bad_header_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosL4SslStats(d *schema.ResourceData) edpt.DdosL4SslStats {
	var ret edpt.DdosL4SslStats

	ret.Stats = getObjectDdosL4SslStatsStats(d.Get("stats").([]interface{}))
	return ret
}
