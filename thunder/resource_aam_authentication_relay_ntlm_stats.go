package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayNtlmStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_relay_ntlm_stats`: Statistics for the object ntlm\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationRelayNtlmStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify NTLM authentication relay name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"success": {
							Type: schema.TypeInt, Optional: true, Description: "Success",
						},
						"failure": {
							Type: schema.TypeInt, Optional: true, Description: "Failure",
						},
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Request",
						},
						"response": {
							Type: schema.TypeInt, Optional: true, Description: "Response",
						},
						"http_code_200": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 200 OK",
						},
						"http_code_400": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 400 Bad Request",
						},
						"http_code_401": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 401 Unauthorized",
						},
						"http_code_403": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 403 Forbidden",
						},
						"http_code_404": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 404 Not Found",
						},
						"http_code_500": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 500 Internal Server Error",
						},
						"http_code_503": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP 503 Service Unavailable",
						},
						"http_code_other": {
							Type: schema.TypeInt, Optional: true, Description: "Other HTTP Response",
						},
						"buffer_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Buffer Allocation Failure",
						},
						"encoding_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Encoding Failure",
						},
						"insert_header_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Insert Header Failure",
						},
						"parse_header_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse Header Failure",
						},
						"internal_error": {
							Type: schema.TypeInt, Optional: true, Description: "Internal Error",
						},
						"ntlm_auth_skipped": {
							Type: schema.TypeInt, Optional: true, Description: "Requests for which NTLM relay is skipped",
						},
						"large_request_processing": {
							Type: schema.TypeInt, Optional: true, Description: "Requests invoking large request processing",
						},
						"large_request_flushed": {
							Type: schema.TypeInt, Optional: true, Description: "Large requests sent to server",
						},
						"head_negotiate_request_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HEAD requests sent with NEGOTIATE header",
						},
						"head_auth_request_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HEAD requests sent with AUTH header",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationRelayNtlmStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayNtlmStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayNtlmStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationRelayNtlmStatsStats := setObjectAamAuthenticationRelayNtlmStatsStats(res)
		d.Set("stats", AamAuthenticationRelayNtlmStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationRelayNtlmStatsStats(ret edpt.DataAamAuthenticationRelayNtlmStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"success":                     ret.DtAamAuthenticationRelayNtlmStats.Stats.Success,
			"failure":                     ret.DtAamAuthenticationRelayNtlmStats.Stats.Failure,
			"request":                     ret.DtAamAuthenticationRelayNtlmStats.Stats.Request,
			"response":                    ret.DtAamAuthenticationRelayNtlmStats.Stats.Response,
			"http_code_200":               ret.DtAamAuthenticationRelayNtlmStats.Stats.HttpCode200,
			"http_code_400":               ret.DtAamAuthenticationRelayNtlmStats.Stats.HttpCode400,
			"http_code_401":               ret.DtAamAuthenticationRelayNtlmStats.Stats.HttpCode401,
			"http_code_403":               ret.DtAamAuthenticationRelayNtlmStats.Stats.HttpCode403,
			"http_code_404":               ret.DtAamAuthenticationRelayNtlmStats.Stats.HttpCode404,
			"http_code_500":               ret.DtAamAuthenticationRelayNtlmStats.Stats.HttpCode500,
			"http_code_503":               ret.DtAamAuthenticationRelayNtlmStats.Stats.HttpCode503,
			"http_code_other":             ret.DtAamAuthenticationRelayNtlmStats.Stats.HttpCodeOther,
			"buffer_alloc_fail":           ret.DtAamAuthenticationRelayNtlmStats.Stats.BufferAllocFail,
			"encoding_fail":               ret.DtAamAuthenticationRelayNtlmStats.Stats.EncodingFail,
			"insert_header_fail":          ret.DtAamAuthenticationRelayNtlmStats.Stats.InsertHeaderFail,
			"parse_header_fail":           ret.DtAamAuthenticationRelayNtlmStats.Stats.ParseHeaderFail,
			"internal_error":              ret.DtAamAuthenticationRelayNtlmStats.Stats.InternalError,
			"ntlm_auth_skipped":           ret.DtAamAuthenticationRelayNtlmStats.Stats.NtlmAuthSkipped,
			"large_request_processing":    ret.DtAamAuthenticationRelayNtlmStats.Stats.LargeRequestProcessing,
			"large_request_flushed":       ret.DtAamAuthenticationRelayNtlmStats.Stats.LargeRequestFlushed,
			"head_negotiate_request_sent": ret.DtAamAuthenticationRelayNtlmStats.Stats.HeadNegotiateRequestSent,
			"head_auth_request_sent":      ret.DtAamAuthenticationRelayNtlmStats.Stats.HeadAuthRequestSent,
		},
	}
}

func getObjectAamAuthenticationRelayNtlmStatsStats(d []interface{}) edpt.AamAuthenticationRelayNtlmStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayNtlmStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Success = in["success"].(int)
		ret.Failure = in["failure"].(int)
		ret.Request = in["request"].(int)
		ret.Response = in["response"].(int)
		ret.HttpCode200 = in["http_code_200"].(int)
		ret.HttpCode400 = in["http_code_400"].(int)
		ret.HttpCode401 = in["http_code_401"].(int)
		ret.HttpCode403 = in["http_code_403"].(int)
		ret.HttpCode404 = in["http_code_404"].(int)
		ret.HttpCode500 = in["http_code_500"].(int)
		ret.HttpCode503 = in["http_code_503"].(int)
		ret.HttpCodeOther = in["http_code_other"].(int)
		ret.BufferAllocFail = in["buffer_alloc_fail"].(int)
		ret.EncodingFail = in["encoding_fail"].(int)
		ret.InsertHeaderFail = in["insert_header_fail"].(int)
		ret.ParseHeaderFail = in["parse_header_fail"].(int)
		ret.InternalError = in["internal_error"].(int)
		ret.NtlmAuthSkipped = in["ntlm_auth_skipped"].(int)
		ret.LargeRequestProcessing = in["large_request_processing"].(int)
		ret.LargeRequestFlushed = in["large_request_flushed"].(int)
		ret.HeadNegotiateRequestSent = in["head_negotiate_request_sent"].(int)
		ret.HeadAuthRequestSent = in["head_auth_request_sent"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayNtlmStats(d *schema.ResourceData) edpt.AamAuthenticationRelayNtlmStats {
	var ret edpt.AamAuthenticationRelayNtlmStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationRelayNtlmStatsStats(d.Get("stats").([]interface{}))
	return ret
}
