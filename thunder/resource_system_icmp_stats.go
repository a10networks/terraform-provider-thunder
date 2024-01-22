package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIcmpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_icmp_stats`: Statistics for the object icmp\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemIcmpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num": {
							Type: schema.TypeInt, Optional: true, Description: "Total number",
						},
						"inmsgs": {
							Type: schema.TypeInt, Optional: true, Description: "In Messages",
						},
						"inerrors": {
							Type: schema.TypeInt, Optional: true, Description: "In Errors",
						},
						"indestunreachs": {
							Type: schema.TypeInt, Optional: true, Description: "In Destination Unreachable",
						},
						"intimeexcds": {
							Type: schema.TypeInt, Optional: true, Description: "In TTL Exceeds",
						},
						"inparmprobs": {
							Type: schema.TypeInt, Optional: true, Description: "In Parameter Problem",
						},
						"insrcquenchs": {
							Type: schema.TypeInt, Optional: true, Description: "In Source Quench Error",
						},
						"inredirects": {
							Type: schema.TypeInt, Optional: true, Description: "In Redirects",
						},
						"inechos": {
							Type: schema.TypeInt, Optional: true, Description: "In Echo requests",
						},
						"inechoreps": {
							Type: schema.TypeInt, Optional: true, Description: "In Echo replies",
						},
						"intimestamps": {
							Type: schema.TypeInt, Optional: true, Description: "In Timestamp",
						},
						"intimestampreps": {
							Type: schema.TypeInt, Optional: true, Description: "In Timestamp Rep",
						},
						"inaddrmasks": {
							Type: schema.TypeInt, Optional: true, Description: "In Address Masks",
						},
						"inaddrmaskreps": {
							Type: schema.TypeInt, Optional: true, Description: "In Address Mask Rep",
						},
						"outmsgs": {
							Type: schema.TypeInt, Optional: true, Description: "Out Message",
						},
						"outerrors": {
							Type: schema.TypeInt, Optional: true, Description: "Out Errors",
						},
						"outdestunreachs": {
							Type: schema.TypeInt, Optional: true, Description: "Out Destination Unreachable",
						},
						"outtimeexcds": {
							Type: schema.TypeInt, Optional: true, Description: "Out TTL Exceeds",
						},
						"outparmprobs": {
							Type: schema.TypeInt, Optional: true, Description: "Out Parameter Problem",
						},
						"outsrcquenchs": {
							Type: schema.TypeInt, Optional: true, Description: "Out Source Quench Error",
						},
						"outredirects": {
							Type: schema.TypeInt, Optional: true, Description: "Out Redirects",
						},
						"outechos": {
							Type: schema.TypeInt, Optional: true, Description: "Out Echo Requests",
						},
						"outechoreps": {
							Type: schema.TypeInt, Optional: true, Description: "Out Echo Replies",
						},
						"outtimestamps": {
							Type: schema.TypeInt, Optional: true, Description: "Out Time Stamp",
						},
						"outtimestampreps": {
							Type: schema.TypeInt, Optional: true, Description: "Out Time Stamp Rep",
						},
						"outaddrmasks": {
							Type: schema.TypeInt, Optional: true, Description: "Out Address Mask",
						},
						"outaddrmaskreps": {
							Type: schema.TypeInt, Optional: true, Description: "Out Address Mask Rep",
						},
					},
				},
			},
		},
	}
}

func resourceSystemIcmpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemIcmpStatsStats := setObjectSystemIcmpStatsStats(res)
		d.Set("stats", SystemIcmpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemIcmpStatsStats(ret edpt.DataSystemIcmpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num":              ret.DtSystemIcmpStats.Stats.Num,
			"inmsgs":           ret.DtSystemIcmpStats.Stats.Inmsgs,
			"inerrors":         ret.DtSystemIcmpStats.Stats.Inerrors,
			"indestunreachs":   ret.DtSystemIcmpStats.Stats.Indestunreachs,
			"intimeexcds":      ret.DtSystemIcmpStats.Stats.Intimeexcds,
			"inparmprobs":      ret.DtSystemIcmpStats.Stats.Inparmprobs,
			"insrcquenchs":     ret.DtSystemIcmpStats.Stats.Insrcquenchs,
			"inredirects":      ret.DtSystemIcmpStats.Stats.Inredirects,
			"inechos":          ret.DtSystemIcmpStats.Stats.Inechos,
			"inechoreps":       ret.DtSystemIcmpStats.Stats.Inechoreps,
			"intimestamps":     ret.DtSystemIcmpStats.Stats.Intimestamps,
			"intimestampreps":  ret.DtSystemIcmpStats.Stats.Intimestampreps,
			"inaddrmasks":      ret.DtSystemIcmpStats.Stats.Inaddrmasks,
			"inaddrmaskreps":   ret.DtSystemIcmpStats.Stats.Inaddrmaskreps,
			"outmsgs":          ret.DtSystemIcmpStats.Stats.Outmsgs,
			"outerrors":        ret.DtSystemIcmpStats.Stats.Outerrors,
			"outdestunreachs":  ret.DtSystemIcmpStats.Stats.Outdestunreachs,
			"outtimeexcds":     ret.DtSystemIcmpStats.Stats.Outtimeexcds,
			"outparmprobs":     ret.DtSystemIcmpStats.Stats.Outparmprobs,
			"outsrcquenchs":    ret.DtSystemIcmpStats.Stats.Outsrcquenchs,
			"outredirects":     ret.DtSystemIcmpStats.Stats.Outredirects,
			"outechos":         ret.DtSystemIcmpStats.Stats.Outechos,
			"outechoreps":      ret.DtSystemIcmpStats.Stats.Outechoreps,
			"outtimestamps":    ret.DtSystemIcmpStats.Stats.Outtimestamps,
			"outtimestampreps": ret.DtSystemIcmpStats.Stats.Outtimestampreps,
			"outaddrmasks":     ret.DtSystemIcmpStats.Stats.Outaddrmasks,
			"outaddrmaskreps":  ret.DtSystemIcmpStats.Stats.Outaddrmaskreps,
		},
	}
}

func getObjectSystemIcmpStatsStats(d []interface{}) edpt.SystemIcmpStatsStats {

	count1 := len(d)
	var ret edpt.SystemIcmpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Num = in["num"].(int)
		ret.Inmsgs = in["inmsgs"].(int)
		ret.Inerrors = in["inerrors"].(int)
		ret.Indestunreachs = in["indestunreachs"].(int)
		ret.Intimeexcds = in["intimeexcds"].(int)
		ret.Inparmprobs = in["inparmprobs"].(int)
		ret.Insrcquenchs = in["insrcquenchs"].(int)
		ret.Inredirects = in["inredirects"].(int)
		ret.Inechos = in["inechos"].(int)
		ret.Inechoreps = in["inechoreps"].(int)
		ret.Intimestamps = in["intimestamps"].(int)
		ret.Intimestampreps = in["intimestampreps"].(int)
		ret.Inaddrmasks = in["inaddrmasks"].(int)
		ret.Inaddrmaskreps = in["inaddrmaskreps"].(int)
		ret.Outmsgs = in["outmsgs"].(int)
		ret.Outerrors = in["outerrors"].(int)
		ret.Outdestunreachs = in["outdestunreachs"].(int)
		ret.Outtimeexcds = in["outtimeexcds"].(int)
		ret.Outparmprobs = in["outparmprobs"].(int)
		ret.Outsrcquenchs = in["outsrcquenchs"].(int)
		ret.Outredirects = in["outredirects"].(int)
		ret.Outechos = in["outechos"].(int)
		ret.Outechoreps = in["outechoreps"].(int)
		ret.Outtimestamps = in["outtimestamps"].(int)
		ret.Outtimestampreps = in["outtimestampreps"].(int)
		ret.Outaddrmasks = in["outaddrmasks"].(int)
		ret.Outaddrmaskreps = in["outaddrmaskreps"].(int)
	}
	return ret
}

func dataToEndpointSystemIcmpStats(d *schema.ResourceData) edpt.SystemIcmpStats {
	var ret edpt.SystemIcmpStats

	ret.Stats = getObjectSystemIcmpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
