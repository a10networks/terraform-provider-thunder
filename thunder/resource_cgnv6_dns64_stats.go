package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Dns64Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_dns64_stats`: Statistics for the object dns64\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Dns64StatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query": {
							Type: schema.TypeInt, Optional: true, Description: "Query",
						},
						"query_bad_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Query Bad Packet",
						},
						"query_chg": {
							Type: schema.TypeInt, Optional: true, Description: "Query Changed",
						},
						"query_parallel": {
							Type: schema.TypeInt, Optional: true, Description: "Query Parallel",
						},
						"query_passive": {
							Type: schema.TypeInt, Optional: true, Description: "Query Passive",
						},
						"resp": {
							Type: schema.TypeInt, Optional: true, Description: "Response",
						},
						"resp_bad_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Response Bad Packet",
						},
						"resp_bad_qr": {
							Type: schema.TypeInt, Optional: true, Description: "Response Bad Query",
						},
						"resp_chg": {
							Type: schema.TypeInt, Optional: true, Description: "Response Changed",
						},
						"resp_err": {
							Type: schema.TypeInt, Optional: true, Description: "Response Error",
						},
						"resp_empty": {
							Type: schema.TypeInt, Optional: true, Description: "Response Empty",
						},
						"resp_local": {
							Type: schema.TypeInt, Optional: true, Description: "Response Local",
						},
						"adjust": {
							Type: schema.TypeInt, Optional: true, Description: "Translated",
						},
						"cache": {
							Type: schema.TypeInt, Optional: true, Description: "Cache",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dropped",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Dns64StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Dns64StatsStats := setObjectCgnv6Dns64StatsStats(res)
		d.Set("stats", Cgnv6Dns64StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Dns64StatsStats(ret edpt.DataCgnv6Dns64Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"query":          ret.DtCgnv6Dns64Stats.Stats.Query,
			"query_bad_pkt":  ret.DtCgnv6Dns64Stats.Stats.QueryBadPkt,
			"query_chg":      ret.DtCgnv6Dns64Stats.Stats.QueryChg,
			"query_parallel": ret.DtCgnv6Dns64Stats.Stats.QueryParallel,
			"query_passive":  ret.DtCgnv6Dns64Stats.Stats.QueryPassive,
			"resp":           ret.DtCgnv6Dns64Stats.Stats.Resp,
			"resp_bad_pkt":   ret.DtCgnv6Dns64Stats.Stats.RespBadPkt,
			"resp_bad_qr":    ret.DtCgnv6Dns64Stats.Stats.RespBadQr,
			"resp_chg":       ret.DtCgnv6Dns64Stats.Stats.RespChg,
			"resp_err":       ret.DtCgnv6Dns64Stats.Stats.RespErr,
			"resp_empty":     ret.DtCgnv6Dns64Stats.Stats.RespEmpty,
			"resp_local":     ret.DtCgnv6Dns64Stats.Stats.RespLocal,
			"adjust":         ret.DtCgnv6Dns64Stats.Stats.Adjust,
			"cache":          ret.DtCgnv6Dns64Stats.Stats.Cache,
			"drop":           ret.DtCgnv6Dns64Stats.Stats.Drop,
		},
	}
}

func getObjectCgnv6Dns64StatsStats(d []interface{}) edpt.Cgnv6Dns64StatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6Dns64StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Query = in["query"].(int)
		ret.QueryBadPkt = in["query_bad_pkt"].(int)
		ret.QueryChg = in["query_chg"].(int)
		ret.QueryParallel = in["query_parallel"].(int)
		ret.QueryPassive = in["query_passive"].(int)
		ret.Resp = in["resp"].(int)
		ret.RespBadPkt = in["resp_bad_pkt"].(int)
		ret.RespBadQr = in["resp_bad_qr"].(int)
		ret.RespChg = in["resp_chg"].(int)
		ret.RespErr = in["resp_err"].(int)
		ret.RespEmpty = in["resp_empty"].(int)
		ret.RespLocal = in["resp_local"].(int)
		ret.Adjust = in["adjust"].(int)
		ret.Cache = in["cache"].(int)
		ret.Drop = in["drop"].(int)
	}
	return ret
}

func dataToEndpointCgnv6Dns64Stats(d *schema.ResourceData) edpt.Cgnv6Dns64Stats {
	var ret edpt.Cgnv6Dns64Stats

	ret.Stats = getObjectCgnv6Dns64StatsStats(d.Get("stats").([]interface{}))
	return ret
}
