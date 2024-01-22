package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayKerberosInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_relay_kerberos_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationRelayKerberosInstanceStatsRead,

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
	}
}

func resourceAamAuthenticationRelayKerberosInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberosInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationRelayKerberosInstanceStatsStats := setObjectAamAuthenticationRelayKerberosInstanceStatsStats(res)
		d.Set("stats", AamAuthenticationRelayKerberosInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationRelayKerberosInstanceStatsStats(ret edpt.DataAamAuthenticationRelayKerberosInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request_send":             ret.DtAamAuthenticationRelayKerberosInstanceStats.Stats.RequestSend,
			"response_receive":         ret.DtAamAuthenticationRelayKerberosInstanceStats.Stats.ResponseReceive,
			"current_requests_of_user": ret.DtAamAuthenticationRelayKerberosInstanceStats.Stats.CurrentRequestsOfUser,
			"tickets":                  ret.DtAamAuthenticationRelayKerberosInstanceStats.Stats.Tickets,
		},
	}
}

func getObjectAamAuthenticationRelayKerberosInstanceStatsStats(d []interface{}) edpt.AamAuthenticationRelayKerberosInstanceStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayKerberosInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestSend = in["request_send"].(int)
		ret.ResponseReceive = in["response_receive"].(int)
		ret.CurrentRequestsOfUser = in["current_requests_of_user"].(int)
		ret.Tickets = in["tickets"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayKerberosInstanceStats(d *schema.ResourceData) edpt.AamAuthenticationRelayKerberosInstanceStats {
	var ret edpt.AamAuthenticationRelayKerberosInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationRelayKerberosInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
