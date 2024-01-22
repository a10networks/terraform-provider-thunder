package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIp6StatsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_ip6_stats_stats`: Statistics for the object ip6-stats\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemIp6StatsStatsRead,

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

func resourceSystemIp6StatsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIp6StatsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIp6StatsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemIp6StatsStatsStats := setObjectSystemIp6StatsStatsStats(res)
		d.Set("stats", SystemIp6StatsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemIp6StatsStatsStats(ret edpt.DataSystemIp6StatsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"inreceives":       ret.DtSystemIp6StatsStats.Stats.Inreceives,
			"inhdrerrors":      ret.DtSystemIp6StatsStats.Stats.Inhdrerrors,
			"intoobigerrors":   ret.DtSystemIp6StatsStats.Stats.Intoobigerrors,
			"innoroutes":       ret.DtSystemIp6StatsStats.Stats.Innoroutes,
			"inaddrerrors":     ret.DtSystemIp6StatsStats.Stats.Inaddrerrors,
			"inunknownprotos":  ret.DtSystemIp6StatsStats.Stats.Inunknownprotos,
			"intruncatedpkts":  ret.DtSystemIp6StatsStats.Stats.Intruncatedpkts,
			"indiscards":       ret.DtSystemIp6StatsStats.Stats.Indiscards,
			"indelivers":       ret.DtSystemIp6StatsStats.Stats.Indelivers,
			"outforwdatagrams": ret.DtSystemIp6StatsStats.Stats.Outforwdatagrams,
			"outrequests":      ret.DtSystemIp6StatsStats.Stats.Outrequests,
			"outdiscards":      ret.DtSystemIp6StatsStats.Stats.Outdiscards,
			"outnoroutes":      ret.DtSystemIp6StatsStats.Stats.Outnoroutes,
			"reasmtimeout":     ret.DtSystemIp6StatsStats.Stats.Reasmtimeout,
			"reasmreqds":       ret.DtSystemIp6StatsStats.Stats.Reasmreqds,
			"reasmoks":         ret.DtSystemIp6StatsStats.Stats.Reasmoks,
			"reasmfails":       ret.DtSystemIp6StatsStats.Stats.Reasmfails,
			"fragoks":          ret.DtSystemIp6StatsStats.Stats.Fragoks,
			"fragfails":        ret.DtSystemIp6StatsStats.Stats.Fragfails,
			"fragcreates":      ret.DtSystemIp6StatsStats.Stats.Fragcreates,
			"inmcastpkts":      ret.DtSystemIp6StatsStats.Stats.Inmcastpkts,
			"outmcastpkts":     ret.DtSystemIp6StatsStats.Stats.Outmcastpkts,
		},
	}
}

func getObjectSystemIp6StatsStatsStats(d []interface{}) edpt.SystemIp6StatsStatsStats {

	count1 := len(d)
	var ret edpt.SystemIp6StatsStatsStats
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

func dataToEndpointSystemIp6StatsStats(d *schema.ResourceData) edpt.SystemIp6StatsStats {
	var ret edpt.SystemIp6StatsStats

	ret.Stats = getObjectSystemIp6StatsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
