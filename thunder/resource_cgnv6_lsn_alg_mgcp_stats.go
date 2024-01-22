package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgMgcpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_alg_mgcp_stats`: Statistics for the object mgcp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnAlgMgcpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auep": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP AUEP",
						},
						"aucx": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP AUCX",
						},
						"crcx": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP CRCX",
						},
						"dlcx": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP DLCX",
						},
						"epcf": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP EPCF",
						},
						"mdcx": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP MDCX",
						},
						"ntfy": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP NTFY",
						},
						"rqnt": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP RQNT",
						},
						"rsip": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP RSIP",
						},
						"parse_error": {
							Type: schema.TypeInt, Optional: true, Description: "MGCP Message Parse Error",
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

func resourceCgnv6LsnAlgMgcpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgMgcpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgMgcpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnAlgMgcpStatsStats := setObjectCgnv6LsnAlgMgcpStatsStats(res)
		d.Set("stats", Cgnv6LsnAlgMgcpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnAlgMgcpStatsStats(ret edpt.DataCgnv6LsnAlgMgcpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"auep":                  ret.DtCgnv6LsnAlgMgcpStats.Stats.Auep,
			"aucx":                  ret.DtCgnv6LsnAlgMgcpStats.Stats.Aucx,
			"crcx":                  ret.DtCgnv6LsnAlgMgcpStats.Stats.Crcx,
			"dlcx":                  ret.DtCgnv6LsnAlgMgcpStats.Stats.Dlcx,
			"epcf":                  ret.DtCgnv6LsnAlgMgcpStats.Stats.Epcf,
			"mdcx":                  ret.DtCgnv6LsnAlgMgcpStats.Stats.Mdcx,
			"ntfy":                  ret.DtCgnv6LsnAlgMgcpStats.Stats.Ntfy,
			"rqnt":                  ret.DtCgnv6LsnAlgMgcpStats.Stats.Rqnt,
			"rsip":                  ret.DtCgnv6LsnAlgMgcpStats.Stats.Rsip,
			"parse_error":           ret.DtCgnv6LsnAlgMgcpStats.Stats.ParseError,
			"tcp_out_of_order_drop": ret.DtCgnv6LsnAlgMgcpStats.Stats.TcpOutOfOrderDrop,
		},
	}
}

func getObjectCgnv6LsnAlgMgcpStatsStats(d []interface{}) edpt.Cgnv6LsnAlgMgcpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnAlgMgcpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Auep = in["auep"].(int)
		ret.Aucx = in["aucx"].(int)
		ret.Crcx = in["crcx"].(int)
		ret.Dlcx = in["dlcx"].(int)
		ret.Epcf = in["epcf"].(int)
		ret.Mdcx = in["mdcx"].(int)
		ret.Ntfy = in["ntfy"].(int)
		ret.Rqnt = in["rqnt"].(int)
		ret.Rsip = in["rsip"].(int)
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgMgcpStats(d *schema.ResourceData) edpt.Cgnv6LsnAlgMgcpStats {
	var ret edpt.Cgnv6LsnAlgMgcpStats

	ret.Stats = getObjectCgnv6LsnAlgMgcpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
