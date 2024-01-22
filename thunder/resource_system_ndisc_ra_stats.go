package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemNdiscRaStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_ndisc_ra_stats`: Statistics for the object ndisc-ra\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemNdiscRaStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"good_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Good Router Solicitations (R.S.) Received",
						},
						"periodic_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Periodic Router Advertisements (R.A.) Sent",
						},
						"rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "R.S. Rate Limited",
						},
						"bad_hop_limit": {
							Type: schema.TypeInt, Optional: true, Description: "R.S. Bad Hop Limit",
						},
						"truncated": {
							Type: schema.TypeInt, Optional: true, Description: "R.S. Truncated",
						},
						"bad_icmpv6_csum": {
							Type: schema.TypeInt, Optional: true, Description: "R.S. Bad ICMPv6 Checksum",
						},
						"bad_icmpv6_code": {
							Type: schema.TypeInt, Optional: true, Description: "R.S. Unknown ICMPv6 Code",
						},
						"bad_icmpv6_option": {
							Type: schema.TypeInt, Optional: true, Description: "R.S. Bad ICMPv6 Option",
						},
						"l2_addr_and_unspec": {
							Type: schema.TypeInt, Optional: true, Description: "R.S. Src Link-Layer Option and Unspecified Address",
						},
						"no_free_buffers": {
							Type: schema.TypeInt, Optional: true, Description: "No Free Buffers to send R.A.",
						},
					},
				},
			},
		},
	}
}

func resourceSystemNdiscRaStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemNdiscRaStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemNdiscRaStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemNdiscRaStatsStats := setObjectSystemNdiscRaStatsStats(res)
		d.Set("stats", SystemNdiscRaStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemNdiscRaStatsStats(ret edpt.DataSystemNdiscRaStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"good_recv":          ret.DtSystemNdiscRaStats.Stats.Good_recv,
			"periodic_sent":      ret.DtSystemNdiscRaStats.Stats.Periodic_sent,
			"rate_limit":         ret.DtSystemNdiscRaStats.Stats.Rate_limit,
			"bad_hop_limit":      ret.DtSystemNdiscRaStats.Stats.Bad_hop_limit,
			"truncated":          ret.DtSystemNdiscRaStats.Stats.Truncated,
			"bad_icmpv6_csum":    ret.DtSystemNdiscRaStats.Stats.Bad_icmpv6_csum,
			"bad_icmpv6_code":    ret.DtSystemNdiscRaStats.Stats.Bad_icmpv6_code,
			"bad_icmpv6_option":  ret.DtSystemNdiscRaStats.Stats.Bad_icmpv6_option,
			"l2_addr_and_unspec": ret.DtSystemNdiscRaStats.Stats.L2_addr_and_unspec,
			"no_free_buffers":    ret.DtSystemNdiscRaStats.Stats.No_free_buffers,
		},
	}
}

func getObjectSystemNdiscRaStatsStats(d []interface{}) edpt.SystemNdiscRaStatsStats {

	count1 := len(d)
	var ret edpt.SystemNdiscRaStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Good_recv = in["good_recv"].(int)
		ret.Periodic_sent = in["periodic_sent"].(int)
		ret.Rate_limit = in["rate_limit"].(int)
		ret.Bad_hop_limit = in["bad_hop_limit"].(int)
		ret.Truncated = in["truncated"].(int)
		ret.Bad_icmpv6_csum = in["bad_icmpv6_csum"].(int)
		ret.Bad_icmpv6_code = in["bad_icmpv6_code"].(int)
		ret.Bad_icmpv6_option = in["bad_icmpv6_option"].(int)
		ret.L2_addr_and_unspec = in["l2_addr_and_unspec"].(int)
		ret.No_free_buffers = in["no_free_buffers"].(int)
	}
	return ret
}

func dataToEndpointSystemNdiscRaStats(d *schema.ResourceData) edpt.SystemNdiscRaStats {
	var ret edpt.SystemNdiscRaStats

	ret.Stats = getObjectSystemNdiscRaStatsStats(d.Get("stats").([]interface{}))
	return ret
}
