package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgH323Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_alg_h323_stats`: Statistics for the object h323\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatAlgH323StatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"h225ras_message": {
							Type: schema.TypeInt, Optional: true, Description: "H323 H225 RAS Message",
						},
						"h225cs_message": {
							Type: schema.TypeInt, Optional: true, Description: "H323 H225 Call Signaling Message",
						},
						"h245ctl_message": {
							Type: schema.TypeInt, Optional: true, Description: "H323 H245 Media Control Message",
						},
						"h245_tunneled": {
							Type: schema.TypeInt, Optional: true, Description: "H323 H245 Tunnelled Message",
						},
						"fast_start": {
							Type: schema.TypeInt, Optional: true, Description: "H323 FastStart",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatAlgH323StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgH323StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgH323Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatAlgH323StatsStats := setObjectCgnv6FixedNatAlgH323StatsStats(res)
		d.Set("stats", Cgnv6FixedNatAlgH323StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatAlgH323StatsStats(ret edpt.DataCgnv6FixedNatAlgH323Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"h225ras_message": ret.DtCgnv6FixedNatAlgH323Stats.Stats.H225rasMessage,
			"h225cs_message":  ret.DtCgnv6FixedNatAlgH323Stats.Stats.H225csMessage,
			"h245ctl_message": ret.DtCgnv6FixedNatAlgH323Stats.Stats.H245ctlMessage,
			"h245_tunneled":   ret.DtCgnv6FixedNatAlgH323Stats.Stats.H245Tunneled,
			"fast_start":      ret.DtCgnv6FixedNatAlgH323Stats.Stats.FastStart,
		},
	}
}

func getObjectCgnv6FixedNatAlgH323StatsStats(d []interface{}) edpt.Cgnv6FixedNatAlgH323StatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatAlgH323StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.H225rasMessage = in["h225ras_message"].(int)
		ret.H225csMessage = in["h225cs_message"].(int)
		ret.H245ctlMessage = in["h245ctl_message"].(int)
		ret.H245Tunneled = in["h245_tunneled"].(int)
		ret.FastStart = in["fast_start"].(int)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgH323Stats(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgH323Stats {
	var ret edpt.Cgnv6FixedNatAlgH323Stats

	ret.Stats = getObjectCgnv6FixedNatAlgH323StatsStats(d.Get("stats").([]interface{}))
	return ret
}
