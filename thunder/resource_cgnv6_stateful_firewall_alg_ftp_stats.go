package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgFtpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_stateful_firewall_alg_ftp_stats`: Statistics for the object ftp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6StatefulFirewallAlgFtpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_port_request": {
							Type: schema.TypeInt, Optional: true, Description: "PORT Requests From Client",
						},
						"client_eprt_request": {
							Type: schema.TypeInt, Optional: true, Description: "EPRT Requests From Client",
						},
						"server_pasv_reply": {
							Type: schema.TypeInt, Optional: true, Description: "PASV Replies From Server",
						},
						"server_epsv_reply": {
							Type: schema.TypeInt, Optional: true, Description: "EPSV Replies From Server",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6StatefulFirewallAlgFtpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgFtpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgFtpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6StatefulFirewallAlgFtpStatsStats := setObjectCgnv6StatefulFirewallAlgFtpStatsStats(res)
		d.Set("stats", Cgnv6StatefulFirewallAlgFtpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6StatefulFirewallAlgFtpStatsStats(ret edpt.DataCgnv6StatefulFirewallAlgFtpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"client_port_request": ret.DtCgnv6StatefulFirewallAlgFtpStats.Stats.ClientPortRequest,
			"client_eprt_request": ret.DtCgnv6StatefulFirewallAlgFtpStats.Stats.ClientEprtRequest,
			"server_pasv_reply":   ret.DtCgnv6StatefulFirewallAlgFtpStats.Stats.ServerPasvReply,
			"server_epsv_reply":   ret.DtCgnv6StatefulFirewallAlgFtpStats.Stats.ServerEpsvReply,
		},
	}
}

func getObjectCgnv6StatefulFirewallAlgFtpStatsStats(d []interface{}) edpt.Cgnv6StatefulFirewallAlgFtpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6StatefulFirewallAlgFtpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClientPortRequest = in["client_port_request"].(int)
		ret.ClientEprtRequest = in["client_eprt_request"].(int)
		ret.ServerPasvReply = in["server_pasv_reply"].(int)
		ret.ServerEpsvReply = in["server_epsv_reply"].(int)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallAlgFtpStats(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgFtpStats {
	var ret edpt.Cgnv6StatefulFirewallAlgFtpStats

	ret.Stats = getObjectCgnv6StatefulFirewallAlgFtpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
