package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgH323Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_alg_h323_stats`: Statistics for the object h323\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnAlgH323StatsRead,

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
						"parse_error": {
							Type: schema.TypeInt, Optional: true, Description: "Message Parse Error",
						},
						"tcp_out_of_order_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Out-of-Order Drop",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnAlgH323StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgH323StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgH323Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnAlgH323StatsStats := setObjectCgnv6LsnAlgH323StatsStats(res)
		d.Set("stats", Cgnv6LsnAlgH323StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnAlgH323StatsStats(ret edpt.DataCgnv6LsnAlgH323Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"h225ras_message":       ret.DtCgnv6LsnAlgH323Stats.Stats.H225rasMessage,
			"h225cs_message":        ret.DtCgnv6LsnAlgH323Stats.Stats.H225csMessage,
			"h245ctl_message":       ret.DtCgnv6LsnAlgH323Stats.Stats.H245ctlMessage,
			"h245_tunneled":         ret.DtCgnv6LsnAlgH323Stats.Stats.H245Tunneled,
			"fast_start":            ret.DtCgnv6LsnAlgH323Stats.Stats.FastStart,
			"parse_error":           ret.DtCgnv6LsnAlgH323Stats.Stats.ParseError,
			"tcp_out_of_order_drop": ret.DtCgnv6LsnAlgH323Stats.Stats.TcpOutOfOrderDrop,
		},
	}
}

func getObjectCgnv6LsnAlgH323StatsStats(d []interface{}) edpt.Cgnv6LsnAlgH323StatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnAlgH323StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.H225rasMessage = in["h225ras_message"].(int)
		ret.H225csMessage = in["h225cs_message"].(int)
		ret.H245ctlMessage = in["h245ctl_message"].(int)
		ret.H245Tunneled = in["h245_tunneled"].(int)
		ret.FastStart = in["fast_start"].(int)
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgH323Stats(d *schema.ResourceData) edpt.Cgnv6LsnAlgH323Stats {
	var ret edpt.Cgnv6LsnAlgH323Stats

	ret.Stats = getObjectCgnv6LsnAlgH323StatsStats(d.Get("stats").([]interface{}))
	return ret
}
