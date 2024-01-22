package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgMgcpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_alg_mgcp_stats`: Statistics for the object mgcp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatAlgMgcpStatsRead,

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
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatAlgMgcpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgMgcpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgMgcpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatAlgMgcpStatsStats := setObjectCgnv6FixedNatAlgMgcpStatsStats(res)
		d.Set("stats", Cgnv6FixedNatAlgMgcpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatAlgMgcpStatsStats(ret edpt.DataCgnv6FixedNatAlgMgcpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"auep": ret.DtCgnv6FixedNatAlgMgcpStats.Stats.Auep,
			"aucx": ret.DtCgnv6FixedNatAlgMgcpStats.Stats.Aucx,
			"crcx": ret.DtCgnv6FixedNatAlgMgcpStats.Stats.Crcx,
			"dlcx": ret.DtCgnv6FixedNatAlgMgcpStats.Stats.Dlcx,
			"epcf": ret.DtCgnv6FixedNatAlgMgcpStats.Stats.Epcf,
			"mdcx": ret.DtCgnv6FixedNatAlgMgcpStats.Stats.Mdcx,
			"ntfy": ret.DtCgnv6FixedNatAlgMgcpStats.Stats.Ntfy,
			"rqnt": ret.DtCgnv6FixedNatAlgMgcpStats.Stats.Rqnt,
			"rsip": ret.DtCgnv6FixedNatAlgMgcpStats.Stats.Rsip,
		},
	}
}

func getObjectCgnv6FixedNatAlgMgcpStatsStats(d []interface{}) edpt.Cgnv6FixedNatAlgMgcpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatAlgMgcpStatsStats
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
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgMgcpStats(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgMgcpStats {
	var ret edpt.Cgnv6FixedNatAlgMgcpStats

	ret.Stats = getObjectCgnv6FixedNatAlgMgcpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
