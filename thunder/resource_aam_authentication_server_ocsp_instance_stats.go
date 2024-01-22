package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerOcspInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_ocsp_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerOcspInstanceStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify OCSP authentication server name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Request",
						},
						"certificate_good": {
							Type: schema.TypeInt, Optional: true, Description: "Good Certificate Response",
						},
						"certificate_revoked": {
							Type: schema.TypeInt, Optional: true, Description: "Revoked Certificate Response",
						},
						"certificate_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Certificate Response",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Timeout",
						},
						"fail": {
							Type: schema.TypeInt, Optional: true, Description: "Handle OCSP response failed",
						},
						"stapling_request": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP Stapling Request Send",
						},
						"stapling_certificate_good": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP Stapling Good Certificate Response",
						},
						"stapling_certificate_revoked": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP Stapling Revoked Certificate Response",
						},
						"stapling_certificate_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP Stapling Unknown Certificate Response",
						},
						"stapling_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "OCSP Stapling Timeout",
						},
						"stapling_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Handle OCSP response failed",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationServerOcspInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcspInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerOcspInstanceStatsStats := setObjectAamAuthenticationServerOcspInstanceStatsStats(res)
		d.Set("stats", AamAuthenticationServerOcspInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationServerOcspInstanceStatsStats(ret edpt.DataAamAuthenticationServerOcspInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request":                      ret.DtAamAuthenticationServerOcspInstanceStats.Stats.Request,
			"certificate_good":             ret.DtAamAuthenticationServerOcspInstanceStats.Stats.CertificateGood,
			"certificate_revoked":          ret.DtAamAuthenticationServerOcspInstanceStats.Stats.CertificateRevoked,
			"certificate_unknown":          ret.DtAamAuthenticationServerOcspInstanceStats.Stats.CertificateUnknown,
			"timeout":                      ret.DtAamAuthenticationServerOcspInstanceStats.Stats.Timeout,
			"fail":                         ret.DtAamAuthenticationServerOcspInstanceStats.Stats.Fail,
			"stapling_request":             ret.DtAamAuthenticationServerOcspInstanceStats.Stats.StaplingRequest,
			"stapling_certificate_good":    ret.DtAamAuthenticationServerOcspInstanceStats.Stats.StaplingCertificateGood,
			"stapling_certificate_revoked": ret.DtAamAuthenticationServerOcspInstanceStats.Stats.StaplingCertificateRevoked,
			"stapling_certificate_unknown": ret.DtAamAuthenticationServerOcspInstanceStats.Stats.StaplingCertificateUnknown,
			"stapling_timeout":             ret.DtAamAuthenticationServerOcspInstanceStats.Stats.StaplingTimeout,
			"stapling_fail":                ret.DtAamAuthenticationServerOcspInstanceStats.Stats.StaplingFail,
		},
	}
}

func getObjectAamAuthenticationServerOcspInstanceStatsStats(d []interface{}) edpt.AamAuthenticationServerOcspInstanceStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOcspInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Request = in["request"].(int)
		ret.CertificateGood = in["certificate_good"].(int)
		ret.CertificateRevoked = in["certificate_revoked"].(int)
		ret.CertificateUnknown = in["certificate_unknown"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.Fail = in["fail"].(int)
		ret.StaplingRequest = in["stapling_request"].(int)
		ret.StaplingCertificateGood = in["stapling_certificate_good"].(int)
		ret.StaplingCertificateRevoked = in["stapling_certificate_revoked"].(int)
		ret.StaplingCertificateUnknown = in["stapling_certificate_unknown"].(int)
		ret.StaplingTimeout = in["stapling_timeout"].(int)
		ret.StaplingFail = in["stapling_fail"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerOcspInstanceStats(d *schema.ResourceData) edpt.AamAuthenticationServerOcspInstanceStats {
	var ret edpt.AamAuthenticationServerOcspInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationServerOcspInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
