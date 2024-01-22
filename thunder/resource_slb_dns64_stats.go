package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDns64Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_dns64_stats`: Statistics for the object dns64\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbDns64StatsRead,

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

func resourceSlbDns64StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDns64StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDns64Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbDns64StatsStats := setObjectSlbDns64StatsStats(res)
		d.Set("stats", SlbDns64StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbDns64StatsStats(ret edpt.DataSlbDns64Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"query":          ret.DtSlbDns64Stats.Stats.Query,
			"query_bad_pkt":  ret.DtSlbDns64Stats.Stats.QueryBadPkt,
			"query_chg":      ret.DtSlbDns64Stats.Stats.QueryChg,
			"query_parallel": ret.DtSlbDns64Stats.Stats.QueryParallel,
			"query_passive":  ret.DtSlbDns64Stats.Stats.QueryPassive,
			"resp":           ret.DtSlbDns64Stats.Stats.Resp,
			"resp_bad_pkt":   ret.DtSlbDns64Stats.Stats.RespBadPkt,
			"resp_bad_qr":    ret.DtSlbDns64Stats.Stats.RespBadQr,
			"resp_chg":       ret.DtSlbDns64Stats.Stats.RespChg,
			"resp_err":       ret.DtSlbDns64Stats.Stats.RespErr,
			"resp_empty":     ret.DtSlbDns64Stats.Stats.RespEmpty,
			"resp_local":     ret.DtSlbDns64Stats.Stats.RespLocal,
			"adjust":         ret.DtSlbDns64Stats.Stats.Adjust,
			"cache":          ret.DtSlbDns64Stats.Stats.Cache,
			"drop":           ret.DtSlbDns64Stats.Stats.Drop,
		},
	}
}

func getObjectSlbDns64StatsStats(d []interface{}) edpt.SlbDns64StatsStats {

	count1 := len(d)
	var ret edpt.SlbDns64StatsStats
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

func dataToEndpointSlbDns64Stats(d *schema.ResourceData) edpt.SlbDns64Stats {
	var ret edpt.SlbDns64Stats

	ret.Stats = getObjectSlbDns64StatsStats(d.Get("stats").([]interface{}))
	return ret
}
