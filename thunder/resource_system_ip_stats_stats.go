package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpStatsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_ip_stats_stats`: Statistics for the object ip-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemIpStatsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inreceives": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming packets received",
						},
						"inhdrerrors": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming packet header errors",
						},
						"intoobigerrors": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming packet too big errors",
						},
						"innoroutes": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming no route packet drops",
						},
						"inaddrerrors": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming packet address errors",
						},
						"inunknownprotos": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming unkown protocol packet drops",
						},
						"intruncatedpkts": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming truncated packets",
						},
						"indiscards": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming packets discarded",
						},
						"indelivers": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming packets delivered",
						},
						"outforwdatagrams": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing forwarded datagrams",
						},
						"outrequests": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing packets",
						},
						"outdiscards": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing packets discarded",
						},
						"outnoroutes": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing no route packet drops",
						},
						"reasmtimeout": {
							Type: schema.TypeInt, Optional: true, Description: "Reassembly timed out packet drops",
						},
						"reasmreqds": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming reassembly requests",
						},
						"reasmoks": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming reassembled packets",
						},
						"reasmfails": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming reassembly requests failed",
						},
						"fragoks": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing packets fragmented",
						},
						"fragfails": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing packets fragmentation failed",
						},
						"fragcreates": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing fragmented packets",
						},
						"inmcastpkts": {
							Type: schema.TypeInt, Optional: true, Description: "Incoming multicast packets",
						},
						"outmcastpkts": {
							Type: schema.TypeInt, Optional: true, Description: "Outgoing multicast packets",
						},
					},
				},
			},
		},
	}
}

func resourceSystemIpStatsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpStatsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpStatsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemIpStatsStatsStats := setObjectSystemIpStatsStatsStats(res)
		d.Set("stats", SystemIpStatsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemIpStatsStatsStats(ret edpt.DataSystemIpStatsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"inreceives":       ret.DtSystemIpStatsStats.Stats.Inreceives,
			"inhdrerrors":      ret.DtSystemIpStatsStats.Stats.Inhdrerrors,
			"intoobigerrors":   ret.DtSystemIpStatsStats.Stats.Intoobigerrors,
			"innoroutes":       ret.DtSystemIpStatsStats.Stats.Innoroutes,
			"inaddrerrors":     ret.DtSystemIpStatsStats.Stats.Inaddrerrors,
			"inunknownprotos":  ret.DtSystemIpStatsStats.Stats.Inunknownprotos,
			"intruncatedpkts":  ret.DtSystemIpStatsStats.Stats.Intruncatedpkts,
			"indiscards":       ret.DtSystemIpStatsStats.Stats.Indiscards,
			"indelivers":       ret.DtSystemIpStatsStats.Stats.Indelivers,
			"outforwdatagrams": ret.DtSystemIpStatsStats.Stats.Outforwdatagrams,
			"outrequests":      ret.DtSystemIpStatsStats.Stats.Outrequests,
			"outdiscards":      ret.DtSystemIpStatsStats.Stats.Outdiscards,
			"outnoroutes":      ret.DtSystemIpStatsStats.Stats.Outnoroutes,
			"reasmtimeout":     ret.DtSystemIpStatsStats.Stats.Reasmtimeout,
			"reasmreqds":       ret.DtSystemIpStatsStats.Stats.Reasmreqds,
			"reasmoks":         ret.DtSystemIpStatsStats.Stats.Reasmoks,
			"reasmfails":       ret.DtSystemIpStatsStats.Stats.Reasmfails,
			"fragoks":          ret.DtSystemIpStatsStats.Stats.Fragoks,
			"fragfails":        ret.DtSystemIpStatsStats.Stats.Fragfails,
			"fragcreates":      ret.DtSystemIpStatsStats.Stats.Fragcreates,
			"inmcastpkts":      ret.DtSystemIpStatsStats.Stats.Inmcastpkts,
			"outmcastpkts":     ret.DtSystemIpStatsStats.Stats.Outmcastpkts,
		},
	}
}

func getObjectSystemIpStatsStatsStats(d []interface{}) edpt.SystemIpStatsStatsStats {

	count1 := len(d)
	var ret edpt.SystemIpStatsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inreceives = in["inreceives"].(int)
		ret.Inhdrerrors = in["inhdrerrors"].(int)
		ret.Intoobigerrors = in["intoobigerrors"].(int)
		ret.Innoroutes = in["innoroutes"].(int)
		ret.Inaddrerrors = in["inaddrerrors"].(int)
		ret.Inunknownprotos = in["inunknownprotos"].(int)
		ret.Intruncatedpkts = in["intruncatedpkts"].(int)
		ret.Indiscards = in["indiscards"].(int)
		ret.Indelivers = in["indelivers"].(int)
		ret.Outforwdatagrams = in["outforwdatagrams"].(int)
		ret.Outrequests = in["outrequests"].(int)
		ret.Outdiscards = in["outdiscards"].(int)
		ret.Outnoroutes = in["outnoroutes"].(int)
		ret.Reasmtimeout = in["reasmtimeout"].(int)
		ret.Reasmreqds = in["reasmreqds"].(int)
		ret.Reasmoks = in["reasmoks"].(int)
		ret.Reasmfails = in["reasmfails"].(int)
		ret.Fragoks = in["fragoks"].(int)
		ret.Fragfails = in["fragfails"].(int)
		ret.Fragcreates = in["fragcreates"].(int)
		ret.Inmcastpkts = in["inmcastpkts"].(int)
		ret.Outmcastpkts = in["outmcastpkts"].(int)
	}
	return ret
}

func dataToEndpointSystemIpStatsStats(d *schema.ResourceData) edpt.SystemIpStatsStats {
	var ret edpt.SystemIpStatsStats

	ret.Stats = getObjectSystemIpStatsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
