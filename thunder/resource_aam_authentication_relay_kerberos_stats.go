package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayKerberosStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_relay_kerberos_stats`: Statistics for the object kerberos\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationRelayKerberosStatsRead,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify Kerberos authentication relay name",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"request_send": {
										Type: schema.TypeInt, Optional: true, Description: "Request Send",
									},
									"response_receive": {
										Type: schema.TypeInt, Optional: true, Description: "Response Receive",
									},
									"current_requests_of_user": {
										Type: schema.TypeInt, Optional: true, Description: "Current Pending Requests of User",
									},
									"tickets": {
										Type: schema.TypeInt, Optional: true, Description: "Tickets",
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
						"request_send": {
							Type: schema.TypeInt, Optional: true, Description: "Total Request Send",
						},
						"response_get": {
							Type: schema.TypeInt, Optional: true, Description: "Total Response Get",
						},
						"timeout_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Timeout",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Other Error",
						},
						"request_normal": {
							Type: schema.TypeInt, Optional: true, Description: "Total Normal Request",
						},
						"request_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Total Dropped Request",
						},
						"response_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Success Response",
						},
						"response_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Failure Response",
						},
						"response_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Error Response",
						},
						"response_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Total Timeout Response",
						},
						"response_other": {
							Type: schema.TypeInt, Optional: true, Description: "Total Other Response",
						},
						"job_start_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Job Start Error",
						},
						"polling_control_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Polling Control Error",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationRelayKerberosStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberosStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationRelayKerberosStatsInstanceList := setSliceAamAuthenticationRelayKerberosStatsInstanceList(res)
		d.Set("instance_list", AamAuthenticationRelayKerberosStatsInstanceList)
		AamAuthenticationRelayKerberosStatsStats := setObjectAamAuthenticationRelayKerberosStatsStats(res)
		d.Set("stats", AamAuthenticationRelayKerberosStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceAamAuthenticationRelayKerberosStatsInstanceList(d edpt.DataAamAuthenticationRelayKerberosStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtAamAuthenticationRelayKerberosStats.InstanceList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectAamAuthenticationRelayKerberosStatsInstanceListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectAamAuthenticationRelayKerberosStatsInstanceListStats(d edpt.AamAuthenticationRelayKerberosStatsInstanceListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["request_send"] = d.RequestSend

	in["response_receive"] = d.ResponseReceive

	in["current_requests_of_user"] = d.CurrentRequestsOfUser

	in["tickets"] = d.Tickets
	result = append(result, in)
	return result
}

func setObjectAamAuthenticationRelayKerberosStatsStats(ret edpt.DataAamAuthenticationRelayKerberosStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request_send":          ret.DtAamAuthenticationRelayKerberosStats.Stats.RequestSend,
			"response_get":          ret.DtAamAuthenticationRelayKerberosStats.Stats.ResponseGet,
			"timeout_error":         ret.DtAamAuthenticationRelayKerberosStats.Stats.TimeoutError,
			"other_error":           ret.DtAamAuthenticationRelayKerberosStats.Stats.OtherError,
			"request_normal":        ret.DtAamAuthenticationRelayKerberosStats.Stats.RequestNormal,
			"request_dropped":       ret.DtAamAuthenticationRelayKerberosStats.Stats.RequestDropped,
			"response_success":      ret.DtAamAuthenticationRelayKerberosStats.Stats.ResponseSuccess,
			"response_failure":      ret.DtAamAuthenticationRelayKerberosStats.Stats.ResponseFailure,
			"response_error":        ret.DtAamAuthenticationRelayKerberosStats.Stats.ResponseError,
			"response_timeout":      ret.DtAamAuthenticationRelayKerberosStats.Stats.ResponseTimeout,
			"response_other":        ret.DtAamAuthenticationRelayKerberosStats.Stats.ResponseOther,
			"job_start_error":       ret.DtAamAuthenticationRelayKerberosStats.Stats.JobStartError,
			"polling_control_error": ret.DtAamAuthenticationRelayKerberosStats.Stats.PollingControlError,
		},
	}
}

func getSliceAamAuthenticationRelayKerberosStatsInstanceList(d []interface{}) []edpt.AamAuthenticationRelayKerberosStatsInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayKerberosStatsInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayKerberosStatsInstanceList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectAamAuthenticationRelayKerberosStatsInstanceListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationRelayKerberosStatsInstanceListStats(d []interface{}) edpt.AamAuthenticationRelayKerberosStatsInstanceListStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayKerberosStatsInstanceListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestSend = in["request_send"].(int)
		ret.ResponseReceive = in["response_receive"].(int)
		ret.CurrentRequestsOfUser = in["current_requests_of_user"].(int)
		ret.Tickets = in["tickets"].(int)
	}
	return ret
}

func getObjectAamAuthenticationRelayKerberosStatsStats(d []interface{}) edpt.AamAuthenticationRelayKerberosStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayKerberosStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestSend = in["request_send"].(int)
		ret.ResponseGet = in["response_get"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
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

func dataToEndpointAamAuthenticationRelayKerberosStats(d *schema.ResourceData) edpt.AamAuthenticationRelayKerberosStats {
	var ret edpt.AamAuthenticationRelayKerberosStats

	ret.InstanceList = getSliceAamAuthenticationRelayKerberosStatsInstanceList(d.Get("instance_list").([]interface{}))

	ret.Stats = getObjectAamAuthenticationRelayKerberosStatsStats(d.Get("stats").([]interface{}))
	return ret
}
