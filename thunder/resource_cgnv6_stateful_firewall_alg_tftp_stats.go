package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgTftpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_stateful_firewall_alg_tftp_stats`: Statistics for the object tftp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6StatefulFirewallAlgTftpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_created": {
							Type: schema.TypeInt, Optional: true, Description: "TFTP Client Sessions created",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6StatefulFirewallAlgTftpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgTftpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgTftpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6StatefulFirewallAlgTftpStatsStats := setObjectCgnv6StatefulFirewallAlgTftpStatsStats(res)
		d.Set("stats", Cgnv6StatefulFirewallAlgTftpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6StatefulFirewallAlgTftpStatsStats(ret edpt.DataCgnv6StatefulFirewallAlgTftpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_created": ret.DtCgnv6StatefulFirewallAlgTftpStats.Stats.SessionCreated,
		},
	}
}

func getObjectCgnv6StatefulFirewallAlgTftpStatsStats(d []interface{}) edpt.Cgnv6StatefulFirewallAlgTftpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6StatefulFirewallAlgTftpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionCreated = in["session_created"].(int)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallAlgTftpStats(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgTftpStats {
	var ret edpt.Cgnv6StatefulFirewallAlgTftpStats

	ret.Stats = getObjectCgnv6StatefulFirewallAlgTftpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
