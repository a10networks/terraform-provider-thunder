package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerOcspStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_ocsp_stats`: Statistics for the object ocsp\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerOcspStatsRead,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stapling_certificate_good": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Good Certificate Response",
						},
						"stapling_certificate_revoked": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Revoked Certificate Response",
						},
						"stapling_certificate_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Unknown Certificate Response",
						},
						"stapling_request_normal": {
							Type: schema.TypeInt, Optional: true, Description: "Total OSCP Stapling Normal Request",
						},
						"stapling_request_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Dropped Request",
						},
						"stapling_response_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Success Response",
						},
						"stapling_response_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Failure Response",
						},
						"stapling_response_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Error Response",
						},
						"stapling_response_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Timeout Response",
						},
						"stapling_response_other": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Other Response",
						},
						"request_normal": {
							Type: schema.TypeInt, Optional: true, Description: "Total OSCP Normal Request",
						},
						"request_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Dropped Request",
						},
						"response_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Success Response",
						},
						"response_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Failure Response",
						},
						"response_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Error Response",
						},
						"response_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Timeout Response",
						},
						"response_other": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Other Response",
						},
						"job_start_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Job Start Error",
						},
						"polling_control_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Polling Control Error",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationServerOcspStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerOcspStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerOcspStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerOcspStatsInstanceList := setSliceAamAuthenticationServerOcspStatsInstanceList(res)
		d.Set("instance_list", AamAuthenticationServerOcspStatsInstanceList)
		AamAuthenticationServerOcspStatsStats := setObjectAamAuthenticationServerOcspStatsStats(res)
		d.Set("stats", AamAuthenticationServerOcspStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceAamAuthenticationServerOcspStatsInstanceList(d edpt.DataAamAuthenticationServerOcspStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtAamAuthenticationServerOcspStats.InstanceList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectAamAuthenticationServerOcspStatsInstanceListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectAamAuthenticationServerOcspStatsInstanceListStats(d edpt.AamAuthenticationServerOcspStatsInstanceListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["request"] = d.Request

	in["certificate_good"] = d.CertificateGood

	in["certificate_revoked"] = d.CertificateRevoked

	in["certificate_unknown"] = d.CertificateUnknown

	in["timeout"] = d.Timeout

	in["fail"] = d.Fail

	in["stapling_request"] = d.StaplingRequest

	in["stapling_certificate_good"] = d.StaplingCertificateGood

	in["stapling_certificate_revoked"] = d.StaplingCertificateRevoked

	in["stapling_certificate_unknown"] = d.StaplingCertificateUnknown

	in["stapling_timeout"] = d.StaplingTimeout

	in["stapling_fail"] = d.StaplingFail
	result = append(result, in)
	return result
}

func setObjectAamAuthenticationServerOcspStatsStats(ret edpt.DataAamAuthenticationServerOcspStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stapling_certificate_good":    ret.DtAamAuthenticationServerOcspStats.Stats.StaplingCertificateGood,
			"stapling_certificate_revoked": ret.DtAamAuthenticationServerOcspStats.Stats.StaplingCertificateRevoked,
			"stapling_certificate_unknown": ret.DtAamAuthenticationServerOcspStats.Stats.StaplingCertificateUnknown,
			"stapling_request_normal":      ret.DtAamAuthenticationServerOcspStats.Stats.StaplingRequestNormal,
			"stapling_request_dropped":     ret.DtAamAuthenticationServerOcspStats.Stats.StaplingRequestDropped,
			"stapling_response_success":    ret.DtAamAuthenticationServerOcspStats.Stats.StaplingResponseSuccess,
			"stapling_response_failure":    ret.DtAamAuthenticationServerOcspStats.Stats.StaplingResponseFailure,
			"stapling_response_error":      ret.DtAamAuthenticationServerOcspStats.Stats.StaplingResponseError,
			"stapling_response_timeout":    ret.DtAamAuthenticationServerOcspStats.Stats.StaplingResponseTimeout,
			"stapling_response_other":      ret.DtAamAuthenticationServerOcspStats.Stats.StaplingResponseOther,
			"request_normal":               ret.DtAamAuthenticationServerOcspStats.Stats.RequestNormal,
			"request_dropped":              ret.DtAamAuthenticationServerOcspStats.Stats.RequestDropped,
			"response_success":             ret.DtAamAuthenticationServerOcspStats.Stats.ResponseSuccess,
			"response_failure":             ret.DtAamAuthenticationServerOcspStats.Stats.ResponseFailure,
			"response_error":               ret.DtAamAuthenticationServerOcspStats.Stats.ResponseError,
			"response_timeout":             ret.DtAamAuthenticationServerOcspStats.Stats.ResponseTimeout,
			"response_other":               ret.DtAamAuthenticationServerOcspStats.Stats.ResponseOther,
			"job_start_error":              ret.DtAamAuthenticationServerOcspStats.Stats.JobStartError,
			"polling_control_error":        ret.DtAamAuthenticationServerOcspStats.Stats.PollingControlError,
		},
	}
}

func getSliceAamAuthenticationServerOcspStatsInstanceList(d []interface{}) []edpt.AamAuthenticationServerOcspStatsInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerOcspStatsInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerOcspStatsInstanceList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectAamAuthenticationServerOcspStatsInstanceListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerOcspStatsInstanceListStats(d []interface{}) edpt.AamAuthenticationServerOcspStatsInstanceListStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOcspStatsInstanceListStats
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

func getObjectAamAuthenticationServerOcspStatsStats(d []interface{}) edpt.AamAuthenticationServerOcspStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerOcspStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaplingCertificateGood = in["stapling_certificate_good"].(int)
		ret.StaplingCertificateRevoked = in["stapling_certificate_revoked"].(int)
		ret.StaplingCertificateUnknown = in["stapling_certificate_unknown"].(int)
		ret.StaplingRequestNormal = in["stapling_request_normal"].(int)
		ret.StaplingRequestDropped = in["stapling_request_dropped"].(int)
		ret.StaplingResponseSuccess = in["stapling_response_success"].(int)
		ret.StaplingResponseFailure = in["stapling_response_failure"].(int)
		ret.StaplingResponseError = in["stapling_response_error"].(int)
		ret.StaplingResponseTimeout = in["stapling_response_timeout"].(int)
		ret.StaplingResponseOther = in["stapling_response_other"].(int)
		ret.RequestNormal = in["request_normal"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseSuccess = in["response_success"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.ResponseOther = in["response_other"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerOcspStats(d *schema.ResourceData) edpt.AamAuthenticationServerOcspStats {
	var ret edpt.AamAuthenticationServerOcspStats

	ret.InstanceList = getSliceAamAuthenticationServerOcspStatsInstanceList(d.Get("instance_list").([]interface{}))

	ret.Stats = getObjectAamAuthenticationServerOcspStatsStats(d.Get("stats").([]interface{}))
	return ret
}
